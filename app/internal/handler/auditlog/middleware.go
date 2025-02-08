package auditlog

import (
	"context"
	"errors"
	"github.com/Siroshun09/logs"
	"github.com/okocraft/monitor/internal/domain/auditlog"
	"github.com/okocraft/monitor/internal/domain/user"
	auditlog2 "github.com/okocraft/monitor/internal/usecases/auditlog"
	user2 "github.com/okocraft/monitor/internal/usecases/user"
	"github.com/okocraft/monitor/lib/ctxlib"
	"net/http"
)

type AuditLogMiddleware struct {
	auditLogUsecase auditlog2.AuditLogUsecase
	userUsecase     user2.UserUsecase
}

func NewAuditLogMiddleware(auditLogUsecase auditlog2.AuditLogUsecase, userUsecase user2.UserUsecase) AuditLogMiddleware {
	return AuditLogMiddleware{
		auditLogUsecase: auditLogUsecase,
		userUsecase:     userUsecase,
	}
}

func (m AuditLogMiddleware) NewHTTPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx, holder := ctxlib.InitAuditLogHolder(ctx)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)

		m.recordAuditLogs(ctx, holder)
	})
}

func (m AuditLogMiddleware) recordAuditLogs(ctx context.Context, holder *ctxlib.AuditLogHolder) {
	records := holder.GetRecords()
	if records == nil || len(records) == 0 {
		return
	}

	operator := auditlog.Operator{
		UserID: holder.GetUserID(),
		Name:   "",
		IP:     ctxlib.GetHTTPAccessLog(ctx).GetIP(),
	}

	if operator.UserID != 0 {
		name, err := m.userUsecase.GetNicknameByID(ctx, operator.UserID)
		if err != nil && !errors.As(err, &user.NotFoundByIDError{}) {
			logs.Error(ctx, err)
			return
		}
		operator.Name = name
	}

	err := m.auditLogUsecase.Record(ctx, operator, records)
	if err != nil {
		logs.Error(ctx, err)
	}
}
