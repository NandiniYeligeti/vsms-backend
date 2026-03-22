package routes

import (
	"context"
	"net/http"
	"os"
	"time"

	"vehiclesales/middleware"
	"vehiclesales/requests"
	"vehiclesales/services"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var req requests.LoginRequest
	if err := req.Validate(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service := services.NewAuthService()
	res, err := service.Login(ctx, &req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Detect environment (Render sets PORT or GIN_MODE)
	isProd := os.Getenv("PORT") != "" || os.Getenv("GIN_MODE") == "release"

	if isProd {
		// IMPORTANT: For cross-site cookies (Vercel -> Render), we need SameSite=None and Secure=true
		c.SetSameSite(http.SameSiteNoneMode)
		c.SetCookie("access_token", res.Token, 86400, "/", "", true, true)
	} else {
		c.SetCookie("access_token", res.Token, 86400, "/", "", false, true)
	}

	// Return only user info (not token)
	c.JSON(http.StatusOK, gin.H{"user": res.User})
}

func CreateCompany(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var req requests.CreateCompanyRequest
	if err := req.Validate(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service := services.NewAuthService()
	user, err := service.CreateCompany(ctx, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func GetCompaniesList(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	service := services.NewAuthService()
	users, err := service.GetCompanies(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// ================= USER MANAGEMENT (Admin creates users for their company) =================

func CreateUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var req requests.CreateUserRequest
	if err := req.Validate(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the admin's company info from JWT claims
	userClaims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	claims := userClaims.(*middleware.Claims)

	service := services.NewAuthService()
	user, err := service.CreateUser(ctx, &req, claims.CompanyCode, "")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func GetUsers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")

	service := services.NewAuthService()
	users, err := service.GetUsers(ctx, companyCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func DeleteUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userID := c.Param("id")

	service := services.NewAuthService()
	err := service.DeleteUser(ctx, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
