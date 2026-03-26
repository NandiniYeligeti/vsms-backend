package routes

import (
	"context"
	"net/http"
	"time"

	// "shared/middleware"
	"vehiclesales/requests"
	"vehiclesales/services"

	"github.com/gin-gonic/gin"
)

func UpdateCustomer(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Validate access token
	// _, err := middleware.GetAccessClaims(c)
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
	// 		"error": "unauthorized access",
	// 	})
	// 	return
	// }

	// Get required params
	companyCode := c.Param("company_code")
	id := c.Param("id")

	if companyCode == "" || id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "company_code and id are required",
		})
		return
	}

	// Validate request body
	req := requests.NewUpdateCustomerRequest()
	if err := req.Validate(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Handle Photo upload if present
	photoFile, err := c.FormFile("photo")
	if err == nil {
		if err := c.SaveUploadedFile(photoFile, "uploads/"+photoFile.Filename); err == nil {
			req.Photo = photoFile
		}
	}

	// Handle multiple Documents upload
	form, _ := c.MultipartForm()
	files := form.File["documents"]
	for _, file := range files {
		if err := c.SaveUploadedFile(file, "uploads/"+file.Filename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save document " + file.Filename})
			return
		}
		req.Documents = append(req.Documents, file)
	}

	// Initialize service
	service := services.NewCustomerService()

	// Update customer
	result, err := service.Update(ctx, companyCode, id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "failed to update customer",
			"details": err.Error(),
		})
		return
	}

	// Success response
	c.JSON(http.StatusOK, result)
}

func UpdateSalesperson(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	id := c.Param("id")

	if companyCode == "" || id == "" {
		c.JSON(400, gin.H{"error": "company_code and id are required"})
		return
	}

	req := requests.NewUpdateSalespersonRequest()
	if err := req.Validate(c); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	service := services.NewSalespersonService()

	result, err := service.Update(ctx, companyCode, id, req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}

func UpdateVehicleModel(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	id := c.Param("id")

	if companyCode == "" || id == "" {
		c.JSON(400, gin.H{"error": "company_code and id are required"})
		return
	}

	req := requests.NewUpdateVehicleModelRequest()
	if err := req.Validate(c); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	service := services.NewVehicleModelService()

	result, err := service.Update(ctx, companyCode, id, req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}

func UpdateVehicleInventory(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	id := c.Param("id")

	if companyCode == "" || id == "" {
		c.JSON(400, gin.H{"error": "company_code and id are required"})
		return
	}

	req := requests.NewUpdateVehicleInventoryRequest()
	if err := req.Validate(c); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	service := services.NewVehicleInventoryService()

	result, err := service.Update(ctx, companyCode, id, req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}


func UpdateSalesOrder(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	id := c.Param("id")

	if companyCode == "" || id == "" {
		c.JSON(400, gin.H{"error": "company_code and id are required"})
		return
	}

	req := requests.NewUpdateSalesOrderRequest()
	if err := req.Validate(c); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	service := services.NewSalesOrderService()

	result, err := service.Update(ctx, companyCode, id, req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}

func UpdatePayment(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	id := c.Param("id")

	if companyCode == "" || id == "" {
		c.JSON(400, gin.H{"error": "company_code and id are required"})
		return
	}

	req := requests.NewUpdatePaymentRequest()
	if err := req.Validate(c); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	service := services.NewPaymentService()

	result, err := service.Update(ctx, companyCode, id, req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}


