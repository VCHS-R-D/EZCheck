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
	ID        string              `json:"id" gorm:"primaryKey"`
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

	if err != nil {
		if err.Error() != "record not found" {
			return err
		}
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	user := User{ID: internal.GenerateUUID(), Username: username, Password: hashedPassword, FirstName: firstName, LastName: lastName, Grade: grade, Code: code}
	err = postgresmanager.Save(&user)

	return err
}

func GetUser(username string) (User, error) {
	var user User
	var machines []*machines.Machine

	err := postgresmanager.Query(&User{Username: username}, &user)

	if err != nil {
		return User{}, err
	}

	user.Password = ""

	err = postgresmanager.ReadAssociation(&user, "Machines", &machines)
	if err != nil {
		return user, err
	}

	user.Machines = machines

	return user, nil
}

func AuthenticateUser(code, machineID string) (string, error) {
	var user *User

	if err := postgresmanager.Query(&User{Code: code}, &user); err != nil {
		return "", err
	}

	var machines []*machines.Machine
	err := postgresmanager.ReadAssociation(&user, "Machines", &machines)

	if err != nil {
		return "", err
	}

	for _, machine := range machines {
		if machine.ID == machineID {
			actions, err := machine.SignIn()
			if err != nil {
				return "", err
			}
			log.Log(fmt.Sprintf("%s %s (Username: %s) signed in to machine %s", user.FirstName, user.LastName, user.Username, machine.ID))
			return fmt.Sprintf("{\"authorized\": true, \"name\": \"%s %s\", actions: %v}", user.FirstName, user.LastName, actions), nil
		}
	}

	log.Log(fmt.Sprintf("%s failed to sign in to machine %s", user.Username, machineID))
	return "{\"authorized\": false}", nil
}

func DeleteUser(id string) error {
	if postgresmanager.ClearAssociations(&User{ID: id}, "Machines") != nil {
		return postgresmanager.ClearAssociations(&User{ID: id}, "Machines")
	}
	
	return postgresmanager.Delete(User{ID: id})
}
