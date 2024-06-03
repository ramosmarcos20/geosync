package services

import (
	"geosync/user-service/internal/models"
	"geosync/user-service/internal/repositories"
)

func CreateUser(user *models.User) error {
	// Add any additional business logic here
	return repositories.CreateUser(user)
}

func GetUsers() []models.User {
	// Add any additional business logic here
	return repositories.GetUsers()
}

func GetUser(id string) (*models.User, error) {
	// Add any additional business logic here
	return repositories.GetUser(id)
}
