-- name: InsertRoleWithIDForTest :exec
INSERT INTO roles (id, name, priority, created_at, updated_at)
VALUES (?, ?, ?, ?, ?);

-- name: InsertUserRoleForTest :exec
INSERT INTO users_role (user_id, role_id, updated_at)
VALUES (?, ?, ?)
