package service

import (
	"context"
	"employee/db"
	"employee/resources"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmployeeService struct {
	EmployeeDao db.EmployeeDao
}

func NewEmployeeService(EmployeeDao db.EmployeeDao) *EmployeeService {
	return &EmployeeService{
		EmployeeDao: EmployeeDao,
	}
}

func AddEmployee(ctx context.Context, input resources.AddEmployeeInput) uint64 {
	return db.AddEmployee(ctx, input)
}

func GetAllEmployees(ctx *gin.Context) resources.ServiceResult {
	result, err := db.GetAllEmployees(ctx)
	if err != nil {
		return resources.ServiceResult{
			Code:              http.StatusInternalServerError,
			ServiceResultData: resources.ServiceResultData{},
			IsError:           true,
			Error:             resources.ServiceError{ErrorMsg: err.Error()},
		}
	}
	return resources.ServiceResult{
		Code:              http.StatusOK,
		ServiceResultData: resources.ServiceResultData{Data: result},
		IsError:           false,
	}
}

func DeleteEmployee(ctx *gin.Context, employeeId int) error {
	return db.DeleteEmployee(ctx, employeeId)
}

func UpdateEmployee(ctx *gin.Context, employeeId int, input resources.AddEmployeeInput) error {
	return db.UpdateEmployee(ctx, employeeId, input)
}
