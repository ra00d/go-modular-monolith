package rbac

import (
	"context"
	"time"

	"modular-monolith/internal/auth/domain/access"
	"modular-monolith/internal/auth/infra/.gen/models"
	"modular-monolith/internal/auth/infra/.gen/table"
	"modular-monolith/lib/db"

	pg "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

var _ access.UserRoleRepository = (*UserRoleRepository)(nil)

type UserRoleRepository struct {
	db.BaseRepository
}

func NewUserRoleRepository(DB *sqlx.DB) *UserRoleRepository {
	return &UserRoleRepository{
		BaseRepository: db.NewRepo(DB),
	}
}

// Assign implements [access.UserRoleRepository].
func (u *UserRoleRepository) Assign(
	ctx context.Context,
	assignment *access.UserRoleAssignment,
) error {
	ur := table.UserRole
	row := models.UserRole{
		UserID:     assignment.UserID,
		RoleID:     assignment.RoleID,
		AssignedAt: time.Now(),
		AssignedBy: assignment.AssignedBy,
		ExpiresAt:  assignment.ExpiresAt,
		IsPrimary:  assignment.IsPrimary,
	}
	stmt := ur.
		INSERT(ur.UserID, ur.RoleID, ur.AssignedAt, ur.AssignedBy, ur.ExpiresAt, ur.IsPrimary).
		MODEL(row)

	if _, err := stmt.ExecContext(ctx, u.DB(ctx)); err != nil {
		if db.IsUniqueViolation(err) {
			return access.ErrAssignmentAlreadyExists
		}
		return err
	}

	// The primary role is unique per user, so demote the incumbent inside the
	// same transaction the caller opened.
	if assignment.IsPrimary {
		return u.clearPrimaryExcept(ctx, assignment.UserID, assignment.RoleID)
	}
	return nil
}

// Unassign implements [access.UserRoleRepository].
func (u *UserRoleRepository) Unassign(
	ctx context.Context,
	userID uuid.UUID,
	roleID access.RoleID,
) error {
	res, err := table.UserRole.DELETE().WHERE(u.assignmentPair(userID, roleID)).
		ExecContext(ctx, u.DB(ctx))
	if err != nil {
		return err
	}
	if n, err := res.RowsAffected(); err == nil && n == 0 {
		return access.ErrAssignmentNotFound
	}
	return nil
}

// Get implements [access.UserRoleRepository].
func (u *UserRoleRepository) Get(
	ctx context.Context,
	userID uuid.UUID,
	roleID access.RoleID,
) (*access.UserRoleAssignment, error) {
	ur := table.UserRole
	var row models.UserRole
	if err := ur.SELECT(ur.AllColumns).WHERE(u.assignmentPair(userID, roleID)).
		LIMIT(1).QueryContext(ctx, u.DB(ctx), &row); err != nil {
		if db.IsNoRowError(err) {
			return nil, access.ErrAssignmentNotFound
		}
		return nil, err
	}
	return toAssignment(row), nil
}

// ListByUser implements [access.UserRoleRepository].
func (u *UserRoleRepository) ListByUser(
	ctx context.Context,
	userID uuid.UUID,
	limit, offset int,
) ([]access.UserRole, int, error) {
	return u.list(ctx, table.UserRole.UserID.EQ(pg.UUID(userID)), limit, offset)
}

// ListByRole implements [access.UserRoleRepository].
func (u *UserRoleRepository) ListByRole(
	ctx context.Context,
	roleID access.RoleID,
	limit, offset int,
) ([]access.UserRole, int, error) {
	return u.list(ctx, table.UserRole.RoleID.EQ(pg.Int(int64(roleID))), limit, offset)
}

func (u *UserRoleRepository) list(
	ctx context.Context,
	filter pg.BoolExpression,
	limit, offset int,
) ([]access.UserRole, int, error) {
	ur := table.UserRole
	assignments := []access.UserRole{}

	stmt := ur.SELECT(ur.AllColumns).WHERE(filter)
	rows, count, err := u.Paginate[[]models.UserRole](ctx, stmt, limit, offset)
	if err != nil {
		return assignments, 0, err
	}
	for _, v := range rows {
		assignments = append(assignments, toUserRole(v))
	}
	return assignments, count, nil
}

// ListActiveByUser implements [access.UserRoleRepository]. It returns the
// assignments that are still in force, which is what the policy engine needs,
// so it is deliberately not paginated.
func (u *UserRoleRepository) ListActiveByUser(
	ctx context.Context,
	userID uuid.UUID,
) ([]access.UserRole, error) {
	ur := table.UserRole
	assignments := []access.UserRole{}
	var rows []models.UserRole

	stmt := ur.SELECT(ur.AllColumns).WHERE(pg.AND(
		ur.UserID.EQ(pg.UUID(userID)),
		pg.OR(ur.ExpiresAt.IS_NULL(), ur.ExpiresAt.GT(pg.LOCALTIMESTAMP())),
	))
	if err := u.QueryMany(ctx, stmt, &rows); err != nil {
		return assignments, err
	}
	for _, v := range rows {
		assignments = append(assignments, toUserRole(v))
	}
	return assignments, nil
}

// UpdateExpiration implements [access.UserRoleRepository].
func (u *UserRoleRepository) UpdateExpiration(
	ctx context.Context,
	userID uuid.UUID,
	roleID access.RoleID,
	expiresAt *time.Time,
) error {
	res, err := table.UserRole.UPDATE(table.UserRole.ExpiresAt).
		MODEL(models.UserRole{ExpiresAt: expiresAt}).
		WHERE(u.assignmentPair(userID, roleID)).
		ExecContext(ctx, u.DB(ctx))
	if err != nil {
		return err
	}
	if n, err := res.RowsAffected(); err == nil && n == 0 {
		return access.ErrAssignmentNotFound
	}
	return nil
}

// SetPrimary implements [access.UserRoleRepository].
func (u *UserRoleRepository) SetPrimary(
	ctx context.Context,
	userID uuid.UUID,
	roleID access.RoleID,
) error {
	if _, err := u.Get(ctx, userID, roleID); err != nil {
		return err
	}
	res, err := table.UserRole.UPDATE(table.UserRole.IsPrimary).
		MODEL(models.UserRole{IsPrimary: true}).
		WHERE(u.assignmentPair(userID, roleID)).
		ExecContext(ctx, u.DB(ctx))
	if err != nil {
		return err
	}
	if n, err := res.RowsAffected(); err == nil && n == 0 {
		return access.ErrAssignmentNotFound
	}
	return u.clearPrimaryExcept(ctx, userID, roleID)
}

// ExistsActive implements [access.UserRoleRepository].
func (u *UserRoleRepository) ExistsActive(
	ctx context.Context,
	userID uuid.UUID,
	roleID access.RoleID,
) (bool, error) {
	ur := table.UserRole
	var count int
	if err := u.QueryOne(ctx, ur.SELECT(pg.COUNT(pg.Int(1))).WHERE(pg.AND(
		ur.UserID.EQ(pg.UUID(userID)),
		ur.RoleID.EQ(pg.Int(int64(roleID))),
		pg.OR(ur.ExpiresAt.IS_NULL(), ur.ExpiresAt.GT(pg.LOCALTIMESTAMP())),
	)), &count); err != nil {
		return false, err
	}
	return count > 0, nil
}

// clearPrimaryExcept demotes every other primary role of the user.
func (u *UserRoleRepository) clearPrimaryExcept(
	ctx context.Context,
	userID uuid.UUID,
	keepRoleID access.RoleID,
) error {
	ur := table.UserRole
	_, err := ur.UPDATE(ur.IsPrimary).
		MODEL(models.UserRole{IsPrimary: false}).
		WHERE(pg.AND(
			ur.UserID.EQ(pg.UUID(userID)),
			ur.RoleID.NOT_EQ(pg.Int(int64(keepRoleID))),
			ur.IsPrimary.EQ(pg.Bool(true)),
		)).
		ExecContext(ctx, u.DB(ctx))
	return err
}

func (u *UserRoleRepository) assignmentPair(
	userID uuid.UUID,
	roleID access.RoleID,
) pg.BoolExpression {
	return pg.AND(
		table.UserRole.UserID.EQ(pg.UUID(userID)),
		table.UserRole.RoleID.EQ(pg.Int(int64(roleID))),
	)
}

func toAssignment(row models.UserRole) *access.UserRoleAssignment {
	return &access.UserRoleAssignment{
		UserID:     row.UserID,
		RoleID:     row.RoleID,
		AssignedAt: row.AssignedAt,
		AssignedBy: row.AssignedBy,
		ExpiresAt:  row.ExpiresAt,
		IsPrimary:  row.IsPrimary,
	}
}

func toUserRole(row models.UserRole) access.UserRole {
	return access.UserRole{
		UserID:     row.UserID,
		RoleID:     row.RoleID,
		AssignedBy: row.AssignedBy,
		AssignedAt: row.AssignedAt,
		ExpiresAt:  row.ExpiresAt,
		IsPrimary:  row.IsPrimary,
	}
}
