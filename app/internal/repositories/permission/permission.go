package permission

import (
	"context"

	"github.com/okocraft/monitor/internal/domain/permission"
	"github.com/okocraft/monitor/internal/domain/role"
	"github.com/okocraft/monitor/internal/domain/user"
	"github.com/okocraft/monitor/internal/repositories/database"
	"github.com/okocraft/monitor/internal/repositories/queries"
)

type PermissionRepository interface {
	GetUserPermissions(ctx context.Context, userID user.ID, perms ...permission.Permission) (permission.ValueMap, error)
	SaveRolePermissions(ctx context.Context, roleID role.ID, valueMap permission.ValueMap) error
}

func NewPermissionRepository(db database.DB) PermissionRepository {
	return permissionRepository{
		db: db,
	}
}

type permissionRepository struct {
	db database.DB
}

func (r permissionRepository) GetUserPermissions(ctx context.Context, userID user.ID, perms ...permission.Permission) (permission.ValueMap, error) {
	q := r.db.Queries(ctx)

	ids := make([]int16, 0, len(perms))
	result := make(permission.ValueMapSource, len(perms))
	for _, perm := range perms {
		ids = append(ids, perm.ID)
		result[perm.ID] = perm.DefaultValue
	}

	records, err := q.GetPermissionsByUserID(ctx, queries.GetPermissionsByUserIDParams{UserID: int32(userID), PermissionIds: ids})
	for _, record := range records {
		result[record.PermissionID] = record.IsAllowed
	}

	if err != nil {
		return permission.EmptyValueMap(), database.NewDBErrorWithStackTrace(err)
	}
	return permission.NewValueMap(result), nil
}

func (r permissionRepository) SaveRolePermissions(ctx context.Context, roleID role.ID, valueMap permission.ValueMap) error {
	conn := r.db.Conn(ctx)
	query, args := queries.BulkUpsertRolePermissions(roleID, valueMap)
	_, err := conn.ExecContext(ctx, query, args...)
	if err != nil {
		return database.NewDBErrorWithStackTrace(err)
	}
	return nil
}
