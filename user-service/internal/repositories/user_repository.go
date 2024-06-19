package repositories

import (
	"geosync/user-service/config"
	"geosync/user-service/internal/models"
	"geosync/user-service/internal/utils"
	"log"

	"gorm.io/gorm"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	err := config.DB.Create(user).Error
	if err != nil {
		log.Printf("Error creating user in DB: %v", err)
	}
	return err
}

func (r *UserRepository) GetUsers() []models.User {
	var users []models.User
	err := config.DB.Find(&users).Error
	if err != nil {
		log.Printf("Error getting users from DB: %v", err)
	}
	return users
}

func (r *UserRepository) GetUser(id string) (*models.User, error) {
	var user models.User
	err := config.DB.First(&user, id).Error
	if err != nil {
		log.Printf("Error getting user from DB: %v", err)
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	err := config.DB.Where("email = ?", email).First(&user).Error

	if err != nil {

		if err == gorm.ErrRecordNotFound {
			log.Printf("User with email %s not found", email)
			return nil, utils.UserNotFound
		}
		log.Printf("Error getting user by email from DB: %v", err)
		return nil, utils.ErrorDataBase
	}

	return &user, nil
}

func (r *UserRepository) GetUserByUserName(userName string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("user_name = ?", userName).First(&user).Error

	if err != nil {

		if err == gorm.ErrRecordNotFound {
			log.Printf("User with username %s not found", userName)
			return nil, utils.UserNotFound
		}

		log.Printf("Error getting user by username from DB: %v", err)
		return nil, utils.ErrorDataBase
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	err := config.DB.Save(user).Error
	if err != nil {
		log.Printf("Error updating user in DB: %v", err)
	}
	return err
}

func (r *UserRepository) DeleteUser(id uint) error {
	err := config.DB.Unscoped().Delete(&models.User{}, id).Error
	if err != nil {
		log.Printf("Error deleting user in DB: %v", err)
	}
	return err
}

func (r *UserRepository) GetUserAuth(id uint) (*models.User, error) {
	var user models.User
	err := config.DB.First(&user, id).Error
	if err != nil {
		log.Printf("Error getting user by ID from DB: %v", err)
		return nil, err
	}
	return &user, nil
}
