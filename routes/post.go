package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"vehiclesales/requests"
	"vehiclesales/services"
	// "shared/middleware"

	"github.com/gin-gonic/gin"
)
func CreateCustomer(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// _, err := middleware.GetAccessClaims(c)
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	// company code from URL
	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "company_code is required",
		})
		return
	}

	// bind request
	req := requests.NewCreateCustomerRequest()
	if err := req.Validate(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Printf("DEBUG: Photo received via struct: %v\n", req.Photo != nil)
	
	photoFile, err := c.FormFile("photo")
	if err == nil {
		fmt.Printf("DEBUG: Photo received via FormFile: %s\n", photoFile.Filename)
		if req.Photo == nil {
			req.Photo = photoFile
		}
	} else {
		fmt.Printf("DEBUG: FormFile error: %v\n", err)
	}

	// save uploaded file if exists
	if req.Photo != nil {
		fmt.Printf("DEBUG: Saving photo to: %s\n", "uploads/"+req.Photo.Filename)
		if err := c.SaveUploadedFile(req.Photo, "uploads/"+req.Photo.Filename); err != nil {
			fmt.Printf("DEBUG: Save Error: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to save photo: " + err.Error(),
			})
			return
		}
	}

	// service call
	service := services.NewCustomerService()

	customer, err := service.Create(ctx, companyCode, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, customer)
}

func CreateSalesperson(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// company code from URL
	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "company_code is required",
		})
		return
	}

	// bind request
	req := requests.NewCreateSalespersonRequest()
	if err := req.Validate(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// service call
	service := services.NewSalespersonService()

	salesperson, err := service.Create(ctx, companyCode, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, salesperson)
}

func CreateVehicleModel(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(400, gin.H{"error": "company_code is required"})
		return
	}

	req := requests.NewCreateVehicleModelRequest()
	if err := req.Validate(c); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	service := services.NewVehicleModelService()

	vehicle, err := service.Create(ctx, companyCode, req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, vehicle)
}

func CreateVehicleInventory(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(400, gin.H{"error": "company_code is required"})
		return
	}

	req := requests.NewCreateVehicleInventoryRequest()
	if err := req.Validate(c); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	service := services.NewVehicleInventoryService()

	vehicle, err := service.Create(ctx, companyCode, req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, vehicle)
}

func CreateSalesOrder(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(400, gin.H{"error": "company_code is required"})
		return
	}

	req := requests.NewCreateSalesOrderRequest()
	if err := req.Validate(c); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	service := services.NewSalesOrderService()

	order, err := service.Create(ctx, companyCode, req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, order)
}

func CreatePayment(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(400, gin.H{"error": "company_code is required"})
		return
	}

	req := requests.NewCreatePaymentRequest()
	if err := req.Validate(c); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	service := services.NewPaymentService()

	payment, err := service.Create(ctx, companyCode, req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, payment)
}

func SeedData(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	if companyCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "company_code is required"})
		return
	}

	err := services.SeedDatabase(ctx, companyCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Database successfully seeded with mock data!"})
}