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
		salesOrder.POST("/:company_code/:id/send-email", SendOrderEmail)
		salesOrder.GET("/:company_code/:id/preview-email", PreviewOrderEmail)
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
		payment.POST("/:company_code/:id/send-email", SendPaymentEmail)
		payment.GET("/:company_code/:id/preview-email", PreviewPaymentEmail)
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

	bankMaster := api.Group("/bank-master")
	{
		bankMaster.POST("/:company_code", CreateBankMaster)
		bankMaster.GET("/:company_code", GetBankMasters)
		bankMaster.PUT("/:company_code/:id", UpdateBankMaster)
		bankMaster.DELETE("/:company_code/:id", DeleteBankMaster)
	}

	// ================= USER MANAGEMENT =================
	users := api.Group("/users")
	users.Use(middleware.AuthMiddleware())
	{
		users.POST("/:company_code", CreateUser)
		users.GET("/:company_code", GetUsers)
		users.PUT("/:id/menus", UpdateUserMenus)
		users.PUT("/:id/password", UpdatePassword)
		users.DELETE("/:company_code/:id", DeleteUser)
	}

	// ================= VEHICLE FEATURES =================
	vehicleFeatures := api.Group("/vehicle-features")
	vehicleFeatures.Use(middleware.AuthMiddleware())
	{
		vehicleFeatures.POST("/type/:company_code", CreateVehicleType)
		vehicleFeatures.GET("/type/:company_code", GetVehicleTypes)
		vehicleFeatures.DELETE("/type/:company_code/:id", DeleteVehicleType)

		vehicleFeatures.POST("/category/:company_code", CreateVehicleCategory)
		vehicleFeatures.GET("/category/:company_code", GetVehicleCategories)
		vehicleFeatures.DELETE("/category/:company_code/:id", DeleteVehicleCategory)

		vehicleFeatures.POST("/accessory/:company_code", CreateVehicleAccessory)
		vehicleFeatures.GET("/accessory/:company_code", GetVehicleAccessories)
		vehicleFeatures.DELETE("/accessory/:company_code/:id", DeleteVehicleAccessory)
	}

	// ================= SETTINGS =================
	settings := api.Group("/settings")
	settings.Use(middleware.AuthMiddleware())
	{
		settings.GET("/:company_code", GetCompanySettings)
		settings.PUT("/:company_code", UpdateCompanySettings)
		settings.POST("/:company_code/test-email", SendTestEmail)
	}

	// ================= COMPANY MASTERS =================
	companyMasters := api.Group("/company-master")
	companyMasters.Use(middleware.AuthMiddleware())
	{
		companyMasters.POST("/:company_code", CreateCompanyMaster)
		companyMasters.GET("/:company_code", GetCompanyMasters)
		companyMasters.DELETE("/:company_code/:id", DeleteCompanyMaster)
	}

}
