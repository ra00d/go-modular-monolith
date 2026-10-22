// Package rest registers the auth module's HTTP controllers.
package rest

import (
	"modular-monolith/internal/modules/auth/api/rest/abac"
	"modular-monolith/internal/modules/auth/api/rest/account"
	"modular-monolith/internal/modules/auth/api/rest/apikeys"
	"modular-monolith/internal/modules/auth/api/rest/audit"
	"modular-monolith/internal/modules/auth/api/rest/authentication"
	"modular-monolith/internal/modules/auth/api/rest/rbac"
	"modular-monolith/internal/modules/auth/api/rest/users"
	authDi "modular-monolith/internal/modules/auth/di"
)

// RegisterRoutes mounts the auth endpoints under /auth. Public endpoints
// (sign-up, sign-in) are open; account, admin, audit and key endpoints
// require a valid bearer token.
func RegisterRoutes(
	deps authDi.AuthDepsIn,
) {
	authentication.RegisterControllers(deps)
	account.RegisterControllers(deps)
	users.RegisterControllers(deps)
	abac.RegisterControllers(deps)
	rbac.RegisterControllers(deps)
	audit.RegisterControllers(deps)
	apikeys.RegisterControllers(deps)
}
