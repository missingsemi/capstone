package model

type Machine struct {
	TitleName string
	Name      string
	Id        string
	Count     int
}

var Machines = []Machine{
	{
		TitleName: "CNC Router",
		Name:      "cnc router",
		Id:        "cnc_router",
		Count:     1,
	},
	{
		TitleName: "Laser Engraver",
		Name:      "laser engraver",
		Id:        "laser_engraver",
		Count:     1,
	},
	{
		TitleName: "3D Printer",
		Name:      "3d printer",
		Id:        "3d_printer",
		Count:     4,
	},
	{
		TitleName: "Resin Printer",
		Name:      "resin printer",
		Id:        "resin_printer",
		Count:     3,
	},
}

func MachineFromId(id string) *Machine {
	for _, machine := range Machines {
		if machine.Id == id {
			return &machine
		}
	}

	return nil
}
