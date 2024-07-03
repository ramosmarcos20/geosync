package services

import (
    "errors"
    "geosync/user-service/internal/models"
    "geosync/user-service/internal/repositories"
    "geosync/user-service/internal/utils"
    "log"
)

type TenantService struct {
    repo repositories.TenantRepository
}

func NewTenantService(repo repositories.TenantRepository) *TenantService {
    return &TenantService{repo: repo}
}

func (s *TenantService) Index() []models.Tenant {
    return s.repo.Index()
}

func (s *TenantService) CreateTenant(tenant *models.Tenant) []error {
    var errs []error

    existsTenant, err := s.repo.ExistTenant(tenant.TaxNumber)
    if err != nil && !errors.Is(err, utils.UserNotFound) {
        log.Printf("Error checking tax number: %v", err)
        errs = append(errs, err)
    }

    if existsTenant != nil {
        errs = append(errs, utils.TaxInUse)
    }

    if len(errs) > 0 {
        return errs
    }

    err = s.repo.CreateTenant(tenant)
    if err != nil {
        log.Printf("Error creating tenant: %v", err)
        errs = append(errs, err)
    }

    return errs
}

func (s *TenantService) GetTenant(id string) (*models.Tenant, error) {
    tenant, err := s.repo.GetTenant(id)
    if err != nil {
        log.Printf("Error getting tenant: %v", err)
        if errors.Is(err, utils.UserNotFound) {
            return nil, utils.UserNotFound
        }
        return nil, utils.UserNotFound
    }
    return tenant, nil
}

