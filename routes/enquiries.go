package routes

import (
	"context"
	"time"

	"vehiclesales/requests"
	"vehiclesales/services"

	"github.com/gin-gonic/gin"
)

func CreateEnquiry(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(400, gin.H{"error": "company_code is required"})
		return
	}

	var req requests.CreateEnquiryRequest
	if err := req.Validate(c); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	service := services.NewEnquiryService()
	enquiry, err := service.Create(ctx, companyCode, &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, enquiry)
}

func GetEnquiries(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(400, gin.H{"error": "company_code is required"})
		return
	}

	service := services.NewEnquiryService()
	enquiries, err := service.GetAll(ctx, companyCode)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, enquiries)
}

func GetEnquiryById(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	id := c.Param("id")
	if companyCode == "" || id == "" {
		c.JSON(400, gin.H{"error": "company_code and id are required"})
		return
	}

	service := services.NewEnquiryService()
	enquiry, err := service.GetByID(ctx, companyCode, id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, enquiry)
}

func UpdateEnquiry(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	id := c.Param("id")
	if companyCode == "" || id == "" {
		c.JSON(400, gin.H{"error": "company_code and id are required"})
		return
	}

	var req requests.UpdateEnquiryRequest
	if err := req.Validate(c); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	service := services.NewEnquiryService()
	enquiry, err := service.Update(ctx, companyCode, id, &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, enquiry)
}

func DeleteEnquiry(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	id := c.Param("id")
	if companyCode == "" || id == "" {
		c.JSON(400, gin.H{"error": "company_code and id are required"})
		return
	}

	service := services.NewEnquiryService()
	err := service.Delete(ctx, companyCode, id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Enquiry deleted successfully"})
}
