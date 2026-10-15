package users

import (
	"context"
	"time"

	"modular-monolith/internal/auth/app/services"
	"modular-monolith/internal/auth/domain/access"
	http "modular-monolith/lib/http"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

type ListUserRolesReq struct {
	g.Meta `path:"/:id/roles" method:"get" tags:"Users" x-group:"Admin/Users" summary:"List the roles assigned to a user" security:"bearer"`
	http.RequestBaseParams
	ID uuid.UUID `                                                                                                                                json:"id" v:"required"`
}
type ListUserRolesRes []access.UserRole

func (c UsersController) ListUserRoles(
	ctx context.Context,
	req *ListUserRolesReq,
) (res http.Response[ListUserRolesRes], err error) {
	p := c.Pagination(ctx)
	roles, count, err := c.usersService.ListRoles(ctx, req.ID, p.PerPage, p.Offset)
	if err != nil {
		c.Fail(ctx, err)
	}
	return c.OKWithPagination(ctx, ListUserRolesRes(roles), count), nil
}

type AssignUserRoleReq struct {
	g.Meta    `path:"/:id/roles" method:"post" tags:"Users" x-group:"Admin/Users" summary:"Assign a role to a user" security:"bearer"`
	ID        uuid.UUID     `                                                                                                                       json:"id"        v:"required"`
	RoleID    access.RoleID `                                                                                                                       json:"roleId"    v:"required" dc:"the role to assign"`
	IsPrimary *bool         `                                                                                                                       json:"isPrimary"              dc:"make this the user's primary role, demoting the previous one"`
	ExpiresAt *time.Time    `                                                                                                                       json:"expiresAt"              dc:"when the assignment lapses, nil grants it permanently"`
}
type AssignUserRoleRes struct {
	g.Meta  `mime:"application/json"`
	ID      uuid.UUID     `                        json:"id"      dc:"the user"`
	RoleID  access.RoleID `                        json:"roleId"  dc:"the assigned role"`
	Primary bool          `                        json:"primary" dc:"true when the role became primary"`
}

func (c UsersController) AssignUserRole(
	ctx context.Context,
	req *AssignUserRoleReq,
) (res http.Response[AssignUserRoleRes], err error) {
	actor, ok := c.UserID(ctx)
	if !ok {
		c.Unauthorized(ctx, "assign user role")
		return res, nil
	}
	primary, err := c.usersService.AssignRole(ctx, services.AssignRoleInput{
		UserID:    req.ID,
		RoleID:    req.RoleID,
		IsPrimary: req.IsPrimary,
		ExpiresAt: req.ExpiresAt,
	}, actor)
	if err != nil {
		c.Fail(ctx, err)
	}
	return c.Created(AssignUserRoleRes{
		ID:      req.ID,
		RoleID:  req.RoleID,
		Primary: primary,
	}), nil
}

type RevokeUserRoleReq struct {
	g.Meta `path:"/:id/roles/:roleId" method:"delete" tags:"Users" x-group:"Admin/Users" summary:"Revoke a role from a user" security:"bearer"`
	ID     uuid.UUID     `                                                                                                                                   json:"id"     v:"required"`
	RoleID access.RoleID `                                                                                                                                   json:"roleId" v:"required"      dc:"the role to revoke"`
	Reason string        `                                                                                                                                   json:"reason" v:"length:10,500" dc:"why, written to the audit log"`
}
type RevokeUserRoleRes struct {
	g.Meta  `mime:"application/json"`
	ID      uuid.UUID     `                        json:"id"      dc:"the user"`
	RoleID  access.RoleID `                        json:"roleId"  dc:"the revoked role"`
	Revoked bool          `                        json:"revoked" dc:"always true on success"`
}

func (c UsersController) RevokeUserRole(
	ctx context.Context,
	req *RevokeUserRoleReq,
) (res http.Response[RevokeUserRoleRes], err error) {
	actor, ok := c.UserID(ctx)
	if !ok {
		c.Unauthorized(ctx, "revoke user role")
		return res, nil
	}
	if err := c.usersService.RevokeRole(ctx, req.ID, req.RoleID, actor); err != nil {
		c.Fail(ctx, err)
	}
	return c.OK(RevokeUserRoleRes{ID: req.ID, RoleID: req.RoleID, Revoked: true}), nil
}
