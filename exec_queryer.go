package dbtyp

import (
	"context"
	"database/sql"

	"github.com/winebarrel/dbtyp/iface"
)

var _ iface.ExecQueryer = &ExecQueryer[struct{}]{}

type ExecQueryer[T any] struct {
	i iface.ExecQueryer
}

// Interface implement

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

// Type converter

func (v *ExecQueryer[T]) Execer() *Execer[T] {
	return &Execer[T]{i: v.i}
}

func (v *ExecQueryer[T]) Queryer() *Queryer[T] {
	return &Queryer[T]{i: v.i}
}
