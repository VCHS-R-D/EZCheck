package machines

import "main/components/postgresmanager"

type Machine struct {
	ID    string `json:"machine_id"`
	Name  string `json:"machine_name"`
	InUSE bool `json:"in_use"`
}

func CreateMachine(id, name string) {
	machine := Machine{ID: id, Name: name, InUSE: false}

	postgresmanager.Save(&machine)
}

func ReadMachines() []Machine {
	var machines []Machine

	postgresmanager.QueryAll(&machines)

	return machines
}

func DeleteMachine(id string) error {
	err := postgresmanager.Delete(Machine{ID: id})

	if err != nil {
		return err
	}
	return err
}
