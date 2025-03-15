package user

import (
	"time"

	"github.com/Siroshun09/serrors"
	"github.com/gofrs/uuid/v5"
	"github.com/okocraft/monitor/internal/domain/sort"
	"github.com/okocraft/monitor/internal/handler/oapi"
	"github.com/okocraft/monitor/lib/errlib"
	"github.com/okocraft/monitor/lib/null"
)

type SortableDataType int8

const (
	SortableDataTypeNickName SortableDataType = iota + 1
	SortableDataTypeLastAccess
	SortableDataTypeCreatedAt
	SortableDataTypeUpdatedAt
	SortableDataTypeRoleName
	SortableDataTypeRolePriority
)

type SearchParams struct {
	Nickname         null.Optional[string]
	LastAccessBefore null.Optional[time.Time]
	LastAccessAfter  null.Optional[time.Time]
	RoleId           null.Optional[uuid.UUID]
	SortedBy         null.Optional[SortableDataType]
	SortType         null.Optional[sort.Type]
}

func NewSearchParamsFromRequest(param oapi.SearchUsersParams) (SearchParams, error) {
	var sortedBy null.Optional[SortableDataType]
	if param.SortedBy != nil {
		switch *param.SortedBy {
		case oapi.SortableUserDataTypeNickName:
			sortedBy = null.FromValue(SortableDataTypeNickName)
		case oapi.SortableUserDataTypeLastAccess:
			sortedBy = null.FromValue(SortableDataTypeLastAccess)
		case oapi.SortableUserDataTypeCreatedAt:
			sortedBy = null.FromValue(SortableDataTypeCreatedAt)
		case oapi.SortableUserDataTypeUpdatedAt:
			sortedBy = null.FromValue(SortableDataTypeUpdatedAt)
		case oapi.SortableUserDataTypeRoleName:
			sortedBy = null.FromValue(SortableDataTypeRoleName)
		case oapi.SortableUserDataTypeRolePriority:
			sortedBy = null.FromValue(SortableDataTypeRolePriority)
		default:
			return SearchParams{}, serrors.Errorf("invalid sortable user data type: %s", *param.SortedBy)
		}
	} else {
		sortedBy = null.Empty[SortableDataType]()
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
