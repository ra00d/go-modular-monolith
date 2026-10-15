// Package storage provides a provider-independent file storage abstraction
// for the Qistas platform. The library exposes:
//
//   - Provider: a small interface implemented by every concrete backend
//     (local filesystem, S3-compatible object storage, ...).
//   - Service:  a higher-level façade that performs validation, MIME
//     detection, safe key generation, and error mapping before delegating
//     to the configured Provider.
//
// Application code should depend on Service or Provider via dependency
// injection. Concrete types from the local/ or s3/ sub-packages must not
// leak into business logic.
package storage

import (
	stderrors "errors"

	"modular-monolith/lib/storage/contract"

	apperrors "modular-monolith/lib/errors"
)

// Stable error codes returned across the library. Callers identify
// failure modes via errors.Is(err, storage.ErrXxx).
const (
	CodeInvalidFile      = "storage.invalid_file"
	CodeFileTooLarge     = "storage.file_too_large"
	CodeUnsupportedType  = "storage.unsupported_type"
	CodeNotFound         = "storage.not_found"
	CodeAlreadyExists    = "storage.already_exists"
	CodeUnavailable      = "storage.unavailable"
	CodeConfigError      = "storage.config_error"
	CodePermissionDenied = "storage.permission_denied"
	CodeValidationFailed = "storage.validation_failed"
	CodeBadKey           = "storage.invalid_key"
)

// Sentinel errors re-exported from the contract package so callers see
// a single import path.
var (
	ErrInvalidFile      = contract.ErrInvalidFile
	ErrFileTooLarge     = contract.ErrFileTooLarge
	ErrUnsupportedType  = contract.ErrUnsupportedType
	ErrNotFound         = contract.ErrNotFound
	ErrAlreadyExists    = contract.ErrAlreadyExists
	ErrUnavailable      = contract.ErrUnavailable
	ErrConfigError      = contract.ErrConfigError
	ErrPermissionDenied = contract.ErrPermissionDenied
	ErrValidationFailed = contract.ErrValidationFailed
	ErrBadKey           = contract.ErrBadKey
)

// AppError converts any storage sentinel into the project's canonical
// AppError type with the appropriate HTTP status mapping. Callers at the
// HTTP boundary should use this to surface stable codes (e.g. via
// apperrors.StatusCode / apperrors.CodeOf).
func AppError(err error) *apperrors.AppError {
	var ae apperrors.AppError
	switch {
	case stderrors.Is(err, ErrInvalidFile):
		ae = apperrors.Invalid(CodeInvalidFile, err.Error())
	case stderrors.Is(err, ErrFileTooLarge):
		ae = apperrors.Invalid(CodeFileTooLarge, err.Error())
	case stderrors.Is(err, ErrUnsupportedType):
		ae = apperrors.Invalid(CodeUnsupportedType, err.Error())
	case stderrors.Is(err, ErrValidationFailed):
		ae = apperrors.Invalid(CodeValidationFailed, err.Error())
	case stderrors.Is(err, ErrBadKey):
		ae = apperrors.Invalid(CodeBadKey, err.Error())
	case stderrors.Is(err, ErrNotFound):
		ae = apperrors.NotFound(CodeNotFound, err.Error())
	case stderrors.Is(err, ErrAlreadyExists):
		ae = apperrors.Conflict(CodeAlreadyExists, err.Error())
	case stderrors.Is(err, ErrPermissionDenied):
		ae = apperrors.Forbidden(CodePermissionDenied, err.Error())
	case stderrors.Is(err, ErrConfigError):
		ae = apperrors.Invalid(CodeConfigError, err.Error())
	case stderrors.Is(err, ErrUnavailable):
		ae = apperrors.New(apperrors.TypeInternal, CodeUnavailable, err.Error())
	default:
		ae = apperrors.Internal("storage.error", err)
	}
	return &ae
}
