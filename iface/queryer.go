package iface

import (
	"context"
	"database/sql"
)

var _ Queryer = &sql.DB{}
var _ Queryer = &sql.Tx{}

type Queryer interface {
	Query(query string, args ...any) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}
