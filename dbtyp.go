package dbtyp

import (
	"database/sql"
)

func New[T any](v *sql.DB) *DB[T] {
	return &DB[T]{DB: v}
}

func New2[T any](v *sql.DB, err error) (*DB[T], error) {
	if err != nil {
		return nil, err
	}

	return New[T](v), nil
}
