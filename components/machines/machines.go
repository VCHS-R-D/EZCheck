package machines

import "main/components/postgresmanager"

type Machine struct {
	ID    string
	Name  string
	InUSE bool
}

func CreateMachine(id, name string) {
	machine := Machine{ID: id, Name: name}

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
