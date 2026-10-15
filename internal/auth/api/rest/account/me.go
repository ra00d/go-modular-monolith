package account

import (
	"context"
	"errors"

	"modular-monolith/internal/auth/domain/users"
	"modular-monolith/lib/auth"
	"modular-monolith/lib/http"

	"github.com/gogf/gf/v2/frame/g"
)

// MeRequest fetches the current user's profile.
type MeRequest struct {
	g.Meta `path:"/me" method:"get" tags:"Account" summary:"Get the current user" security:"bearer"`
}

// MeResponse is the current user's profile.
type MeResponse struct {
	users.User
}

// Me returns the authenticated user's profile.
func (c *AccountController) Me(
	ctx context.Context,
	_ *MeRequest,
) (res http.Response[MeResponse], err error) {
	id, ok := auth.UserID(ctx)
	if !ok {
		return res, errors.New("unauthenticated")
	}
	user, err := c.userService.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, users.ErrUserNotFound) {
			return res, errors.New("user not found")
		}
		return res, err
	}
	return c.OK(MeResponse{User: user}), nil
}
