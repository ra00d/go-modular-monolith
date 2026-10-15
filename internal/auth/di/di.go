// Package di
package di

import (
	"modular-monolith/internal/auth/app/services"
	"modular-monolith/lib/auth"
	"modular-monolith/lib/http"

	"github.com/gogf/gf/v2/net/ghttp"
	"go.uber.org/fx"
)

type AuthDepsIn struct {
	fx.In
	// APIGrp is a group wrapping all routes. under one prefix the default is "/api"
	APIGrp             *ghttp.RouterGroup
	Srv                *ghttp.Server
	Jwt                *auth.JWTManager
	Guard              *auth.Guard
	Az                 *auth.Authorizer
	BaseController     *http.BaseController
	AccountService     *services.AccountService
	UsersService       *services.UsersService
	RolesService       *services.RolesService
	PermissionsService *services.PermissionsService
	PoliciesService    *services.PoliciesService
	AuditService       *services.AuditService
	APIKeysService     *services.APIKeyService
}
