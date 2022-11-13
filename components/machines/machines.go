package machines

import (
	"fmt"
	"main/components/internal"
	"main/components/log"
	"main/components/postgresmanager"
	"time"
)

type Machine struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"uniqueIndex"`
	InUSE     bool      `json:"in_use" gorm:"index"`
	Actions   []Action  `json:"actions" gorm:"foreignKey:ActionID"`
	CreatedAt time.Time `json:"-" gorm:"index"`
	UpdatedAt time.Time `json:"-" gorm:"index"`
}

func CreateMachine(name string) error {
	actions := make([]Action, 0)
	machine := Machine{ID: internal.GenerateUUID(), Name: name, InUSE: false, Actions: actions}
	return postgresmanager.Save(&machine)
}

func ReadMachines() []Machine {
	var machines []Machine

	postgresmanager.QueryAll(&machines)

	return machines
}

func (m *Machine) SignIn() ([]Action, error) {
	return m.Actions, postgresmanager.Update(m, &Machine{InUSE: true})
}

func SignOut(machineID string) error {

	var m *Machine
	err := postgresmanager.Query(Machine{ID: machineID}, &m)
	if err != nil {
		return err
	}

	log.Log(fmt.Sprintf("User signed out of machine: %s", m.Name))
	return postgresmanager.Update(m, &Machine{InUSE: false})
}

func DeleteMachine(id string) error {
	return postgresmanager.Delete(Machine{ID: id})
}
