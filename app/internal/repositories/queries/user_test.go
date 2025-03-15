package queries

import (
	"testing"
	"time"

	"github.com/okocraft/monitor/internal/domain/user"
	"github.com/stretchr/testify/assert"
)

func TestSaveLoginKeyForUserID(t *testing.T) {
	tests := []struct {
		name      string
		userID    user.ID
		loginKey  int64
		now       time.Time
		wantQuery string
		wantArgs  []any
	}{
		{
			name:      "success",
			userID:    1,
			loginKey:  1234,
			now:       time.Date(2025, 3, 15, 0, 0, 0, 0, time.UTC),
			wantQuery: "INSERT INTO users_login_key VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE login_key = VALUES(login_key), created_at = VALUES(created_at)",
			wantArgs:  []any{1, int64(1234), time.Date(2025, 3, 15, 0, 0, 0, 0, time.UTC)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotQuery, gotArgs := SaveLoginKeyForUserID(tt.userID, tt.loginKey, tt.now)
			assert.Equal(t, tt.wantQuery, gotQuery)
			assert.Equal(t, tt.wantArgs, gotArgs)
		})
	}
}
