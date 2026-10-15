package rbac

import (
	"context"

	"modular-monolith/internal/auth/app/services"
	"modular-monolith/internal/auth/domain/access"
	"modular-monolith/lib/http"

	"github.com/gogf/gf/v2/frame/g"
)

type CreateRoleReq struct {
	g.Meta       `                                                                                     path:"/" method:"post" tags:"RBAC" x-group:"Admin/Roles" summary:"Create a role" security:"bearer"`
	Name         string         `                                                                     json:"name"        v:"required|length:2,100" dc:"display name"`
	Slug         string         `                                                                     json:"slug"        v:"required|regex:^[a-z0-9_]{2,50}$" dc:"stable machine name, unique across roles"`
	Description  string         `                                                                     json:"description" dc:"what the role is for"`
	ParentRoleID *access.RoleID `                                                                     json:"parentRoleId" dc:"role this one inherits permissions from"`
	Metadata     map[string]any `                                                                     json:"metadata"    dc:"JSON attributes the policy engine can read for this role"`
}
type CreateRoleRes struct {
	g.Meta `mime:"application/json"`
	ID     access.RoleID `json:"id" dc:"the created role id"`
	Slug   string        `json:"slug" dc:"the created role slug"`
}

func (c RolesCntrl) CreateRole(
	ctx context.Context,
	req *CreateRoleReq,
) (res http.Response[CreateRoleRes], err error) {
	actor, ok := c.UserID(ctx)
	if !ok {
		c.Unauthorized(ctx, "create role")
		return res, nil
	}
	id, slug, err := c.rolesService.Create(ctx, services.CreateRoleInput{
		Name:         req.Name,
		Slug:         req.Slug,
		Description:  req.Description,
		ParentRoleID: req.ParentRoleID,
		Metadata:     req.Metadata,
	}, actor)
	if err != nil {
		c.Fail(ctx, err)
	}
	return c.Created(CreateRoleRes{ID: id, Slug: slug}), nil
}
