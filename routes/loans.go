package routes

import (
	"context"
	"net/http"
	"time"

	"vehiclesales/requests"
	"vehiclesales/services"

	"github.com/gin-gonic/gin"
)

func CreateLoan(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(400, gin.H{"error": "company_code is required"})
		return
	}

	req := requests.NewCreateLoanRequest()
	if err := req.Validate(c); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	service := services.NewLoanService()

	loan, err := service.Create(ctx, companyCode, req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, loan)
}

func GetLoans(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(400, gin.H{"error": "company_code is required"})
		return
	}

	service := services.NewLoanService()

	data, err := service.GetAll(ctx, companyCode)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

func GetLoanById(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(400, gin.H{"error": "company_code is required"})
		return
	}

	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "id is required"})
		return
	}

	service := services.NewLoanService()

	data, err := service.GetByID(ctx, companyCode, id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

func UpdateLoan(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(400, gin.H{"error": "company_code is required"})
		return
	}

	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "id is required"})
		return
	}

	req := requests.NewUpdateLoanRequest()
	if err := req.Validate(c); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	service := services.NewLoanService()

	loan, err := service.Update(ctx, companyCode, id, req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, loan)
}

func DeleteLoan(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(400, gin.H{"error": "company_code is required"})
		return
	}

	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "id is required"})
		return
	}

	service := services.NewLoanService()

	err := service.Delete(ctx, companyCode, id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
