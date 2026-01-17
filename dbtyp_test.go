package dbtyp_test

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/winebarrel/dbtyp"
	"github.com/winebarrel/dbtyp/iface"
	_ "modernc.org/sqlite"
)

func TestQueryer(t *testing.T) {
	type AliceDB struct{}

	assert := assert.New(t)
	require := require.New(t)

	db0, err := sql.Open("sqlite", "file::memory:")
	require.NoError(err)
	db := dbtyp.New[AliceDB](db0)

	var n int
	err = db.QueryRow("select 1").Scan(&n)
	require.NoError(err)
	assert.Equal(1, n)

	n = 0
	queryer := db.Queryer()
	var _ iface.Queryer = queryer
	queryer.QueryRow("select 1").Scan(&n)
	require.NoError(err)
	assert.Equal(1, n)
}

func TestExecer(t *testing.T) {
	type AliceDB struct{}

	assert := assert.New(t)
	require := require.New(t)

	db0, err := sql.Open("sqlite", "file::memory:")
	require.NoError(err)
	db := dbtyp.New[AliceDB](db0)

	_, err = db.Exec("create table foo (id int)")
	require.NoError(err)

	execer := db.Execer()
	var _ iface.Execer = execer
	execer.Exec("insert into foo values (1)")

	var n int
	db.QueryRow("select 1").Scan(&n)
	require.NoError(err)
	assert.Equal(1, n)
}

func TestExecQueryer(t *testing.T) {
	type AliceDB struct{}

	assert := assert.New(t)
	require := require.New(t)

	db0, err := sql.Open("sqlite", "file::memory:")
	require.NoError(err)
	db := dbtyp.New[AliceDB](db0)

	_, err = db.Exec("create table foo (id int)")
	require.NoError(err)

	execQueryer := db.ExecQueryer()
	var _ iface.ExecQueryer = execQueryer
	var _ iface.Execer = execQueryer
	var _ iface.Queryer = execQueryer
	execQueryer.Exec("insert into foo values (1)")

	var n int
	execQueryer.QueryRow("select 1").Scan(&n)
	require.NoError(err)
	assert.Equal(1, n)
}

func TestNew2(t *testing.T) {
	type AliceDB struct{}

	assert := assert.New(t)
	require := require.New(t)

	db, err := dbtyp.New2[AliceDB](sql.Open("sqlite", "file::memory:"))

	var n int
	err = db.QueryRow("select 1").Scan(&n)
	require.NoError(err)
	assert.Equal(1, n)
}
