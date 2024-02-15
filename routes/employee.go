package routes

import (
	"employee/controller"

	"github.com/gin-gonic/gin"
)

func EmployeeRoutes(router *gin.RouterGroup) {
	employeeGroup := router.Group("/employee")
	{
		employeeGroup.POST("/", controller.AddEmployee)
		employeeGroup.PUT("/:employee_id", controller.UpdateEmployee)
		employeeGroup.DELETE("/:employee_id", controller.DeleteEmployee)
		employeeGroup.GET("/", controller.GetAllEmployees)
		employeeGroup.GET("/:employee_id", controller.GetEmployeeById)
	}
}
