// Package authentication group all related controlers for authentication
package authentication

import (
	"modular-monolith/internal/modules/auth/app/services"
	"modular-monolith/internal/modules/auth/di"
	libauth "modular-monolith/lib/auth"
	"modular-monolith/lib/http"

	"github.com/gogf/gf/v2/net/ghttp"
)

// PublicController serves unauthenticated endpoints.
type PublicController struct {
	*http.BaseController
	userService *services.AccountService
}

// AuthController serves endpoints that require authentication.
type AuthController struct {
	*http.BaseController
	userService *services.AccountService
	jwt         *libauth.JWTManager
}

func RegisterControllers(deps di.AuthDepsIn) {
	grp := deps.APIGrp
	baseController := deps.BaseController
	userService := deps.AccountService
	guard := deps.Guard
	public := &PublicController{
		BaseController: baseController,
		userService:    userService,
	}
	auth := &AuthController{
		BaseController: baseController,
		userService:    userService,
		jwt:            deps.Jwt,
	}
	grp.Group("/auth", func(group *ghttp.RouterGroup) {
		group.Bind(public.SignIn)
		group.Bind(public.SignUp)
		group.Bind(public.VerifyEmail)
		group.Bind(public.VerifyPhone)
		group.Bind(public.ResendVerification)
		// Password recovery cannot require a session: the caller lost it.
		group.Bind(public.ForgotPassword)
		group.Bind(public.VerifyResetPassword)
		group.Bind(public.ResetPassword)
		group.Bind(public.ChallengeMFA)
		group.Bind(public.OAuthAuthorize)
		group.Bind(public.OAuthCallback)
		group.Group("/", func(p *ghttp.RouterGroup) {
			p.Middleware(guard.Authenticate).
				Bind(auth.Refresh).
				Bind(auth.Logout).
				Bind(auth.LogoutAll).
				Bind(auth.ChangePassword).
				Bind(auth.ListSessions).
				Bind(auth.RevokeSession).
				Bind(auth.RevokeAllSessions).
				Bind(auth.EnableMFA).
				Bind(auth.ConfirmMFA).
				Bind(auth.DisableMFA)
		})
	})
}
