package users

import (
	"context"

	"modular-monolith/internal/auth/app/services"
	"modular-monolith/lib/http"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

type CreateUserReq struct {
	g.Meta   `        path:"/" method:"post" tags:"Users" x-group:"Admin/Users" summary:"Create new User" security:"bearer"`
	Name     string  `                                                                                                      json:"name"     v:"required|length:2,50"       dc:"Display name"`
	Email    *string `                                                                                                      json:"email"    v:"email"                      dc:"Email (optional if phone or username given)"`
	Phone    *string `                                                                                                      json:"phone"                                   dc:"Phone (optional if email or username given)"`
	Username *string `                                                                                                      json:"username" v:"regex:^[a-zA-Z0-9_]{3,30}$" dc:"Username (optional if email or phone given)"`
	Password string  `                                                                                                      json:"password" v:"required"                   dc:"Account password (subject to the configured policy)"`
	Role     int     `                                                                                                      json:"role"                                    dc:"the role id of the user"                             d:"1"`
}

type CreateUserRes struct {
	ID uuid.UUID `json:"id"`
}

func (c UsersController) Create(
	ctx context.Context,
	req *CreateUserReq,
) (res http.Response[CreateUserRes], err error) {
	userID, err := c.usersService.Create(ctx, services.CreateUserInput{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Username: req.Username,
		Password: req.Password,
		Role:     req.Role,
	})
	if err != nil {
		c.Fail(ctx, err)
	}
	return c.OK(CreateUserRes{
		ID: userID,
	}), nil
}
