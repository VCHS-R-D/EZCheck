package server

import (
	"github.com/labstack/echo/v4"
)

func Init() {
	e := echo.New()
	e.GET("/", testRouter)
	e.Logger.Fatal(e.Start(":1323"))
}
