package audit

import (
	"context"
	"time"

	"modular-monolith/internal/modules/auth/domain/access"
	"modular-monolith/lib/http"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

// SecurityEvent is one row of the security_event table: the high priority
// subset of the audit trail (logins, lockouts, MFA changes, denials).
type SecurityEvent = access.SecurityEvent

type ListSecurityEventsReq struct {
	g.Meta `                                                                                                                          path:"/" method:"get" tags:"Audit" x-group:"Admin/SecurityEvents" summary:"Search the security events" security:"bearer"`
	http.RequestBaseParams
	EventType  string     `                                                                                                                  json:"eventType" dc:"filter on one event type"`
	Severity   string     `                                                                                                                  json:"severity"  v:"in:info,warning,critical" dc:"filter on one severity"`
	UserID     *uuid.UUID `                                                                                                                 json:"userId"    dc:"events of one user"`
	Unresolved *bool      `                                                                                                                  json:"unresolved" dc:"true returns only events nobody has resolved yet"`
	From       *time.Time `                                                                                                                  json:"from"      dc:"only events at or after this time"`
	To         *time.Time `                                                                                                                  json:"to"        dc:"only events at or before this time"`
}

// ListSecurityEventsRes is the page of matching events, newest first. The total
// is carried in the response meta rather than duplicated in the payload.
type ListSecurityEventsRes []SecurityEvent

type GetSecurityEventReq struct {
	g.Meta `                                                    path:"/:id" method:"get" tags:"Audit" x-group:"Admin/SecurityEvents" summary:"Get one security event" security:"bearer"`
	ID     int64 ` v:"required" json:"id" dc:"the event id"`
}
type GetSecurityEventRes SecurityEvent

type ResolveSecurityEventReq struct {
	g.Meta `                                                                   path:"/:id/resolve" method:"post" tags:"Audit" x-group:"Admin/SecurityEvents" summary:"Mark a security event as handled" security:"bearer"`
	ID     int64  `                                                             json:"id" v:"required" dc:"the event id"`
	Note   string `                                                             json:"note" v:"required|length:10,1000" dc:"what was done about it, written to the audit log"`
}
type ResolveSecurityEventRes struct {
	g.Meta     `mime:"application/json"`
	ID         int64     `json:"id"         dc:"the resolved event"`
	ResolvedAt time.Time `json:"resolvedAt" dc:"when it was resolved"`
	ResolvedBy uuid.UUID `json:"resolvedBy" dc:"who resolved it"`
}

func (c AuditController) ListSecurityEvents(
	ctx context.Context,
	req *ListSecurityEventsReq,
) (res http.Response[ListSecurityEventsRes], err error) {
	p := c.Pagination(ctx)
	events, count, err := c.auditService.ListEvents(ctx, access.SecurityEventFilter{
		EventType:  req.EventType,
		Severity:   req.Severity,
		UserID:     req.UserID,
		Unresolved: req.Unresolved,
		Since:      req.From,
		Until:      req.To,
	}, p.PerPage, p.Offset)
	if err != nil {
		c.Fail(ctx, err)
	}
	return c.OKWithPagination(ctx, ListSecurityEventsRes(events), count), nil
}

func (c AuditController) GetSecurityEvent(
	ctx context.Context,
	req *GetSecurityEventReq,
) (res http.Response[GetSecurityEventRes], err error) {
	event, err := c.auditService.GetEvent(ctx, req.ID)
	if err != nil {
		c.Fail(ctx, err)
	}
	return c.OK(GetSecurityEventRes(event)), nil
}

func (c AuditController) ResolveSecurityEvent(
	ctx context.Context,
	req *ResolveSecurityEventReq,
) (res http.Response[ResolveSecurityEventRes], err error) {
	actor, ok := c.UserID(ctx)
	if !ok {
		c.Unauthorized(ctx, "resolve security event")
		return res, nil
	}
	if err := c.auditService.ResolveEvent(ctx, req.ID, actor); err != nil {
		c.Fail(ctx, err)
	}
	event, err := c.auditService.GetEvent(ctx, req.ID)
	if err != nil {
		c.Fail(ctx, err)
	}
	return c.OK(ResolveSecurityEventRes{
		ID:         event.ID,
		ResolvedAt: *event.ResolvedAt,
		ResolvedBy: *event.ResolvedBy,
	}), nil
}
