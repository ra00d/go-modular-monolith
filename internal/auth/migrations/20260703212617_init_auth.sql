-- +goose Up
-- ============================================
-- CORE TABLES (Enhanced from your base)
-- ============================================

-- Enable UUID extension
-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Users table with multi-login support
CREATE TABLE "user" (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL,
    
    -- Profile info
    "name" TEXT NOT NULL,
    "email" TEXT NULL,
    "phone" TEXT NULL,
    "username" TEXT NULL,
    "email_verified" BOOLEAN NOT NULL DEFAULT FALSE,
    "phone_verified" BOOLEAN NOT NULL DEFAULT FALSE,
    "image" TEXT NULL,
    
    -- Security
    "password_hash" TEXT NULL,           -- NULL for social-only accounts
    "mfa_enabled" BOOLEAN NOT NULL DEFAULT FALSE,
    "mfa_secret" TEXT NULL,              -- Encrypted TOTP secret
    "mfa_backup_codes" TEXT NULL,        -- JSON array of hashed backup codes
    
    -- Account status
    "status" TEXT NOT NULL DEFAULT 'active' CHECK ("status" IN ('active', 'suspended', 'locked', 'pending_verification')),
    "failed_login_attempts" INTEGER NOT NULL DEFAULT 0,
    "locked_until" TIMESTAMP NULL,
    "last_login_at" TIMESTAMP NULL,
    "last_login_ip" TEXT NULL,
    
    -- Metadata
    "metadata" TEXT NULL                 -- JSON for ABAC attributes
);

CREATE INDEX "user_deleted_at_idx" ON "user" ("deleted_at");
CREATE UNIQUE INDEX "user_email_unique" ON "user" ("email") WHERE "email" IS NOT NULL;
CREATE UNIQUE INDEX "user_phone_unique" ON "user" ("phone") WHERE "phone" IS NOT NULL;
CREATE UNIQUE INDEX "user_username_unique" ON "user" ("username") WHERE "username" IS NOT NULL;
CREATE INDEX "user_status_idx" ON "user" ("status");

-- Social login accounts (OAuth/OIDC)
CREATE TABLE "account" (
    "id" SERIAL PRIMARY KEY,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL,
    
    "user_id" UUID NOT NULL,
    "provider_id" TEXT NOT NULL,          -- 'google', 'github', 'apple', etc.
    "provider_account_id" TEXT NOT NULL,  -- ID from provider
    
    -- OAuth tokens
    "access_token" TEXT NULL,
    "refresh_token" TEXT NULL,
    "access_token_expires_at" TIMESTAMP NULL,
    "refresh_token_expires_at" TIMESTAMP NULL,
    "scope" TEXT NULL,
    "id_token" TEXT NULL,
    
    -- Provider profile cache
    "provider_data" TEXT NULL,            -- JSON: email, name, image from provider
    
    FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE
);

CREATE INDEX "account_deleted_at_idx" ON "account" ("deleted_at");
CREATE INDEX "account_user_id_idx" ON "account" ("user_id");
CREATE UNIQUE INDEX "account_provider_unique" ON "account" ("provider_id", "provider_account_id");

-- ============================================
-- VERIFICATION SYSTEM
-- ============================================

CREATE TABLE "verification" (
    "id" SERIAL PRIMARY KEY,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL,
    
    "user_id" UUID NULL,                  -- NULL for pre-registration verification
    "identifier" TEXT NOT NULL,           -- email, phone, or username
    "identifier_type" TEXT NOT NULL CHECK ("identifier_type" IN ('email', 'phone', 'username')),
    
    -- Verification details
    "purpose" TEXT NOT NULL CHECK ("purpose" IN (
        'email_verification', 
        'phone_verification', 
        'password_reset', 
        'account_recovery',
        'mfa_setup',
        'login_challenge',
        'email_change',
        'phone_change'
    )),
    "code_hash" TEXT NOT NULL,            -- Hashed OTP/code
    "attempts" INTEGER NOT NULL DEFAULT 0,
    "max_attempts" INTEGER NOT NULL DEFAULT 5,
    
    -- Expiration
    "expires_at" TIMESTAMP NOT NULL,
    "verified_at" TIMESTAMP NULL,
    
    -- Context
    "ip_address" TEXT NULL,
    "user_agent" TEXT NULL,
    
    FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE
);

CREATE INDEX "verification_deleted_at_idx" ON "verification" ("deleted_at");
CREATE INDEX "verification_identifier_idx" ON "verification" ("identifier", "purpose", "expires_at");
CREATE INDEX "verification_user_id_idx" ON "verification" ("user_id");

-- Password reset tokens (separate for extra security)
CREATE TABLE "password_reset" (
    "id" SERIAL PRIMARY KEY,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    "user_id" UUID NOT NULL,
    "token_hash" TEXT NOT NULL,
    "expires_at" TIMESTAMP NOT NULL,
    "used_at" TIMESTAMP NULL,
    "ip_address" TEXT NULL,
    
    FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE
);

CREATE INDEX "password_reset_user_id_idx" ON "password_reset" ("user_id");
CREATE INDEX "password_reset_token_idx" ON "password_reset" ("token_hash", "expires_at");

-- ============================================
-- RBAC: ROLE-BASED ACCESS CONTROL
-- ============================================

-- Hierarchical roles with inheritance
CREATE TABLE "role" (
    "id" SERIAL PRIMARY KEY,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL,
    
    "name" TEXT NOT NULL,
    "slug" TEXT NOT NULL,
    "description" TEXT NOT NULL DEFAULT '',
    
    -- Hierarchy
    "parent_role_id" INTEGER NULL,
    "level" INTEGER NOT NULL DEFAULT 0,    -- Depth in hierarchy for query optimization
    
    -- Metadata
    "is_system" BOOLEAN NOT NULL DEFAULT FALSE,  -- Protected roles (admin, user, etc.)
    "metadata" TEXT NULL,                  -- JSON for ABAC context on roles
    
    FOREIGN KEY ("parent_role_id") REFERENCES "role" ("id")
);

CREATE INDEX "role_deleted_at_idx" ON "role" ("deleted_at");
CREATE INDEX "role_parent_role_id_idx" ON "role" ("parent_role_id");
CREATE UNIQUE INDEX "role_slug_unique" ON "role" ("slug");

-- Permissions registry (granular actions)
CREATE TABLE "permission" (
    "id" SERIAL PRIMARY KEY,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL,
    
    "name" TEXT NOT NULL,
    "slug" TEXT NOT NULL,
    "description" TEXT NOT NULL DEFAULT '',
    "resource" TEXT NOT NULL,              -- 'user', 'post', 'order', etc.
    "action" TEXT NOT NULL,                -- 'create', 'read', 'update', 'delete', 'list', 'manage'
    
    -- Grouping
    "category" TEXT NULL,                  -- 'user-management', 'content', 'billing'
    "is_system" BOOLEAN NOT NULL DEFAULT FALSE,
    
    -- ABAC conditions (optional JSON for dynamic evaluation)
    "conditions" TEXT NULL               -- JSON: {"owner_only": true, "time_restricted": false}
);

CREATE INDEX "permission_deleted_at_idx" ON "permission" ("deleted_at");
CREATE UNIQUE INDEX "permission_slug_unique" ON "permission" ("slug");
CREATE INDEX "permission_resource_action_idx" ON "permission" ("resource", "action");

-- Role-Permission mapping (many-to-many with conditions)
CREATE TABLE "role_permission" (
    "id" SERIAL PRIMARY KEY,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    "role_id" INTEGER NOT NULL,
    "permission_id" INTEGER NOT NULL,
    
    -- ABAC: Conditional grants
    "conditions" TEXT NULL,               -- JSON override for this role-permission pair
    "expires_at" TIMESTAMP NULL,           -- Temporary permissions
    "granted_by" UUID NULL,                -- Admin who granted it (UUID references user)
    
    FOREIGN KEY ("role_id") REFERENCES "role" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("permission_id") REFERENCES "permission" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("granted_by") REFERENCES "user" ("id") ON DELETE SET NULL,
    
    UNIQUE ("role_id", "permission_id")
);

CREATE INDEX "role_permission_role_id_idx" ON "role_permission" ("role_id");
CREATE INDEX "role_permission_permission_id_idx" ON "role_permission" ("permission_id");

-- User-Role mapping (many-to-many)
CREATE TABLE "user_role" (
    "user_id" UUID NOT NULL,
    "role_id" INTEGER NOT NULL,
    
    "assigned_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "assigned_by" UUID NULL,
    "expires_at" TIMESTAMP NULL,
    "is_primary" BOOLEAN NOT NULL DEFAULT FALSE,  -- Primary role for default context
    
    PRIMARY KEY ("user_id", "role_id"),
    FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("role_id") REFERENCES "role" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("assigned_by") REFERENCES "user" ("id") ON DELETE SET NULL
);

CREATE INDEX "user_role_role_id_idx" ON "user_role" ("role_id");
CREATE INDEX "user_role_user_id_idx" ON "user_role" ("user_id");

-- ============================================
-- ABAC: ATTRIBUTE-BASED ACCESS CONTROL
-- ============================================

-- Dynamic policies for fine-grained authorization
CREATE TABLE "abac_policy" (
    "id" SERIAL PRIMARY KEY,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL,
    
    "name" TEXT NOT NULL,
    "slug" TEXT NOT NULL,
    "description" TEXT NULL,
    
    -- Target
    "resource" TEXT NOT NULL,             -- Entity being accessed
    "action" TEXT NOT NULL,               -- Action being performed
    
    -- Policy logic (JSON for flexible rules)
    "rules" TEXT NOT NULL,                -- JSON: {
                                          --   "subject": {"role": "editor", "department": "engineering"},
                                          --   "resource": {"status": "published", "owner_id": "{{subject.id}}"},
                                          --   "environment": {"time": {"start": "09:00", "end": "18:00"}}
                                          -- }
    
    -- Effect
    "effect" TEXT NOT NULL DEFAULT 'allow' CHECK ("effect" IN ('allow', 'deny')),
    "priority" INTEGER NOT NULL DEFAULT 100,  -- Lower = higher priority
    
    "is_active" BOOLEAN NOT NULL DEFAULT TRUE,
    "expires_at" TIMESTAMP NULL,
    
    "created_by" UUID NULL,
    FOREIGN KEY ("created_by") REFERENCES "user" ("id") ON DELETE SET NULL
);

CREATE INDEX "abac_policy_deleted_at_idx" ON "abac_policy" ("deleted_at");
CREATE UNIQUE INDEX "abac_policy_slug_unique" ON "abac_policy" ("slug");
CREATE INDEX "abac_policy_resource_action_idx" ON "abac_policy" ("resource", "action", "is_active");

-- User attributes for ABAC evaluation
CREATE TABLE "user_attribute" (
    "id" SERIAL PRIMARY KEY,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    "user_id" UUID NOT NULL,
    "namespace" TEXT NOT NULL DEFAULT 'default',  -- 'org', 'project', 'system'
    "key" TEXT NOT NULL,
    "value" TEXT NOT NULL,                -- Stored as string, parsed based on type
    "value_type" TEXT NOT NULL DEFAULT 'string' CHECK ("value_type" IN ('string', 'number', 'boolean', 'json', 'date')),
    
    FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE,
    UNIQUE ("user_id", "namespace", "key")
);

CREATE INDEX "user_attribute_user_id_idx" ON "user_attribute" ("user_id");
CREATE INDEX "user_attribute_namespace_key_idx" ON "user_attribute" ("namespace", "key");

-- ============================================
-- AUDIT LOGGING SYSTEM
-- ============================================

-- Comprehensive audit trail
CREATE TABLE "audit_log" (
    "id" SERIAL PRIMARY KEY,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    -- Actor
    "actor_type" TEXT NOT NULL DEFAULT 'user' CHECK ("actor_type" IN ('user', 'system', 'api_key', 'service')),
    "actor_id" TEXT NULL,                 -- User UUID or service name
    "actor_role" TEXT NULL,               -- Role at time of action
    
    -- Action
    "action" TEXT NOT NULL,               -- 'user.login', 'user.update', 'post.delete'
    "category" TEXT NOT NULL DEFAULT 'general',  -- 'auth', 'data', 'admin', 'security'
    
    -- Target
    "resource_type" TEXT NOT NULL,        -- 'user', 'post', 'role', 'permission'
    "resource_id" TEXT NULL,
    
    -- Context
    "ip_address" TEXT NULL,
    "user_agent" TEXT NULL,
    "session_id" TEXT NULL,
    "request_id" TEXT NULL,
    
    -- Details
    "before_state" TEXT NULL,             -- JSON: previous values
    "after_state" TEXT NULL,              -- JSON: new values
    "metadata" TEXT NULL,                 -- JSON: additional context
    
    -- Result
    "status" TEXT NOT NULL DEFAULT 'success' CHECK ("status" IN ('success', 'failure', 'denied', 'error')),
    "error_message" TEXT NULL,
    
    -- Integrity
    "checksum" TEXT NULL                  -- HMAC for tamper detection
);

CREATE INDEX "audit_log_created_at_idx" ON "audit_log" ("created_at");
CREATE INDEX "audit_log_actor_idx" ON "audit_log" ("actor_type", "actor_id");
CREATE INDEX "audit_log_action_idx" ON "audit_log" ("action", "created_at");
CREATE INDEX "audit_log_resource_idx" ON "audit_log" ("resource_type", "resource_id");
CREATE INDEX "audit_log_category_status_idx" ON "audit_log" ("category", "status");
CREATE INDEX "audit_log_session_idx" ON "audit_log" ("session_id");

-- Security events (high-priority audit subset)
CREATE TABLE "security_event" (
    "id" SERIAL PRIMARY KEY,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    "event_type" TEXT NOT NULL CHECK ("event_type" IN (
        'login_success', 'login_failure', 'logout', 'session_expired',
        'password_changed', 'password_reset_requested', 'password_reset_completed',
        'mfa_enabled', 'mfa_disabled', 'mfa_challenge_failed',
        'permission_denied', 'role_changed', 'account_locked',
        'suspicious_activity', 'brute_force_detected', 'token_refresh'
    )),
    
    "user_id" UUID NULL,
    "severity" TEXT NOT NULL DEFAULT 'info' CHECK ("severity" IN ('info', 'warning', 'critical')),
    
    "details" TEXT NULL,                  -- JSON
    "ip_address" TEXT NULL,
    "resolved_at" TIMESTAMP NULL,
    "resolved_by" UUID NULL,
    
    FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE SET NULL,
    FOREIGN KEY ("resolved_by") REFERENCES "user" ("id") ON DELETE SET NULL
);

CREATE INDEX "security_event_created_at_idx" ON "security_event" ("created_at");
CREATE INDEX "security_event_user_id_idx" ON "security_event" ("user_id");
CREATE INDEX "security_event_type_idx" ON "security_event" ("event_type", "severity");

-- ============================================
-- API KEYS & SERVICE ACCOUNTS
-- ============================================

CREATE TABLE "api_key" (
    "id" SERIAL PRIMARY KEY,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL,
    
    "name" TEXT NOT NULL,
    "key_prefix" TEXT NOT NULL,           -- First 8 chars for identification
    "key_hash" TEXT NOT NULL,             -- Full hash of the key
    
    "user_id" UUID NULL,                  -- NULL for service accounts
    "role_id" INTEGER NULL,               -- Assigned role for this key
    
    -- Scopes (RBAC override)
    "permissions" TEXT NULL,              -- JSON array of permission slugs
    
    -- Limits
    "rate_limit" INTEGER NULL,          -- Requests per minute
    "expires_at" TIMESTAMP NULL,
    "last_used_at" TIMESTAMP NULL,
    
    "is_active" BOOLEAN NOT NULL DEFAULT TRUE,
    "created_by" UUID NULL,
    
    FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("role_id") REFERENCES "role" ("id") ON DELETE SET NULL,
    FOREIGN KEY ("created_by") REFERENCES "user" ("id") ON DELETE SET NULL
);

CREATE INDEX "api_key_deleted_at_idx" ON "api_key" ("deleted_at");
CREATE INDEX "api_key_user_id_idx" ON "api_key" ("user_id");
CREATE UNIQUE INDEX "api_key_prefix_unique" ON "api_key" ("key_prefix");

-- ============================================
-- DYNAMIC PERMISSION MANAGEMENT
-- ============================================

-- Permission templates for quick assignment
CREATE TABLE "permission_template" (
    "id" SERIAL PRIMARY KEY,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    "name" TEXT NOT NULL,
    "description" TEXT NULL,
    "permissions" TEXT NOT NULL,          -- JSON array of permission IDs/slugs
    
    "is_system" BOOLEAN NOT NULL DEFAULT FALSE
);

-- Permission change history
CREATE TABLE "permission_change_log" (
    "id" SERIAL PRIMARY KEY,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    "changed_by" UUID NOT NULL,
    "change_type" TEXT NOT NULL CHECK ("change_type" IN ('grant', 'revoke', 'modify')),
    
    "target_type" TEXT NOT NULL CHECK ("target_type" IN ('role', 'user', 'api_key')),
    "target_id" TEXT NOT NULL,            -- Can be UUID (user) or INTEGER (role/api_key)
    
    "permission_id" INTEGER NOT NULL,
    "previous_conditions" TEXT NULL,
    "new_conditions" TEXT NULL,
    
    FOREIGN KEY ("changed_by") REFERENCES "user" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("permission_id") REFERENCES "permission" ("id") ON DELETE CASCADE
);

-- ============================================
-- SEED DATA
-- ============================================

-- Core roles
INSERT INTO "role" ("name", "slug", "description", "is_system", "level") VALUES
('Super Administrator', 'super_admin', 'Full system access', TRUE, 0),
('Administrator', 'admin', 'Administrative access with some restrictions', TRUE, 1),
('Moderator', 'moderator', 'Content and user moderation', TRUE, 2),
('User', 'user', 'Standard authenticated user', TRUE, 3),
('Guest', 'guest', 'Unauthenticated visitor', TRUE, 4);

-- Core permissions
INSERT INTO "permission" ("name", "slug", "description", "resource", "action", "category", "is_system") VALUES
-- User management
('Create Users', 'user.create', 'Create new user accounts', 'user', 'create', 'user-management', TRUE),
('Read Users', 'user.read', 'View user profiles', 'user', 'read', 'user-management', TRUE),
('Update Users', 'user.update', 'Modify user accounts', 'user', 'update', 'user-management', TRUE),
('Delete Users', 'user.delete', 'Delete user accounts', 'user', 'delete', 'user-management', TRUE),
('List Users', 'user.list', 'View user lists', 'user', 'list', 'user-management', TRUE),
('Manage Users', 'user.manage', 'Full user management', 'user', 'manage', 'user-management', TRUE),

-- Role management
('Create Roles', 'role.create', 'Create new roles', 'role', 'create', 'admin', TRUE),
('Read Roles', 'role.read', 'View roles', 'role', 'read', 'admin', TRUE),
('Update Roles', 'role.update', 'Modify roles', 'role', 'update', 'admin', TRUE),
('Delete Roles', 'role.delete', 'Delete roles', 'role', 'delete', 'admin', TRUE),
('Assign Roles', 'role.assign', 'Assign roles to users', 'role', 'assign', 'admin', TRUE),

-- Permission management
('Create Permissions', 'permission.create', 'Create permissions', 'permission', 'create', 'admin', TRUE),
('Read Permissions', 'permission.read', 'View permissions', 'permission', 'read', 'admin', TRUE),
('Update Permissions', 'permission.update', 'Modify permissions', 'permission', 'update', 'admin', TRUE),
('Delete Permissions', 'permission.delete', 'Delete permissions', 'permission', 'delete', 'admin', TRUE),
('Assign Permissions', 'permission.assign', 'Assign permissions to roles', 'permission', 'assign', 'admin', TRUE),

-- Audit
('Read Audit Logs', 'audit.read', 'View audit logs', 'audit', 'read', 'admin', TRUE),
('Export Audit Logs', 'audit.export', 'Export audit logs', 'audit', 'export', 'admin', TRUE),

-- System
('Manage System', 'system.manage', 'System configuration', 'system', 'manage', 'admin', TRUE),
('View Analytics', 'analytics.read', 'View system analytics', 'analytics', 'read', 'admin', TRUE);

-- Assign all permissions to super_admin
INSERT INTO "role_permission" ("role_id", "permission_id", "conditions")
SELECT 1, "id", NULL FROM "permission";

-- Assign basic permissions to user role
INSERT INTO "role_permission" ("role_id", "permission_id", "conditions")
SELECT 4, "id", '{"owner_only": true}' FROM "permission" 
WHERE "resource" = 'user' AND "action" IN ('read', 'update');

-- Guest permissions (minimal)
INSERT INTO "role_permission" ("role_id", "permission_id", "conditions")
VALUES (5, (SELECT "id" FROM "permission" WHERE "slug" = 'user.read'), '{"public_profile": true}');

-- Sample ABAC policy: Users can only edit their own posts during business hours
INSERT INTO "abac_policy" ("name", "slug", "description", "resource", "action", "rules", "effect", "priority") VALUES
('Own Content Editor', 'own_content_edit', 'Users can edit their own content', 'post', 'update', '{
    "subject": {"role": "user"},
    "resource": {"author_id": "{{subject.id}}"},
    "environment": {"time": {"start": "00:00", "end": "23:59"}}
}', 'allow', 100);

-- Sample ABAC policy: Admin access restricted to office IP during off-hours
INSERT INTO "abac_policy" ("name", "slug", "description", "resource", "action", "rules", "effect", "priority") VALUES
('Admin Office Hours', 'admin_office_hours', 'Admin access from office IP only after hours', 'system', 'manage', '{
    "subject": {"role": "admin"},
    "environment": {
        "time": {"start": "18:00", "end": "09:00"},
        "ip_range": ["10.0.0.0/8", "172.16.0.0/12"]
    }
}', 'allow', 50);
-- +goose Down
-- ============================================
-- DOWN MIGRATION WITH TRANSACTION
-- ============================================

BEGIN;

-- Disable foreign key checks temporarily (PostgreSQL handles this differently)
-- In PostgreSQL, foreign key constraints are checked at transaction commit
-- So we can drop in any order within a transaction, but it's cleaner to drop child tables first

-- Drop all tables in reverse dependency order
DROP TABLE IF EXISTS "permission_change_log";
DROP TABLE IF EXISTS "permission_template";
DROP TABLE IF EXISTS "api_key";
DROP TABLE IF EXISTS "security_event";
DROP TABLE IF EXISTS "audit_log";
DROP TABLE IF EXISTS "user_attribute";
DROP TABLE IF EXISTS "abac_policy";
DROP TABLE IF EXISTS "user_role";
DROP TABLE IF EXISTS "role_permission";
DROP TABLE IF EXISTS "permission";
DROP TABLE IF EXISTS "role";
DROP TABLE IF EXISTS "password_reset";
DROP TABLE IF EXISTS "verification";
DROP TABLE IF EXISTS "account";
DROP TABLE IF EXISTS "user";

-- Drop all sequences if they still exist
DO $$
DECLARE
    seq_name TEXT;
BEGIN
    FOR seq_name IN 
        SELECT sequence_name 
        FROM information_schema.sequences 
        WHERE sequence_schema = 'public'
        AND sequence_name LIKE '%_id_seq'
    LOOP
        EXECUTE 'DROP SEQUENCE IF EXISTS ' || seq_name;
    END LOOP;
END $$;

-- Optionally drop any remaining indexes (though they're dropped with tables)
-- This is just for cleanup if needed

COMMIT;

-- ============================================
-- Verification: Check that everything is dropped
-- ============================================

-- Uncomment to verify:
-- SELECT table_name FROM information_schema.tables 
-- WHERE table_schema = 'public' 
-- AND table_name IN ('user', 'account', 'verification', 'password_reset', 'role', 'permission', 'role_permission', 'user_role', 'abac_policy', 'user_attribute', 'audit_log', 'security_event', 'api_key', 'permission_template', 'permission_change_log');
