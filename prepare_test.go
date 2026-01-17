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

func TestDBPrepare(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	db, err := dbtyp.New2[AliceDB](sql.Open("sqlite", "file::memory:"))
	require.NoError(err)

	_, err = db.Exec("create table foo (id int)")
	require.NoError(err)

	stmt, err := db.PrepareT("insert into foo values (?)")
	var _ iface.Stmt = stmt
	require.NoError(err)

	_, err = stmt.Exec(100)
	require.NoError(err)

	var n int
	err = db.QueryRow("select id from foo").Scan(&n)
	require.NoError(err)
	assert.Equal(100, n)
}

func TestTxPrepare(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	db, err := dbtyp.New2[AliceDB](sql.Open("sqlite", "file::memory:"))
	require.NoError(err)

	ctx := t.Context()

	_, err = db.ExecContext(ctx, "create table foo (id int)")
	require.NoError(err)

	tx, err := db.BeginTxT(ctx, nil)
	require.NoError(err)

	stmt, err := tx.PrepareContextT(ctx, "insert into foo values (?)")
	var _ iface.Stmt = stmt
	require.NoError(err)
	_, err = stmt.ExecContext(ctx, 100)
	require.NoError(err)

	err = tx.Commit()
	require.NoError(err)

	var n int
	err = db.QueryRowContext(ctx, "select id from foo").Scan(&n)
	require.NoError(err)
	assert.Equal(100, n)
}
