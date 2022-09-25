package users

import (
	"main/components/internal"
	"main/components/machines"
	"main/components/postgresmanager"
)

var adminRegistrationTable = make(map[string]string)

type Admin struct {
	ID              string
	FingerprintHash string `json:"-"`
	Username        string `json:"username"`
	Password        string `json:"-"`
	FirstName       string `json:"first"`
	LastName        string `json:"last"`
}

func GenerateTemporaryAdmin(fingerprintHash string) string {
	code := internal.GenerateCode()
	adminRegistrationTable[code] = fingerprintHash
	return code
}

func CreateAdmin(username, password, firstName, lastName, code string) error {
	fingerprintHash, ok := adminRegistrationTable[code]
	if !ok {
		return nil
	}

	defer delete(adminRegistrationTable, code)

	admin := Admin{ID: internal.GenerateUUID(), Username: username, Password: password, FirstName: firstName, LastName: lastName, FingerprintHash: fingerprintHash}
	return postgresmanager.Save(&admin)
}

func ReadAdmin(id string) Admin {
	var admin Admin
	postgresmanager.Query(Admin{ID: id}, &admin)

	return admin
}

func CertifyUser(username, machineID string) error {
	var user User
	var machine machines.Machine
	err := postgresmanager.Query(User{Username: username}, &user)
	if err != nil {
		return err
	}

	err = postgresmanager.Query(machines.Machine{ID: machineID}, &machine)

	if err != nil {
		return err
	}

	err = postgresmanager.CreateAssociation(&user, "Machines", &machine)

	return err
}
