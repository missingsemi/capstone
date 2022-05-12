package database

import "log"

// Sets up an empty .db file
func Configure() {
	_, err := db.Exec(`
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

	if err != nil {
		log.Fatalf("Failed to configure new database: %v", err)
	} else {
		log.Println("Successfully configured new database.")
	}
}
