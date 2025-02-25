package sort

import (
	"github.com/Siroshun09/serrors"
	"github.com/okocraft/monitor/internal/handler/oapi"
)

type Type int8

const (
	ASC Type = iota + 1
	DESC
)

func FromRequest(sortType oapi.SortType) (Type, error) {
	switch sortType {
	case oapi.SortTypeASC:
		return ASC, nil
	case oapi.SortTypeDESC:
		return DESC, nil
	default:
		return 0, serrors.Errorf("invalid sort type: %v", sortType)
	}
}
