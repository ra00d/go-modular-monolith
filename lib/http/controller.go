// Package http provides http  
package http

import (
	"context"

	"modular-monolith/lib/auth"
	"modular-monolith/lib/common"
	"modular-monolith/lib/errors"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/google/uuid"
)

type BaseController struct {
	Jwt   *auth.JWTManager
	Authz *auth.Authorizer
}

func NewBaseController(
	jwt *auth.JWTManager,
	authz *auth.Authorizer,
) *BaseController {
	return &BaseController{
		Jwt:   jwt,
		Authz: authz,
	}
}

func (c BaseController) GetRequest(ctx context.Context) *ghttp.Request {
	r := g.RequestFromCtx(ctx)

	if r == nil {
		panic("request is nil")
	}
	return r
}

func (c BaseController) SetCookie(ctx context.Context, key, value string) {
	req := c.GetRequest(ctx)
	req.Cookie.Set(key, value)
	// req.Cookie.SetHttpCookie(&http.Cookie{
	// 	Name:  key,
	// 	Value: value,
	// 	// Secure:   true,
	// 	SameSite: http.SameSiteNoneMode,
	// })
}

func (c BaseController) GetCookie(ctx context.Context, key string) string {
	req := c.GetRequest(ctx)
	return req.Cookie.Get(key).String()
}

func (c *BaseController) Pagination(
	ctx context.Context,
) common.Pagination {
	r := c.GetRequest(ctx)
	page := r.GetQuery("page", 1).Int()
	pageSize := r.GetQuery("perPage", 10).Int()
	cursor := r.GetQuery("cursor", "").String()

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 20
	}

	return common.Pagination{
		Page:    page,
		PerPage: pageSize,
		Offset:  (page - 1) * pageSize,
		Cursor:  cursor,
	}
}

func (c *BaseController) PaginationMeta(
	ctx context.Context,
	count int,
	cursor ...string,
) *common.PaginationMeta {
	pg := c.Pagination(ctx)
	pagination := common.GetPaginationMeta(count, pg.Page, pg.PerPage)
	if len(cursor) > 0 {
		pagination.SetCursor(cursor[len(cursor)-1])
	}
	return &pagination
}

func (c BaseController) Log() {
	//fmt.Println("loging from")
}

func (c BaseController) Fail(ctx context.Context, err error) {
	r := c.GetRequest(ctx)
	r.SetError(errors.From(err))
	r.Exit()
}

func (c BaseController) FailWithStatus(ctx context.Context, status int, err error) {
	r := c.GetRequest(ctx)
	r.Response.Status = status
	r.SetError(errors.From(err))
	r.Exit()
}

func (c BaseController) Forbidden(ctx context.Context, msg string) {
	r := c.GetRequest(ctx)
	// r.Response.Status = status
	r.SetError(errors.Forbidden("", msg))
	r.Exit()
}

func (c BaseController) Unauthorized(ctx context.Context, msg string) {
	r := c.GetRequest(ctx)
	r.SetError(errors.Unauthorized("", msg))
	r.Exit()
}

func (c BaseController) NotFound(ctx context.Context, msg string) {
	r := c.GetRequest(ctx)
	r.SetError(errors.NotFound("", msg))
	r.Exit()
}

func (c BaseController) Internal(ctx context.Context, err error) {
	r := c.GetRequest(ctx)
	// r.Response.Status = status
	r.SetError(errors.Internal("Internal server error", err))
	r.Exit()
}

func (c BaseController) Invalid(ctx context.Context, msg string) {
	r := c.GetRequest(ctx)
	// r.Response.Status = status
	r.SetError(errors.Invalid("Internal server error", msg))
	r.Exit()
}

func (c BaseController) Error(
	ctx context.Context,
	err error,
) {
	c.Fail(ctx, err)
}

func (c BaseController) UserID(ctx context.Context) (uuid.UUID, bool) {
	return auth.UserID(ctx)
}

func (c BaseController) MustUserID(ctx context.Context) uuid.UUID {
	id, _ := auth.UserID(ctx)
	return id
}

func (c BaseController) Can(ctx context.Context, permission string) bool {
	auth, ok := auth.GetAuthCtxFromCtx(ctx)
	if !ok {
		return false
	}

	return auth.HasPermission(permission)
}

// Is check if the user has a certain role
func (c BaseController) Is(ctx context.Context, role string) bool {
	auth, ok := auth.GetAuthCtxFromCtx(ctx)
	if !ok {
		return false
	}

	return auth.HasRole(role)
}

func (c BaseController) HasPolicy(ctx context.Context, resourceType, action string) bool {
	r := g.RequestFromCtx(ctx)
	if r == nil {
		return false
	}
	userID, ok := auth.UserID(ctx)
	if !ok {
		return false
	}
	decision, err := c.Authz.Authorize(ctx, auth.Request{
		UserID:       userID,
		Action:       action,
		ResourceType: resourceType,
		IP:           r.GetClientIp(),
		UserAgent:    r.Header.Get("User-Agent"),
		RequestID:    r.Header.Get("X-Request-Id"),
	})
	if err != nil {
		c.Forbidden(ctx, "authorization failed")
		return false
	}
	if !decision.Allowed {
		c.Unauthorized(
			ctx,
			auth.PermissionDeniedError(resourceType, action).Message,
		)
		return false
	}
	return true
}
