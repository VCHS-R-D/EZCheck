package server

import (
	"main/components/machines"
	"main/components/postgresmanager"
	"main/components/users"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func CreateAdmin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	firstName := c.FormValue("first")
	lastName := c.FormValue("last")
	code := c.FormValue("code")

	if err := users.CreateAdmin(username, password, firstName, lastName, code); err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, "success")
}

func GetAdmin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	admin, err := users.GetAdmin(username, password)
	if err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, admin)
}

func CertifyUser(c echo.Context) error {
	userID := c.FormValue("userID")
	machineID := c.FormValue("machineID")

	if err := users.CertifyUser(userID, machineID); err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, "success")
}

func UncertifyUser(c echo.Context) error {
	userID := c.FormValue("userID")
	machineID := c.FormValue("machineID")

	if err := users.UncertifyUser(userID, machineID); err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, "success")
}

func SearchUsers(c echo.Context) error {
	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		return c.JSON(400, err)
	}

	if err := users.SearchUsers(m); err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, "success")
}

func DeleteAdmin(c echo.Context) error {
	if err := users.DeleteAdmin(c.FormValue("id")); err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, "success")
}

func AdminAuth(username, password string, c echo.Context) (bool, error) {
	var admin users.Admin
	if err := postgresmanager.Query(users.Admin{Username: username}, &admin); err != nil {
		return false, nil
	} else {
		if CheckPasswordHash(password, admin.Password) {
			return true, nil
		}
	}

	return false, nil
}

func CreateUser(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	firstName := c.FormValue("first")
	lastName := c.FormValue("last")
	code := c.FormValue("code")
	grade := c.FormValue("grade")

	if err := users.CreateUser(username, password, firstName, lastName, grade, code); err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, "success")
}

func GetUser(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := users.GetUser(username, password)
	if err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, user)
}

func DeleteUser(c echo.Context) error {
	if err := users.DeleteUser(c.FormValue("id")); err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, "success")
}

func UserAuth(username, password string, c echo.Context) (bool, error) {
	var user users.User
	if err := postgresmanager.Query(users.User{Username: username}, &user); err != nil {
		return false, nil
	} else {
		if CheckPasswordHash(password, user.Password) {
			return true, nil
		}
	}

	return false, nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateMachine(c echo.Context) error {
	name := c.FormValue("name")

	if err := machines.CreateMachine(name); err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, "success")
}

func GetMachines(c echo.Context) error {
	machines := machines.ReadMachines()
	
	return c.JSON(200, machines)
}

func SignOut(c echo.Context) error {
	if err := machines.SignOut(c.FormValue("id")); err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, "success")
}

func DeleteMachine(c echo.Context) error {
	if err := machines.DeleteMachine(c.FormValue("id")); err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, "success")
}

func Authenticate(c echo.Context) error {
	code := c.FormValue("code")
	machineID := c.FormValue("machineID")

	output, err := users.AuthenticateUser(code, machineID)

	if err != nil {
		output, err = users.AuthenticateAdmin(code, machineID)
		if err != nil {
			return c.String(400, "could not authenticate this person")
		}
		return c.JSON(200, output)
	}

	return c.JSON(200, output)
}
	
