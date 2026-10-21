// Package users implements the auth domain Repository against PostgreSQL
// using jet query builders. All access is scoped to the "auth" schema.
package users

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"modular-monolith/internal/auth/domain/users"
	model "modular-monolith/internal/auth/infra/.gen/models"
	table "modular-monolith/internal/auth/infra/.gen/table"
	"modular-monolith/internal/auth/infra/mappers"
	"modular-monolith/lib/db"
	"modular-monolith/lib/utils"

	pg "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// maxFailedLogins locks an account once reached.
// ponytail: fixed threshold; make it config-driven if product needs per-tenant tuning.
const maxFailedLogins = 5

var (
	userT     = table.User
	roleT     = table.Role
	userRoleT = table.UserRole
	rolePermT = table.RolePermission
	permT     = table.Permission
)

var _ users.AccountRepository = (*AccountRepository)(nil)

var nullable = utils.Nullable[string]

// AccountRepository is the sqlx/jet-backed implementation of users.AccountRepository.
type AccountRepository struct {
	db.BaseRepository
}

// NewAccountRepository builds the auth repository.
func NewAccountRepository(DB *sqlx.DB) *AccountRepository {
	return &AccountRepository{BaseRepository: db.NewRepo(DB)}
}

// Create inserts a new user row. Local (password) signups do not create an
// account row; that table is reserved for OAuth providers.
func (r *AccountRepository) Create(ctx context.Context, in *users.User) error {
	row := model.User{
		Name:          in.Name,
		Email:         nullable(in.Email),
		Phone:         nullable(in.Phone),
		Username:      nullable(in.Username),
		Image:         nullable(in.Image),
		PasswordHash:  nullable(in.PasswordHash),
		EmailVerified: in.EmailVerified,
		Status:        users.StatusActive,
	}

	var created model.User
	stmt := userT.
		INSERT(
			userT.Name, userT.Email, userT.Phone, userT.Username,
			userT.Image, userT.PasswordHash, userT.EmailVerified, userT.Status,
		).
		MODEL(row).
		RETURNING(userT.AllColumns)

	if err := stmt.QueryContext(ctx, r.DB(ctx), &created); err != nil {
		return err
	}

	// The role is looked up by slug, not hardcoded id: the seed inserts
	// super_admin first, so an id assumption here silently made every signup an
	// administrator.
	roleID, err := r.roleIDBySlug(ctx, DefaultRoleSlug)
	if err != nil {
		return err
	}
	userRoleRow := model.UserRole{
		RoleID:     roleID,
		UserID:     created.ID,
		AssignedAt: time.Now(),
		IsPrimary:  true,
	}
	userRoleStmt := userRoleT.INSERT(userRoleT.RoleID, userRoleT.UserID, userRoleT.AssignedAt, userRoleT.IsPrimary).
		MODEL(userRoleRow)
	if _, err := userRoleStmt.ExecContext(ctx, r.DB(ctx)); err != nil {
		return err
	}

	in.ID = created.ID
	in.CreatedAt = created.CreatedAt
	in.UpdatedAt = created.UpdatedAt

	return nil
}

// roleIDBySlug resolves a seeded role. A fresh database without the seed
// migration has no roles, so an unknown slug yields 0 rather than an error and
// the caller can decide whether that is fatal.
func (r *AccountRepository) roleIDBySlug(ctx context.Context, slug string) (int32, error) {
	var id int32
	stmt := roleT.SELECT(roleT.ID).
		WHERE(roleT.Slug.EQ(pg.String(slug))).
		LIMIT(1)
	if err := stmt.QueryContext(ctx, r.DB(ctx), &id); err != nil {
		if db.IsNoRowError(err) {
			return 0, nil
		}
		return 0, err
	}
	return id, nil
}

// GetByIdentifier matches email, phone, or username.
func (r *AccountRepository) GetByIdentifier(
	ctx context.Context,
	identifier string,
) (users.User, error) {
	val := pg.String(identifier)
	stmt := pg.SELECT(
		userT.AllColumns,
	).
		FROM(userT).
		WHERE(
			userT.DeletedAt.IS_NULL().AND(
				userT.Email.EQ(val).
					OR(userT.Phone.EQ(val)).
					OR(userT.Username.EQ(val)),
			),
		).
		LIMIT(1)

	return r.queryOne(ctx, stmt)
}

// GetByID matches the primary key.
func (r *AccountRepository) GetByID(ctx context.Context, id uuid.UUID) (users.User, error) {
	stmt := pg.SELECT(userT.AllColumns).
		FROM(userT).
		WHERE(userT.DeletedAt.IS_NULL().AND(userT.ID.EQ(pg.UUID(id)))).
		LIMIT(1)

	return r.queryOne(ctx, stmt)
}

func (r *AccountRepository) queryOne(
	ctx context.Context,
	stmt pg.SelectStatement,
) (users.User, error) {
	var row model.User
	if err := stmt.QueryContext(ctx, r.DB(ctx), &row); err != nil {
		if errors.Is(err, sql.ErrNoRows) || errors.Is(err, qrm.ErrNoRows) {
			return users.User{}, users.ErrUserNotFound
		}
		return users.User{}, err
	}
	return mappers.ToDomainUser(row), nil
}

// Exists reports whether a non-deleted user already uses value in the column.
func (r *AccountRepository) Exists(ctx context.Context, column, value string) (bool, error) {
	var col pg.ColumnString
	switch column {
	case "email":
		col = userT.Email
	case "phone":
		col = userT.Phone
	case "username":
		col = userT.Username
	default:
		return false, users.ErrInvalidInput
	}

	stmt := pg.SELECT(userT.ID).
		FROM(userT).
		WHERE(userT.DeletedAt.IS_NULL().AND(col.EQ(pg.String(value)))).
		LIMIT(1)

	var row model.User
	err := stmt.QueryContext(ctx, r.DB(ctx), &row)
	if errors.Is(err, sql.ErrNoRows) || errors.Is(err, qrm.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// RecordSuccessfulLogin clears failures and stamps the login.
func (r *AccountRepository) RecordSuccessfulLogin(
	ctx context.Context,
	id uuid.UUID,
	ip string,
) error {
	now := time.Now()
	stmt := userT.UPDATE(
		userT.FailedLoginAttempts, userT.LockedUntil,
		userT.LastLoginAt, userT.LastLoginIP, userT.UpdatedAt,
	).SET(
		pg.Int(0), pg.TimestampExp(pg.NULL),
		pg.TimestampT(now), pg.String(ip), pg.TimestampT(now),
	).WHERE(userT.ID.EQ(pg.UUID(id)))

	_, err := stmt.ExecContext(ctx, r.DB(ctx))
	return err
}

// RecordFailedLogin increments the counter and locks on threshold.
func (r *AccountRepository) RecordFailedLogin(ctx context.Context, id uuid.UUID) error {
	attempts := userT.FailedLoginAttempts.ADD(pg.Int(1))
	stmt := userT.UPDATE(
		userT.FailedLoginAttempts, userT.Status, userT.UpdatedAt,
	).SET(
		attempts,
		pg.CASE().
			WHEN(attempts.GT_EQ(pg.Int(maxFailedLogins))).
			THEN(pg.String(users.StatusLocked)).
			ELSE(userT.Status),
		pg.TimestampT(time.Now()),
	).WHERE(userT.ID.EQ(pg.String(id.String())))

	_, err := stmt.ExecContext(ctx, r.DB(ctx))
	return err
}

// GetRoles returns assigned, non-expired role slugs.
func (r *AccountRepository) GetRoles(ctx context.Context, id uuid.UUID) ([]string, error) {
	stmt := pg.SELECT(roleT.Slug).
		FROM(userRoleT.INNER_JOIN(roleT, roleT.ID.EQ(userRoleT.RoleID))).
		WHERE(
			userRoleT.UserID.EQ(pg.UUID(id)).
				AND(roleT.DeletedAt.IS_NULL()).
				AND(userRoleT.ExpiresAt.IS_NULL().OR(userRoleT.ExpiresAt.GT(pg.TimestampT(time.Now())))),
		)

	var rows []string
	if err := stmt.QueryContext(ctx, r.DB(ctx), &rows); err != nil {
		if errors.Is(err, qrm.ErrNoRows) || errors.Is(err, sql.ErrNoRows) {
			return []string{}, nil
		}
		return nil, err
	}
	return rows, nil
}

// GetPermissions returns distinct permission slugs granted via the user's roles.
func (r *AccountRepository) GetPermissions(ctx context.Context, id uuid.UUID) ([]string, error) {
	stmt := pg.SELECT(permT.Slug).DISTINCT().
		FROM(
			userRoleT.
				INNER_JOIN(rolePermT, rolePermT.RoleID.EQ(userRoleT.RoleID)).
				INNER_JOIN(permT, permT.ID.EQ(rolePermT.PermissionID)),
		).
		WHERE(
			userRoleT.UserID.EQ(pg.UUID(id)).
				AND(permT.DeletedAt.IS_NULL()).
				AND(rolePermT.ExpiresAt.IS_NULL().OR(rolePermT.ExpiresAt.GT(pg.TimestampT(time.Now())))),
		)

	var rows []string
	if err := stmt.QueryContext(ctx, r.DB(ctx), &rows); err != nil {
		if errors.Is(err, qrm.ErrNoRows) || errors.Is(err, sql.ErrNoRows) {
			return []string{}, nil
		}
		return nil, err
	}
	return rows, nil
}

// GetPreference implements [users.AccountRepository].
func (r *AccountRepository) GetPreference(
	ctx context.Context,
	id uuid.UUID,
) (map[string]any, error) {
	var preference model.User
	stmt := pg.SELECT(userT.Preferences).FROM(userT).WHERE(userT.ID.EQ(pg.UUID(id)))

	err := stmt.QueryContext(ctx, r.DB(ctx), &preference)
	if err != nil {
		if db.IsNoRowError(err) {
			return map[string]any{}, nil
		}
		return nil, fmt.Errorf("GetPreference: %w", err)
	}
	return *preference.Preferences, nil
}

// UpdatePreference implements [users.AccountRepository].
func (r *AccountRepository) UpdatePreference(
	ctx context.Context,
	id uuid.UUID,
	in map[string]any,
) (map[string]any, error) {
	var updated model.User
	stmt := userT.UPDATE(userT.Preferences, userT.UpdatedAt).
		SET(in, time.Now()).
		WHERE(userT.ID.EQ(pg.UUID(id))).
		RETURNING(userT.AllColumns)
	if err := stmt.QueryContext(ctx, r.DB(ctx), &updated); err != nil {
		if db.IsNoRowError(err) {
			return nil, fmt.Errorf("UpdatePreference: %w", users.ErrUserNotFound)
		}
		return nil, fmt.Errorf("UpdatePreference: %w", err)
	}
	if updated.Preferences == nil {
		return map[string]any{}, nil
	}
	return *updated.Preferences, nil
}

// ListConnections implements [users.AccountRepository]. Tokens are deliberately
// not selected: nothing outside the OAuth flow needs them.
func (r *AccountRepository) ListConnections(
	ctx context.Context,
	id uuid.UUID,
) ([]users.Connection, error) {
	acct := table.Account
	connections := []users.Connection{}

	var rows []model.Account
	stmt := acct.SELECT(
		acct.ID, acct.CreatedAt, acct.ProviderID,
		acct.ProviderAccountID, acct.Scope, acct.AccessTokenExpiresAt,
	).WHERE(pg.AND(
		acct.UserID.EQ(pg.UUID(id)),
		acct.DeletedAt.IS_NULL(),
	)).ORDER_BY(acct.CreatedAt.ASC())

	if err := r.QueryMany(ctx, stmt, &rows); err != nil {
		return connections, err
	}
	for _, row := range rows {
		connections = append(connections, users.Connection{
			ID:                 int64(row.ID),
			Provider:           row.ProviderID,
			ProviderAccountID:  row.ProviderAccountID,
			Scopes:             splitScopes(utils.Deref(row.Scope)),
			AccessTokenExpires: row.AccessTokenExpiresAt,
			CreatedAt:          row.CreatedAt,
		})
	}
	return connections, nil
}

// UnlinkConnection implements [users.AccountRepository]. Unlinking the only
// remaining sign-in method would lock the user out, so it is rejected.
func (r *AccountRepository) UnlinkConnection(
	ctx context.Context,
	id uuid.UUID,
	provider string,
) error {
	acct := table.Account

	var remaining int
	if err := r.QueryOne(ctx, acct.SELECT(pg.COUNT(pg.Int(1))).WHERE(pg.AND(
		acct.UserID.EQ(pg.UUID(id)),
		acct.DeletedAt.IS_NULL(),
	)), &remaining); err != nil {
		return err
	}
	if remaining <= 1 {
		return users.ErrConnectionProtected
	}

	res, err := acct.DELETE().WHERE(pg.AND(
		acct.UserID.EQ(pg.UUID(id)),
		acct.ProviderID.EQ(pg.String(provider)),
	)).ExecContext(ctx, r.DB(ctx))
	if err != nil {
		return err
	}
	if n, err := res.RowsAffected(); err == nil && n == 0 {
		return users.ErrConnectionNotFound
	}
	return nil
}

// splitScopes reads the space separated scope column. Some providers write a
// JSON array instead, so both shapes are accepted.
func splitScopes(raw string) []string {
	if raw == "" {
		return []string{}
	}
	if strings.HasPrefix(raw, "[") {
		var out []string
		if err := json.Unmarshal([]byte(raw), &out); err == nil {
			return out
		}
	}
	return strings.Fields(raw)
}

//
// func slugs(rows []struct{ Slug string }) []string {
// 	out := make([]string, len(rows))
// 	for i, r := range rows {
// 		out[i] = r.Slug
// 	}
// 	return out
// }
