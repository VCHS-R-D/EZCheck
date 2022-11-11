package users

import (
	"fmt"
	"main/components/internal"
	"main/components/log"
	"main/components/machines"
	"main/components/postgresmanager"
	"time"
)

type Admin struct {
	ID        string    `json:"-" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"uniqueIndex"`
	Password  string    `json:"password"`
	FirstName string    `json:"first" gorm:"index"`
	LastName  string    `json:"last" gorm:"index"`
	Code      string    `json:"code" gorm:"uniqueIndex"`
	CreatedAt time.Time `json:"-" gorm:"index"`
	UpdatedAt time.Time `json:"-" gorm:"index"`
}

func CreateAdmin(username, password, firstName, lastName, code string) error {

	var a *Admin
	err := postgresmanager.Query(&Admin{Code: code}, &a)

	if err.Error() != "record not found" {
		return err
	}

	admin := Admin{ID: internal.GenerateUUID(), Username: username, Password: password, FirstName: firstName, LastName: lastName, Code: code}

	err = postgresmanager.Save(&admin)

	return err
}

func ReadAdmin(id string) Admin {
	var admin Admin
	postgresmanager.Query(Admin{ID: id}, &admin)
	admin.Password = ""
	return admin
}

func CertifyUser(userID, machineID string) error {
	var user *User
	var machine *machines.Machine
	err := postgresmanager.Query(User{ID: userID}, &user)
	if err != nil {
		return err
	}

	err = postgresmanager.Query(machines.Machine{ID: machineID}, &machine)

	if err != nil {
		return err
	}

	err = postgresmanager.CreateAssociation(&user, "Machines", machine)
	log.Log(fmt.Sprintf("Authorized user %s to machine %s", user.Username, machine.Name))

	return err
}

func UncertifyUser(userID, machineID string) error {
	var user *User
	var machine *machines.Machine
	err := postgresmanager.Query(User{ID: userID}, &user)
	if err != nil {
		return err
	}

	err = postgresmanager.Query(machines.Machine{ID: machineID}, &machine)

	if err != nil {
		return err
	}

	err = postgresmanager.DeleteAssociation(&user, "Machines", machine)
	log.Log(fmt.Sprintf("Unauthorized user %s to machine %s", user.Username, machine.Name))

	return err
}

func SearchUsers(query map[string]interface{}) []User {
	var users []User
	postgresmanager.GroupQuery(query, &users)
	for _, u := range users {
		u.Password = ""
	}
	return users
}

func AuthenticateAdmin(code, machineID string) string {
	var admin *Admin

	if postgresmanager.Query(&Admin{Code: code}, &admin) != nil {
		return "{\"error\": \"admin not found\"}"
	}

	var machine *machines.Machine
	if postgresmanager.Query(&machines.Machine{ID: machineID}, &machine) != nil {
		return "{\"error\": \"machine not found\"}"
	}

	actions, err := machine.SignIn()
	if err != nil {
		log.Log(fmt.Sprintf("%s failed to sign in to machine %s", admin.Username, machineID))
		return "{\"error\": \"could not sign in\"}"
	}

	log.Log(fmt.Sprintf("%s signed in to machine %s", admin.Username, machine.Name))
	return fmt.Sprintf("{\"authorized\": true, \"name\": \"%s %s\", actions: %v}", admin.FirstName, admin.LastName, actions)
}

func DeleteAdmin(id string) error {
	return postgresmanager.Delete(Admin{ID: id})
}
