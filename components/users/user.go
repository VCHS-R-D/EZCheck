package users

import (
	"fmt"
	"main/components/internal"
	"main/components/log"
	"main/components/machines"
	"main/components/postgresmanager"
	"main/components/types"
)

func CreateUser(username, password, firstName, lastName, grade, code string) error {
	var u *types.User
	err := postgresmanager.Query(&types.User{Code: code}, &u)

	if err != nil {
		if err.Error() != "record not found" {
			return err
		}
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	user := types.User{ID: internal.GenerateUUID(), Username: username, Password: hashedPassword, FirstName: firstName, LastName: lastName, Grade: grade, Code: code}
	err = postgresmanager.Save(&user)

	return err
}

func GetUser(username string) (types.User, error) {
	var user types.User
	var machines []*types.Machine

	err := postgresmanager.Query(&types.User{Username: username}, &user)

	if err != nil {
		return types.User{}, err
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
	var user *types.User

	if err := postgresmanager.Query(&types.User{Code: code}, &user); err != nil {
		return "", err
	}

	var myMachines []*types.Machine
	err := postgresmanager.ReadAssociation(&user, "Machines", &myMachines)

	if err != nil {
		return "", err
	}

	for _, machine := range myMachines {
		if machine.ID == machineID {
			actions, err := machines.SignIn(machineID)
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
	err := postgresmanager.ClearAssociations(&types.User{ID: id}, "Machines")
	if err != nil {
		return postgresmanager.ClearAssociations(&types.User{ID: id}, "Machines")
	}

	return postgresmanager.Delete(types.User{ID: id})
}
