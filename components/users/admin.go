package users

import (
	"main/components/internal"
	"main/components/machines"
	"main/components/postgresmanager"
)

type Admin struct {
	ID        string `json:"-" gorm:"primaryKey"`
	Username  string `json:"username" gorm:"uniqueIndex"`
	Password  string `json:"-"`
	FirstName string `json:"first" gorm:"index"`
	LastName  string `json:"last" gorm:"index"`
	Code	  string `json:"code" gorm:"uniqueIndex"`
	CreatedAt    time.Time  `json:"-" gorm:"index"`
	UpdatedAt    time.Time  `json:"-" gorm:"index"`
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
			code = codes.GenerateCode()
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
		return err
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

	return admin
}

func CertifyUser(code, machineID string) error {
	var user *User
	var machine *machines.Machine
	err := postgresmanager.Query(User{Code: code}, &user)
	if err != nil {
		return err
	}

	err = postgresmanager.Query(machines.Machine{ID: machineID}, &machine)

	if err != nil {
		return err
	}

	err = postgresmanager.CreateAssociation(&user, "Machines", machine)

	return err
}
