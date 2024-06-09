package controllers

import (
	"geosync/user-service/config"
	"geosync/user-service/internal/models"
	"geosync/user-service/internal/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Login attempt for email: %s", input.Email)

	var user models.User
	// Log before the query
	log.Println("Attempting to query user from database")
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		// Log the detailed error
		log.Printf("Error finding user: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	log.Println("User found, checking password")

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
