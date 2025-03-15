package permission

import (
	"context"
	"testing"

	"github.com/Siroshun09/testrecords"
	"github.com/okocraft/monitor/internal/domain/permission"
	"github.com/okocraft/monitor/internal/domain/role"
	"github.com/okocraft/monitor/internal/domain/user"
	"github.com/okocraft/monitor/internal/repositories/database"
	"github.com/okocraft/monitor/internal/repositories/queries"
	"github.com/okocraft/monitor/internal/repositories/records"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_permissionRepository_HasPermission(t *testing.T) {
	type initial struct {
		userRecord     queries.User
		roleRecord     queries.Role
		userRoleRecord queries.UsersRole
		dbValue        bool
	}
	tests := []struct {
		name    string
		initial *initial
		userID  user.ID
		perm    permission.Permission
		want    bool
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "ok: true from db (default allowed)",
			initial: &initial{
				userRecord:     records.User1,
				roleRecord:     records.Role1,
				userRoleRecord: records.UserRole1,
				dbValue:        true,
			},
			userID:  user.ID(records.User1.ID),
			perm:    records.TestPermissionAllowed,
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name: "ok: true from db (default not allowed)",
			initial: &initial{
				userRecord:     records.User1,
				roleRecord:     records.Role1,
				userRoleRecord: records.UserRole1,
				dbValue:        true,
			},
			userID:  user.ID(records.User1.ID),
			perm:    records.TestPermissionNotAllowed,
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name: "ok: false from db (default allowed)",
			initial: &initial{
				userRecord:     records.User1,
				roleRecord:     records.Role1,
				userRoleRecord: records.UserRole1,
				dbValue:        false,
			},
			userID:  user.ID(records.User1.ID),
			perm:    records.TestPermissionAllowed,
			want:    false,
			wantErr: assert.NoError,
		},
		{
			name: "ok: false from db (default not allowed)",
			initial: &initial{
				userRecord:     records.User1,
				roleRecord:     records.Role1,
				userRoleRecord: records.UserRole1,
				dbValue:        false,
			},
			userID:  user.ID(records.User1.ID),
			perm:    records.TestPermissionNotAllowed,
			want:    false,
			wantErr: assert.NoError,
		},
		{
			name:    "ok: true (default allowed)",
			userID:  user.ID(records.User1.ID),
			perm:    records.TestPermissionAllowed,
			want:    true,
			wantErr: assert.NoError,
		},

		{
			name:    "ok: false (default not allowed)",
			userID:  user.ID(records.User1.ID),
			perm:    records.TestPermissionNotAllowed,
			want:    false,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testDB.Run(t, func(ctx context.Context, db database.DB) {
				if tt.initial != nil {
					require.NoError(t, testrecords.NewInserterForMySQL().
						Add(queries.UsersTable.TableName, tt.initial.userRecord).
						Add(queries.RolesTable.TableName, tt.initial.roleRecord).
						Add(queries.UsersRoleTable.TableName, tt.initial.userRoleRecord).
						Add(queries.RolesPermissionsTable.TableName, queries.RolesPermission{
							RoleID:       tt.initial.roleRecord.ID,
							PermissionID: tt.perm.ID,
							IsAllowed:    tt.initial.dbValue,
						}).InsertAll(ctx, db.Conn(ctx)))
				}

				r := permissionRepository{db: db}

				got, err := r.HasPermission(ctx, tt.userID, tt.perm)
				if !tt.wantErr(t, err) {
					return
				}

				assert.Equal(t, tt.want, got)
			})
		})
	}
}

func Test_permissionRepository_GetPermissions(t *testing.T) {
	type initial struct {
		userRecord     queries.User
		roleRecord     queries.Role
		userRoleRecord queries.UsersRole
		dbValueMap     permission.ValueMapSource
	}
	tests := []struct {
		name    string
		initial *initial
		userID  user.ID
		perms   []permission.Permission
		want    permission.ValueMapSource
		wantErr assert.ErrorAssertionFunc
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
			initial: &initial{
				userRecord:     records.User1,
				roleRecord:     records.Role1,
				userRoleRecord: records.UserRole1,
				dbValueMap: permission.ValueMapSource{
					records.TestPermissionNotAllowed.ID: true,
				},
			},
			userID: user.ID(records.User1.ID),
			perms:  []permission.Permission{records.TestPermissionNotAllowed},
			want: permission.ValueMapSource{
				records.TestPermissionNotAllowed.ID: true,
			},
			wantErr: assert.NoError,
		},
		{
			name: "ok: single permission false (default allowed)",
			initial: &initial{
				userRecord:     records.User1,
				roleRecord:     records.Role1,
				userRoleRecord: records.UserRole1,
				dbValueMap: permission.ValueMapSource{
					records.TestPermissionAllowed.ID: false,
				},
			},
			userID: user.ID(records.User1.ID),
			perms:  []permission.Permission{records.TestPermissionAllowed},
			want: permission.ValueMapSource{
				records.TestPermissionAllowed.ID: false,
			},
			wantErr: assert.NoError,
		},
		{
			name: "ok: multiple permissions (all inserted)",
			initial: &initial{
				userRecord:     records.User1,
				roleRecord:     records.Role1,
				userRoleRecord: records.UserRole1,
				dbValueMap: permission.ValueMapSource{
					records.TestPermissionAllowed.ID:    false,
					records.TestPermissionNotAllowed.ID: true,
				},
			},
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
			initial: &initial{
				userRecord:     records.User1,
				roleRecord:     records.Role1,
				userRoleRecord: records.UserRole1,
				dbValueMap: permission.ValueMapSource{
					records.TestPermissionNotAllowed.ID: true,
				},
			},
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
			initial: &initial{
				userRecord:     records.User1,
				roleRecord:     records.Role1,
				userRoleRecord: records.UserRole1,
				dbValueMap: permission.ValueMapSource{
					records.TestPermissionAllowed.ID: false,
				},
			},
			userID: user.ID(records.User1.ID),
			perms:  []permission.Permission{records.TestPermissionAllowed, records.TestPermissionNotAllowed},
			want: permission.ValueMapSource{
				records.TestPermissionAllowed.ID:    false,
				records.TestPermissionNotAllowed.ID: false,
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testDB.Run(t, func(ctx context.Context, db database.DB) {
				if tt.initial != nil {
					require.NoError(t, testrecords.NewInserterForMySQL().
						Add(queries.UsersTable.TableName, tt.initial.userRecord).
						Add(queries.RolesTable.TableName, tt.initial.roleRecord).
						Add(queries.UsersRoleTable.TableName, tt.initial.userRoleRecord).
						InsertAll(ctx, db.Conn(ctx)))

					sql, args := queries.BulkUpsertRolePermissions(role.ID(tt.initial.roleRecord.ID), permission.NewValueMap(tt.initial.dbValueMap))
					_, err := db.Conn(ctx).ExecContext(ctx, sql, args...)
					require.NoError(t, err, sql)
				}

				r := permissionRepository{db: db}

				got, err := r.GetPermissions(ctx, tt.userID, tt.perms...)
				if !tt.wantErr(t, err) {
					return
				}

				assert.Equal(t, permission.NewValueMap(tt.want), got)
			})
		})
	}
}
