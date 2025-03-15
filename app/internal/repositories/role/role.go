package role

import (
	"context"
	"database/sql"
	"time"

	"github.com/Siroshun09/logs"
	"github.com/okocraft/monitor/internal/domain/role"
	"github.com/okocraft/monitor/internal/domain/sort"
	"github.com/okocraft/monitor/internal/domain/user"
	"github.com/okocraft/monitor/internal/repositories/database"
	"github.com/okocraft/monitor/internal/repositories/queries"
	"github.com/okocraft/monitor/lib/null"
)

type RoleRepository interface {
	GetAllRoles(ctx context.Context, sortedBy null.Optional[role.SortableDataType], sortType null.Optional[sort.Type]) (role.Roles, error)
	ExistsRoleByID(ctx context.Context, id role.ID) (bool, error)

	CreateRoleWithIDIfNotExists(ctx context.Context, role role.Role) error

	SaveUserRole(ctx context.Context, userID user.ID, roleID role.ID, now time.Time) error
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

func (r roleRepository) ExistsRoleByID(ctx context.Context, id role.ID) (bool, error) {
	q := r.db.Queries(ctx)
	exists, err := q.ExistsRoleByID(ctx, int32(id))
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r roleRepository) CreateRoleWithIDIfNotExists(ctx context.Context, role role.Role) error {
	q := r.db.Queries(ctx)
	err := q.CreateRoleWithIDIfNotExists(ctx, queries.CreateRoleWithIDIfNotExistsParams{
		ID:        int32(role.ID),
		Uuid:      role.UUID.Bytes(),
		Name:      role.Name,
		Priority:  role.Priority,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	})
	if err != nil {
		return database.NewDBErrorWithStackTrace(err)
	}
	return nil
}

func (r roleRepository) SaveUserRole(ctx context.Context, userID user.ID, roleID role.ID, now time.Time) error {
	conn := r.db.Conn(ctx)
	query, args := queries.SaveUsersRole(userID, roleID, now)
	_, err := conn.ExecContext(ctx, query, args...)
	if err != nil {
		return database.NewDBErrorWithStackTrace(err)
	}
	return nil
}
