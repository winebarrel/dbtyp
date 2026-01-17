# dbtyp

[![CI](https://github.com/winebarrel/dbtyp/actions/workflows/ci.yml/badge.svg)](https://github.com/winebarrel/dbtyp/actions/workflows/ci.yml)

dbtyp is a library that associates types with `*sql.DB`.

## Installation

```sh
go get github.com/winebarrel/dbtyp
```

## Usage

```go
package main

import (
	"database/sql"
	"fmt"

	"github.com/winebarrel/dbtyp"
	"github.com/winebarrel/dbtyp/types"
	_ "modernc.org/sqlite"
)

type AliceDB struct{}
type BobDB struct{}

type MyDB interface {
	AliceDB | BobDB
}

func OpenDB[T MyDB]() (*types.DB[T], error) {
	return dbtyp.New2[T](sql.Open("sqlite", "file::memory:"))
}

func main() {
	alice, err := OpenDB[AliceDB]()
	if err != nil {
		panic(err)
	}
	alice.Exec("create table foo (id int)")

	bob, err := OpenDB[BobDB]()
	if err != nil {
		panic(err)
	}
	bob.Exec("create table bar (id int)")
	// bob = alice // COMPILE ERROR!

	procForAlice(alice.ExecQueryer())
	// procForAlice(bob.ExecQueryer()) // COMPILE ERROR!

	procForBob(bob.Queryer())
	// procForBob(alice.Queryer()) // COMPILE ERROR!
}

func procForAlice(eq *types.ExecQueryer[AliceDB]) {
	_, err := eq.Exec("insert into foo values (1)")
	if err != nil {
		panic(err)
	}

	var n int
	err = eq.QueryRow("select count(*) from foo").Scan(&n)
	if err != nil {
		panic(err)
	}
	fmt.Println("foo rows count:", n)
}

func procForBob(q *types.Queryer[BobDB]) {
	var n int
	err := q.QueryRow("select count(*) from bar").Scan(&n)
	if err != nil {
		panic(err)
	}
	fmt.Println("bar rows count:", n)
}
```
