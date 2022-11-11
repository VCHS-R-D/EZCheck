package server

import (
	"github.com/labstack/echo/v4"
)

func Init() {
	e := echo.New()

	e.Logger.Fatal(e.Start(":3000"))
}
