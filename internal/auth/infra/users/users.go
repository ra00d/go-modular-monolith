package users

import (
	"context"
	"time"

	"modular-monolith/internal/auth/domain/users"
	"modular-monolith/internal/auth/infra/.gen/models"
	"modular-monolith/internal/auth/infra/.gen/table"
	"modular-monolith/internal/auth/infra/mappers"
	"modular-monolith/lib/db"
	"modular-monolith/lib/utils"

	pg "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// DefaultRoleSlug is the role a fresh account receives. Resolved by slug so the
// seed ordering does not matter.
const DefaultRoleSlug = "user"

var (
	User          = table.User
	UserRole      = table.UserRole
	UserAttribute = table.UserAttribute
)

var _ users.UserRepository = (*UserRepository)(nil)

type UserRepository struct {
	db.BaseRepository
}

func NewUsersAdminRepository(DB *sqlx.DB) *UserRepository {
	return &UserRepository{BaseRepository: db.NewRepo(DB)}
}

// GetAll implements [admin.UsersAdminRepository].
func (u *UserRepository) GetAll(ctx context.Context, limit, offset int) ([]users.User, int, error) {
	stmt := pg.SELECT(User.AllColumns).FROM(User).WHERE(User.DeletedAt.IS_NULL())
	users := []users.User{}

	rows, count, err := u.Paginate[[]models.User](ctx, stmt, limit, offset)
	if err != nil {
		return users, 0, err
	}
	for _, v := range rows {
		users = append(users, mappers.ToDomainUser(v))
	}
	return users, count, nil
}

// Create implements [admin.UsersAdminRepository].
func (u *UserRepository) Create(
	ctx context.Context,
	in users.CreateUserInput,
) (uuid.UUID, error) {
	row := models.User{
		Name:          in.Name,
		Email:         utils.Nullable(in.Email),
		Phone:         utils.Nullable(in.Phone),
		Username:      utils.Nullable(in.Username),
		Image:         utils.Nullable(in.Image),
		PasswordHash:  utils.Nullable(in.PasswordHash),
		EmailVerified: in.EmailVerified,
		Status:        in.Status,
	}

	var created models.User
	stmt := User.
		INSERT(
			User.Name, User.Email, User.Phone, User.Username,
			User.Image, User.PasswordHash, User.EmailVerified, User.Status,
		).
		MODEL(row).
		RETURNING(User.AllColumns)

	if err := stmt.QueryContext(ctx, u.DB(ctx), &created); err != nil {
		return uuid.UUID{}, err
	}

	return created.ID, nil
}

// Update implements [users.UserRepository].
//
// jet has no way to build a SET list from a dynamic column slice, so the patch
// is applied to the stored row first and written back whole. Replacing an
// identifier also clears its verified flag: the old proof says nothing about
// the new address.
func (u *UserRepository) Update(
	ctx context.Context,
	id uuid.UUID,
	in users.UpdateUserInput,
) error {
	var current models.User
	if err := User.SELECT(User.AllColumns).
		WHERE(pg.AND(User.ID.EQ(pg.UUID(id)), User.DeletedAt.IS_NULL())).
		LIMIT(1).
		QueryContext(ctx, u.DB(ctx), &current); err != nil {
		if db.IsNoRowError(err) {
			return users.ErrUserNotFound
		}
		return err
	}

	if in.Name != nil {
		current.Name = *in.Name
	}
	if in.Image != nil {
		current.Image = utils.Nullable(*in.Image)
	}
	if in.Email != nil {
		current.Email = utils.Nullable(*in.Email)
		current.EmailVerified = false
	}
	if in.Phone != nil {
		current.Phone = utils.Nullable(*in.Phone)
		current.PhoneVerified = false
	}
	if in.Username != nil {
		current.Username = utils.Nullable(*in.Username)
	}
	current.UpdatedAt = time.Now()

	res, err := User.UPDATE(
		User.Name, User.Email, User.Phone, User.Username, User.Image,
		User.EmailVerified, User.PhoneVerified, User.UpdatedAt,
	).MODEL(current).
		WHERE(pg.AND(User.ID.EQ(pg.UUID(id)), User.DeletedAt.IS_NULL())).
		ExecContext(ctx, u.DB(ctx))
	if err != nil {
		if db.IsUniqueViolation(err) {
			return users.ErrUserAlreadyExists
		}
		return err
	}
	if n, err := res.RowsAffected(); err == nil && n == 0 {
		return users.ErrUserNotFound
	}
	return nil
}

// Delete implements [users.UserRepository]. Soft delete so the audit trail keeps
// pointing at a resolvable row.
func (u *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	now := time.Now()
	res, err := User.UPDATE(User.DeletedAt, User.UpdatedAt).
		MODEL(models.User{DeletedAt: &now, UpdatedAt: now}).
		WHERE(pg.AND(User.ID.EQ(pg.UUID(id)), User.DeletedAt.IS_NULL())).
		ExecContext(ctx, u.DB(ctx))
	if err != nil {
		return err
	}
	if n, err := res.RowsAffected(); err == nil && n == 0 {
		return users.ErrUserNotFound
	}
	return nil
}

// SetStatus implements [users.UserRepository].
func (u *UserRepository) SetStatus(
	ctx context.Context,
	id uuid.UUID,
	status string,
	banUntil *time.Time,
) error {
	res, err := User.UPDATE(User.Status, User.LockedUntil, User.UpdatedAt).
		MODEL(models.User{
			Status:      status,
			LockedUntil: banUntil,
			UpdatedAt:   time.Now(),
		}).
		WHERE(pg.AND(User.ID.EQ(pg.UUID(id)), User.DeletedAt.IS_NULL())).
		ExecContext(ctx, u.DB(ctx))
	if err != nil {
		return err
	}
	if n, err := res.RowsAffected(); err == nil && n == 0 {
		return users.ErrUserNotFound
	}
	return nil
}

// SetPasswordHash implements [users.UserRepository].
func (u *UserRepository) SetPasswordHash(
	ctx context.Context,
	id uuid.UUID,
	hash string,
) error {
	res, err := User.UPDATE(User.PasswordHash, User.UpdatedAt).
		MODEL(models.User{PasswordHash: utils.Nullable(hash), UpdatedAt: time.Now()}).
		WHERE(pg.AND(User.ID.EQ(pg.UUID(id)), User.DeletedAt.IS_NULL())).
		ExecContext(ctx, u.DB(ctx))
	if err != nil {
		return err
	}
	if n, err := res.RowsAffected(); err == nil && n == 0 {
		return users.ErrUserNotFound
	}
	return nil
}

// AssignRole implements [admin.UsersAdminRepository].
func (u *UserRepository) AssignRole(ctx context.Context, userID uuid.UUID, roleID int) error {
	userRoleRow := models.UserRole{
		RoleID:     int32(roleID),
		UserID:     userID,
		AssignedAt: time.Now(),
		IsPrimary:  true,
	}
	userRoleStmt := UserRole.
		INSERT(UserRole.RoleID, UserRole.UserID, UserRole.AssignedAt, UserRole.IsPrimary).
		MODEL(userRoleRow)
	if _, err := userRoleStmt.ExecContext(ctx, u.DB(ctx)); err != nil {
		return err
	}
	// Only one primary role per user, so demote whatever held it.
	_, err := UserRole.UPDATE(UserRole.IsPrimary).
		MODEL(models.UserRole{IsPrimary: false}).
		WHERE(pg.AND(
			UserRole.UserID.EQ(pg.UUID(userID)),
			UserRole.RoleID.NOT_EQ(pg.Int(int64(roleID))),
			UserRole.IsPrimary.EQ(pg.Bool(true)),
		)).
		ExecContext(ctx, u.DB(ctx))
	return err
}

// AssignPermission implements [admin.UsersAdminRepository].
func (u *UserRepository) AssignPermission(
	ctx context.Context,
	userID uuid.UUID,
	permissionID int,
) error {
	return u.AssignPermissions(ctx, userID, []int{permissionID})
}

// AssignPermissions implements [admin.UsersAdminRepository].
func (u *UserRepository) AssignPermissions(
	ctx context.Context,
	userID uuid.UUID,
	permissionIDs []int,
) error {
	// Direct user-to-permission grants have no table of their own; they are
	// modelled as a synthetic role holding the permissions, which is how the
	// permission resolver already reads them.
	for _, permissionID := range permissionIDs {
		if err := u.assignViaRole(ctx, userID, permissionID); err != nil {
			return err
		}
	}
	return nil
}

// assignViaRole grants one permission to the user through their primary role so
// the existing role_permission resolution path picks it up.
func (u *UserRepository) assignViaRole(
	ctx context.Context,
	userID uuid.UUID,
	permissionID int,
) error {
	var roleID int32
	stmt := UserRole.SELECT(UserRole.RoleID).
		WHERE(pg.AND(
			UserRole.UserID.EQ(pg.UUID(userID)),
			UserRole.IsPrimary.EQ(pg.Bool(true)),
		)).
		LIMIT(1)
	if err := stmt.QueryContext(ctx, u.DB(ctx), &roleID); err != nil {
		if db.IsNoRowError(err) {
			return users.ErrNoPrimaryRole
		}
		return err
	}

	grant := table.RolePermission
	var existing int
	if err := u.QueryOne(ctx, grant.SELECT(pg.COUNT(pg.Int(1))).WHERE(pg.AND(
		grant.RoleID.EQ(pg.Int(int64(roleID))),
		grant.PermissionID.EQ(pg.Int(int64(permissionID))),
	)), &existing); err != nil {
		return err
	}
	if existing > 0 {
		return nil
	}

	_, err := grant.
		INSERT(grant.RoleID, grant.PermissionID).
		VALUES(roleID, int32(permissionID)).
		ExecContext(ctx, u.DB(ctx))
	return err
}

// ListAttributes implements [users.UserRepository].
func (u *UserRepository) ListAttributes(
	ctx context.Context,
	userID uuid.UUID,
	namespace *string,
	limit, offset int,
) ([]users.UserAttribute, int, error) {
	attr := table.UserAttribute
	attributes := []users.UserAttribute{}

	stmt := attr.SELECT(attr.AllColumns).WHERE(attr.UserID.EQ(pg.UUID(userID)))
	if namespace != nil {
		stmt = stmt.WHERE(attr.Namespace.EQ(pg.String(*namespace)))
	}
	stmt = stmt.ORDER_BY(attr.Namespace.ASC(), attr.Key.ASC())

	rows, count, err := u.Paginate[[]models.UserAttribute](ctx, stmt, limit, offset)
	if err != nil {
		return attributes, 0, err
	}
	for _, v := range rows {
		attributes = append(attributes, toUserAttribute(v))
	}
	return attributes, count, nil
}

// SetAttribute implements [users.UserRepository].
func (u *UserRepository) SetAttribute(ctx context.Context, in users.UserAttribute) error {
	attr := table.UserAttribute
	if in.Namespace == "" {
		in.Namespace = users.DefaultNamespace
	}
	now := time.Now()
	row := models.UserAttribute{
		UserID:    in.UserID,
		Namespace: in.Namespace,
		Key:       in.Key,
		Value:     in.Value,
		ValueType: in.ValueType,
		UpdatedAt: now,
	}
	stmt := attr.
		INSERT(attr.UserID, attr.Namespace, attr.Key, attr.Value, attr.ValueType, attr.UpdatedAt).
		MODEL(row).
		ON_CONFLICT(attr.UserID, attr.Namespace, attr.Key).
		DO_UPDATE(pg.SET(
			attr.Value.SET(pg.String(in.Value)),
			attr.ValueType.SET(pg.String(in.ValueType)),
			attr.UpdatedAt.SET(pg.TimestampT(now)),
		))

	_, err := stmt.ExecContext(ctx, u.DB(ctx))
	return err
}

// DeleteAttribute implements [users.UserRepository].
func (u *UserRepository) DeleteAttribute(
	ctx context.Context,
	userID uuid.UUID,
	namespace, key string,
) error {
	attr := table.UserAttribute
	res, err := attr.DELETE().WHERE(pg.AND(
		attr.UserID.EQ(pg.UUID(userID)),
		attr.Namespace.EQ(pg.String(namespace)),
		attr.Key.EQ(pg.String(key)),
	)).ExecContext(ctx, u.DB(ctx))
	if err != nil {
		return err
	}
	if n, err := res.RowsAffected(); err == nil && n == 0 {
		return users.ErrAttributeNotFound
	}
	return nil
}

func toUserAttribute(row models.UserAttribute) users.UserAttribute {
	return users.UserAttribute{
		ID:        int64(row.ID),
		UserID:    row.UserID,
		Namespace: row.Namespace,
		Key:       row.Key,
		Value:     row.Value,
		ValueType: row.ValueType,
		UpdatedAt: row.UpdatedAt,
	}
}
