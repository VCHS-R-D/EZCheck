package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func testRouter(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
