package tenancy

import (
	"context"
	"errors"
	"net/http"

	"modular-monolith/lib/config"

	"github.com/gogf/gf/v2/net/ghttp"
)

// TenantVerifier validates the resolved tenant against request state, e.g.
// once JWT claims carry tenant_id: compare against auth.TenantID(ctx) and
// return TenantMismatchError() on mismatch. Optional; nil skips verification.
type TenantVerifier func(ctx context.Context, t *Tenant) error

// globalVerifier is the process-wide verifier applied when a call site
// passes nil. The companies module sets this so every existing
// TenantMiddleware(cfg, nil) call site gains a live-company check without
// having to be edited.
var globalVerifier TenantVerifier

// SetGlobalVerifier installs the process-wide verifier. Replaces any
// previous value; safe to call once during fx Invoke.
func SetGlobalVerifier(v TenantVerifier) { globalVerifier = v }

// TenantMiddleware resolves the tenant from X-TENANT-ID and injects it into
// the request context. No-op when tenancy is disabled or in single mode.
//
// Mount it on route groups before controllers:
//
//	group.Middleware(tenancy.TenantMiddleware(cfg.Tenancy), guard.Authenticate)
//
// Read it downstream with tenancy.TenantID(ctx) / MustTenantID(ctx).
func TenantMiddleware(cfg config.TenancyConfig, verify TenantVerifier) ghttp.HandlerFunc {
	if verify == nil {
		verify = globalVerifier
	}
	return func(r *ghttp.Request) {
		if !cfg.Enabled || cfg.Mode == "" || cfg.Mode == ModeSingle {
			r.Middleware.Next()
			return
		}

		t, err := Resolve(r.Header.Get(HeaderTenantID))
		if err != nil {
			fail(r, err)
			return
		}

		if verify != nil {
			if err := verify(r.Context(), t); err != nil {
				fail(r, err)
				return
			}
		}

		r.SetCtxVar(ctxTenant, t)
		r.Middleware.Next()
	}
}

func fail(r *ghttp.Request, err error) {
	status := http.StatusForbidden
	if te, ok := errors.AsType[*TenancyError](err); ok {
		status = te.Status
	}
	r.Response.Status = status
	r.SetError(err)
}
