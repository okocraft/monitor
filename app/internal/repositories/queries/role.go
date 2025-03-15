package queries

import (
	"database/sql"

	"github.com/Siroshun09/serrors"
	"github.com/gofrs/uuid/v5"
	"github.com/huandu/go-sqlbuilder"
	"github.com/okocraft/monitor/internal/domain/role"
	"github.com/okocraft/monitor/internal/domain/sort"
	"github.com/okocraft/monitor/lib/null"
)

func (record Role) ToDomain() role.Role {
	return role.Role{
		ID:        role.ID(record.ID),
		UUID:      uuid.UUID(record.Uuid),
		Name:      record.Name,
		Priority:  record.Priority,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
	}
}

var roleStruct = sqlbuilder.NewStruct(new(Role)).For(sqlbuilder.MySQL)

func GetSortedAllRoles(sortedBy null.Optional[role.SortableDataType], sortType null.Optional[sort.Type]) (string, []any) {
	b := roleStruct.SelectFrom(RolesTable.TableName)

	if sortedBy.Valid() {
		switch sortedBy.Get() {
		case role.SortableDataTypeName:
			b.OrderBy(RolesTable.Name)
		case role.SortableDataTypePriority:
			b.OrderBy(RolesTable.Priority)
		case role.SortableDataTypeCreatedAt:
			b.OrderBy(RolesTable.CreatedAt)
		case role.SortableDataTypeUpdatedAt:
			b.OrderBy(RolesTable.UpdatedAt)
		}

		switch {
		case sortType.Valid() && sortType.Get() == sort.ASC:
			b.Asc()
		case sortType.Valid() && sortType.Get() == sort.DESC:
			b.Desc()
		}
	}

	return b.Build()
}

func ScanRoleRow(rows *sql.Rows) (Role, error) {
	record := Role{}
	err := rows.Scan(roleStruct.Addr(&record)...)
	if err != nil {
		return Role{}, serrors.WithStackTrace(err)
	}
	return record, nil
}
