package users

import (
	"fmt"
	"main/components/internal"
	"main/components/log"
	"main/components/machines"
	"main/components/postgresmanager"
	"main/components/types"

	"golang.org/x/crypto/bcrypt"
)

func CreateAdmin(username, password, firstName, lastName, code string) error {
	var a *types.Admin
	err := postgresmanager.Query(&types.Admin{Code: code}, &a)

	if err != nil {
		if err.Error() != "record not found" {
			return err
		}
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	admin := types.Admin{ID: internal.GenerateUUID(), Username: username, Password: hashedPassword, FirstName: firstName, LastName: lastName, Code: code}

	err = postgresmanager.Save(&admin)

	return err
}

func GetAdmin(username string) (*types.Admin, error) {
	var admin *types.Admin
	err := postgresmanager.Query(&types.Admin{Username: username}, &admin)
	if err != nil {
		return nil, err
	}

	admin.Password = ""
	return admin, nil
}

func CertifyUser(adminID, userID, machineID string) error {
	var user *types.User
	var admin *types.Admin
	var machine *types.Machine
	err := postgresmanager.Query(types.User{ID: userID}, &user)
	if err != nil {
		return err
	}

	err = postgresmanager.Query(types.Admin{ID: adminID}, &admin)
	if err != nil {
		return err
	}

	err = postgresmanager.Query(types.Machine{ID: machineID}, &machine)
	if err != nil {
		return err
	}

	err = postgresmanager.CreateAssociation(&user, "Machines", machine)
	log.Log(fmt.Sprintf("admin %s authorized user %s to use machine %s", admin.Username, user.Username, machine.ID))

	return err
}

func UncertifyUser(adminID, userID, machineID string) error {
	var user *types.User
	var admin *types.Admin
	var machine *types.Machine
	err := postgresmanager.Query(types.User{ID: userID}, &user)
	if err != nil {
		return err
	}

	err = postgresmanager.Query(types.Admin{ID: adminID}, &admin)
	if err != nil {
		return err
	}

	err = postgresmanager.Query(types.Machine{ID: machineID}, &machine)
	if err != nil {
		return err
	}

	err = postgresmanager.DeleteAssociation(&user, "Machines", machine)
	log.Log(fmt.Sprintf("admin %s deauthorized user %s from machine %s", admin.Username, user.Username, machine.ID))

	return err
}

func SearchUsers(query map[string]interface{}) []types.User {
	var users []types.User

	if len(query) == 0 {
		postgresmanager.QueryAll(&users)
		for i, u := range users {
			var machines []*types.Machine
			u.Password = ""
			postgresmanager.ReadAssociation(&u, "Machines", &machines)
			u.Machines = machines
			users[i] = u
		}
		return users
	}

	postgresmanager.GroupQuery(query, &users)
	for i, u := range users {
		var machines []*types.Machine
		u.Password = ""
		postgresmanager.ReadAssociation(&u, "Machines", &machines)
		u.Machines = machines
		users[i] = u
	}

	return users
}

func SearchAdmins(query map[string]interface{}) []*types.Admin {
	var admins []*types.Admin

	if len(query) == 0 {
		postgresmanager.QueryAll(&admins)
		for i, a := range admins {
			a.Password = ""
			admins[i] = a
		}
		return admins
	}

	postgresmanager.GroupQuery(query, &admins)
	for i, a := range admins {
		a.Password = ""
		admins[i] = a
	}

	return admins
}

func AuthenticateAdmin(code, machineID string) (string, error) {
	var admin *types.Admin

	if err := postgresmanager.Query(&types.Admin{Code: code}, &admin); err != nil {
		return "", err
	}

	actions, err := machines.SignIn(machineID)
	if err != nil {
		log.Log(fmt.Sprintf("%s failed to sign in to machine %s", admin.Username, machineID))
		return "", err
	}

	log.Log(fmt.Sprintf("%s %s (Username: %s) signed in to machine %s", admin.FirstName, admin.LastName, admin.Username, machineID))
	return fmt.Sprintf("{\"authorized\": true, \"name\": \"%s %s\", actions: %v}", admin.FirstName, admin.LastName, actions), nil
}

func DeleteAdmin(id string) error {
	return postgresmanager.Delete(types.Admin{ID: id})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
