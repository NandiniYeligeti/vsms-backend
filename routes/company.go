package routes

import (
	"context"
	"net/http"
	"vehiclesales/middleware"
	"vehiclesales/models"
	"vehiclesales/services"

	"github.com/gin-gonic/gin"
)

func GetCompanySettings(c *gin.Context) {
	companyCode := c.Param("company_code")
	service := services.NewCompanySettingsService()

	settings, err := service.Get(context.Background(), companyCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, settings)
}

func UpdateCompanySettings(c *gin.Context) {
	companyCode := c.Param("company_code")
	
	// Role check: Only admin or super_admin can update settings
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	claims := user.(*middleware.Claims)
	if claims.Role != "admin" && claims.Role != "super_admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: Only admins can update settings"})
		return
	}

	var settings models.CompanySettings
	if err := c.ShouldBindJSON(&settings); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service := services.NewCompanySettingsService()
	err := service.Update(context.Background(), companyCode, &settings)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Settings updated successfully", "data": settings})
}
