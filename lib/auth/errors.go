package auth

import (
	"fmt"

	"modular-monolith/lib/errors"
)

// ErrorType represents the category of error
type ErrorType string

const (
	ErrorTypeAuth       ErrorType = "AUTHENTICATION"
	ErrorTypeAuthz      ErrorType = "AUTHORIZATION"
	ErrorTypeValidation ErrorType = "VALIDATION"
	ErrorTypeInternal   ErrorType = "INTERNAL"
)

// NewAuthError Authentication error constructors
func NewAuthError(message string, err error) errors.AppError {
	return errors.Internal(message, err)
}

func InvalidCredentialsError() errors.AppError {
	return errors.Forbidden("invalid_credentials", "invalid username or password")
}

func TokenExpiredError() errors.AppError {
	return errors.AppError{
		Type:    errors.TypeForbidden,
		Code:    "TOKEN_EXPIRED",
		Message: "authentication token has expired",
	}
}

func InvalidTokenError(err error) errors.AppError {
	return errors.AppError{
		Type:    errors.TypeForbidden,
		Code:    "INVALID_TOKEN",
		Message: "invalid authentication token",
		Err:     err,
	}
}

func MissingTokenError() errors.AppError {
	return errors.AppError{
		Type:    errors.TypeForbidden,
		Code:    "MISSING_TOKEN",
		Message: "missing authentication token",
	}
}

func TokenMalformedError() errors.AppError {
	return errors.AppError{
		Type:    errors.TypeForbidden,
		Code:    "TOKEN_MALFORMED",
		Message: "malformed authentication token",
	}
}

func SessionExpiredError() errors.AppError {
	return errors.AppError{
		Type:    errors.TypeForbidden,
		Code:    "SESSION_EXPIRED",
		Message: "session has expired",
	}
}

// Authorization error constructors
func PermissionDeniedError(resource, action string) errors.AppError {
	return errors.AppError{
		Type:    errors.TypeUnauthorized,
		Code:    "PERMISSION_DENIED",
		Message: fmt.Sprintf("permission denied: cannot %s %s", action, resource),
	}
}

func RoleUnauthorizedError(requiredRole string) errors.AppError {
	return errors.AppError{
		Type:    errors.TypeUnauthorized,
		Code:    "ROLE_UNAUTHORIZED",
		Message: fmt.Sprintf("requires role: %s", requiredRole),
	}
}

func ResourceAccessError(resourceID string) errors.AppError {
	return errors.AppError{
		Type:    errors.TypeUnauthorized,
		Code:    "RESOURCE_ACCESS_DENIED",
		Message: fmt.Sprintf("access denied to resource: %s", resourceID),
	}
}

func ScopeError(requiredScopes []string) errors.AppError {
	return errors.AppError{
		Type:    errors.TypeUnauthorized,
		Code:    "INSUFFICIENT_SCOPE",
		Message: fmt.Sprintf("requires scopes: %v", requiredScopes),
	}
}

func AccountLockedError() errors.AppError {
	return errors.AppError{
		Type:    errors.TypeForbidden,
		Code:    "ACCOUNT_LOCKED",
		Message: "account has been locked",
	}
}

func AccountDisabledError() errors.AppError {
	return errors.AppError{
		Type:    errors.TypeForbidden,
		Code:    "ACCOUNT_DISABLED",
		Message: "account has been disabled",
	}
}
