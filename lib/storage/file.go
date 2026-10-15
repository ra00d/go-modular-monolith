package storage

import (
	"io"
)

// UploadInput is the streaming upload descriptor accepted by Service.Put.
// Content is required; Size and ContentType are optional hints. When
// Size is -1 the service enforces the configured maximum while streaming.
// Detection of the actual MIME type happens after the first bytes have
// been buffered.
type UploadInput struct {
	Content           io.Reader         `json:"-"`
	Size              int64             `json:"size,omitempty"`
	Filename          string            `json:"filename"`
	ContentType       string            `json:"content_type,omitempty"`
	Metadata          map[string]string `json:"metadata,omitempty"`
	Prefix            string            `json:"prefix,omitempty"`
	DetectContentType bool              `json:"detect_content_type,omitempty"`
}

// Options aggregates per-request behaviour that callers may override
// (size limits, key prefix, etc.). Zero values fall back to the service's
// configured defaults.
type Options struct {
	Prefix      string
	MaxSize     int64
	AllowedMime []string
	AllowedExt  []string
	Overwrite   *bool
	Metadata    map[string]string
}