package repositories

import (
	"github.com/Siroshun09/serrors"
	"github.com/okocraft/monitor/internal/repositories/database"
)

func asDBError(err error) error {
	return serrors.WithStackTrace(database.NewDBError(err))
}
