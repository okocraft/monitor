package role_test

import (
	"context"
	"testing"

	"github.com/Siroshun09/testrecords"
	"github.com/okocraft/monitor/internal/domain/role"
	"github.com/okocraft/monitor/internal/domain/sort"
	"github.com/okocraft/monitor/internal/repositories/database"
	"github.com/okocraft/monitor/internal/repositories/queries"
	"github.com/okocraft/monitor/internal/repositories/records"
	roleRepo "github.com/okocraft/monitor/internal/repositories/role"
	"github.com/okocraft/monitor/lib/null"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_roleRepository_GetAllRoles(t *testing.T) {
	tests := []struct {
		name     string
		initial  testrecords.Inserter
		sortedBy null.Optional[role.SortableDataType]
		sortType null.Optional[sort.Type]
		want     role.Roles
		wantErr  bool
	}{
		{
			name:    "success: 1 record",
			initial: testrecords.NewInserterForMySQL().Add(queries.RolesTable.TableName, records.Role1),
			want:    role.Roles{records.Role1.ToDomain()},
			wantErr: false,
		},
		{
			name:     "success: 3 records, sorted by name (asc)",
			initial:  testrecords.NewInserterForMySQL().Add(queries.RolesTable.TableName, records.Role1, records.Role2, records.Role3),
			sortedBy: null.FromValue(role.SortableDataTypeName),
			sortType: null.FromValue(sort.ASC),
			want:     role.Roles{records.Role1.ToDomain(), records.Role2.ToDomain(), records.Role3.ToDomain()},
			wantErr:  false,
		},
		{
			name:     "success: 3 records, sorted by name (desc)",
			initial:  testrecords.NewInserterForMySQL().Add(queries.RolesTable.TableName, records.Role1, records.Role2, records.Role3),
			sortedBy: null.FromValue(role.SortableDataTypeName),
			sortType: null.FromValue(sort.DESC),
			want:     role.Roles{records.Role3.ToDomain(), records.Role2.ToDomain(), records.Role1.ToDomain()},
			wantErr:  false,
		},
		{
			name:     "success: 3 records, sorted by priority (asc)",
			initial:  testrecords.NewInserterForMySQL().Add(queries.RolesTable.TableName, records.Role1, records.Role2, records.Role3),
			sortedBy: null.FromValue(role.SortableDataTypePriority),
			sortType: null.FromValue(sort.ASC),
			want:     role.Roles{records.Role1.ToDomain(), records.Role2.ToDomain(), records.Role3.ToDomain()},
			wantErr:  false,
		},
		{
			name:     "success: 3 records, sorted by priority (desc)",
			initial:  testrecords.NewInserterForMySQL().Add(queries.RolesTable.TableName, records.Role1, records.Role2, records.Role3),
			sortedBy: null.FromValue(role.SortableDataTypePriority),
			sortType: null.FromValue(sort.DESC),
			want:     role.Roles{records.Role3.ToDomain(), records.Role2.ToDomain(), records.Role1.ToDomain()},
			wantErr:  false,
		},
		{
			name:     "success: 3 records, sorted by createdAt (asc)",
			initial:  testrecords.NewInserterForMySQL().Add(queries.RolesTable.TableName, records.Role1, records.Role2, records.Role3),
			sortedBy: null.FromValue(role.SortableDataTypeCreatedAt),
			sortType: null.FromValue(sort.ASC),
			want:     role.Roles{records.Role1.ToDomain(), records.Role2.ToDomain(), records.Role3.ToDomain()},
			wantErr:  false,
		},
		{
			name:     "success: 3 records, sorted by createdAt (desc)",
			initial:  testrecords.NewInserterForMySQL().Add(queries.RolesTable.TableName, records.Role1, records.Role2, records.Role3),
			sortedBy: null.FromValue(role.SortableDataTypeCreatedAt),
			sortType: null.FromValue(sort.DESC),
			want:     role.Roles{records.Role3.ToDomain(), records.Role2.ToDomain(), records.Role1.ToDomain()},
			wantErr:  false,
		},
		{
			name:     "success: 3 records, sorted by updatedAt (asc)",
			initial:  testrecords.NewInserterForMySQL().Add(queries.RolesTable.TableName, records.Role1, records.Role2, records.Role3),
			sortedBy: null.FromValue(role.SortableDataTypeUpdatedAt),
			sortType: null.FromValue(sort.ASC),
			want:     role.Roles{records.Role1.ToDomain(), records.Role2.ToDomain(), records.Role3.ToDomain()},
			wantErr:  false,
		},
		{
			name:     "success: 3 records, sorted by updatedAt (desc)",
			initial:  testrecords.NewInserterForMySQL().Add(queries.RolesTable.TableName, records.Role1, records.Role2, records.Role3),
			sortedBy: null.FromValue(role.SortableDataTypeUpdatedAt),
			sortType: null.FromValue(sort.DESC),
			want:     role.Roles{records.Role3.ToDomain(), records.Role2.ToDomain(), records.Role1.ToDomain()},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testDB.Run(t, func(ctx context.Context, db database.DB) {
				require.NoError(t, tt.initial.InsertAll(ctx, db.Conn(ctx)))

				r := roleRepo.NewRoleRepository(db)
				got, err := r.GetAllRoles(ctx, tt.sortedBy, tt.sortType)
				if tt.wantErr {
					assert.Error(t, err)
					assert.Zero(t, got)
				} else {
					require.NoError(t, err)
					assert.Equal(t, tt.want, got)
				}
			})
		})
	}
}
