-- name: InsertOperator :execresult
INSERT INTO audit_log_operators (user_id, name, ip, created_at)
VALUES (?, ?, ?, ?);
