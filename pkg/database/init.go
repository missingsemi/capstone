package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB = nil

func DbInit() {
	if db != nil {
		return
	}
	_db, err := sql.Open("sqlite3", "file:schedule.db?mode=rwc")
	if err != nil {
		log.Fatal(err)
	}
	db = _db
}

func DbDeinit() {
	db.Close()
}
