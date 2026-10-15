package rbac

import (
	"context"
	"time"

	"modular-monolith/internal/auth/app/services"
	"modular-monolith/internal/auth/domain/access"
	"modular-monolith/lib/http"

	"github.com/gogf/gf/v2/frame/g"
)

type GetRoleReq struct {
	g.Meta `                                                    path:"/:id" method:"get" tags:"RBAC" x-group:"Admin/Roles" summary:"Get a role" security:"bearer"`
	ID     access.RoleID ` v:"required" json:"id" dc:"the role id"`
}
type GetRoleRes access.Role

func (c RolesCntrl) GetRole(
	ctx context.Context,
	req *GetRoleReq,
) (res http.Response[GetRoleRes], err error) {
	role, err := c.rolesService.GetByID(ctx, req.ID)
	if err != nil {
		c.Fail(ctx, err)
	}
	return c.OK(GetRoleRes(role)), nil
}

type UpdateRoleReq struct {
	g.Meta       `                                                                            path:"/:id" method:"patch" tags:"RBAC" x-group:"Admin/Roles" summary:"Update a role" security:"bearer"`
	ID           access.RoleID  `                                                            v:"required" json:"id"`
	Name         *string        `                                                           json:"name"        v:"length:2,100" dc:"display name"`
	Description  *string        `                                                           json:"description" dc:"what the role is for"`
	ParentRoleID *access.RoleID `                                                           json:"parentRoleId" dc:"inheritance parent, nil detaches, must not create a cycle"`
	Metadata     map[string]any `                                                           json:"metadata"    dc:"replaces the stored attributes"`
}
type UpdateRoleRes struct {
	g.Meta `mime:"application/json"`
	ID     access.RoleID `json:"id" dc:"the updated role"`
}

func (c RolesCntrl) UpdateRole(
	ctx context.Context,
	req *UpdateRoleReq,
) (res http.Response[UpdateRoleRes], err error) {
	actor, ok := c.UserID(ctx)
	if !ok {
		c.Unauthorized(ctx, "update role")
		return res, nil
	}
	err = c.rolesService.Update(ctx, req.ID, services.UpdateRoleInput{
		Name:         req.Name,
		Description:  req.Description,
		ParentRoleID: req.ParentRoleID,
		Metadata:     req.Metadata,
	}, actor)
	if err != nil {
		c.Fail(ctx, err)
	}
	return c.OK(UpdateRoleRes{ID: req.ID}), nil
}

type DeleteRoleReq struct {
	g.Meta `                                                                         path:"/:id" method:"delete" tags:"RBAC" x-group:"Admin/Roles" summary:"Delete a role, system roles are rejected" security:"bearer"`
	ID     access.RoleID ` v:"required" json:"id" dc:"the role id"`
	Reason string        ` v:"length:10,500" json:"reason" dc:"why, written to the audit log"`
}
type DeleteRoleRes struct {
	g.Meta  `mime:"application/json"`
	ID      access.RoleID `json:"id"      dc:"the deleted role"`
	Deleted bool          `json:"deleted" dc:"always true on success"`
}

func (c RolesCntrl) DeleteRole(
	ctx context.Context,
	req *DeleteRoleReq,
) (res http.Response[DeleteRoleRes], err error) {
	actor, ok := c.UserID(ctx)
	if !ok {
		c.Unauthorized(ctx, "delete role")
		return res, nil
	}
	if err := c.rolesService.Delete(ctx, req.ID, actor); err != nil {
		c.Fail(ctx, err)
	}
	return c.OK(DeleteRoleRes{ID: req.ID, Deleted: true}), nil
}

type ListRolePermissionsReq struct {
	g.Meta `path:"/:id/permissions" method:"get" tags:"RBAC" x-group:"Admin/Roles" summary:"List the permissions granted to a role" security:"bearer"`
	http.RequestBaseParams
	ID access.RoleID `v:"required" json:"id" dc:"the role id"`
}
type ListRolePermissionsRes []access.RolePermission

func (c RolesCntrl) ListRolePermissions(
	ctx context.Context,
	req *ListRolePermissionsReq,
) (res http.Response[ListRolePermissionsRes], err error) {
	p := c.Pagination(ctx)
	grants, count, err := c.rolesService.ListPermissions(ctx, req.ID, p.PerPage, p.Offset)
	if err != nil {
		c.Fail(ctx, err)
	}
	return c.OKWithPagination(ctx, ListRolePermissionsRes(grants), count), nil
}

type AssignRolePermissionReq struct {
	g.Meta       `                                                                                path:"/:id/permissions" method:"post" tags:"RBAC" x-group:"Admin/Roles" summary:"Grant a permission to a role" security:"bearer"`
	ID           access.RoleID       `                                                                 v:"required" json:"id" dc:"the role id"`
	PermissionID access.PermissionID `                                                             json:"permissionId" v:"required" dc:"the permission to grant"`
	Conditions   map[string]any      `                                                                 json:"conditions" dc:"ABAC conditions narrowing the grant, e.g. {\"owner_only\":true}"`
	ExpiresAt    *time.Time          `                                                                 json:"expiresAt"  dc:"when the grant lapses, nil grants it permanently"`
}
type AssignRolePermissionRes struct {
	g.Meta       `mime:"application/json"`
	RoleID       access.RoleID       `json:"roleId"        dc:"the role"`
	PermissionID access.PermissionID `json:"permissionId"  dc:"the granted permission"`
	Granted      bool                `json:"granted"       dc:"always true on success"`
}

func (c RolesCntrl) AssignRolePermission(
	ctx context.Context,
	req *AssignRolePermissionReq,
) (res http.Response[AssignRolePermissionRes], err error) {
	actor, ok := c.UserID(ctx)
	if !ok {
		c.Unauthorized(ctx, "grant permission")
		return res, nil
	}
	err = c.rolesService.GrantPermission(ctx, services.GrantPermissionInput{
		RoleID:       req.ID,
		PermissionID: req.PermissionID,
		Conditions:   req.Conditions,
		ExpiresAt:    req.ExpiresAt,
	}, actor)
	if err != nil {
		c.Fail(ctx, err)
	}
	return c.OK(AssignRolePermissionRes{
		RoleID:       req.ID,
		PermissionID: req.PermissionID,
		Granted:      true,
	}), nil
}

type RevokeRolePermissionReq struct {
	g.Meta       `                                                                                path:"/:id/permissions/:permissionId" method:"delete" tags:"RBAC" x-group:"Admin/Roles" summary:"Revoke a permission from a role" security:"bearer"`
	ID           access.RoleID       `                                                            v:"required" json:"id" dc:"the role id"`
	PermissionID access.PermissionID `                                                            v:"required" json:"permissionId" dc:"the permission to revoke"`
	Reason       string              `                                                            v:"length:10,500" json:"reason" dc:"why, written to the audit log"`
}
type RevokeRolePermissionRes struct {
	g.Meta       `mime:"application/json"`
	RoleID       access.RoleID       `json:"roleId"        dc:"the role"`
	PermissionID access.PermissionID `json:"permissionId"  dc:"the revoked permission"`
	Revoked      bool                `json:"revoked"       dc:"always true on success"`
}

func (c RolesCntrl) RevokeRolePermission(
	ctx context.Context,
	req *RevokeRolePermissionReq,
) (res http.Response[RevokeRolePermissionRes], err error) {
	actor, ok := c.UserID(ctx)
	if !ok {
		c.Unauthorized(ctx, "revoke permission")
		return res, nil
	}
	if err := c.rolesService.RevokePermission(ctx, req.ID, req.PermissionID, actor); err != nil {
		c.Fail(ctx, err)
	}
	return c.OK(RevokeRolePermissionRes{
		RoleID:       req.ID,
		PermissionID: req.PermissionID,
		Revoked:      true,
	}), nil
}
