package requests

import (
	"time"

	"github.com/gin-gonic/gin"
)

type CreateSalesOrderRequest struct {
	CompanyID string `json:"company_id" binding:"required"`
	BranchID  string `json:"branch_id" binding:"required"`

	CustomerID         string `json:"customer_id" binding:"required"`
	VehicleInventoryID string `json:"vehicle_inventory_id" binding:"required"`
	SalespersonID      string `json:"salesperson_id" binding:"required"`

	SaleDate     time.Time `json:"sale_date" binding:"required"`
	DeliveryDate time.Time `json:"delivery_date" binding:"required"`

	VehiclePrice        float64 `json:"vehicle_price" binding:"required"`
	RegistrationCharges float64 `json:"registration_charges"`
	Insurance           float64 `json:"insurance"`
	Accessories         float64 `json:"accessories"`

	TotalAmount   float64 `json:"total_amount"`
	DownPayment   float64 `json:"down_payment"`
	LoanAmount    float64 `json:"loan_amount"`
	BalanceAmount float64 `json:"balance_amount"`
}

type UpdateSalesOrderRequest struct {
	DeliveryDate  *time.Time `json:"delivery_date,omitempty"`
	DownPayment   *float64   `json:"down_payment,omitempty"`
	LoanAmount    *float64   `json:"loan_amount,omitempty"`
	BalanceAmount *float64   `json:"balance_amount,omitempty"`
	Status        *string    `json:"status,omitempty"`
}

func NewCreateSalesOrderRequest() *CreateSalesOrderRequest {
	return &CreateSalesOrderRequest{}
}

func NewUpdateSalesOrderRequest() *UpdateSalesOrderRequest {
	return &UpdateSalesOrderRequest{}
}

func (r *CreateSalesOrderRequest) Validate(c *gin.Context) error {
	return c.ShouldBindJSON(r)
}

func (r *UpdateSalesOrderRequest) Validate(c *gin.Context) error {
	return c.ShouldBindJSON(r)
}