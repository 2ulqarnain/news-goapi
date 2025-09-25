package repository

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // or postgres driver
)

var db *sql.DB

func InitDB(dbFilePath string) {
	var err error
	db, err = sql.Open("sqlite3", dbFilePath)
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Unable to ping DB: %v", err)
	}

	log.Println("Database successfully connected!")
}

func Close() error {
	return db.Close()
}
