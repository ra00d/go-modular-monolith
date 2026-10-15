// Package auth provides
package auth

import (
	"context"
	"strings"

	"modular-monolith/lib/logger"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

type AuthCtx struct {
	UserID uuid.UUID

	Status string

	Permissions []string
	Roles       []string
}

type AuthClaims struct {
	UserID      uuid.UUID
	Status      string
	Roles       []string
	Permissions []string
}

func NewAuthContext(claims AuthClaims) *AuthCtx {
	return &AuthCtx{
		UserID:      claims.UserID,
		Status:      claims.Status,
		Roles:       claims.Roles,
		Permissions: claims.Permissions,
	}
}

func GetAuthCtxFromCtx(ctx context.Context) (AuthCtx, bool) {
	r := g.RequestFromCtx(ctx)

	userID := r.GetCtxVar(ctxUserID)
	permissions := r.GetCtxVar(ctxPermissions)
	roles := r.GetCtxVar(ctxRoles)
	status := r.GetCtxVar(ctxUserStatus)
	switch {
	case userID.IsNil():
	case status.IsNil():
	case permissions.IsNil():
	case roles.IsNil():
		return AuthCtx{}, false
	case userID.IsEmpty():
	case status.IsEmpty():
	case permissions.IsEmpty():
	case roles.IsEmpty():
		return AuthCtx{}, false
	default:
		userUUID, err := uuid.Parse(userID.String())
		if err != nil {
			logger.Error(err.Error())
			return AuthCtx{}, false
		}

		return AuthCtx{
			UserID:      userUUID,
			Status:      status.String(),
			Roles:       roles.Strings(),
			Permissions: permissions.Strings(),
		}, true
	}
	return AuthCtx{}, false
}

// CombineAuthErrors Combine multiple errors
func CombineAuthErrors(errs ...error) error {
	var messages []string
	for _, err := range errs {
		if err != nil {
			messages = append(messages, err.Error())
		}
	}
	if len(messages) == 0 {
		return nil
	}
	return NewAuthError(strings.Join(messages, "; "), nil)
}
