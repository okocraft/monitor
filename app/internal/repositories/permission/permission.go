package permission

import (
	"context"
	"database/sql"
	"errors"
	"github.com/okocraft/monitor/internal/domain/permission"
	"github.com/okocraft/monitor/internal/domain/user"
	"github.com/okocraft/monitor/internal/repositories/database"
	"github.com/okocraft/monitor/internal/repositories/queries"
)

type PermissionRepository interface {
	HasPermission(ctx context.Context, userID user.ID, perm permission.Permission) (bool, error)
	GetPermissions(ctx context.Context, userID user.ID, perms ...permission.Permission) (permission.ValueMap, error)
}

func NewPermissionRepository(db database.DB) PermissionRepository {
	return permissionRepository{
		db: db,
	}
}

type permissionRepository struct {
	db database.DB
}

func (r permissionRepository) HasPermission(ctx context.Context, userID user.ID, perm permission.Permission) (bool, error) {
	q := r.db.Queries(ctx)
	result, err := q.CheckPermissionByUserID(ctx, queries.CheckPermissionByUserIDParams{UserID: int32(userID), PermissionID: perm.ID})
	if errors.Is(err, sql.ErrNoRows) {
		return perm.DefaultValue, nil
	} else if err != nil {
		return false, database.NewDBErrorWithStackTrace(err)
	}
	return result, nil
}

func (r permissionRepository) GetPermissions(ctx context.Context, userID user.ID, perms ...permission.Permission) (permission.ValueMap, error) {
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
