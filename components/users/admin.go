package users

import (
	"fmt"
	"main/components/internal"
	"main/components/log"
	"main/components/machines"
	"main/components/postgresmanager"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	ID        string    `json:"id" gorm:"primaryKey"`
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

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	admin := Admin{ID: internal.GenerateUUID(), Username: username, Password: hashedPassword, FirstName: firstName, LastName: lastName, Code: code}

	err = postgresmanager.Save(&admin)

	return err
}

func GetAdmin(id string) (Admin, error) {
	var admin Admin
	
	err := postgresmanager.Query(&Admin{ID: id}, &admin)
	if err != nil {
		return Admin{}, err
	}

	admin.Password = ""
	return admin, nil
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

func AuthenticateAdmin(code, machineID string) (string, error) {
	var admin *Admin

	if err := postgresmanager.Query(&Admin{Code: code}, &admin); err != nil {
		return "{\"error\": \"admin not found\"}", err
	}

	var machine *machines.Machine
	if err := postgresmanager.Query(&machines.Machine{ID: machineID}, &machine); err != nil {
		return "{\"error\": \"machine not found\"}", nil
	}

	actions, err := machine.SignIn()
	if err != nil {
		log.Log(fmt.Sprintf("%s failed to sign in to machine %s", admin.Username, machineID))
		return "{\"error\": \"could not sign in\"}", nil
	}

	log.Log(fmt.Sprintf("%s signed in to machine %s", admin.Username, machine.Name))
	return fmt.Sprintf("{\"authorized\": true, \"name\": \"%s %s\", actions: %v}", admin.FirstName, admin.LastName, actions), nil
}

func DeleteAdmin(id string) error {
	return postgresmanager.Delete(Admin{ID: id})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
