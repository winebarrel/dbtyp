package dbtyp

import (
	"github.com/kanmu/dbtyp/iface"
)

type Stmt[T any] struct {
	iface.Stmt
}
