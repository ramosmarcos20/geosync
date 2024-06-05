package main

import (
	"geosync/user-service/config"
	"geosync/user-service/internal/routes"
)

func main() {
	config.ConnectDatabase()
	config.RunMigrations()
	config.SeedDatabase() // Llamar a la función de prellenado de datos

	// Configurar el enrutador
	router := routes.SetupRouter()

	// Iniciar el servidor
	router.Run(":8080")
}
