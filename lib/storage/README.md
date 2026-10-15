# File Storage Library

`lib/storage` is a provider-independent file storage library for the
Qistas platform. It supports local filesystem and S3-compatible object
storage (AWS S3, MinIO, others) behind a single, stable API.

## Architecture

```
Application / Handlers / Services
             │
             ▼
    lib/storage.Service        ← validation, key generation, MIME detection
             │
             ▼
    lib/storage.Provider       ← streaming Put/Open/Stat/Delete/Exists
             │
   ┌─────────┴─────────┐
   ▼                   ▼
local.Provider     s3.Provider
```

Application code should depend only on `storage.Service` or
`storage.Provider`. The `local` and `s3` packages are infrastructure
details and must not leak into business logic.

## Configuration

`configs/base/storage.yaml`:

```yaml
storage:
  default: local                  # local | s3
  max_file_size: 52428800         # 50 MiB global cap, 0 = unlimited
  key_prefix: ""                  # prepended to every generated key
  allowed_mime: []                # empty = accept any detected MIME
  allowed_ext:  []                # empty = accept any extension
  overwrite: false                # reject uploads to existing keys
  detect_mime: true               # sniff first bytes, trust the sniffer

  local:
    path: ./storage
    create_if_missing: true

  s3:
    enabled: true
    endpoint: https://s3.amazonaws.com   # set to MinIO URL for MinIO
    bucket: qistas
    region: us-east-1
    access_key: ${S3_ACCESS_KEY}
    secret_key: ${S3_SECRET_KEY}
    session_token: ""             # optional
    path_style: false             # true for MinIO and most non-AWS
    disable_tls: false
```

When `default: s3` the `s3` block must validate. When `default: local`
the `local` block must validate. Both blocks are always available — only
the active one is consulted by `NewProvider`.

## Public interfaces

```go
type Service struct { /* unexported */ }

func NewService(cfg Config, provider Provider, logger *glog.Logger) *Service
func (s *Service) Put(ctx context.Context, in UploadInput) (*Object, error)
func (s *Service) Open(ctx context.Context, key string) (*Download, error)
func (s *Service) Stat(ctx context.Context, key string) (*Stat, error)
func (s *Service) Delete(ctx context.Context, key string) error
func (s *Service) Exists(ctx context.Context, key string) (bool, error)
func (s *Service) Provider() Provider
```

`Provider` is an interface — every backend implements the same five
methods (`Name`, `Put`, `Open`, `Stat`, `Delete`, `Exists`).

## Uploads

```go
obj, err := svc.Put(ctx, storage.UploadInput{
    Content:     file,           // io.Reader; io.ReadSeeker is honoured for size probes
    Size:        fileSize,       // -1 if unknown
    Filename:    "report.pdf",
    ContentType: "application/pdf",
    Prefix:      "documents",
})
```

The service:

1. Validates the key (`storage.ValidateKey`).
2. Drains the reader into a temp buffer to enforce size and compute
   the sha256 checksum when `cfg.Checksum` is enabled.
3. Sniffs the MIME type via `net/http.DetectContentType` when
   `cfg.DetectMime` is true.
4. Generates a collision-resistant key via the configured
   `KeyGenerator`.
5. Persists the bytes through the underlying `Provider`.

## Downloads

```go
dl, err := svc.Open(ctx, key)
if err != nil { /* map via storage.AppError */ }
defer dl.Reader.Close()
io.Copy(w, dl.Reader)
```

`dl.Reader` is streaming — never load it into memory in business code.

## HTTP integration

```go
// Server-side: serve a stored object with proper headers.
storage.ServeFromService(svc, w, r, key, storage.DispositionAttachment)

// Server-side: handle an upload from a multipart form.
file, err := storage.FromGFRequest(r, "file")
if err != nil { storage.WriteError(r, err); return }
defer file.Reader.Close()
obj, err := svc.Put(r.Context(), storage.UploadInput{
    Content: file.Reader, Size: file.Size,
    Filename: file.Filename, ContentType: file.ContentType,
})
if err != nil { storage.WriteError(r, err); return }
r.Response.WriteJson(obj)
```

`Content-Disposition` honours RFC 5987 so non-ASCII filenames round-trip
safely. `ETag`, `Last-Modified`, `Cache-Control` and `Content-Length`
are set automatically when known.

## Errors

Storage errors are sentinel values that map cleanly to HTTP status
codes via `storage.AppError`:

| Sentinel                     | Code                       | HTTP |
|------------------------------|----------------------------|------|
| `ErrInvalidFile`             | `storage.invalid_file`     | 400  |
| `ErrFileTooLarge`            | `storage.file_too_large`   | 400  |
| `ErrUnsupportedType`         | `storage.unsupported_type` | 400  |
| `ErrBadKey`                  | `storage.invalid_key`      | 400  |
| `ErrValidationFailed`        | `storage.validation_failed`| 400  |
| `ErrConfigError`             | `storage.config_error`     | 400  |
| `ErrNotFound`                | `storage.not_found`        | 404  |
| `ErrAlreadyExists`           | `storage.already_exists`   | 409  |
| `ErrPermissionDenied`        | `storage.permission_denied`| 403  |
| `ErrUnavailable`             | `storage.unavailable`      | 500  |

Use `errors.Is(err, storage.ErrXxx)` for comparison; `storage.AppError`
converts to `lib/errors.AppError` so handlers can rely on the project's
HTTP status mapping.

## Dependency injection

```go
// cmd/server/main.go
import "modular-monolith/lib/storage"

func main() {
    a := app.NewApp(
        storage.Module,   // ← adds the storage module
        // ...
    )
    a.Run()
}
```

The module registers `storage.Provider` and `storage.Service` for fx.
Other modules consume them via constructors:

```go
fx.Provide(func(svc *storage.Service) myDomain.FileStore {
    return myDomain.NewFileStore(svc)
}),
```

## Adding a new provider

1. Create a package `lib/storage/<name>/`.
2. Implement the `contract.Provider` interface.
3. Add a constructor `New<Name>Provider(ctx, cfg)`.
4. Wire it in `lib/storage/module.go`:
   ```go
   case ProviderX:
       return xprov.NewProvider(ctx, *lc.X)
   ```
5. Add tests; reuse the existing `storage.FakeProvider` for service
   tests so application code paths are exercised without your backend.

## Testing

- `lib/storage/*_test.go`: validation, key generation, errors, MIME,
  filename sanitisation, HTTP helpers.
- `lib/storage/local/local_test.go`: filesystem lifecycle, path
  traversal, tmp-file cleanup, concurrent writes.
- `lib/storage/s3/s3_test.go`: endpoint validation, SDK error mapping.
- `FakeProvider` (in `lib/storage/fake.go`): in-memory provider for
  application tests.

For integration tests against a real S3-compatible service, run a
local MinIO container and set `STORAGE_DEFAULT=s3` plus the MinIO
endpoint in the test environment.

## Security

- Uploaded filenames are never used as filesystem paths; only the
  generated key is.
- Keys are validated against path traversal, absolute paths, NUL bytes
  and ".." segments.
- MIME detection runs against the actual bytes — the client's
  `Content-Type` is treated as a hint.
- Optional size limits prevent unbounded uploads.
- Credentials are never logged; the storage logger omits key material
  and treats the bucket / key as the only identifying fields.