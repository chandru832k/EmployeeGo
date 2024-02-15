package routes

import "github.com/gin-gonic/gin"

func CreateRoutes(router *gin.RouterGroup) {
	EmployeeRoutes(router);
}