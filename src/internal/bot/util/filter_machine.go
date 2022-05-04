package util

import "github.com/missingsemi/capstone/pkg/model"

func FilterMachine(id string, machines []model.Machine) model.Machine {
	for _, machine := range machines {
		if machine.Id == id {
			return machine
		}
	}

	return model.Machine{}
}
