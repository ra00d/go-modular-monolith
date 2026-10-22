// Package access holds the auth module's authorization domain errors.
package access

import (
	"errors"
)

// Domain errors for the access bounded context. Callers compare with errors.Is.
var (
	ErrRoleNotFound       = errors.New("role not found")
	ErrRoleAlreadyExists  = errors.New("role already exists")
	ErrRoleSlugTaken      = errors.New("role slug already in use")
	ErrRoleHasChildren    = errors.New("cannot delete role with child roles")
	ErrRoleHasAssignments = errors.New("cannot delete role with user assignments")
	ErrRoleHasPermissions = errors.New("cannot delete role with permission grants")
	ErrRoleIsSystem       = errors.New("cannot modify or delete a system role")
	ErrRoleParentNotFound = errors.New("parent role not found")
	ErrRoleParentCycle    = errors.New("role parent would create a cycle")
	ErrRoleInvalidLevel   = errors.New("role level inconsistent with parent")
	ErrRoleNoPrimary      = errors.New("user must have a primary role")

	ErrPermissionNotFound      = errors.New("permission not found")
	ErrPermissionAlreadyExists = errors.New("permission already exists")
	ErrPermissionSlugTaken     = errors.New("permission slug already in use")
	ErrPermissionIsSystem      = errors.New("cannot modify or delete a system permission")
	ErrPermissionHasGrants     = errors.New("permission is still granted to a role")

	ErrGrantNotFound      = errors.New("role-permission grant not found")
	ErrGrantAlreadyExists = errors.New("role already has this permission")
	ErrGrantExpired       = errors.New("grant has expired")

	ErrAssignmentNotFound      = errors.New("user-role assignment not found")
	ErrAssignmentAlreadyExists = errors.New("user already has this role")
	ErrAssignmentExpired       = errors.New("assignment has expired")

	ErrPolicyNotFound      = errors.New("policy not found")
	ErrPolicyAlreadyExists = errors.New("policy already exists")
	ErrPolicySlugTaken     = errors.New("policy slug already in use")
	ErrPolicyInvalidRules  = errors.New("policy rules JSON is invalid")

	ErrAuditEntryNotFound = errors.New("audit entry not found")
	ErrAuditReadOnly      = errors.New("audit log is append-only")

	ErrSecurityEventNotFound    = errors.New("security event not found")
	ErrSecurityEventResolved    = errors.New("security event is already resolved")
	ErrSecurityEventTypeUnknown = errors.New("unknown security event type")

	ErrAPIKeyNotFound = errors.New("api key not found")
	ErrAPIKeyInactive = errors.New("api key is revoked")
	ErrAPIKeyExpired  = errors.New("api key has expired")
	ErrAPIKeyInUse    = errors.New("api key is assigned and cannot be rotated")

	ErrInvalidInput = errors.New("invalid input")
)
