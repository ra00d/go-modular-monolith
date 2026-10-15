package http

import (
	"context"
	"net/http"

	"modular-monolith/lib/common"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/goai"
)

type (
	EmptyResponse struct{}
)

type GeneralError struct {
	Code    string `json:"code" dc:"A code describing the error if any"`
	Message string `json:"message,omitempty" dc:"the error  message"`
	Status  int    `json:"status" dc:"The http status code of the response 400,500"`
	Stack   any    `json:"-"`
}
type InternalErrorResponse struct {
	GeneralError
}
type UnAuthenticatedResponse struct {
	g.Meta `dc:"User is not authenticated" status:"401"`
	GeneralError
}
type UnAuthorizedResponse struct {
	g.Meta `dc:"User is not authorized to access this resource" status:"401"`
	GeneralError
}

type BadRequestResponse struct {
	g.Meta `status:"400" dc:"Invalid input"`
	GeneralError
}

type InvalidInputErrorResponse struct {
	g.Meta `               dc:"User input are invalid and not processable" status:"422"`
	GeneralError
	Errors map[string]any `json:"errors" dc:"returned errors in post requests with body"`
}

type Response[T any] struct {
	// g.Meta  `               mime:"application/json"`
	Data    T                      `json:"data,omitempty"              dc:"The returned data if the request is OK"`
	Message string                 `json:"message,omitempty"`
	Meta    *common.PaginationMeta `json:"meta,omitempty" dc:"The pagination data for list response"`
	Errors  map[string]any         `json:"-"  dc:"returned errors in post requests with body"`
	Code    string                 `json:"-" dc:"A code describing the error if any"`
	Status  int                    `json:"status" dc:"The http status code of the response 200,300,400,500"`
	Stack   any                    `json:"-"`
}

type MockResponse struct {
	// g.Meta  `               mime:"application/json"`
	Data    any                    `json:"data,omitempty"              dc:"The returned data if the request is OK"`
	Message string                 `json:"message,omitempty"`
	Meta    *common.PaginationMeta `json:"meta,omitempty" dc:"The pagination data for list response"`
	Errors  map[string]any         `json:"-"  dc:"returned errors in post requests with body"`
	Code    string                 `json:"-" dc:"A code describing the error if any"`
	Status  int                    `json:"status" dc:"The http status code of the response 200,300,400,500"`
	Stack   any                    `json:"-"`
}

func (c Response[any]) EnhanceResponseStatus() map[goai.EnhancedStatusCode]goai.EnhancedStatusType {
	return map[goai.EnhancedStatusCode]goai.EnhancedStatusType{
		400: {
			Response: BadRequestResponse{},
		},
		500: {
			Response: InternalErrorResponse{},
		},
		401: {
			Response: UnAuthenticatedResponse{},
		},
		403: {
			Response: UnAuthenticatedResponse{},
		},
		422: {
			Response: InvalidInputErrorResponse{},
		},
	}
}

func (c BaseController) Created[T any](data T) Response[T] {
	return Response[T]{
		Data:    data,
		Message: "",
		Status:  http.StatusCreated,
	}
}

func (c BaseController) OK[T any](data T) Response[T] {
	return Response[T]{
		Data:   data,
		Status: http.StatusAccepted,
	}
}

func (c BaseController) OKWithPagination[T ~[]S, S any](ctx context.Context, data T, count int, cursor ...string) Response[T] {
	return Response[T]{
		Data:   data,
		Meta:   c.PaginationMeta(ctx, count, cursor...),
		Status: http.StatusAccepted,
	}
}
