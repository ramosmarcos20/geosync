package repositories

import (
    "geosync/user-service/internal/models"
    "geosync/user-service/internal/utils"
    "log"

    "gorm.io/gorm"
)

type tenantRepository struct {
    db *gorm.DB
}

func NewTenantRepository(db *gorm.DB) TenantRepository {
    return &tenantRepository{db: db}
}

func (r *tenantRepository) Index() []models.Tenant {
    var tenants []models.Tenant
    err := r.db.Find(&tenants).Error
    if err != nil {
        log.Printf("Error en obtener tenant: %v", err)
    }
    return tenants
}

func (r *tenantRepository) CreateTenant(tenant *models.Tenant) error {
    err := r.db.Create(tenant).Error
    if err != nil {
        log.Printf("Error creating tenant in DB: %v", err)
    }
    return err
}

func (r *tenantRepository) GetTenant(id string) (*models.Tenant, error) {
    var tenant models.Tenant
    err := r.db.First(&tenant, id).Error
    if err != nil {
        log.Printf("Error getting tenant from DB: %v", err)
        return nil, err
    }
    return &tenant, nil
}


/* VALIDACIONES */
func (r *tenantRepository) ExistTenant(taxNumber string) (*models.Tenant, error) {
    var tenant models.Tenant
    err := r.db.Where("tax_number = ?", taxNumber).First(&tenant).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            log.Printf("Tenant with tax number %s not found", taxNumber)
            return nil, utils.UserNotFound
        }
        log.Printf("Error getting tenant by tax number from DB: %v", err)
        return nil, utils.ErrorDataBase
    }
    return &tenant, nil
}


