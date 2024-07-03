package repositories

import "geosync/user-service/internal/models"

type UserRepository interface {
    CreateUser(user *models.User) error
    GetUsers() []models.User
    GetUser(id string) (*models.User, error)
    GetUserByEmail(email string) (*models.User, error)
    GetUserByUserName(userName string) (*models.User, error)
    UpdateUser(user *models.User) error
    DeleteUser(id uint) error
    GetUserAuth(id uint) (*models.User, error)
}

type TenantRepository interface {
    Index() []models.Tenant
    CreateTenant(tenant *models.Tenant) error
    ExistTenant(taxNumber string) (*models.Tenant, error)
    GetTenant(id string) (*models.Tenant, error)
}
// Aquí puedes agregar más interfaces para otros repositorios si es necesario
