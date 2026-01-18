package dbtyp

import (
	"context"
	"database/sql"

	"github.com/winebarrel/dbtyp/iface"
)

var _ iface.Tx = &Tx[struct{}]{}

type Tx[T any] struct {
	*sql.Tx
}

// Type converter

func (v *Tx[T]) ExecQueryer() *ExecQueryer[T] {
	return &ExecQueryer[T]{i: v}
}

func (v *Tx[T]) Execer() *Execer[T] {
	return &Execer[T]{i: v}
}

func (v *Tx[T]) Queryer() *Queryer[T] {
	return &Queryer[T]{i: v}
}

// Typed prepare

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
