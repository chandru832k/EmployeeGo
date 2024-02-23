package controller

import (
	"employee/resources"
	"employee/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	ProjectService service.ProjectService
}


func NewProjectController(ProjectService service.ProjectService) ProjectController {
	return ProjectController{
		ProjectService: ProjectService,
	}
}

func (pc *ProjectController) AddProject(c *gin.Context) {
	functionDesc := "Add Project Controller"
	log.Println(functionDesc)

	var input resources.ProjectInput

	err := c.BindJSON(&input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error Reading Input",
		})
		return
	}

	projectId := pc.ProjectService.AddProject(c, input)

	if projectId == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error Creating Project",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"project_id": projectId,
		"message": "project created successfully",
	})
}

func  (pc *ProjectController) EditProject(c *gin.Context) {
	functionDesc := "Edit Project Controller"
	log.Println(functionDesc)

	var input resources.ProjectInput
	err := c.BindJSON(&input)
	projectId, err := strconv.Atoi(c.Param("project_id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error Reading Input",
		})
		return
	}

	result := pc.ProjectService.EditProject(c, input, uint64(projectId))

	if result != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error Editing Project",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "project updated successfully",
	})
}

func (pc *ProjectController) GetAllProjects(c *gin.Context) {
	functionDesc := "Get All Projects"
	log.Println(functionDesc)

	result := pc.ProjectService.GetAllProjects(c)

	if result.Code != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error Fetching Project",
		})
		return
	}
	c.JSON(http.StatusOK, result.ServiceResultData.Data)
}

func (pc *ProjectController) GetProjectById(c *gin.Context) {
	functionDesc := "Get Project By Id"
	log.Println(functionDesc)

	projectId, err := strconv.Atoi(c.Param("project_id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error Reading Project Id",
		})
		return
	}

	result := pc.ProjectService.GetProjectById(c, uint64(projectId))

	if result.Code != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error Fetching Project",
		})
		return
	}

	c.JSON(http.StatusOK, result.ServiceResultData.Data)
}

func  (pc *ProjectController) DeleteProject(c *gin.Context) {
	functionDesc := "Delete Project"
	log.Println(functionDesc)

	projectId, err := strconv.Atoi(c.Param("project_id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error Reading Project Id",
		})
		return
	}

	result := pc.ProjectService.DeleteProject(c, uint64(projectId))

	if result != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error Deleting Project",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Project Deleted Successfully",
	})
}

