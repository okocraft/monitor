package user

import (
	"context"
	"testing"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/okocraft/monitor/internal/domain/sort"
	"github.com/okocraft/monitor/internal/repositories/records"
	"github.com/okocraft/monitor/lib/null"

	"github.com/okocraft/monitor/internal/domain/user"
	"github.com/okocraft/monitor/internal/repositories/database"
	"github.com/okocraft/monitor/internal/repositories/queries"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_userRepository_GetUserByID(t *testing.T) {
	tests := []struct {
		name          string
		initialRecord *queries.User
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
					require.NoError(t, records.NewInitialRecords().Table(queries.UsersTable.TableName, *tt.initialRecord).InsertAll(ctx, db))
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

func Test_userRepository_SearchForUserUUIDs(t *testing.T) {
	tests := []struct {
		name         string
		params       user.SearchParams
		initialTable *records.InitialRecords
		want         []uuid.UUID
		wantErr      assert.ErrorAssertionFunc
	}{
		{
			name:    "success: no records",
			params:  user.SearchParams{},
			want:    []uuid.UUID{},
			wantErr: assert.NoError,
		},
		{
			name:         "success: one record",
			params:       user.SearchParams{},
			initialTable: records.NewInitialRecords().Table(queries.UsersTable.TableName, records.User1),
			want:         []uuid.UUID{uuid.UUID(records.User1.Uuid)},
			wantErr:      assert.NoError,
		},
		{
			name: "success: one record (filter by NickName)",
			params: user.SearchParams{
				Nickname: null.FromValue("Test"),
			},
			initialTable: records.NewInitialRecords().Table(queries.UsersTable.TableName, records.User1),
			want:         []uuid.UUID{uuid.UUID(records.User1.Uuid)},
			wantErr:      assert.NoError,
		},
		{
			name: "success: no record (filter by NickName)",
			params: user.SearchParams{
				Nickname: null.FromValue("not match"),
			},
			initialTable: records.NewInitialRecords().Table(queries.UsersTable.TableName, records.User1),
			want:         []uuid.UUID{},
			wantErr:      assert.NoError,
		},
		{
			name: "success: one record (filter by LastAccessBefore)",
			params: user.SearchParams{
				LastAccessBefore: null.FromValue(records.User1.LastAccess.Add(1 * time.Second)),
			},
			initialTable: records.NewInitialRecords().Table(queries.UsersTable.TableName, records.User1),
			want:         []uuid.UUID{uuid.UUID(records.User1.Uuid)},
			wantErr:      assert.NoError,
		},
		{
			name: "success: one record (filter by LastAccessBefore, same time)",
			params: user.SearchParams{
				LastAccessBefore: null.FromValue(records.User1.LastAccess),
			},
			initialTable: records.NewInitialRecords().Table(queries.UsersTable.TableName, records.User1),
			want:         []uuid.UUID{uuid.UUID(records.User1.Uuid)},
			wantErr:      assert.NoError,
		},
		{
			name: "success: no record (filter by LastAccessBefore)",
			params: user.SearchParams{
				LastAccessBefore: null.FromValue(records.User1.LastAccess.Add(-1 * time.Second)),
			},
			initialTable: records.NewInitialRecords().Table(queries.UsersTable.TableName, records.User1),
			want:         []uuid.UUID{},
			wantErr:      assert.NoError,
		},
		{
			name: "success: one record (filter by LastAccessAfter)",
			params: user.SearchParams{
				LastAccessAfter: null.FromValue(records.User1.LastAccess.Add(-1 * time.Second)),
			},
			initialTable: records.NewInitialRecords().Table(queries.UsersTable.TableName, records.User1),
			want:         []uuid.UUID{uuid.UUID(records.User1.Uuid)},
			wantErr:      assert.NoError,
		},
		{
			name: "success: one record (filter by LastAccessAfter, same time)",
			params: user.SearchParams{
				LastAccessAfter: null.FromValue(records.User1.LastAccess),
			},
			initialTable: records.NewInitialRecords().Table(queries.UsersTable.TableName, records.User1),
			want:         []uuid.UUID{uuid.UUID(records.User1.Uuid)},
			wantErr:      assert.NoError,
		},
		{
			name: "success: no record (filter by LastAccessAfter)",
			params: user.SearchParams{
				LastAccessAfter: null.FromValue(records.User1.LastAccess.Add(1 * time.Second)),
			},
			initialTable: records.NewInitialRecords().Table(queries.UsersTable.TableName, records.User1),
			want:         []uuid.UUID{},
			wantErr:      assert.NoError,
		},
		{
			name: "success: one record (filter by RoleID)",
			params: user.SearchParams{
				RoleId: null.FromValue(uuid.UUID(records.Role1.Uuid)),
			},
			initialTable: records.NewInitialRecords().
				Table(queries.UsersTable.TableName, records.User1).
				Table(queries.RolesTable.TableName, records.Role1).
				Table(queries.UsersRoleTable.TableName, records.UserRole1),
			want:    []uuid.UUID{uuid.UUID(records.User1.Uuid)},
			wantErr: assert.NoError,
		},
		{
			name: "success: no record (filter by RoleID, user does not have role)",
			params: user.SearchParams{
				RoleId: null.FromValue(uuid.UUID(records.Role1.Uuid)),
			},
			initialTable: records.NewInitialRecords().Table(queries.UsersTable.TableName, records.User1),
			want:         []uuid.UUID{},
			wantErr:      assert.NoError,
		},
		{
			name: "success: no record (filter by RoleID, different role)",
			params: user.SearchParams{
				RoleId: null.FromValue(uuid.Nil),
			},
			initialTable: records.NewInitialRecords().
				Table(queries.UsersTable.TableName, records.User1).
				Table(queries.RolesTable.TableName, records.Role1).
				Table(queries.UsersRoleTable.TableName, records.UserRole1),
			want:    []uuid.UUID{},
			wantErr: assert.NoError,
		},
		{
			name: "success: 3 records (sorted by NickName, not specify sort type)",
			params: user.SearchParams{
				SortedBy: null.FromValue(user.SortableDataTypeNickName),
			},
			initialTable: records.NewInitialRecords().Table(queries.UsersTable.TableName, records.User1, records.User2, records.User3),
			want:         []uuid.UUID{uuid.UUID(records.User1.Uuid), uuid.UUID(records.User2.Uuid), uuid.UUID(records.User3.Uuid)},
			wantErr:      assert.NoError,
		},
		{
			name: "success: 3 records (sorted by NickName, asc)",
			params: user.SearchParams{
				SortedBy: null.FromValue(user.SortableDataTypeNickName),
				SortType: null.FromValue(sort.ASC),
			},
			initialTable: records.NewInitialRecords().Table(queries.UsersTable.TableName, records.User1, records.User2, records.User3),
			want:         []uuid.UUID{uuid.UUID(records.User1.Uuid), uuid.UUID(records.User2.Uuid), uuid.UUID(records.User3.Uuid)},
			wantErr:      assert.NoError,
		},
		{
			name: "success: 3 records (sorted by NickName, desc)",
			params: user.SearchParams{
				SortedBy: null.FromValue(user.SortableDataTypeNickName),
				SortType: null.FromValue(sort.DESC),
			},
			initialTable: records.NewInitialRecords().Table(queries.UsersTable.TableName, records.User1, records.User2, records.User3),
			want:         []uuid.UUID{uuid.UUID(records.User3.Uuid), uuid.UUID(records.User2.Uuid), uuid.UUID(records.User1.Uuid)},
			wantErr:      assert.NoError,
		},
		{
			name: "success: 3 records (sorted by LastAccess, asc)",
			params: user.SearchParams{
				SortedBy: null.FromValue(user.SortableDataTypeLastAccess),
				SortType: null.FromValue(sort.ASC),
			},
			initialTable: records.NewInitialRecords().Table(queries.UsersTable.TableName, records.User1, records.User2, records.User3),
			want:         []uuid.UUID{uuid.UUID(records.User1.Uuid), uuid.UUID(records.User2.Uuid), uuid.UUID(records.User3.Uuid)},
			wantErr:      assert.NoError,
		},
		{
			name: "success: 3 records (sorted by LastAccess, desc)",
			params: user.SearchParams{
				SortedBy: null.FromValue(user.SortableDataTypeLastAccess),
				SortType: null.FromValue(sort.DESC),
			},
			initialTable: records.NewInitialRecords().Table(queries.UsersTable.TableName, records.User1, records.User2, records.User3),
			want:         []uuid.UUID{uuid.UUID(records.User3.Uuid), uuid.UUID(records.User2.Uuid), uuid.UUID(records.User1.Uuid)},
			wantErr:      assert.NoError,
		},
		{
			name: "success: 3 records (sorted by CreatedAt, asc)",
			params: user.SearchParams{
				SortedBy: null.FromValue(user.SortableDataTypeCreatedAt),
				SortType: null.FromValue(sort.ASC),
			},
			initialTable: records.NewInitialRecords().Table(queries.UsersTable.TableName, records.User1, records.User2, records.User3),
			want:         []uuid.UUID{uuid.UUID(records.User1.Uuid), uuid.UUID(records.User2.Uuid), uuid.UUID(records.User3.Uuid)},
			wantErr:      assert.NoError,
		},
		{
			name: "success: 3 records (sorted by CreatedAt, desc)",
			params: user.SearchParams{
				SortedBy: null.FromValue(user.SortableDataTypeCreatedAt),
				SortType: null.FromValue(sort.DESC),
			},
			initialTable: records.NewInitialRecords().Table(queries.UsersTable.TableName, records.User1, records.User2, records.User3),
			want:         []uuid.UUID{uuid.UUID(records.User3.Uuid), uuid.UUID(records.User2.Uuid), uuid.UUID(records.User1.Uuid)},
			wantErr:      assert.NoError,
		},
		{
			name: "success: 3 records (sorted by UpdatedAt, asc)",
			params: user.SearchParams{
				SortedBy: null.FromValue(user.SortableDataTypeUpdatedAt),
				SortType: null.FromValue(sort.ASC),
			},
			initialTable: records.NewInitialRecords().Table(queries.UsersTable.TableName, records.User1, records.User2, records.User3),
			want:         []uuid.UUID{uuid.UUID(records.User1.Uuid), uuid.UUID(records.User2.Uuid), uuid.UUID(records.User3.Uuid)},
			wantErr:      assert.NoError,
		},
		{
			name: "success: 3 records (sorted by UpdatedAt, desc)",
			params: user.SearchParams{
				SortedBy: null.FromValue(user.SortableDataTypeUpdatedAt),
				SortType: null.FromValue(sort.DESC),
			},
			initialTable: records.NewInitialRecords().Table(queries.UsersTable.TableName, records.User1, records.User2, records.User3),
			want:         []uuid.UUID{uuid.UUID(records.User3.Uuid), uuid.UUID(records.User2.Uuid), uuid.UUID(records.User1.Uuid)},
			wantErr:      assert.NoError,
		},
		{
			name: "success: 3 records (sorted by RoleName, asc)",
			params: user.SearchParams{
				SortedBy: null.FromValue(user.SortableDataTypeRoleName),
				SortType: null.FromValue(sort.ASC),
			},
			initialTable: records.NewInitialRecords().
				Table(queries.UsersTable.TableName, records.User1, records.User2, records.User3).
				Table(queries.RolesTable.TableName, records.Role1, records.Role2, records.Role3).
				Table(queries.UsersRoleTable.TableName, records.UserRole1, records.UserRole2, records.UserRole3),
			want:    []uuid.UUID{uuid.UUID(records.User1.Uuid), uuid.UUID(records.User2.Uuid), uuid.UUID(records.User3.Uuid)},
			wantErr: assert.NoError,
		},
		{
			name: "success: 3 records (sorted by RoleName, desc)",
			params: user.SearchParams{
				SortedBy: null.FromValue(user.SortableDataTypeRoleName),
				SortType: null.FromValue(sort.DESC),
			},
			initialTable: records.NewInitialRecords().
				Table(queries.UsersTable.TableName, records.User1, records.User2, records.User3).
				Table(queries.RolesTable.TableName, records.Role1, records.Role2, records.Role3).
				Table(queries.UsersRoleTable.TableName, records.UserRole1, records.UserRole2, records.UserRole3),
			want:    []uuid.UUID{uuid.UUID(records.User3.Uuid), uuid.UUID(records.User2.Uuid), uuid.UUID(records.User1.Uuid)},
			wantErr: assert.NoError,
		},
		{
			name: "success: 3 records (sorted by RoleName, someone has no role)",
			params: user.SearchParams{
				SortedBy: null.FromValue(user.SortableDataTypeRoleName),
				SortType: null.FromValue(sort.ASC),
			},
			initialTable: records.NewInitialRecords().
				Table(queries.UsersTable.TableName, records.User1, records.User2, records.User3).
				Table(queries.RolesTable.TableName, records.Role1, records.Role2).
				Table(queries.UsersRoleTable.TableName, records.UserRole1, records.UserRole2),
			want:    []uuid.UUID{uuid.UUID(records.User3.Uuid), uuid.UUID(records.User1.Uuid), uuid.UUID(records.User2.Uuid)},
			wantErr: assert.NoError,
		},
		{
			name: "success: 3 records (sorted by RolePriority, asc)",
			params: user.SearchParams{
				SortedBy: null.FromValue(user.SortableDataTypeRolePriority),
				SortType: null.FromValue(sort.ASC),
			},
			initialTable: records.NewInitialRecords().
				Table(queries.UsersTable.TableName, records.User1, records.User2, records.User3).
				Table(queries.RolesTable.TableName, records.Role1, records.Role2, records.Role3).
				Table(queries.UsersRoleTable.TableName, records.UserRole1, records.UserRole2, records.UserRole3),
			want:    []uuid.UUID{uuid.UUID(records.User1.Uuid), uuid.UUID(records.User2.Uuid), uuid.UUID(records.User3.Uuid)},
			wantErr: assert.NoError,
		},
		{
			name: "success: 3 records (sorted by RolePriority, desc)",
			params: user.SearchParams{
				SortedBy: null.FromValue(user.SortableDataTypeRolePriority),
				SortType: null.FromValue(sort.DESC),
			},
			initialTable: records.NewInitialRecords().
				Table(queries.UsersTable.TableName, records.User1, records.User2, records.User3).
				Table(queries.RolesTable.TableName, records.Role1, records.Role2, records.Role3).
				Table(queries.UsersRoleTable.TableName, records.UserRole1, records.UserRole2, records.UserRole3),
			want:    []uuid.UUID{uuid.UUID(records.User3.Uuid), uuid.UUID(records.User2.Uuid), uuid.UUID(records.User1.Uuid)},
			wantErr: assert.NoError,
		},
		{
			name: "success: 3 records (sorted by RolePriority, someone has no role)",
			params: user.SearchParams{
				SortedBy: null.FromValue(user.SortableDataTypeRolePriority),
				SortType: null.FromValue(sort.ASC),
			},
			initialTable: records.NewInitialRecords().
				Table(queries.UsersTable.TableName, records.User1, records.User2, records.User3).
				Table(queries.RolesTable.TableName, records.Role1, records.Role2).
				Table(queries.UsersRoleTable.TableName, records.UserRole1, records.UserRole2),
			want:    []uuid.UUID{uuid.UUID(records.User3.Uuid), uuid.UUID(records.User1.Uuid), uuid.UUID(records.User2.Uuid)},
			wantErr: assert.NoError,
		},
		{
			name: "success: 1 record (all filter)",
			params: user.SearchParams{
				Nickname:         null.FromValue("Test"),
				LastAccessBefore: null.FromValue(records.User1.LastAccess.Add(1 * time.Second)),
				LastAccessAfter:  null.FromValue(records.User1.LastAccess.Add(-1 * time.Second)),
				RoleId:           null.FromValue(uuid.UUID(records.Role1.Uuid)),
			},
			initialTable: records.NewInitialRecords().
				Table(queries.UsersTable.TableName, records.User1, records.User2, records.User3).
				Table(queries.RolesTable.TableName, records.Role1, records.Role2).
				Table(queries.UsersRoleTable.TableName, records.UserRole1, records.UserRole2),
			want:    []uuid.UUID{uuid.UUID(records.User1.Uuid)},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testDB.Run(t, func(ctx context.Context, db database.DB) {
				if tt.initialTable != nil {
					require.NoError(t, tt.initialTable.InsertAll(ctx, db))
				}

				r := userRepository{db: db}
				u, err := r.SearchForUserUUIDs(ctx, tt.params)
				tt.wantErr(t, err)
				assert.Equal(t, tt.want, u)
			})
		})
	}
}
