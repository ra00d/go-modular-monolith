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

var _ access.PermissionRepository = (*PermissionRepository)(nil)

type PermissionRepository struct {
	db.BaseRepository
}

func NewPermissionRepository(DB *sqlx.DB) *PermissionRepository {
	return &PermissionRepository{
		BaseRepository: db.NewRepo(DB),
	}
}

// Create implements [access.PermissionRepository].
func (p *PermissionRepository) Create(ctx context.Context, in *access.Permission) error {
	perm := table.Permission
	row := models.Permission{
		Name:        in.Name,
		Slug:        in.Slug,
		Description: in.Description,
		Resource:    in.Resource,
		Action:      in.Action,
		Category:    utils.Nullable(in.Category),
		Conditions:  utils.Nullable(in.Conditions),
	}
	stmt := perm.
		INSERT(
			perm.Name, perm.Slug, perm.Description, perm.Resource,
			perm.Action, perm.Category, perm.Conditions,
		).
		MODEL(row).
		RETURNING(perm.ID)

	if err := stmt.QueryContext(ctx, p.DB(ctx), &in.ID); err != nil {
		if db.IsUniqueViolation(err) {
			return access.ErrPermissionSlugTaken
		}
		return err
	}
	return nil
}

// Delete implements [access.PermissionRepository].
func (p *PermissionRepository) Delete(ctx context.Context, id access.PermissionID) error {
	// Grants cascade at the database level, but a permission still held by a
	// role is a mistake worth surfacing rather than silently cascading.
	var grants int
	if err := p.QueryOne(ctx, table.RolePermission.SELECT(pg.COUNT(pg.Int(1))).
		WHERE(table.RolePermission.PermissionID.EQ(pg.Int(int64(id)))), &grants); err != nil {
		return err
	}
	if grants > 0 {
		return access.ErrPermissionHasGrants
	}

	res, err := table.Permission.DELETE().
		WHERE(table.Permission.ID.EQ(pg.Int(int64(id)))).
		ExecContext(ctx, p.DB(ctx))
	if err != nil {
		return err
	}
	if n, err := res.RowsAffected(); err == nil && n == 0 {
		return access.ErrPermissionNotFound
	}
	return nil
}

// ExistsID implements [access.PermissionRepository].
func (p *PermissionRepository) ExistsID(
	ctx context.Context,
	id access.PermissionID,
) (bool, error) {
	return p.exists(ctx, table.Permission.ID.EQ(pg.Int(int64(id))))
}

// ExistsSlug implements [access.PermissionRepository].
func (p *PermissionRepository) ExistsSlug(ctx context.Context, slug string) (bool, error) {
	return p.exists(ctx, table.Permission.Slug.EQ(pg.String(slug)))
}

func (p *PermissionRepository) exists(
	ctx context.Context,
	cond pg.BoolExpression,
) (bool, error) {
	var count int
	if err := p.QueryOne(
		ctx,
		table.Permission.SELECT(pg.COUNT(pg.Int(1))).WHERE(cond),
		&count,
	); err != nil {
		return false, err
	}
	return count > 0, nil
}

// GetByID implements [access.PermissionRepository].
func (p *PermissionRepository) GetByID(
	ctx context.Context,
	id access.PermissionID,
) (access.Permission, error) {
	return p.get(ctx, table.Permission.ID.EQ(pg.Int(int64(id))))
}

// GetBySlug implements [access.PermissionRepository].
func (p *PermissionRepository) GetBySlug(
	ctx context.Context,
	slug string,
) (access.Permission, error) {
	return p.get(ctx, table.Permission.Slug.EQ(pg.String(slug)))
}

func (p *PermissionRepository) get(
	ctx context.Context,
	cond pg.BoolExpression,
) (access.Permission, error) {
	perm := table.Permission
	var row models.Permission
	if err := perm.SELECT(perm.AllColumns).WHERE(cond).LIMIT(1).
		QueryContext(ctx, p.DB(ctx), &row); err != nil {
		if db.IsNoRowError(err) {
			return access.Permission{}, access.ErrPermissionNotFound
		}
		return access.Permission{}, err
	}
	return mappers.ToDomainPermission(row), nil
}

// List implements [access.PermissionRepository].
func (p *PermissionRepository) List(ctx context.Context, limit, offset int) ([]access.Permission, int, error) {
	return p.list(ctx, nil, limit, offset)
}

// ListByCategory implements [access.PermissionRepository].
func (p *PermissionRepository) ListByCategory(
	ctx context.Context,
	category string,
	limit, offset int,
) ([]access.Permission, int, error) {
	return p.list(ctx, table.Permission.Category.EQ(pg.String(category)), limit, offset)
}

// ListByResource implements [access.PermissionRepository].
func (p *PermissionRepository) ListByResource(
	ctx context.Context,
	resource string,
	limit, offset int,
) ([]access.Permission, int, error) {
	return p.list(ctx, table.Permission.Resource.EQ(pg.String(resource)), limit, offset)
}

func (p *PermissionRepository) list(
	ctx context.Context,
	filter pg.BoolExpression,
	limit, offset int,
) ([]access.Permission, int, error) {
	permission := table.Permission
	permissions := []access.Permission{}

	stmt := permission.SELECT(permission.AllColumns).WHERE(permission.DeletedAt.IS_NULL())
	if filter != nil {
		stmt = stmt.WHERE(filter)
	}
	rows, count, err := p.Paginate[[]models.Permission](ctx, stmt, limit, offset)
	if err != nil {
		return permissions, 0, err
	}
	for _, v := range rows {
		permissions = append(permissions, mappers.ToDomainPermission(v))
	}
	return permissions, count, nil
}

// Update implements [access.PermissionRepository].
func (p *PermissionRepository) Update(ctx context.Context, in *access.Permission) error {
	perm := table.Permission
	row := models.Permission{
		Name:        in.Name,
		Description: in.Description,
		Category:    utils.Nullable(in.Category),
		Conditions:  utils.Nullable(in.Conditions),
		UpdatedAt:   time.Now(),
	}
	stmt := perm.UPDATE(
		perm.Name, perm.Description, perm.Category, perm.Conditions, perm.UpdatedAt,
	).MODEL(row).WHERE(pg.AND(
		perm.ID.EQ(pg.Int(int64(in.ID))),
		perm.DeletedAt.IS_NULL(),
	))

	res, err := stmt.ExecContext(ctx, p.DB(ctx))
	if err != nil {
		return err
	}
	if n, err := res.RowsAffected(); err == nil && n == 0 {
		return access.ErrPermissionNotFound
	}
	return nil
}
