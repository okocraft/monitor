-- name: InsertUserWithIDForTest :exec
INSERT INTO users (id, uuid, nickname, last_access, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?);
