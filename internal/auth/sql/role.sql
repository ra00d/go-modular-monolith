-- name: CreateRole :one
INSERT INTO role (
    created_at,
    updated_at,
    name,
    description,
    parent_role_id
)
VALUES (
    NOW(),
    NOW(),
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetRoleByID :one
SELECT *
FROM role
WHERE id = $1
  AND deleted_at IS NULL;

-- name: GetRoleByName :one
SELECT *
FROM role
WHERE name = $1
  AND deleted_at IS NULL;

-- name: ListRoles :many
SELECT *
FROM role
WHERE deleted_at IS NULL
ORDER BY id;

-- name: UpdateRole :one
UPDATE role
SET
    updated_at = NOW(),
    name = $2,
    description = $3,
    parent_role_id = $4
WHERE id = $1
RETURNING *;

-- name: DeleteRole :exec
UPDATE role
SET
    deleted_at = NOW(),
    updated_at = NOW()
WHERE id = $1;
