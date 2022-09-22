package users

import (
	"main/components/internal"
	"main/components/machines"
	"main/components/postgresmanager"
)

var userRegistrationTable = make(map[string]string)

type User struct {
	ID              string
	FingerprintHash string `json:"-"`
	Username        string
	Password        string `json:"-"`
	FirstName       string
	LastName        string
	Machines        []*machines.Machine `gorm:"many2many:users_machines"`
}

func GenerateTemporaryUser(fingerprintHash string) string {
	code := internal.GenerateCode()
	userRegistrationTable[code] = fingerprintHash
	return code
}

func CreateUser(username, password, firstName, lastName, code string) error {
	fingerprintHash, ok := userRegistrationTable[code]
	if !ok {
		return nil
	}

	defer delete(userRegistrationTable, code)

	user := User{ID: internal.GenerateUUID(), Username: username, Password: password, FirstName: firstName, LastName: lastName, FingerprintHash: fingerprintHash}
	return postgresmanager.Save(&user)
}

func ReadUser(id string) User {
	var user User
	postgresmanager.Query(User{ID: id}, &user)

	return user
}
