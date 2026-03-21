package routes

import (
	"vehiclesales/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(api *gin.RouterGroup) {

	// ================= AUTH/SUPER-ADMIN =================
	auth := api.Group("")
	{
		auth.POST("/login", Login)
		auth.POST("/super-admin/company", CreateCompany)
		auth.GET("/super-admin/companies", GetCompaniesList)
	}

	// ================= CUSTOMERS =================
	customer := api.Group("/customer")
	customer.Use(middleware.AuthMiddleware())
	{
		customer.POST("/:company_code", func(c *gin.Context) {
			user, exists := c.Get("user")
			if !exists {
				c.JSON(401, gin.H{"error": "Unauthorized"})
				return
			}
			claims := user.(*middleware.Claims)
			requestedCompany := c.Param("company_code")
			if claims.CompanyCode != requestedCompany {
				c.JSON(403, gin.H{"error": "Forbidden: company mismatch"})
				return
			}
			CreateCustomer(c)
		})
		customer.GET("/:company_code", GetCustomers)
		customer.GET("/:company_code/:id", GetCustomerById)
		customer.PUT("/:company_code/:id", UpdateCustomer)
		customer.DELETE("/:company_code/:id", DeleteCustomer)
		customer.GET("/ledger/:company_code/:id", GetCustomerLedger)
	}

	// ================= SALESPERSONS =================
	salesperson := api.Group("/salesperson")
	salesperson.Use(middleware.AuthMiddleware())
	{
		salesperson.POST("/:company_code", CreateSalesperson)
		salesperson.GET("/:company_code", GetSalespersons)
		salesperson.GET("/:company_code/:id", GetSalespersonById)
		salesperson.PUT("/:company_code/:id", UpdateSalesperson)
		salesperson.DELETE("/:company_code/:id", DeleteSalesperson)
	}

	// ================= VEHICLE MODELS =================
	vehicleModel := api.Group("/vehicle-model")
	vehicleModel.Use(middleware.AuthMiddleware())
	{
		vehicleModel.POST("/:company_code", CreateVehicleModel)
		vehicleModel.GET("/:company_code", GetVehicleModels)
		vehicleModel.GET("/:company_code/:id", GetVehicleModelById)
		vehicleModel.PUT("/:company_code/:id", UpdateVehicleModel)
		vehicleModel.DELETE("/:company_code/:id", DeleteVehicleModel)
	}

	// ================= VEHICLE INVENTORY =================
	vehicleInventory := api.Group("/vehicle-inventory")
	vehicleInventory.Use(middleware.AuthMiddleware())
	{
		vehicleInventory.POST("/:company_code", CreateVehicleInventory)
		vehicleInventory.GET("/:company_code", GetVehicleInventories)
		vehicleInventory.GET("/:company_code/:id", GetVehicleInventoryById)
		vehicleInventory.PUT("/:company_code/:id", UpdateVehicleInventory)
		vehicleInventory.DELETE("/:company_code/:id", DeleteVehicleInventory)
	}

	// ================= SALES ORDERS =================
	salesOrder := api.Group("/sales-order")
	salesOrder.Use(middleware.AuthMiddleware())
	{
		salesOrder.POST("/:company_code", CreateSalesOrder)
		salesOrder.GET("/:company_code", GetSalesOrders)
		salesOrder.GET("/:company_code/:id", GetSalesOrderById)
		salesOrder.PUT("/:company_code/:id", UpdateSalesOrder)
		salesOrder.DELETE("/:company_code/:id", DeleteSalesOrder)
	}

	// ================= PAYMENTS =================
	payment := api.Group("/payment")
	payment.Use(middleware.AuthMiddleware())
	{
		payment.POST("/:company_code", CreatePayment)
		payment.GET("/:company_code", GetPayments)
		payment.GET("/:company_code/:id", GetPaymentById)
		payment.PUT("/:company_code/:id", UpdatePayment)
		payment.DELETE("/:company_code/:id", DeletePayment)
	}

	// ================= LOANS =================
	loan := api.Group("/loan")
	loan.Use(middleware.AuthMiddleware())
	{
		loan.POST("/:company_code", CreateLoan)
		loan.GET("/:company_code", GetLoans)
		loan.GET("/:company_code/:id", GetLoanById)
		loan.PUT("/:company_code/:id", UpdateLoan)
		loan.DELETE("/:company_code/:id", DeleteLoan)
	}

	// ================= DASHBOARD =================
	dashboard := api.Group("/dashboard")
	dashboard.Use(middleware.AuthMiddleware())
	{
		dashboard.GET("/:company_code", GetDashboardStats)
		dashboard.POST("/seed/:company_code", SeedData)
	}

}
