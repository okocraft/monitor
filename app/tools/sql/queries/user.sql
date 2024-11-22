-- name: GetUserByID :one
SELECT uuid, nickname, last_access
FROM users
WHERE id = ?
LIMIT 1;
