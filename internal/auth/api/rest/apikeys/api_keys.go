// Package apikeys manages the service accounts and machine credentials of the
// platform (the api_key table). Keys are shown in plaintext exactly once, at
// creation and at rotation; only a hash is stored.
package apikeys

import (
	"context"
	"time"

	"modular-monolith/internal/auth/app/services"
	"modular-monolith/internal/auth/di"
	"modular-monolith/internal/auth/domain/access"
	"modular-monolith/lib/errors"
	"modular-monolith/lib/http"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/google/uuid"
)

// APIKey is the safe projection of a key: no hash, no secret.
type APIKey = access.APIKey

// APIKeyController serves the admin-only key management endpoints.
type APIKeyController struct {
	*http.BaseController
	apiKeysService *services.APIKeyService
}

func RegisterControllers(deps di.AuthDepsIn) {
	ctlr := &APIKeyController{
		BaseController: deps.BaseController,
		apiKeysService: deps.APIKeysService,
	}

	deps.APIGrp.Middleware(deps.Guard.Authenticate).Group("/admin", func(group *ghttp.RouterGroup) {
		group.Group("/api-keys", func(kg *ghttp.RouterGroup) {
			kg.
				Bind(ctlr.ListAPIKeys).
				Bind(ctlr.CreateAPIKey).
				Bind(ctlr.GetAPIKey).
				Bind(ctlr.UpdateAPIKey).
				Bind(ctlr.RevokeAPIKey).
				Bind(ctlr.RotateAPIKey)
		})
	})
}

type ListAPIKeysReq struct {
	g.Meta `                                                                                        path:"/" method:"get" tags:"APIKeys" x-group:"Admin/APIKeys" summary:"List the issued api keys" security:"bearer"`
	http.RequestBaseParams
	UserID   *uuid.UUID     `                                                                                      json:"userId"    dc:"restrict to the keys of one user"`
	RoleID   *access.RoleID `                                                                                 json:"roleId"    dc:"restrict to the keys acting under one role"`
	IsActive *bool          `                                                                                      json:"isActive"  dc:"false returns revoked or expired keys too"`
}

// ListAPIKeysRes is the page of matching keys. The total is carried in the
// response meta rather than duplicated in the payload.
type ListAPIKeysRes []APIKey

type CreateAPIKeyReq struct {
	g.Meta      `                                                                       path:"/" method:"post" tags:"APIKeys" x-group:"Admin/APIKeys" summary:"Issue an api key" security:"bearer"`
	Name        string         `                                                             json:"name"       v:"required|length:2,100" dc:"human label, shows up in the audit log"`
	UserID      *uuid.UUID     `                                                             json:"userId"     dc:"owning user, omit for a service account"`
	RoleID      *access.RoleID `                                                             json:"roleId"     dc:"role the key acts under"`
	Permissions []string       `                                                             json:"permissions" dc:"permission slugs, they narrow the role grants"`
	RateLimit   *int           `                                                             json:"rateLimit"  dc:"requests per minute"`
	ExpiresAt   *time.Time     `                                                             json:"expiresAt"  dc:"when the key stops working, prefer a short life"`
}
type CreateAPIKeyRes struct {
	g.Meta   `mime:"application/json"`
	APIKey   `dc:"the created key, without the secret"`
	Key      string    `json:"key"      dc:"the plaintext key, returned only once and never recoverable"`
	IssuedAt time.Time `json:"issuedAt" dc:"when the key was issued"`
}

type GetAPIKeyReq struct {
	g.Meta `                                                path:"/:id" method:"get" tags:"APIKeys" x-group:"Admin/APIKeys" summary:"Get one api key" security:"bearer"`
	ID     int64 ` v:"required" json:"id" dc:"the key id"`
}
type GetAPIKeyRes APIKey

type UpdateAPIKeyReq struct {
	g.Meta      `                                                                          path:"/:id" method:"patch" tags:"APIKeys" x-group:"Admin/APIKeys" summary:"Update an api key, the secret itself cannot be changed" security:"bearer"`
	ID          int64      `                                                            v:"required" json:"id"`
	Name        *string    `                                                            json:"name"        v:"length:2,100" dc:"human label"`
	Permissions *[]string  `                                                            json:"permissions" dc:"replaces the permission slugs"`
	RateLimit   *int       `                                                            json:"rateLimit"   dc:"requests per minute"`
	ExpiresAt   *time.Time `                                                            json:"expiresAt"   dc:"when the key stops working"`
	IsActive    *bool      `                                                            json:"isActive"    dc:"suspend or resume the key"`
	Reason      string     `                                                            json:"reason"      v:"length:10,500" dc:"why, written to the audit log"`
}
type UpdateAPIKeyRes struct {
	g.Meta `mime:"application/json"`
	ID     int64 `json:"id" dc:"the updated key id"`
}

type RevokeAPIKeyReq struct {
	g.Meta `                                                                             path:"/:id" method:"delete" tags:"APIKeys" x-group:"Admin/APIKeys" summary:"Revoke an api key" security:"bearer"`
	ID     int64  `                                                                       json:"id" v:"required" dc:"the key id"`
	Reason string `                                                                       json:"reason" v:"required|length:10,500" dc:"why, written to the audit log"`
}
type RevokeAPIKeyRes struct {
	g.Meta  `mime:"application/json"`
	ID      int64 `json:"id"      dc:"the revoked key id"`
	Revoked bool  `json:"revoked" dc:"always true on success"`
}

type RotateAPIKeyReq struct {
	g.Meta `                                                                    path:"/:id/rotate" method:"post" tags:"APIKeys" x-group:"Admin/APIKeys" summary:"Replace the secret of a key, the old one stops working at once" security:"bearer"`
	ID     int64  `                                                                  json:"id" v:"required" dc:"the key id"`
	Reason string `                                                                  json:"reason" v:"length:10,500" dc:"why the secret is replaced, written to the audit log"`
}
type RotateAPIKeyRes struct {
	g.Meta    `mime:"application/json"`
	ID        int64     `json:"id"       dc:"the rotated key id"`
	Key       string    `json:"key"      dc:"the new plaintext key, returned only once"`
	Prefix    string    `json:"prefix"   dc:"prefix of the new key"`
	RotatedAt time.Time `json:"rotatedAt" dc:"when the secret was replaced"`
}

func (c APIKeyController) ListAPIKeys(
	ctx context.Context,
	req *ListAPIKeysReq,
) (res http.Response[ListAPIKeysRes], err error) {
	p := c.Pagination(ctx)
	keys, count, err := c.apiKeysService.List(ctx, access.APIKeyFilter{
		UserID:   req.UserID,
		RoleID:   req.RoleID,
		IsActive: req.IsActive,
	}, p.PerPage, p.Offset)
	if err != nil {
		c.Fail(ctx, err)
	}
	return c.OKWithPagination(ctx, ListAPIKeysRes(keys), count), nil
}

func (c APIKeyController) CreateAPIKey(
	ctx context.Context,
	req *CreateAPIKeyReq,
) (res http.Response[CreateAPIKeyRes], err error) {
	return res, errors.NotImplemented("apikeys.create", "not implemented")
}

func (c APIKeyController) GetAPIKey(
	ctx context.Context,
	req *GetAPIKeyReq,
) (res http.Response[GetAPIKeyRes], err error) {
	key, err := c.apiKeysService.GetByID(ctx, req.ID)
	if err != nil {
		c.Fail(ctx, err)
	}
	return c.OK(GetAPIKeyRes(key)), nil
}

func (c APIKeyController) UpdateAPIKey(
	ctx context.Context,
	req *UpdateAPIKeyReq,
) (res http.Response[UpdateAPIKeyRes], err error) {
	return res, errors.NotImplemented("apikeys.update", "not implemented")
}

func (c APIKeyController) RevokeAPIKey(
	ctx context.Context,
	req *RevokeAPIKeyReq,
) (res http.Response[RevokeAPIKeyRes], err error) {
	actor, ok := c.UserID(ctx)
	if !ok {
		c.Unauthorized(ctx, "revoke api key")
		return res, nil
	}
	if err := c.apiKeysService.Revoke(ctx, req.ID, actor); err != nil {
		c.Fail(ctx, err)
	}
	return c.OK(RevokeAPIKeyRes{ID: req.ID, Revoked: true}), nil
}

func (c APIKeyController) RotateAPIKey(
	ctx context.Context,
	req *RotateAPIKeyReq,
) (res http.Response[RotateAPIKeyRes], err error) {
	return res, errors.NotImplemented("apikeys.rotate", "not implemented")
}
