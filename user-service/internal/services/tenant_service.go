package services

import (
	"errors"
	"geosync/user-service/internal/models"
	"geosync/user-service/internal/repositories"
	"geosync/user-service/internal/utils"
	"log"
)

type TenantService struct{
	repo *repositories.TenantRepository
}

func NewTenantService() *TenantService {
	return &TenantService{
		repo: repositories.NewTenantRepository(),
	}
}

func (s *TenantService) Index() []models.Tenant  {
	tenants := s.repo.Index()
	return tenants
}

func (s *TenantService) CreateTenant(tenant *models.Tenant) []error  {
	var errs []error

	existeTenant, err := s.repo.ExistTenant(tenant.TaxNumber);

	if err != nil && !errors.Is(err, utils.UserNotFound) {
		log.Printf("Error checking tax number: %v", err)
		errs = append(errs, err)
	}

	if existeTenant != nil {
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