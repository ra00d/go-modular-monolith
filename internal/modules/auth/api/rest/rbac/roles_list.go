package rbac

import (
	"context"

	"modular-monolith/internal/modules/auth/domain/access"
	"modular-monolith/lib/http"

	"github.com/gogf/gf/v2/frame/g"
)

type ListRolesReq struct {
	g.Meta `path:"/" method:"get" tags:"RBAC" security:"bearer"`
	http.RequestBaseParams
}

func (c RolesCntrl) ListRoles(
	ctx context.Context,
	req *ListRolesReq,
) (res http.Response[[]access.Role], err error) {
	p := c.Pagination(ctx)
	data, count, err := c.rolesService.GetAll(ctx, p.PerPage, p.Offset)
	if err != nil {
		c.Fail(ctx, err)
	}
	return c.OKWithPagination(ctx, data, count), nil
}
