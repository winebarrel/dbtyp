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

func TestTxCommit(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	db, err := dbtyp.New2[AliceDB](sql.Open("sqlite", "file::memory:"))
	require.NoError(err)

	_, err = db.Exec("create table foo (id int)")
	require.NoError(err)

	tx, err := db.BeginT()
	var _ iface.Tx = tx
	require.NoError(err)
	_, err = tx.Exec("insert into foo values (100)")
	require.NoError(err)

	err = tx.Commit()
	require.NoError(err)

	var n int
	err = db.QueryRow("select id from foo").Scan(&n)
	require.NoError(err)
	assert.Equal(100, n)
}

func TestTxRollback(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	db, err := dbtyp.New2[AliceDB](sql.Open("sqlite", "file::memory:"))
	require.NoError(err)

	ctx := t.Context()

	_, err = db.ExecContext(ctx, "create table foo (id int)")
	require.NoError(err)

	tx, err := db.BeginTxT(ctx, nil)
	var _ iface.Tx = tx
	require.NoError(err)
	_, err = tx.ExecContext(ctx, "insert into foo values (100)")
	require.NoError(err)

	err = tx.Rollback()
	require.NoError(err)

	n := -1
	err = db.QueryRowContext(ctx, "select count(*) from foo").Scan(&n)
	require.NoError(err)
	assert.Equal(0, n)
}
