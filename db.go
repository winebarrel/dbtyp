package dbtyp

import (
	"context"
	"database/sql"

	"github.com/winebarrel/dbtyp/iface"
)

type DB[T any] struct {
	iface.DB
}

// Type converter

func (v *DB[T]) ExecQueryer() *ExecQueryer[T] {
	return &ExecQueryer[T]{i: v}
}

func (v *DB[T]) Execer() *Execer[T] {
	return &Execer[T]{i: v}
}

func (v *DB[T]) Queryer() *Queryer[T] {
	return &Queryer[T]{i: v}
}

// Typed begin

func (v *DB[T]) BeginT() (*Tx[T], error) {
	tx, err := v.Begin()

	if err != nil {
		return nil, err
	}

	return &Tx[T]{Tx: tx}, nil
}

func (v *DB[T]) BeginTxT(ctx context.Context, opts *sql.TxOptions) (*Tx[T], error) {
	tx, err := v.BeginTx(ctx, opts)

	if err != nil {
		return nil, err
	}

	return &Tx[T]{Tx: tx}, nil
}

// Typed prepare

func (v *DB[T]) PrepareT(query string) (*Stmt[T], error) {
	stmt, err := v.Prepare(query)

	if err != nil {
		return nil, err
	}

	return &Stmt[T]{Stmt: stmt}, nil
}

func (v *DB[T]) PrepareContextT(ctx context.Context, query string) (*Stmt[T], error) {
	stmt, err := v.PrepareContext(ctx, query)

	if err != nil {
		return nil, err
	}

	return &Stmt[T]{Stmt: stmt}, nil
}
