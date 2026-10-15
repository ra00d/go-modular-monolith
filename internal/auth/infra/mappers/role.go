package mappers

import (
	"modular-monolith/internal/auth/domain/access"
	"modular-monolith/internal/auth/infra/.gen/models"
	"modular-monolith/lib/utils"
)

func ToDomainRole(row models.Role) access.Role {
	return access.Role{
		ID:           row.ID,
		Name:         row.Name,
		Slug:         row.Slug,
		Description:  row.Description,
		ParentRoleID: row.ParentRoleID,
		Level:        row.Level,
		IsSystem:     row.IsSystem,
		Metadata:     utils.Deref(row.Metadata),
		CreatedAt:    row.CreatedAt,
		UpdatedAt:    row.UpdatedAt,
	}
}
