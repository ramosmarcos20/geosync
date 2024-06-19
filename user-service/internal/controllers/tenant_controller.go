package controllers

import (
	"geosync/user-service/internal/services"
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
