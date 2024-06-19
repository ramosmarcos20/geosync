package repositories

import (
	"geosync/user-service/config"
	"geosync/user-service/internal/models"
	"log"
)

type TenantRepository struct{}

func NewTenantRepository() *TenantRepository {
	return &TenantRepository{}
}

func (r *TenantRepository) Index() []models.Tenant {
	var tenant []models.Tenant
	err := config.DB.Find(&tenant).Error
	if err != nil {
		log.Printf("Error en obtener tenant: %v", err)
	}

	return tenant
}

