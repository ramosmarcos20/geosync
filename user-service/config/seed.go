package config

import (
	"geosync/user-service/internal/models"
	"geosync/user-service/internal/utils"
	"log"
)

// Check if the database is empty and seed data if necessary
func SeedDatabase() {
	seedRoles()
	seedPermissions()
	seedAdminUser()
}

func seedRoles() {
	var count int64
	DB.Model(&models.Role{}).Count(&count)
	if count == 0 {
		roles := []models.Role{
			{Name: "admin"},
			{Name: "user"},
		}
		for _, role := range roles {
			if err := DB.Create(&role).Error; err != nil {
				log.Fatalf("failed to seed roles: %v", err)
			}
		}
	}
}

func seedPermissions() {
	var count int64
	DB.Model(&models.Permission{}).Count(&count)
	if count == 0 {
		permissions := []models.Permission{
			{Name: "read"},
			{Name: "write"},
			{Name: "delete"},
		}
		for _, permission := range permissions {
			if err := DB.Create(&permission).Error; err != nil {
				log.Fatalf("failed to seed permissions: %v", err)
			}
		}
	}
}

func seedAdminUser() {
	var count int64
	DB.Model(&models.User{}).Count(&count)
	if count == 0 {
		password, _ := utils.HashPassword("admin123") // Hashear la contraseña
		adminUser := models.User{
			Email:    "admin@admin.com",
			UserName: "admin",
			Password: password,
			RoleID:   1, // Suponiendo que el ID del rol admin es 1
		}
		if err := DB.Create(&adminUser).Error; err != nil {

			log.Fatalf("failed to seed admin user: %v", err)
		}
	}
}
