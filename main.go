package main

import (
	"flag"
	"main/components/machines"
	"main/components/postgresmanager"
	"main/components/users"
	"main/server"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	dhostPtr := flag.String("dbhost", "localhost", "host")
	dbnamePtr := flag.String("dbname", "postgres", "name")
	dbportPtr := flag.String("dbport", "5432", "port")
	dbuserPtr := flag.String("dbuser", "postgres", "user")
	dbpassPtr := flag.String("dbpassword", "password", "password")
	serverPtr := flag.String("port", "8080", "port")
	flag.Parse()

	err := postgresmanager.Open(*dhostPtr, *dbnamePtr, *dbportPtr, *dbuserPtr, *dbpassPtr)
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

	err = godotenv.Load("variables.env")
	if err != nil {
		panic(err)
	}

	if os.Getenv("ADMIN_PASS") == "" {
		panic("ADMIN_PASS environment variable is not set")
	}

	server.Init(*serverPtr)
}
