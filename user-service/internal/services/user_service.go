package services

import (
	"errors"
	"geosync/user-service/internal/models"
	"geosync/user-service/internal/repositories"
	"log"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		repo: repositories.NewUserRepository(),
	}
}

func (s *UserService) CreateUser(user *models.User) error {
	// Validar que el email no exista ya en la base de datos
	existingUserByEmail, err := s.repo.GetUserByEmail(user.Email)
	if err == nil && existingUserByEmail != nil {
		log.Printf("Email already in use: %s", user.Email)
		return errors.New("email already in use")
	}

	// Validar que el nombre de usuario no exista ya en la base de datos
	existingUserByUserName, err := s.repo.GetUserByUserName(user.UserName)
	if err == nil && existingUserByUserName != nil {
		log.Printf("Username already in use: %s", user.UserName)
		return errors.New("username already in use")
	}

	// Crear el usuario en la base de datos
	err = s.repo.CreateUser(user)
	if err != nil {
		log.Printf("Error creating user in repository: %v", err)
	}
	return err
}

func (s *UserService) GetUsers() []models.User {
	users := s.repo.GetUsers()
	return users
}

func (s *UserService) GetUser(id string) (*models.User, error) {
	user, err := s.repo.GetUser(id)
	if err != nil {
		log.Printf("Error getting user: %v", err)
	}
	return user, err
}

func (s *UserService) UpdateUser(user *models.User) error {
	// Validar que el email no exista ya en la base de datos para otro usuario
	existingUserByEmail, err := s.repo.GetUserByEmail(user.Email)
	if err != nil && err.Error() != "record not found" {
		log.Printf("Error finding user by email: %v", err)
		return err
	}
	if existingUserByEmail != nil && existingUserByEmail.ID != user.ID {
		log.Printf("Email already in use: %s", user.Email)
		return errors.New("email already in use")
	}

	// Validar que el nombre de usuario no exista ya en la base de datos para otro usuario
	existingUserByUserName, err := s.repo.GetUserByUserName(user.UserName)
	if err != nil && err.Error() != "record not found" {
		log.Printf("Error finding user by username: %v", err)
		return err
	}
	if existingUserByUserName != nil && existingUserByUserName.ID != user.ID {
		log.Printf("Username already in use: %s", user.UserName)
		return errors.New("username already in use")
	}

	// Actualizar el usuario en la base de datos
	err = s.repo.UpdateUser(user)
	if err != nil {
		log.Printf("Error updating user in repository: %v", err)
		return err
	}
	return nil
}

func (s *UserService) DeleteUser(id uint) error {
	err := s.repo.DeleteUser(id)
	if err != nil {
		log.Printf("Error deleting user in repository: %v", err)
	}
	return err
}

func (s *UserService) GetUserAuth(id uint) (*models.User, error) {
	user, err := s.repo.GetUserAuth(id)
	if err != nil {
		log.Printf("Error getting user by ID: %v", err)
		return nil, err
	}
	return user, nil
}
