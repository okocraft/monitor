package auditlog

import (
	"context"
	"github.com/huandu/go-sqlbuilder"
	"github.com/okocraft/monitor/internal/domain/auditlog"
	"github.com/okocraft/monitor/internal/repositories/database"
	"github.com/okocraft/monitor/internal/repositories/queries"
	"iter"
	"time"
)

type AuditLogRepository interface {
	RecordOperator(ctx context.Context, operator auditlog.Operator) (auditlog.OperatorID, error)
	RecordUserAuditLog(ctx context.Context, operatorID auditlog.OperatorID, itr iter.Seq[auditlog.UserLogRecord]) error
}

type auditLogRepository struct {
	db database.DB
}

func NewAuditLogRepository(db database.DB) AuditLogRepository {
	return auditLogRepository{
		db: db,
	}
}

func (r auditLogRepository) RecordOperator(ctx context.Context, operator auditlog.Operator) (auditlog.OperatorID, error) {
	q := r.db.Queries(ctx)
	result, err := q.InsertOperator(ctx, queries.InsertOperatorParams{
		UserID:    int32(operator.UserID),
		Name:      operator.Name,
		Ip:        operator.IP,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return 0, database.NewDBErrorWithStackTrace(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, database.NewDBErrorWithStackTrace(err)
	}
	return auditlog.OperatorID(id), nil
}

func (r auditLogRepository) RecordUserAuditLog(ctx context.Context, operatorID auditlog.OperatorID, itr iter.Seq[auditlog.UserLogRecord]) error {
	ib := sqlbuilder.NewInsertBuilder()
	ib.InsertInto("audit_log_user")
	ib.Cols("operator_id", "action_type", "changed_from", "changed_to", "created_at")

	for log := range itr {
		ib.Values(operatorID, log.Action, log.ChangedFrom, log.ChangedTo, log.GetTimestamp())
	}

	sql, args := ib.Build()
	conn := r.db.Conn(ctx)
	if _, err := conn.ExecContext(ctx, sql, args...); err != nil {
		return database.NewDBErrorWithStackTrace(err)
	}
	return nil
}
