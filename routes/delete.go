package routes

import (
	"context"
	"net/http"
	"time"

	// "shared/middleware"
	"vehiclesales/services"

	"github.com/gin-gonic/gin"
)

func DeleteCustomer(c *gin.Context) {
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

	// Get params
	companyCode := c.Param("company_code")
	id := c.Param("id")

	// Validate required params
	if companyCode == "" || id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "company_code and id are required",
		})
		return
	}

	// Initialize service
	service := services.NewCustomerService()

	// Delete customer
	err := service.Delete(ctx, companyCode, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "failed to delete customer",
			"details": err.Error(),
		})
		return
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Customer deleted successfully",
	})
}

func DeleteSalesperson(c *gin.Context) {
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

	service := services.NewSalespersonService()

	err := service.Delete(ctx, companyCode, id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Salesperson deleted successfully"})
}

func DeleteVehicleModel(c *gin.Context) {
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

	service := services.NewVehicleModelService()

	err := service.Delete(ctx, companyCode, id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Vehicle model deleted successfully"})
}

func DeleteVehicleInventory(c *gin.Context) {
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

	service := services.NewVehicleInventoryService()

	err := service.Delete(ctx, companyCode, id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Vehicle inventory deleted successfully"})
}

func DeleteSalesOrder(c *gin.Context) {
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

	service := services.NewSalesOrderService()

	err := service.Delete(ctx, companyCode, id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Sales order deleted successfully"})
}

func DeletePayment(c *gin.Context) {
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

	service := services.NewPaymentService()

	err := service.Delete(ctx, companyCode, id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Payment deleted successfully"})
}
