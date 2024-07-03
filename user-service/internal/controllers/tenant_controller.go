package controllers

import (
	"errors"
	"geosync/user-service/internal/models"
	"geosync/user-service/internal/services"
	"geosync/user-service/internal/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TenantController struct {
    service *services.TenantService
}

func NewTenantController(service *services.TenantService) *TenantController {
    return &TenantController{service: service}
}

func (tc *TenantController) Index(c *gin.Context) {
    tenants := tc.service.Index()
    c.JSON(http.StatusOK, tenants)
}

func (tc *TenantController) CreateTenant(c *gin.Context) {
    var tenant models.Tenant
    if err := c.ShouldBindJSON(&tenant); err != nil {
        log.Printf("Error binding JSON: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": utils.InvalidInput})
        return
    }

    errs := tc.service.CreateTenant(&tenant)
    if len(errs) > 0 {
        var errorMessages []string
        for _, e := range errs {
            errorMessages = append(errorMessages, e.Error())
        }
        c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
        log.Printf("Errors creating tenant: %v", errs)
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Tenant creado exitosamente", "tenant": tenant})
}

func (tc *TenantController) GetTenant(c *gin.Context)  {
    id          := c.Param("id")
    tenant, err := tc.service.GetTenant(id)

    if err != nil {
        log.Printf("Error getting tenant: %v", err)
        if errors.Is(err, utils.UserNotFound) {
            c.JSON(http.StatusNotFound, gin.H{"error": utils.UserNotFound.Error()})
        }else{
            c.JSON(http.StatusInternalServerError, gin.H{"error": utils.ErrorDataBase.Error()})
        }
        return
    }
    c.JSON(http.StatusOK, tenant)
}