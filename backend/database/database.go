package database

import (
	"database/sql"
	"fmt"
	"os"
)

func Conn() *sql.DB {
	// todo env должно быть только в main.go
	dbURL := os.Getenv("GO_DB_URL")
	dbconn, err := sql.Open("pgx", dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return dbconn
}
