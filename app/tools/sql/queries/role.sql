-- name: ExistsRoleByID :one
SELECT EXISTS(
    SELECT id FROM roles WHERE id = ?
);

-- name: CreateRoleWithIDIfNotExists :exec
INSERT IGNORE INTO roles (id, uuid, name, priority, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?);
