package sort

import (
	"fmt"

	"github.com/Siroshun09/serrors/v2"
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
		return 0, serrors.Wrap(fmt.Errorf("invalid sort type: %v", sortType))
	}
}
