package service

import (
	"employee/db"
	"employee/resources"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmployeeService interface {
	AddEmployee(ctx *gin.Context, input resources.AddEmployeeInput) uint64
	GetAllEmployees(ctx *gin.Context) resources.ServiceResult
	GetEmployeeById(ctx *gin.Context, employeeId int) resources.ServiceResult
	DeleteEmployee(ctx *gin.Context, employeeId int) error
	UpdateEmployee(ctx *gin.Context, employeeId int, input resources.AddEmployeeInput) error
}
type employeeService struct {
	EmployeeDao db.EmployeeDao
}

func NewEmployeeService(EmployeeDao db.EmployeeDao) *employeeService {
	return &employeeService{
		EmployeeDao: EmployeeDao,
	}
}

func (es *employeeService) AddEmployee(ctx *gin.Context, input resources.AddEmployeeInput) uint64 {
	functionDesc := "Add Employee Service"
	fmt.Println(functionDesc)
	return es.EmployeeDao.AddEmployee(ctx, input)
}

func (es *employeeService) GetAllEmployees(ctx *gin.Context) resources.ServiceResult {
	result, err := es.EmployeeDao.GetAllEmployees(ctx)
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

func (es *employeeService) GetEmployeeById(ctx *gin.Context, employeeId int) resources.ServiceResult {
	result, err := es.EmployeeDao.GetEmployeeById(ctx, employeeId)
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

func (es *employeeService) DeleteEmployee(ctx *gin.Context, employeeId int) error {
	return es.EmployeeDao.DeleteEmployee(ctx, employeeId)
}

func (es *employeeService) UpdateEmployee(ctx *gin.Context, employeeId int, input resources.AddEmployeeInput) error {
	return es.EmployeeDao.UpdateEmployee(ctx, employeeId, input)
}
