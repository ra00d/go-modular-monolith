package services

import (
	"context"

	"modular-monolith/internal/auth/domain/access"
	"modular-monolith/lib/db"
	"modular-monolith/lib/errors"
	"modular-monolith/lib/utils"

	"github.com/google/uuid"
)

type PermissionsService struct {
	permissionRepo access.PermissionRepository
	db             *db.Manager
	auditor        auditRecorder
}

func NewPermissionsService(
	repo access.PermissionRepository,
	auditRepo access.AuditRepository,
	db *db.Manager,
) *PermissionsService {
	return &PermissionsService{
		permissionRepo: repo,
		db:             db,
		auditor:        newAuditRecorder(auditRepo),
	}
}

func (p PermissionsService) GetAll(
	ctx context.Context,
	limit, offset int,
) ([]access.Permission, int, error) {
	return p.permissionRepo.List(ctx, limit, offset)
}

func (p PermissionsService) GetByID(
	ctx context.Context,
	id access.PermissionID,
) (access.Permission, error) {
	permission, err := p.permissionRepo.GetByID(ctx, id)
	if err != nil {
		return access.Permission{}, errors.WrapRaw("get permission", err)
	}
	return permission, nil
}

// CreatePermissionInput is the write-side input for a new permission.
type CreatePermissionInput struct {
	Name        string
	Slug        string
	Description string
	Resource    string
	Action      string
	Category    string
	Conditions  map[string]any
}

func (p PermissionsService) Create(
	ctx context.Context,
	in CreatePermissionInput,
	actor uuid.UUID,
) (access.PermissionID, string, error) {
	conditions, err := utils.MarshalJSON(in.Conditions)
	if err != nil {
		return 0, "", errors.WrapRaw("create permission", err)
	}
	permission := &access.Permission{
		Name:        in.Name,
		Slug:        in.Slug,
		Description: in.Description,
		Resource:    in.Resource,
		Action:      in.Action,
		Category:    in.Category,
		Conditions:  conditions,
	}

	err = p.db.Run(ctx, func(ctx context.Context) error {
		if err := p.permissionRepo.Create(ctx, permission); err != nil {
			return errors.WrapRaw("create permission", err)
		}
		p.auditor.Record(ctx, actor, "permission.create", "permission", resourceID(int64(permission.ID)), nil)
		return nil
	})
	if err != nil {
		return 0, "", err
	}
	return permission.ID, in.Slug, nil
}

// UpdatePermissionInput is a partial update of a permission.
type UpdatePermissionInput struct {
	Name        *string
	Description *string
	Category    *string
	Conditions  map[string]any
}

func (p PermissionsService) Update(
	ctx context.Context,
	id access.PermissionID,
	in UpdatePermissionInput,
	actor uuid.UUID,
) error {
	current, err := p.permissionRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if current.IsProtected() {
		return access.ErrPermissionIsSystem
	}

	before := current
	current.Name = utils.Pick(in.Name, current.Name)
	current.Description = utils.Pick(in.Description, current.Description)
	current.Category = utils.Pick(in.Category, current.Category)
	if in.Conditions != nil {
		conditions, err := utils.MarshalJSON(in.Conditions)
		if err != nil {
			return errors.WrapRaw("update permission", err)
		}
		current.Conditions = conditions
	}

	err = p.db.Run(ctx, func(ctx context.Context) error {
		if err := p.permissionRepo.Update(ctx, &current); err != nil {
			return errors.WrapRaw("update permission", err)
		}
		p.auditor.Record(ctx, actor, "permission.update", "permission", resourceID(int64(id)), &before)
		return nil
	})
	return err
}

func (p PermissionsService) Delete(
	ctx context.Context,
	id access.PermissionID,
	actor uuid.UUID,
) error {
	permission, err := p.permissionRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if permission.IsProtected() {
		return access.ErrPermissionIsSystem
	}
	return p.db.Run(ctx, func(ctx context.Context) error {
		if err := p.permissionRepo.Delete(ctx, id); err != nil {
			return errors.WrapRaw("delete permission", err)
		}
		p.auditor.Record(ctx, actor, "permission.delete", "permission", resourceID(int64(id)), &permission)
		return nil
	})
}
