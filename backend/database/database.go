package database

import (
	"database/sql"
	"fmt"
	"os"
)

func Conn() *sql.DB {
	var dbURL = "postgres://postgres:postgres@localhost:6432/bank"
	dbconn, err := sql.Open("pgx", dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return dbconn
}
