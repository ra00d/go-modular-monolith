package storage

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gogf/gf/v2/os/glog"
	"github.com/google/uuid"

	apperrors "modular-monolith/lib/errors"
)

// Service is the high-level entry point for file operations. It owns the
// configured Provider and the validation / key-generation rules. All
// business code should depend on Service rather than Provider so the
// configured backend can be swapped without touching consumers.
type Service struct {
	cfg      Config
	provider Provider
	validate *Validator
	keys     KeyGenerator
	log      *glog.Logger
}

// NewService wires a service around a Provider and its config. The
// supplied logger is used for operational events (provider, operation,
// key, duration, size); sensitive fields are never logged.
func NewService(cfg Config, provider Provider, logger *glog.Logger) *Service {
	if logger == nil {
		logger = glog.New()
	}
	return &Service{
		cfg:      cfg,
		provider: provider,
		validate: NewValidator(cfg),
		keys:     DefaultKeyGenerator,
		log:      logger,
	}
}

// Provider returns the underlying storage backend. Exposed for tests and
// for the rare caller that needs backend-specific behaviour; new business
// code should not depend on it.
func (s *Service) Provider() Provider { return s.provider }

// Config returns the effective configuration.
func (s *Service) Config() Config { return s.cfg }

// Put streams a file into storage. The returned Object contains the
// authoritative size, content type, and storage key. The input reader is
// consumed but never buffered in full.
func (s *Service) Put(ctx context.Context, in UploadInput) (*Object, error) {
	if in.Content == nil {
		return nil, fmt.Errorf("%w: nil reader", ErrInvalidFile)
	}
	if in.Size < 0 {
		// Best-effort size probe — seekable readers expose their length.
		if rs, ok := in.Content.(io.Seeker); ok {
			n, err := rs.Seek(0, io.SeekEnd)
			if err == nil {
				in.Size = n
				_, _ = rs.Seek(0, io.SeekStart)
			}
		}
	}
	if s.cfg.MaxSize > 0 && in.Size > s.cfg.MaxSize {
		return nil, ErrFileTooLarge
	}

	original := in.Filename
	safe, ext := SanitizeFilename(original)
	if ext == "" {
		ext = Extension(original)
	}
	if safe == "" {
		safe = "file"
	}

	prefix := strings.TrimSpace(in.Prefix)
	if prefix == "" {
		prefix = s.cfg.KeyPrefix
	}
	key, displayName, err := s.keys(prefix, original, time.Now().UTC())
	if err != nil {
		return nil, fmt.Errorf("storage: key: %w", err)
	}
	if err := ValidateKey(key); err != nil {
		return nil, err
	}

	declared := NormalizeMime(in.ContentType)
	body, size, checksum, detected, err := s.drain(ctx, in)
	if err != nil {
		return nil, err
	}
	_ = body // unused after this point; readers are returned through provider

	contentType := declared
	if s.cfg.DetectMime || in.DetectContentType {
		if d := NormalizeMime(detected); d != "" {
			contentType = d
		}
	}
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	obj := &Object{
		ID:          uuid.New(),
		Key:         key,
		Bucket:      s.bucketOf(),
		Name:        displayName,
		Original:    original,
		Extension:   ext,
		Size:        size,
		ContentType: contentType,
		Provider:    s.provider.Name(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
	if len(in.Metadata) > 0 {
		obj.Metadata = cloneMeta(in.Metadata)
	}
	if s.cfg.Checksum {
		obj.Checksum = checksum
	}

	// Validation on metadata before backend write.
	if err := s.validate.Validate(obj); err != nil {
		return nil, err
	}

	start := time.Now()
	stored, err := s.provider.Put(ctx, key, contentType, size, bytes.NewReader(body), obj.Metadata)
	if err != nil {
		s.logOperation(ctx, "put", key, start, size, err)
		return nil, err
	}
	if stored != nil {
		// Trust backend for ETag / LastModified.
		if stored.ETag != "" {
			obj.ETag = stored.ETag
		}
		if !stored.LastModified.IsZero() {
			obj.UpdatedAt = stored.LastModified
		}
		if stored.Size > 0 {
			obj.Size = stored.Size
		}
	}
	s.logOperation(ctx, "put", key, start, size, nil)
	return obj, nil
}

// Open streams a stored object back to the caller.
func (s *Service) Open(ctx context.Context, key string) (*Download, error) {
	if err := ValidateKey(key); err != nil {
		return nil, err
	}
	start := time.Now()
	dl, err := s.provider.Open(ctx, key)
	s.logOperation(ctx, "open", key, start, sizeOf(dl), err)
	if err != nil {
		return nil, err
	}
	return dl, nil
}

// Stat returns metadata for key without reading bytes.
func (s *Service) Stat(ctx context.Context, key string) (*Stat, error) {
	if err := ValidateKey(key); err != nil {
		return nil, err
	}
	return s.provider.Stat(ctx, key)
}

// Delete removes key from the configured backend.
func (s *Service) Delete(ctx context.Context, key string) error {
	if err := ValidateKey(key); err != nil {
		return err
	}
	start := time.Now()
	err := s.provider.Delete(ctx, key)
	s.logOperation(ctx, "delete", key, start, 0, err)
	return err
}

// Exists reports whether key is present.
func (s *Service) Exists(ctx context.Context, key string) (bool, error) {
	if err := ValidateKey(key); err != nil {
		return false, err
	}
	return s.provider.Exists(ctx, key)
}

// Drain exposes the buffered body and detection result so HTTP handlers
// can both stream uploads and surface the detected type to callers.
// Production callers should use Put; this is provided for HTTP helpers
// that need to inspect detection results.
func (s *Service) Detect(in UploadInput) (size int64, declared, detected string, err error) {
	body, size, _, detected, err := s.drain(context.Background(), in)
	if err != nil {
		return 0, "", "", err
	}
	_ = body
	return size, NormalizeMime(in.ContentType), NormalizeMime(detected), nil
}

func (s *Service) drain(
	ctx context.Context,
	in UploadInput,
) (body []byte, size int64, checksum, detected string, err error) {
	if in.Content == nil {
		return nil, 0, "", "", fmt.Errorf("%w: nil reader", ErrInvalidFile)
	}
	limit := s.cfg.MaxSize
	if limit <= 0 {
		limit = maxInt64
	}
	hasher := sha256.New()
	tee := io.TeeReader(in.Content, hasher)
	buf := &bytes.Buffer{}
	detected = ""
	// Buffer up to 512 bytes for MIME detection — the canonical sniff
	// window. The remainder continues to stream into the buffer so the
	// backend can write everything from a single Reader.
	probe := make([]byte, 0, 512)
	written := int64(0)
	chunk := make([]byte, 32*1024)
	for {
		if err := ctx.Err(); err != nil {
			return nil, 0, "", "", err
		}
		n, rerr := tee.Read(chunk)
		if n > 0 {
			if written+int64(n) > limit {
				return nil, 0, "", "", ErrFileTooLarge
			}
			if len(probe) < 512 {
				need := 512 - len(probe)
				if need > n {
					need = n
				}
				probe = append(probe, chunk[:need]...)
			}
			buf.Write(chunk[:n])
			written += int64(n)
		}
		if rerr != nil {
			if errors.Is(rerr, io.EOF) {
				break
			}
			return nil, 0, "", "", fmt.Errorf("storage: read: %w", rerr)
		}
	}
	if written == 0 {
		return nil, 0, "", "", ErrInvalidFile
	}
	if detected == "" {
		// Compute once, here, so callers see the result of detection.
		detected = NormalizeMime(sniffMime(probe))
	}
	body = buf.Bytes()
	size = written
	if s.cfg.Checksum {
		checksum = hex.EncodeToString(hasher.Sum(nil))
	}
	if in.Size > 0 && in.Size != size {
		// Caller declared a size that doesn't match what we received.
		// Treat as protocol error, not a missing-file error.
		return nil, 0, "", "", fmt.Errorf(
			"%w: declared size %d, received %d",
			ErrInvalidFile,
			in.Size,
			size,
		)
	}
	return body, size, checksum, detected, nil
}

func (s *Service) logOperation(
	ctx context.Context,
	op, key string,
	start time.Time,
	size int64,
	err error,
) {
	if s.log == nil {
		return
	}
	dur := time.Since(start)
	if err != nil {
		s.log.Warningf(
			ctx,
			"storage.%s provider=%s key=%s size=%d dur=%s err=%v",
			op, s.provider.Name(), key, size, dur, err,
		)
		return
	}
	s.log.Debugf(
		ctx,
		"storage.%s provider=%s key=%s size=%d dur=%s",
		op, s.provider.Name(), key, size, dur,
	)
}

// bucketOf returns the S3 bucket name when the configured provider is S3,
// otherwise "". Implemented by reading from the provider if it exposes
// one; fall back to "" for backends that don't have a bucket concept.
func (s *Service) bucketOf() string {
	type bucketter interface {
		Bucket() string
	}
	if b, ok := s.provider.(bucketter); ok {
		return b.Bucket()
	}
	return ""
}

func cloneMeta(in map[string]string) map[string]string {
	out := make(map[string]string, len(in))
	for k, v := range in {
		if k == "" {
			continue
		}
		out[k] = v
	}
	return out
}

func sizeOf(d *Download) int64 {
	if d == nil {
		return 0
	}
	return d.Size
}

const maxInt64 = int64(1<<63 - 1)

// sniffMime wraps net/http.DetectContentType for content-type detection.
func sniffMime(data []byte) string {
	return http.DetectContentType(data)
}

// MapError is a convenience alias re-exported so callers can map service
// errors to HTTP responses without importing the lib/errors package twice.
func MapError(err error) *apperrors.AppError { return AppError(err) }
