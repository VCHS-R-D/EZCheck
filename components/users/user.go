package users

import (
	"main/components/internal"
	"main/components/machines"
	"main/components/postgresmanager"
)

type User struct {
	ID        string              `json:"-"`
	Username  string              `json:"username"`
	Password  string              `json:"-"`
	FirstName string              `json:"first"`
	LastName  string              `json:"last"`
	Machines  []*machines.Machine `gorm:"many2many:users_machines"`
}

func CreateUser(username, password, firstName, lastName string) error {

	user := User{ID: internal.GenerateUUID(), Username: username, Password: password, FirstName: firstName, LastName: lastName}
	return postgresmanager.Save(&user)
}

func ReadUser(id string) User {
	var user User
	postgresmanager.Query(User{ID: id}, &user)

	return user
}
