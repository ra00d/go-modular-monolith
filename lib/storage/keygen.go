package storage

import (
	"path"
	"strings"
	"time"

	"github.com/google/uuid"

	"modular-monolith/lib/storage/contract"
)

// KeyGenerator builds storage keys from a logical prefix, the upload's
// original filename and the upload timestamp. The default generator
// produces:
//
//	{prefix}/{YYYY/MM/DD}/{uuid}/{sanitized-name}.{ext}
//
// Implementations must guarantee:
//
//   - The returned key never contains "..", leading "/", or a NUL byte.
//   - The key never starts with an absolute path component.
//   - Two concurrent uploads with the same input never collide (UUID).
type KeyGenerator func(prefix, filename string, now time.Time) (key string, safeName string, err error)

// DefaultKeyGenerator is the canonical key generator. It is exported so
// callers can wrap or replace it without re-implementing the safety rules.
func DefaultKeyGenerator(prefix, filename string, now time.Time) (string, string, error) {
	safe, ext := SanitizeFilename(filename)
	prefix = cleanPrefix(prefix)
	id := uuid.New().String()
	date := now.UTC().Format("2006/01/02")
	var b strings.Builder
	if prefix != "" {
		b.WriteString(prefix)
		b.WriteByte('/')
	}
	b.WriteString(date)
	b.WriteByte('/')
	b.WriteString(id)
	b.WriteByte('/')
	b.WriteString(safe)
	if ext != "" {
		b.WriteByte('.')
		b.WriteString(ext)
	}
	key := b.String()
	if err := contract.ValidateKey(key); err != nil {
		return "", "", err
	}
	return key, path.Base(key), nil
}

// cleanPrefix trims slashes, collapses empty segments and removes "." /
// ".." so a malicious or accidental prefix cannot change the resolved
// location of generated keys.
func cleanPrefix(p string) string {
	p = strings.TrimSpace(p)
	if p == "" {
		return ""
	}
	p = strings.ReplaceAll(p, `\`, "/")
	p = strings.Trim(p, "/")
	if p == "" {
		return ""
	}
	parts := make([]string, 0, 4)
	for _, seg := range strings.Split(p, "/") {
		switch seg {
		case "", ".":
			continue
		case "..":
			if len(parts) > 0 {
				parts = parts[:len(parts)-1]
			}
		default:
			parts = append(parts, seg)
		}
	}
	return strings.Join(parts, "/")
}