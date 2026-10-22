package authentication

import (
	"context"

	"modular-monolith/lib/errors"
	"modular-monolith/lib/http"

	"github.com/gogf/gf/v2/frame/g"
)

// RefreshRequest exchanges the current token for a fresh one.
type RefreshRequest struct {
	g.Meta `path:"/refresh" method:"post" tags:"Auth" summary:"Refresh the session token" security:"BearerAuth"`
}

// RefreshResponse carries the new token.
type RefreshResponse struct {
	g.Meta    `       mime:"application/json"`
	Token     string `                        json:"token"`
	TokenType string `                        json:"token_type"`
	ExpiresAt int64  `                        json:"expires_at"`
}

// Refresh rotates the session token.
func (c *AuthController) Refresh(
	ctx context.Context,
	_ *RefreshRequest,
) (res http.Response[RefreshResponse], err error) {
	token, err := c.Jwt.RefreshToken(ctx)
	if err != nil {
		return res, errors.Internal("refresh failed", err).Err
	}
	return c.OK(RefreshResponse{
		Token:     token.AccessToken,
		TokenType: token.TokenType,
		ExpiresAt: token.ExpiresAt.Unix(),
	}), nil
}
