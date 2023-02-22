package machines

import (
	"main/components/postgresmanager"
	"main/components/types"
)

func AddAction(machineID string, actionID string) error {
	var machine types.Machine
	err := postgresmanager.Query(&types.Machine{ID: machineID}, &machine)

	if err != nil {
		return err
	}

	machine.Actions = append(machine.Actions, actionID)

	return postgresmanager.Update(machine, &machine)
}

func DeleteAction(machineID string, actionID string) error {
	var machine types.Machine
	err := postgresmanager.Query(&types.Machine{ID: machineID}, &machine)
	if err != nil {
		return err
	}

	for i, action := range machine.Actions {
		if action == actionID {
			machine.Actions = append(machine.Actions[:i], machine.Actions[i+1:]...)
			break
		}
	}
	return postgresmanager.Update(machine, &machine)
}
