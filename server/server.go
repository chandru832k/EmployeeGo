// package server

// import (
// 	"employee/controller"
// 	"employee/db"
// 	"employee/routes"
// 	"employee/service"

// 	"github.com/gin-gonic/gin"
// )

// var (
// 	EmployeeController controller.EmployeeController
// 	ListController     controller.ListController
// )

// var (
// 	EmployeeService service.EmployeeService
// 	ListService     service.ListService
// )

// var (
// 	EmployeeDao *db.EmployeeDao
// 	ListDao     *db.ListDao
// )

// func buildDB() {
// 	EmployeeDao = db.NewEmployeeDao()
// 	ListDao = db.NewListDao()
// }

// func buildService() {
// 	EmployeeService = service.NewEmployeeService(*EmployeeDao)
// 	ListService = service.NewListService(*ListDao)
// }

// func Start() {
// 	router := gin.Default()
// 	mainRouter := router.Group("/chan/")
// 	routes.CreateRoutes(mainRouter)
// 	router.Run()
// }

// func buildHandler() {
// 	buildDB()
// 	buildService()
// 	EmployeeController = controller.NewEmployeeController(EmployeeService)
// 	ListController = controller.NewListController(ListService)
// }
