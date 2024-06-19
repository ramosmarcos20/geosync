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
			userController := controllers.NewUserController()
			users.POST("/create", userController.CreateUser)       // Crear un nuevo usuario
			users.GET("/list", userController.GetUsers)            // Obtener lista de usuarios
			users.GET("/details/:id", userController.GetUser)      // Obtener detalles de un usuario específico
			users.POST("/update/:id", userController.UpdateUser)   // Actualizar un usuario
			users.DELETE("/delete/:id", userController.DeleteUser) // Eliminar un usuario
			users.GET("/profile", userController.UserAuth)         // Obtener perfil del usuario autenticado
		}

		tenants := root.Group("/tenants")
		{
			tenantController := controllers.NewTenantController()
			tenants.GET("/index", tenantController.Index)
		}
	}


	return router
}
