package tenancy

import (
	"fmt"
	"net/http"
)

// TenancyError is a typed tenancy failure carrying an HTTP status.
type TenancyError struct {
	Code    string
	Message string
	Status  int
}

func (e *TenancyError) Error() string { return e.Code + ": " + e.Message }

func MissingTenantError() *TenancyError {
	return &TenancyError{
		Code:    "MISSING_TENANT",
		Message: fmt.Sprintf("missing %s header", HeaderTenantID),
		Status:  http.StatusBadRequest,
	}
}

func InvalidTenantError(id string) *TenancyError {
	return &TenancyError{
		Code:    "INVALID_TENANT",
		Message: fmt.Sprintf("invalid tenant identifier: %q", id),
		Status:  http.StatusBadRequest,
	}
}

func TenantMismatchError() *TenancyError {
	return &TenancyError{
		Code:    "TENANT_MISMATCH",
		Message: "authenticated principal does not belong to the requested tenant",
		Status:  http.StatusForbidden,
	}
}
