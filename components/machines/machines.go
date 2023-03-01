package machines

import (
	"fmt"
	"main/components/log"
	"main/components/postgresmanager"
	"main/components/types"

	"github.com/lib/pq"
)

func CreateMachine(id string) error {
	actions := make(pq.StringArray, 0)
	machine := types.Machine{ID: id, InUSE: false, Actions: actions}
	return postgresmanager.Save(&machine)
}

func GetMachines() []*types.Machine {
	var machines []*types.Machine

	postgresmanager.QueryAll(&machines)

	return machines
}

func SignIn(machineID string) ([]string, error) {
	var m *types.Machine
	err := postgresmanager.Query(types.Machine{ID: machineID}, &m)
	if err != nil {
		return nil, err
	}

	return m.Actions, postgresmanager.Update(m, &types.Machine{InUSE: true})
}

func SignOut(name, machineID string) error {
	var m *types.Machine
	err := postgresmanager.Query(types.Machine{ID: machineID}, &m)
	if err != nil {
		return err
	}

	log.Log(fmt.Sprintf("%s signed out of machine: %s", name, m.ID))
	return postgresmanager.Update(m, &types.Machine{InUSE: false})
}

func DeleteMachine(id string) error {
	err := postgresmanager.ClearAssociations(&types.Machine{ID: id}, "Users")
	if err != nil {
		return err
	}

	return postgresmanager.Delete(types.Machine{ID: id})
}
