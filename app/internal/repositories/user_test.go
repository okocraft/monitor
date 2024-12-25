package repositories

import (
	"context"
	"github.com/gofrs/uuid/v5"
	"github.com/okocraft/monitor/internal/repositories/records"
	"testing"

	"github.com/okocraft/monitor/internal/domain/user"
	"github.com/okocraft/monitor/internal/repositories/database"
	"github.com/okocraft/monitor/internal/repositories/queries"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_userRepository_GetUserByID(t *testing.T) {
	tests := []struct {
		name          string
		initialRecord *queries.InsertUserWithIDForTestParams
		id            user.ID
		want          user.User
		wantErr       bool
	}{
		{
			name:          "success",
			initialRecord: &records.User1,
			id:            1,
			want: user.User{
				ID:         1,
				UUID:       uuid.UUID(records.User1.Uuid),
				Nickname:   records.User1.Nickname,
				LastAccess: records.User1.LastAccess,
			},
		},
		{
			name:    "not found (no records)",
			id:      1,
			wantErr: true,
		},
		{
			name:          "not found (different id)",
			initialRecord: &records.User1,
			id:            0,
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testDB.Run(t, func(ctx context.Context, db database.DB) {
				if tt.initialRecord != nil {
					require.NoError(t, db.Queries(ctx).InsertUserWithIDForTest(ctx, *tt.initialRecord))
				}

				r := userRepository{db: db}
				u, err := r.GetUserByID(ctx, tt.id)
				if tt.wantErr {
					assert.EqualError(t, err, user.NotFoundByIDError{ID: tt.id}.Error())
				} else {
					require.NoError(t, err)
					assert.Equal(t, tt.want, u)
				}
			})
		})
	}
}
