// Package access holds the access-management domain repository interfaces.
package access

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// ============================================================
// ROLE REPOSITORY
// ============================================================

type RoleRepository interface {
	Create(ctx context.Context, in Role) error
	GetByID(ctx context.Context, id RoleID) (Role, error)
	GetBySlug(ctx context.Context, slug string) (Role, error)
	List(ctx context.Context, limit, offset int) ([]Role, int, error)
	ListByParent(ctx context.Context, parentID RoleID) ([]Role, error)
	Update(ctx context.Context, in *Role) error
	Delete(ctx context.Context, id RoleID) error
	ExistsSlug(ctx context.Context, slug string) (bool, error)
	ExistsID(ctx context.Context, id RoleID) (bool, error)
}

// ============================================================
// PERMISSION REPOSITORY
// ============================================================

type PermissionRepository interface {
	Create(ctx context.Context, in *Permission) error
	GetByID(ctx context.Context, id PermissionID) (Permission, error)
	GetBySlug(ctx context.Context, slug string) (Permission, error)
	List(ctx context.Context, limit, offset int) ([]Permission, int, error)
	ListByResource(ctx context.Context, resource string, limit, offset int) ([]Permission, int, error)
	ListByCategory(ctx context.Context, category string, limit, offset int) ([]Permission, int, error)
	Update(ctx context.Context, in *Permission) error
	Delete(ctx context.Context, id PermissionID) error
	ExistsSlug(ctx context.Context, slug string) (bool, error)
	ExistsID(ctx context.Context, id PermissionID) (bool, error)
}

// ============================================================
// ROLE-PERMISSION GRANT REPOSITORY
// ============================================================

type RolePermissionRepository interface {
	Grant(ctx context.Context, grant *RolePermissionGrant) error
	Revoke(ctx context.Context, roleID RoleID, permissionID PermissionID) error
	Get(ctx context.Context, roleID RoleID, permissionID PermissionID) (*RolePermissionGrant, error)
	ListByRole(ctx context.Context, roleID RoleID, limit, offset int) ([]RolePermission, int, error)
	ListByPermission(ctx context.Context, permissionID PermissionID, limit, offset int) ([]RolePermission, int, error)
	UpdateConditions(
		ctx context.Context,
		roleID RoleID,
		permissionID PermissionID,
		conditions string,
	) error
	Exists(ctx context.Context, roleID RoleID, permissionID PermissionID) (bool, error)
}

// ============================================================
// USER-ROLE ASSIGNMENT REPOSITORY
// ============================================================

type UserRoleRepository interface {
	Assign(ctx context.Context, assignment *UserRoleAssignment) error
	Unassign(ctx context.Context, userID uuid.UUID, roleID RoleID) error
	Get(ctx context.Context, userID uuid.UUID, roleID RoleID) (*UserRoleAssignment, error)
	ListByUser(ctx context.Context, userID uuid.UUID, limit, offset int) ([]UserRole, int, error)
	ListByRole(ctx context.Context, roleID RoleID, limit, offset int) ([]UserRole, int, error)
	ListActiveByUser(ctx context.Context, userID uuid.UUID) ([]UserRole, error)
	UpdateExpiration(
		ctx context.Context,
		userID uuid.UUID,
		roleID RoleID,
		expiresAt *time.Time,
	) error
	SetPrimary(ctx context.Context, userID uuid.UUID, roleID RoleID) error
	ExistsActive(ctx context.Context, userID uuid.UUID, roleID RoleID) (bool, error)
}

// ============================================================
// ABAC POLICY REPOSITORY
// ============================================================

type PolicyRepository interface {
	Create(ctx context.Context, in *Policy) error
	GetByID(ctx context.Context, id PolicyID) (Policy, error)
	GetBySlug(ctx context.Context, slug string) (Policy, error)
	List(ctx context.Context, limit, offset int) ([]Policy, int, error)
	ListByResourceAction(ctx context.Context, resource, action string) ([]Policy, error)
	Update(ctx context.Context, in *Policy) error
	Delete(ctx context.Context, id PolicyID) error
	ExistsSlug(ctx context.Context, slug string) (bool, error)
	ExistsID(ctx context.Context, id PolicyID) (bool, error)
}

// ============================================================
// AUDIT LOG REPOSITORY
// ============================================================

type AuditRepository interface {
	Create(ctx context.Context, in *AuditLog) error
	GetByID(ctx context.Context, id AuditID) (*AuditLog, error)
	List(ctx context.Context, filter AuditFilter, limit, offset int) ([]*AuditLog, int, error)
	Count(ctx context.Context, filter AuditFilter) (int, error)
}

// AuditFilter provides flexible querying of audit logs.
type AuditFilter struct {
	ActorType    string
	ActorID      string
	Action       string
	Category     string
	ResourceType string
	ResourceID   string
	Status       string
	Since        *time.Time
	Until        *time.Time
	RequestID    string
	SessionID    string
}

// ============================================================
// PERMISSION CHANGE LOG REPOSITORY
// ============================================================

type PermissionChangeRepository interface {
	Create(ctx context.Context, in *PermissionChangeLog) error
	GetByID(ctx context.Context, id int32) (*PermissionChangeLog, error)
	List(
		ctx context.Context,
		filter ChangeLogFilter,
		limit, offset int,
	) ([]*PermissionChangeLog, int, error)
}

// ChangeLogFilter provides flexible querying of permission changes.
type ChangeLogFilter struct {
	ChangedBy    uuid.UUID
	ChangeType   string
	TargetType   string
	TargetID     string
	PermissionID PermissionID
	Since        *time.Time
	Until        *time.Time
}

// ============================================================
// SECURITY EVENT REPOSITORY
// ============================================================

type SecurityEventRepository interface {
	Create(ctx context.Context, in *SecurityEvent) error
	GetByID(ctx context.Context, id int64) (SecurityEvent, error)
	List(ctx context.Context, filter SecurityEventFilter, limit, offset int) ([]SecurityEvent, int, error)
	// Resolve stamps the event as triaged. Returns ErrSecurityEventResolved
	// when it was already resolved.
	Resolve(ctx context.Context, id int64, by uuid.UUID) error
}

// ============================================================
// API KEY REPOSITORY
// ============================================================

// IssuedAPIKey is the write-side input for a new key. KeyHash is the only
// secret; the plaintext key is never persisted.
type IssuedAPIKey struct {
	ID          int64
	Name        string
	KeyPrefix   string
	KeyHash     string
	UserID      *uuid.UUID
	RoleID      *RoleID
	Permissions []string
	RateLimit   *int
	ExpiresAt   *time.Time
	CreatedBy   *uuid.UUID
}

type APIKeyRepository interface {
	// Create persists a new key and populates in.ID.
	Create(ctx context.Context, in *IssuedAPIKey) (int64, error)
	GetByID(ctx context.Context, id int64) (APIKey, error)
	GetByPrefix(ctx context.Context, prefix string) (APIKey, error)
	List(ctx context.Context, filter APIKeyFilter, limit, offset int) ([]APIKey, int, error)
	// Update writes the mutable fields of the key. Callers read the current
	// state first and pass a merged value: nil columns are written as NULL.
	Update(ctx context.Context, in APIKey) error
	// Revoke flips is_active off. Returns ErrAPIKeyNotFound when absent.
	Revoke(ctx context.Context, id int64) error
	// TouchLastUsed stamps last_used_at, called by the API key middleware.
	TouchLastUsed(ctx context.Context, id int64) error
}
