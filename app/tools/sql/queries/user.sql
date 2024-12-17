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

-- name: InsertUser :exec
INSERT INTO users (uuid, nickname, last_access, created_at, updated_at)
VALUES (?, ?, ?, ?, ?);

-- name: InsertUserWithID :exec
INSERT INTO users (id, uuid, nickname, last_access, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?);

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
