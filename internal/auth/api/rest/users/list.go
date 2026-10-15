package users

import (
	"context"

	"modular-monolith/internal/auth/domain/users"
	http "modular-monolith/lib/http"

	"github.com/gogf/gf/v2/frame/g"
)

type GetAllUsersReq struct {
	g.Meta `path:"/" method:"get" tags:"Users" x-sort:"1" x-group:"Admin/Users" security:"bearer"`
	http.RequestBaseParams
}

func (c UsersController) GetAllUsers(
	ctx context.Context,
	_ *GetAllUsersReq,
) (res http.Response[[]users.User], err error) {
	p := c.Pagination(ctx)
	data, count, err := c.usersService.GetAll(ctx, p.PerPage, p.Offset)
	if err != nil {
		c.Fail(ctx, err)
	}
	return c.OKWithPagination(ctx, data, count), nil
}
