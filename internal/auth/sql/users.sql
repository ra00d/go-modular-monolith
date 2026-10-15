-- name: CreateUser :one
INSERT INTO "user" (
    created_at,
    updated_at,
    name,
    email,
    email_verified,
    image
)
VALUES (
    NOW(),
    NOW(),
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUserByID :one
SELECT *
FROM "user"
WHERE id = $1
  AND deleted_at IS NULL;

-- name: GetUserByEmail :one
SELECT *
FROM "user"
WHERE email = $1
  AND deleted_at IS NULL;

-- name: ListUsers :many
SELECT *
FROM "user"
WHERE deleted_at IS NULL
ORDER BY id;

-- name: UpdateUser :one
UPDATE "user"
SET
    updated_at = NOW(),
    name = $2,
    email = $3,
    email_verified = $4,
    image = $5
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
UPDATE "user"
SET
    deleted_at = NOW(),
    updated_at = NOW()
WHERE id = $1;

-- name: VerifyUserEmail :exec
UPDATE "user"
SET
    updated_at = NOW(),
    email_verified = TRUE
WHERE id = $1;
