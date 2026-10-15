package users

import (
	"context"

	"modular-monolith/lib/errors"

	"github.com/google/uuid"
)

// Domain errors for the auth module. Callers compare with errors.Is.
var (
	ErrUserNotFoundCode      = "user.not_found"
	ErrUserAlreadyExistsCode = "user.already_exists"
	ErrInvalidInputCode      = "user.invalid_input"
	ErrInvalidPasswordCode   = "user.invalid_password"
	ErrUserNotActiveCode     = "user.not_action"
	ErrAccountLockedCode     = "user.account_locked"

	ErrUserNotFound      = errors.NotFound(ErrUserNotFoundCode, "user not found")
	ErrUserAlreadyExists = errors.Conflict(ErrUserAlreadyExistsCode, "user already exists")
	ErrInvalidInput      = errors.Invalid(ErrInvalidInputCode, "Invalid input")
	ErrInvalidPassword   = errors.Forbidden(ErrInvalidInputCode, "invalid username or password")
	ErrUserNotActive     = errors.Forbidden(ErrUserNotActiveCode, "user is not active")
	ErrAccountLocked     = errors.Forbidden(ErrAccountLockedCode, "account is locked")

	ErrConnectionNotFoundCode  = "account.connection.not_found"
	ErrConnectionProtectedCode = "account.connection.protected"
	ErrConnectionNotFound      = errors.NotFound(ErrConnectionNotFoundCode, "no linked account for this provider")
	ErrConnectionProtected     = errors.Conflict(
		ErrConnectionProtectedCode,
		"cannot unlink the only sign-in method on this account",
	)

	ErrAttributeNotFoundCode = "user.attribute.not_found"
	ErrNoPrimaryRoleCode     = "user.role.no_primary"

	ErrAttributeNotFound = errors.NotFound(ErrAttributeNotFoundCode, "user attribute not found")
	ErrNoPrimaryRole     = errors.Invalid(ErrNoPrimaryRoleCode, "user has no primary role")
)

// AccountRepository is the single persistence contract for the auth module.
// Its implementation lives in infra.
type AccountRepository interface {
	// Create persists a new user. On success in.ID is populated.
	Create(ctx context.Context, in *User) error

	// GetByIdentifier finds a non-deleted user by email, phone, or username.
	// Returns ErrUserNotFound when no match exists.
	GetByIdentifier(ctx context.Context, identifier string) (User, error)

	// GetByID finds a non-deleted user by primary key.
	GetByID(ctx context.Context, id uuid.UUID) (User, error)

	// Exists reports whether a non-deleted user already uses the identifier
	// in the given column ("email", "phone", or "username").
	Exists(ctx context.Context, column, value string) (bool, error)

	// RecordSuccessfulLogin resets the failure counter and stamps last login.
	RecordSuccessfulLogin(ctx context.Context, id uuid.UUID, ip string) error

	// RecordFailedLogin increments the failure counter, locking the account
	// once the threshold is reached.
	RecordFailedLogin(ctx context.Context, id uuid.UUID) error

	// GetRoles returns the role slugs currently assigned to the user.
	GetRoles(ctx context.Context, id uuid.UUID) ([]string, error)

	// GetPermissions returns the distinct permission slugs the user holds
	// through their roles.
	GetPermissions(ctx context.Context, id uuid.UUID) ([]string, error)

	// GetPreference returns the user preference
	GetPreference(ctx context.Context, id uuid.UUID) (map[string]any, error)

	// UpdatePreference updates the user preference
	UpdatePreference(ctx context.Context, id uuid.UUID, in map[string]any) (map[string]any, error)

	// ListConnections returns the external identities linked to the user,
	// without the stored access or refresh tokens.
	ListConnections(ctx context.Context, id uuid.UUID) ([]Connection, error)

	// UnlinkConnection removes one linked identity. Returns
	// ErrConnectionNotFound when the user has no such link.
	UnlinkConnection(ctx context.Context, id uuid.UUID, provider string) error
}
