package dbtyp

import (
	"github.com/winebarrel/dbtyp/iface"
)

type Stmt[T any] struct {
	iface.Stmt
}
