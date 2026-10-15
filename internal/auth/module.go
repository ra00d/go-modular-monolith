// Package auth wires the authentication & authorization module into the app.
package auth

import (
	"time"

	"modular-monolith/internal/auth/api/rest"
	"modular-monolith/internal/auth/app"
	"modular-monolith/internal/auth/infra/abac"
	"modular-monolith/internal/auth/infra/rbac"
	users_infra "modular-monolith/internal/auth/infra/users"
	"modular-monolith/internal/auth/migrations"
	"modular-monolith/lib/auth"
	"modular-monolith/lib/db"
	"modular-monolith/lib/di"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
)

// policyCacheTTL bounds how long ABAC policies are cached in Redis.
// ponytail: 5m is fine until policy edits need to take effect faster; then
// invalidate the cache key on write instead of shortening the TTL.
const policyCacheTTL = 5 * time.Minute

type MoudleRepos struct {
	fx.In
	Account        *users_infra.AccountRepository
	Users          *users_infra.UserRepository
	Roles          *rbac.RoleRepository
	Permissions    *rbac.PermissionRepository
	RolePermission *rbac.RolePermissionRepository
	UserRole       *rbac.UserRoleRepository
	Policies       *abac.PolicyRepository
	Audit          *rbac.AuditRepository
	SecurityEvents *rbac.SecurityEventRepository
	APIKeys        *rbac.APIKeyRepository
}

var CoreModule = fx.Module("Core", fx.Provide(func(in di.AppBaseParams) *auth.JWTManager {
	jwt := auth.NewJWTManager(in.Redis, in.Config.Jwt)
	return jwt
}), fx.Provide(func(in di.AppBaseParams, jwt *auth.JWTManager, az *auth.Authorizer) *auth.Guard {
	guard := auth.NewGuard(jwt, az)
	return guard
}))

var Module = fx.Module(
	"auth",
	fx.Provide(users_infra.NewAccountRepository),
	fx.Provide(users_infra.NewUsersAdminRepository),
	fx.Provide(abac.NewPolicyRepository),
	fx.Provide(rbac.NewRoleRepository),
	fx.Provide(rbac.NewPermissionRepository),
	fx.Provide(rbac.NewRolePermissionRepository),
	fx.Provide(rbac.NewUserRoleRepository),
	fx.Provide(rbac.NewAuditRepository),
	fx.Provide(rbac.NewSecurityEventRepository),
	fx.Provide(rbac.NewAPIKeyRepository),
	fx.Provide(func(repos MoudleRepos, db *db.Manager, jwt *auth.JWTManager) app.AppServices {
		return app.NewAppServices(db, jwt, app.AppRepos{
			Account:        repos.Account,
			Users:          repos.Users,
			Roles:          repos.Roles,
			Permission:     repos.Permissions,
			RolePermission: repos.RolePermission,
			UserRole:       repos.UserRole,
			Policies:       repos.Policies,
			Audit:          repos.Audit,
			SecurityEvents: repos.SecurityEvents,
			APIKeys:        repos.APIKeys,
		})
	}),
	// fx.Provide(
	// 	func(repo *users_infra.AccountRepository, db *db.Manager) *services.AccountService {
	// 		return services.NewAccountService(repo, db)
	// 	},
	// ),
	// fx.Provide(
	// 	func(repo *admin_infra.UsersAdminRepository, db *db.Manager, userRepo *users_infra.AccountRepository, rolesRepo *rbac.RoleRepository, permRepo *rbac.PermissionRepository) *services.AdminService {
	// 		return services.NewAdminService(db, repo, userRepo, rolesRepo, permRepo)
	// 	},
	// ),
	// fx.Provide(
	// 	func(repo *abac.PolicyRepository, db *db.Manager) *services.PoliciesService {
	// 		return services.NewPoliciesService(repo, db)
	// 	},
	// ),
	fx.Provide(
		func(db *sqlx.DB, rdb *redis.Client, repo *users_infra.AccountRepository) *auth.Authorizer {
			store := auth.NewCachedPolicyStore(
				abac.NewPolicyStore(db),
				abac.NewRedisCache(rdb),
				policyCacheTTL,
			)
			engine := auth.NewPolicyEngine(
				store,
				abac.NewAttributeResolver(db),
				auth.NewDefaultConditionEvaluator(),
			)
			return auth.NewAuthorizer(engine, repo)
		},
	),
	CoreModule,
	fx.Invoke(rest.RegisterRoutes),
	fx.Invoke(db.MigrateAndGenerateTypes(migrations.Migrations)),
)
