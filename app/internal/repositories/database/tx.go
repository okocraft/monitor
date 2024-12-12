package database

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Siroshun09/serrors"
)

var (
	ErrFailedToBegin    = errors.New("tx begin err")
	ErrFailedToRollback = errors.New("rollback err")
	ErrFunctionError    = errors.New("function err")
	ErrFailedToCommit   = errors.New("commit err")
)

type Transaction interface {
	WithTx(ctx context.Context, fn func(ctx context.Context) error) error
}

func NewTransaction(db DB) Transaction {
	return transaction{db: db.Base(), opts: &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	}}
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
		return serrors.WithStackTrace(errors.Join(ErrFailedToBegin, beginErr))
	}

	var fnErr error
	defer func() {
		if fnErr != nil {
			rbErr := tx.Rollback()
			if rbErr != nil {
				returnErr = serrors.WithStackTrace(errors.Join(ErrFailedToRollback, fnErr, rbErr))
				return
			}
		}
	}()

	ctx = SetTx(ctx, tx)
	if fnErr = fn(ctx); fnErr != nil {
		return serrors.WithStackTrace(errors.Join(ErrFunctionError, fnErr))
	}

	if err := tx.Commit(); err != nil {
		return serrors.WithStackTrace(errors.Join(ErrFailedToCommit, err))
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
