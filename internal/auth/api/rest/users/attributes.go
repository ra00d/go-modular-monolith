package users

import (
	"context"

	"modular-monolith/internal/auth/domain/users"
	"modular-monolith/lib/http"
	"modular-monolith/lib/utils"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

// UserAttribute is one ABAC attribute row (the user_attribute table).
// Attributes feed the policy engine; values are stored as strings and parsed
// according to ValueType.
type UserAttribute = users.UserAttribute

type ListUserAttributesReq struct {
	g.Meta `path:"/:id/attributes" method:"get" tags:"Users" x-group:"Admin/Users" summary:"List the ABAC attributes of a user" security:"bearer"`
	http.RequestBaseParams
	ID        uuid.UUID `json:"id"        v:"required"`
	Namespace *string   `json:"namespace"              dc:"restrict to one namespace"`
}
type ListUserAttributesRes []UserAttribute

func (c UsersController) ListUserAttributes(
	ctx context.Context,
	req *ListUserAttributesReq,
) (res http.Response[ListUserAttributesRes], err error) {
	p := c.Pagination(ctx)
	attributes, count, err := c.usersService.ListAttributes(
		ctx, req.ID, req.Namespace, p.PerPage, p.Offset,
	)
	if err != nil {
		c.Fail(ctx, err)
	}
	return c.OKWithPagination(ctx, ListUserAttributesRes(attributes), count), nil
}

type SetUserAttributeReq struct {
	g.Meta    `path:"/:id/attributes" method:"put" tags:"Users" x-group:"Admin/Users" summary:"Create or replace one ABAC attribute of a user" security:"bearer"`
	ID        uuid.UUID `                                                                                                                                                  json:"id"        v:"required"`
	Namespace string    `                                                                                                                                                  json:"namespace" v:"length:1,64"                          d:"default" dc:"org, project, system or default"`
	Key       string    `                                                                                                                                                  json:"key"       v:"required|regex:^[a-zA-Z0-9_.]{1,64}$"             dc:"attribute name"`
	Value     string    `                                                                                                                                                  json:"value"     v:"required|max-length:4096"                         dc:"the value, stored as a string"`
	ValueType string    `                                                                                                                                                  json:"valueType" v:"in:string,number,boolean,json,date"   d:"string"  dc:"how the value is parsed during policy evaluation"`
}
type SetUserAttributeRes struct {
	g.Meta    `mime:"application/json"`
	UserID    uuid.UUID `                        json:"userId"    dc:"the user"`
	Namespace string    `                        json:"namespace" dc:"the namespace written"`
	Key       string    `                        json:"key"       dc:"the attribute written"`
}

func (c UsersController) SetUserAttribute(
	ctx context.Context,
	req *SetUserAttributeReq,
) (res http.Response[SetUserAttributeRes], err error) {
	actor, ok := c.UserID(ctx)
	if !ok {
		c.Unauthorized(ctx, "set user attribute")
		return res, nil
	}
	namespace := utils.DerefOr(&req.Namespace, users.DefaultNamespace)
	err = c.usersService.SetAttribute(ctx, users.UserAttribute{
		UserID:    req.ID,
		Namespace: namespace,
		Key:       req.Key,
		Value:     req.Value,
		ValueType: req.ValueType,
	}, actor)
	if err != nil {
		c.Fail(ctx, err)
	}
	return c.OK(SetUserAttributeRes{
		UserID:    req.ID,
		Namespace: namespace,
		Key:       req.Key,
	}), nil
}

type DeleteUserAttributeReq struct {
	g.Meta    `path:"/:id/attributes/:key" method:"delete" tags:"Users" x-group:"Admin/Users" summary:"Delete one ABAC attribute of a user" security:"bearer"`
	ID        uuid.UUID `                                                                                                                                               json:"id"        v:"required"`
	Key       string    `                                                                                                                                               json:"key"       v:"required" dc:"the attribute name"`
	Namespace *string   `                                                                                                                                               json:"namespace"              dc:"namespace the attribute lives in" d:"default"`
}
type DeleteUserAttributeRes struct {
	g.Meta  `mime:"application/json"`
	UserID  uuid.UUID `                        json:"userId"  dc:"the user"`
	Key     string    `                        json:"key"     dc:"the deleted attribute"`
	Deleted bool      `                        json:"deleted" dc:"always true on success"`
}

func (c UsersController) DeleteUserAttribute(
	ctx context.Context,
	req *DeleteUserAttributeReq,
) (res http.Response[DeleteUserAttributeRes], err error) {
	actor, ok := c.UserID(ctx)
	if !ok {
		c.Unauthorized(ctx, "delete user attribute")
		return res, nil
	}
	namespace := utils.Deref(req.Namespace)
	if err := c.usersService.DeleteAttribute(ctx, req.ID, namespace, req.Key, actor); err != nil {
		c.Fail(ctx, err)
	}
	return c.OK(DeleteUserAttributeRes{
		UserID:  req.ID,
		Key:     req.Key,
		Deleted: true,
	}), nil
}
