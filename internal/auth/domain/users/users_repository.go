package users

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// UserRepository is the admin-side persistence contract for the user
// aggregate. Authentication reads live on [AccountRepository].
type UserRepository interface {
	GetAll(ctx context.Context, limit, offset int) ([]User, int, error)
	Create(ctx context.Context, in CreateUserInput) (uuid.UUID, error)

	// Update applies a partial update. Nil fields are left untouched. Replacing
	// an identifier clears its verified flag, because the address changed.
	Update(ctx context.Context, id uuid.UUID, in UpdateUserInput) error

	// Delete soft deletes the user.
	Delete(ctx context.Context, id uuid.UUID) error

	// SetStatus moves the user to another account status. banUntil is only
	// meaningful for a locked account: nil lifts the lock, a timestamp makes
	// it expire on its own.
	SetStatus(ctx context.Context, id uuid.UUID, status string, banUntil *time.Time) error

	// SetPasswordHash replaces the stored hash.
	SetPasswordHash(ctx context.Context, id uuid.UUID, hash string) error

	// RBAC
	AssignRole(ctx context.Context, userID uuid.UUID, roleID int) error
	AssignPermission(ctx context.Context, userID uuid.UUID, permissionID int) error
	AssignPermissions(ctx context.Context, userID uuid.UUID, permissionIDs []int) error

	// ABAC
	ListAttributes(
		ctx context.Context,
		userID uuid.UUID,
		namespace *string,
		limit, offset int,
	) ([]UserAttribute, int, error)
	// SetAttribute upserts one attribute keyed by (user, namespace, key).
	SetAttribute(ctx context.Context, in UserAttribute) error
	DeleteAttribute(ctx context.Context, userID uuid.UUID, namespace, key string) error
}

// UpdateUserInput is a partial update of the user aggregate. Every pointer
// field is optional; nil means "leave as is".
type UpdateUserInput struct {
	Name     *string
	Image    *string
	Email    *string
	Phone    *string
	Username *string
}
