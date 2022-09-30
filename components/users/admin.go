package users

import (
	"main/components/internal"
	"main/components/machines"
	"main/components/postgresmanager"
)

type Admin struct {
	ID        string `json:"-"`
	Username  string `json:"username"`
	Password  string `json:"-"`
	FirstName string `json:"first"`
	LastName  string `json:"last"`
}

func CreateAdmin(username, password, firstName, lastName string) error {
	admin := Admin{ID: internal.GenerateUUID(), Username: username, Password: password, FirstName: firstName, LastName: lastName}
	return postgresmanager.Save(&admin)
}

func ReadAdmin(id string) Admin {
	var admin Admin
	postgresmanager.Query(Admin{ID: id}, &admin)

	return admin
}

func CertifyUser(username, machineID string) error {
	var user User
	var machine machines.Machine
	err := postgresmanager.Query(User{Username: username}, &user)
	if err != nil {
		return err
	}

	err = postgresmanager.Query(machines.Machine{ID: machineID}, &machine)

	if err != nil {
		return err
	}

	err = postgresmanager.CreateAssociation(&user, "Machines", &machine)

	return err
}
