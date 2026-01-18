package dbtyp

import (
	"database/sql"

	"github.com/winebarrel/dbtyp/iface"
)

var _ iface.Stmt = &Stmt[struct{}]{}

type Stmt[T any] struct {
	*sql.Stmt
}
