package server

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init(port string) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/auth", Authenticate)

	gAdmin := e.Group("/admin", middleware.BasicAuth(AdminAuth))
	gAdmin.POST("/create", CreateAdmin)
	gAdmin.POST("/get", GetAdmin)
	gAdmin.POST("/certify", CertifyUser)
	gAdmin.POST("/uncertify", UncertifyUser)
	gAdmin.POST("/search", SearchUsers)
	gAdmin.DELETE("/delete", DeleteAdmin)
	gAdmin.POST("/machines/create", CreateMachine)
	gAdmin.POST("/machines/get", GetMachines)
	gAdmin.POST("/machines/signout", SignOut)
	gAdmin.DELETE("/machines/delete", DeleteMachine)

	gUser := e.Group("/user", middleware.BasicAuth(UserAuth))
	gUser.POST("/create", CreateUser)
	gUser.POST("/get", GetUser)
	gUser.DELETE("/delete", DeleteUser)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
