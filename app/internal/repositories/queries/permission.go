package queries

import (
	"github.com/huandu/go-sqlbuilder"
	"github.com/okocraft/monitor/internal/domain/permission"
	"github.com/okocraft/monitor/internal/domain/role"
)

var rolePermissionStruct = sqlbuilder.NewStruct(new(RolesPermission)).For(sqlbuilder.MySQL)

func BulkInsertRolePermissions(roleID role.ID, valueMap permission.ValueMap) (string, []any) {
	values := make([]any, 0, valueMap.Len())
	for id, state := range valueMap.Iter {
		values = append(values, RolesPermission{
			RoleID:       roleID,
			PermissionID: id,
			IsAllowed:    state,
		})
	}
	return rolePermissionStruct.InsertInto("roles_permissions", values...).Build()
}
