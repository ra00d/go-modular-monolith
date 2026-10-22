package rbac

import (
	"context"
	"encoding/json"
	"time"

	"modular-monolith/internal/modules/auth/domain/access"
	"modular-monolith/internal/modules/auth/infra/.gen/models"
	"modular-monolith/internal/modules/auth/infra/.gen/table"
	"modular-monolith/lib/db"
	"modular-monolith/lib/utils"

	pg "github.com/go-jet/jet/v2/postgres"
	"github.com/jmoiron/sqlx"
)

var _ access.APIKeyRepository = (*APIKeyRepository)(nil)

type APIKeyRepository struct {
	db.BaseRepository
}

func NewAPIKeyRepository(DB *sqlx.DB) *APIKeyRepository {
	return &APIKeyRepository{
		BaseRepository: db.NewRepo(DB),
	}
}

// Create implements [access.APIKeyRepository].
func (a *APIKeyRepository) Create(
	ctx context.Context,
	in *access.IssuedAPIKey,
) (int64, error) {
	ak := table.APIKey
	permissions, err := marshalPermissions(in.Permissions)
	if err != nil {
		return 0, err
	}
	row := models.APIKey{
		Name:        in.Name,
		KeyPrefix:   in.KeyPrefix,
		KeyHash:     in.KeyHash,
		UserID:      in.UserID,
		RoleID:      in.RoleID,
		Permissions: permissions,
		RateLimit:   int32Ptr(in.RateLimit),
		ExpiresAt:   in.ExpiresAt,
		IsActive:    true,
		CreatedBy:   in.CreatedBy,
	}
	stmt := ak.
		INSERT(
			ak.Name, ak.KeyPrefix, ak.KeyHash, ak.UserID, ak.RoleID,
			ak.Permissions, ak.RateLimit, ak.ExpiresAt, ak.CreatedBy,
		).
		MODEL(row).
		RETURNING(ak.ID)

	if err := stmt.QueryContext(ctx, a.DB(ctx), &in.ID); err != nil {
		if db.IsUniqueViolation(err) {
			return 0, access.ErrAPIKeyInUse
		}
		return 0, err
	}
	return in.ID, nil
}

// GetByID implements [access.APIKeyRepository].
func (a *APIKeyRepository) GetByID(ctx context.Context, id int64) (access.APIKey, error) {
	return a.get(ctx, table.APIKey.ID.EQ(pg.Int(id)))
}

// GetByPrefix implements [access.APIKeyRepository].
func (a *APIKeyRepository) GetByPrefix(
	ctx context.Context,
	prefix string,
) (access.APIKey, error) {
	return a.get(ctx, table.APIKey.KeyPrefix.EQ(pg.String(prefix)))
}

func (a *APIKeyRepository) get(
	ctx context.Context,
	cond pg.BoolExpression,
) (access.APIKey, error) {
	ak := table.APIKey
	var row models.APIKey
	if err := ak.SELECT(ak.AllColumns).WHERE(pg.AND(cond, ak.DeletedAt.IS_NULL())).
		LIMIT(1).QueryContext(ctx, a.DB(ctx), &row); err != nil {
		if db.IsNoRowError(err) {
			return access.APIKey{}, access.ErrAPIKeyNotFound
		}
		return access.APIKey{}, err
	}
	return toAPIKey(row)
}

// List implements [access.APIKeyRepository].
func (a *APIKeyRepository) List(
	ctx context.Context,
	filter access.APIKeyFilter,
	limit, offset int,
) ([]access.APIKey, int, error) {
	ak := table.APIKey
	keys := []access.APIKey{}

	stmt := ak.SELECT(ak.AllColumns).WHERE(apiKeyFilters(filter))
	rows, count, err := a.Paginate[[]models.APIKey](ctx, stmt, limit, offset)
	if err != nil {
		return keys, 0, err
	}
	for _, v := range rows {
		key, err := toAPIKey(v)
		if err != nil {
			return keys, 0, err
		}
		keys = append(keys, key)
	}
	return keys, count, nil
}

// Update implements [access.APIKeyRepository].
func (a *APIKeyRepository) Update(ctx context.Context, in access.APIKey) error {
	ak := table.APIKey
	row := models.APIKey{
		Name:        in.Name,
		Permissions: marshalPermissionsOrNil(in.Permissions),
		RateLimit:   int32Ptr(in.RateLimit),
		ExpiresAt:   in.ExpiresAt,
		UpdatedAt:   time.Now(),
	}
	res, err := ak.UPDATE(ak.Name, ak.Permissions, ak.RateLimit, ak.ExpiresAt, ak.UpdatedAt).
		MODEL(row).
		WHERE(ak.ID.EQ(pg.Int(in.ID))).
		ExecContext(ctx, a.DB(ctx))
	if err != nil {
		return err
	}
	if n, err := res.RowsAffected(); err == nil && n == 0 {
		return access.ErrAPIKeyNotFound
	}
	return nil
}

// Revoke implements [access.APIKeyRepository].
func (a *APIKeyRepository) Revoke(ctx context.Context, id int64) error {
	ak := table.APIKey
	res, err := ak.UPDATE(ak.IsActive, ak.UpdatedAt).
		MODEL(models.APIKey{IsActive: false, UpdatedAt: time.Now()}).
		WHERE(ak.ID.EQ(pg.Int(id))).
		ExecContext(ctx, a.DB(ctx))
	if err != nil {
		return err
	}
	if n, err := res.RowsAffected(); err == nil && n == 0 {
		return access.ErrAPIKeyNotFound
	}
	return nil
}

// TouchLastUsed implements [access.APIKeyRepository].
func (a *APIKeyRepository) TouchLastUsed(ctx context.Context, id int64) error {
	ak := table.APIKey
	touchAt := time.Now()
	_, err := ak.UPDATE(ak.LastUsedAt).
		MODEL(models.APIKey{LastUsedAt: &touchAt}).
		WHERE(ak.ID.EQ(pg.Int(id))).
		ExecContext(ctx, a.DB(ctx))
	return err
}

func apiKeyFilters(filter access.APIKeyFilter) pg.BoolExpression {
	ak := table.APIKey
	conds := []pg.BoolExpression{ak.DeletedAt.IS_NULL()}
	if filter.UserID != nil {
		conds = append(conds, ak.UserID.EQ(pg.UUID(*filter.UserID)))
	}
	if filter.RoleID != nil {
		conds = append(conds, ak.RoleID.EQ(pg.Int(int64(*filter.RoleID))))
	}
	if filter.IsActive != nil {
		conds = append(conds, ak.IsActive.EQ(pg.Bool(*filter.IsActive)))
	}
	return pg.AND(conds...)
}

func marshalPermissions(permissions []string) (*string, error) {
	if permissions == nil {
		return nil, nil
	}
	raw, err := json.Marshal(permissions)
	if err != nil {
		return nil, err
	}
	return utils.Nullable(string(raw)), nil
}

// marshalPermissionsOrNil swallows the marshal error: a []string of slugs
// cannot fail to encode, and Update has no way to report it usefully.
func marshalPermissionsOrNil(permissions []string) *string {
	raw, _ := marshalPermissions(permissions)
	return raw
}

func toAPIKey(row models.APIKey) (access.APIKey, error) {
	var permissions []string
	if raw := utils.Deref(row.Permissions); raw != "" {
		if err := json.Unmarshal([]byte(raw), &permissions); err != nil {
			return access.APIKey{}, err
		}
	}
	rateLimit := int32PtrToInt(row.RateLimit)
	return access.APIKey{
		ID:          int64(row.ID),
		Name:        row.Name,
		KeyPrefix:   row.KeyPrefix,
		UserID:      row.UserID,
		RoleID:      row.RoleID,
		Permissions: permissions,
		RateLimit:   rateLimit,
		ExpiresAt:   row.ExpiresAt,
		LastUsedAt:  row.LastUsedAt,
		IsActive:    row.IsActive,
		CreatedBy:   row.CreatedBy,
		CreatedAt:   row.CreatedAt,
	}, nil
}
func int32Ptr(v *int) *int32 {
	if v == nil {
		return nil
	}
	out := int32(*v)
	return &out
}

// int32PtrToInt preserves the "no limit" distinction a nil column means, so
// the response carries null rather than a misleading zero.
func int32PtrToInt(v *int32) *int {
	if v == nil {
		return nil
	}
	out := int(*v)
	return &out
}
