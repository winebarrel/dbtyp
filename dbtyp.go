package dbtyp

import (
	"database/sql"

	"github.com/winebarrel/dbtyp/types"
)

func New[T any](v *sql.DB) *types.DB[T] {
	return &types.DB[T]{DB: v}
}

func New2[T any](v *sql.DB, err error) (*types.DB[T], error) {
	if err != nil {
		return nil, err
	}

	return New[T](v), nil
}
