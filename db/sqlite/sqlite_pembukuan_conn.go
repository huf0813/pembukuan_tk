package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type ConnSqlite struct{}

type ConnSqliteInterface interface {
	SqliteConn() *sql.DB
}

func (sqliteConn *ConnSqlite) SqliteConn() *sql.DB {
	result, err := sql.Open("sqlite3", "./db/sqlite/pembukuan_db")
	if err != nil {
		panic(err)
	}
	return result
}
