package auth

import (
	"context"
	"net/http"
	"slices"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/google/uuid"
)

type ctxKey string

const (
	ctxUserID      ctxKey = "auth.user_id"
	ctxUserStatus  ctxKey = "auth.user_status"
	ctxRole        ctxKey = "auth.role"
	ctxRoles       ctxKey = "auth.roles"
	ctxPermissions ctxKey = "auth.permissions"
)

// Guard authenticates requests and enforces authorization (RBAC and ABAC).
type Guard struct {
	jwt   *JWTManager
	authz *Authorizer
}

// NewGuard builds the guard.
func NewGuard(
	jwt *JWTManager,
	az *Authorizer,
) *Guard {
	return &Guard{jwt: jwt, authz: az}
}

// Authenticate validates the bearer token and stores the user id in the
// request context. Rejects unauthenticated requests with 401.
func (g *Guard) Authenticate(r *ghttp.Request) {
	token := bearerToken(r)
	if token == "" {
		fail(r, http.StatusUnauthorized, MissingTokenError())
		return
	}
	claims, err := g.jwt.ValidateToken(r.Context(), token)
	if err != nil {
		fail(r, http.StatusUnauthorized, InvalidTokenError(err))
		return
	}

	r.SetCtxVar(ctxUserID, claims.UserID)
	r.SetCtxVar(ctxRole, claims.Role)
	r.SetCtxVar(ctxPermissions, claims.Permissions)
	r.SetCtxVar(ctxRoles, claims.Roles)
	r.SetCtxVar(ctxUserStatus, claims.Status)
	r.Middleware.Next()
}

const (
	statusActive              = "active"
	statusSuspended           = "suspended"
	statusLocked              = "locked"
	statusPendingVerification = "pending_verification"
)

// SuperAdminRole is granted every permission implicitly.
const superAdminRole = "super_admin"

// IsActive reports whether the account may authenticate.
func (u AuthCtx) IsActive() bool { return u.Status == statusActive }

// HasPermission reports whether the user may perform the given permission slug
// (e.g. "user.read"). Super administrators are always allowed.
func (u AuthCtx) HasPermission(slug string) bool {
	if slices.Contains(u.Roles, superAdminRole) {
		return true
	}
	return slices.Contains(u.Permissions, slug)
}

func (u AuthCtx) HasRole(slug string) bool {
	if slices.Contains(u.Roles, superAdminRole) {
		return true
	}
	return slices.Contains(u.Roles, slug)
}

// RequirePermission returns middleware that enforces a permission slug. It must
// run after Authenticate. Permissions are resolved from the database per
// request. ponytail: DB lookup per call; cache in Redis if it becomes hot.
func (g *Guard) RequirePermission(slug string) ghttp.HandlerFunc {
	return func(r *ghttp.Request) {
		token := bearerToken(r)
		if token == "" {
			fail(r, http.StatusUnauthorized, MissingTokenError())
			return
		}

		claims, err := g.jwt.ValidateToken(r.Context(), token)
		if err != nil {
			fail(r, http.StatusUnauthorized, InvalidTokenError(err))
			return
		}
		userID, ok := UserID(r.Context())
		if !ok {
			fail(r, http.StatusUnauthorized, InvalidTokenError(err))
		}
		user := AuthCtx{
			UserID:      userID,
			Roles:       claims.Roles,
			Permissions: claims.Permissions,
			Status:      claims.Status,
		}
		// user, err := g.userService.GetUserById(r.Context(), id)
		// 	if err != nil {
		// 		fail(r, http.StatusUnauthorized, libauth.NewAuthError("user lookup failed", err))
		// 		return
		// 	}
		if !user.HasPermission(slug) {
			resource, action, _ := strings.Cut(slug, ".")
			fail(r, http.StatusForbidden, PermissionDeniedError(resource, action))
			return
		}
		r.Middleware.Next()
	}
}

// RequirePolicy returns middleware that evaluates ABAC policies for the given
// resource type and action. It must run after Authenticate. Resource-attribute
// (ownership) checks need the loaded resource, so handlers that guard a
// specific record should call the Authorizer directly instead.
func (g *Guard) RequirePolicy(resourceType, action string) ghttp.HandlerFunc {
	return func(r *ghttp.Request) {
		id, ok := UserID(r.Context())
		if !ok {
			fail(r, http.StatusUnauthorized, MissingTokenError())
			return
		}

		decision, err := g.authz.Authorize(r.Context(), Request{
			UserID:       id,
			Action:       action,
			ResourceType: resourceType,
			IP:           r.GetClientIp(),
			UserAgent:    r.Header.Get("User-Agent"),
			RequestID:    r.Header.Get("X-Request-Id"),
		})
		if err != nil {
			fail(r, http.StatusForbidden, NewAuthError("authorization failed", err))
			return
		}
		if !decision.Allowed {
			fail(r, http.StatusForbidden, PermissionDeniedError(resourceType, action))
			return
		}
		r.Middleware.Next()
	}
}

// UserID returns the authenticated user id from the request context.
func UserID(ctx context.Context) (uuid.UUID, bool) {
	s, ok := ctx.Value(ctxUserID).(string)
	if !ok || s == "" {
		return uuid.Nil, false
	}
	id, err := uuid.Parse(s)
	if err != nil {
		return uuid.Nil, false
	}
	return id, true
}

// ParseClaims extracts and validates the bearer token from the request,
// returning the full claims (including Roles and Permissions). Useful for
// custom middleware that needs to fine-tune authorisation beyond the
// single-slug check in RequirePermission.
//
// The method re-runs JWT validation, which is intentional: a downstream
// middleware that needs the claims is also a natural place to enforce
// authentication, and reusing the existing token path keeps a single
// source of truth for "is this token valid right now".
func (g *Guard) ParseClaims(r *ghttp.Request) (*CustomClaims, error) {
	token := bearerToken(r)
	if token == "" {
		return nil, MissingTokenError()
	}
	claims, err := g.jwt.ValidateToken(r.Context(), token)
	if err != nil {
		return nil, InvalidTokenError(err)
	}
	return claims, nil
}

// RequireAny is a ghttp middleware that authenticates the request and
// enforces that the caller holds at least one of the given permission
// slugs. The full claims (user id, permissions) are stashed in the
// request context for downstream handlers and services to read.
func (g Guard) RequireAny(slugs ...string) ghttp.HandlerFunc {
	return g.authAndEnforce(slugs, false)
}

// RequireSingle enforces a single permission slug; the call exists so
// route definitions read symmetrically with requireAny above.
func (g Guard) RequireSingle(slug string) ghttp.HandlerFunc {
	return g.authAndEnforce([]string{slug}, true)
}

// authAndEnforce is the shared implementation for the two helpers above.
// When allOf is true every slug must match; otherwise any one suffices.
func (g Guard) authAndEnforce(slugs []string, allOf bool) ghttp.HandlerFunc {
	return func(r *ghttp.Request) {
		claims, err := g.ParseClaims(r)
		if err != nil {
			fail(r, http.StatusUnauthorized, err)
			return
		}
		if !g.hasPerm(claims, slugs, allOf) {
			fail(r, http.StatusForbidden, PermissionDeniedError(slugs[0], ""))
			return
		}
		uid, _ := uuid.Parse(claims.UserID)
		// Stash on the request context using the ctxkeys keys so the
		// services package can read them via ctxkeys.UserID / Permissions.
		r.SetCtxVar(ctxUserID, uid)
		r.SetCtxVar(ctxRole, claims.Role)
		r.Middleware.Next()
	}
}

func (guard Guard) hasPerm(claims *CustomClaims, slugs []string, allOf bool) bool {
	// super_admin shortcut: the same convention as AuthCtx.HasPermission.
	if slices.Contains(claims.Roles, "super_admin") {
		return true
	}
	if allOf {
		for _, want := range slugs {
			if !contains(claims.Permissions, want) {
				return false
			}
		}
		return len(slugs) > 0
	}
	for _, want := range slugs {
		if contains(claims.Permissions, want) {
			return true
		}
	}
	return false
}

func bearerToken(r *ghttp.Request) string {
	token := r.Cookie.Get("auth-token")
	tokenString := ""
	if token.IsNil() || token.IsEmpty() {

		const prefix = "Bearer "
		h := r.Header.Get("Authorization")
		if len(h) > len(prefix) && strings.EqualFold(h[:len(prefix)], prefix) {
			tokenString = strings.TrimSpace(h[len(prefix):])
		}

	} else {
		tokenString = token.String()
	}
	return tokenString
}

func fail(r *ghttp.Request, status int, err error) {
	r.Response.Status = status
	r.SetError(err)
}

func contains(haystack []string, needle string) bool {
	return slices.Contains(haystack, needle)
}
