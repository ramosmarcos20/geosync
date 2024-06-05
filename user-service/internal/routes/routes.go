package routes

import (
	"geosync/user-service/internal/controllers"
	"geosync/user-service/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Aplicar middleware de autenticación
	router.Use(middlewares.AuthMiddleware())

	// Definir rutas
	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("/", controllers.CreateUser)
			users.GET("/", controllers.GetUsers)
			users.GET("/:id", controllers.GetUser)
		}
	}

	return router
}
