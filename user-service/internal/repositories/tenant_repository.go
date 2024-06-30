package repositories

import (
	"geosync/user-service/config"
	"geosync/user-service/internal/models"
	"geosync/user-service/internal/utils"
	"log"

	"gorm.io/gorm"
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

func (r *TenantRepository) CreateTenant(tenant *models.Tenant) error   {
	err := config.DB.Create(tenant).Error
	if err != nil {
		log.Printf("Error creating user in DB: %v", err)
	}
	return err
}

/* VALIDACIONES */
func (r *TenantRepository) ExistTenant(tax_number string) (*models.Tenant, error)  {
	var tenant models.Tenant

	exist := config.DB.Where("tax_number = ?", tax_number).First(&tenant).Error

	if exist != nil {

		if exist == gorm.ErrRecordNotFound {
			log.Printf("Tenant with Tax %s no found", tax_number)
			return nil, utils.UserNotFound
		}
		log.Printf("Error getting user by tax from DB: %v", exist)
		return nil, utils.ErrorDataBase
	}

	return &tenant, nil
}


