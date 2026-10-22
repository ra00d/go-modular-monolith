package authentication

import (
	"context"

	"modular-monolith/lib/errors"
	"modular-monolith/lib/http"

	"github.com/gogf/gf/v2/frame/g"
)

// LogoutRequest revokes the current session token.
type LogoutRequest struct {
	g.Meta `path:"/logout" method:"post" tags:"Auth" summary:"Revoke the current session" security:"BearerAuth"`
}

// LogoutResponse confirms logout.
type LogoutResponse struct {
	g.Meta  `       mime:"application/json"`
	Message string `                        json:"message"`
}

// Logout revokes the presented token.
func (c *AuthController) Logout(
	ctx context.Context,
	_ *LogoutRequest,
) (http.Response[LogoutResponse], error) {
	if err := c.Jwt.RevokeCurrentToken(ctx); err != nil {
		c.Internal(ctx, errors.Internal("logout failed", err).Err)
	}
	return c.OK(LogoutResponse{Message: "logged out"}), nil
}
