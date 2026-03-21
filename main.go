package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"vehiclesales/middleware"
	"vehiclesales/routes"
	"vehiclesales/storage"

	// "shared/pkgs/jwtmanager"

	// "shared/pkgs/keys"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Connect to MongoDB

	// if err := keys.InitKeyPair(); err != nil {
	// 	log.Fatalf("failed to initiate key pair %v", err)
	// }

	if err := storage.InitMongo(); err != nil {
		log.Fatalf("Mongo connection failed: %v", err)
	}
	fmt.Println("MongoDB connected")

	// // Initialize Vault-backed JWT signer and cache public key
	// if err := jwtmanager.InitVaultJWT(); err != nil {
	// 	log.Fatalf("JWT initialization failed: %v", err)
	// }
	// fmt.Println("✅ JWT signer initialized")

	// Create main app router
	app := gin.Default()

	// Enable CORS globally
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000", "http://localhost:4173", "http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type, Accept, Origin, X-Requested-With, X-CSRF-Token, X-Company-Code, Authorization, X-Forwarded-Proto"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	// Logging middleware
	// app.Use(middleware.LogMiddleware())
	app.Use(gin.Recovery())

	// API group
	api := app.Group("/api/v1")

	// Static files
	app.Static("/uploads", "./uploads")

	// Public routes (no auth)
	routes.Routes(api)

	// Protected routes (add middleware)
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware())
	// Example: add a protected test route
	protected.GET("/protected-test/:company_code", func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		claims := user.(*middleware.Claims)
		requestedCompany := c.Param("company_code")
		if claims.CompanyCode != requestedCompany {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: company mismatch"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Access granted", "user": claims})
	})

	// Start server
	server := &http.Server{
		Addr:    ":4001",
		Handler: app,
	}

	fmt.Println("🚀 Server running on :4001")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}
}
