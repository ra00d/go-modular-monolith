// Package audit exposes the read side of the security record: the audit trail
// and the security events emitted alongside it. Both are read-only here;
// writes happen in the services that perform the audited action.
package audit

import (
	"modular-monolith/internal/auth/app/services"
	"modular-monolith/internal/auth/di"
	"modular-monolith/lib/http"

	"github.com/gogf/gf/v2/net/ghttp"
)

// AuditController serves the admin-only read endpoints.
type AuditController struct {
	*http.BaseController
	auditService *services.AuditService
}

func RegisterControllers(deps di.AuthDepsIn) {
	audit := &AuditController{
		BaseController: deps.BaseController,
		auditService:   deps.AuditService,
	}
	guard := deps.Guard

	deps.APIGrp.Middleware(guard.Authenticate).Group("/admin", func(group *ghttp.RouterGroup) {
		group.Group("/audit-logs", func(ag *ghttp.RouterGroup) {
			ag.
				Bind(audit.ListAuditLogs).
				Bind(audit.GetAuditLog)
		})
		group.Group("/security-events", func(sg *ghttp.RouterGroup) {
			sg.
				Bind(audit.ListSecurityEvents).
				Bind(audit.GetSecurityEvent).
				Bind(audit.ResolveSecurityEvent)
		})
	})
}
