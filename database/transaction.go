package database

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/constant"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/exception"
)

type transactor struct {
	db *sql.DB
}

type Runner struct {
	db *sql.DB
	tx *sql.Tx
}

type txKey struct{}

type Transactor interface {
	WithinTransaction(context.Context, func(context.Context) (any, error)) error
}

type RunnerImpl interface {
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, q string, args ...interface{}) *sql.Row
}

func NewRunner(db *sql.DB, tx *sql.Tx) *Runner {
	return &Runner{
		db: db,
		tx: tx,
	}
}

func NewTransaction(db *sql.DB) *transactor {
	return &transactor{
		db: db,
	}
}

func GetQueryRunner(ctx context.Context) *sql.Tx {
	tx := ExtractTx(ctx)
	if tx == nil {
		return nil
	}

	return tx
}

func injectTx(ctx context.Context, tx *sql.Tx) context.Context {

	return context.WithValue(ctx, txKey{}, tx)
}

func ExtractTx(ctx context.Context) *sql.Tx {
	if tx, ok := ctx.Value(txKey{}).(*sql.Tx); ok {
		return tx
	}
	return nil
}

func (r *Runner) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	if r.tx != nil {
		return r.tx.QueryContext(ctx, query, args...)
	}
	return r.db.QueryContext(ctx, query, args...)
}

func (r *Runner) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	if r.tx != nil {
		return r.tx.QueryRowContext(ctx, query, args...)
	}
	return r.db.QueryRowContext(ctx, query, args...)
}

func (tr *transactor) WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) (any, error)) error {

	tx, err := tr.db.Begin()
	if err != nil {
		return exception.NewErrorType(http.StatusBadRequest, constant.ResponseMsgBadRequest)
	}

	defer func() {
		if errTx := tr.db.Close(); errTx != nil {
			log.Printf("close transaction: %v", errTx)
		}
	}()

	_, err = tFunc(injectTx(ctx, tx))
	if err != nil {

		if errRollback := tx.Rollback(); errRollback != nil {
			log.Printf("rollback transaction: %v", errRollback)
		}
		return exception.NewErrorType(http.StatusBadRequest, constant.ResponseMsgBadRequest)
	}
	if errCommit := tx.Commit(); errCommit != nil {
		log.Printf("commit transaction: %v", errCommit)
	}
	return nil
}
