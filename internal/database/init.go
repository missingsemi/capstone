package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB = nil

// Connects to the specified database file
func DbInit(file string, init bool) {
	if db != nil {
		return
	}

	var dsn string
	if init {
		dsn = fmt.Sprintf("file:%s?mode=rwc", file)
	} else {
		dsn = fmt.Sprintf("file:%s?mode=rw", file)
	}
	db_, err := sql.Open("sqlite3", dsn)
	if err != nil {
		log.Fatalf("Failed to open %s. If this is your first time running the program, use the --init flag.", file)
	}

	if init {
		db_.Exec(`
			-- table containing information about every machine available
			CREATE TABLE machine (
				id TEXT PRIMARY KEY,
				name TEXT NOT NULL,
				title_name TEXT NOT NULL,
				count INTEGER NOT NULL
			);

			-- table containing fully created sessions
			CREATE TABLE schedule (
				id INTEGER PRIMARY KEY,
				user_id TEXT NOT NULL,
				group_ids TEXT NOT NULL,
				machine_id TEXT NOT NULL,
				reason TEXT NOT NULL,
				duration INTEGER NOT NULL,
				time DATETIME NOT NULL,
				stage INTEGER NOT NULL DEFAULT 0,
				FOREIGN KEY(machine_id) REFERENCES machine(id) ON DELETE CASCADE ON UPDATE CASCADE
			);

			-- table with a list of admins
			CREATE TABLE admins (
				id INTEGER PRIMARY KEY,
				user_id TEXT NOT NULL
			);
		`)
	}

	db = db_
}

func DbDeinit() {
	if db != nil {
		db.Close()
	}
}
