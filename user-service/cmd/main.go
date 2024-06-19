package main

import (
	"geosync/user-service/config"
	"geosync/user-service/internal/routes"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config.ConnectDatabase()
	config.RunMigrations()
	config.SeedDatabase()

	router := routes.SetupRouter()

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
