package queries

import (
	"time"

	"github.com/huandu/go-sqlbuilder"
	"github.com/okocraft/monitor/internal/domain/role"
	"github.com/okocraft/monitor/internal/domain/user"
)

var usersRoleStruct = sqlbuilder.NewStruct(new(UsersRole)).For(sqlbuilder.MySQL)

func SaveUsersRole(userID user.ID, roleID role.ID, now time.Time) (string, []any) {
	b := usersRoleStruct.InsertInto(UsersRoleTable.TableName).Values(int(userID), int(roleID), now)
	b = ToUpsert(b, UsersRoleTable.RoleId, UsersRoleTable.UpdatedAt)
	return b.Build()
}
