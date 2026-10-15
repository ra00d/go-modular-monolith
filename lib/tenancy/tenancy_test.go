package tenancy

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/google/uuid"
)

func TestResolve(t *testing.T) {
	valid := uuid.NewString()

	tests := []struct {
		name    string
		raw     string
		wantID  uuid.UUID
		wantErr error
	}{
		{name: "valid id", raw: valid, wantID: uuid.MustParse(valid)},
		{name: "padded with spaces", raw: "  " + valid + "  ", wantID: uuid.MustParse(valid)},
		{name: "missing header", raw: "", wantErr: &TenancyError{Code: "MISSING_TENANT", Status: http.StatusBadRequest}},
		{name: "blank header", raw: "   ", wantErr: &TenancyError{Code: "MISSING_TENANT"}},
		{
			name:    "not a uuid",
			raw:     "acme-corp",
			wantErr: &TenancyError{Code: "INVALID_TENANT", Status: http.StatusBadRequest},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Resolve(tt.raw)

			if tt.wantErr != nil {
				if err == nil {
					t.Fatalf("Resolve(%q) = %v, want error", tt.raw, got)
				}
				var te *TenancyError
				if !errors.As(err, &te) {
					t.Fatalf("Resolve(%q) error type = %T, want *TenancyError", tt.raw, err)
				}
				if te.Code != tt.wantErr.(*TenancyError).Code {
					t.Errorf("code = %q, want %q", te.Code, tt.wantErr.(*TenancyError).Code)
				}
				if te.Status != tt.wantErr.(*TenancyError).Status && tt.wantErr.(*TenancyError).Status != 0 {
					t.Errorf("status = %d, want %d", te.Status, tt.wantErr.(*TenancyError).Status)
				}
				return
			}

			if err != nil {
				t.Fatalf("Resolve(%q) unexpected error: %v", tt.raw, err)
			}
			if got.ID != tt.wantID {
				t.Errorf("id = %v, want %v", got.ID, tt.wantID)
			}
		})
	}
}

func TestTenantContext(t *testing.T) {
	ctx := context.Background()

	if _, ok := FromContext(ctx); ok {
		t.Error("FromContext on empty context = ok, want !ok")
	}
	if _, ok := TenantID(ctx); ok {
		t.Error("TenantID on empty context = ok, want !ok")
	}

	id := uuid.New()
	got := WithTenant(ctx, &Tenant{ID: id})

	fromCtx, ok := TenantID(got)
	if !ok || fromCtx != id {
		t.Errorf("TenantID(WithTenant(...)) = %v, %v; want %v, true", fromCtx, ok, id)
	}
	if _, ok := TenantID(ctx); ok {
		t.Error("original context mutated by WithTenant")
	}

	if MustTenantID(got) != id {
		t.Error("MustTenantID returned wrong id")
	}
}

func TestMustTenantIDPanicsWithoutTenant(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("MustTenantID without tenant did not panic")
		}
	}()
	MustTenantID(context.Background())
}
