package testdb

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Siroshun09/serrors"
	"github.com/gofrs/uuid/v5"
	"github.com/okocraft/monitor/internal/config"
	"github.com/okocraft/monitor/internal/repositories/database"
	"github.com/okocraft/monitor/lib/testutils"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

type TestDB interface {
	Run(t *testing.T, f func(ctx context.Context, db database.DB))
	Cleanup() error
}

func NewTestDB(useTx bool) (TestDB, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, serrors.WithStackTrace(err)
	}

	dbConfig, err := config.NewDBConfigFromEnv()
	if err != nil {
		dbConfig = config.DBConfig{
			Host:     "localhost",
			Port:     "3306",
			User:     "monitor_user",
			Password: "monitor_pw",
		}
	}

	cfg := database.GenerateConfig(dbConfig)
	conn, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, serrors.WithStackTrace(err)
	}

	dbConfig.DBName = "testdb_" + strings.ReplaceAll(id.String(), "-", "")
	createDB := "CREATE " + "DATABASE " + dbConfig.DBName // Prevent detection that this is SQL statement because this would be recognized as an error
	_, err = conn.Exec(createDB)
	if err != nil {
		return nil, serrors.WithStackTrace(err)
	}

	_, err = conn.Exec("USE " + dbConfig.DBName)
	if err != nil {
		return nil, serrors.WithStackTrace(err)
	}

	err = createTables(conn)
	if err != nil {
		return nil, serrors.WithStackTrace(err)
	}

	return &testDB{
		cfg:   dbConfig,
		useTx: useTx,
	}, nil
}

func createTables(db *sql.DB) error {
	rootDir, err := testutils.GetProjectRoot()
	if err != nil {
		return serrors.WithStackTrace(err)
	}

	schema, err := os.ReadFile(filepath.Join(rootDir, "../schema/database/schema.sql"))
	if err != nil {
		return serrors.WithStackTrace(err)
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		return serrors.WithStackTrace(err)
	}

	return nil
}

type testDB struct {
	cfg   config.DBConfig
	useTx bool
}

func (db *testDB) Run(t *testing.T, f func(ctx context.Context, db database.DB)) {
	ctx := context.Background()
	d, err := database.New(db.cfg, 15*time.Minute)
	require.NoError(t, err)

	if !db.useTx {
		f(ctx, d)
		return
	}

	tx, err := d.Base().BeginTx(context.Background(), &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		require.NoError(t, tx.Rollback())
	})

	ctx = database.SetTx(ctx, tx)
	f(ctx, d)
}

func (db *testDB) Cleanup() (err error) {
	cfg := db.cfg
	cfg.DBName = ""
	conn, err := sql.Open("mysql", database.GenerateConfig(cfg).FormatDSN())
	if err != nil {
		return serrors.WithStackTrace(err)
	}

	defer func(conn *sql.DB) {
		closeErr := conn.Close()
		if closeErr != nil {
			err = errors.Join(err, closeErr)
		}
	}(conn)

	_, err = conn.Exec("DROP " + "DATABASE " + db.cfg.DBName)
	if err != nil {
		return serrors.WithStackTrace(err)
	}

	return nil
}
