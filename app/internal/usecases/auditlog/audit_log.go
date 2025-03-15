package auditlog

import (
	"context"

	"github.com/Siroshun09/serrors"
	"github.com/okocraft/monitor/internal/domain/auditlog"
	auditlog2 "github.com/okocraft/monitor/internal/repositories/auditlog"
	"github.com/okocraft/monitor/lib/errlib"
)

type AuditLogUsecase interface {
	Record(ctx context.Context, operator auditlog.Operator, records auditlog.AuditLogRecords) error
}

func NewAuditLogUsecase(auditLogRepo auditlog2.AuditLogRepository) AuditLogUsecase {
	return auditLogUsecase{
		auditLogRepo: auditLogRepo,
	}
}

type auditLogUsecase struct {
	auditLogRepo auditlog2.AuditLogRepository
}

func (u auditLogUsecase) Record(ctx context.Context, operator auditlog.Operator, records auditlog.AuditLogRecords) error {
	operatorID, err := u.auditLogRepo.RecordOperator(ctx, operator)
	if err != nil {
		return errlib.AsIs(err)
	}

	byType := records.KeyByType()

	for logType, rs := range byType {
		var err error
		switch logType {
		case auditlog.AuditLogTypeUser:
			err = u.auditLogRepo.RecordUserAuditLog(ctx, operatorID, auditlog.ToTypedIter[auditlog.UserLogRecord](rs))
		default:
			return serrors.Errorf("unknown audit log type: %d", logType)
		}
		if err != nil {
			return errlib.AsIs(err)
		}
	}

	return nil
}
