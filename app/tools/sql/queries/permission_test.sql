-- name: InsertPermissionForTest :exec
INSERT INTO roles_permissions (role_id, permission_id, is_allowed)
VALUES (?, ?, ?);
