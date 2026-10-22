package users

import (
	"context"

	"modular-monolith/internal/modules/auth/domain/users"
	"modular-monolith/lib/http"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

type GetUserReq struct {
	g.Meta `path:"/:id" method:"get" tags:"Users" x-group:"Admin/Users" security:"bearer"`
	ID     uuid.UUID `                                                                              json:"id"`
}
type GetUserRes users.User

func (c UsersController) GetUser(
	ctx context.Context,
	req *GetUserReq,
) (res http.Response[users.User], err error) {
	data, err := c.usersService.GetByID(ctx, req.ID)
	if err != nil {
		c.Fail(ctx, err)
	}
	return c.OK(data), nil
}
