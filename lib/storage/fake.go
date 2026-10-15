package storage

import (
	"bytes"
	"context"
	stderrors "errors"
	"io"
	"strings"
	"sync"
	"time"
)

// FakeProvider is an in-memory Provider suitable for tests. The
// implementation is concurrency-safe and never touches disk or network.
// Application tests should depend on FakeProvider through Service so
// production paths are exercised even when the backend is fake.
type FakeProvider struct {
	mu        sync.Mutex
	objects   map[string]*fakeObject
	putHook   func(key string, size int64) error
	openHook  func(key string) error
	statHook  func(key string) error
	bucket    string
	createdAt time.Time
}

type fakeObject struct {
	contentType string
	body        []byte
	metadata    map[string]string
	etag        string
	createdAt   time.Time
}

// NewFakeProvider returns an empty fake provider.
func NewFakeProvider() *FakeProvider {
	return &FakeProvider{
		objects:   make(map[string]*fakeObject),
		createdAt: time.Now(),
	}
}

// SetBucket configures the bucket name reported by Bucket().
func (f *FakeProvider) SetBucket(b string) { f.bucket = b }

// SetPutHook installs a hook invoked before each Put. The hook can
// return an error to simulate transient backend failures.
func (f *FakeProvider) SetPutHook(fn func(key string, size int64) error) { f.putHook = fn }

// SetOpenHook installs a hook invoked before each Open.
func (f *FakeProvider) SetOpenHook(fn func(key string) error) { f.openHook = fn }

// SetStatHook installs a hook invoked before each Stat.
func (f *FakeProvider) SetStatHook(fn func(key string) error) { f.statHook = fn }

// Name implements Provider.
func (f *FakeProvider) Name() string { return "fake" }

// Bucket returns the configured bucket (empty for fake by default).
func (f *FakeProvider) Bucket() string { return f.bucket }

// Put stores the body in memory.
func (f *FakeProvider) Put(ctx context.Context, key, contentType string, size int64, src io.Reader, meta map[string]string) (*PutResult, error) {
	if err := ValidateKey(key); err != nil {
		return nil, err
	}
	if f.putHook != nil {
		if err := f.putHook(key, size); err != nil {
			return nil, err
		}
	}
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	body, err := io.ReadAll(src)
	if err != nil {
		return nil, err
	}
	if size > 0 && int64(len(body)) != size {
		return nil, ErrInvalidFile
	}
	f.mu.Lock()
	defer f.mu.Unlock()
	if _, exists := f.objects[key]; exists {
		return nil, ErrAlreadyExists
	}
	f.objects[key] = &fakeObject{
		contentType: contentType,
		body:        body,
		metadata:    cloneMap(meta),
		etag:        "etag-" + key,
		createdAt:   time.Now(),
	}
	return &PutResult{
		Size:         int64(len(body)),
		ETag:         "etag-" + key,
		LastModified: f.objects[key].createdAt,
	}, nil
}

// Open returns a reader over the stored body.
func (f *FakeProvider) Open(ctx context.Context, key string) (*Download, error) {
	if err := ValidateKey(key); err != nil {
		return nil, err
	}
	if f.openHook != nil {
		if err := f.openHook(key); err != nil {
			return nil, err
		}
	}
	f.mu.Lock()
	obj, ok := f.objects[key]
	f.mu.Unlock()
	if !ok {
		return nil, ErrNotFound
	}
	return &Download{
		Reader:       io.NopCloser(bytes.NewReader(obj.body)),
		Size:         int64(len(obj.body)),
		ContentType:  obj.contentType,
		Name:         key,
		ETag:         obj.etag,
		LastModified: obj.createdAt,
	}, nil
}

// Stat returns the stored metadata.
func (f *FakeProvider) Stat(ctx context.Context, key string) (*Stat, error) {
	if err := ValidateKey(key); err != nil {
		return nil, err
	}
	if f.statHook != nil {
		if err := f.statHook(key); err != nil {
			return nil, err
		}
	}
	f.mu.Lock()
	obj, ok := f.objects[key]
	f.mu.Unlock()
	if !ok {
		return nil, ErrNotFound
	}
	return &Stat{
		Key:          key,
		Size:         int64(len(obj.body)),
		ContentType:  obj.contentType,
		ETag:         obj.etag,
		LastModified: obj.createdAt,
		Metadata:     cloneMap(obj.metadata),
	}, nil
}

// Delete removes the key. Missing keys are not errors.
func (f *FakeProvider) Delete(ctx context.Context, key string) error {
	if err := ValidateKey(key); err != nil {
		return err
	}
	f.mu.Lock()
	defer f.mu.Unlock()
	delete(f.objects, key)
	return nil
}

// Exists reports whether the key is present.
func (f *FakeProvider) Exists(ctx context.Context, key string) (bool, error) {
	if err := ValidateKey(key); err != nil {
		return false, err
	}
	f.mu.Lock()
	defer f.mu.Unlock()
	_, ok := f.objects[key]
	return ok, nil
}

// Keys returns a snapshot of stored keys (sorted-ish, for assertions).
func (f *FakeProvider) Keys() []string {
	f.mu.Lock()
	defer f.mu.Unlock()
	out := make([]string, 0, len(f.objects))
	for k := range f.objects {
		out = append(out, k)
	}
	return out
}

func cloneMap(in map[string]string) map[string]string {
	if in == nil {
		return nil
	}
	out := make(map[string]string, len(in))
	for k, v := range in {
		out[k] = v
	}
	return out
}

// errFake is a sentinel used by tests to trigger transient failures.
var errFake = stderrors.New("fake: simulated")

// IsFake reports whether err originated from one of the FakeProvider hooks.
func IsFake(err error) bool { return stderrors.Is(err, errFake) }

// matchKey is a tiny helper used by tests to assert on key substrings.
func matchKey(haystack, needle string) bool {
	return strings.Contains(haystack, needle)
}