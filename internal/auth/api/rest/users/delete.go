package users

import (
	"context"

	"modular-monolith/lib/http"
	"modular-monolith/lib/utils"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

type DeleteUserReq struct {
	g.Meta         `                                                                                  path:"/:id" method:"delete" tags:"Users" x-group:"Admin/Users" summary:"Soft delete a user and revoke its sessions" security:"bearer"`
	ID             uuid.UUID `                                                                        json:"id" v:"required"`
	Reason         string    `                                                                        json:"reason" v:"required|length:10,500" dc:"why the account is removed, written to the audit log"`
	RevokeSessions *bool     `                                                                       json:"revokeSessions" d:"true" dc:"revoke every active session of that user"`
}
type DeleteUserRes struct {
	g.Meta  `mime:"application/json"`
	ID      uuid.UUID `json:"id"      dc:"the soft deleted user"`
	Deleted bool      `json:"deleted" dc:"always true on success"`
}

func (c UsersController) DeleteUser(
	ctx context.Context,
	req *DeleteUserReq,
) (res http.Response[DeleteUserRes], err error) {
	actor, ok := c.UserID(ctx)
	if !ok {
		c.Unauthorized(ctx, "delete user")
		return res, nil
	}
	if err := c.usersService.Delete(
		ctx, req.ID, actor, utils.DerefOr(req.RevokeSessions, true),
	); err != nil {
		c.Fail(ctx, err)
	}
	return c.OK(DeleteUserRes{ID: req.ID, Deleted: true}), nil
}
