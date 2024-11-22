package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/okocraft/monitor/lib/errlib"
)

type FailType string

const (
	FailedToBegin    FailType = "tx begin err"
	FailedToRollback          = "rollback err"
	FunctionError             = "function err"
	FailedToCommit            = "commit err"
)

type Transaction interface {
	WithTx(ctx context.Context, fn func(ctx context.Context) error) error
}

func NewTransaction(db *sql.DB, opts *sql.TxOptions) Transaction {
	return transaction{db: db, opts: opts}
}

type transaction struct {
	db   *sql.DB
	opts *sql.TxOptions
}

func (t transaction) WithTx(ctx context.Context, fn func(ctx context.Context) error) (returnErr error) {
	if _, hasTx := getTx(ctx); hasTx {
		return fn(ctx)
	}

	tx, beginErr := t.db.BeginTx(ctx, t.opts)
	if beginErr != nil {
		return errlib.WithDetail(FailedToBegin, beginErr)
	}

	var fnErr error
	defer func() {
		if fnErr != nil {
			rbErr := tx.Rollback()
			if rbErr != nil {
				returnErr = errlib.WithDetail(FailedToRollback, errors.Join(fnErr, rbErr))
				return
			}
		}
	}()

	ctx = SetTx(ctx, tx)
	if fnErr = fn(ctx); fnErr != nil {
		return errlib.WithDetail(FunctionError, fnErr)
	}

	if err := tx.Commit(); err != nil {
		return errlib.WithDetail(FailedToCommit, err)
	}

	return nil
}

type contextKey int8

const transactionKey contextKey = 1

func getTx(ctx context.Context) (*sql.Tx, bool) {
	tx, ok := ctx.Value(transactionKey).(*sql.Tx)
	if !ok {
		return nil, false
	}
	return tx, true
}

func SetTx(ctx context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, transactionKey, tx)
}
