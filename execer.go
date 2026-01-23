package dbtyp

import (
	"context"
	"database/sql"

	"github.com/kanmu/dbtyp/iface"
)

var _ iface.Execer = &Execer[struct{}]{}

type Execer[T any] struct {
	i iface.Execer
}

// Interface implement

func (v *Execer[T]) Exec(query string, args ...any) (sql.Result, error) {
	return v.i.Exec(query, args...)
}

func (v *Execer[T]) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return v.i.ExecContext(ctx, query, args...)
}
