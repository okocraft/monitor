package repositories

import (
	"github.com/Siroshun09/serrors"
	"github.com/gofrs/uuid/v5"
	"github.com/okocraft/monitor/internal/repositories/database"
)

func asDBError(err error) error {
	return serrors.WithStackTrace(database.NewDBError(err))
}

func toBytesSlice(uuids []uuid.UUID) [][]byte {
	result := make([][]byte, 0, len(uuids))
	for _, id := range uuids {
		result = append(result, id.Bytes())
	}
	return result
}
