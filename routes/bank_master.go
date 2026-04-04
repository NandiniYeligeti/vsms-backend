package routes

import (
	"context"
	"net/http"
	"time"

	"vehiclesales/requests"
	"vehiclesales/services"

	"github.com/gin-gonic/gin"
)

func CreateBankMaster(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "company_code is required"})
		return
	}

	req := requests.NewCreateBankMasterRequest()
	if err := req.Validate(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.CompanyID = companyCode

	service := services.NewBankMasterService()
	bank, err := service.Create(ctx, companyCode, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, bank)
}

func GetBankMasters(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "company_code is required"})
		return
	}

	service := services.NewBankMasterService()
	banks, err := service.GetAll(ctx, companyCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, banks)
}

func UpdateBankMaster(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	id := c.Param("id")
	if companyCode == "" || id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "company_code and id are required"})
		return
	}

	req := requests.NewUpdateBankMasterRequest()
	if err := req.Validate(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service := services.NewBankMasterService()
	bank, err := service.Update(ctx, companyCode, id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bank)
}

func DeleteBankMaster(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	id := c.Param("id")
	if companyCode == "" || id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "company_code and id are required"})
		return
	}

	service := services.NewBankMasterService()
	err := service.Delete(ctx, companyCode, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bank master record deleted"})
}
