package controller

import (
	"employee/resources"
	"employee/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	EmployeeService service.EmployeeService
}

func NewEmployeeController(EmployeeService service.EmployeeService) EmployeeController {
	return EmployeeController{
		EmployeeService: EmployeeService,
	}
}

func AddEmployee(c *gin.Context) {
	ctx := c.Request.Context()
	functionDesc := "Add Employee Controller"
	fmt.Println(functionDesc)
	fmt.Println(ctx)
	// result := c
	var input resources.AddEmployeeInput
	err := c.BindJSON(&input)
	if err != nil {
		return
	}
	employeeId := service.AddEmployee(ctx, input)
	if employeeId <= 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error While Creating Employee",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"employee_id": employeeId,
		"message":     "Employee added Successfully",
	})
}

func GetAllEmployees(c *gin.Context) {
	functionDesc := "Get All Employees Controller"
	fmt.Println(functionDesc)

	result := service.GetAllEmployees(c)

	if result.Code != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error Fetching Employee Details",
		})
		return
	}
	c.JSON(http.StatusOK, result.ServiceResultData.Data)
}

func GetEmployeeById(c *gin.Context) {
	functionDesc := "Get Employee By Id Controller"
	fmt.Println(functionDesc)

	employeeId, err := strconv.Atoi(c.Param("employee_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Employee Id Not Found",
		})
		return
	}

	result :=  service.GetEmployeeById(c, employeeId)

	if result.Code != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error Fetching Employee Details",
		})
		return
	}
	c.JSON(http.StatusOK,  result.ServiceResultData.Data)
}

func UpdateEmployee(c *gin.Context) {
	ctx := c.Request.Context()
	functionDesc := "Update Employee Controller"
	fmt.Println(functionDesc)
	fmt.Println(ctx)

	var input resources.AddEmployeeInput
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error reading payload",
		})
		return
	}

	employeeId, err := strconv.Atoi(c.Param("employee_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Employee Id Not Found",
		})
		return
	}

	result := service.UpdateEmployee(c, employeeId, input)

	if result != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error While Updating Employee",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Employee Details Updated Successfully",
	})
	// fmt.Println(result)
}

func DeleteEmployee(c *gin.Context) {
	functionDesc := "Delete Employee Controller"
	fmt.Println(functionDesc)

	employeeId, err := strconv.Atoi(c.Param("employee_id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Employee Id Not Found",
		})
		return
	}

	result := service.DeleteEmployee(c, employeeId)

	if result != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error While Deleting Employee",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Employee Details Deleted Successfully",
	})
}
