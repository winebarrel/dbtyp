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
