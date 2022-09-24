package server

import (
	"main/components/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

func testRouter(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func registerAdminRouter(c echo.Context) error {
	user := new(users.Admin)
	err := c.Bind(&user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}
