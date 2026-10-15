package storage

import (
	"time"

	"github.com/google/uuid"
	"modular-monolith/lib/storage/contract"
)

// ProviderName constants re-exported for caller convenience.
const (
	ProviderLocal = contract.ProviderLocal
	ProviderS3    = contract.ProviderS3
)

// Provider is the low-level byte-storage interface every backend
// implements. It is a type alias for contract.Provider so application
// code can refer to storage.Provider without reaching into the contract
// sub-package.
type Provider = contract.Provider

// PutResult, Download and Stat are re-exported aliases for the contract
// types so callers see a single import.
type (
	PutResult = contract.PutResult
	Download  = contract.Download
	Stat      = contract.Stat
)

// ValidateKey enforces the storage key safety rules. Re-exported from
// the contract package.
func ValidateKey(key string) error { return contract.ValidateKey(key) }

// Object is the provider-independent file metadata model. It deliberately
// only carries information relevant to byte storage (path, size, content
// type, provider-specific identifiers). Business attributes such as
// ownership or visibility belong in the application's domain model.
type Object struct {
	ID          uuid.UUID         `json:"id"`
	Key         string            `json:"key"`
	Bucket      string            `json:"bucket,omitempty"`
	Name        string            `json:"name"`
	Original    string            `json:"original,omitempty"`
	Extension   string            `json:"extension,omitempty"`
	Size        int64             `json:"size"`
	ContentType string            `json:"content_type"`
	Checksum    string            `json:"checksum,omitempty"`
	ETag        string            `json:"etag,omitempty"`
	Provider    string            `json:"provider"`
	Metadata    map[string]string `json:"metadata,omitempty"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}