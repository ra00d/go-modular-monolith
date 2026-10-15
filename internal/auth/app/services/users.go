package services

import (
	"context"
	"time"

	"modular-monolith/internal/auth/domain/access"
	"modular-monolith/internal/auth/domain/account"
	"modular-monolith/internal/auth/domain/users"
	"modular-monolith/lib/auth"
	"modular-monolith/lib/db"
	"modular-monolith/lib/errors"
	"modular-monolith/lib/logger"
	"modular-monolith/lib/utils"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UsersService struct {
	usersRepo      users.UserRepository
	accountRepo    users.AccountRepository
	permissionRepo access.PermissionRepository
	userRoleRepo   access.UserRoleRepository
	db             *db.Manager
	auditor        auditRecorder
	jwt            *auth.JWTManager
}

func NewUsersService(
	db *db.Manager,
	usersRepo users.UserRepository,
	accountRepo users.AccountRepository,
	permissionRepo access.PermissionRepository,
	userRoleRepo access.UserRoleRepository,
	auditRepo access.AuditRepository,
	jwt *auth.JWTManager,
) *UsersService {
	return &UsersService{
		usersRepo,
		accountRepo,
		permissionRepo,
		userRoleRepo,
		db,
		newAuditRecorder(auditRepo),
		jwt,
	}
}

func (u UsersService) GetAll(ctx context.Context, limit, offset int) ([]users.User, int, error) {
	return u.usersRepo.GetAll(ctx, limit, offset)
}

func (u UsersService) GetByID(ctx context.Context, id uuid.UUID) (users.User, error) {
	user, err := u.accountRepo.GetByID(ctx, id)
	if err != nil {
		return users.User{}, err
	}
	if user.Roles, err = u.accountRepo.GetRoles(ctx, user.ID); err != nil {
		return users.User{}, err
	}
	if user.Permissions, err = u.accountRepo.GetPermissions(ctx, user.ID); err != nil {
		return users.User{}, err
	}
	return user, nil
}

type CreateUserInput struct {
	Name     string
	Email    *string
	Phone    *string
	Username *string
	Password string
	Role     int
}

func (u UsersService) Exists(ctx context.Context, in account.AccountIdentifier) (bool, error) {
	if in.Value() == "" || in.Column() == "" {
		return false, errors.Internal("invalid Identifier params", nil)
	}
	exists, err := u.accountRepo.Exists(ctx, in.Column(), in.Value())
	if err != nil {
		return false, err
	}
	if exists {
		return true, nil
	}
	return false, nil
}

func (u UsersService) Create(ctx context.Context, in CreateUserInput) (uuid.UUID, error) {
	idn := account.AccountIdentifier{
		Email:    in.Email,
		Phone:    in.Phone,
		Username: in.Username,
	}
	exist, err := u.Exists(ctx, idn)
	if exist {
		return uuid.UUID{}, errors.Invalid(
			"",
			"user with "+idn.Column()+" "+idn.Value()+" already exists",
		)
	}
	if err != nil {
		return uuid.UUID{}, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return uuid.UUID{}, err
	}
	var userID uuid.UUID
	err = u.db.Run(ctx, func(ctx context.Context) error {
		userID, err = u.usersRepo.Create(ctx, users.CreateUserInput{
			Name:          in.Name,
			Email:         utils.Deref(in.Email),
			Phone:         utils.Deref(in.Phone),
			Username:      utils.Deref(in.Username),
			PasswordHash:  string(hash),
			EmailVerified: false,
			PhoneVerified: false,
			Image:         "",
			Status:        users.StatusActive,
			Preference:    map[string]any{},
			MFAEnabled:    false,
			RoleID:        in.Role,
		})
		if err != nil {
			return errors.WrapRaw("Create user ", err)
		}

		err = u.usersRepo.AssignRole(ctx, userID, in.Role)
		if err != nil {
			return errors.WrapRaw("Asign user role", err)
		}

		return nil
	})
	if err != nil {
		return uuid.UUID{}, err
	}
	return userID, nil
}

// Update applies a partial update. Replacing an identifier drops its verified
// flag in the repository, so the new address has to be re-verified.
func (u UsersService) Update(
	ctx context.Context,
	id uuid.UUID,
	in users.UpdateUserInput,
	actor uuid.UUID,
) error {
	before, err := u.accountRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	err = u.db.Run(ctx, func(ctx context.Context) error {
		if err := u.usersRepo.Update(ctx, id, in); err != nil {
			return errors.WrapRaw("update user", err)
		}
		u.auditor.Record(ctx, actor, "user.update", "user", id.String(), &before)
		return nil
	})
	return err
}

func (u UsersService) Delete(
	ctx context.Context,
	id uuid.UUID,
	actor uuid.UUID,
	revokeSessions bool,
) error {
	before, err := u.accountRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	err = u.db.Run(ctx, func(ctx context.Context) error {
		if err := u.usersRepo.Delete(ctx, id); err != nil {
			return errors.WrapRaw("delete user", err)
		}
		u.auditor.Record(ctx, actor, "user.delete", "user", id.String(), &before)
		return nil
	})
	if err != nil {
		return err
	}
	if revokeSessions {
		return u.revokeSessions(id)
	}
	return nil
}

// SetStatus moves a user to another account status.
func (u UsersService) SetStatus(
	ctx context.Context,
	id uuid.UUID,
	status string,
	actor uuid.UUID,
) error {
	before, err := u.accountRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	err = u.db.Run(ctx, func(ctx context.Context) error {
		if err := u.usersRepo.SetStatus(ctx, id, status, nil); err != nil {
			return errors.WrapRaw("set user status", err)
		}
		u.auditor.Record(ctx, actor, "user.status", "user", id.String(), &before)
		return nil
	})
	if err != nil {
		return err
	}
	// A suspended or locked account must not keep an authenticated session.
	if status != users.StatusActive {
		return u.revokeSessions(id)
	}
	return nil
}

// Ban locks the account until until, or indefinitely when until is nil, and
// clears the failed login counter so the lock is the only thing in force.
func (u UsersService) Ban(
	ctx context.Context,
	id uuid.UUID,
	until *time.Time,
	reason string,
	actor uuid.UUID,
	revokeSessions bool,
) (string, error) {
	if _, err := u.accountRepo.GetByID(ctx, id); err != nil {
		return "", err
	}
	err := u.db.Run(ctx, func(ctx context.Context) error {
		if err := u.usersRepo.SetStatus(ctx, id, users.StatusLocked, until); err != nil {
			return errors.WrapRaw("ban user", err)
		}
		entry := &access.AuditLog{
			CreatedAt:    time.Now(),
			ActorType:    "user",
			ActorID:      actor.String(),
			Action:       "user.ban",
			Category:     access.CategorySecurity,
			ResourceType: "user",
			ResourceID:   id.String(),
			Status:       string(access.AuditStatusSuccess),
			ErrorMessage: reason,
		}
		u.auditor.RecordRaw(ctx, entry)
		return nil
	})
	if err != nil {
		return "", err
	}
	if revokeSessions {
		if err := u.revokeSessions(id); err != nil {
			return users.StatusLocked, err
		}
	}
	return users.StatusLocked, nil
}

// Unban lifts a lock and restores the active status.
func (u UsersService) Unban(
	ctx context.Context,
	id uuid.UUID,
	reason string,
	actor uuid.UUID,
) (string, error) {
	before, err := u.accountRepo.GetByID(ctx, id)
	if err != nil {
		return "", err
	}
	err = u.db.Run(ctx, func(ctx context.Context) error {
		if err := u.usersRepo.SetStatus(ctx, id, users.StatusActive, nil); err != nil {
			return errors.WrapRaw("unban user", err)
		}
		u.auditor.Record(ctx, actor, "user.unban", "user", id.String(), &before)
		return nil
	})
	if err != nil {
		return "", err
	}
	return users.StatusActive, nil
}

// ResetPassword sets a new password directly. This is the admin path, so no
// verification token is involved.
func (u UsersService) ResetPassword(
	ctx context.Context,
	id uuid.UUID,
	password string,
	actor uuid.UUID,
	revokeSessions bool,
) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	err = u.db.Run(ctx, func(ctx context.Context) error {
		if err := u.usersRepo.SetPasswordHash(ctx, id, string(hash)); err != nil {
			return errors.WrapRaw("reset user password", err)
		}
		u.auditor.Record(ctx, actor, "user.password_reset", "user", id.String(), nil)
		return nil
	})
	if err != nil {
		return err
	}
	if revokeSessions {
		return u.revokeSessions(id)
	}
	return nil
}

// Impersonate issues a short-lived token acting as the given user.
func (u UsersService) Impersonate(
	ctx context.Context,
	id uuid.UUID,
	durationMinutes int,
	reason string,
	actor uuid.UUID,
) (*auth.TokenResponse, error) {
	user, err := u.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if !user.IsActive() {
		return nil, users.ErrUserNotActive
	}

	token, err := u.jwt.CreateAndStoreToken(ctx, auth.CreatTokenParams{
		UserID:      user.ID.String(),
		Identifier:  user.Email,
		Status:      user.Status,
		Roles:       user.Roles,
		Permissions: user.Permissions,
	})
	if err != nil {
		return nil, errors.WrapRaw("impersonate user", err)
	}

	u.auditor.Record(ctx, actor, "user.impersonate", "user", id.String(), map[string]any{
		"reason":   reason,
		"duration": durationMinutes,
	})
	return token, nil
}

//
// ----------------------------- roles -----------------------------
//

func (u UsersService) ListRoles(
	ctx context.Context,
	userID uuid.UUID,
	limit, offset int,
) ([]access.UserRole, int, error) {
	if _, err := u.accountRepo.GetByID(ctx, userID); err != nil {
		return nil, 0, err
	}
	return u.userRoleRepo.ListByUser(ctx, userID, limit, offset)
}

// AssignRoleInput is the write-side input for a role assignment.
type AssignRoleInput struct {
	UserID    uuid.UUID
	RoleID    access.RoleID
	IsPrimary *bool
	ExpiresAt *time.Time
}

func (u UsersService) AssignRole(
	ctx context.Context,
	in AssignRoleInput,
	actor uuid.UUID,
) (bool, error) {
	if _, err := u.accountRepo.GetByID(ctx, in.UserID); err != nil {
		return false, err
	}
	isPrimary := in.IsPrimary == nil || *in.IsPrimary

	err := u.db.Run(ctx, func(ctx context.Context) error {
		if isPrimary {
			// Promote first so the new role is primary and the old one is
			// demoted, then make sure the assignment itself exists.
			if err := u.userRoleRepo.Assign(ctx, &access.UserRoleAssignment{
				UserID:     in.UserID,
				RoleID:     in.RoleID,
				AssignedBy: &actor,
				ExpiresAt:  in.ExpiresAt,
				IsPrimary:  false,
			}); err != nil {
				return errors.WrapRaw("assign user role", err)
			}
			if err := u.userRoleRepo.SetPrimary(ctx, in.UserID, in.RoleID); err != nil {
				return errors.WrapRaw("set primary user role", err)
			}
			return nil
		}
		return u.userRoleRepo.Assign(ctx, &access.UserRoleAssignment{
			UserID:     in.UserID,
			RoleID:     in.RoleID,
			AssignedBy: &actor,
			ExpiresAt:  in.ExpiresAt,
			IsPrimary:  false,
		})
	})
	if err != nil {
		return false, err
	}
	u.auditor.Record(ctx, actor, "role.assign", "user", in.UserID.String(), map[string]any{
		"roleId": in.RoleID,
	})
	return isPrimary, nil
}

func (u UsersService) RevokeRole(
	ctx context.Context,
	userID uuid.UUID,
	roleID access.RoleID,
	actor uuid.UUID,
) error {
	err := u.db.Run(ctx, func(ctx context.Context) error {
		if err := u.userRoleRepo.Unassign(ctx, userID, roleID); err != nil {
			return errors.WrapRaw("revoke user role", err)
		}
		u.auditor.Record(ctx, actor, "role.revoke", "user", userID.String(), map[string]any{
			"roleId": roleID,
		})
		return nil
	})
	return err
}

//
// --------------------------- attributes ---------------------------
//

func (u UsersService) ListAttributes(
	ctx context.Context,
	userID uuid.UUID,
	namespace *string,
	limit, offset int,
) ([]users.UserAttribute, int, error) {
	if _, err := u.accountRepo.GetByID(ctx, userID); err != nil {
		return nil, 0, err
	}
	return u.usersRepo.ListAttributes(ctx, userID, namespace, limit, offset)
}

func (u UsersService) SetAttribute(
	ctx context.Context,
	in users.UserAttribute,
	actor uuid.UUID,
) error {
	if _, err := u.accountRepo.GetByID(ctx, in.UserID); err != nil {
		return err
	}
	if err := u.usersRepo.SetAttribute(ctx, in); err != nil {
		return errors.WrapRaw("set user attribute", err)
	}
	u.auditor.Record(ctx, actor, "user.attribute_set", "user", in.UserID.String(), map[string]any{
		"namespace": in.Namespace,
		"key":       in.Key,
	})
	return nil
}

func (u UsersService) DeleteAttribute(
	ctx context.Context,
	userID uuid.UUID,
	namespace, key string,
	actor uuid.UUID,
) error {
	if namespace == "" {
		namespace = users.DefaultNamespace
	}
	if err := u.usersRepo.DeleteAttribute(ctx, userID, namespace, key); err != nil {
		return errors.WrapRaw("delete user attribute", err)
	}
	u.auditor.Record(ctx, actor, "user.attribute_delete", "user", userID.String(), map[string]any{
		"namespace": namespace,
		"key":       key,
	})
	return nil
}

func (u UsersService) ListPermissions(
	ctx context.Context,
	limit, offset int,
) ([]access.Permission, int, error) {
	return u.permissionRepo.List(ctx, limit, offset)
}

// revokeSessions drops every live session for a user. Best effort: the account
// state change already succeeded, and a Redis hiccup should not report the
// whole admin action as failed.
func (u UsersService) revokeSessions(userID uuid.UUID) error {
	if u.jwt == nil {
		return nil
	}
	if err := u.jwt.RevokeAllUserTokens(context.Background(), userID.String()); err != nil {
		logger.Error("revoke sessions for ", userID, ": ", err)
	}
	return nil
}
