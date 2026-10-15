package local

import (
	"bytes"
	"context"
	stderrors "errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"modular-monolith/lib/config"
	"modular-monolith/lib/storage/contract"
)

func newProvider(t *testing.T) *Provider {
	t.Helper()
	root := t.TempDir()
	p, err := NewProvider(config.LocalStorageCfg{
		Path:            root,
		CreateIfMissing: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	return p
}

func TestNewProviderValidates(t *testing.T) {
	if _, err := NewProvider(config.LocalStorageCfg{}); err == nil {
		t.Fatal("expected error for empty path")
	}
	if _, err := NewProvider(config.LocalStorageCfg{Path: filepath.Join(t.TempDir(), "file")}); err == nil {
		t.Fatal("expected error when path is not a directory")
	}
}

func TestPutAndOpen(t *testing.T) {
	p := newProvider(t)
	ctx := context.Background()
	body := []byte("hello world")
	key := "uploads/file.txt"
	res, err := p.Put(ctx, key, "text/plain", int64(len(body)), bytes.NewReader(body), nil)
	if err != nil {
		t.Fatal(err)
	}
	if res.Size != int64(len(body)) {
		t.Errorf("Size = %d, want %d", res.Size, len(body))
	}
	if res.LastModified.IsZero() {
		t.Errorf("LastModified must be set")
	}

	dl, err := p.Open(ctx, key)
	if err != nil {
		t.Fatal(err)
	}
	defer dl.Reader.Close()
	got, _ := io.ReadAll(dl.Reader)
	if !bytes.Equal(got, body) {
		t.Errorf("round-trip mismatch: %q vs %q", got, body)
	}
	if dl.ContentType != "text/plain" {
		t.Errorf("ContentType = %q", dl.ContentType)
	}
}

func TestOpenMissing(t *testing.T) {
	p := newProvider(t)
	_, err := p.Open(context.Background(), "missing.txt")
	if !stderrors.Is(err, contract.ErrNotFound) {
		t.Errorf("err = %v, want ErrNotFound", err)
	}
}

func TestStat(t *testing.T) {
	p := newProvider(t)
	ctx := context.Background()
	body := []byte("abc")
	if _, err := p.Put(ctx, "k", "text/plain", int64(len(body)), bytes.NewReader(body), nil); err != nil {
		t.Fatal(err)
	}
	stat, err := p.Stat(ctx, "k")
	if err != nil {
		t.Fatal(err)
	}
	if stat.Size != int64(len(body)) {
		t.Errorf("Size = %d", stat.Size)
	}
	if stat.ContentType != "text/plain" {
		t.Errorf("ContentType = %q", stat.ContentType)
	}
}

func TestDelete(t *testing.T) {
	p := newProvider(t)
	ctx := context.Background()
	if _, err := p.Put(ctx, "k", "text/plain", 1, strings.NewReader("x"), nil); err != nil {
		t.Fatal(err)
	}
	if err := p.Delete(ctx, "k"); err != nil {
		t.Fatal(err)
	}
	exists, _ := p.Exists(ctx, "k")
	if exists {
		t.Error("file still exists after delete")
	}
	// Deleting again must be a no-op.
	if err := p.Delete(ctx, "k"); err != nil {
		t.Errorf("delete of missing key should be silent, got %v", err)
	}
}

func TestExists(t *testing.T) {
	p := newProvider(t)
	ctx := context.Background()
	if _, err := p.Put(ctx, "k", "text/plain", 1, strings.NewReader("x"), nil); err != nil {
		t.Fatal(err)
	}
	exists, err := p.Exists(ctx, "k")
	if err != nil || !exists {
		t.Errorf("Exists = %v, %v", exists, err)
	}
	exists, _ = p.Exists(ctx, "missing")
	if exists {
		t.Error("Exists returned true for missing key")
	}
}

func TestPathTraversalRejected(t *testing.T) {
	p := newProvider(t)
	ctx := context.Background()
	bad := []string{
		"../etc/passwd",
		"/etc/passwd",
		"foo/../../bar",
		"..",
		".",
	}
	for _, k := range bad {
		if _, err := p.Put(ctx, k, "x", 1, strings.NewReader("x"), nil); err == nil {
			t.Errorf("Put(%q) must reject path traversal", k)
		}
		if _, err := p.Open(ctx, k); err == nil {
			t.Errorf("Open(%q) must reject path traversal", k)
		}
		if _, err := p.Stat(ctx, k); err == nil {
			t.Errorf("Stat(%q) must reject path traversal", k)
		}
	}
}

func TestPutOverwritesByDefault(t *testing.T) {
	p := newProvider(t)
	ctx := context.Background()
	if _, err := p.Put(ctx, "k", "text/plain", 1, strings.NewReader("x"), nil); err != nil {
		t.Fatal(err)
	}
	// Provider level is overwrite-friendly; service-level policy enforces
	// ErrAlreadyExists when overwrite=false.
	res, err := p.Put(ctx, "k", "text/plain", 1, strings.NewReader("y"), nil)
	if err != nil {
		t.Fatalf("provider overwrite must succeed: %v", err)
	}
	dl, err := p.Open(ctx, "k")
	if err != nil {
		t.Fatal(err)
	}
	defer dl.Reader.Close()
	got, _ := io.ReadAll(dl.Reader)
	if string(got) != "y" {
		t.Errorf("body = %q, want y", got)
	}
	if res == nil || res.Size != 1 {
		t.Errorf("unexpected result: %+v", res)
	}
}

func TestPutDoesNotLeaveTmpFiles(t *testing.T) {
	p := newProvider(t)
	ctx := context.Background()
	_, err := p.Put(ctx, "k", "text/plain", 1, strings.NewReader("x"), nil)
	if err != nil {
		t.Fatal(err)
	}
	matches, err := filepath.Glob(filepath.Join(p.Root(), "k.tmp.*"))
	if err != nil {
		t.Fatal(err)
	}
	if len(matches) != 0 {
		t.Errorf("leftover tmp files: %v", matches)
	}
}

func TestPutFailureRemovesTmp(t *testing.T) {
	p := newProvider(t)
	ctx := context.Background()
	// Reader that errors on second read by returning a wrapped error.
	r := &failingReader{}
	_, err := p.Put(ctx, "k", "text/plain", -1, r, nil)
	if err == nil {
		t.Fatal("expected error from failing reader")
	}
	matches, _ := filepath.Glob(filepath.Join(p.Root(), "*.tmp.*"))
	if len(matches) != 0 {
		t.Errorf("tmp file leaked after failure: %v", matches)
	}
	if _, err := os.Stat(filepath.Join(p.Root(), "k")); err == nil {
		t.Error("destination file should not exist after failed upload")
	}
}

type failingReader struct {
	calls int
}

func (r *failingReader) Read(p []byte) (int, error) {
	r.calls++
	if r.calls == 1 {
		copy(p, []byte("x"))
		return 1, nil
	}
	return 0, os.ErrClosed
}

func TestConcurrentPuts(t *testing.T) {
	p := newProvider(t)
	ctx := context.Background()
	const n = 50
	mu := sync.Mutex{}
	keys := make(map[string]struct{}, n)
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("concurrent/%d.bin", i)
			if _, err := p.Put(ctx, key, "text/plain", 4, strings.NewReader("data"), nil); err != nil {
				t.Errorf("put %d failed: %v", i, err)
				return
			}
			mu.Lock()
			keys[key] = struct{}{}
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	if len(keys) != n {
		t.Errorf("unique keys = %d, want %d", len(keys), n)
	}
}