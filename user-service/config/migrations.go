package config

import (
	"geosync/user-service/internal/models"
)

func RunMigrations() {
	DB.AutoMigrate(
		&models.Tenant{}, &models.Plan{},
		&models.User{}, &models.Role{},
		&models.Permission{}, &models.RolePermission{},
	)
}
