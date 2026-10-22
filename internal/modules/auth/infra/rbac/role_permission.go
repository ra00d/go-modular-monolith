package rbac

import (
	"context"

	"modular-monolith/internal/modules/auth/domain/access"
	"modular-monolith/internal/modules/auth/infra/.gen/models"
	"modular-monolith/internal/modules/auth/infra/.gen/table"
	"modular-monolith/lib/db"
	"modular-monolith/lib/utils"

	pg "github.com/go-jet/jet/v2/postgres"
	"github.com/jmoiron/sqlx"
)

var _ access.RolePermissionRepository = (*RolePermissionRepository)(nil)

type RolePermissionRepository struct {
	db.BaseRepository
}

func NewRolePermissionRepository(DB *sqlx.DB) *RolePermissionRepository {
	return &RolePermissionRepository{
		BaseRepository: db.NewRepo(DB),
	}
}

// Grant implements [access.RolePermissionRepository].
func (r *RolePermissionRepository) Grant(
	ctx context.Context,
	grant *access.RolePermissionGrant,
) error {
	rp := table.RolePermission
	row := models.RolePermission{
		RoleID:       grant.RoleID,
		PermissionID: grant.PermissionID,
		Conditions:   utils.Nullable(grant.Conditions),
		ExpiresAt:    grant.ExpiresAt,
		GrantedBy:    grant.GrantedBy,
	}
	stmt := rp.
		INSERT(rp.RoleID, rp.PermissionID, rp.Conditions, rp.ExpiresAt, rp.GrantedBy).
		MODEL(row).
		RETURNING(rp.ID)

	if err := stmt.QueryContext(ctx, r.DB(ctx), &grant.ID); err != nil {
		if db.IsUniqueViolation(err) {
			return access.ErrGrantAlreadyExists
		}
		return err
	}
	return nil
}

// Revoke implements [access.RolePermissionRepository].
func (r *RolePermissionRepository) Revoke(
	ctx context.Context,
	roleID access.RoleID,
	permissionID access.PermissionID,
) error {
	stmt := table.RolePermission.DELETE().WHERE(pg.AND(
		table.RolePermission.RoleID.EQ(pg.Int(int64(roleID))),
		table.RolePermission.PermissionID.EQ(pg.Int(int64(permissionID))),
	))
	res, err := stmt.ExecContext(ctx, r.DB(ctx))
	if err != nil {
		return err
	}
	if n, err := res.RowsAffected(); err == nil && n == 0 {
		return access.ErrGrantNotFound
	}
	return nil
}

// Get implements [access.RolePermissionRepository].
func (r *RolePermissionRepository) Get(
	ctx context.Context,
	roleID access.RoleID,
	permissionID access.PermissionID,
) (*access.RolePermissionGrant, error) {
	rp := table.RolePermission
	var row models.RolePermission
	if err := rp.SELECT(rp.AllColumns).WHERE(pg.AND(
		rp.RoleID.EQ(pg.Int(int64(roleID))),
		rp.PermissionID.EQ(pg.Int(int64(permissionID))),
	)).LIMIT(1).QueryContext(ctx, r.DB(ctx), &row); err != nil {
		if db.IsNoRowError(err) {
			return nil, access.ErrGrantNotFound
		}
		return nil, err
	}
	return &access.RolePermissionGrant{
		ID:           row.ID,
		RoleID:       row.RoleID,
		PermissionID: row.PermissionID,
		Conditions:   utils.Deref(row.Conditions),
		ExpiresAt:    row.ExpiresAt,
		GrantedBy:    row.GrantedBy,
		CreatedAt:    row.CreatedAt,
	}, nil
}

// ListByRole implements [access.RolePermissionRepository].
func (r *RolePermissionRepository) ListByRole(
	ctx context.Context,
	roleID access.RoleID,
	limit, offset int,
) ([]access.RolePermission, int, error) {
	return r.list(ctx, table.RolePermission.RoleID.EQ(pg.Int(int64(roleID))), limit, offset)
}

// ListByPermission implements [access.RolePermissionRepository].
func (r *RolePermissionRepository) ListByPermission(
	ctx context.Context,
	permissionID access.PermissionID,
	limit, offset int,
) ([]access.RolePermission, int, error) {
	return r.list(
		ctx,
		table.RolePermission.PermissionID.EQ(pg.Int(int64(permissionID))),
		limit,
		offset,
	)
}

func (r *RolePermissionRepository) list(
	ctx context.Context,
	filter pg.BoolExpression,
	limit, offset int,
) ([]access.RolePermission, int, error) {
	rp := table.RolePermission
	grants := []access.RolePermission{}

	stmt := rp.SELECT(rp.AllColumns).WHERE(filter)
	rows, count, err := r.Paginate[[]models.RolePermission](ctx, stmt, limit, offset)
	if err != nil {
		return grants, 0, err
	}
	for _, v := range rows {
		grants = append(grants, toRolePermission(v))
	}
	return grants, count, nil
}

// UpdateConditions implements [access.RolePermissionRepository].
func (r *RolePermissionRepository) UpdateConditions(
	ctx context.Context,
	roleID access.RoleID,
	permissionID access.PermissionID,
	conditions string,
) error {
	rp := table.RolePermission
	res, err := rp.UPDATE(rp.Conditions).MODEL(models.RolePermission{
		Conditions: utils.Nullable(conditions),
	}).WHERE(pg.AND(
		rp.RoleID.EQ(pg.Int(int64(roleID))),
		rp.PermissionID.EQ(pg.Int(int64(permissionID))),
	)).ExecContext(ctx, r.DB(ctx))
	if err != nil {
		return err
	}
	if n, err := res.RowsAffected(); err == nil && n == 0 {
		return access.ErrGrantNotFound
	}
	return nil
}

// Exists implements [access.RolePermissionRepository].
func (r *RolePermissionRepository) Exists(
	ctx context.Context,
	roleID access.RoleID,
	permissionID access.PermissionID,
) (bool, error) {
	rp := table.RolePermission
	var count int
	if err := r.QueryOne(ctx, rp.SELECT(pg.COUNT(pg.Int(1))).WHERE(pg.AND(
		rp.RoleID.EQ(pg.Int(int64(roleID))),
		rp.PermissionID.EQ(pg.Int(int64(permissionID))),
	)), &count); err != nil {
		return false, err
	}
	return count > 0, nil
}

func toRolePermission(row models.RolePermission) access.RolePermission {
	return access.RolePermission{
		ID:           row.ID,
		RoleID:       row.RoleID,
		PermissionID: row.PermissionID,
		Conditions:   utils.Deref(row.Conditions),
		ExpiresAt:    row.ExpiresAt,
		GrantedBy:    row.GrantedBy,
		CreatedAt:    row.CreatedAt,
	}
}
