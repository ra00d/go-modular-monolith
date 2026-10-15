// Package access holds the auth module's authorization domain events.
package access

import (
	"github.com/google/uuid"
)

// ============================================
// ROLE EVENTS
// ============================================

// RoleCreated is emitted when a new role is created.
type RoleCreated struct {
	RoleID   RoleID
	Name     string
	Slug     string
	ParentID *RoleID
	Level    int
	IsSystem bool
}

// RoleUpdated is emitted when a role is modified.
type RoleUpdated struct {
	RoleID   RoleID
	Name     *string
	Slug     *string
	ParentID *RoleID
	Level    *int
	Metadata *string
}

// RoleDeleted is emitted when a role is soft-deleted.
type RoleDeleted struct {
	RoleID RoleID
}

// RoleAssignedToUser is emitted when a user receives a role.
type RoleAssignedToUser struct {
	UserID     uuid.UUID
	RoleID     RoleID
	AssignedBy *uuid.UUID
	ExpiresAt  *int64
	IsPrimary  bool
}

// RoleRevokedFromUser is emitted when a role is removed from a user.
type RoleRevokedFromUser struct {
	UserID uuid.UUID
	RoleID RoleID
}

// ============================================
// PERMISSION EVENTS
// ============================================

// PermissionCreated is emitted when a new permission is registered.
type PermissionCreated struct {
	PermissionID PermissionID
	Name         string
	Slug         string
	Resource     string
	Action       string
	Category     string
	IsSystem     bool
}

// PermissionUpdated is emitted when a permission is modified.
type PermissionUpdated struct {
	PermissionID PermissionID
	Name         *string
	Slug         *string
	Description  *string
	Category     *string
	Conditions   *string
}

// PermissionDeleted is emitted when a permission is soft-deleted.
type PermissionDeleted struct {
	PermissionID PermissionID
}

// PermissionGrantedToRole is emitted when a role receives a permission.
type PermissionGrantedToRole struct {
	RoleID       RoleID
	PermissionID PermissionID
	Conditions   string
	ExpiresAt    *int64
	GrantedBy    *uuid.UUID
}

// PermissionRevokedFromRole is emitted when a permission is removed from a role.
type PermissionRevokedFromRole struct {
	RoleID       RoleID
	PermissionID PermissionID
	RevokedBy    *uuid.UUID
}

// ============================================
// POLICY EVENTS
// ============================================

// PolicyCreated is emitted when an ABAC policy is created.
type PolicyCreated struct {
	PolicyID  PolicyID
	Name      string
	Slug      string
	Resource  string
	Action    string
	Effect    string
	Priority  int
	CreatedBy *uuid.UUID
}

// PolicyUpdated is emitted when a policy is modified.
type PolicyUpdated struct {
	PolicyID  PolicyID
	Name      *string
	Slug      *string
	Rules     *string
	Effect    *string
	Priority  *int
	IsActive  *bool
	ExpiresAt *int64
}

// PolicyDeleted is emitted when a policy is soft-deleted.
type PolicyDeleted struct {
	PolicyID PolicyID
}

// ============================================
// AUDIT EVENTS
// ============================================

// AccessAudited is emitted for every security-relevant action. It mirrors
// the AuditEntry fields so consumers can react to specific events.
type AccessAudited struct {
	AuditEntry
}

// ============================================
// PERMISSION CHANGE EVENTS
// ============================================

// PermissionChangeLogged is emitted when a permission grant/revoke/modify
// is recorded on a role, user or API key.
type PermissionChangeLogged struct {
	ChangeID           int32
	ChangedBy          uuid.UUID
	ChangeType         string
	TargetType         string
	TargetID           string
	PermissionID       PermissionID
	PreviousConditions string
	NewConditions      string
}
