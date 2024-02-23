package routes

import (
	"github.com/gin-gonic/gin"
)

func EmployeeRoutes(router *gin.RouterGroup) {
	employeeGroup := router.Group("/employee")
	{
		employeeGroup.POST("/", EmployeeController.AddEmployee)
		employeeGroup.PUT("/:employee_id", EmployeeController.UpdateEmployee)
		employeeGroup.DELETE("/:employee_id", EmployeeController.DeleteEmployee)
		employeeGroup.GET("/", EmployeeController.GetAllEmployees)
		employeeGroup.GET("/:employee_id", EmployeeController.GetEmployeeById)
	}
}
