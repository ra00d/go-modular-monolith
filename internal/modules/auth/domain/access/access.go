// Package access holds the auth module's authorization domain: roles,
// permissions, ABAC policies and the access audit log. It is a bounded context
// within the auth module, separate from the users context which owns the
// account aggregate.
//
// The package only contains domain logic plus repository contracts that the
// infra layer implements. It depends on nothing outside the standard library,
// google/uuid and lib/auth's policy value objects.
package access

import (
	"time"

	"github.com/google/uuid"
)

// Role status constants
const (
	RoleStatusActive   = "active"
	RoleStatusArchived = "archived"
)

// Standard actions (aligned with migration seed)
const (
	ActionExport = "export"
)

// Standard categories
const (
	CategoryUserManagement = "user-management"
	CategoryContent        = "content"
	CategoryBilling        = "billing"
	CategorySystem         = "system"
)

// ============================================================
// ROLE-PERMISSION GRANT
// ============================================================

// RolePermissionGrant links a role to a permission with optional
// ABAC conditions and expiration.
type RolePermissionGrant struct {
	ID           int32
	RoleID       RoleID
	PermissionID PermissionID
	Conditions   string
	ExpiresAt    *time.Time
	GrantedBy    *uuid.UUID
	CreatedAt    time.Time
}

// ============================================================
// USER-ROLE ASSIGNMENT
// ============================================================

// UserRoleAssignment links a user to a role with optional
// expiration and primary-role flag.
type UserRoleAssignment struct {
	UserID     uuid.UUID
	RoleID     RoleID
	AssignedAt time.Time
	AssignedBy *uuid.UUID
	ExpiresAt  *time.Time
	IsPrimary  bool
}

// ============================================================
// AUDIT LOG
// ============================================================

// AuditLog is an immutable audit trail entry. Actors may be users,
// system components, API keys, or services.
type AuditLog struct {
	ID           AuditID
	CreatedAt    time.Time
	ActorType    string // "user" | "system" | "api_key" | "service"
	ActorID      string
	ActorRole    string
	Action       string // e.g. "role.create", "permission.assign"
	Category     string
	ResourceType string
	ResourceID   string
	IPAddress    string
	UserAgent    string
	SessionID    string
	RequestID    string
	BeforeState  string // JSON
	AfterState   string // JSON
	Metadata     string // JSON
	Status       string // "success" | "failure" | "denied" | "error"
	ErrorMessage string
	Checksum     string
}

var (
// Generic
)

// DomainError wraps a message and implements error.
type DomainError string

func (e DomainError) Error() string { return string(e) }

// NewError creates a typed domain error.
func NewError(msg string) DomainError { return DomainError(msg) }

// Effect is the allow/deny outcome of a policy.
type Effect string

const (
	EffectAllow Effect = "allow"
	EffectDeny  Effect = "deny"
)

// Action performed on a resource. Used for audit entries and policy targets.
const (
	ActionCreate = "create"
	ActionRead   = "read"
	ActionUpdate = "update"
	ActionDelete = "delete"
	ActionList   = "list"
	ActionManage = "manage"
	ActionAssign = "assign"
	ActionRevoke = "revoke"
)

// AuditStatus records the outcome of an audited action.
type AuditStatus string

const (
	AuditStatusSuccess AuditStatus = "success"
	AuditStatusFailure AuditStatus = "failure"
	AuditStatusDenied  AuditStatus = "denied"
	AuditStatusError   AuditStatus = "error"
)

// AuditCategory groups audit entries for filtering.
const (
	CategoryAuth     = "auth"
	CategoryAdmin    = "admin"
	CategoryData     = "data"
	CategorySecurity = "security"
	CategoryGeneral  = "general"
)

// Severity for security events co-emitted with audit entries.
type Severity string

const (
	SeverityInfo     Severity = "info"
	SeverityWarning  Severity = "warning"
	SeverityCritical Severity = "critical"
)

// Domain errors for the access context. Callers compare with errors.Is.
var ()

// Entity identifiers. Roles and permissions use int32 (SERIAL primary keys);
// users are uuid.UUID; assignments are identified by the (user, role) or
// (role, permission) pair.

// RoleID is the primary key of a role.
type RoleID = int32

// PermissionID is the primary key of a permission.
type PermissionID = int32

// PolicyID is the primary key of an ABAC policy.
type PolicyID = int32

// AuditID is the primary key of an audit log entry.
type AuditID = int32

// now is overridable in tests via the package-level clock.
var now = time.Now

// Role is the authorization domain aggregate for a hierarchical role.
type Role struct {
	ID           RoleID
	Name         string
	Slug         string
	Description  string
	ParentRoleID *RoleID
	Level        int32
	IsSystem     bool
	Metadata     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// IsProtected reports whether the role is a seeded system role that must not
// be deleted or have its slug changed.
func (r Role) IsProtected() bool { return r.IsSystem }

// HasParent reports whether the role inherits from another role.
func (r Role) HasParent() bool { return r.ParentRoleID != nil }

// Permission is a granular action a role may be granted on a resource.
type Permission struct {
	ID          PermissionID
	Name        string
	Slug        string
	Description string
	Resource    string
	Action      string
	Category    string
	IsSystem    bool
	Conditions  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// IsProtected reports whether the permission is a seeded system permission.
func (p Permission) IsProtected() bool { return p.IsSystem }

// RolePermission is the granted-permission-on-role relation. It carries
// optional ABAC conditions and an expiration; both are enforced by the policy
// engine, not the database.
type RolePermission struct {
	ID           int32
	RoleID       RoleID
	PermissionID PermissionID
	Conditions   string
	ExpiresAt    *time.Time
	GrantedBy    *uuid.UUID
	CreatedAt    time.Time
}

// IsExpired reports whether the grant has lapse-past its expiry. nil means no
// expiration.
func (rp RolePermission) IsExpired() bool {
	return rp.ExpiresAt != nil && rp.ExpiresAt.Before(now())
}

// UserRole is the assigned-role-to-user relation. It carries the assigning
// admin and an optional expiry.
type UserRole struct {
	UserID     uuid.UUID
	RoleID     RoleID
	AssignedBy *uuid.UUID
	AssignedAt time.Time
	ExpiresAt  *time.Time
	IsPrimary  bool
}

// IsExpired reports whether the assignment has lapsed.
func (ur UserRole) IsExpired() bool {
	return ur.ExpiresAt != nil && ur.ExpiresAt.Before(now())
}

// Policy is the authorization domain aggregate for an ABAC policy. The rules
// field is a JSON document shaped after lib/auth.RuleSet, kept as a raw string
// so the domain does not import the engine's value objects transitively here.
type Policy struct {
	ID          PolicyID
	Name        string
	Slug        string
	Description string
	Resource    string
	Action      string
	Rules       map[string]any
	Effect      Effect
	Priority    int32
	IsActive    bool
	ExpiresAt   *time.Time
	CreatedBy   *uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// IsExpired reports whether the policy has lapsed.
func (p Policy) IsExpired() bool {
	return p.ExpiresAt != nil && p.ExpiresAt.Before(now())
}

// IsDeny reports whether the policy is an explicit deny.
func (p Policy) IsDeny() bool { return p.Effect == EffectDeny }

// AuditEntry is an immutable record of a security-relevant action. It is the
// source of truth for the access-log read model exposed by the AccessRead API.
type AuditEntry struct {
	ID           AuditID
	CreatedAt    time.Time
	ActorType    string
	ActorID      string
	ActorRole    string
	Action       string
	Category     string
	ResourceType string
	ResourceID   string
	IPAddress    string
	UserAgent    string
	RequestID    string
	SessionID    string
	BeforeState  string
	AfterState   string
	Metadata     string
	Status       AuditStatus
	ErrorMessage string
}

// PermissionChangeLog is the domain entity for the permission_change_log table.
type PermissionChangeLog struct {
	ID                 int32
	CreatedAt          time.Time
	ChangedBy          uuid.UUID
	ChangeType         string // "grant", "revoke", "modify"
	TargetType         string // "role", "user", "api_key"
	TargetID           string // UUID or stringified int
	PermissionID       PermissionID
	PreviousConditions string
	NewConditions      string
}

// SecurityEvent is one row of the security_event table: the high-priority
// subset of the audit trail (logins, lockouts, MFA changes, denials).
type SecurityEvent struct {
	ID         int64          `json:"id"          dc:"internal event id"`
	EventType  string         `json:"eventType"   dc:"login_success, login_failure, account_locked, mfa_enabled, ..."`
	UserID     *uuid.UUID     `json:"userId"      dc:"the user the event belongs to, nil for system wide events"`
	Severity   string         `json:"severity"    dc:"info, warning or critical"`
	Details    map[string]any `json:"details"     dc:"JSON payload of the event"`
	IPAddress  string         `json:"ipAddress"   dc:"client address the event came from"`
	ResolvedAt *time.Time     `json:"resolvedAt"  dc:"when an operator resolved it"`
	ResolvedBy *uuid.UUID     `json:"resolvedBy"  dc:"who resolved it"`
	CreatedAt  time.Time      `json:"createdAt"   dc:"when the event was recorded"`
}

// IsResolved reports whether an operator has already triaged the event.
func (e SecurityEvent) IsResolved() bool { return e.ResolvedAt != nil }

// SecurityEventFilter narrows a security event query. Every field is optional.
type SecurityEventFilter struct {
	EventType  string
	Severity   string
	UserID     *uuid.UUID
	Unresolved *bool
	Since      *time.Time
	Until      *time.Time
}

// APIKey is the safe projection of an issued key: it never carries the hash.
type APIKey struct {
	ID          int64      `json:"id"           dc:"internal key id"`
	Name        string     `json:"name"         dc:"human label"`
	KeyPrefix   string     `json:"keyPrefix"    dc:"first characters of the key, enough to recognise it"`
	UserID      *uuid.UUID `json:"userId"       dc:"owning user, nil for a service account"`
	RoleID      *RoleID    `json:"roleId"       dc:"role the key acts under"`
	Permissions []string   `json:"permissions"  dc:"permission slugs overriding the role grants"`
	RateLimit   *int       `json:"rateLimit"    dc:"requests per minute"`
	ExpiresAt   *time.Time `json:"expiresAt"    dc:"when the key stops working"`
	LastUsedAt  *time.Time `json:"lastUsedAt"   dc:"last time the key authenticated a request"`
	IsActive    bool       `json:"isActive"     dc:"false once revoked or expired"`
	CreatedBy   *uuid.UUID `json:"createdBy"    dc:"who issued the key"`
	CreatedAt   time.Time  `json:"createdAt"    dc:"when the key was issued"`
}

// IsUsable reports whether the key may authenticate a request right now.
func (k APIKey) IsUsable() bool {
	if !k.IsActive {
		return false
	}
	return k.ExpiresAt == nil || k.ExpiresAt.After(now())
}

// APIKeyFilter narrows an api key query. Every field is optional.
type APIKeyFilter struct {
	UserID   *uuid.UUID
	RoleID   *RoleID
	IsActive *bool
}
