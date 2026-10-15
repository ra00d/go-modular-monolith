-- name: CreatePermission :one
INSERT INTO permission (
    created_at,
    updated_at,
    name,
    description,
    role_id
)
VALUES (
    NOW(),
    NOW(),
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetPermissionByID :one
SELECT *
FROM permission
WHERE id = $1
  AND deleted_at IS NULL;

-- name: ListPermissions :many
SELECT *
FROM permission
WHERE deleted_at IS NULL
ORDER BY id;

-- name: ListPermissionsByRole :many
SELECT *
FROM permission
WHERE role_id = $1
  AND deleted_at IS NULL
ORDER BY id;

-- name: UpdatePermission :one
UPDATE permission
SET
    updated_at = NOW(),
    name = $2,
    description = $3,
    role_id = $4
WHERE id = $1
RETURNING *;

-- name: DeletePermission :exec
UPDATE permission
SET
    deleted_at = NOW(),
    updated_at = NOW()
WHERE id = $1;
