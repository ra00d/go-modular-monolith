package rbac

import (
	"context"

	"modular-monolith/internal/auth/domain/access"
	"modular-monolith/lib/http"

	"github.com/gogf/gf/v2/frame/g"
)

type ListPermissionReq struct {
	g.Meta `path:"/" method:"get" tags:"RBAC" x-group:"Admin/Permissions" security:"bearer"`
	http.RequestBaseParams
}
type ListPermissionRes []access.Permission

func (c PermissionsController) ListPermissions(
	ctx context.Context,
	_ *ListPermissionReq,
) (res http.Response[[]access.Permission], err error) {
	p := c.Pagination(ctx)
	data, count, err := c.usersService.ListPermissions(ctx, p.PerPage, p.Offset)
	if err != nil {
		c.Fail(ctx, err)
	}
	return c.OKWithPagination(ctx, data, count), nil
}
