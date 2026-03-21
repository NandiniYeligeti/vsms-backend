package routes

import (
	"context"
	"time"

	// "shared/middleware"
	"vehiclesales/services"

	"github.com/gin-gonic/gin"
)

func GetCustomers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// _, err := middleware.GetAccessClaims(c)
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(400, gin.H{"error": "company_code is required"})
		return
	}

	service := services.NewCustomerService()

	data, err := service.GetAll(ctx, companyCode)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

func GetCustomerById(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// _, err := middleware.GetAccessClaims(c)
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

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

	service := services.NewCustomerService()

	data, err := service.GetByID(ctx, companyCode, id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

func GetSalespersons(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(400, gin.H{"error": "company_code is required"})
		return
	}

	service := services.NewSalespersonService()

	data, err := service.GetAll(ctx, companyCode)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

func GetSalespersonById(c *gin.Context) {
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

	data, err := service.GetByID(ctx, companyCode, id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

func GetVehicleModels(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(400, gin.H{"error": "company_code is required"})
		return
	}

	service := services.NewVehicleModelService()

	data, err := service.GetAll(ctx, companyCode)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

func GetVehicleModelById(c *gin.Context) {
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

	data, err := service.GetByID(ctx, companyCode, id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

func GetVehicleInventories(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(400, gin.H{"error": "company_code is required"})
		return
	}

	service := services.NewVehicleInventoryService()

	data, err := service.GetAll(ctx, companyCode)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

func GetVehicleInventoryById(c *gin.Context) {
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

	data, err := service.GetByID(ctx, companyCode, id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

func GetSalesOrders(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(400, gin.H{"error": "company_code is required"})
		return
	}

	service := services.NewSalesOrderService()

	data, err := service.GetAll(ctx, companyCode)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

func GetSalesOrderById(c *gin.Context) {
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

	data, err := service.GetByID(ctx, companyCode, id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

func GetPayments(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(400, gin.H{"error": "company_code is required"})
		return
	}

	service := services.NewPaymentService()

	data, err := service.GetAll(ctx, companyCode)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

func GetPaymentById(c *gin.Context) {
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

	data, err := service.GetByID(ctx, companyCode, id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

func GetDashboardStats(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(400, gin.H{"error": "company_code is required"})
		return
	}

	service := services.NewDashboardService()

	data, err := service.GetStats(ctx, companyCode)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

func GetCustomerLedger(c *gin.Context) {
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

	service := services.NewCustomerService()

	data, err := service.GetLedger(ctx, companyCode, id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}
