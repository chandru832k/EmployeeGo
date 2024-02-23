package main

import (
	"employee/config"
	"employee/db"
	"employee/routes"
	"fmt"
)


func main() {
	config.InitializeEnv()
	err := db.ConnectToDb()
	if err != nil {
		fmt.Println(err)
	}
	routes.StartServer()
}
