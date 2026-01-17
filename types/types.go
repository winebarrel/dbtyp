package types

import (
	"context"
	"database/sql"

	"github.com/winebarrel/dbtyp/iface"
)

/////////////////////////////////////////////////////////////////////
// DB
/////////////////////////////////////////////////////////////////////

var _ iface.DB = &DB[struct{}]{}

type DB[T any] struct {
	*sql.DB
}

func (v *DB[T]) ExecQueryer() *ExecQueryer[T] {
	return &ExecQueryer[T]{
		ExecQueryer: v,
	}
}

func (v *DB[T]) Execer() *Execer[T] {
	return &Execer[T]{
		Execer: v,
	}
}

func (v *DB[T]) Queryer() *Queryer[T] {
	return &Queryer[T]{
		Queryer: v,
	}
}

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

func (v *DB[T]) PrepareT(query string) (*Stmt[T], error) {
	stmt, err := v.DB.Prepare(query)

	if err != nil {
		return nil, err
	}

	return &Stmt[T]{Stmt: stmt}, nil
}

func (v *DB[T]) PrepareContextT(ctx context.Context, query string) (*Stmt[T], error) {
	stmt, err := v.DB.PrepareContext(ctx, query)

	if err != nil {
		return nil, err
	}

	return &Stmt[T]{Stmt: stmt}, nil
}

/////////////////////////////////////////////////////////////////////
// Tx
/////////////////////////////////////////////////////////////////////

var _ iface.Tx = &Tx[struct{}]{}

type Tx[T any] struct {
	*sql.Tx
}

func (v *Tx[T]) ExecQueryer() *ExecQueryer[T] {
	return &ExecQueryer[T]{
		ExecQueryer: v,
	}
}

func (v *Tx[T]) Execer() *Execer[T] {
	return &Execer[T]{
		Execer: v,
	}
}

func (v *Tx[T]) Queryer() *Queryer[T] {
	return &Queryer[T]{
		Queryer: v,
	}
}

func (v *Tx[T]) PrepareT(query string) (*Stmt[T], error) {
	stmt, err := v.Tx.Prepare(query)

	if err != nil {
		return nil, err
	}

	return &Stmt[T]{Stmt: stmt}, nil
}

func (v *Tx[T]) PrepareContextT(ctx context.Context, query string) (*Stmt[T], error) {
	stmt, err := v.Tx.PrepareContext(ctx, query)

	if err != nil {
		return nil, err
	}

	return &Stmt[T]{Stmt: stmt}, nil
}

/////////////////////////////////////////////////////////////////////
// Misc
/////////////////////////////////////////////////////////////////////

var _ iface.Stmt = &Stmt[struct{}]{}

type Stmt[T any] struct {
	*sql.Stmt
}

type ExecQueryer[T any] struct {
	iface.ExecQueryer
}

type Execer[T any] struct {
	iface.Execer
}

type Queryer[T any] struct {
	iface.Queryer
}
