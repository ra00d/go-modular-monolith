// Package contract holds the type definitions shared between the
// storage façade (lib/storage) and its provider implementations
// (lib/storage/local, lib/storage/s3). It exists as a separate package
// so providers can import the interface without creating an import cycle
// with the top-level storage package.
package contract

import (
	"context"
	"errors"
	"io"
	"time"
)

// ProviderName constants — recognised by every backend.
const (
	ProviderLocal = "local"
	ProviderS3    = "s3"
)

// Sentinel errors shared with the top-level storage package. The
// storage façade re-exports these under its own names so callers see
// only storage.ErrXxx.
var (
	ErrInvalidFile      = errors.New("storage: invalid file")
	ErrFileTooLarge     = errors.New("storage: file too large")
	ErrUnsupportedType  = errors.New("storage: unsupported file type")
	ErrNotFound         = errors.New("storage: not found")
	ErrAlreadyExists    = errors.New("storage: already exists")
	ErrUnavailable      = errors.New("storage: backend unavailable")
	ErrConfigError      = errors.New("storage: invalid configuration")
	ErrPermissionDenied = errors.New("storage: permission denied")
	ErrValidationFailed = errors.New("storage: validation failed")
	ErrBadKey           = errors.New("storage: invalid key")
)

// PutResult is the metadata returned by a successful Put. Fields the
// provider does not know about are left at their zero value.
type PutResult struct {
	ETag         string
	LastModified time.Time
	Size         int64
}

// Download is the streaming result of an Open call.
type Download struct {
	Reader       io.ReadCloser
	Size         int64
	ContentType  string
	Name         string
	ETag         string
	LastModified time.Time
}

// Stat is the metadata-only result of a Head/Stat call.
type Stat struct {
	Key          string
	Size         int64
	ContentType  string
	ETag         string
	LastModified time.Time
	Metadata     map[string]string
}

// Provider is the contract every backend implements.
type Provider interface {
	Name() string
	Put(ctx context.Context, key, contentType string, size int64, src io.Reader, metadata map[string]string) (*PutResult, error)
	Open(ctx context.Context, key string) (*Download, error)
	Stat(ctx context.Context, key string) (*Stat, error)
	Delete(ctx context.Context, key string) error
	Exists(ctx context.Context, key string) (bool, error)
}

// ValidateKey enforces the safety rules for a storage key. It is
// duplicated from the top-level storage package so providers can use it
// without creating an import cycle.
func ValidateKey(key string) error {
	if key == "" {
		return ErrBadKey
	}
	for _, r := range key {
		if r == 0 {
			return ErrBadKey
		}
	}
	// Reject any segment that is "." or ".." or starts with an absolute
	// path separator. "/" and backslash are not allowed together.
	if key[0] == '/' {
		return ErrBadKey
	}
	// Scan for ".." segments.
	for i := 0; i < len(key); i++ {
		if key[i] == '\\' {
			return ErrBadKey
		}
	}
	// Split on "/" and validate each segment.
	start := 0
	for i := 0; i <= len(key); i++ {
		if i == len(key) || key[i] == '/' {
			seg := key[start:i]
			if seg == "" {
				return ErrBadKey
			}
			if seg == "." || seg == ".." {
				return ErrBadKey
			}
			start = i + 1
		}
	}
	return nil
}