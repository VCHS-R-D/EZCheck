package users

import (
	"fmt"
	"main/components/internal"
	"main/components/log"
	"main/components/machines"
	"main/components/postgresmanager"
	"time"
)

type User struct {
	ID        string              `json:"-" gorm:"primaryKey"`
	Username  string              `json:"username" gorm:"uniqueIndex"`
	Password  string              `json:"password"`
	FirstName string              `json:"first" gorm:"index"`
	LastName  string              `json:"last" gorm:"index"`
	Grade     string              `json:"grade" gorm:"index"`
	Code      string              `json:"code" gorm:"uniqueIndex"`
	Machines  []*machines.Machine `gorm:"many2many:users_machines"`
	CreatedAt time.Time           `json:"-" gorm:"index"`
	UpdatedAt time.Time           `json:"-" gorm:"index"`
}

func CreateUser(username, password, firstName, lastName, grade, code string) error {

	var u *User
	err := postgresmanager.Query(&User{Code: code}, &u)

	if checkErr != "record not found" {
		return error
	}

	user := User{ID: internal.GenerateUUID(), Username: username, Password: password, FirstName: firstName, LastName: lastName, Grade: grade, Code: code}
	if postgresmanager.Save(&user) != nil {
		return "", postgresmanager.Save(&user)
	}

	return err
}

func GetUser(id string) (error, User) {
	var user User
	var machines []*machines.Machine

	if err := postgresmanager.Query(User{ID: id}, &user); err != nil {
		return err, User{}
	} else if err := postgresmanager.ReadAssociation(&user, "Machines", &machines); err != nil {
		return err, User{}
	}

	user.Machines = machines
	user.Password = ""

	return nil, user
}

func Authenticate(code, machineID string) string {
	var user *User

	if postgresmanager.Query(&User{Code: code}, &user) != nil {
		return "{\"error\": \"user not found\"}"
	}

	var machines []*machines.Machine
	err := postgresmanager.ReadAssociation(&user, "Machines", &machines)

	if err != nil {
		return "{\"error\": \"could not read machines\"}"
	}

	for _, machine := range machines {
		if machine.ID == machineID {
			actions, err := machine.SignIn()
			if err != nil {
				return "{\"error\": \"could not sign in\"}"
			}
			log.Log(fmt.Sprintf("%s signed in to machine %s", user.Username, machine.Name))
			return fmt.Sprintf("{\"authorized\": true, \"name\": \"%s %s\", actions: %v}", user.FirstName, user.LastName, actions)
		}
	}

	log.Log(fmt.Sprintf("%s failed to sign in to machine %s", user.Username, machineID))
	return "{\"authorized\": false}"
}

func DeleteUser(id string) error {
	return postgresmanager.Delete(User{ID: id})
}
