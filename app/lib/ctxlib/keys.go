package ctxlib

type contextKey int8

const (
	accessLogKey contextKey = iota + 1
	auditLogKey
	userIDKey
)
