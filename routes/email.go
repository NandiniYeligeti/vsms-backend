package routes

import (
	"context"
	"net/http"
	"time"
	"vehiclesales/services"

	"github.com/gin-gonic/gin"
)

func SendOrderEmail(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	id := c.Param("id")

	service := services.NewSalesOrderService()
	err := service.SendEmail(ctx, companyCode, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order email sent successfully"})
}

func SendPaymentEmail(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	id := c.Param("id")

	service := services.NewPaymentService()
	err := service.SendEmail(ctx, companyCode, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment receipt email sent successfully"})
}

func PreviewOrderEmail(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	id := c.Param("id")

	service := services.NewSalesOrderService()
	preview, err := service.PreviewEmail(ctx, companyCode, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, preview)
}

func PreviewPaymentEmail(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	companyCode := c.Param("company_code")
	id := c.Param("id")

	service := services.NewPaymentService()
	preview, err := service.PreviewEmail(ctx, companyCode, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, preview)
}

