package config

import (
	"geosync/user-service/internal/models"
	"geosync/user-service/internal/utils"
	"log"
)

func SeedDatabase() {
	seedTenant()
	seedRoles()
	seedPermissions()
	seedPlans()
	seedAdminUser()
}

func seedTenant() {
	var count int64
	DB.Model(&models.Tenant{}).Count(&count)
	if count == 0 {
		tenant := models.Tenant{
			Name:      "Marcos Ramos",
			TaxNumber: "0953331675001",
			Email:     "marcos@marcos.com",
			Phone:     "+593968319032",
			Address:   "Guayaquil, Ecuador",
		}
		if err := DB.Create(&tenant).Error; err != nil {
			log.Fatalf("failed to seed tenant: %v", err)
		}
	}
}

func seedRoles() {
	var count int64
	DB.Model(&models.Role{}).Count(&count)
	if count == 0 {
		roles := []models.Role{
			{Name: "superAdmin", Description: "Acceso primordial"},
			{Name: "admin", Description: "Administrador de empresa"},
			{Name: "user", Description: "Usuario común"},
			{Name: "technical", Description: "Técnico para soporte"},
			{Name: "support", Description: "Soporte, servicio al cliente"},
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

func seedPlans() {
	var count int64
	DB.Model(&models.Plan{}).Count(&count)
	if count == 0 {
		plans := []models.Plan{
			{Name: "Basic", IsFree: true, NumberUser: 2, NumberDevice: 2},
			{Name: "Standard", IsFree: false, NumberUser: 0, NumberDevice: 5},
			{Name: "Premium", IsFree: false, NumberUser: 0, NumberDevice: 0},
		}
		for _, plan := range plans {
			if err := DB.Create(&plan).Error; err != nil {
				log.Fatalf("failed to seed plans: %v", err)
			}
		}
	}
}

func seedAdminUser() {
	var count int64
	DB.Model(&models.User{}).Count(&count)
	if count == 0 {
		password, _ := utils.HashPassword("superadmin") // Hashear la contraseña
		adminUser := models.User{
			Email:    "superadmin@correo.com",
			UserName: "superadmin",
			Password: password,
			RoleID:   1,
			PlanId:   1,
			TenantID: 1,
		}
		if err := DB.Create(&adminUser).Error; err != nil {
			log.Fatalf("failed to seed admin user: %v", err)
		}
	}
}
