package queries

import "github.com/gofrs/uuid/v5"

func ToAnySlice[T any](ts []T) []any {
	a := make([]any, 0, len(ts))
	for _, v := range ts {
		a = append(a, v)
	}
	return a
}

func ToBytesSlice(uuids []uuid.UUID) [][]byte {
	result := make([][]byte, 0, len(uuids))
	for _, id := range uuids {
		result = append(result, id.Bytes())
	}
	return result
}
