package authentication

import (
	"context"
	"errors"

	"modular-monolith/internal/modules/auth/app/services"
	"modular-monolith/internal/modules/auth/domain/users"
	"modular-monolith/lib/auth"
	"modular-monolith/lib/http"

	"github.com/gogf/gf/v2/frame/g"
)

// SignInRequest supports email, phone, or username login.
type SignInRequest struct {
	Identifier string `json:"identifier" v:"required" dc:"Email, phone, or username" example:"raad@email.com"`
	Password   string `json:"password"   v:"required" dc:"Account password"          example:"Password123#"`
	g.Meta     `                                                                                              path:"/sign-in" method:"post" tags:"Auth" sm:"Sign in with identifier and password"`
}

// SignInResponse carries the issued token and the user profile.
type SignInResponse struct {
	// common.CommonResponses
	Token      string `json:"token"      dc:"Bearer access token"`
	TokenType  string `json:"token_type" dc:"Always \"Bearer\""`
	ExpiresAt  int64  `json:"expires_at" dc:"Unix expiry of the access token"`
	users.User `                         dc:"Authenticated user"`
}

// SignIn authenticates the user and issues a session token.
func (c PublicController) SignIn(
	ctx context.Context,
	req *SignInRequest,
) (http.Response[SignInResponse], error) {
	r := g.RequestFromCtx(ctx)
	user, err := c.userService.SignIn(ctx, services.GetAuthAccount{
		Identifier: req.Identifier,
		Password:   req.Password,
		IP:         r.GetClientIp(),
	})
	if err != nil {
		switch {
		case errors.Is(err, users.ErrAccountLocked):
			c.Forbidden(ctx, "account is locked due to too many failed attempts")
		case errors.Is(err, users.ErrUserNotActive):
			c.Forbidden(ctx, "account is not active")
		case errors.Is(err, users.ErrInvalidPassword):
			c.Forbidden(ctx, "invalid password")
		default:
			c.Fail(ctx, err)
		}
	}

	token, err := c.Jwt.CreateAndStoreToken(
		ctx,
		auth.CreatTokenParams{
			UserID:      user.ID.String(),
			Identifier:  req.Identifier,
			Role:        primaryRole(user),
			Roles:       user.Roles,
			Permissions: user.Permissions,
		},
	)
	if err != nil {
		c.Fail(ctx, err)
	}
	c.SetCookie(ctx, "auth-token", token.AccessToken)
	data := SignInResponse{
		Token:     token.AccessToken,
		TokenType: token.TokenType,
		ExpiresAt: token.ExpiresAt.Unix(),
		User:      user,
	}
	return c.OK(data), nil
}

// primaryRole returns the user's first role, defaulting to "user".
func primaryRole(u users.User) string {
	if len(u.Roles) > 0 {
		return u.Roles[0]
	}
	return "user"
}
