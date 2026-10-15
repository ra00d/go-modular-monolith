package rbac

import (
	"context"
	"time"

	"modular-monolith/internal/auth/domain/access"
	"modular-monolith/internal/auth/infra/.gen/models"
	"modular-monolith/internal/auth/infra/.gen/table"
	"modular-monolith/internal/auth/infra/mappers"
	"modular-monolith/lib/db"
	"modular-monolith/lib/utils"

	pg "github.com/go-jet/jet/v2/postgres"
	"github.com/jmoiron/sqlx"
)

var _ access.RoleRepository = (*RoleRepository)(nil)

type RoleRepository struct {
	db.BaseRepository
}

func NewRoleRepository(DB *sqlx.DB) *RoleRepository {
	return &RoleRepository{
		BaseRepository: db.NewRepo(DB),
	}
}

// Create implements [access.RoleRepository].
func (r *RoleRepository) Create(ctx context.Context, in access.Role) error {
	row := models.Role{
		Name:         in.Name,
		Slug:         in.Slug,
		Description:  in.Description,
		ParentRoleID: in.ParentRoleID,
		Level:        in.Level,
		Metadata:     utils.Nullable(in.Metadata),
	}
	stmt := table.Role.
		INSERT(
			table.Role.Name, table.Role.Slug, table.Role.Description,
			table.Role.ParentRoleID, table.Role.Level, table.Role.Metadata,
		).
		MODEL(row).
		RETURNING(table.Role.ID)

	if err := stmt.QueryContext(ctx, r.DB(ctx), &in.ID); err != nil {
		if db.IsUniqueViolation(err) {
			return access.ErrRoleSlugTaken
		}
		return err
	}
	return nil
}

// Delete implements [access.RoleRepository].
func (r *RoleRepository) Delete(ctx context.Context, id access.RoleID) error {
	// Deleting a role that still has children or assignments would either
	// orphan the subtree or silently strip somebody's access.
	var dependents int
	if err := r.QueryOne(ctx, table.Role.SELECT(pg.COUNT(pg.Int(1))).
		WHERE(table.Role.ParentRoleID.EQ(pg.Int(int64(id)))), &dependents); err != nil {
		return err
	}
	if dependents > 0 {
		return access.ErrRoleHasChildren
	}
	if err := r.QueryOne(ctx, table.UserRole.SELECT(pg.COUNT(pg.Int(1))).
		WHERE(table.UserRole.RoleID.EQ(pg.Int(int64(id)))), &dependents); err != nil {
		return err
	}
	if dependents > 0 {
		return access.ErrRoleHasAssignments
	}

	res, err := table.Role.DELETE().
		WHERE(table.Role.ID.EQ(pg.Int(int64(id)))).
		ExecContext(ctx, r.DB(ctx))
	if err != nil {
		return err
	}
	if n, err := res.RowsAffected(); err == nil && n == 0 {
		return access.ErrRoleNotFound
	}
	return nil
}

// ExistsID implements [access.RoleRepository].
func (r *RoleRepository) ExistsID(ctx context.Context, id access.RoleID) (bool, error) {
	return r.exists(ctx, table.Role.ID.EQ(pg.Int(int64(id))))
}

// ExistsSlug implements [access.RoleRepository].
func (r *RoleRepository) ExistsSlug(ctx context.Context, slug string) (bool, error) {
	return r.exists(ctx, table.Role.Slug.EQ(pg.String(slug)))
}

func (r *RoleRepository) exists(
	ctx context.Context,
	cond pg.BoolExpression,
) (bool, error) {
	var count int
	if err := r.QueryOne(ctx, table.Role.SELECT(pg.COUNT(pg.Int(1))).WHERE(cond), &count); err != nil {
		return false, err
	}
	return count > 0, nil
}

// GetByID implements [access.RoleRepository].
func (r *RoleRepository) GetByID(
	ctx context.Context,
	id access.RoleID,
) (access.Role, error) {
	return r.get(ctx, table.Role.ID.EQ(pg.Int(int64(id))))
}

// GetBySlug implements [access.RoleRepository].
func (r *RoleRepository) GetBySlug(
	ctx context.Context,
	slug string,
) (access.Role, error) {
	return r.get(ctx, table.Role.Slug.EQ(pg.String(slug)))
}

func (r *RoleRepository) get(
	ctx context.Context,
	cond pg.BoolExpression,
) (access.Role, error) {
	role := table.Role
	var row models.Role
	if err := role.SELECT(role.AllColumns).WHERE(cond).LIMIT(1).
		QueryContext(ctx, r.DB(ctx), &row); err != nil {
		if db.IsNoRowError(err) {
			return access.Role{}, access.ErrRoleNotFound
		}
		return access.Role{}, err
	}
	return mappers.ToDomainRole(row), nil
}

// List implements [access.RoleRepository].
func (r *RoleRepository) List(
	ctx context.Context,
	limit, offset int,
) ([]access.Role, int, error) {
	role := table.Role
	filters := []pg.BoolExpression{role.DeletedAt.IS_NULL()}
	roles := []access.Role{}

	stmt := role.SELECT(role.AllColumns).WHERE(pg.AND(filters...))
	// fil := role.ID.EQ(pg.Int(5)).AND(role.Slug.EQ(pg.String("guest")))
	rows, count, err := r.Paginate[[]models.Role](ctx, stmt, limit, offset)
	if err != nil {
		return roles, 0, err
	}
	for _, v := range rows {
		roles = append(roles, mappers.ToDomainRole(v))
	}
	return roles, count, nil
}

// ListByParent implements [access.RoleRepository].
func (r *RoleRepository) ListByParent(
	ctx context.Context,
	parentID access.RoleID,
) ([]access.Role, error) {
	role := table.Role
	roles := []access.Role{}
	var rows []models.Role

	stmt := role.SELECT(role.AllColumns).WHERE(pg.AND(
		role.DeletedAt.IS_NULL(),
		role.ParentRoleID.EQ(pg.Int(int64(parentID))),
	))
	if err := r.QueryMany(ctx, stmt, &rows); err != nil {
		return roles, err
	}
	for _, v := range rows {
		roles = append(roles, mappers.ToDomainRole(v))
	}
	return roles, nil
}

// Update implements [access.RoleRepository].
func (r *RoleRepository) Update(ctx context.Context, in *access.Role) error {
	role := table.Role
	row := models.Role{
		Name:         in.Name,
		Description:  in.Description,
		ParentRoleID: in.ParentRoleID,
		Level:        in.Level,
		Metadata:     utils.Nullable(in.Metadata),
		UpdatedAt:    time.Now(),
	}
	stmt := role.UPDATE(
		role.Name, role.Description, role.ParentRoleID, role.Level, role.Metadata, role.UpdatedAt,
	).MODEL(row).WHERE(pg.AND(
		role.ID.EQ(pg.Int(int64(in.ID))),
		role.DeletedAt.IS_NULL(),
	))

	res, err := stmt.ExecContext(ctx, r.DB(ctx))
	if err != nil {
		return err
	}
	if n, err := res.RowsAffected(); err == nil && n == 0 {
		return access.ErrRoleNotFound
	}
	return nil
}
