package iface

import (
	"context"
	"database/sql"
)

var _ Execer = &sql.DB{}
var _ Execer = &sql.Tx{}

type Execer interface {
	Exec(query string, args ...any) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}
