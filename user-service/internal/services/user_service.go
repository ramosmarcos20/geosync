package services

import (
    "errors"
    "geosync/user-service/internal/models"
    "geosync/user-service/internal/repositories"
    "geosync/user-service/internal/utils"
    "log"
)

type UserService struct {
    repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *models.User) []error {
    var errs []error

    // Validar que el email no exista ya en la base de datos
    existingUserByEmail, err := s.repo.GetUserByEmail(user.Email)
    if err != nil && !errors.Is(err, utils.UserNotFound) {
        log.Printf("Error checking email: %v", err)
        errs = append(errs, err)
    }
    if existingUserByEmail != nil {
        errs = append(errs, utils.EmailInUse)
    }

    // Validar que el nombre de usuario no exista ya en la base de datos
    existingUserByUserName, err := s.repo.GetUserByUserName(user.UserName)
    if err != nil && !errors.Is(err, utils.UserNotFound) {
        log.Printf("Error checking username: %v", err)
        errs = append(errs, err)
    }
    if existingUserByUserName != nil {
        errs = append(errs, utils.UserNameInUse)
    }

    if len(errs) > 0 {
        return errs
    }

    // Crear el usuario en la base de datos
    err = s.repo.CreateUser(user)
    if err != nil {
        log.Printf("Error creating user in repository: %v", err)
        errs = append(errs, err)
    }
    return errs
}

func (s *UserService) UpdateUser(user *models.User) []error {
    var errs []error

    // Validar que el email no exista ya en la base de datos para otro usuario
    existingUserByEmail, err := s.repo.GetUserByEmail(user.Email)
    if err != nil && !errors.Is(err, utils.UserNotFound) {
        log.Printf("Error finding user by email: %v", err)
        errs = append(errs, err)
    }
    if existingUserByEmail != nil && existingUserByEmail.ID != user.ID {
        log.Printf("Email already in use: %s", user.Email)
        errs = append(errs, utils.EmailInUse)
    }

    // Validar que el nombre de usuario no exista ya en la base de datos para otro usuario
    existingUserByUserName, err := s.repo.GetUserByUserName(user.UserName)
    if err != nil && !errors.Is(err, utils.UserNotFound) {
        log.Printf("Error finding user by username: %v", err)
        errs = append(errs, err)
    }
    if existingUserByUserName != nil && existingUserByUserName.ID != user.ID {
        log.Printf("Username already in use: %s", user.UserName)
        errs = append(errs, utils.UserNameInUse)
    }

    if len(errs) > 0 {
        return errs
    }

    // Actualizar el usuario en la base de datos
    err = s.repo.UpdateUser(user)
    if err != nil {
        log.Printf("Error updating user in repository: %v", err)
        errs = append(errs, utils.ErrorDataBase)
    }
    return errs
}

func (s *UserService) GetUsers() []models.User {
    return s.repo.GetUsers()
}

func (s *UserService) GetUser(id string) (*models.User, error) {
    user, err := s.repo.GetUser(id)
    if err != nil {
        log.Printf("Error getting user: %v", err)
        if errors.Is(err, utils.UserNotFound) {
            return nil, utils.UserNotFound
        }
        return nil, utils.ErrorDataBase
    }
    return user, nil
}

func (s *UserService) DeleteUser(id uint) error {
    err := s.repo.DeleteUser(id)
    if err != nil {
        log.Printf("Error deleting user in repository: %v", err)
        return utils.ErrorDataBase
    }
    return nil
}

func (s *UserService) GetUserAuth(id uint) (*models.User, error) {
    user, err := s.repo.GetUserAuth(id)
    if err != nil {
        log.Printf("Error getting user by ID: %v", err)
        if errors.Is(err, utils.UserNotFound) {
            return nil, utils.UserNotFound
        }
        return nil, utils.ErrorDataBase
    }
    return user, nil
}
