-- name: CheckPermissionByUserID :one
SELECT is_allowed
FROM roles_permissions
         INNER JOIN users_role ON user_id = ?
WHERE roles_permissions.role_id = users_role.role_id
  AND permission_id = ?;

-- name: GetPermissionsByUserID :many
SELECT permission_id, is_allowed
FROM roles_permissions
         INNER JOIN users_role ON user_id = ?
WHERE roles_permissions.role_id = users_role.role_id
  AND permission_id IN (sqlc.slice('permission_ids'));
