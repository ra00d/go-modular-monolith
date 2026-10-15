package services

import (
	"context"
	"fmt"

	"modular-monolith/internal/auth/domain/access"
	"modular-monolith/lib/db"
)

type PoliciesService struct {
	repo access.PolicyRepository
	db   *db.Manager
}

func NewPoliciesService(repo access.PolicyRepository, db *db.Manager) *PoliciesService {
	return &PoliciesService{
		repo, db,
	}
}

func (p PoliciesService) GetAll(ctx context.Context, limit, offset int) ([]access.Policy, int, error) {
	policies, count, err := p.repo.List(ctx, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("%w", err)
	}
	return policies, count, nil
}

func (p PoliciesService) GetByID(ctx context.Context, id access.PolicyID) (access.Policy, error) {
	policy, err := p.repo.GetByID(ctx, id)
	if err != nil {
		return access.Policy{}, fmt.Errorf("%w", err)
	}
	return policy, nil
}
