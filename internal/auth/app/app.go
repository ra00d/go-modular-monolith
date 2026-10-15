// Package app provides application services for the auth module.
package app

import (
	"modular-monolith/internal/auth/app/services"
	"modular-monolith/internal/auth/domain/access"
	"modular-monolith/internal/auth/domain/users"
	"modular-monolith/lib/auth"
	"modular-monolith/lib/db"

	"go.uber.org/fx"
)

type AppRepos struct {
	Users          users.UserRepository
	Account        users.AccountRepository
	Roles          access.RoleRepository
	Permission     access.PermissionRepository
	RolePermission access.RolePermissionRepository
	UserRole       access.UserRoleRepository
	Policies       access.PolicyRepository
	Audit          access.AuditRepository
	SecurityEvents access.SecurityEventRepository
	APIKeys        access.APIKeyRepository
}

type AppServices struct {
	fx.Out
	Users      *services.UsersService
	Account    *services.AccountService
	Roles      *services.RolesService
	Permission *services.PermissionsService
	Policies   *services.PoliciesService
	Audit      *services.AuditService
	APIKeys    *services.APIKeyService
}

func NewAppServices(db *db.Manager, jwt *auth.JWTManager, repos AppRepos) AppServices {
	return AppServices{
		Users: services.NewUsersService(
			db,
			repos.Users,
			repos.Account,
			repos.Permission,
			repos.UserRole,
			repos.Audit,
			jwt,
		),
		Account:    services.NewAccountService(repos.Account, repos.Users, db),
		Roles:      services.NewRolesService(repos.Roles, repos.RolePermission, repos.Audit, db),
		Permission: services.NewPermissionsService(repos.Permission, repos.Audit, db),
		Policies:   services.NewPoliciesService(repos.Policies, db),
		Audit:      services.NewAuditService(repos.Audit, repos.SecurityEvents),
		APIKeys:    services.NewAPIKeyService(repos.APIKeys),
	}
}
