package storage

import (
	"bytes"
	"context"
	stderrors "errors"
	"io"
	"strings"
	"testing"
)

func newTestService(provider Provider, cfg Config) *Service {
	return NewService(cfg, provider, nil)
}

func TestServicePutAndGet(t *testing.T) {
	fake := NewFakeProvider()
	svc := newTestService(fake, Config{
		Provider: ProviderLocal,
		MaxSize:  1 << 20,
	})

	payload := []byte("%PDF-1.4 hello world")
	obj, err := svc.Put(context.Background(), UploadInput{
		Content:     bytes.NewReader(payload),
		Size:        int64(len(payload)),
		Filename:    "report.pdf",
		ContentType: "application/pdf",
		Prefix:      "documents",
	})
	if err != nil {
		t.Fatal(err)
	}
	if obj.Size != int64(len(payload)) {
		t.Errorf("size = %d, want %d", obj.Size, len(payload))
	}
	if obj.ContentType != "application/pdf" {
		t.Errorf("content type = %q", obj.ContentType)
	}
	if !strings.HasPrefix(obj.Key, "documents/") {
		t.Errorf("key prefix missing: %s", obj.Key)
	}
	if obj.Extension != "pdf" {
		t.Errorf("extension = %q", obj.Extension)
	}

	dl, err := svc.Open(context.Background(), obj.Key)
	if err != nil {
		t.Fatal(err)
	}
	defer dl.Reader.Close()
	got, _ := io.ReadAll(dl.Reader)
	if !bytes.Equal(got, payload) {
		t.Errorf("round-trip mismatch: %q vs %q", got, payload)
	}
}

func TestServicePutRejectsPathTraversal(t *testing.T) {
	fake := NewFakeProvider()
	svc := newTestService(fake, Config{})
	_, err := svc.Put(context.Background(), UploadInput{
		Content:  strings.NewReader("x"),
		Filename: "../../../etc/passwd",
	})
	if err != nil {
		t.Fatalf("unexpected error from Put: %v", err)
	}
	// Generated key must not contain ".." or escape the prefix.
	for _, k := range fake.Keys() {
		if strings.Contains(k, "..") {
			t.Errorf("generated key %q contains ..", k)
		}
	}
}

func TestServicePutRejectsOversize(t *testing.T) {
	fake := NewFakeProvider()
	svc := newTestService(fake, Config{MaxSize: 10})
	_, err := svc.Put(context.Background(), UploadInput{
		Content:  bytes.NewReader(make([]byte, 100)),
		Filename: "big.bin",
	})
	if !stderrors.Is(err, ErrFileTooLarge) {
		t.Errorf("err = %v, want ErrFileTooLarge", err)
	}
}

func TestServiceRejectsZeroBody(t *testing.T) {
	fake := NewFakeProvider()
	svc := newTestService(fake, Config{})
	_, err := svc.Put(context.Background(), UploadInput{
		Content:  bytes.NewReader(nil),
		Filename: "empty.bin",
	})
	if !stderrors.Is(err, ErrInvalidFile) {
		t.Errorf("err = %v, want ErrInvalidFile", err)
	}
}

func TestServiceMimeDetection(t *testing.T) {
	fake := NewFakeProvider()
	svc := newTestService(fake, Config{
		MaxSize:    1 << 20,
		DetectMime: true,
	})
	// PNG header
	header := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
	body := append(header, bytes.Repeat([]byte{0}, 64)...)
	obj, err := svc.Put(context.Background(), UploadInput{
		Content:     bytes.NewReader(body),
		Size:        int64(len(body)),
		Filename:    "img.png",
		ContentType: "application/octet-stream", // intentionally wrong
	})
	if err != nil {
		t.Fatal(err)
	}
	if obj.ContentType != "image/png" {
		t.Errorf("detected type = %q, want image/png", obj.ContentType)
	}
}

func TestServiceMimeWhitelist(t *testing.T) {
	fake := NewFakeProvider()
	svc := newTestService(fake, Config{
		MaxSize:     1 << 20,
		DetectMime:  true,
		AllowedMime: []string{"image/png"},
	})
	header := []byte{0xFF, 0xD8, 0xFF} // JPEG
	body := append(header, bytes.Repeat([]byte{0}, 32)...)
	_, err := svc.Put(context.Background(), UploadInput{
		Content:  bytes.NewReader(body),
		Size:     int64(len(body)),
		Filename: "img.jpg",
	})
	if !stderrors.Is(err, ErrUnsupportedType) {
		t.Errorf("err = %v, want ErrUnsupportedType", err)
	}
}

func TestServiceExtWhitelist(t *testing.T) {
	fake := NewFakeProvider()
	svc := newTestService(fake, Config{
		MaxSize:    1 << 20,
		AllowedExt: []string{"pdf"},
	})
	_, err := svc.Put(context.Background(), UploadInput{
		Content:  strings.NewReader("hello"),
		Size:     5,
		Filename: "file.exe",
	})
	if !stderrors.Is(err, ErrUnsupportedType) {
		t.Errorf("err = %v, want ErrUnsupportedType", err)
	}
}

func TestServiceBadKeyRejected(t *testing.T) {
	fake := NewFakeProvider()
	svc := newTestService(fake, Config{})
	_, err := svc.Open(context.Background(), "../etc/passwd")
	if !stderrors.Is(err, ErrBadKey) {
		t.Errorf("err = %v, want ErrBadKey", err)
	}
}

func TestServiceDelete(t *testing.T) {
	fake := NewFakeProvider()
	svc := newTestService(fake, Config{})
	obj, err := svc.Put(context.Background(), UploadInput{
		Content:  strings.NewReader("x"),
		Filename: "x.txt",
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := svc.Delete(context.Background(), obj.Key); err != nil {
		t.Fatalf("delete: %v", err)
	}
	exists, _ := svc.Exists(context.Background(), obj.Key)
	if exists {
		t.Error("object still exists after delete")
	}
}

func TestServiceNotFoundMapped(t *testing.T) {
	fake := NewFakeProvider()
	svc := newTestService(fake, Config{})
	_, err := svc.Open(context.Background(), "nope/missing.txt")
	if !stderrors.Is(err, ErrNotFound) {
		t.Errorf("err = %v, want ErrNotFound", err)
	}
}

func TestServicePutChecksum(t *testing.T) {
	fake := NewFakeProvider()
	svc := newTestService(fake, Config{
		MaxSize:  1024,
		Checksum: true,
	})
	body := []byte("checksum me")
	obj, err := svc.Put(context.Background(), UploadInput{
		Content:  bytes.NewReader(body),
		Size:     int64(len(body)),
		Filename: "x.bin",
	})
	if err != nil {
		t.Fatal(err)
	}
	want := HashSHA256(body)
	if obj.Checksum != want {
		t.Errorf("checksum = %q, want %q", obj.Checksum, want)
	}
}

func TestServiceDeleteAlreadyGone(t *testing.T) {
	fake := NewFakeProvider()
	svc := newTestService(fake, Config{})
	if err := svc.Delete(context.Background(), "missing/file"); err != nil {
		t.Errorf("delete of missing key must be no-op, got %v", err)
	}
}