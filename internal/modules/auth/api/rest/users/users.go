// Package users
package users

import (
	"modular-monolith/internal/modules/auth/app/services"
	"modular-monolith/internal/modules/auth/di"
	"modular-monolith/lib/http"

	"github.com/gogf/gf/v2/net/ghttp"
)

type PermissionStr string

var PermissionGetUser PermissionStr = "user.list"

type baseCntrl struct {
	*http.BaseController
	usersService *services.UsersService
}
type UsersController struct {
	baseCntrl
}

func RegisterControllers(deps di.AuthDepsIn) {
	grp := deps.APIGrp
	baseController := deps.BaseController
	srv := deps.UsersService
	guard := deps.Guard
	bCntrl := baseCntrl{BaseController: baseController, usersService: srv}
	ctlr := &UsersController{
		baseCntrl: bCntrl,
	}

	grp.Middleware(guard.Authenticate).Group("/admin", func(group *ghttp.RouterGroup) {
		group.Group("/users", func(ug *ghttp.RouterGroup) {
			ug.
				Bind(ctlr.Create).
				Bind(ctlr.UpdateUser).
				Bind(ctlr.DeleteUser).
				Bind(ctlr.BanUser).
				Bind(ctlr.UnbanUser).
				Bind(ctlr.SetUserStatus).
				Bind(ctlr.ResetUserPassword).
				Bind(ctlr.ImpersonateUser).
				Bind(ctlr.ListUserRoles).
				Bind(ctlr.AssignUserRole).
				Bind(ctlr.RevokeUserRole).
				Bind(ctlr.ListUserAttributes).
				Bind(ctlr.SetUserAttribute).
				Bind(ctlr.DeleteUserAttribute).
				Map(map[string]any{
					"/":     ctlr.GetAllUsers,
					"/{id}": ctlr.GetUser,
				})
		})
	})
}
