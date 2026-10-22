// Package abac provides the PostgreSQL-backed ABAC policy store, attribute
// resolver and policy cache that wire the lib/auth policy engine to the auth
// module's schema.
package abac

import (
	"modular-monolith/internal/modules/auth/domain/access"
	"modular-monolith/internal/modules/auth/infra/.gen/models"
)

type Policy = models.AbacPolicy

func toDomainPolicy(p Policy) access.Policy {
	return access.Policy{
		ID:          p.ID,
		Name:        p.Name,
		Slug:        p.Slug,
		Description: *p.Description,
		Resource:    p.Resource,
		Action:      p.Action,
		Rules:       *p.Rules,
		Effect:      access.Effect(p.Effect),
		Priority:    p.Priority,
		IsActive:    p.IsActive,
		ExpiresAt:   p.ExpiresAt,
		CreatedBy:   p.CreatedBy,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}
