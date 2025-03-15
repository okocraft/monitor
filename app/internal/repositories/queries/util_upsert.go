package queries

import (
	"strings"

	"github.com/huandu/go-sqlbuilder"
)

func ToUpsert(ib *sqlbuilder.InsertBuilder, updateCols ...string) *sqlbuilder.InsertBuilder {
	if len(updateCols) == 0 {
		panic("no update column set")
	}
	cols := make([]string, len(updateCols))
	for i, col := range updateCols {
		cols[i] = col + " = VALUES(" + col + ")"
	}

	b := *ib
	b.SQL("ON DUPLICATE KEY UPDATE " + strings.Join(cols, ", "))
	return &b
}
