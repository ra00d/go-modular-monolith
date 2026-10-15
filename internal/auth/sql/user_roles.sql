-- name: AssignRoleToUser :exec
INSERT INTO user_roles (
    user_id,
    role_id
)
VALUES (
    $1,
    $2
);

-- name: RemoveRoleFromUser :exec
DELETE
FROM user_roles
WHERE user_id = $1
  AND role_id = $2;

-- name: ListRolesForUser :many
SELECT r.*
FROM role r
JOIN user_roles ur
    ON ur.role_id = r.id
WHERE ur.user_id = $1
  AND r.deleted_at IS NULL
ORDER BY r.name;

-- name: ListUsersForRole :many
SELECT u.*
FROM "user" u
JOIN user_roles ur
    ON ur.user_id = u.id
WHERE ur.role_id = $1
  AND u.deleted_at IS NULL
ORDER BY u.id;

-- name: UserHasRole :one
SELECT EXISTS (
    SELECT 1
    FROM user_roles
    WHERE user_id = $1
      AND role_id = $2
);
