package server

import (
	"database/sql"
)

// setupDatabase will be called when creating a new server.  This should return
// the *sql.DB that will be used for every request.  Note, the server will
// panic if you return a non-nil error.
func setupDatabase() (*sql.DB, error) {
	// Insert database initialization code here.  For example, you can use a
	// MySQL database by including the following import:
	//
	// 	_ "github.com/go-sql-driver/mysql"
	//
	// And then by replacing the body of this function with
	//
	// 	return sql.Open("mysql", "")
	return nil, nil
}
