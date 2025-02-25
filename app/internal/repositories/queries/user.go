package queries

import (
	"github.com/gofrs/uuid/v5"
	"github.com/huandu/go-sqlbuilder"
	"github.com/okocraft/monitor/internal/domain/sort"
	"github.com/okocraft/monitor/internal/domain/user"
)

func SearchAndGetUUIDs(params user.SearchParams) (string, []any) {
	b := sqlbuilder.MySQL.NewSelectBuilder()

	b.Select(tableColumnRef(UsersTable.TableName, UsersTable.Uuid)).From(UsersTable.TableName)
	b.JoinWithOption(sqlbuilder.LeftOuterJoin, UsersRoleTable.TableName, joinCondition(UsersRoleTable.TableName, UsersRoleTable.UserId, UsersTable.TableName, UsersTable.Id))
	b.JoinWithOption(sqlbuilder.LeftOuterJoin, RolesTable.TableName, joinCondition(RolesTable.TableName, RolesTable.Id, UsersRoleTable.TableName, UsersRoleTable.RoleId))

	if params.Nickname.Valid() {
		b.Where(b.Like(UsersTable.Nickname, likePartialMatch(params.Nickname.Get())))
	}

	if params.LastAccessBefore.Valid() {
		b.Where(b.LTE(UsersTable.LastAccess, params.LastAccessBefore.Get()))
	}

	if params.LastAccessAfter.Valid() {
		b.Where(b.GTE(UsersTable.LastAccess, params.LastAccessAfter.Get()))
	}

	if params.RoleId.Valid() {
		if params.RoleId.Get() == uuid.Nil {
			b.Where(b.IsNull(tableColumnRef(RolesTable.TableName, RolesTable.Uuid)))
		} else {
			b.Where(b.EQ(tableColumnRef(RolesTable.TableName, RolesTable.Uuid), params.RoleId.Get().Bytes()))
		}
	}

	if params.SortedBy.Valid() {
		var order string
		switch {
		case params.SortType.Valid() && params.SortType.Get() == sort.ASC:
			order = " ASC"
		case params.SortType.Valid() && params.SortType.Get() == sort.DESC:
			order = " DESC"
		}
		switch params.SortedBy.Get() {
		case user.SortableDataTypeNickName:
			b.OrderBy(tableColumnRef(UsersTable.TableName, UsersTable.Nickname) + order)
		case user.SortableDataTypeLastAccess:
			b.OrderBy(tableColumnRef(UsersTable.TableName, UsersTable.LastAccess) + order)
		case user.SortableDataTypeCreatedAt:
			b.OrderBy(tableColumnRef(UsersTable.TableName, UsersTable.CreatedAt) + order)
		case user.SortableDataTypeUpdatedAt:
			b.OrderBy(tableColumnRef(UsersTable.TableName, UsersTable.UpdatedAt) + order)
		case user.SortableDataTypeRoleName:
			b.OrderBy(tableColumnRef(RolesTable.TableName, RolesTable.Name) + order)
		case user.SortableDataTypeRolePriority:
			b.OrderBy(tableColumnRef(RolesTable.TableName, RolesTable.Priority) + order)
		}
	}

	query, args := b.Build()
	return query, args
}
