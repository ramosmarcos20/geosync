package main

import (
	"geosync/user-service/config"
	"geosync/user-service/internal/routes"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// Cargar las variables de entorno desde el archivo .env

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Conectar a la base de datos
	config.ConnectDatabase()

	// Ejecutar migraciones
	config.RunMigrations()

	// Llamar a la función de prellenado de datos
	config.SeedDatabase()

	// Configurar el enrutador
	router := routes.SetupRouter()

	// Iniciar el servidor
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
