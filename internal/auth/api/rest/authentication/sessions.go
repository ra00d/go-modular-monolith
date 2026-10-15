package authentication

import (
	"context"

	libauth "modular-monolith/lib/auth"
	"modular-monolith/lib/http"

	"github.com/gogf/gf/v2/frame/g"
)

type ListSessionsReq struct {
	g.Meta `path:"/sessions" method:"get" tags:"Auth" summary:"List the active sessions of the signed-in user" security:"bearer"`
	http.RequestBaseParams
}
type ListSessionsRes []libauth.SessionData

func (c *AuthController) ListSessions(
	ctx context.Context,
	_ *ListSessionsReq,
) (res http.Response[ListSessionsRes], err error) {
	userID, ok := c.UserID(ctx)
	if !ok {
		c.Unauthorized(ctx, "list sessions")
		return res, nil
	}
	// Sessions live in redis, not postgres, so there is no count query to run:
	// the whole set is small and bounded by the token TTL.
	sessions, err := c.jwt.ListSessions(ctx, userID.String())
	if err != nil {
		c.Fail(ctx, err)
	}
	total := len(sessions)
	start := c.Pagination(ctx).Offset
	if start > total {
		start = total
	}
	end := min(start+c.Pagination(ctx).PerPage, total)

	return c.OKWithPagination(ctx, ListSessionsRes(sessions[start:end]), total), nil
}

type RevokeSessionReq struct {
	g.Meta `                                      path:"/sessions/:id" method:"delete" tags:"Auth" summary:"Revoke one session by id" security:"bearer"`
	ID     string `                                json:"id" v:"required" dc:"the session id (token id)"`
}
type RevokeSessionRes struct {
	g.Meta  `mime:"application/json"`
	ID      string `json:"id"      dc:"the revoked session id"`
	Revoked bool   `json:"revoked" dc:"always true on success"`
}

func (c *AuthController) RevokeSession(
	ctx context.Context,
	req *RevokeSessionReq,
) (res http.Response[RevokeSessionRes], err error) {
	userID, ok := c.UserID(ctx)
	if !ok {
		c.Unauthorized(ctx, "revoke session")
		return res, nil
	}
	if err := c.jwt.RevokeSession(ctx, userID.String(), req.ID); err != nil {
		c.Fail(ctx, err)
	}
	return c.OK(RevokeSessionRes{ID: req.ID, Revoked: true}), nil
}

type RevokeAllSessionsReq struct {
	g.Meta `path:"/sessions" method:"delete" tags:"Auth" summary:"Revoke every session of the signed-in user" security:"bearer"`
}
type RevokeAllSessionsRes struct {
	g.Meta  `mime:"application/json"`
	Message string `json:"message" dc:"all sessions revoked"`
}

func (c *AuthController) RevokeAllSessions(
	ctx context.Context,
	_ *RevokeAllSessionsReq,
) (res http.Response[RevokeAllSessionsRes], err error) {
	userID, ok := c.UserID(ctx)
	if !ok {
		c.Unauthorized(ctx, "revoke sessions")
		return res, nil
	}
	if err := c.jwt.RevokeAllUserTokens(ctx, userID.String()); err != nil {
		c.Fail(ctx, err)
	}
	return c.OK(RevokeAllSessionsRes{Message: "all sessions revoked"}), nil
}

type LogoutAllReq struct {
	g.Meta `path:"/logout-all" method:"post" tags:"Auth" summary:"Sign out of every session" security:"bearer"`
}
type LogoutAllRes struct {
	g.Meta  `mime:"application/json"`
	Message string `json:"message" dc:"signed out of all sessions"`
}

func (c *AuthController) LogoutAll(
	ctx context.Context,
	_ *LogoutAllReq,
) (res http.Response[LogoutAllRes], err error) {
	userID, ok := c.UserID(ctx)
	if !ok {
		c.Unauthorized(ctx, "logout all")
		return res, nil
	}
	if err := c.jwt.RevokeAllUserTokens(ctx, userID.String()); err != nil {
		c.Fail(ctx, err)
	}
	return c.OK(LogoutAllRes{Message: "signed out of all sessions"}), nil
}
