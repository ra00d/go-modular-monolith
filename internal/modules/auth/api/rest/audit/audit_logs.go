package audit

import (
	"context"
	"time"

	"modular-monolith/internal/modules/auth/domain/access"
	"modular-monolith/lib/http"

	"github.com/gogf/gf/v2/frame/g"
)

type ListAuditLogsReq struct {
	g.Meta `                                                                                                                                                     path:"/" method:"get" tags:"Audit" x-group:"Admin/AuditLogs" summary:"Search the audit trail" security:"bearer"`
	http.RequestBaseParams
	ActorID      string     `                                                                                                                                                                       json:"actorId"     dc:"user id, api key or service name that acted"`
	Action       string     `                                                                                                                                                                       json:"action"      dc:"action slug, e.g. user.login or role.update"`
	Category     string     `                                                                                                                                                                       json:"category"    v:"in:auth,admin,data,security,general" dc:"audit category"`
	ResourceType string     `                                                                                                                                                                       json:"resourceType" dc:"affected resource type"`
	ResourceID   string     `                                                                                                                                                                       json:"resourceId"  dc:"affected resource id"`
	Status       string     `                                                                                                                                                                       json:"status"      v:"in:success,failure,denied,error" dc:"outcome of the audited action"`
	SessionID    string     `                                                                                                                                                                       json:"sessionId"   dc:"session the action belongs to"`
	From         *time.Time `                                                                                                                                                                       json:"from"        dc:"only entries at or after this time"`
	To           *time.Time `                                                                                                                                                                       json:"to"          dc:"only entries at or before this time"`
}

// ListAuditLogsRes is the page of matching entries, newest first. The total is
// carried in the response meta rather than duplicated in the payload.
type ListAuditLogsRes []access.AuditEntry

type GetAuditLogReq struct {
	g.Meta `                                                   path:"/:id" method:"get" tags:"Audit" x-group:"Admin/AuditLogs" summary:"Get one audit entry" security:"bearer"`
	ID     access.AuditID ` v:"required" json:"id" dc:"the audit entry id"`
}
type GetAuditLogRes access.AuditEntry

func (c AuditController) ListAuditLogs(
	ctx context.Context,
	req *ListAuditLogsReq,
) (res http.Response[ListAuditLogsRes], err error) {
	p := c.Pagination(ctx)
	entries, count, err := c.auditService.List(ctx, access.AuditFilter{
		ActorID:      req.ActorID,
		Action:       req.Action,
		Category:     req.Category,
		ResourceType: req.ResourceType,
		ResourceID:   req.ResourceID,
		Status:       req.Status,
		SessionID:    req.SessionID,
		Since:        req.From,
		Until:        req.To,
	}, p.PerPage, p.Offset)
	if err != nil {
		c.Fail(ctx, err)
	}
	return c.OKWithPagination(ctx, ListAuditLogsRes(entries), count), nil
}

func (c AuditController) GetAuditLog(
	ctx context.Context,
	req *GetAuditLogReq,
) (res http.Response[GetAuditLogRes], err error) {
	entry, err := c.auditService.GetByID(ctx, req.ID)
	if err != nil {
		c.Fail(ctx, err)
	}
	return c.OK(GetAuditLogRes(entry)), nil
}
