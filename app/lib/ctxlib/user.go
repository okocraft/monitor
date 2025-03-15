package ctxlib

import (
	"context"

	"github.com/okocraft/monitor/internal/domain/user"
)

func WithUserID(ctx context.Context, userID user.ID) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

func GetUserID(ctx context.Context) (user.ID, bool) {
	userID, ok := ctx.Value(userIDKey).(user.ID)
	if !ok {
		return 0, false
	}
	return userID, true
}
