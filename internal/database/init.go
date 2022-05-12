package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB = nil

// Connects to the specified database file
func DbInit(file string) {
	if db != nil {
		return
	}

	dsn := fmt.Sprintf("file:%s?mode=rwc", file)
	db_, err := sql.Open("sqlite3", dsn)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	db = db_
}

func DbDeinit() {
	if db != nil {
		db.Close()
	}
}
