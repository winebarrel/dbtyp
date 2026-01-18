package iface

import (
	"context"
	"database/sql"
)

var _ Stmt = &sql.Stmt{}

type Stmt interface {
	Close() error
	Exec(args ...any) (sql.Result, error)
	ExecContext(ctx context.Context, args ...any) (sql.Result, error)
	Query(args ...any) (*sql.Rows, error)
	QueryContext(ctx context.Context, args ...any) (*sql.Rows, error)
	QueryRow(args ...any) *sql.Row
	QueryRowContext(ctx context.Context, args ...any) *sql.Row
}
