package users

import (
	"context"
	"time"

	"modular-monolith/lib/http"
	"modular-monolith/lib/utils"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

type BanUserReq struct {
	g.Meta         `                                                                          path:"/:id/ban" method:"post" tags:"Users" x-group:"Admin/Users" summary:"Ban a user until a given time or indefinitely" security:"bearer"`
	ID             uuid.UUID  `                                                                   json:"id" v:"required"`
	Reason         string     `                                                                   json:"reason" v:"required|length:10,500" dc:"why the ban is applied, written to the audit log"`
	Until          *time.Time `                                                                   json:"until" dc:"when the ban lifts, nil bans indefinitely"`
	RevokeSessions *bool      `                                                                  json:"revokeSessions" d:"true" dc:"revoke every active session of that user"`
}
type BanUserRes struct {
	g.Meta `mime:"application/json"`
	ID     uuid.UUID  `json:"id"      dc:"the banned user"`
	Status string     `json:"status"  dc:"the resulting user status"`
	Until  *time.Time `json:"until"   dc:"when the ban lifts"`
}

func (c UsersController) BanUser(
	ctx context.Context,
	req *BanUserReq,
) (res http.Response[BanUserRes], err error) {
	actor, ok := c.UserID(ctx)
	if !ok {
		c.Unauthorized(ctx, "ban user")
		return res, nil
	}
	status, err := c.usersService.Ban(
		ctx, req.ID, req.Until, req.Reason, actor,
		utils.DerefOr(req.RevokeSessions, true),
	)
	if err != nil {
		c.Fail(ctx, err)
	}
	return c.OK(BanUserRes{ID: req.ID, Status: status, Until: req.Until}), nil
}

type UnbanUserReq struct {
	g.Meta `                                                                                path:"/:id/unban" method:"post" tags:"Users" x-group:"Admin/Users" summary:"Lift a ban and clear the failed login counter" security:"bearer"`
	ID     uuid.UUID `                                                                      json:"id" v:"required"`
	Reason string    `                                                                      json:"reason" v:"length:10,500" dc:"why the ban is lifted, written to the audit log"`
}
type UnbanUserRes struct {
	g.Meta `mime:"application/json"`
	ID     uuid.UUID `json:"id"     dc:"the unbanned user"`
	Status string    `json:"status" dc:"the resulting user status"`
}

func (c UsersController) UnbanUser(
	ctx context.Context,
	req *UnbanUserReq,
) (res http.Response[UnbanUserRes], err error) {
	actor, ok := c.UserID(ctx)
	if !ok {
		c.Unauthorized(ctx, "unban user")
		return res, nil
	}
	status, err := c.usersService.Unban(ctx, req.ID, req.Reason, actor)
	if err != nil {
		c.Fail(ctx, err)
	}
	return c.OK(UnbanUserRes{ID: req.ID, Status: status}), nil
}

type SetUserStatusReq struct {
	g.Meta `                                                                path:"/:id/status" method:"patch" tags:"Users" x-group:"Admin/Users" summary:"Move a user to another account status" security:"bearer"`
	ID     uuid.UUID `                                                                          json:"id"     v:"required"`
	Status string    `                                                                          json:"status" v:"required|in:active,suspended,locked,pending_verification" dc:"the target status"`
	Reason string    `                                                                          json:"reason" v:"required|length:10,500" dc:"why, written to the audit log"`
}
type SetUserStatusRes struct {
	g.Meta `mime:"application/json"`
	ID     uuid.UUID `json:"id"     dc:"the user"`
	Status string    `json:"status" dc:"the resulting status"`
}

func (c UsersController) SetUserStatus(
	ctx context.Context,
	req *SetUserStatusReq,
) (res http.Response[SetUserStatusRes], err error) {
	actor, ok := c.UserID(ctx)
	if !ok {
		c.Unauthorized(ctx, "set user status")
		return res, nil
	}
	if err := c.usersService.SetStatus(ctx, req.ID, req.Status, actor); err != nil {
		c.Fail(ctx, err)
	}
	return c.OK(SetUserStatusRes{ID: req.ID, Status: req.Status}), nil
}
