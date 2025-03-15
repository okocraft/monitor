package role

import (
	"context"
	"database/sql"

	"github.com/Siroshun09/logs"
	"github.com/okocraft/monitor/internal/domain/role"
	"github.com/okocraft/monitor/internal/domain/sort"
	"github.com/okocraft/monitor/internal/repositories/database"
	"github.com/okocraft/monitor/internal/repositories/queries"
	"github.com/okocraft/monitor/lib/null"
)

type RoleRepository interface {
	GetAllRoles(ctx context.Context, sortedBy null.Optional[role.SortableDataType], sortType null.Optional[sort.Type]) (role.Roles, error)
}

func NewRoleRepository(db database.DB) RoleRepository {
	return roleRepository{db: db}
}

type roleRepository struct {
	db database.DB
}

func (r roleRepository) GetAllRoles(ctx context.Context, sortedBy null.Optional[role.SortableDataType], sortType null.Optional[sort.Type]) (role.Roles, error) {
	conn := r.db.Conn(ctx)
	query, args := queries.GetSortedAllRoles(sortedBy, sortType)
	rows, err := conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		closeErr := rows.Close()
		if closeErr != nil {
			logs.Error(ctx, closeErr)
		}
	}(rows)

	roles := []role.Role{}
	for rows.Next() {
		record, err := queries.ScanRoleRow(rows)
		if err != nil {
			return nil, database.NewDBErrorWithStackTrace(err)
		}
		roles = append(roles, record.ToDomain())
	}
	return roles, nil
}
