// Package mappers handles mapping db models to domain models
package mappers

import (
	"modular-monolith/internal/auth/domain/users"
	"modular-monolith/internal/auth/infra/.gen/models"
	"modular-monolith/lib/utils"
)

func ToDomainUser(m models.User) users.User {
	return users.User{
		ID:            m.ID,
		Name:          m.Name,
		Email:         utils.Deref(m.Email),
		Phone:         utils.Deref(m.Phone),
		Username:      utils.Deref(m.Username),
		PasswordHash:  utils.Deref(m.PasswordHash),
		EmailVerified: m.EmailVerified,
		PhoneVerified: m.PhoneVerified,
		Image:         utils.Deref(m.Image),
		Status:        m.Status,
		MFAEnabled:    m.MfaEnabled,
		Preference:    *m.Preferences,
		CreatedAt:     m.CreatedAt,
		UpdatedAt:     m.UpdatedAt,
	}
}
