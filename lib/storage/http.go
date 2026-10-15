package storage

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"path/filepath"
	"strings"

	apperrors "modular-monolith/lib/errors"

	"github.com/gogf/gf/v2/net/ghttp"
)

// HTTPDisposition controls how the browser renders a downloaded file.
type HTTPDisposition int

const (
	// DispositionAttachment prompts a "Save As" dialog with the filename.
	DispositionAttachment HTTPDisposition = iota
	// DispositionInline renders the file inline when the browser can.
	DispositionInline
)

// ContentDisposition builds a Content-Disposition header value honouring
// non-ASCII filenames (RFC 5987) so modern browsers keep the original
// name even when it contains UTF-8 characters.
func ContentDisposition(disposition HTTPDisposition, filename string) string {
	disp := "attachment"
	if disposition == DispositionInline {
		disp = "inline"
	}
	if filename == "" {
		return disp
	}
	// Strip directory components and CRLF injection attempts.
	filename = filepath.Base(strings.ReplaceAll(filename, "\r", ""))
	filename = strings.ReplaceAll(filename, "\n", "")
	ascii := sanitizeASCII(filename)
	encoded := mime.QEncoding.Encode("utf-8", filename)
	if ascii == filename {
		return fmt.Sprintf("%s; filename=%q", disp, filename)
	}
	return fmt.Sprintf("%s; filename=%q; filename*=UTF-8''%s", disp, ascii, encoded)
}

// sanitizeASCII replaces non-ASCII runes with '?' so the legacy
// filename= parameter stays valid per RFC 2616.
func sanitizeASCII(s string) string {
	var b strings.Builder
	for _, r := range s {
		if r > 0x7F {
			b.WriteByte('?')
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}

// WriteHTTP streams a stored object to an HTTP response. It sets the
// content type, length (when known), filename, ETag, last-modified and
// cache-control headers, then copies the body using io.Copy so the file
// is not buffered in memory.
func WriteHTTP(w http.ResponseWriter, stat *Stat, body io.Reader, disp HTTPDisposition, cacheControl string) {
	if stat == nil {
		http.Error(w, "missing object", http.StatusInternalServerError)
		return
	}
	ct := stat.ContentType
	if ct == "" {
		ct = "application/octet-stream"
	}
	w.Header().Set("Content-Type", ct)
	w.Header().Set("Content-Disposition", ContentDisposition(disp, stat.Key))
	if stat.Size > 0 {
		w.Header().Set("Content-Length", fmt.Sprintf("%d", stat.Size))
	}
	if stat.ETag != "" {
		w.Header().Set("ETag", `"`+stat.ETag+`"`)
	}
	if !stat.LastModified.IsZero() {
		w.Header().Set("Last-Modified", stat.LastModified.UTC().Format(http.TimeFormat))
	}
	if cacheControl == "" {
		cacheControl = "private, max-age=300"
	}
	w.Header().Set("Cache-Control", cacheControl)
	if body != nil {
		_, _ = io.Copy(w, body)
	}
}

// ServeFromService is a convenience wrapper that handles Stat + Open +
// body copy + defer-close for HTTP handlers.
func ServeFromService(svc *Service, w http.ResponseWriter, r *http.Request, key string, disp HTTPDisposition) error {
	if svc == nil {
		http.Error(w, "storage unavailable", http.StatusServiceUnavailable)
		return ErrUnavailable
	}
	dl, err := svc.Open(r.Context(), key)
	if err != nil {
		return err
	}
	defer dl.Reader.Close()
	stat := &Stat{
		Key:          key,
		Size:         dl.Size,
		ContentType:  dl.ContentType,
		ETag:         dl.ETag,
		LastModified: dl.LastModified,
	}
	WriteHTTP(w, stat, dl.Reader, disp, "")
	return nil
}

// HTTPRequestFile wraps the ghttp request to extract an uploaded file
// without forcing business code to import the ghttp package directly.
type HTTPRequestFile struct {
	Filename    string
	ContentType string
	Size        int64
	Reader      io.ReadCloser
}

// FromMultipart extracts the first uploaded file from the request, using
// the supplied form-field name. The returned reader must be closed by
// the caller. Returns ErrInvalidFile when no file is present.
func FromMultipart(r *http.Request, field string) (*HTTPRequestFile, error) {
	if r.MultipartForm == nil {
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			return nil, fmt.Errorf("%w: parse multipart: %v", ErrInvalidFile, err)
		}
	}
	if r.MultipartForm == nil || r.MultipartForm.File == nil {
		return nil, fmt.Errorf("%w: empty form", ErrInvalidFile)
	}
	files := r.MultipartForm.File[field]
	if len(files) == 0 {
		return nil, fmt.Errorf("%w: field %q not found", ErrInvalidFile, field)
	}
	fh := files[0]
	file, err := fh.Open()
	if err != nil {
		return nil, fmt.Errorf("%w: open part: %v", ErrInvalidFile, err)
	}
	ct := fh.Header.Get("Content-Type")
	return &HTTPRequestFile{
		Filename:    fh.Filename,
		ContentType: ct,
		Size:        fh.Size,
		Reader:      file,
	}, nil
}

// FromGFRequest is the ghttp-flavored equivalent of FromMultipart. It
// preserves the project's HTTP-framework dependency rather than forcing
// callers to drop to net/http directly.
func FromGFRequest(r *ghttp.Request, field string) (*HTTPRequestFile, error) {
	files := r.GetUploadFiles(field)
	if len(files) == 0 {
		return nil, fmt.Errorf("%w: field %q not found", ErrInvalidFile, field)
	}
	f := files[0]
	ct := ""
	if f.Header != nil {
		ct = f.Header.Get("Content-Type")
	}
	// ghttp.UploadFile exposes the underlying multipart.File via Open().
	rc, err := f.Open()
	if err != nil {
		return nil, fmt.Errorf("%w: open part: %v", ErrInvalidFile, err)
	}
	return &HTTPRequestFile{
		Filename:    f.Filename,
		ContentType: ct,
		Size:        f.Size,
		Reader:      rc,
	}, nil
}

// WriteError renders a storage error to the supplied ghttp response. It
// uses the project's error model so the response body matches every
// other endpoint (Code / Error / Data shape).
func WriteError(r *ghttp.Request, err error) {
	if err == nil {
		return
	}
	ae := MapError(err)
	status := 0
	if ae != nil {
		status = apperrors.StatusCode(ae)
	}
	if status == 0 {
		status = http.StatusInternalServerError
	}
	r.Response.Status = status
	r.Response.WriteJsonExit(struct {
		Code  string `json:"code"`
		Error string `json:"error"`
	}{
		Code:  ae.Code,
		Error: ae.Message,
	})
}