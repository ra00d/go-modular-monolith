// Package rbac
package rbac

import (
	"modular-monolith/internal/auth/app/services"
	"modular-monolith/internal/auth/di"
	"modular-monolith/lib/http"

	"github.com/gogf/gf/v2/net/ghttp"
)

type PermissionStr string

type baseCntrl struct {
	*http.BaseController
	usersService       *services.UsersService
	rolesService       *services.RolesService
	permissionsService *services.PermissionsService
}
type RolesCntrl struct {
	baseCntrl
}

type PermissionsController struct {
	baseCntrl
}

func RegisterControllers(deps di.AuthDepsIn) {
	grp := deps.APIGrp
	baseController := deps.BaseController
	srv := deps.UsersService
	guard := deps.Guard
	bCntrl := baseCntrl{
		BaseController:     baseController,
		usersService:       srv,
		rolesService:       deps.RolesService,
		permissionsService: deps.PermissionsService,
	}
	rctlr := &RolesCntrl{
		baseCntrl: bCntrl,
	}
	pctlr := &PermissionsController{
		baseCntrl: bCntrl,
	}

	grp.Middleware(guard.Authenticate).Group("/admin", func(group *ghttp.RouterGroup) {
		group.Group("/roles", func(rg *ghttp.RouterGroup) {
			rg.
				Bind(rctlr.ListRoles).
				Bind(rctlr.CreateRole).
				Bind(rctlr.GetRole).
				Bind(rctlr.UpdateRole).
				Bind(rctlr.DeleteRole).
				Bind(rctlr.ListRolePermissions).
				Bind(rctlr.AssignRolePermission).
				Bind(rctlr.RevokeRolePermission)
		})
		group.Group("/permissions", func(pg *ghttp.RouterGroup) {
			pg.
				Bind(pctlr.ListPermissions).
				Bind(pctlr.CreatePermission).
				Bind(pctlr.GetPermission).
				Bind(pctlr.UpdatePermission).
				Bind(pctlr.DeletePermission)
		})
	})
}
