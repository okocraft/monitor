-- name: GetUserByID :one
SELECT uuid, nickname, last_access
FROM users
WHERE id = ?
LIMIT 1;

-- name: GetUserNicknameByID :one
SELECT nickname
FROM users
WHERE id = ?
LIMIT 1;

-- name: GetUserByUUID :one
SELECT id, nickname, last_access
FROM users
WHERE uuid = ?
LIMIT 1;

-- name: UpdateUserNicknameByID :exec
UPDATE users
SET nickname=?,
    updated_at=?
WHERE id = ?;

-- name: UpdateUserLastAccessByID :exec
UPDATE users
SET last_access=?,
    updated_at=?
WHERE id = ?;

-- name: GetUserIDBySub :one
SELECT user_id
FROM users_sub
WHERE sub = ?;

-- name: InsertSubForUserID :exec
INSERT INTO users_sub (user_id, sub)
VALUES (?, ?);

-- name: InsertLoginKey :exec
INSERT INTO users_login_key (user_id, login_key, created_at)
VALUES (?, ?, ?);

-- name: GetUserIDByLoginKey :one
SELECT user_id
FROM users_login_key
WHERE login_key = ?;

-- name: DeleteLoginKey :exec
DELETE
FROM users_login_key
WHERE user_id = ?;

-- name: GetUsersWithRoleByUUIDs :many
SELECT users.id AS user_id, users.uuid AS user_uuid, users.nickname AS user_nickname, users.last_access AS user_last_access, users.created_at AS user_created_at, users.updated_at AS user_updated_at,
       roles.id AS role_id, roles.uuid AS role_uuid, roles.name AS role_name, roles.priority AS role_priority, roles.created_at AS role_created_at, roles.updated_at AS role_updated_at
FROM users
LEFT OUTER JOIN users_role on users.id = users_role.user_id
LEFT OUTER JOIN roles on users_role.role_id = roles.id
WHERE users.uuid IN(sqlc.slice('uuids'));
