package machines

import "main/components/postgresmanager"

type Machine struct {
	ID    string `json:"machine_id" gorm:"primaryKey"`
	Name  string `json:"machine_name" gorm:"uniqueIndex"`
	InUSE bool `json:"in_use" gorm:"index"`
	Students  []*users.User `gorm:"many2many:users_machines"`
	CreatedAt    time.Time  `json:"-" gorm:"index"`
	UpdatedAt    time.Time  `json:"-" gorm:"index"`
}

func CreateMachine(id, name string) error {
	machine := Machine{ID: id, Name: name, InUSE: false}

	return postgresmanager.Save(&machine)
}

func ReadMachines() []Machine {
	var machines []Machine

	postgresmanager.QueryAll(&machines)

	return machines
}

func SignOut(id string) error {
	postgresmanager.Update(ID: id, InUSE: false)
}

func DeleteMachine(id string) error {
	return postgresmanager.Delete(Machine{ID: id})
}
