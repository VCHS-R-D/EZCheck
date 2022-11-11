package server

import (
	"main/components/postgresmanager"
	"main/components/users"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
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
