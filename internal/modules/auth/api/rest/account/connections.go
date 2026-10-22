package account

import (
	"context"

	"modular-monolith/internal/modules/auth/domain/users"
	"modular-monolith/lib/http"

	"github.com/gogf/gf/v2/frame/g"
)

// Connection is an external identity linked to the account (the account
// table: provider_id + provider_account_id).
type Connection = users.Connection

type ListConnectionsReq struct {
	g.Meta `path:"/connections" method:"get" tags:"Account" summary:"List the social identities linked to the account" security:"bearer"`
	http.RequestBaseParams
}
type ListConnectionsRes []Connection

func (c AccountController) ListConnections(
	ctx context.Context,
	_ *ListConnectionsReq,
) (res http.Response[ListConnectionsRes], err error) {
	userID, ok := c.UserID(ctx)
	if !ok {
		c.Unauthorized(ctx, "list connections")
		return res, nil
	}
	// A user has a handful of linked identities at most, so the page is sliced
	// here rather than pushed into a count query.
	connections, err := c.userService.ListConnections(ctx, userID)
	if err != nil {
		c.Fail(ctx, err)
	}
	total := len(connections)
	start := c.Pagination(ctx).Offset
	if start > total {
		start = total
	}
	end := min(start+c.Pagination(ctx).PerPage, total)

	return c.OKWithPagination(ctx, ListConnectionsRes(connections[start:end]), total), nil
}

type UnlinkConnectionReq struct {
	g.Meta   `                                                                                      path:"/connections/:provider" method:"delete" tags:"Account" summary:"Unlink a social identity" security:"bearer"`
	Provider string `                                                                                json:"provider" v:"required|in:google,github,apple,microsoft" dc:"the provider to unlink"`
	Password string `                                                                                json:"password" v:"required" dc:"the current password, proves possession"`
	Revoke   *bool  `                                                                                json:"revoke"   d:"true" dc:"also revoke the tokens held at the provider"`
}
type UnlinkConnectionRes struct {
	g.Meta   `mime:"application/json"`
	Provider string `json:"provider"  dc:"the unlinked provider"`
	Unlinked bool   `json:"unlinked"  dc:"always true on success"`
}

func (c AccountController) UnlinkConnection(
	ctx context.Context,
	req *UnlinkConnectionReq,
) (res http.Response[UnlinkConnectionRes], err error) {
	userID, ok := c.UserID(ctx)
	if !ok {
		c.Unauthorized(ctx, "unlink connection")
		return res, nil
	}
	// ponytail: req.Revoke is ignored, revoking the tokens at the provider needs
	// an OAuth client that this module does not have yet.
	if err := c.userService.UnlinkConnection(ctx, userID, req.Provider, req.Password); err != nil {
		c.Fail(ctx, err)
	}
	return c.OK(UnlinkConnectionRes{Provider: req.Provider, Unlinked: true}), nil
}
