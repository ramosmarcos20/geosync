package repositories

import (
	"geosync/user-service/config"
	"geosync/user-service/internal/models"
)

func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

func GetUsers() []models.User {
	var users []models.User
	config.DB.Find(&users)
	return users
}

func GetUser(id string) (*models.User, error) {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
