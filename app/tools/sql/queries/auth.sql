-- name: InsertRefreshToken :exec
INSERT INTO users_refresh_tokens (user_id, jti, ip, user_agent, created_at)
VALUES (?, ?, ?, ?, ?);

-- name: InsertAccessToken :exec
INSERT INTO users_access_tokens (user_id, refresh_token_id, jti, created_at)
VALUES (?, ?, ?, ?);

-- name: GetUserIDAndRefreshTokenIDByJTI :one
SELECT id, user_id FROM users_refresh_tokens
WHERE jti = ?;

-- name: DeleteRefreshTokenAndAccessToken :exec
DELETE FROM users_refresh_tokens
WHERE id = ?;

-- name: GetUserIDByAccessTokenJTI :one
SELECT user_id FROM users_access_tokens WHERE jti = ?;

-- name: DeleteExpiredAccessTokens :execrows
DELETE
FROM users_access_tokens
WHERE created_at < ?;

-- name: DeleteAccessTokensByExpiredRefreshTokens :execrows
DELETE
FROM users_access_tokens
WHERE refresh_token_id IN (SELECT users_refresh_tokens.created_at
                           FROM users_refresh_tokens
                           WHERE users_refresh_tokens.created_at < ?);

-- name: DeleteExpiredRefreshTokens :execrows
DELETE
FROM users_refresh_tokens
WHERE created_at < ?;
