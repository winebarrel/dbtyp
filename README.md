# dbtyp

[![CI](https://github.com/winebarrel/dbtyp/actions/workflows/ci.yml/badge.svg)](https://github.com/winebarrel/dbtyp/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/winebarrel/dbtyp.svg)](https://pkg.go.dev/github.com/winebarrel/dbtyp)
[![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/winebarrel/dbtyp)](https://pkg.go.dev/github.com/winebarrel/dbtyp?tab=versions)
[![Go Report Card](https://goreportcard.com/badge/github.com/winebarrel/dbtyp)](https://goreportcard.com/report/github.com/winebarrel/dbtyp)

dbtyp is a library that associates types with `*sql.DB`.

[*dbtyp.DB[T]](https://pkg.go.dev/github.com/winebarrel/dbtyp#DB) and [*dbtyp.Tx[T]](https://pkg.go.dev/github.com/winebarrel/dbtyp#Tx) have the same methods as [*sql.DB](https://pkg.go.dev/database/sql#DB) and [*sql.Tx](https://pkg.go.dev/database/sql#Tx), but it can define different types with the same interface using [generics](https://go.dev/doc/tutorial/generics).

Additionally, it can generate instances of restricted types [*dbtyp.ExecQueryer[T]](https://pkg.go.dev/github.com/winebarrel/dbtyp#ExecQueryer), [*dbtyp.Execer[T]](https://pkg.go.dev/github.com/winebarrel/dbtyp#Execer), and [*dbtyp.Queryer[T]](https://pkg.go.dev/github.com/winebarrel/dbtyp#Queryer).

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
	_ "modernc.org/sqlite"
)

type AliceDB struct{}
type BobDB struct{}

type MyDB interface {
	AliceDB | BobDB
}

func main() {
	aliceDB := openDB[AliceDB]()
	bobDB := openDB[BobDB]()

	// bob = alice // COMPILE ERROR!

	createTable(aliceDB, "foo")
	createTable(bobDB, "bar")

	procForAlice(aliceDB.ExecQueryer())
	// procForAlice(bob.ExecQueryer()) // COMPILE ERROR!

	procForBob(bobDB.Queryer())
	// procForBob(alice.Queryer()) // COMPILE ERROR!
}

func openDB[T MyDB]() *dbtyp.DB[T] {
	db, err := dbtyp.New2[T](sql.Open("sqlite", "file::memory:"))

	if err != nil {
		panic(err)
	}

	return db
}

func createTable[T MyDB](db *dbtyp.DB[T], name string) {
	_, err := db.Exec("create table " + name + " (id int)")

	if err != nil {
		panic(err)
	}
}

func procForAlice(eq *dbtyp.ExecQueryer[AliceDB]) {
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

func procForBob(q *dbtyp.Queryer[BobDB]) {
	var n int
	err := q.QueryRow("select count(*) from bar").Scan(&n)
	if err != nil {
		panic(err)
	}
	fmt.Println("bar rows count:", n)
}
```
