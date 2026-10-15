package storage

import (
	stderrors "errors"
	"net/http"
	"testing"

	apperrors "modular-monolith/lib/errors"
)

func TestAppErrorMapping(t *testing.T) {
	cases := []struct {
		err      error
		wantCode string
		wantHTTP int
	}{
		{ErrInvalidFile, CodeInvalidFile, http.StatusBadRequest},
		{ErrFileTooLarge, CodeFileTooLarge, http.StatusBadRequest},
		{ErrUnsupportedType, CodeUnsupportedType, http.StatusBadRequest},
		{ErrValidationFailed, CodeValidationFailed, http.StatusBadRequest},
		{ErrBadKey, CodeBadKey, http.StatusBadRequest},
		{ErrNotFound, CodeNotFound, http.StatusNotFound},
		{ErrAlreadyExists, CodeAlreadyExists, http.StatusConflict},
		{ErrPermissionDenied, CodePermissionDenied, http.StatusForbidden},
		{ErrConfigError, CodeConfigError, http.StatusBadRequest},
		{ErrUnavailable, CodeUnavailable, http.StatusInternalServerError},
	}
	for _, c := range cases {
		t.Run(c.err.Error(), func(t *testing.T) {
			ae := AppError(c.err)
			if ae.Code != c.wantCode {
				t.Errorf("code = %q, want %q", ae.Code, c.wantCode)
			}
			if got := apperrors.StatusCode(ae); got != c.wantHTTP {
				t.Errorf("status = %d, want %d", got, c.wantHTTP)
			}
		})
	}
}

func TestAppErrorWraps(t *testing.T) {
	root := stderrors.New("boom")
	ae := AppError(root)
	if ae.Err == nil || ae.Err.Error() != "boom" {
		t.Errorf("expected wrapped cause, got %v", ae.Err)
	}
	if !stderrors.Is(ae, root) {
		t.Error("expected original error to be in chain")
	}
}

func TestAppErrorPlainBecomesInternal(t *testing.T) {
	plain := stderrors.New("kaboom")
	ae := AppError(plain)
	if got := apperrors.StatusCode(ae); got != http.StatusInternalServerError {
		t.Errorf("plain error status = %d, want 500", got)
	}
}