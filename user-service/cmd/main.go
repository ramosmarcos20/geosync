package main

import (
	"geosync/user-service/config"
	"geosync/user-service/internal/controllers"
	"geosync/user-service/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config.ConnectDatabase()
	router.Use(middlewares.AuthMiddleware())

	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("/", controllers.CreateUser)
			users.GET("/", controllers.GetUsers)
			users.GET("/:id", controllers.GetUser)
		}
	}

	router.Run(":8080")
}
