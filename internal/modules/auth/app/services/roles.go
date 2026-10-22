package services

import (
	"context"
	"time"

	"modular-monolith/internal/modules/auth/domain/access"
	"modular-monolith/lib/db"
	"modular-monolith/lib/errors"
	"modular-monolith/lib/utils"

	"github.com/google/uuid"
)

type RolesService struct {
	rolesRepo  access.RoleRepository
	grantsRepo access.RolePermissionRepository
	db         *db.Manager
	auditor    auditRecorder
}

func NewRolesService(
	repo access.RoleRepository,
	grantsRepo access.RolePermissionRepository,
	auditRepo access.AuditRepository,
	db *db.Manager,
) *RolesService {
	return &RolesService{
		rolesRepo:  repo,
		grantsRepo: grantsRepo,
		db:         db,
		auditor:    newAuditRecorder(auditRepo),
	}
}

func (r RolesService) GetAll(
	ctx context.Context,
	limit, offset int,
) ([]access.Role, int, error) {
	return r.rolesRepo.List(ctx, limit, offset)
}

func (r RolesService) GetByID(ctx context.Context, id access.RoleID) (access.Role, error) {
	role, err := r.rolesRepo.GetByID(ctx, id)
	if err != nil {
		return access.Role{}, errors.WrapRaw("get role", err)
	}
	return role, nil
}

// CreateRoleInput is the write-side input for a new role.
type CreateRoleInput struct {
	Name         string
	Slug         string
	Description  string
	ParentRoleID *access.RoleID
	Metadata     map[string]any
}

func (r RolesService) Create(
	ctx context.Context,
	in CreateRoleInput,
	actor uuid.UUID,
) (access.RoleID, string, error) {
	metadata, err := utils.MarshalJSON(in.Metadata)
	if err != nil {
		return 0, "", errors.WrapRaw("create role", err)
	}

	level, err := r.resolveLevel(ctx, in.ParentRoleID)
	if err != nil {
		return 0, "", err
	}

	role := access.Role{
		Name:         in.Name,
		Slug:         in.Slug,
		Description:  in.Description,
		ParentRoleID: in.ParentRoleID,
		Level:        level,
		Metadata:     metadata,
	}

	var id access.RoleID
	if err := r.db.Run(ctx, func(ctx context.Context) error {
		if err := r.rolesRepo.Create(ctx, role); err != nil {
			return errors.WrapRaw("create role", err)
		}
		id = role.ID
		r.auditor.Record(ctx, actor, "role.create", "role", resourceID(int64(id)), nil)
		return nil
	}); err != nil {
		return 0, "", err
	}
	return id, in.Slug, nil
}

func (r RolesService) Update(
	ctx context.Context,
	id access.RoleID,
	in UpdateRoleInput,
	actor uuid.UUID,
) error {
	current, err := r.rolesRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if current.IsProtected() {
		return access.ErrRoleIsSystem
	}
	if in.ParentRoleID != nil {
		if err := r.assertNoCycle(ctx, id, *in.ParentRoleID); err != nil {
			return err
		}
	}

	level, err := r.resolveLevel(ctx, in.ParentRoleID)
	if err != nil {
		return err
	}

	before := current
	current.Name = utils.Pick(in.Name, current.Name)
	current.Description = utils.Pick(in.Description, current.Description)
	current.ParentRoleID = in.ParentRoleID
	current.Level = level
	if in.Metadata != nil {
		metadata, err := utils.MarshalJSON(in.Metadata)
		if err != nil {
			return errors.WrapRaw("update role", err)
		}
		current.Metadata = metadata
	}

	return r.db.Run(ctx, func(ctx context.Context) error {
		if err := r.rolesRepo.Update(ctx, &current); err != nil {
			return errors.WrapRaw("update role", err)
		}
		r.auditor.Record(ctx, actor, "role.update", "role", resourceID(int64(id)), &before)
		return nil
	})
}

// UpdateRoleInput is a partial update. ParentRoleID is distinct from the other
// fields because nil means "detach from parent", not "leave unchanged".
type UpdateRoleInput struct {
	Name         *string
	Description  *string
	ParentRoleID *access.RoleID
	Metadata     map[string]any
}

func (r RolesService) Delete(ctx context.Context, id access.RoleID, actor uuid.UUID) error {
	role, err := r.rolesRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if role.IsProtected() {
		return access.ErrRoleIsSystem
	}
	return r.db.Run(ctx, func(ctx context.Context) error {
		if err := r.rolesRepo.Delete(ctx, id); err != nil {
			return errors.WrapRaw("delete role", err)
		}
		r.auditor.Record(ctx, actor, "role.delete", "role", resourceID(int64(id)), &role)
		return nil
	})
}

func (r RolesService) ListPermissions(
	ctx context.Context,
	roleID access.RoleID,
	limit, offset int,
) ([]access.RolePermission, int, error) {
	if _, err := r.rolesRepo.GetByID(ctx, roleID); err != nil {
		return nil, 0, err
	}
	return r.grantsRepo.ListByRole(ctx, roleID, limit, offset)
}

// GrantPermissionInput is the write-side input for a role permission grant.
type GrantPermissionInput struct {
	RoleID       access.RoleID
	PermissionID access.PermissionID
	Conditions   map[string]any
	ExpiresAt    *time.Time
}

func (r RolesService) GrantPermission(
	ctx context.Context,
	in GrantPermissionInput,
	actor uuid.UUID,
) error {
	conditions, err := utils.MarshalJSON(in.Conditions)
	if err != nil {
		return errors.WrapRaw("grant permission", err)
	}
	grant := &access.RolePermissionGrant{
		RoleID:       in.RoleID,
		PermissionID: in.PermissionID,
		Conditions:   conditions,
		ExpiresAt:    in.ExpiresAt,
		GrantedBy:    &actor,
	}
	if err := r.grantsRepo.Grant(ctx, grant); err != nil {
		return errors.WrapRaw("grant permission", err)
	}
	return nil
}

func (r RolesService) RevokePermission(
	ctx context.Context,
	roleID access.RoleID,
	permissionID access.PermissionID,
	actor uuid.UUID,
) error {
	if err := r.grantsRepo.Revoke(ctx, roleID, permissionID); err != nil {
		return errors.WrapRaw("revoke permission", err)
	}
	r.auditor.Record(ctx, actor, "permission.revoke", "role", resourceID(int64(roleID)), nil)
	return nil
}

// resolveLevel computes the depth of a role in the hierarchy. A root role is
// level 0; a child sits one below its parent.
func (r RolesService) resolveLevel(
	ctx context.Context,
	parentID *access.RoleID,
) (int32, error) {
	if parentID == nil {
		return 0, nil
	}
	parent, err := r.rolesRepo.GetByID(ctx, *parentID)
	if err != nil {
		return 0, errors.WrapRaw("resolve role level", err)
	}
	return parent.Level + 1, nil
}

// assertNoCycle walks up from the proposed parent and refuses to link a role to
// one of its own descendants.
func (r RolesService) assertNoCycle(
	ctx context.Context,
	roleID access.RoleID,
	parentID access.RoleID,
) error {
	if roleID == parentID {
		return access.ErrRoleParentCycle
	}
	seen := map[access.RoleID]bool{roleID: true}
	for cursor := parentID; cursor != 0; {
		if seen[cursor] {
			return access.ErrRoleParentCycle
		}
		seen[cursor] = true
		parent, err := r.rolesRepo.GetByID(ctx, cursor)
		if err != nil {
			return err
		}
		if parent.ParentRoleID == nil {
			return nil
		}
		cursor = *parent.ParentRoleID
	}
	return nil
}
