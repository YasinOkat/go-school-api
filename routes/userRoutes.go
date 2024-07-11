package routes

import (
	"github.com/YasinOkat/go-school-api/controllers"
	"github.com/YasinOkat/go-school-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	userRoutes := router.Group("/users").Use(middlewares.AuthMiddleware())
	{
		userRoutes.POST("/", middlewares.AdminMiddleware(), controllers.CreateUser)
		userRoutes.GET("/", middlewares.AdminMiddleware(), controllers.GetUsers)
		userRoutes.DELETE("/users/:id", middlewares.AdminMiddleware(), controllers.DeleteUser)

	}

	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/login", controllers.Login)
	}

}
