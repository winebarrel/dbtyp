package dbtyp

import (
	"database/sql"

	"github.com/winebarrel/dbtyp/types"
)

func New[T any](v *sql.DB) *types.DB[T] {
	return &types.DB[T]{DB: v}
}
