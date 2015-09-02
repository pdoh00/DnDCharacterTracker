package interfaces

import "database/sql"

// DBHandler is a generic interface for database operations
type DBHandler interface {
	Execute(statement string) (r sql.Result, e error)
	Query(statement string) Row
}

// Row is a generic interface for the result of a DBHandler operation
type Row interface {
	Scan(dest ...interface{})
	Next() bool
}

//DBRepo is the generic definition of a database repository
type DBRepo struct {
	dbHandlers map[string]DBHandler
	dbHandler  DBHandler
}
