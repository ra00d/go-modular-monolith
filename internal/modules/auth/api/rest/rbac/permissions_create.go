package rbac

import (
	"context"

	"modular-monolith/internal/modules/auth/app/services"
	"modular-monolith/internal/modules/auth/domain/access"
	"modular-monolith/lib/http"

	"github.com/gogf/gf/v2/frame/g"
)

type CreatePermissionReq struct {
	g.Meta      `                                                                                        path:"/" method:"post" tags:"RBAC" x-group:"Admin/Permissions" summary:"Create a permission" security:"bearer"`
	Name        string         `                                                                        json:"name"        v:"required|length:2,100" dc:"display name"`
	Slug        string         `                                                                        json:"slug"        v:"required|regex:^[a-z0-9_]+[.][a-z0-9_]+$" dc:"resource.action, e.g. invoice.export"`
	Description string         `                                                                        json:"description" dc:"what the permission allows"`
	Resource    string         `                                                                        json:"resource"    v:"required|length:2,50" dc:"the resource the action applies to"`
	Action      string         `                                                                        json:"action"      v:"required|in:create,read,update,delete,list,manage,assign,revoke,export" dc:"the action"`
	Category    string         `                                                                        json:"category"    v:"in:user-management,content,billing,system,admin" dc:"grouping used by the permission registry ui"`
	Conditions  map[string]any `                                                                        json:"conditions"  dc:"ABAC conditions applied whenever the permission is granted"`
}
type CreatePermissionRes struct {
	g.Meta `mime:"application/json"`
	ID     access.PermissionID `json:"id" dc:"the created permission id"`
	Slug   string              `json:"slug" dc:"the created permission slug"`
}

func (c PermissionsController) CreatePermission(
	ctx context.Context,
	req *CreatePermissionReq,
) (res http.Response[CreatePermissionRes], err error) {
	actor, ok := c.UserID(ctx)
	if !ok {
		c.Unauthorized(ctx, "create permission")
		return res, nil
	}
	id, slug, err := c.permissionsService.Create(ctx, services.CreatePermissionInput{
		Name:        req.Name,
		Slug:        req.Slug,
		Description: req.Description,
		Resource:    req.Resource,
		Action:      req.Action,
		Category:    req.Category,
		Conditions:  req.Conditions,
	}, actor)
	if err != nil {
		c.Fail(ctx, err)
	}
	return c.Created(CreatePermissionRes{ID: id, Slug: slug}), nil
}
