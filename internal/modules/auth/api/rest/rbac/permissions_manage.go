package rbac

import (
	"context"

	"modular-monolith/internal/modules/auth/app/services"
	"modular-monolith/internal/modules/auth/domain/access"
	"modular-monolith/lib/http"

	"github.com/gogf/gf/v2/frame/g"
)

type GetPermissionReq struct {
	g.Meta `                                                             path:"/:id" method:"get" tags:"RBAC" x-group:"Admin/Permissions" summary:"Get a permission" security:"bearer"`
	ID     access.PermissionID ` v:"required" json:"id" dc:"the permission id"`
}
type GetPermissionRes access.Permission

func (c PermissionsController) GetPermission(
	ctx context.Context,
	req *GetPermissionReq,
) (res http.Response[GetPermissionRes], err error) {
	permission, err := c.permissionsService.GetByID(ctx, req.ID)
	if err != nil {
		c.Fail(ctx, err)
	}
	return c.OK(GetPermissionRes(permission)), nil
}

type UpdatePermissionReq struct {
	g.Meta      `                                                                                              path:"/:id" method:"patch" tags:"RBAC" x-group:"Admin/Permissions" summary:"Update a permission" security:"bearer"`
	ID          access.PermissionID `                                                                           v:"required" json:"id"`
	Name        *string             `                                                                          json:"name"        v:"length:2,100" dc:"display name"`
	Description *string             `                                                                          json:"description" dc:"what the permission allows"`
	Category    *string             `                                                                          json:"category"    dc:"grouping"`
	Conditions  map[string]any      `                                                                          json:"conditions"  dc:"replaces the ABAC conditions of the permission"`
}
type UpdatePermissionRes struct {
	g.Meta `mime:"application/json"`
	ID     access.PermissionID `json:"id" dc:"the updated permission"`
}

func (c PermissionsController) UpdatePermission(
	ctx context.Context,
	req *UpdatePermissionReq,
) (res http.Response[UpdatePermissionRes], err error) {
	actor, ok := c.UserID(ctx)
	if !ok {
		c.Unauthorized(ctx, "update permission")
		return res, nil
	}
	err = c.permissionsService.Update(ctx, req.ID, services.UpdatePermissionInput{
		Name:        req.Name,
		Description: req.Description,
		Category:    req.Category,
		Conditions:  req.Conditions,
	}, actor)
	if err != nil {
		c.Fail(ctx, err)
	}
	return c.OK(UpdatePermissionRes{ID: req.ID}), nil
}

type DeletePermissionReq struct {
	g.Meta `                                                                                 path:"/:id" method:"delete" tags:"RBAC" x-group:"Admin/Permissions" summary:"Delete a permission, system permissions are rejected" security:"bearer"`
	ID     access.PermissionID ` v:"required" json:"id" dc:"the permission id"`
	Reason string              ` v:"length:10,500" json:"reason" dc:"why, written to the audit log"`
}
type DeletePermissionRes struct {
	g.Meta  `mime:"application/json"`
	ID      access.PermissionID `json:"id"      dc:"the deleted permission"`
	Deleted bool                `json:"deleted" dc:"always true on success"`
}

func (c PermissionsController) DeletePermission(
	ctx context.Context,
	req *DeletePermissionReq,
) (res http.Response[DeletePermissionRes], err error) {
	actor, ok := c.UserID(ctx)
	if !ok {
		c.Unauthorized(ctx, "delete permission")
		return res, nil
	}
	if err := c.permissionsService.Delete(ctx, req.ID, actor); err != nil {
		c.Fail(ctx, err)
	}
	return c.OK(DeletePermissionRes{ID: req.ID, Deleted: true}), nil
}
