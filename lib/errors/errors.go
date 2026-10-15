// Package errors provides the application-wide typed error: a stable code,
// an error category with an HTTP status mapping, and cause wrapping.
// Domain packages keep their own sentinel errors; map them here at the edge.
package errors

import (
	stderrors "errors"
	"fmt"
	"net/http"
)

// Type categorizes an application error and drives its HTTP status.
type Type string

const (
	TypeNotFound     Type = "not_found"
	TypeInvalid      Type = "invalid_input"
	TypeUnauthorized Type = "unauthorized"
	TypeForbidden    Type = "forbidden"
	TypeConflict     Type = "conflict"
	TypeInternal     Type = "internal"
	// TypeNotImplemented marks a declared but unimplemented endpoint. Handlers
	// return it instead of a half-built response so callers see 501, not 500.
	TypeNotImplemented Type = "not_implemented"
)

var statusByType = map[Type]int{
	TypeNotFound:       http.StatusNotFound,
	TypeInvalid:        http.StatusBadRequest,
	TypeUnauthorized:   http.StatusUnauthorized,
	TypeForbidden:      http.StatusForbidden,
	TypeConflict:       http.StatusConflict,
	TypeInternal:       http.StatusInternalServerError,
	TypeNotImplemented: http.StatusNotImplemented,
}

// AppError is the canonical error crossing module boundaries.
type AppError struct {
	Type    Type   // category
	Code    string // stable machine code, e.g. "plan.not_found"
	Message string // human-readable message (i18n key by convention)
	Err     error  // wrapped cause, if any
}

func (e AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s [%s]: %s: %v", e.Type, e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%s [%s]: %s", e.Type, e.Code, e.Message)
}

func (e AppError) Unwrap() error { return e.Err }

// Is compares by Code; two AppErrors match when codes are equal.
func (e AppError) Is(target error) bool {
	t, ok := target.(AppError)
	return ok && t.Code == e.Code
}

// StatusCode maps the error's type to its HTTP status.
// map to 500 via [StatusCode].
func (e AppError) StatusCode() int {
	return StatusCode(e)
}

// ErrCode returns the machine code of the error ("internal.error" for plain errors).
func (e AppError) ErrCode() string {
	return e.Code
}

// New builds a typed error.
func New(t Type, code, message string) AppError {
	return AppError{Type: t, Code: code, Message: message}
}

// Wrap builds a typed error carrying cause. Wrapping a nil err degrades to New.
func Wrap(err error, t Type, code, message string) AppError {
	if err == nil {
		return New(t, code, message)
	}
	return AppError{Type: t, Code: code, Message: message, Err: err}
}

// WrapRaw returns a wraped error using fmt.Errorf(caller+" %w", err)
// caller it could be the function producing the error
// err the actual error
func WrapRaw(caller string, err error) error {
	return fmt.Errorf(caller+" %w", err)
}

// Convenience constructors for the common categories.

func NotFound(code, message string) AppError {
	return New(TypeNotFound, code, message)
}

func Invalid(code, message string) AppError {
	return New(TypeInvalid, code, message)
}

func Unauthorized(code, message string) AppError {
	return New(TypeUnauthorized, code, message)
}

func Forbidden(code, message string) AppError {
	return New(TypeForbidden, code, message)
}

func Conflict(code, message string) AppError {
	return New(TypeConflict, code, message)
}

func Internal(message string, err error) AppError {
	return Wrap(err, TypeInternal, "internal.error", message)
}

// NotImplemented marks an endpoint whose handler is still a stub. The code
// should be the endpoint identifier, e.g. "auth.verify_email".
func NotImplemented(code, message string) AppError {
	return New(TypeNotImplemented, code, message)
}

// From coerces any error into an *AppError; plain errors become internal ones.
func From(err error) AppError {
	ae, ok := stderrors.AsType[AppError](err)
	if ok {
		return ae
	}
	return Internal(err.Error(), err)
}

// StatusCode maps the error's type to its HTTP status. Non-AppError inputs
// map to 500 via [From].
func StatusCode(err error) int {
	return statusByType[From(err).Type]
}

// CodeOf returns the machine code of the error ("internal.error" for plain errors).
func CodeOf(err error) string {
	return From(err).Code
}

// ErrInvalidInput Keep the historical sentinel working for existing imports.
var ErrInvalidInput = stderrors.New("invalid input")
