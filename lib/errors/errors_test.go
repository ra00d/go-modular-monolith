package errors

import (
	stderrors "errors"
	"net/http"
	"testing"
)

func TestWrapAndUnwrap(t *testing.T) {
	cause := stderrors.New("db: connection refused")
	err := Wrap(cause, TypeNotFound, "plan.missing", "plan not found")

	if !stderrors.Is(err, cause) {
		t.Error("cause chain broken")
	}
	var ae *AppError
	if !stderrors.As(err, &ae) || ae.Code != "plan.missing" {
		t.Errorf("As failed: %+v", ae)
	}
}

func TestIsByCode(t *testing.T) {
	a := NotFound("plan.missing", "plan not found")
	b := Wrap(stderrors.New("x"), TypeNotFound, "plan.missing", "different message")
	c := NotFound("other.code", "plan not found")

	if !stderrors.Is(a, b) {
		t.Error("same code should match")
	}
	if stderrors.Is(a, c) {
		t.Error("different codes must not match")
	}
}

func TestStatusCode(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want int
	}{
		{"not found", NotFound("a.b", "m"), http.StatusNotFound},
		{"invalid", Invalid("a.b", "m"), http.StatusBadRequest},
		{"conflict", Conflict("a.b", "m"), http.StatusConflict},
		{"not implemented", NotImplemented("a.b", "m"), http.StatusNotImplemented},
		{"plain error", stderrors.New("boom"), http.StatusInternalServerError},
		{"wrapped internal", Internal("op failed", stderrors.New("x")), 500},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StatusCode(tt.err); got != tt.want {
				t.Errorf("StatusCode = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestFromPlainError(t *testing.T) {
	plain := stderrors.New("boom")
	got := From(plain)
	if got.Type != TypeInternal || got.Code != CodeOf(plain) {
		t.Errorf("From = %+v", got)
	}
	if From(got) != got {
		t.Error("From on AppError should be identity")
	}
}

func TestWrapNilCause(t *testing.T) {
	e := Wrap(nil, TypeInvalid, "a.b", "msg")
	if e.Err != nil {
		t.Errorf("Err = %v, want nil", e.Err)
	}
	if StatusCode(e) != http.StatusBadRequest {
		t.Errorf("status = %d", StatusCode(e))
	}
}
