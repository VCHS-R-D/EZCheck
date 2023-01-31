package machines

import (
	"main/components/postgresmanager"
	"strconv"
)

type Action struct {
	ActionID int `json:"action_id"`
}

func AddAction(machineID string, actionID string) error {
	machine := Machine{ID: machineID}
	err := postgresmanager.Query(&machine, &machine)
	if err != nil {
		return err
	}

	actionInt, err := strconv.Atoi(actionID)
	if err != nil {
		return err
	}

	machine.Actions = append(machine.Actions, Action{ActionID: actionInt})

	return postgresmanager.Update(machine, &machine)
}

func DeleteAction(machineID string, actionID string) error {
	machine := Machine{ID: machineID}
	err := postgresmanager.Query(&machine, &machine)
	if err != nil {
		return err
	}

	actionInt, err := strconv.Atoi(actionID)
	if err != nil {
		return err
	}

	for i, action := range machine.Actions {
		if action.ActionID == actionInt {
			machine.Actions = append(machine.Actions[:i], machine.Actions[i+1:]...)
			break
		}
	}
	return postgresmanager.Update(machine, &machine)
}
