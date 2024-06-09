package routes

import (
	"geosync/user-service/internal/controllers"
	"geosync/user-service/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Ruta de inicio de sesión sin prefijo
	router.POST("/login", controllers.Login)

	// Rutas protegidas
	root := router.Group("/")
	root.Use(middlewares.AuthMiddleware())
	{
		users := root.Group("/users")
		{
			users.POST("/create", controllers.CreateUser)  // Crear un nuevo usuario
			users.GET("/list", controllers.GetUsers)       // Obtener lista de usuarios
			users.GET("/details/:id", controllers.GetUser) // Obtener detalles de un usuario específico
		}
	}

	return router
}
