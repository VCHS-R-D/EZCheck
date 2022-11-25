package main

import (
	"main/components/machines"
	"main/components/postgresmanager"
	"main/components/users"
	"main/server"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("variables.env")
	if err != nil {
		panic(err)
	}

	if os.Getenv("ADMIN_PASS") == "" {
		panic("ADMIN_PASS environment variable is not set")
	}

	err = postgresmanager.Open(os.Getenv("DB_HOST"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"))
	if err != nil {
		panic(err)
	}

	err = postgresmanager.AutoCreateStruct(users.Admin{})
	if err != nil {
		panic(err)
	}

	err = postgresmanager.AutoCreateStruct(users.User{})
	if err != nil {
		panic(err)
	}

	err = postgresmanager.AutoCreateStruct(machines.Machine{})
	if err != nil {
		panic(err)
	}

	server.Init(os.Getenv("HTTP_PORT"))
}
