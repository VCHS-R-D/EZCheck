package server

import (
	"github.com/labstack/echo/v4"
)

func Init() {
	e := echo.New()

	e.POST("/admin/register", registerAdminRouter)
	e.POST("/student/register", registerStudentRouter)

	e.Logger.Fatal(e.Start(":3000"))
}
