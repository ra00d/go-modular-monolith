// Package tenancy provides tenant resolution, request-scoped tenant context,
// and HTTP middleware for multi-tenant request handling.
//
// Strategy (docs/plan/PLAN.md): shared database with tenant_id columns;
// tenants are addressed via the X-TENANT-ID header. This package owns
// identity resolution and propagation only — persistence (tenant store,
// subscription state) lives in the tenancy module.
package tenancy

import (
	"context"

	"github.com/google/uuid"
)

// HeaderTenantID is the custom request header carrying the caller's tenant id.
const HeaderTenantID = "X-TENANT-ID"

// Tenancy modes (mirrors config.TenancyConfig.Mode).
const (
	ModeSingle    = "single"
	ModeShared    = "shared"
	ModeDedicated = "dedicated"
)

type ctxKey string

const ctxTenant ctxKey = "tenancy.tenant"

// Tenant is the resolved tenant of the current request.
// ponytail: ID only until the tenants table exists; extend with slug/status then.
type Tenant struct {
	ID uuid.UUID
}

// WithTenant returns a copy of ctx carrying t.
func WithTenant(ctx context.Context, t *Tenant) context.Context {
	return context.WithValue(ctx, ctxTenant, t)
}

// FromContext returns the tenant resolved for the current request, if any.
func FromContext(ctx context.Context) (*Tenant, bool) {
	t, ok := ctx.Value(ctxTenant).(*Tenant)
	return t, ok && t != nil
}

// TenantID returns the resolved tenant id, if any.
func TenantID(ctx context.Context) (uuid.UUID, bool) {
	t, ok := FromContext(ctx)
	if !ok {
		return uuid.Nil, false
	}
	return t.ID, true
}

// MustTenantID returns the resolved tenant id or panics. For code paths where
// the middleware guarantees a tenant (protected route groups).
func MustTenantID(ctx context.Context) uuid.UUID {
	id, ok := TenantID(ctx)
	if !ok {
		panic("tenancy: no tenant in context; TenantMiddleware must run first")
	}
	return id
}
