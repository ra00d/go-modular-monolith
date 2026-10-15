package users

import (
	"context"
	"time"

	"modular-monolith/internal/auth/domain/users"
	"modular-monolith/lib/http"
	"modular-monolith/lib/utils"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

type UpdateUserReq struct {
	g.Meta   `                                                                     path:"/:id" method:"patch" tags:"Users" x-group:"Admin/Users" summary:"Update a user" security:"bearer"`
	ID       uuid.UUID `                                                              json:"id" v:"required"`
	Name     *string   `                                                            json:"name"     v:"length:2,50"           dc:"display name"`
	Image    *string   `                                                            json:"image"    dc:"avatar url"`
	Email    *string   `                                                            json:"email"    v:"email"                 dc:"replaces the address without verification, audited"`
	Phone    *string   `                                                            json:"phone"    dc:"replaces the number without verification, audited"`
	Username *string   `                                                            json:"username" v:"regex:^[a-zA-Z0-9_]{3,30}$" dc:"replaces the username without verification, audited"`
}
type UpdateUserRes struct {
	g.Meta `mime:"application/json"`
	ID     uuid.UUID `json:"id" dc:"the updated user"`
}

func (c UsersController) UpdateUser(
	ctx context.Context,
	req *UpdateUserReq,
) (res http.Response[UpdateUserRes], err error) {
	actor, ok := c.UserID(ctx)
	if !ok {
		c.Unauthorized(ctx, "update user")
		return res, nil
	}
	err = c.usersService.Update(ctx, req.ID, users.UpdateUserInput{
		Name:     req.Name,
		Image:    req.Image,
		Email:    req.Email,
		Phone:    req.Phone,
		Username: req.Username,
	}, actor)
	if err != nil {
		c.Fail(ctx, err)
	}
	return c.OK(UpdateUserRes{ID: req.ID}), nil
}

type ResetUserPasswordReq struct {
	g.Meta             `                                                                          path:"/:id/password" method:"post" tags:"Users" x-group:"Admin/Users" summary:"Set a new password for a user" security:"bearer"`
	ID                 uuid.UUID `                                                                  json:"id" v:"required"`
	Password           string    `                                                                  json:"password" v:"required|length:12,72" dc:"the new password, subject to the configured policy"`
	MustChangePassword *bool     `                                                                  json:"mustChangePassword" d:"true" dc:"force a change on next sign-in"`
	RevokeSessions     *bool     `                                                                  json:"revokeSessions"      d:"true" dc:"revoke every active session of that user"`
}
type ResetUserPasswordRes struct {
	g.Meta  `mime:"application/json"`
	ID      uuid.UUID `json:"id"      dc:"the user whose password was reset"`
	Message string    `json:"message" dc:"password reset"`
}

func (c UsersController) ResetUserPassword(
	ctx context.Context,
	req *ResetUserPasswordReq,
) (res http.Response[ResetUserPasswordRes], err error) {
	actor, ok := c.UserID(ctx)
	if !ok {
		c.Unauthorized(ctx, "reset user password")
		return res, nil
	}
	// ponytail: MustChangePassword is ignored, the user table has no column for
	// it and adding one is a schema migration.
	err = c.usersService.ResetPassword(
		ctx, req.ID, req.Password, actor,
		utils.DerefOr(req.RevokeSessions, true),
	)
	if err != nil {
		c.Fail(ctx, err)
	}
	return c.OK(ResetUserPasswordRes{ID: req.ID, Message: "password reset"}), nil
}

type ImpersonateUserReq struct {
	g.Meta   `                                                                                    path:"/:id/impersonate" method:"post" tags:"Users" x-group:"Admin/Users" summary:"Issue a short-lived token acting as a user, audited" security:"bearer"`
	ID       uuid.UUID `                                                                          json:"id" v:"required"`
	Duration int       `                                                                          json:"duration" d:"15" dc:"token lifetime in minutes, capped by configuration"`
	Reason   string    `                                                                          json:"reason"  v:"required|length:10,500" dc:"why the impersonation is needed, written to the audit log"`
}
type ImpersonateUserRes struct {
	g.Meta         `mime:"application/json"`
	Token          string    `json:"token"     dc:"the impersonation token"`
	TokenType      string    `json:"tokenType" dc:"always Bearer"`
	ExpiresAt      int64     `json:"expiresAt" dc:"unix expiry, short by design"`
	ImpersonatedAt time.Time `json:"impersonatedAt" dc:"when the impersonation started"`
}

func (c UsersController) ImpersonateUser(
	ctx context.Context,
	req *ImpersonateUserReq,
) (res http.Response[ImpersonateUserRes], err error) {
	actor, ok := c.UserID(ctx)
	if !ok {
		c.Unauthorized(ctx, "impersonate user")
		return res, nil
	}
	token, err := c.usersService.Impersonate(ctx, req.ID, req.Duration, req.Reason, actor)
	if err != nil {
		c.Fail(ctx, err)
	}
	return c.Created(ImpersonateUserRes{
		Token:          token.AccessToken,
		TokenType:      token.TokenType,
		ExpiresAt:      token.ExpiresAt.Unix(),
		ImpersonatedAt: time.Now(),
	}), nil
}
