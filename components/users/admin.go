package users

import (
	"errors"
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

func CreateAdmin(username, password, firstName, lastName string) (string, error) {
	code := internal.GenerateCode()

	var checkErr error
	var err error

	var count int

	var a *Admin
	checkErr = postgresmanager.Query(&Admin{Code: code}, &a)
	checkErrStr := ""
	if checkErr != nil {
		checkErrStr = checkErr.Error()
	}

	for checkErrStr != "record not found" {
		if count < 1000 {
			code = internal.GenerateCode()
			checkErr = postgresmanager.Query(&Admin{Code: code}, &a)
			if checkErr != nil {
				checkErrStr = checkErr.Error()
			} else {
				checkErrStr = ""
			}
			count++
		} else {
			err = errors.New("could not generate new code for admin")
			break
		}
	}

	if err != nil {
		return "", err
	}

	admin := Admin{ID: internal.GenerateUUID(), Username: username, Password: password, FirstName: firstName, LastName: lastName, Code: code}
	if postgresmanager.Save(&admin) != nil {
		return "", postgresmanager.Save(&admin)
	}
	return code, nil
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
	return users
}

func DeleteAdmin(id string) error {
	return postgresmanager.Delete(Admin{ID: id})
}
