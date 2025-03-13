package role

import (
	"github.com/Siroshun09/serrors"
	"github.com/okocraft/monitor/internal/handler/oapi"
)

type SortableDataType int8

const (
	SortableDataTypeName SortableDataType = iota + 1
	SortableDataTypePriority
	SortableDataTypeCreatedAt
	SortableDataTypeUpdatedAt
)

func ConvertSortableDataTypeFromRequest(req oapi.SortableRoleDataType) (SortableDataType, error) {
	switch req {
	case oapi.SortableRoleDataTypeName:
		return SortableDataTypeName, nil
	case oapi.SortableRoleDataTypePriority:
		return SortableDataTypePriority, nil
	case oapi.SortableRoleDataTypeCreatedAt:
		return SortableDataTypeCreatedAt, nil
	case oapi.SortableRoleDataTypeUpdatedAt:
		return SortableDataTypeUpdatedAt, nil
	default:
		return 0, serrors.New("invalid SortableRoleDataType")
	}
}
