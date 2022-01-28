// Code generated by sqlc. DO NOT EDIT.

package generated

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createStockStmt, err = db.PrepareContext(ctx, createStock); err != nil {
		return nil, fmt.Errorf("error preparing query CreateStock: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.getNotUpdateStockCodeListStmt, err = db.PrepareContext(ctx, getNotUpdateStockCodeList); err != nil {
		return nil, fmt.Errorf("error preparing query GetNotUpdateStockCodeList: %w", err)
	}
	if q.getNotUpdateStockListStmt, err = db.PrepareContext(ctx, getNotUpdateStockList); err != nil {
		return nil, fmt.Errorf("error preparing query GetNotUpdateStockList: %w", err)
	}
	if q.getStockCountStmt, err = db.PrepareContext(ctx, getStockCount); err != nil {
		return nil, fmt.Errorf("error preparing query GetStockCount: %w", err)
	}
	if q.getStockListStmt, err = db.PrepareContext(ctx, getStockList); err != nil {
		return nil, fmt.Errorf("error preparing query GetStockList: %w", err)
	}
	if q.getUserByIdStmt, err = db.PrepareContext(ctx, getUserById); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserById: %w", err)
	}
	if q.getUserByNameStmt, err = db.PrepareContext(ctx, getUserByName); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByName: %w", err)
	}
	if q.updateUserStmt, err = db.PrepareContext(ctx, updateUser); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUser: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createStockStmt != nil {
		if cerr := q.createStockStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createStockStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.getNotUpdateStockCodeListStmt != nil {
		if cerr := q.getNotUpdateStockCodeListStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getNotUpdateStockCodeListStmt: %w", cerr)
		}
	}
	if q.getNotUpdateStockListStmt != nil {
		if cerr := q.getNotUpdateStockListStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getNotUpdateStockListStmt: %w", cerr)
		}
	}
	if q.getStockCountStmt != nil {
		if cerr := q.getStockCountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getStockCountStmt: %w", cerr)
		}
	}
	if q.getStockListStmt != nil {
		if cerr := q.getStockListStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getStockListStmt: %w", cerr)
		}
	}
	if q.getUserByIdStmt != nil {
		if cerr := q.getUserByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByIdStmt: %w", cerr)
		}
	}
	if q.getUserByNameStmt != nil {
		if cerr := q.getUserByNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByNameStmt: %w", cerr)
		}
	}
	if q.updateUserStmt != nil {
		if cerr := q.updateUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                            DBTX
	tx                            *sql.Tx
	createStockStmt               *sql.Stmt
	createUserStmt                *sql.Stmt
	getNotUpdateStockCodeListStmt *sql.Stmt
	getNotUpdateStockListStmt     *sql.Stmt
	getStockCountStmt             *sql.Stmt
	getStockListStmt              *sql.Stmt
	getUserByIdStmt               *sql.Stmt
	getUserByNameStmt             *sql.Stmt
	updateUserStmt                *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                            tx,
		tx:                            tx,
		createStockStmt:               q.createStockStmt,
		createUserStmt:                q.createUserStmt,
		getNotUpdateStockCodeListStmt: q.getNotUpdateStockCodeListStmt,
		getNotUpdateStockListStmt:     q.getNotUpdateStockListStmt,
		getStockCountStmt:             q.getStockCountStmt,
		getStockListStmt:              q.getStockListStmt,
		getUserByIdStmt:               q.getUserByIdStmt,
		getUserByNameStmt:             q.getUserByNameStmt,
		updateUserStmt:                q.updateUserStmt,
	}
}
