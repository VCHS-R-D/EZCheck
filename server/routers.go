package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	e := echo.New()

	gAdmin := e.Group("/admin", middleware.BasicAuth(AdminAuth))
	gAdmin.POST("/create", CreateAdmin)

	e.Logger.Fatal(e.Start(":3000"))
}
