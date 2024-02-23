package service

import (
	"employee/db"
	"employee/resources"
	"net/http"

	"github.com/gin-gonic/gin"
)

type projectService struct {
	ProjectDao db.ProjectDao
}

type ProjectService interface {
	AddProject(c *gin.Context, input resources.ProjectInput) uint64
	EditProject(c *gin.Context, input resources.ProjectInput, projectId uint64) error
	GetAllProjects(c *gin.Context) resources.ServiceResult
	GetProjectById(c *gin.Context, projectId uint64) resources.ServiceResult
	DeleteProject(c *gin.Context, projectId uint64) error
}

func NewProjectService(ProjectDao db.ProjectDao) *projectService {
	return &projectService{
		ProjectDao: ProjectDao,
	}
}

func (ps *projectService) AddProject(c *gin.Context, input resources.ProjectInput) uint64{
	return ps.ProjectDao.AddProject(c, input)
}

func (ps *projectService) EditProject(c *gin.Context, input resources.ProjectInput, projectId uint64) error {
	return ps.ProjectDao.EditProject(c, input, projectId)
} 

func (ps *projectService) GetAllProjects(c *gin.Context) resources.ServiceResult {
	result, err := ps.ProjectDao.GetAllProjects(c)
	if (err != nil) {
		return resources.ServiceResult{
			Code:              http.StatusInternalServerError,
			// ServiceResultData: resources.ServiceResultData{},
			IsError:           true,
			Error:             resources.ServiceError{ErrorMsg: err.Error()},
		}
	}
	return resources.ServiceResult{
		Code: http.StatusOK,
		IsError: false,
		ServiceResultData: resources.ServiceResultData{
			Data: result,
		},
	}
}

func (ps *projectService) GetProjectById(c *gin.Context, projectId uint64) resources.ServiceResult {
	result, err := ps.ProjectDao.GetProjectById(c, projectId)
	if (err != nil) {
		return resources.ServiceResult{
			Code:              http.StatusInternalServerError,
			// ServiceResultData: resources.ServiceResultData{},
			IsError:           true,
			Error:             resources.ServiceError{ErrorMsg: err.Error()},
		}
	}
	return resources.ServiceResult{
		Code: http.StatusOK,
		IsError: false,
		ServiceResultData: resources.ServiceResultData{
			Data: result,
		},
	}
}

func (ps *projectService) DeleteProject (c *gin.Context, projectId uint64) error {
	return ps.ProjectDao.DeleteProject(c, projectId)
}