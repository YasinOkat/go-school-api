package routes

import (
	"github.com/YasinOkat/go-school-api/controllers"
	"github.com/YasinOkat/go-school-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterStudentRoutes(router *gin.Engine) {
	studentRoutes := router.Group("/students").Use(middlewares.AuthMiddleware())
	{
		studentRoutes.POST("/", middlewares.AdminMiddleware(), controllers.CreateStudent)
	}
}
