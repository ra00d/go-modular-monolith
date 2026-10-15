-- +goose Up
-- Seed the tenants module's permissions and bind them to the
-- admin role. super_admin inherits all permissions implicitly (see
-- lib/auth).
--
-- These tables live in the "auth" schema (set by the auth module's
-- MigrateAndGenerateTypes via search_path); the INSERTs are unqualified
-- so they target the active search_path.

-- Drop the legacy business.* permissions seeded before this module was
-- renamed to tenants. Safe on fresh installs (no rows yet).
DELETE FROM "role_permission"
    WHERE "permission_id" IN (SELECT "id" FROM "permission" WHERE "slug" IN ('business.manage','business.read'));
DELETE FROM "permission" WHERE "slug" IN ('business.manage','business.read');

INSERT INTO "permission" ("name", "slug", "description", "resource", "action", "category", "is_system")
VALUES
    ('Manage Tenants',     'tenant.manage', 'Full tenant & location management', 'tenant', 'manage', 'admin', TRUE),
    ('Read Tenant',         'tenant.read',   'View the current tenant profile',   'tenant', 'read',   'admin', TRUE)
ON CONFLICT ("slug") DO NOTHING;

-- Grant tenant.manage to the 'admin' role (id = 2 in the seed).
INSERT INTO "role_permission" ("role_id", "permission_id")
SELECT 2, p."id" FROM "permission" p WHERE p."slug" = 'tenant.manage'
ON CONFLICT ("role_id", "permission_id") DO NOTHING;

-- +goose Down
DELETE FROM "role_permission"
    WHERE "permission_id" IN (SELECT "id" FROM "permission" WHERE "slug" IN ('tenant.manage','tenant.read'));
DELETE FROM "permission" WHERE "slug" IN ('tenant.manage','tenant.read');
