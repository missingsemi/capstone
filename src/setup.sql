-- this file contains the commands required to set up the sqlite3 database
-- you can either type it in manually or by piping it into the sqlite3 command

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
  FOREIGN KEY(machine_id) REFERENCES machine(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- table with a list of admins
CREATE TABLE admins (
  id INTEGER PRIMARY KEY,
  user_id TEXT NOT NULL
);