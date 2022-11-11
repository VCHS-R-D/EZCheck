package server

import (
	"github.com/labstack/echo/v4"
)

func Init() {
	e := echo.New()

<<<<<<< HEAD
	e.POST("/register/admin", registerAdminRouter)
	e.POST("/register/student", registerStudentRouter)

	g := e.Group("/admin")
	g.Use(middleware(BasicAuth())

=======
>>>>>>> eb011a258ab3e973868b3652481219b572955586
	e.Logger.Fatal(e.Start(":3000"))
}
