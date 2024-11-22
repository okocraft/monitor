package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/Siroshun09/serrors"
	"github.com/go-sql-driver/mysql"
	"github.com/okocraft/monitor/internal/config"
	"github.com/okocraft/monitor/internal/repositories/queries"
)

type DB interface {
	Queries(ctx context.Context) *queries.Queries
	Base() *sql.DB
	Close() error
}

func GenerateConfig(c config.DBConfig) *mysql.Config {
	cfg := mysql.NewConfig()
	cfg.User = c.User
	cfg.Passwd = c.Password
	cfg.Addr = c.Host + ":" + c.Port
	cfg.DBName = c.DBName
	cfg.MultiStatements = true
	cfg.ParseTime = true
	return cfg
}

func New(c config.DBConfig, maxLifeTime time.Duration) (DB, error) {
	conn, err := sql.Open("mysql", GenerateConfig(c).FormatDSN())
	if err != nil {
		return nil, serrors.WithStackTrace(err)
	}
	conn.SetConnMaxLifetime(maxLifeTime)
	return db{conn: conn}, nil
}

type db struct {
	conn *sql.DB
}

func (db db) Queries(ctx context.Context) *queries.Queries {
	q := queries.New(db.conn)
	tx, ok := getTx(ctx)
	if ok {
		return q.WithTx(tx)
	}
	return q
}

func (db db) Base() *sql.DB {
	return db.conn
}

func (db db) Close() error {
	err := db.conn.Close()
	if err != nil {
		return serrors.WithStackTrace(err)
	}
	return nil
}
