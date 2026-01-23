package dbtyp

import (
	"context"
	"database/sql"

	"github.com/kanmu/dbtyp/iface"
)

var _ iface.Queryer = &Queryer[struct{}]{}

type Queryer[T any] struct {
	i iface.Queryer
}

// Interface implement

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
