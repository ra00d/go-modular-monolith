package mappers

import (
	"modular-monolith/internal/auth/domain/access"
	"modular-monolith/internal/auth/infra/.gen/models"
	"modular-monolith/lib/utils"
)

func ToDomainPermission(row models.Permission) access.Permission {
	return access.Permission{
		ID:          row.ID,
		Name:        row.Name,
		Slug:        row.Slug,
		Description: row.Description,
		IsSystem:    row.IsSystem,
		Resource:    row.Resource,
		Action:      row.Action,
		Category:    utils.Deref(row.Category),
		Conditions:  utils.Deref(row.Conditions),
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}
}
