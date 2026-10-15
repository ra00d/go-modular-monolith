package storage

import (
	"context"
	"fmt"
	"log"

	"modular-monolith/lib/config"
	"modular-monolith/lib/storage/contract"
	"modular-monolith/lib/storage/local"
	s3prov "modular-monolith/lib/storage/s3"

	"github.com/gogf/gf/v2/os/glog"
	"go.uber.org/fx"
)

// Module wires the storage library into the fx container. It depends on
// *config.Config (always present via lib/app) and exposes:
//
//   - storage.Service     high-level façade with validation + key generation
//   - storage.Provider    the underlying backend, also injectable for tests
//
// Module does NOT register HTTP routes — handlers consume Service from
// their own controllers.
var Module = fx.Module("storage",
	fx.Provide(NewProvider),
	fx.Provide(NewServiceFromAppParams),
)

// NewProvider selects the configured backend and returns a ready
// Provider. The chosen implementation never leaks into application code:
// callers see only the storage.Provider interface.
func NewProvider(cfg *config.Config) (Provider, error) {
	lc := normaliseConfig(cfg.Storage)
	switch lc.Provider {
	case ProviderLocal:
		if lc.Local == nil {
			return nil, fmt.Errorf("%w: local section missing", contract.ErrConfigError)
		}
		return local.NewProvider(*lc.Local)
	case ProviderS3:
		if lc.S3 == nil {
			return nil, fmt.Errorf("%w: s3 section missing", contract.ErrConfigError)
		}
		ctx := context.Background()
		return s3prov.NewProvider(ctx, *lc.S3)
	default:
		return nil, fmt.Errorf("%w: unknown provider %q", contract.ErrConfigError, lc.Provider)
	}
}

// NewServiceFromAppParams wires a high-level service around the
// configured Provider. It also pulls the glog.Logger so storage
// operations can log timing and size without leaking credentials.
func NewServiceFromAppParams(p Provider, cfg *config.Config) *Service {
	return NewService(normaliseConfig(cfg.Storage), p, glog.New())
}

// normaliseConfig builds the library's Config from the project's
// typed configuration. It is also the place to apply sensible defaults
// so callers don't have to spell every option out.
func normaliseConfig(c config.StorageConfig) Config {
	out := Config{
		Provider:    c.Default,
		MaxSize:     c.MaxFileSize,
		KeyPrefix:   c.KeyPrefix,
		AllowedMime: c.AllowedMime,
		AllowedExt:  c.AllowedExt,
		Overwrite:   c.Overwrite,
		DetectMime:  c.DetectMime,
		Local:       c.Local,
		S3:          c.S3,
	}
	if out.Provider == "" {
		out.Provider = ProviderLocal
	}
	return out
}

// MustNewProvider is the imperative constructor used by tests and
// tooling that don't go through fx. It fatals on configuration errors
// because misconfiguration at startup is always a programming error.
func MustNewProvider(cfg *config.Config) Provider {
	p, err := NewProvider(cfg)
	if err != nil {
		log.Fatalf("storage: provider: %v", err)
	}
	return p
}

// MustNewService builds both the Provider and the Service in one call.
func MustNewService(cfg *config.Config) *Service {
	return NewServiceFromAppParams(MustNewProvider(cfg), cfg)
}

// Compile-time assertions live here (instead of in the providers) so the
// provider packages can implement the interface without importing the
// top-level storage package (which would create an import cycle).
var (
	_ Provider = (*local.Provider)(nil)
	_ Provider = (*s3prov.Provider)(nil)
)