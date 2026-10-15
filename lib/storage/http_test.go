package storage

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestContentDisposition(t *testing.T) {
	cases := []struct {
		disp HTTPDisposition
		name string
		want []string
	}{
		{DispositionAttachment, "report.pdf", []string{"attachment", "filename=\"report.pdf\""}},
		{DispositionInline, "image.png", []string{"inline", "filename=\"image.png\""}},
		{DispositionAttachment, "with spaces.pdf", []string{"attachment", "filename=\"with spaces.pdf\""}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := ContentDisposition(c.disp, c.name)
			for _, want := range c.want {
				if !strings.Contains(got, want) {
					t.Errorf("ContentDisposition(%v, %q) = %q, missing %q", c.disp, c.name, got, want)
				}
			}
		})
	}

	// UTF-8 filename must be RFC 5987 encoded.
	got := ContentDisposition(DispositionAttachment, "中文.pdf")
	if !strings.Contains(got, "filename*=UTF-8''") {
		t.Errorf("UTF-8 filename not RFC 5987 encoded: %q", got)
	}
	// Path traversal attempts must be sanitised.
	got = ContentDisposition(DispositionAttachment, "../../etc/passwd")
	if strings.Contains(got, "..") {
		t.Errorf("path traversal survived: %q", got)
	}
	// Empty filename should still produce a valid disposition.
	got = ContentDisposition(DispositionAttachment, "")
	if got != "attachment" {
		t.Errorf("empty filename produced %q", got)
	}
}

func TestWriteHTTP(t *testing.T) {
	rec := httptest.NewRecorder()
	body := bytes.NewReader([]byte("hello"))
	stat := &Stat{
		Key:          "foo/bar.txt",
		Size:         5,
		ContentType:  "text/plain",
		ETag:         "abc",
		LastModified: time.Unix(1700000000, 0).UTC(),
	}
	WriteHTTP(rec, stat, body, DispositionInline, "public, max-age=60")

	if rec.Header().Get("Content-Type") != "text/plain" {
		t.Errorf("Content-Type = %q", rec.Header().Get("Content-Type"))
	}
	if rec.Header().Get("Content-Length") != "5" {
		t.Errorf("Content-Length = %q", rec.Header().Get("Content-Length"))
	}
	if !strings.Contains(rec.Header().Get("Content-Disposition"), "inline") {
		t.Errorf("Content-Disposition = %q", rec.Header().Get("Content-Disposition"))
	}
	if rec.Header().Get("ETag") != `"abc"` {
		t.Errorf("ETag = %q", rec.Header().Get("ETag"))
	}
	if rec.Header().Get("Cache-Control") != "public, max-age=60" {
		t.Errorf("Cache-Control = %q", rec.Header().Get("Cache-Control"))
	}
	if rec.Body.String() != "hello" {
		t.Errorf("body = %q", rec.Body.String())
	}
}

func TestWriteHTTPDefaultCache(t *testing.T) {
	rec := httptest.NewRecorder()
	WriteHTTP(rec, &Stat{Key: "k", ContentType: "x/x"}, bytes.NewReader(nil), DispositionAttachment, "")
	if !strings.HasPrefix(rec.Header().Get("Cache-Control"), "private") {
		t.Errorf("Cache-Control default = %q", rec.Header().Get("Cache-Control"))
	}
}

func TestServeFromService(t *testing.T) {
	fake := NewFakeProvider()
	svc := newTestService(fake, Config{MaxSize: 1024})
	body := []byte("streaming")
	obj, err := svc.Put(context.Background(), UploadInput{
		Content:  bytes.NewReader(body),
		Size:     int64(len(body)),
		Filename: "stream.bin",
	})
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/files/"+obj.Key, nil)
	if err := ServeFromService(svc, rec, req, obj.Key, DispositionAttachment); err != nil {
		t.Fatalf("ServeFromService: %v", err)
	}
	if rec.Body.String() != "streaming" {
		t.Errorf("body = %q", rec.Body.String())
	}
	if !strings.Contains(rec.Header().Get("Content-Disposition"), "attachment") {
		t.Errorf("Content-Disposition = %q", rec.Header().Get("Content-Disposition"))
	}
}

func TestServeFromServiceMissing(t *testing.T) {
	fake := NewFakeProvider()
	svc := newTestService(fake, Config{})
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/files/missing", nil)
	if err := ServeFromService(svc, rec, req, "missing/file.bin", DispositionAttachment); !isError(err) {
		t.Errorf("expected error, got %v", err)
	}
}

func isError(err error) bool { return err != nil }