package queries

import (
	"testing"
	"time"

	"github.com/okocraft/monitor/internal/domain/role"
	"github.com/okocraft/monitor/internal/domain/user"
	"github.com/stretchr/testify/assert"
)

func TestSaveUsersRole(t *testing.T) {
	tests := []struct {
		name      string
		userID    user.ID
		roleID    role.ID
		now       time.Time
		wantQuery string
		wantArgs  []any
	}{
		{
			name:      "success",
			userID:    1,
			roleID:    1,
			now:       time.Date(2025, 3, 15, 0, 0, 0, 0, time.UTC),
			wantQuery: "INSERT INTO users_role VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE role_id = VALUES(role_id), updated_at = VALUES(updated_at)",
			wantArgs:  []any{1, 1, time.Date(2025, 3, 15, 0, 0, 0, 0, time.UTC)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotQuery, gotArgs := SaveUsersRole(tt.userID, tt.roleID, tt.now)
			assert.Equal(t, tt.wantQuery, gotQuery)
			assert.Equal(t, tt.wantArgs, gotArgs)
		})
	}
}
