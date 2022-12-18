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

func CertifyUser(adminID, userID, machineID string) error {
	var user *User
	var admin *Admin
	var machine *machines.Machine
	err := postgresmanager.Query(User{ID: userID}, &user)
	if err != nil {
		return err
	}

	err = postgresmanager.Query(Admin{ID: adminID}, &admin)
	if err != nil {
		return err
	}

	err = postgresmanager.Query(machines.Machine{ID: machineID}, &machine)
	if err != nil {
		return err
	}

	err = postgresmanager.CreateAssociation(&user, "Machines", machine)
	log.Log(fmt.Sprintf("%s authorized user %s to use machine %s", admin.Username, user.Username, machine.Name))

	return err
}

func UncertifyUser(adminID, userID, machineID string) error {
	var user *User
	var admin *Admin
	var machine *machines.Machine
	err := postgresmanager.Query(User{ID: userID}, &user)
	if err != nil {
		return err
	}

	err = postgresmanager.Query(Admin{ID: adminID}, &admin)
	if err != nil {
		return err
	}

	err = postgresmanager.Query(machines.Machine{ID: machineID}, &machine)
	if err != nil {
		return err
	}

	err = postgresmanager.DeleteAssociation(&user, "Machines", machine)
	log.Log(fmt.Sprintf("%s deauthorized user %s to use machine %s", admin.Username, user.Username, machine.Name))

	return err
}

func SearchUsers(query map[string]interface{}) []User {
	var users []User

	if len(query) == 0 {
		postgresmanager.QueryAll(&users)
		for i, u := range users {
			u.Password = ""
			users[i] = u
		}
		return users
	}

	postgresmanager.GroupQuery(query, &users)
	for i, u := range users {
		u.Password = ""
		users[i] = u
	}
	
	return users
}

func SearchAdmins(query map[string]interface{}) []Admin {
	var admins []Admin

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
	var admin *Admin

	if err := postgresmanager.Query(&Admin{Code: code}, &admin); err != nil {
		return "", err
	}

	var machine *machines.Machine
	if err := postgresmanager.Query(&machines.Machine{ID: machineID}, &machine); err != nil {
		return "", err
	}

	actions, err := machine.SignIn()
	if err != nil {
		log.Log(fmt.Sprintf("%s failed to sign in to machine %s", admin.Username, machineID))
		return "", err
	}

	log.Log(fmt.Sprintf("%s %s (Username: %s) signed in to machine %s", admin.FirstName, admin.LastName, admin.Username, machine.Name))
	return fmt.Sprintf("{\"authorized\": true, \"name\": \"%s %s\", actions: %v}", admin.FirstName, admin.LastName, actions), nil
}

func DeleteAdmin(id string) error {
	return postgresmanager.Delete(Admin{ID: id})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
