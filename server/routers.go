package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init(port string) {
	e := echo.New()

	e.HideBanner = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:          middleware.DefaultSkipper,
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))

	// FOR MACHINES
	e.POST("/auth", Authenticate)
	e.POST("/signout", SignOut)

	// AUDIT LOG
	e.GET("/log", ReadLog)

	// REGISTER NEW ADMINS AND USERS
	e.POST("/admin/create", CreateAdmin)
	e.POST("/user/create", CreateUser)

	// GET LIST OF ALL MACHINES
	e.POST("/machines", GetMachines)

	gAdmin := e.Group("/admin", middleware.BasicAuth(AdminAuth))
	gAdmin.POST("/get", GetAdmin)
	gAdmin.POST("/certify", CertifyUser)
	gAdmin.POST("/uncertify", UncertifyUser)
	gAdmin.POST("/search/users", SearchUsers)
	gAdmin.POST("/search/admins", SearchAdmins)
	gAdmin.DELETE("/delete", DeleteAdmin)
	gAdmin.POST("/machines/create", CreateMachine)
	gAdmin.DELETE("/machines/delete", DeleteMachine)
	gAdmin.POST("/machines/actions/add", AddAction)
	gAdmin.POST("/machines/actions/delete", DeleteAction)

	gUser := e.Group("/user", middleware.BasicAuth(UserAuth))
	gUser.POST("/get", GetUser)
	gUser.DELETE("/delete", DeleteUser)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
