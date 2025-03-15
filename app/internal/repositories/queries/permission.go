package queries

import (
	"github.com/huandu/go-sqlbuilder"
	"github.com/okocraft/monitor/internal/domain/permission"
	"github.com/okocraft/monitor/internal/domain/role"
)

var rolePermissionStruct = sqlbuilder.NewStruct(new(RolesPermission)).For(sqlbuilder.MySQL)

func BulkUpsertRolePermissions(roleID role.ID, valueMap permission.ValueMap) (string, []any) {
	values := make([]any, 0, valueMap.Len())
	for id, state := range valueMap.Iter {
		values = append(values, RolesPermission{
			RoleID:       int32(roleID),
			PermissionID: id,
			IsAllowed:    state,
		})
	}
	b := rolePermissionStruct.InsertInto(RolesPermissionsTable.TableName, values...)
	b = ToUpsert(b, RolesPermissionsTable.IsAllowed)
	return b.Build()
}
