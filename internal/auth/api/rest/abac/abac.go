// Package abac group all controllers related to attribute based access control
package abac

import (
	"modular-monolith/internal/auth/app/services"
	"modular-monolith/internal/auth/di"
	"modular-monolith/lib/http"

	"github.com/gogf/gf/v2/net/ghttp"
)

type RBACController struct {
	*http.BaseController
}
type PoliciesController struct {
	*http.BaseController
	polciesService *services.PoliciesService
}

func RegisterControllers(deps di.AuthDepsIn) {
	grp := deps.APIGrp
	guard := deps.Guard
	plctlr := PoliciesController{
		deps.BaseController,
		deps.PoliciesService,
	}
	grp.Middleware(guard.Authenticate).Group("/admin", func(group *ghttp.RouterGroup) {
		group.Group("/policies", func(p *ghttp.RouterGroup) {
			p.
				Bind(plctlr.List).
				Bind(plctlr.Create).
				Bind(plctlr.GetByID).
				Bind(plctlr.Update).
				Bind(plctlr.Delete).
				Bind(plctlr.Evaluate)
		})
	})
}
