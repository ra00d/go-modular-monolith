package storage

import (
	"net/http"
	"path/filepath"
	"strings"
)

// Validator runs configured checks against a finished upload. It is
// stateless and safe for concurrent use.
type Validator struct {
	maxSize      int64
	allowedMime  []string
	allowedExt   []string
	detectMime   bool
}

// NewValidator constructs a validator from normalised configuration.
// Empty whitelists disable the corresponding check.
func NewValidator(cfg Config) *Validator {
	return &Validator{
		maxSize:     cfg.MaxSize,
		allowedMime: normaliseList(cfg.AllowedMime),
		allowedExt:  normaliseList(cfg.AllowedExt),
		detectMime:  cfg.DetectMime,
	}
}

// Validate enforces the configured rules against a fully-written object.
// It is invoked once the bytes are in storage so size enforcement reflects
// what was actually persisted.
func (v *Validator) Validate(o *Object) error {
	if o == nil {
		return ErrInvalidFile
	}
	if o.Size < 0 {
		return ErrInvalidFile
	}
	if v.maxSize > 0 && o.Size > v.maxSize {
		return ErrFileTooLarge
	}
	if len(v.allowedExt) > 0 && !inSet(v.allowedExt, o.Extension) {
		return ErrUnsupportedType
	}
	if len(v.allowedMime) > 0 && !inSet(v.allowedMime, NormalizeMime(o.ContentType)) {
		return ErrUnsupportedType
	}
	return nil
}

// MergeOptions returns a new Options whose unset fields fall back to the
// validator's configured defaults. Used by Service.Put to apply per-call
// overrides without forcing every caller to repeat defaults.
func (v *Validator) MergeOptions(in Options) Options {
	out := Options{
		Prefix:      in.Prefix,
		MaxSize:     in.MaxSize,
		AllowedMime: in.AllowedMime,
		AllowedExt:  in.AllowedExt,
		Metadata:    in.Metadata,
	}
	if out.MaxSize <= 0 {
		out.MaxSize = v.maxSize
	}
	if len(out.AllowedMime) == 0 {
		out.AllowedMime = v.allowedMime
	}
	if len(out.AllowedExt) == 0 {
		out.AllowedExt = v.allowedExt
	}
	if in.Overwrite != nil {
		out.Overwrite = in.Overwrite
	}
	return out
}

// Extension returns the lowercase extension (no leading dot) for the
// supplied filename, or "" if none. Exported so HTTP integration code can
// pre-fill an Object's Extension before validation.
func Extension(name string) string {
	ext := filepath.Ext(name)
	return strings.ToLower(strings.TrimPrefix(ext, "."))
}

// DetectContentType inspects the first 512 bytes using net/http's
// detector and returns the canonicalised MIME token. Returns
// "application/octet-stream" when detection fails.
func DetectContentType(data []byte) string {
	if len(data) == 0 {
		return "application/octet-stream"
	}
	return NormalizeMime(http.DetectContentType(data))
}

func normaliseList(in []string) []string {
	if len(in) == 0 {
		return nil
	}
	seen := make(map[string]struct{}, len(in))
	out := make([]string, 0, len(in))
	for _, v := range in {
		v = NormalizeMime(v)
		if v == "" {
			continue
		}
		v = strings.ToLower(strings.TrimPrefix(v, "."))
		if v == "" {
			continue
		}
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		out = append(out, v)
	}
	return out
}

func inSet(set []string, v string) bool {
	v = strings.ToLower(v)
	for _, s := range set {
		if s == v {
			return true
		}
	}
	return false
}