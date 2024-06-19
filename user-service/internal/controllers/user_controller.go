package controllers

import (
	"errors"
	"geosync/user-service/internal/models"
	"geosync/user-service/internal/services"
	"geosync/user-service/internal/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{service: service}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.InvalidInput.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = hashedPassword

	errs := uc.service.CreateUser(&user)
	if len(errs) > 0 {
		var errorMessages []string
		for _, e := range errs {
			errorMessages = append(errorMessages, e.Error())
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		log.Printf("Errors creating user: %v", errs)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario creado exitosamente", "user": user})
}

func (uc *UserController) GetUsers(c *gin.Context) {
	users := uc.service.GetUsers()
	c.JSON(http.StatusOK, users)
}

func (uc *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := uc.service.GetUser(id)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		if errors.Is(err, utils.UserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": utils.UserNotFound.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": utils.ErrorDataBase.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Printf("Error parsing user ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.InvalidInput.Error()})
		return
	}

	user.ID = uint(uid)
	errs := uc.service.UpdateUser(&user)
	if len(errs) > 0 {
		var errorMessages []string
		for _, e := range errs {
			errorMessages = append(errorMessages, e.Error())
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		log.Printf("Errors updating user: %v", errs)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado exitosamente", "user": user})
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Printf("Error parsing user ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = uc.service.DeleteUser(uint(uid))
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": utils.ErrorDataBase.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado exitosamente"})
}

func (uc *UserController) UserAuth(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user from context"})
		return
	}

	user, err := uc.service.GetUserAuth(userID.(uint))
	if err != nil {
		log.Printf("Error getting user: %v", err)
		if errors.Is(err, utils.UserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": utils.UserNotFound.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": utils.ErrorDataBase.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}
