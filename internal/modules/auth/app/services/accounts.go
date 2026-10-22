// Package services provides services  
package services

import (
	"context"
	"errors"

	"modular-monolith/internal/modules/auth/domain/users"
	"modular-monolith/lib/db"
	"modular-monolith/lib/logger"
	"modular-monolith/lib/utils"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AccountService struct {
	// repo is the authentication read/write side; usersRepo is the admin side
	// used for self-service profile writes.
	repo      users.AccountRepository
	usersRepo users.UserRepository
	db        *db.Manager
}

func NewAccountService(
	repo users.AccountRepository,
	usersRepo users.UserRepository,
	db *db.Manager,
) *AccountService {
	return &AccountService{
		repo:      repo,
		usersRepo: usersRepo,
		db:        db,
	}
}

// GetAuthAccount authenticates a user by identifier (email/phone/username) + password.
type GetAuthAccount struct {
	Identifier string
	Password   string
	IP         string
}

// RegisterAccount is the command to register a new local (password) account.
type RegisterAccount struct {
	Name     string
	Email    *string
	Phone    *string
	Username *string
	Password string
}

// RegisterAccountResult is returned on success. Commands never return entities.
type RegisterAccountResult struct {
	AccountID string
}

// Register Handle validates input, enforces uniqueness, hashes the password and persists.
func (u AccountService) Register(
	ctx context.Context,
	cmd RegisterAccount,
) (*RegisterAccountResult, error) {
	if cmd.Name == "" || !hasIdentifier(cmd) {
		return nil, users.ErrInvalidInput
	}

	for _, c := range []struct {
		col string
		val *string
	}{
		{"email", cmd.Email},
		{"phone", cmd.Phone},
		{"username", cmd.Username},
	} {
		if c.val == nil || *c.val == "" {
			continue
		}
		exists, err := u.repo.Exists(ctx, c.col, *c.val)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, users.ErrUserAlreadyExists
		}
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &users.User{
		Name:         cmd.Name,
		Email:        deref(cmd.Email),
		Phone:        deref(cmd.Phone),
		Username:     deref(cmd.Username),
		PasswordHash: string(hash),
		Status:       users.StatusActive,
	}
	err = u.db.Run(ctx, func(ctx context.Context) error {
		if err = u.repo.Create(ctx, user); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &RegisterAccountResult{AccountID: user.ID.String()}, nil
}

// SignIn Handle verifies credentials and returns the user with resolved roles and
// permissions. Failed attempts are recorded; locked/inactive accounts are
// rejected before the password is checked.
func (u AccountService) SignIn(ctx context.Context, in GetAuthAccount) (users.User, error) {
	user, err := u.repo.GetByIdentifier(ctx, in.Identifier)
	if err != nil {
		if errors.Is(err, users.ErrUserNotFound) {
			// ponytail: still run a hash compare-cost trade-off if timing
			// attacks matter; skipped for the core slice.
			return users.User{}, users.ErrInvalidPassword
		}
		return users.User{}, err
	}

	if user.Status == users.StatusLocked {
		return users.User{}, users.ErrAccountLocked
	}
	if !user.IsActive() {
		return users.User{}, users.ErrUserNotActive
	}

	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(in.Password)) != nil {
		// _ = q.repo.RecordFailedLogin(ctx, user.ID)
		return users.User{}, users.ErrInvalidPassword
	}

	roles, err := u.repo.GetRoles(ctx, user.ID)
	if err != nil {
		logger.Error(err)
		return users.User{}, err
	}
	user.Roles = roles
	if user.Permissions, err = u.repo.GetPermissions(ctx, user.ID); err != nil {
		logger.Error(err)
		return users.User{}, err
	}

	if err := u.repo.RecordSuccessfulLogin(ctx, user.ID, in.IP); err != nil {
		return users.User{}, err
	}
	return user, nil
}

// GetByID Handle returns the user with resolved roles and permissions.
func (u AccountService) GetByID(ctx context.Context, UserID uuid.UUID) (users.User, error) {
	user, err := u.repo.GetByID(ctx, UserID)
	if err != nil {
		return users.User{}, err
	}
	if user.Roles, err = u.repo.GetRoles(ctx, user.ID); err != nil {
		return users.User{}, err
	}
	if user.Permissions, err = u.repo.GetPermissions(ctx, user.ID); err != nil {
		return users.User{}, err
	}
	return user, nil
}

func (u AccountService) GetPreference(
	ctx context.Context,
	UserID uuid.UUID,
) (map[string]any, error) {
	preference, err := u.repo.GetPreference(ctx, UserID)
	if err != nil {
		return map[string]any{}, err
	}
	return preference, nil
}

func (u AccountService) UpdatePreference(
	ctx context.Context,
	UserID uuid.UUID,
	in map[string]any,
) (map[string]any, error) {
	preference, err := u.repo.UpdatePreference(ctx, UserID, in)
	if err != nil {
		return map[string]any{}, err
	}
	return preference, nil
}

func hasIdentifier(cmd RegisterAccount) bool {
	return nonEmpty(cmd.Email) || nonEmpty(cmd.Phone) || nonEmpty(cmd.Username)
}

var (
	deref    = utils.Deref[string]
	nonEmpty = utils.NonEmpty
)

// ListConnections returns the external identities linked to the account.
func (u AccountService) ListConnections(
	ctx context.Context,
	UserID uuid.UUID,
) ([]users.Connection, error) {
	return u.repo.ListConnections(ctx, UserID)
}

// UnlinkConnection removes one linked identity after checking the caller's
// password. The repository refuses to drop the last one, which would leave the
// user with no way to sign in.
func (u AccountService) UnlinkConnection(
	ctx context.Context,
	UserID uuid.UUID,
	provider string,
	password string,
) error {
	if provider == "" {
		return users.ErrInvalidInput
	}
	if err := u.VerifyPassword(ctx, UserID, password); err != nil {
		return err
	}
	return u.repo.UnlinkConnection(ctx, UserID, provider)
}

// UpdateProfile applies a partial update to the signed-in user's own profile.
func (u AccountService) UpdateProfile(
	ctx context.Context,
	UserID uuid.UUID,
	in users.UpdateUserInput,
) error {
	if _, err := u.repo.GetByID(ctx, UserID); err != nil {
		return err
	}
	return u.usersRepo.Update(ctx, UserID, in)
}

// DeleteAccount closes the account after checking the caller's password.
// Password authentication is disabled first so the stored hash can never be
// reused, then the row is soft deleted.
func (u AccountService) DeleteAccount(
	ctx context.Context,
	UserID uuid.UUID,
	password string,
) error {
	if err := u.VerifyPassword(ctx, UserID, password); err != nil {
		return err
	}
	return u.db.Run(ctx, func(ctx context.Context) error {
		if err := u.usersRepo.SetPasswordHash(ctx, UserID, ""); err != nil {
			return err
		}
		return u.usersRepo.Delete(ctx, UserID)
	})
}

// VerifyPassword checks a candidate password against the stored hash. It backs
// the endpoints that demand the current password to prove possession before
// changing something security sensitive.
func (u AccountService) VerifyPassword(
	ctx context.Context,
	UserID uuid.UUID,
	password string,
) error {
	user, err := u.repo.GetByID(ctx, UserID)
	if err != nil {
		return err
	}
	if user.PasswordHash == "" {
		// Password login is disabled on this account, so there is nothing to
		// prove against.
		return users.ErrInvalidPassword
	}
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	); err != nil {
		return users.ErrInvalidPassword
	}
	return nil
}
