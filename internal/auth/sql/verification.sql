-- name: CreateVerification :one
INSERT INTO verification (
    created_at,
    updated_at,
    identifier,
    value,
    expires_at
)
VALUES (
    NOW(),
    NOW(),
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetVerification :one
SELECT *
FROM verification
WHERE identifier = $1
  AND value = $2
  AND deleted_at IS NULL;

-- name: DeleteVerification :exec
UPDATE verification
SET
    deleted_at = NOW(),
    updated_at = NOW()
WHERE id = $1;

-- name: DeleteExpiredVerifications :exec
DELETE
FROM verification
WHERE expires_at < NOW();
