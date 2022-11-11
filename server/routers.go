package server

import (
	"github.com/labstack/echo/v4"
)

func Init() {
	e := echo.New()

	e.POST("/register/admin", registerAdminRouter)
	e.POST("/register/student", registerStudentRouter)

	g := e.Group("/admin")
	g.Use(middleware(BasicAuth())

	e.Logger.Fatal(e.Start(":3000"))
}
