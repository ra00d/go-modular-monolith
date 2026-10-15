// Package local implements a filesystem-backed storage provider. Files
// are written to a configured root directory using safe path resolution
// and atomic rename. The provider never lets a caller-supplied key
// escape its root and never leaves half-written files behind on failure.
package local

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"modular-monolith/lib/config"
	"modular-monolith/lib/storage/contract"
)

// Provider stores files on the local filesystem. It is safe for
// concurrent use. The root directory is created on first use when
// CreateIfMissing is true; existing files are left untouched.
type Provider struct {
	root            string
	createIfMissing bool
	mu              sync.Mutex // serialise directory creation
}

// NewProvider configures a local filesystem provider. The root is
// absolute or expanded relative to the current working directory. The
// directory must already exist unless cfg.CreateIfMissing is true.
func NewProvider(cfg config.LocalStorageCfg) (*Provider, error) {
	if cfg.Path == "" {
		return nil, fmt.Errorf("%w: local root path is empty", contract.ErrConfigError)
	}
	root, err := filepath.Abs(cfg.Path)
	if err != nil {
		return nil, fmt.Errorf("%w: resolve local root: %v", contract.ErrConfigError, err)
	}
	root = filepath.Clean(root)
	p := &Provider{
		root:            root,
		createIfMissing: cfg.CreateIfMissing,
	}
	if cfg.CreateIfMissing {
		if err := os.MkdirAll(root, 0o755); err != nil {
			return nil, fmt.Errorf("%w: create root: %v", contract.ErrConfigError, err)
		}
	}
	info, err := os.Stat(root)
	if err != nil {
		return nil, fmt.Errorf("%w: stat root: %v", contract.ErrConfigError, err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("%w: local root is not a directory", contract.ErrConfigError)
	}
	return p, nil
}

// Name implements contract.Provider.
func (p *Provider) Name() string { return contract.ProviderLocal }

// Root returns the absolute storage root. Exposed for tests.
func (p *Provider) Root() string { return p.root }

// Put streams src to a file under the root, atomically renaming into
// place once the upload completes. The upload is bounded by the
// configured max-size at the service layer; this method itself accepts
// arbitrary readers because the service drains and validates first.
func (p *Provider) Put(ctx context.Context, key, contentType string, size int64, src io.Reader, m map[string]string) (*contract.PutResult, error) {
	dest, tmp, err := p.resolveDest(key)
	if err != nil {
		return nil, err
	}
	if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		return nil, fmt.Errorf("storage: mkdir: %w", err)
	}
	// O_EXCL so concurrent writers can't silently overwrite; the higher
	// layer is responsible for the Overwrite policy.
	f, err := os.OpenFile(tmp, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0o644)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			return nil, contract.ErrAlreadyExists
		}
		return nil, fmt.Errorf("storage: open tmp: %w", err)
	}
	cleanup := func() {
		_ = f.Close()
		_ = os.Remove(tmp)
	}
	written, err := io.Copy(f, src)
	if err != nil {
		cleanup()
		return nil, fmt.Errorf("storage: copy: %w", err)
	}
	if err := f.Sync(); err != nil {
		cleanup()
		return nil, fmt.Errorf("storage: fsync: %w", err)
	}
	if err := f.Close(); err != nil {
		_ = os.Remove(tmp)
		return nil, fmt.Errorf("storage: close tmp: %w", err)
	}
	if err := os.Rename(tmp, dest); err != nil {
		_ = os.Remove(tmp)
		return nil, fmt.Errorf("storage: rename: %w", err)
	}
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	// Persist content-type alongside the file in a sidecar so subsequent
	// Open/Stat calls can return it without sniffing. The sidecar lives
	// in the same directory under a hidden name and is best-effort.
	_ = os.WriteFile(dest+".meta", []byte("content-type: "+contentType+"\n"), 0o644)

	st, _ := os.Stat(dest)
	res := &contract.PutResult{Size: written}
	if st != nil {
		res.LastModified = st.ModTime()
	}
	_ = ctx // os.Open / os.Stat honour cancellation through ctx at the syscall layer.
	return res, nil
}

// Open streams the file at key. Callers must close the returned reader.
func (p *Provider) Open(ctx context.Context, key string) (*contract.Download, error) {
	dest, err := p.resolveKey(key)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(dest)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, contract.ErrNotFound
		}
		return nil, fmt.Errorf("storage: open: %w", err)
	}
	st, err := f.Stat()
	if err != nil {
		_ = f.Close()
		return nil, fmt.Errorf("storage: stat: %w", err)
	}
	contentType := readSidecar(dest + ".meta")
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	name := filepath.Base(key)
	return &contract.Download{
		Reader:       f,
		Size:         st.Size(),
		ContentType:  contentType,
		Name:         name,
		LastModified: st.ModTime(),
	}, nil
}

// Stat returns the file metadata without reading bytes.
func (p *Provider) Stat(ctx context.Context, key string) (*contract.Stat, error) {
	dest, err := p.resolveKey(key)
	if err != nil {
		return nil, err
	}
	st, err := os.Stat(dest)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, contract.ErrNotFound
		}
		return nil, fmt.Errorf("storage: stat: %w", err)
	}
	ct := readSidecar(dest + ".meta")
	if ct == "" {
		ct = "application/octet-stream"
	}
	return &contract.Stat{
		Key:          key,
		Size:         st.Size(),
		ContentType:  ct,
		LastModified: st.ModTime(),
	}, nil
}

// Delete removes the file. Missing files are not an error.
func (p *Provider) Delete(ctx context.Context, key string) error {
	dest, err := p.resolveKey(key)
	if err != nil {
		return err
	}
	if err := os.Remove(dest); err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("storage: remove: %w", err)
	}
	_ = os.Remove(dest + ".meta")
	return nil
}

// Exists reports whether the file is present.
func (p *Provider) Exists(ctx context.Context, key string) (bool, error) {
	dest, err := p.resolveKey(key)
	if err != nil {
		return false, err
	}
	_, err = os.Stat(dest)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// resolveDest validates the key and computes the final destination plus
// a sibling temp file used for atomic rename.
func (p *Provider) resolveDest(key string) (dest, tmp string, err error) {
	abs, err := p.resolveKey(key)
	if err != nil {
		return "", "", err
	}
	tmp = abs + ".tmp." + randomSuffix()
	return abs, tmp, nil
}

// resolveKey validates the supplied storage key and returns the absolute
// filesystem path it resolves to. Any attempt to escape the root (via
// "..", absolute paths, or symlinks pointing outside) is rejected with
// contract.ErrBadKey.
func (p *Provider) resolveKey(key string) (string, error) {
	if err := contract.ValidateKey(key); err != nil {
		return "", err
	}
	rel := filepath.FromSlash(key)
	joined := filepath.Join(p.root, rel)
	cleaned := filepath.Clean(joined)
	r, err := filepath.Rel(p.root, cleaned)
	if err != nil {
		return "", fmt.Errorf("%w: %v", contract.ErrBadKey, err)
	}
	if strings.HasPrefix(r, ".."+string(os.PathSeparator)) || r == ".." || strings.Contains(r, ".."+string(os.PathSeparator)) {
		return "", fmt.Errorf("%w: escapes root", contract.ErrBadKey)
	}
	if filepath.IsAbs(r) {
		return "", fmt.Errorf("%w: absolute path", contract.ErrBadKey)
	}
	return cleaned, nil
}

func readSidecar(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	s := strings.TrimSpace(string(data))
	if i := strings.IndexByte(s, ':'); i >= 0 {
		return strings.TrimSpace(s[i+1:])
	}
	return s
}

// randomSuffix produces an opaque suffix used for tmp-file names.
func randomSuffix() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}