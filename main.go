package main

import (
	"employee/config"
	"employee/db"
	"employee/server"
	"fmt"
)

func main() {
	config.InitializeEnv()
	err := db.ConnectToDb()
	if err != nil {
		fmt.Println(err)
	}
	result := db.NewEmployeeDao()
	server.Start()
	fmt.Println("result", result)
}
