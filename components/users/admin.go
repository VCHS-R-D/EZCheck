package users

import (
	"main/components/postgresmanager"
	"main/components/internal"
)

var registrationTable = make(map[string]string)

type Admin struct {
	ID              string
	FingerprintHash string
	Username        string
	Password        string
	FirstName       string
	LastName        string
}

func GenerateTemporaryAdmin(fingerprintHash string) (string) {
	code := internal.GenerateCode()
	registrationTable[code] = fingerprintHash
	return code
}

func CreateAdmin(username, password, firstName, lastName, code string) (error) {
	fingerprintHash, ok := registrationTable[code]
	if !ok {
		return nil
	}

	defer delete(registrationTable, code)
	
	admin := Admin{ID: internal.GenerateUUID(), Username: username, Password: password, FirstName: firstName, LastName: lastName, FingerprintHash: fingerprintHash}
	return postgresmanager.Save(&admin)
}
