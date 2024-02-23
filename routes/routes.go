package routes

import (
	"employee/controller"
	"employee/db"
	"employee/service"

	"github.com/gin-gonic/gin"
)

var (
	EmployeeController controller.EmployeeController
	ProjectController  controller.ProjectController
)

var (
	EmployeeService service.EmployeeService
	ProjectService  service.ProjectService
)

var (
	EmployeeDao db.EmployeeDao
	projectDao  db.ProjectDao
)

func buildDB() {
	EmployeeDao = db.NewEmployeeDao()
	// ListDao = db.NewListDao()
	projectDao = db.NewProjectDao()
}

func buildService() {
	EmployeeService = service.NewEmployeeService(EmployeeDao)
	ProjectService = service.NewProjectService(projectDao)
}

func buildHandlers() {
	buildDB()
	buildService()
	EmployeeController = controller.NewEmployeeController(EmployeeService)
	ProjectController = controller.NewProjectController(ProjectService)
}

func CreateRoutes(router *gin.RouterGroup) {
	EmployeeRoutes(router)
	ListRoutes(router)
}

func StartServer() {
	router := gin.Default()
	mainRouter := router.Group("/chan/")
	buildHandlers()
	CreateRoutes(mainRouter)
	router.Run()
}
