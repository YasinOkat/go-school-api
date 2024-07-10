package main

import (
	_ "github.com/YasinOkat/go-school-api/docs"
	"github.com/YasinOkat/go-school-api/routes"
	"github.com/YasinOkat/go-school-api/utils"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           School Management API
// @version         1.0
// @description     This is a sample school management server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   https://www.gnu.org/licenses/gpl-3.0.en.html

// @host      localhost:3002
// @BasePath  /

func main() {
	LoadConfig()

	utils.ConnectDatabase()

	router := gin.Default()

	routes.RegisterUserRoutes(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":3002")
}
