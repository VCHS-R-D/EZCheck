package users

import (
	"main/components/internal"
	"main/components/machines"
	"main/components/postgresmanager"
)

type User struct {
	ID        string              `json:"-" gorm:"primaryKey"`
	Username  string              `json:"username" gorm:"uniqueIndex"`
	Password  string              `json:"-"`
	FirstName string              `json:"first" gorm:"index"`
	LastName  string              `json:"last" gorm:"index"`
	Code	  string			  `json:"code" gorm:"uniqueIndex"`
	Machines  []*machines.Machine `gorm:"many2many:users_machines"`
	CreatedAt    time.Time  `json:"-" gorm:"index"`
	UpdatedAt    time.Time  `json:"-" gorm:"index"`
}

func CreateUser(username, password, firstName, lastName string) (string, error) {
	code := internal.GenerateCode()

	var checkErr error
	var err error

	var count int

	var u *User
	checkErr = postgresmanager.Query(&User{Code: code}, &u)
	checkErrStr := ""
	if checkErr != nil {
		checkErrStr = checkErr.Error()
	}

	for checkErrStr != "record not found" {
		if count < 1000 {
			code = codes.GenerateCode()
			checkErr = postgresmanager.Query(&User{Code: code}, &u)
			if checkErr != nil {
				checkErrStr = checkErr.Error()
			} else {
				checkErrStr = ""
			}
			count++
		} else {
			err = errors.New("could not generate new code for user")
			break
		}
	}

	if err != nil {
		return err
	}

	user := User{ID: internal.GenerateUUID(), Username: username, Password: password, FirstName: firstName, LastName: lastName, Code: code}
	if postgresmanager.Save(&user) != nil {
		return "", postgresmanager.Save(&user)
	}
	return code, nil
}

func ReadUser(id string) User {
	var user User
	postgresmanager.Query(User{ID: id}, &user)

	return user
}