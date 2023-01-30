package server

import (
	"main/components/log"
	"main/components/machines"
	"main/components/postgresmanager"
	"main/components/users"
	"os"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func CreateAdmin(c echo.Context) error {
	if c.FormValue("adminPass") != os.Getenv("ADMIN_PASS") {
		return c.JSON(400, "Invalid admin password")
	}

	if err := users.CreateAdmin(c.FormValue("username"), c.FormValue("password"), c.FormValue("first"), c.FormValue("last"), c.FormValue("code")); err != nil {
		return c.JSON(400, err)
	}

	return c.String(200, "success")
}

func CertifyUser(c echo.Context) error {
	if err := users.CertifyUser(c.FormValue("adminID"), c.FormValue("userID"), c.FormValue("machineID")); err != nil {
		return c.JSON(400, err)
	}

	return c.String(200, "success")
}

func UncertifyUser(c echo.Context) error {
	if err := users.UncertifyUser(c.FormValue("adminID"), c.FormValue("userID"), c.FormValue("machineID")); err != nil {
		return c.JSON(400, err)
	}

	return c.String(200, "success")
}

func SearchUsers(c echo.Context) error {
	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, users.SearchUsers(m))
}

func SearchAdmins(c echo.Context) error {
	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, users.SearchAdmins(m))
}

func GetAdmin(c echo.Context) error {
	admin, err := users.GetAdmin(c.FormValue("username"))
	if err != nil {
		return c.JSON(400, err)
	}
	return c.JSON(200, admin)
}

func DeleteAdmin(c echo.Context) error {
	if err := users.DeleteAdmin(c.FormValue("id")); err != nil {
		return c.JSON(400, err)
	}

	return c.String(200, "success")
}

func AdminAuth(username, password string, c echo.Context) (bool, error) {
	var admin users.Admin
	if err := postgresmanager.Query(users.Admin{Username: username}, &admin); err != nil {
		return false, c.JSON(400, err)
	} else {
		if CheckPasswordHash(password, admin.Password) {
			return true, nil
		}
	}

	return false, c.String(400, "Invalid username or password")
}

func CreateUser(c echo.Context) error {
	if err := users.CreateUser(c.FormValue("username"), c.FormValue("password"), c.FormValue("first"), c.FormValue("last"), c.FormValue("grade"), c.FormValue("code")); err != nil {
		return c.JSON(400, err)
	}

	return c.String(200, "success")
}

func GetUser(c echo.Context) error {
	user, err := users.GetUser(c.FormValue("username"))
	if err != nil {
		return c.JSON(400, err)
	}
	return c.JSON(200, user)
}

func DeleteUser(c echo.Context) error {
	if err := users.DeleteUser(c.FormValue("id")); err != nil {
		return c.JSON(400, err)
	}

	return c.String(200, "success")
}

func UserAuth(username, password string, c echo.Context) (bool, error) {
	var user users.User
	if err := postgresmanager.Query(users.User{Username: username}, &user); err != nil {
		return false, c.JSON(400, err)
	} else {
		if CheckPasswordHash(password, user.Password) {
			return true, nil
		}
	}

	return false, c.String(400, "Invalid username or password")
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateMachine(c echo.Context) error {
	if err := machines.CreateMachine(c.FormValue("machineID")); err != nil {
		return c.JSON(400, err)
	}

	return c.String(200, "success")
}

func GetMachines(c echo.Context) error {
	return c.JSON(200, machines.GetMachines())
}

func SignOut(c echo.Context) error {
	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		return c.JSON(400, err)
	}

	var nameString string
	var machineIDString string

	if name, ok := m["name"]; ok {
		nameString = name.(string)
	} else {
		return c.String(400, "name key not found")
	}

	if machineID, ok := m["machineID"]; ok {
		machineIDString = machineID.(string)
	} else {
		return c.String(400, "machineID key not found")
	}

	if err := machines.SignOut(nameString, machineIDString); err != nil {
		return c.JSON(400, err)
	}

	return c.String(200, "success")
}

func DeleteMachine(c echo.Context) error {
	if err := machines.DeleteMachine(c.FormValue("machineID")); err != nil {
		return c.JSON(400, err)
	}

	return c.String(200, "success")
}

func Authenticate(c echo.Context) error {
	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		return c.JSON(400, err)
	}

	var codeString string
	var machineIDString string

	if code, ok := m["code"]; ok {
		codeString = code.(string)
	} else {
		return c.String(400, "code key not found")
	}

	if machineID, ok := m["machineID"]; ok {
		machineIDString = machineID.(string)
	} else {
		return c.String(400, "machineID key not found")
	}

	if codeString == "" || machineIDString == "" {
		return c.String(400, "code or machineID is empty")
	}

	output, err := users.AuthenticateUser(codeString, machineIDString)

	if err != nil {
		if err.Error() == "record not found" {
			output, err = users.AuthenticateAdmin(codeString, machineIDString)
			if err != nil {
				if err.Error() == "record not found" {
					return c.String(400, "could not authenticate this person")
				} else {
					return c.JSON(400, err)
				}
			}
			return c.String(200, output)
		} else {
			return c.JSON(400, err)
		}
	}

	return c.String(200, output)
}

func ReadLog(c echo.Context) error {
	log, err := log.Read()
	if err != nil {
		return c.JSON(400, err)
	}
	return c.String(200, log)
}

func AddAction(c echo.Context) error {
	if err := machines.AddAction(c.FormValue("machineID"), c.FormValue("actionID")); err != nil {
		return c.JSON(400, err)
	}

	return c.String(200, "success")
}

func DeleteAction(c echo.Context) error {
	if err := machines.DeleteAction(c.FormValue("machineID"), c.FormValue("actionID")); err != nil {
		return c.JSON(400, err)
	}

	return c.String(200, "success")
}
