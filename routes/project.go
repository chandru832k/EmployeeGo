package routes

import (
	"github.com/gin-gonic/gin"
)

func ListRoutes(router *gin.RouterGroup) {
	projectGroup := router.Group("/project")
	{
		projectGroup.POST("/", ProjectController.AddProject)
		projectGroup.PUT("/:project_id", ProjectController.EditProject)
		projectGroup.GET("/", ProjectController.GetAllProjects)
		projectGroup.GET("/:project_id", ProjectController.GetProjectById)
		projectGroup.DELETE("/:project_id", ProjectController.DeleteProject)
	}
}
