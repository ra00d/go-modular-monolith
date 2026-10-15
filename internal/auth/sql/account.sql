-- name: CreateAccount :one
INSERT INTO account (
    created_at,
    updated_at,
    user_id,
    account_id,
    provider_id,
    access_token,
    refresh_token,
    access_token_expires_at,
    refresh_token_expires_at,
    scope,
    id_token,
    password
)
VALUES (
    NOW(),
    NOW(),
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10
)
RETURNING *;

-- name: GetAccountByID :one
SELECT *
FROM account
WHERE id = $1
  AND deleted_at IS NULL;

-- name: GetAccountByProvider :one
SELECT *
FROM account
WHERE provider_id = $1
  AND account_id = $2
  AND deleted_at IS NULL;

-- name: ListAccountsByUser :many
SELECT *
FROM account
WHERE user_id = $1
  AND deleted_at IS NULL;

-- name: DeleteAccount :exec
UPDATE account
SET
    deleted_at = NOW(),
    updated_at = NOW()
WHERE id = $1;
