package routes

import (
	"context"
	"net/http"
	"time"

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

	// Set JWT as httpOnly, Secure cookie (adjust domain as needed)
	c.SetCookie("access_token", res.Token, 3600, "/", "localhost", false, true) // Set Secure=true in production

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
