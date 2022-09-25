package users

import (
	"main/components/internal"
	"main/components/machines"
	"main/components/postgresmanager"
)


type User struct {
	ID              string
	Username        string
	Password        string `json:"-"`
	FirstName       string
	LastName        string
	Machines        []*machines.Machine `gorm:"many2many:users_machines"`
}


func CreateUser(username, password, firstName, lastName, code string) error {

	user := User{ID: internal.GenerateUUID(), Username: username, Password: password, FirstName: firstName, LastName: lastName}
	return postgresmanager.Save(&user)
}

func ReadUser(id string) User {
	var user User
	postgresmanager.Query(User{ID: id}, &user)

	return user
}
