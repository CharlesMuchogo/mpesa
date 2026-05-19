package main

import (
	"main/api"
	"main/controllers"
	"main/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	//connectionString := database.GetPostgresConnectionString()
	// Initialize Database
	//database.Connect(connectionString)
	//database.Migrate()
	// Initialize Router
	router := initRouter()
	err := router.Run(":8000")
	if err != nil {
		panic(err)
		return
	}
}

func initRouter() *gin.Engine {
	router := gin.Default()
	apis := router.Group("/api")
	apis.POST("/mpesa-express", api.MpesaExpress)
	apis.POST("/callback", api.CallbackUrl)
	{
		apis.POST("/token", controllers.GenerateToken)
		apis.POST("/user/register", controllers.RegisterUser)
		secured := apis.Group("/secured").Use(middlewares.Auth())
		{
			secured.POST("/mpesa-express", api.MpesaExpress)
		}
	}
	return router
}
