package iface

import (
	"database/sql"
)

var _ ExecQueryer = &sql.DB{}
var _ ExecQueryer = &sql.Tx{}

type ExecQueryer interface {
	Execer
	Queryer
}
