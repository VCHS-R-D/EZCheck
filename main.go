package main

import (
	"main/components/postgresmanager"
	"main/server"
)

func main() {
	postgresmanager.Open("localhost", "EZCheck", "4000", "root", "root")
	server.Init()
}
