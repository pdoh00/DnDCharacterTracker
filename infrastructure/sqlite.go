package infrastructure

import (
	"fmt"

	"github.com/pdoh00/dndAdventuresLeague/interfaces"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// SqliteHandler exposes the sqlite3 connection
type SqliteHandler struct {
	Conn *sql.DB
}

//SqliteRow is Sqlite implementation fo the Row interface
type SqliteRow struct {
	Rows *sql.Rows
}

// Scan reads values into given paramaters
//e.g. Scan(&id, &some_field1, &some_field2)
func (r *SqliteRow) Scan(dest ...interface{}) {
	r.Rows.Scan(dest)
}

// Next iterates to the next row result.
// Returns true when row exists false otherwise
func (r *SqliteRow) Next() bool {
	return r.Rows.Next()
}

// Execute performs the sql statement in a transaction
func (handler *SqliteHandler) Execute(statement string) (r sql.Result, e error) {
	tx, err := handler.Conn.Begin()
	if err != nil {
		panic(err)
	}
	stmt, err := tx.Prepare(statement)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	res, err := stmt.Exec()
	if err != nil {
		panic(err)
	}

	tx.Commit()
	return res, err
}

// Query performs the given statement on the DB
func (handler *SqliteHandler) Query(statement string) interfaces.Row {
	rows, err := handler.Conn.Query(statement)
	if err != nil {
		fmt.Println(err)
		return new(SqliteRow)
	}
	row := new(SqliteRow)
	row.Rows = rows
	return row
}

// NewSqliteHandler creates a new SqliteHandler
func NewSqliteHandler(dbFileName string) *SqliteHandler {
	conn, err := sql.Open("sqlite3", dbFileName)
	if err != nil {
		panic(err)
	}
	sqliteHandler := new(SqliteHandler)
	sqliteHandler.Conn = conn
	return sqliteHandler
}
