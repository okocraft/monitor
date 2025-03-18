package permission

import (
	"context"
	"testing"
	"time"

	"github.com/Siroshun09/testrecords"
	"github.com/okocraft/monitor/internal/domain/permission"
	"github.com/okocraft/monitor/internal/domain/user"
	"github.com/okocraft/monitor/internal/repositories/database"
	"github.com/okocraft/monitor/internal/repositories/queries"
	"github.com/okocraft/monitor/internal/repositories/records"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_permissionRepository_GetUserPermissions(t *testing.T) {
	newInserter := func(user queries.User, role queries.Role, perms ...queries.RolesPermission) testrecords.Inserter {
		return testrecords.NewInserterForMySQL().
			Add(queries.UsersTable.TableName, user).
			Add(queries.RolesTable.TableName, role).
			Add(queries.UsersRoleTable.TableName, queries.UsersRole{
				UserID:    user.ID,
				RoleID:    role.ID,
				UpdatedAt: time.Date(2025, 3, 18, 0, 0, 0, 0, time.UTC),
			}).
			Add(queries.RolesPermissionsTable.TableName, queries.ToAnySlice(perms)...)
	}
	tests := []struct {
		name     string
		inserter testrecords.Inserter
		userID   user.ID
		perms    []permission.Permission
		want     permission.ValueMapSource
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			name:   "ok: no records inserted (default allowed)",
			userID: user.ID(records.User1.ID),
			perms:  []permission.Permission{records.TestPermissionAllowed},
			want: permission.ValueMapSource{
				records.TestPermissionAllowed.ID: true,
			},
			wantErr: assert.NoError,
		},
		{
			name:   "ok: no records inserted (default not allowed)",
			userID: user.ID(records.User1.ID),
			perms:  []permission.Permission{records.TestPermissionNotAllowed},
			want: permission.ValueMapSource{
				records.TestPermissionNotAllowed.ID: false,
			},
			wantErr: assert.NoError,
		},
		{
			name: "ok: single permission true (default not allowed)",
			inserter: newInserter(
				records.User1, records.Role1,
				queries.RolesPermission{
					RoleID:       records.Role1.ID,
					PermissionID: records.TestPermissionNotAllowed.ID,
					IsAllowed:    true,
				},
			),
			userID: user.ID(records.User1.ID),
			perms:  []permission.Permission{records.TestPermissionNotAllowed},
			want: permission.ValueMapSource{
				records.TestPermissionNotAllowed.ID: true,
			},
			wantErr: assert.NoError,
		},
		{
			name: "ok: single permission false (default allowed)",
			inserter: newInserter(
				records.User1, records.Role1,
				queries.RolesPermission{
					RoleID:       records.Role1.ID,
					PermissionID: records.TestPermissionAllowed.ID,
					IsAllowed:    false,
				},
			),
			userID: user.ID(records.User1.ID),
			perms:  []permission.Permission{records.TestPermissionAllowed},
			want: permission.ValueMapSource{
				records.TestPermissionAllowed.ID: false,
			},
			wantErr: assert.NoError,
		},
		{
			name: "ok: multiple permissions (all inserted)",
			inserter: newInserter(
				records.User1, records.Role1,
				queries.RolesPermission{
					RoleID:       records.Role1.ID,
					PermissionID: records.TestPermissionAllowed.ID,
					IsAllowed:    false,
				},
				queries.RolesPermission{
					RoleID:       records.Role1.ID,
					PermissionID: records.TestPermissionNotAllowed.ID,
					IsAllowed:    true,
				},
			),
			userID: user.ID(records.User1.ID),
			perms:  []permission.Permission{records.TestPermissionAllowed, records.TestPermissionNotAllowed},
			want: permission.ValueMapSource{
				records.TestPermissionAllowed.ID:    false,
				records.TestPermissionNotAllowed.ID: true,
			},
			wantErr: assert.NoError,
		},
		{
			name: "ok: multiple permissions (not all inserted: default true)",
			inserter: newInserter(
				records.User1, records.Role1,
				queries.RolesPermission{
					RoleID:       records.Role1.ID,
					PermissionID: records.TestPermissionNotAllowed.ID,
					IsAllowed:    true,
				},
			),
			userID: user.ID(records.User1.ID),
			perms:  []permission.Permission{records.TestPermissionAllowed, records.TestPermissionNotAllowed},
			want: permission.ValueMapSource{
				records.TestPermissionAllowed.ID:    true,
				records.TestPermissionNotAllowed.ID: true,
			},
			wantErr: assert.NoError,
		},
		{
			name: "ok: multiple permissions (not all inserted: default false)",
			inserter: newInserter(
				records.User1, records.Role1,
				queries.RolesPermission{
					RoleID:       records.Role1.ID,
					PermissionID: records.TestPermissionAllowed.ID,
					IsAllowed:    false,
				},
			),
			userID: user.ID(records.User1.ID),
			perms:  []permission.Permission{records.TestPermissionAllowed, records.TestPermissionNotAllowed},
			want: permission.ValueMapSource{
				records.TestPermissionAllowed.ID:    false,
				records.TestPermissionNotAllowed.ID: false,
			},
			wantErr: assert.NoError,
		},
		{
			name: "ok: multiple permissions (all inserted, is admin)",
			inserter: newInserter(
				records.User1, records.Role1,
				queries.RolesPermission{
					RoleID:       records.Role1.ID,
					PermissionID: permission.Admin.ID,
					IsAllowed:    true,
				},
				queries.RolesPermission{
					RoleID:       records.Role1.ID,
					PermissionID: records.TestPermissionAllowed.ID,
					IsAllowed:    false,
				},
				queries.RolesPermission{
					RoleID:       records.Role1.ID,
					PermissionID: records.TestPermissionNotAllowed.ID,
					IsAllowed:    true,
				},
			),
			userID: user.ID(records.User1.ID),
			perms:  []permission.Permission{records.TestPermissionAllowed, records.TestPermissionNotAllowed},
			want: permission.ValueMapSource{
				permission.Admin.ID:                 true,
				records.TestPermissionAllowed.ID:    false,
				records.TestPermissionNotAllowed.ID: true,
			},
			wantErr: assert.NoError,
		},
		{
			name: "ok: multiple permissions (all inserted, is not admin)",
			inserter: newInserter(
				records.User1, records.Role1,
				queries.RolesPermission{
					RoleID:       records.Role1.ID,
					PermissionID: permission.Admin.ID,
					IsAllowed:    false,
				},
				queries.RolesPermission{
					RoleID:       records.Role1.ID,
					PermissionID: records.TestPermissionAllowed.ID,
					IsAllowed:    false,
				},
				queries.RolesPermission{
					RoleID:       records.Role1.ID,
					PermissionID: records.TestPermissionNotAllowed.ID,
					IsAllowed:    true,
				},
			),
			userID: user.ID(records.User1.ID),
			perms:  []permission.Permission{records.TestPermissionAllowed, records.TestPermissionNotAllowed},
			want: permission.ValueMapSource{
				permission.Admin.ID:                 false,
				records.TestPermissionAllowed.ID:    false,
				records.TestPermissionNotAllowed.ID: true,
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testDB.Run(t, func(ctx context.Context, db database.DB) {
				require.NoError(t, tt.inserter.InsertAll(ctx, db.Conn(ctx)))

				r := permissionRepository{db: db}

				got, err := r.GetUserPermissions(ctx, tt.userID, tt.perms...)
				if !tt.wantErr(t, err) {
					return
				}

				assert.Equal(t, permission.NewValueMap(tt.want), got)
			})
		})
	}
}
