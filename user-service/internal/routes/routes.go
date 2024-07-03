package routes

import (
    "geosync/user-service/internal/controllers"
    "geosync/user-service/internal/middlewares"
    "geosync/user-service/internal/repositories"
    "geosync/user-service/internal/services"
    "geosync/user-service/config"
    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    // Ruta de inicio de sesión sin prefijo
    router.POST("/login", controllers.Login)

    // Crear instancias del repositorio y servicio
    db := config.DB
    userRepo := repositories.NewUserRepository(db)
    userService := services.NewUserService(userRepo)
    userController := controllers.NewUserController(userService)

    tenantRepo := repositories.NewTenantRepository(db)
    tenantService := services.NewTenantService(tenantRepo)
    tenantController := controllers.NewTenantController(tenantService)

    // Rutas protegidas
    root := router.Group("/")
    root.Use(middlewares.AuthMiddleware())
    {
        users := root.Group("/user")
        {
            users.GET("/", userController.GetUsers)
            users.POST("/create", userController.CreateUser)
            users.GET("/edit/:id", userController.GetUser)
            users.POST("/update/:id", userController.UpdateUser)
            users.DELETE("/delete/:id", userController.DeleteUser)
            users.GET("/profile", userController.UserAuth)
        }

        tenants := root.Group("/tenant")
        {
            tenants.GET("/", tenantController.Index)
            tenants.POST("/create", tenantController.CreateTenant)
            tenants.GET("/edit/:id", tenantController.GetTenant)
        }
    }

    return router
}
