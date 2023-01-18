package machines

import (
	"fmt"
	"main/components/log"
	"main/components/postgresmanager"
	"time"
)

type Machine struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	InUSE     bool      `json:"in_use" gorm:"index"`
	Actions   []Action  `json:"actions" gorm:"foreignKey:ActionID"`
	CreatedAt time.Time `json:"-" gorm:"index"`
	UpdatedAt time.Time `json:"-" gorm:"index"`
}

func CreateMachine(id string) error {
	actions := make([]Action, 0)
	machine := Machine{ID: id, InUSE: false, Actions: actions}
	return postgresmanager.Save(&machine)
}

func GetMachines() []Machine {
	var machines []Machine

	postgresmanager.QueryAll(&machines)

	return machines
}

func (m *Machine) SignIn() ([]Action, error) {
	return m.Actions, postgresmanager.Update(m, &Machine{InUSE: true})
}

func SignOut(name, machineID string) error {
	var m *Machine
	err := postgresmanager.Query(Machine{ID: machineID}, &m)
	if err != nil {
		return err
	}

	log.Log(fmt.Sprintf("%s signed out of machine: %s", name, m.ID))
	return postgresmanager.Update(m, &Machine{InUSE: false})
}

func DeleteMachine(id string) error {
	return postgresmanager.Delete(Machine{ID: id})
}
