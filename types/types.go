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
	return &ExecQueryer[T]{i: v}
}

func (v *DB[T]) Execer() *Execer[T] {
	return &Execer[T]{i: v}
}

func (v *DB[T]) Queryer() *Queryer[T] {
	return &Queryer[T]{i: v}
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

/////////////////////////////////////////////////////////////////////
// Tx
/////////////////////////////////////////////////////////////////////

var _ iface.Tx = &Tx[struct{}]{}

type Tx[T any] struct {
	*sql.Tx
}

func (v *Tx[T]) ExecQueryer() *ExecQueryer[T] {
	return &ExecQueryer[T]{i: v}
}

func (v *Tx[T]) Execer() *Execer[T] {
	return &Execer[T]{i: v}
}

func (v *Tx[T]) Queryer() *Queryer[T] {
	return &Queryer[T]{i: v}
}

func (v *Tx[T]) PrepareT(query string) (*Stmt[T], error) {
	stmt, err := v.Prepare(query)

	if err != nil {
		return nil, err
	}

	return &Stmt[T]{Stmt: stmt}, nil
}

func (v *Tx[T]) PrepareContextT(ctx context.Context, query string) (*Stmt[T], error) {
	stmt, err := v.PrepareContext(ctx, query)

	if err != nil {
		return nil, err
	}

	return &Stmt[T]{Stmt: stmt}, nil
}

/////////////////////////////////////////////////////////////////////
// Stmt
/////////////////////////////////////////////////////////////////////

var _ iface.Stmt = &Stmt[struct{}]{}

type Stmt[T any] struct {
	*sql.Stmt
}

/////////////////////////////////////////////////////////////////////
// ExecQueryer
/////////////////////////////////////////////////////////////////////

var _ iface.ExecQueryer = &ExecQueryer[struct{}]{}

type ExecQueryer[T any] struct {
	i iface.ExecQueryer
}

func (v *ExecQueryer[T]) Exec(query string, args ...any) (sql.Result, error) {
	return v.i.Exec(query, args...)
}

func (v *ExecQueryer[T]) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return v.i.ExecContext(ctx, query, args...)
}

func (v *ExecQueryer[T]) Query(query string, args ...any) (*sql.Rows, error) {
	return v.i.Query(query, args...)
}

func (v *ExecQueryer[T]) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return v.i.QueryContext(ctx, query, args...)
}

func (v *ExecQueryer[T]) QueryRow(query string, args ...any) *sql.Row {
	return v.i.QueryRow(query, args...)
}

func (v *ExecQueryer[T]) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	return v.i.QueryRowContext(ctx, query, args...)
}

/////////////////////////////////////////////////////////////////////
// Execer
/////////////////////////////////////////////////////////////////////

var _ iface.Execer = &Execer[struct{}]{}

type Execer[T any] struct {
	i iface.Execer
}

func (v *Execer[T]) Exec(query string, args ...any) (sql.Result, error) {
	return v.i.Exec(query, args...)
}

func (v *Execer[T]) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return v.i.ExecContext(ctx, query, args...)
}

/////////////////////////////////////////////////////////////////////
// Queryer
/////////////////////////////////////////////////////////////////////

var _ iface.Queryer = &Queryer[struct{}]{}

type Queryer[T any] struct {
	i iface.Queryer
}

func (v *Queryer[T]) Query(query string, args ...any) (*sql.Rows, error) {
	return v.i.Query(query, args...)
}

func (v *Queryer[T]) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return v.i.QueryContext(ctx, query, args...)
}

func (v *Queryer[T]) QueryRow(query string, args ...any) *sql.Row {
	return v.i.QueryRow(query, args...)
}

func (v *Queryer[T]) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	return v.i.QueryRowContext(ctx, query, args...)
}
