package machines

import "main/components/postgresmanager"

const (
	PowerOn int = iota
	PowerOff
)

type Action struct {
	ID int `json:"action_id"`
}

func AddAction(id string, actionInt ...int) error {
	machine := Machine{ID: id}
	err := postgresmanager.Query(&machine, &machine)
	if err != nil {
		return err
	}

	for _, action := range actionInt {
		machine.Actions = append(machine.Actions, Action{ID: action})
	}
	
	return postgresmanager.Update(machine, &machine)
}

func DeleteAction(id string, actionInt int) error {
	machine := Machine{ID: id}
	err := postgresmanager.Query(&machine, &machine)
	if err != nil {
		return err
	}
	for i, action := range machine.Actions {
		if action.ID == actionInt {
			machine.Actions = append(machine.Actions[:i], machine.Actions[i+1:]...)
			break
		}
	}
	return postgresmanager.Update(machine, &machine)
}
