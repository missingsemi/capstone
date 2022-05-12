package database

func IsUserAdmin(id string) (bool, error) {
	stmt, err := db.Prepare("SELECT EXISTS(SELECT * FROM admins WHERE user_id=?);")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	result := 0
	err = stmt.QueryRow(id).Scan(&result)
	return result == 1, err
}

func CreateAdmin(id string) error {
	stmt, err := db.Prepare("INSERT INTO admins (user_id) VALUES (?);")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	return err
}
