package services

import (
	"geosync/user-service/internal/models"
	"geosync/user-service/internal/repositories"
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