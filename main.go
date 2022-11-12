package main

import (
	"flag"
	"main/components/postgresmanager"
	"main/server"
)

func main() {

	dhostPtr := flag.String("dbhost", "localhost", "host")
	dbnamePtr := flag.String("dbname", "postgres", "name")
	dbportPtr := flag.String("dbport", "5432", "port")
	dbuserPtr := flag.String("dbuser", "postgres", "user")
	dbpassPtr := flag.String("dbpassword", "password", "password")
	serverPtr := flag.String("port", "8080", "port")
	flag.Parse()

	postgresmanager.Open(*dhostPtr, *dbnamePtr, *dbportPtr, *dbuserPtr, *dbpassPtr)

	server.Init(*serverPtr)
}
