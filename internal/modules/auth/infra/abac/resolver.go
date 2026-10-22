package abac

import (
	"context"
	"encoding/json"
	"strconv"

	model "modular-monolith/internal/modules/auth/infra/.gen/models"
	table "modular-monolith/internal/modules/auth/infra/.gen/table"
	libauth "modular-monolith/lib/auth"

	pg "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

var _ libauth.AttributeResolver = (*AttributeResolver)(nil)

// AttributeResolver loads dynamic subject attributes from auth.user_attribute.
// Resource attributes are supplied by the caller (they are resource-specific),
// so ResolveResourceAttributes is a no-op.
type AttributeResolver struct {
	db *sqlx.DB
}

// NewAttributeResolver builds the resolver.
func NewAttributeResolver(db *sqlx.DB) *AttributeResolver { return &AttributeResolver{db: db} }

// ResolveSubjectAttributes returns the user's stored attributes keyed by
// "namespace.key" (and bare "key" for the default namespace), typed per
// value_type so policy comparisons behave correctly.
func (r *AttributeResolver) ResolveSubjectAttributes(
	ctx context.Context,
	subject libauth.Subject,
) (map[string]any, error) {
	if subject.ID == "" {
		return map[string]any{}, nil
	}

	t := table.UserAttribute
	stmt := pg.SELECT(t.Namespace, t.Key, t.Value, t.ValueType).
		FROM(t).
		WHERE(t.UserID.EQ(pg.UUID(uuid.MustParse(subject.ID))))

	var rows []model.UserAttribute
	if err := stmt.QueryContext(ctx, r.db, &rows); err != nil {
		return nil, err
	}

	attrs := make(map[string]any, len(rows)*2)
	for _, row := range rows {
		v := parseValue(row.Value, row.ValueType)
		attrs[row.Namespace+"."+row.Key] = v
		if row.Namespace == "default" {
			attrs[row.Key] = v
		}
	}
	return attrs, nil
}

// ResolveResourceAttributes is a no-op; callers pass resource attributes.
func (r *AttributeResolver) ResolveResourceAttributes(
	ctx context.Context,
	resource libauth.Resource,
) (map[string]any, error) {
	return nil, nil
}

// parseValue converts the stored string to its typed representation.
func parseValue(raw, valueType string) any {
	switch valueType {
	case "number":
		if f, err := strconv.ParseFloat(raw, 64); err == nil {
			return f
		}
	case "boolean":
		if b, err := strconv.ParseBool(raw); err == nil {
			return b
		}
	case "json":
		var v any
		if err := json.Unmarshal([]byte(raw), &v); err == nil {
			return v
		}
	}
	// string, date, or failed parse: keep the raw string.
	return raw
}
