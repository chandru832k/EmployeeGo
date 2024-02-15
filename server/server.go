package server

import (
	"employee/controller"
	"employee/db"
	"employee/routes"
	"employee/service"

	"github.com/gin-gonic/gin"
)

var (
	EmployeeController controller.EmployeeController
)

var (
	EmployeeService *service.EmployeeService
)

var (
	EmployeeDao *db.EmployeeDao
)

func buildDB() {
	EmployeeDao = db.NewEmployeeDao()
}

func buildService() {
	EmployeeService = service.NewEmployeeService(*EmployeeDao)
}

func Start() {
	router := gin.Default()
	mainRouter := router.Group("/chan/")
	routes.CreateRoutes(mainRouter)
	router.Run()
}

func buildHandler() {
	buildDB()
	buildService()
	EmployeeController = controller.NewEmployeeController(*EmployeeService)
}
