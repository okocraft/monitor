package ctxlib

import (
	"context"

	"github.com/okocraft/monitor/internal/domain/auditlog"
	"github.com/okocraft/monitor/internal/domain/user"
)

type AuditLogHolder struct {
	userID user.ID
	logs   auditlog.AuditLogRecords
}

func InitAuditLogHolder(ctx context.Context) (context.Context, *AuditLogHolder) {
	a := AuditLogHolder{}
	ctx = context.WithValue(ctx, auditLogKey, &a)
	return ctx, &a
}

func GetAuditLogHolder(ctx context.Context) *AuditLogHolder {
	a, ok := ctx.Value(auditLogKey).(*AuditLogHolder)
	if ok {
		return a
	}
	return &AuditLogHolder{}
}

func AddAuditLogRecord(ctx context.Context, record auditlog.AuditLogRecord) {
	a := GetAuditLogHolder(ctx)

	if a.logs == nil {
		a.logs = auditlog.AuditLogRecords{record}
	} else {
		a.logs = append(a.logs, record)
	}
}

func SetUserIDForAuditLog(ctx context.Context, userID user.ID) {
	a := GetAuditLogHolder(ctx)
	a.userID = userID
}

func (a *AuditLogHolder) GetUserID() user.ID {
	return a.userID
}

func (a *AuditLogHolder) GetRecords() auditlog.AuditLogRecords {
	return a.logs
}
