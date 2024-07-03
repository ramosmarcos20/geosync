package repositories

import (
    "geosync/user-service/internal/models"
    "geosync/user-service/internal/utils"
    "log"

    "gorm.io/gorm"
)

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *models.User) error {
    err := r.db.Create(user).Error
    if err != nil {
        log.Printf("Error creating user in DB: %v", err)
    }
    return err
}

func (r *userRepository) GetUsers() []models.User {
    var users []models.User
    err := r.db.Find(&users).Error
    if err != nil {
        log.Printf("Error getting users from DB: %v", err)
    }
    return users
}

func (r *userRepository) GetUser(id string) (*models.User, error) {
    var user models.User
    err := r.db.First(&user, id).Error
    if err != nil {
        log.Printf("Error getting user from DB: %v", err)
        return nil, err
    }
    return &user, nil
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
    var user models.User
    err := r.db.Where("email = ?", email).First(&user).Error
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

func (r *userRepository) GetUserByUserName(userName string) (*models.User, error) {
    var user models.User
    err := r.db.Where("user_name = ?", userName).First(&user).Error
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

func (r *userRepository) UpdateUser(user *models.User) error {
    err := r.db.Save(user).Error
    if err != nil {
        log.Printf("Error updating user in DB: %v", err)
    }
    return err
}

func (r *userRepository) DeleteUser(id uint) error {
    err := r.db.Unscoped().Delete(&models.User{}, id).Error
    if err != nil {
        log.Printf("Error deleting user in DB: %v", err)
    }
    return err
}

func (r *userRepository) GetUserAuth(id uint) (*models.User, error) {
    var user models.User
    err := r.db.First(&user, id).Error
    if err != nil {
        log.Printf("Error getting user by ID from DB: %v", err)
        return nil, err
    }
    return &user, nil
}
