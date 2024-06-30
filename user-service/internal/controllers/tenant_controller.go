package controllers

import (
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

func NewTenantController() *TenantController  {
	return &TenantController{
		service: services.NewTenantService(),
	}
}

func (c *TenantController) Index(data *gin.Context)  {
	tenants := c.service.Index()
	data.JSON(http.StatusOK, tenants)
}

func (tc *TenantController) CreateTenant(c *gin.Context) {
	var tenant models.Tenant
	if err := c.ShouldBindJSON(&tenant); err != nil{
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

	c.JSON(http.StatusOK, gin.H{"message": "Tenant creadi exitosament", "user": tenant})
}
