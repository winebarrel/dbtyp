package dbtyp

import (
	"github.com/kanmu/dbtyp/iface"
)

func New[T any](v iface.DB) *DB[T] {
	return &DB[T]{DB: v}
}

func New2[T any](v iface.DB, err error) (*DB[T], error) {
	if err != nil {
		return nil, err
	}

	return New[T](v), nil
}
