package testdb_test

import (
	"context"
	"testing"

	"github.com/okocraft/monitor/lib/testutils/testdb"
	"github.com/stretchr/testify/require"
)

func TestNewTestDB(t *testing.T) {
	db, err := testdb.NewTestDB(true)
	require.NoError(t, err)
	require.NotNil(t, db)

	t.Cleanup(func() {
		require.NoError(t, db.Cleanup())
	})

	conn, err := db.Conn()
	require.NoError(t, err)
	require.NotNil(t, conn)

	err = conn.Base().PingContext(context.Background())
	require.NoError(t, err)
}
