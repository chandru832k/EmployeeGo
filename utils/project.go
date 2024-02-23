package utils

import (
	"employee/models"
	"employee/resources"
	"time"

	"github.com/gin-gonic/gin"
)

func MapCreateProjectPayloadFields(c *gin.Context, input resources.ProjectInput, ) models.Project {
	data := models.Project{
		Name:        input.Name,
		Status:      input.Status,
		Description: input.Description,
		CreatedBy:   c.GetHeader("email"),
		CreatedAt:   uint64(time.Now().Unix()),
	}
	return data
}

func MapEditProjectPayloadFields(c *gin.Context, input resources.ProjectInput) models.Project {
	data := models.Project{
		Name:        input.Name,
		Status:      input.Status,
		Description: input.Description,
		UpdatedBy:   c.GetHeader("email"),
		UpdatedAt:   uint64(time.Now().Unix()),
	}
	return data
}