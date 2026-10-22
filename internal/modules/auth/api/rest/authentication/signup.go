package authentication

import (
	"context"
	"errors"

	"modular-monolith/internal/modules/auth/app/services"
	"modular-monolith/internal/modules/auth/domain/users"
	"modular-monolith/lib/http"

	"github.com/gogf/gf/v2/frame/g"
)

// SignUpRequest is the registration payload.
type SignUpRequest struct {
	g.Meta   `        path:"/sign-up" method:"post" tags:"Auth" summary:"Register a new account"`
	Name     string  `                                                                           json:"name"     v:"required|length:2,50"       dc:"Display name"`
	Email    *string `                                                                           json:"email"    v:"email"                      dc:"Email (optional if phone or username given)"`
	Phone    *string `                                                                           json:"phone"                                   dc:"Phone (optional if email or username given)"`
	Username *string `                                                                           json:"username" v:"regex:^[a-zA-Z0-9_]{3,30}$" dc:"Username (optional if email or phone given)"`
	Password string  `                                                                           json:"password" v:"required"                   dc:"Account password (subject to the configured policy)"`
}

// SignUpResponse is returned after successful registration.
type SignUpResponse struct {
	g.Meta  `       mime:"application/json"`
	Message string `                        json:"message"`
	UserID  string `                        json:"user_id"`
}

// SignUp registers a new local account.
func (c *PublicController) SignUp(
	ctx context.Context,
	req *SignUpRequest,
) (res http.Response[SignUpResponse], err error) {
	result, err := c.userService.Register(ctx, services.RegisterAccount{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		switch {
		case errors.Is(err, users.ErrUserAlreadyExists):
			return res, errors.New("a user with these credentials already exists")
		default:
			return res, err
		}
	}
	return c.OK(SignUpResponse{
		Message: "user registered successfully",
		UserID:  result.AccountID,
	}), nil
}
