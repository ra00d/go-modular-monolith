package abac

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"modular-monolith/internal/modules/auth/domain/access"
	"modular-monolith/internal/modules/auth/infra/.gen/table"
	"modular-monolith/lib/db"

	pg "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/jmoiron/sqlx"
)

var (
	// pt    AbacPolicy table
	pt                         = table.AbacPolicy
	_  access.PolicyRepository = (*PolicyRepository)(nil)
)

type PolicyRepository struct {
	db.BaseRepository
}

func NewPolicyRepository(DB *sqlx.DB) *PolicyRepository {
	return &PolicyRepository{
		BaseRepository: db.NewRepo(DB),
	}
}

// Create implements [access.PolicyRepository].
func (p PolicyRepository) Create(ctx context.Context, in *access.Policy) error {
	panic("unimplemented")
}

// Delete implements [access.PolicyRepository].
func (p PolicyRepository) Delete(ctx context.Context, id access.PolicyID) error {
	panic("unimplemented")
}

// ExistsID implements [access.PolicyRepository].
func (p PolicyRepository) ExistsID(ctx context.Context, id access.PolicyID) (bool, error) {
	panic("unimplemented")
}

// ExistsSlug implements [access.PolicyRepository].
func (p PolicyRepository) ExistsSlug(ctx context.Context, slug string) (bool, error) {
	panic("unimplemented")
}

// GetByID implements [access.PolicyRepository].
func (p PolicyRepository) GetByID(
	ctx context.Context,
	id access.PolicyID,
) (access.Policy, error) {
	stmt := pt.SELECT(pt.AllColumns).WHERE(pt.ID.EQ(pg.Int(int64(id))))
	policy, err := p.queryOne(ctx, stmt)
	if err != nil {
		return access.Policy{}, fmt.Errorf(": %w", err)
	}
	return policy, nil
}

// GetBySlug implements [access.PolicyRepository].
func (p PolicyRepository) GetBySlug(ctx context.Context, slug string) (access.Policy, error) {
	panic("unimplemented")
}

// List implements [access.PolicyRepository].
func (p PolicyRepository) List(ctx context.Context, limit, offset int) ([]access.Policy, int, error) {
	rows, count, err := p.Paginate[[]Policy](ctx, pt.SELECT(pt.AllColumns), limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf(": %w", err)
	}
	policies := make([]access.Policy, 0, len(rows))
	for _, v := range rows {
		policies = append(policies, toDomainPolicy(v))
	}
	return policies, count, nil
}

// ListByResourceAction implements [access.PolicyRepository].
func (p PolicyRepository) ListByResourceAction(
	ctx context.Context,
	resource string,
	action string,
) ([]access.Policy, error) {
	panic("unimplemented")
}

// Update implements [access.PolicyRepository].
func (p PolicyRepository) Update(ctx context.Context, in *access.Policy) error {
	panic("unimplemented")
}

func (p PolicyRepository) queryOne(
	ctx context.Context,
	stmt pg.SelectStatement,
) (access.Policy, error) {
	var row Policy
	if err := stmt.QueryContext(ctx, p.DB(ctx), &row); err != nil {
		if errors.Is(err, sql.ErrNoRows) || errors.Is(err, qrm.ErrNoRows) {
			return access.Policy{}, access.ErrPolicyNotFound
		}
		return access.Policy{}, err
	}
	return toDomainPolicy(row), nil
}
