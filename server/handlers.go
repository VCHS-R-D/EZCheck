package server

import (
	"main/components/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

func registerAdminRouter(c echo.Context) error {
	user := new(users.Admin)
	err := c.Bind(&user)
	if err != nil {
		return err
	}

	users.CreateAdmin(user.Username, user.Password, user.FirstName, user.LastName)

	return c.JSON(http.StatusCreated, user)
}

func registerStudentRouter(c echo.Context) error {
	user := new(users.User)

	if err := c.Bind(&user); err != nil {
		return err
	}

	users.CreateUser(user.Username, user.Password, user.FirstName, user.LastName)

	return c.JSON(http.StatusCreated, user)
}
