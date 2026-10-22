package services

import (
	"context"

	"modular-monolith/internal/modules/auth/domain/access"
	"modular-monolith/lib/errors"

	"github.com/google/uuid"
)

// APIKeyService serves the admin read model over issued api keys.
type APIKeyService struct {
	repo access.APIKeyRepository
}

func NewAPIKeyService(repo access.APIKeyRepository) *APIKeyService {
	return &APIKeyService{repo: repo}
}

func (s *APIKeyService) List(
	ctx context.Context,
	filter access.APIKeyFilter,
	limit, offset int,
) ([]access.APIKey, int, error) {
	keys, count, err := s.repo.List(ctx, filter, limit, offset)
	if err != nil {
		return nil, 0, errors.WrapRaw("list api keys", err)
	}
	return keys, count, nil
}

func (s *APIKeyService) GetByID(ctx context.Context, id int64) (access.APIKey, error) {
	key, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return access.APIKey{}, errors.WrapRaw("get api key", err)
	}
	return key, nil
}

// Revoke flips a key off. The stored hash is left in place so the audit trail
// can still identify which key was revoked.
func (s *APIKeyService) Revoke(ctx context.Context, id int64, actor uuid.UUID) error {
	if _, err := s.repo.GetByID(ctx, id); err != nil {
		return err
	}
	if err := s.repo.Revoke(ctx, id); err != nil {
		return errors.WrapRaw("revoke api key", err)
	}
	return nil
}
