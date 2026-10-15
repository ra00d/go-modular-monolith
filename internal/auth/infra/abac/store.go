package abac

import (
	"context"
	"fmt"
	"strconv"
	"time"

	model "modular-monolith/internal/auth/infra/.gen/models"
	table "modular-monolith/internal/auth/infra/.gen/table"
	libauth "modular-monolith/lib/auth"

	pg "github.com/go-jet/jet/v2/postgres"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/jmoiron/sqlx"
)

var _ libauth.PolicyStore = (*PolicyStore)(nil)

// PolicyStore reads ABAC policies from the auth.abac_policy table.
type PolicyStore struct {
	db *sqlx.DB
}

// NewPolicyStore builds the store.
func NewPolicyStore(db *sqlx.DB) *PolicyStore { return &PolicyStore{db: db} }

// GetPoliciesForResource returns active, non-expired policies for a resource
// type, ordered by priority (highest first). Policies whose rules JSON fails to
// parse are skipped so one malformed row cannot break authorization.
func (s *PolicyStore) GetPoliciesForResource(
	ctx context.Context,
	resourceType string,
) ([]libauth.Policy, error) {
	t := table.AbacPolicy
	stmt := pg.SELECT(t.AllColumns).
		FROM(t).
		WHERE(
			t.Resource.EQ(pg.String(resourceType)).
				AND(t.IsActive.IS_TRUE()).
				AND(t.DeletedAt.IS_NULL()).
				AND(t.ExpiresAt.IS_NULL().OR(t.ExpiresAt.GT(pg.TimestampT(time.Now())))),
		).
		ORDER_BY(t.Priority.ASC(), t.CreatedAt.ASC())

	var rows []model.AbacPolicy
	if err := stmt.QueryContext(ctx, s.db, &rows); err != nil {
		return nil, fmt.Errorf("query abac policies: %w", err)
	}

	policies := make([]libauth.Policy, 0, len(rows))
	for _, r := range rows {
		p, err := toPolicy(r)
		if err != nil {
			// ponytail: skip a malformed policy rather than deny-all the app;
			// surface it via logs/monitoring when an admin editor exists.
			continue
		}
		policies = append(policies, p)
	}
	return policies, nil
}

func toPolicy(m model.AbacPolicy) (libauth.Policy, error) {
	var rules libauth.RuleSet

	if err := gconv.Struct(m.Rules, &rules); err != nil {
		return libauth.Policy{}, fmt.Errorf("parse rules for policy %s: %w", m.Slug, err)
	}

	p := libauth.Policy{
		ID:       strconv.Itoa(int(m.ID)),
		Name:     m.Name,
		Slug:     m.Slug,
		Resource: m.Resource,
		Action:   m.Action,
		Rules:    rules,
		Effect:   libauth.Effect(m.Effect),
		Priority: int(m.Priority),
		IsActive: m.IsActive,
	}
	if m.Description != nil {
		p.Description = *m.Description
	}
	if m.ExpiresAt != nil {
		p.ExpiresAt = m.ExpiresAt
	}
	if m.CreatedBy != nil {
		p.CreatedBy = m.CreatedBy.String()
	}
	return p, nil
}
