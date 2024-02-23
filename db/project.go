package db

import (
	"employee/models"
	"employee/resources"
	"employee/utils"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type projectDao struct {
	db *gorm.DB
}

type ProjectDao interface {
	AddProject(c *gin.Context, input resources.ProjectInput) uint64
	EditProject(c *gin.Context, input resources.ProjectInput, projectId uint64) error
	GetAllProjects(c *gin.Context) ([]models.Project, error)
	GetProjectById(c *gin.Context, projectId uint64) (models.Project, error)
	DeleteProject(c *gin.Context, projectId uint64) error 
}

func NewProjectDao() *projectDao {
	return &projectDao{
		db: DbClient,
	}
}

func (pd *projectDao) AddProject(c *gin.Context, input resources.ProjectInput) uint64 {
	data := utils.MapCreateProjectPayloadFields(c, input)

	fmt.Println(data)

	db := pd.db.WithContext(c)

	queryPtr := db.Debug().Table("project").Create((&data))
	if queryPtr.Error != nil {
		log.Println("Error creating project in DB  | ", queryPtr.Error.Error())
		return 0
	}
	return data.ProjectId
}

func (pd *projectDao) EditProject(c *gin.Context, input resources.ProjectInput, projectId uint64) error {
	db := pd.db.WithContext(c)

	data := utils.MapEditProjectPayloadFields(c, input)

	queryPtr := db.Debug().Table("Project").Where("project_id = ?", projectId).Updates(&data)

	if err := queryPtr.Error; err != nil {
		log.Println("Error updating project in DB  | ", queryPtr.Error.Error())
		return err
	}
	return nil
}

func (pd *projectDao) GetAllProjects(c *gin.Context) ([]models.Project, error) {
	db := pd.db.WithContext(c)

	result := []models.Project{}

	queryPtr := db.Debug().Table("project").Find(&result)

	if err := queryPtr.Error; err != nil {
		log.Println("Error fetching projects in DB  | ", queryPtr.Error.Error())
		return nil, err
	}

	return result, nil
}

func (pd *projectDao) GetProjectById(c *gin.Context, projectId uint64) (models.Project, error) {
	db := pd.db.WithContext(c)

	result := models.Project{}

	queryPtr := db.Debug().Table("project").Where("project_id = ?", projectId).Find(&result)

	if err := queryPtr.Error; err != nil {
		log.Println("Error fetching projects in DB  | ", queryPtr.Error.Error())
		return result, err
	}
	return result, nil
}

func (pd *projectDao) DeleteProject(c *gin.Context, projectId uint64) error {
	db := pd.db.WithContext(c)

	queryPtr := db.Table("project").Delete(&models.Project{}, projectId)

	if err := queryPtr.Error; err != nil {
		log.Println("Error Deleting project in DB  | ", queryPtr.Error.Error())
		return err
	}

	return nil
}

