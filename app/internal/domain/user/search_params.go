package user

import (
	"github.com/Siroshun09/serrors"
	"github.com/gofrs/uuid/v5"
	"github.com/okocraft/monitor/internal/domain/sort"
	"github.com/okocraft/monitor/internal/handler/oapi"
	"github.com/okocraft/monitor/lib/errlib"
	"github.com/okocraft/monitor/lib/null"
	"time"
)

type SortableUserDataType int8

const (
	NickName SortableUserDataType = iota + 1
	LastAccess
	CreatedAt
	UpdatedAt
	RoleName
	RolePriority
)

type SearchParams struct {
	Nickname         null.Optional[string]
	LastAccessBefore null.Optional[time.Time]
	LastAccessAfter  null.Optional[time.Time]
	RoleId           null.Optional[uuid.UUID]
	SortedBy         null.Optional[SortableUserDataType]
	SortType         null.Optional[sort.Type]
}

func NewSearchParamsFromRequest(param oapi.SearchUsersParams) (SearchParams, error) {
	var sortedBy null.Optional[SortableUserDataType]
	if param.SortedBy != nil {
		switch *param.SortedBy {
		case oapi.SortableUserDataTypeNickName:
			sortedBy = null.FromValue(NickName)
		case oapi.SortableUserDataTypeLastAccess:
			sortedBy = null.FromValue(LastAccess)
		case oapi.SortableUserDataTypeCreatedAt:
			sortedBy = null.FromValue(CreatedAt)
		case oapi.SortableUserDataTypeUpdatedAt:
			sortedBy = null.FromValue(UpdatedAt)
		case oapi.SortableUserDataTypeRoleName:
			sortedBy = null.FromValue(RoleName)
		case oapi.SortableUserDataTypeRolePriority:
			sortedBy = null.FromValue(RolePriority)
		default:
			return SearchParams{}, serrors.Errorf("invalid sortable user data type: %s", *param.SortedBy)
		}
	} else {
		sortedBy = null.Empty[SortableUserDataType]()
	}

	var sortType null.Optional[sort.Type]
	if param.SortType != nil {
		parsedSortType, err := sort.FromRequest(*param.SortType)
		if err != nil {
			return SearchParams{}, errlib.AsIs(err)
		}
		sortType = null.FromValue(parsedSortType)
	} else {
		sortType = null.Empty[sort.Type]()
	}

	return SearchParams{
		Nickname:         null.FromPtr(param.Nickname),
		LastAccessBefore: null.FromPtr(param.LastAccessBefore),
		LastAccessAfter:  null.FromPtr(param.LastAccessAfter),
		RoleId:           null.FromPtr(param.RoleId),
		SortedBy:         sortedBy,
		SortType:         sortType,
	}, nil
}
