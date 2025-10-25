package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // or postgres driver
)

func Init(dbFilePath string) *sql.DB {
	sqliteDB, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}

	if err := sqliteDB.Ping(); err != nil {
		log.Fatalf("unable to ping DB, likely the sqliteDB doesn't exist: %v", err)
	}

	log.Println("Database successfully connected!")
	return sqliteDB
}
