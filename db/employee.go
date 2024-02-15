package db

import (
	"context"
	"employee/models"
	"employee/resources"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EmployeeDao struct {
	db *gorm.DB
}

func NewEmployeeDao() *EmployeeDao {
	// fmt.Println("dbClient, %v", DbClient)
	return &EmployeeDao{
		db: DbClient,
	}
}

func AddEmployee(ctx context.Context, input resources.AddEmployeeInput) uint64 {
	data := models.Employee{
		EmployeeName:   input.EmployeeName,
		EmployeeRole:   input.EmployeeRole,
		EmployeeMail:   input.EmployeeMail,
		EmployeeAge:    uint(input.EmployeeAge),
		EmployeeGender: input.EmployeeGender,
	}
	queryPtr := DbClient.Debug().Table("employee").Create(&data)

	if queryPtr.Error != nil {
		fmt.Println("Error in DB Add Employee | ", queryPtr.Error.Error())
		return 0
	}

	return data.EmployeeId
}

func GetAllEmployees(ctx *gin.Context) ([]models.Employee, error) {
	result := []models.Employee{}
	queryPtr := DbClient.Debug().Table("employee").Find(&result)

	if err := queryPtr.Error; err != nil {
		fmt.Print("Error updating Employee Details | ", queryPtr.Error.Error())
		// return resources.ServiceResult{
		// 	Code:              http.StatusInternalServerError,
		// 	ServiceResultData: resources.ServiceResultData{},
		// 	IsError:           true,
		// 	Error:             resources.ServiceError{ErrorMsg: err.Error()},
		// }
		return nil, err
	}

	return result, nil
}

func GetEmployeeById(ctx *gin.Context, employeeId int) (models.Employee, error) {
	result := models.Employee{}
	queryPtr := DbClient.Debug().Table("employee").First(&result, employeeId)

	if err := queryPtr.Error; err != nil {
		fmt.Print("Error updating Employee Details | ", queryPtr.Error.Error())
	}
	if result.EmployeeId == 0 {
		fmt.Print("Error finding Employee Details")
		return result, errors.New("Error finding Employee Details") 
	}
	return result, nil
}

func UpdateEmployee(ctx *gin.Context, employeeId int, input resources.AddEmployeeInput) error {
	data := models.Employee{
		EmployeeName:   input.EmployeeName,
		EmployeeRole:   input.EmployeeRole,
		EmployeeMail:   input.EmployeeMail,
		EmployeeAge:    uint(input.EmployeeAge),
		EmployeeGender: input.EmployeeGender,
	}

	queryPtr := DbClient.Debug().Table("employee").Where("employee_id = ?", employeeId).Updates(&data)

	if err := queryPtr.Error; err != nil {
		fmt.Print("Error updating Employee Details | ", queryPtr.Error.Error())
		return err
	}

	return nil
}

func DeleteEmployee(ctx *gin.Context, employeeId int) error {
	data := models.Employee{
		EmployeeId: uint64(employeeId),
	}
	queryPtr := DbClient.Debug().Table("employee").Where("employee_id =?", employeeId).Delete(&data)

	if err := queryPtr.Error; err != nil {
		fmt.Println("Error Deleting Employee Details | ", queryPtr.Error.Error())
		return err
	}
	return nil
}
