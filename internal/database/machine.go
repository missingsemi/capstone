package database

import (
	"github.com/missingsemi/capstone/internal/model"
)

func GetMachines() ([]model.Machine, error) {
	rows, err := db.Query("SELECT * FROM machine;")
	if err != nil {
		return []model.Machine{}, err
	}
	defer rows.Close()

	machines := make([]model.Machine, 0)

	for rows.Next() {
		machine := model.Machine{}
		err := rows.Scan(&machine.Id, &machine.Name, &machine.TitleName, &machine.Count)
		if err != nil {
			return machines, err
		}
		machines = append(machines, machine)
	}

	return machines, nil
}

func GetMachineById(id string) (model.Machine, error) {
	stmt, err := db.Prepare("SELECT * FROM machine WHERE id = ?;")
	if err != nil {
		return model.Machine{}, err
	}
	defer stmt.Close()

	machine := model.Machine{}
	err = stmt.QueryRow(id).Scan(&machine.Id, &machine.Name, &machine.TitleName, &machine.Count)
	return machine, err
}

func CreateMachine(machine model.Machine) error {
	stmt, err := db.Prepare("INSERT INTO machine (id, name, title_name, count) VALUES (?, ?, ?, ?);")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(machine.Id, machine.Name, machine.TitleName, machine.Count)
	return err
}

func ModifyMachine(id string, machine model.Machine) error {
	stmt, err := db.Prepare("UPDATE machine SET id = ?, name = ?, title_name = ?, count = ? WHERE id = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(machine.Id, machine.Name, machine.TitleName, machine.Count, id)
	return err
}

func DeleteMachine(id string) error {
	stmt, err := db.Prepare("DELETE FROM machine WHERE id = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}
