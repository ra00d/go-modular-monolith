package storage

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"unicode"

	"modular-monolith/lib/config"
)

// Config is the runtime configuration consumed by the library. It is
// derived from the project's storage configuration in lib/config and
// normalised here so providers and services receive a single shape.
//
// Provider-specific blocks (Local, S3) are kept on the struct so the
// Fx factory can dispatch to the right provider without re-parsing the
// raw configuration.
type Config struct {
	// Provider selects the active backend: "local" or "s3".
	Provider string

	// MaxSize is the upper bound for a single upload. Zero disables the
	// global limit; individual calls may set a stricter bound via Options.
	MaxSize int64

	// KeyPrefix is prepended to every generated storage key.
	KeyPrefix string

	// AllowedMime is the optional whitelist of accepted MIME types. An
	// empty list means "accept any detected type".
	AllowedMime []string

	// AllowedExt is the optional whitelist of accepted extensions. An
	// empty list means "accept any extension".
	AllowedExt []string

	// Overwrite controls behaviour when an upload targets an existing
	// key. When false, the service returns ErrAlreadyExists.
	Overwrite bool

	// DetectMime, when true, forces content-based MIME detection and
	// rejects uploads whose detected type is not in AllowedMime.
	DetectMime bool

	// Checksum enables server-side sha256 checksumming of uploaded bytes.
	Checksum bool

	// Local and S3 are the underlying provider-specific blocks. They are
	// only consumed by NewProvider; the rest of the library ignores them.
	Local *config.LocalStorageCfg
	S3    *config.S3StorageCfg
}

// NormalizeMime returns a lower-cased, parameter-stripped MIME token.
// "text/html; charset=utf-8" becomes "text/html". Returns an empty string
// for blank input.
func NormalizeMime(v string) string {
	v = strings.TrimSpace(strings.ToLower(v))
	if v == "" {
		return ""
	}
	if i := strings.IndexByte(v, ';'); i >= 0 {
		v = v[:i]
	}
	return strings.TrimSpace(v)
}

// SanitizeFilename returns a filesystem- and HTTP-safe filename. It strips
// directory components, control characters, reserved device names, and
// collapses whitespace. The original extension (if any) is preserved and
// returned separately so callers can reconstruct a display name.
func SanitizeFilename(name string) (safe string, ext string) {
	name = strings.TrimSpace(name)
	if name == "" {
		return "", ""
	}
	// Strip any directory components — both / and \ are unsafe.
	if i := strings.LastIndexAny(name, `/\`); i >= 0 {
		name = name[i+1:]
	}
	// Split extension before further sanitisation so we can preserve it.
	ext = ""
	if dot := strings.LastIndexByte(name, '.'); dot > 0 && dot < len(name)-1 {
		ext = strings.ToLower(name[dot+1:])
		name = name[:dot]
	}
	// Replace control characters and forbidden runes.
	var b strings.Builder
	b.Grow(len(name))
	for _, r := range name {
		switch {
		case unicode.IsControl(r):
			b.WriteByte('_')
		case r == 0:
			b.WriteByte('_')
		case r == ' ' || r == '-' || r == '_' || r == '.':
			b.WriteRune(r)
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			b.WriteRune(r)
		default:
			b.WriteByte('_')
		}
	}
	clean := strings.Trim(b.String(), ". _-")
	if clean == "" {
		clean = "file"
	}
	// Reserved Windows device names (case-insensitive).
	lower := strings.ToLower(clean)
	switch lower {
	case "con", "prn", "aux", "nul",
		"com1", "com2", "com3", "com4", "com5", "com6", "com7", "com8", "com9",
		"lpt1", "lpt2", "lpt3", "lpt4", "lpt5", "lpt6", "lpt7", "lpt8", "lpt9":
		clean = "_" + clean
	}
	return clean, ext
}

// HashSHA256 returns the hex-encoded sha256 digest of the supplied bytes.
// Returns an empty string for an empty input.
func HashSHA256(data []byte) string {
	if len(data) == 0 {
		return ""
	}
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}