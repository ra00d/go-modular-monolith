// Package users holds the auth module's user domain: entity, repository
// interfaces and the pure authorization helpers.
package users

import (
	"slices"
	"time"

	"github.com/google/uuid"
)

// User is the auth domain aggregate. It carries the resolved role and
// permission slugs so callers do not need to re-query for authorization.
type User struct {
	ID            uuid.UUID      `json:"id"`
	Name          string         `json:"name"`
	Email         string         `json:"email"`
	Phone         string         `json:"phone"`
	Username      string         `json:"username"`
	PasswordHash  string         `json:"-"` // Never serialize the password hash.
	EmailVerified bool           `json:"email_verified"`
	PhoneVerified bool           `json:"phone_verified"`
	Image         string         `json:"image,omitempty"`
	Roles         []string       `json:"roles"`
	Permissions   []string       `json:"permissions"`
	Status        string         `json:"status"`
	Preference    map[string]any `json:"preference"`
	MFAEnabled    bool           `json:"mfa_enabled"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

// Account status values (see the "status" CHECK constraint on the user table).
const (
	StatusActive              = "active"
	StatusSuspended           = "suspended"
	StatusLocked              = "locked"
	StatusPendingVerification = "pending_verification"
)

// SuperAdminRole is granted every permission implicitly.
const SuperAdminRole = "super_admin"

// IsActive reports whether the account may authenticate.
func (u User) IsActive() bool { return u.Status == StatusActive }

// UserAttribute is one ABAC attribute row (the user_attribute table).
// Attributes feed the policy engine; values are stored as strings and parsed
// according to ValueType.
type UserAttribute struct {
	ID        int64     `json:"id"        dc:"internal attribute id"`
	UserID    uuid.UUID `json:"userId"    dc:"the user the attribute belongs to"`
	Namespace string    `json:"namespace" dc:"org, project, system or default"`
	Key       string    `json:"key"       dc:"attribute name"`
	Value     string    `json:"value"     dc:"the raw stored value"`
	ValueType string    `json:"valueType" dc:"string, number, boolean, json or date"`
	UpdatedAt time.Time `json:"updatedAt" dc:"when the attribute was written"`
}

// DefaultNamespace is used when a caller does not scope an attribute.
const DefaultNamespace = "default"

// Attribute value types, mirroring the CHECK constraint on the column.
const (
	AttrTypeString  = "string"
	AttrTypeNumber  = "number"
	AttrTypeBoolean = "boolean"
	AttrTypeJSON    = "json"
	AttrTypeDate    = "date"
)

// Connection is an external identity linked to the account (the account
// table: provider_id + provider_account_id).
type Connection struct {
	ID                 int64      `json:"id"                 dc:"internal link id"`
	Provider           string     `json:"provider"           dc:"google, github, apple, ..."`
	ProviderAccountID  string     `json:"providerAccountId"  dc:"the identifier the provider assigned"`
	Scopes             []string   `json:"scopes"             dc:"granted scopes"`
	AccessTokenExpires *time.Time `json:"accessTokenExpires" dc:"unix expiry of the stored access token"`
	CreatedAt          time.Time  `json:"createdAt"          dc:"when the link was created"`
}

// HasPermission reports whether the user may perform the given permission slug
// (e.g. "user.read"). Super administrators are always allowed.
func (u User) HasPermission(slug string) bool {
	if slices.Contains(u.Roles, SuperAdminRole) {
		return true
	}
	return slices.Contains(u.Permissions, slug)
}
