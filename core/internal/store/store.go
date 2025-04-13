package store

import (
	"context"
	"database/sql"
	"fmt"
)

type QueryExecutor interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

type QueryExecutorTx interface {
	QueryExecutor
	Commit() error
	Rollback() error
}

type ExecFn func(QueryExecutor) error

type DBTX struct {
	db *sql.DB
}

func New(db *sql.DB) *DBTX {
	return &DBTX{
		db: db,
	}
}

func (d *DBTX) WithTransaction(ctx context.Context, fn ExecFn) (err error) {
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and re-panic
			_ = tx.Rollback() // nolint: errcheck

			panic(p)
		}
	}()

	return handleTxError(tx, fn)
}

func (d *DBTX) DB() QueryExecutor {
	return d.db
}

func (d *DBTX) TX(ctx context.Context) (QueryExecutorTx, error) {
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("begin transaction: %w", err)
	}

	return tx, nil
}

func handleTxError(tx *sql.Tx, fn ExecFn) error {
	err := fn(tx)
	if err != nil {
		if errRb := tx.Rollback(); errRb != nil {
			return fmt.Errorf("cannot rollback transaction db: %w", err)
		}

		return err
	}

	if errComm := tx.Commit(); errComm != nil {
		return fmt.Errorf("cannot commit transaction db: %w", errComm)
	}

	return err
}
