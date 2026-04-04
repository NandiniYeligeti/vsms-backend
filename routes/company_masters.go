package routes

import (
	"context"
	"net/http"
	"vehiclesales/requests"
	"vehiclesales/services"
	"github.com/gin-gonic/gin"
)

var companyMasterService = services.NewCompanyMasterService()

func CreateCompanyMaster(c *gin.Context) {
	companyCode := c.Param("company_code")
	var req requests.CreateCompanyMasterRequest
	if err := req.Validate(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := companyMasterService.Create(context.Background(), companyCode, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

func GetCompanyMasters(c *gin.Context) {
	companyCode := c.Param("company_code")
	res, err := companyMasterService.GetAll(context.Background(), companyCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func DeleteCompanyMaster(c *gin.Context) {
	companyCode := c.Param("company_code")
	id := c.Param("id")
	err := companyMasterService.Delete(context.Background(), companyCode, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
