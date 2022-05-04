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
