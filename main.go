package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/api"
	"main/controllers"
	"main/database"
	"main/middlewares"
	"main/utils"
)

func main() {
	fmt.Println("Starting server")

	// Initialize Database
	database.Connect(utils.GoDotEnvVariable("databaseUrl"))
	//database.Migrate()

	// Initialize Router
	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	apis := router.Group("/api")
	{
		apis.POST("/token", controllers.GenerateToken)
		apis.POST("/user/register", controllers.RegisterUser)
		secured := apis.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
			secured.POST("/mpesa-express", api.MpesaExpress)
		}
	}
	return router
}
