package records

import (
	"context"

	"github.com/Siroshun09/serrors"
	"github.com/huandu/go-sqlbuilder"
	"github.com/okocraft/monitor/internal/repositories/database"
)

type InitialRecords struct {
	tables         []string
	recordsByTable map[string][]any
}

func NewInitialRecords() *InitialRecords {
	return &InitialRecords{
		tables:         []string{},
		recordsByTable: map[string][]any{},
	}
}

func (b *InitialRecords) Table(tableName string, records ...any) *InitialRecords {
	if existingRecords, exists := b.recordsByTable[tableName]; exists {
		existingRecords = append(existingRecords, records...)
		b.recordsByTable[tableName] = existingRecords
		return b
	}

	b.tables = append(b.tables, tableName)
	b.recordsByTable[tableName] = records
	return b
}

func (b *InitialRecords) InsertAll(ctx context.Context, db database.DB) error {
	for _, tableName := range b.tables {
		records, ok := b.recordsByTable[tableName]
		if !ok || len(records) == 0 {
			continue
		}

		sql, args := sqlbuilder.NewStruct(records[0]).For(sqlbuilder.MySQL).InsertInto(tableName, records...).Build()
		_, err := db.Conn(ctx).ExecContext(ctx, sql, args...)
		if err != nil {
			return serrors.WithStackTrace(err)
		}
	}

	return nil
}
