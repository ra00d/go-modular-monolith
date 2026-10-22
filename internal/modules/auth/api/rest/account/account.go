// Package account
package account

import (
	"modular-monolith/internal/modules/auth/app/services"
	"modular-monolith/internal/modules/auth/di"
	"modular-monolith/lib/http"

	"github.com/gogf/gf/v2/net/ghttp"
)

// AccountController serves endpoints that require authentication.
type AccountController struct {
	*http.BaseController
	userService *services.AccountService
}

// RegisterControllers register all controllers of account module
func RegisterControllers(deps di.AuthDepsIn) {
	grp := deps.APIGrp
	guard := deps.Guard
	bctl := deps.BaseController
	//auhtz:=deps.Az
	account := AccountController{
		BaseController: bctl,
		userService:    deps.AccountService,
	}

	grp.Group("/account", func(group *ghttp.RouterGroup) {
		group.Middleware(guard.Authenticate).
			Bind(account.Me).
			Bind(account.UpdateProfile).
			Bind(account.ChangeEmail).
			Bind(account.ChangePhone).
			Bind(account.UploadAvatar).
			Bind(account.DeleteAccount).
			Bind(account.GetPreference).
			Bind(account.UpdatePreference).
			Bind(account.ListConnections).
			Bind(account.UnlinkConnection)
	})
}
