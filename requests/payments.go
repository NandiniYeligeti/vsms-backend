package requests

import (
	"time"

	"github.com/gin-gonic/gin"
)

type CreatePaymentRequest struct {
	CompanyID string `json:"company_id" binding:"required"`
	BranchID  string `json:"branch_id" binding:"required"`

	CustomerID    string `json:"customer_id" binding:"required"`
	SalesOrderID  string `json:"sales_order_id" binding:"required"`
	InvoiceNumber string `json:"invoice_number" binding:"required"`

	PaymentDate   time.Time `json:"payment_date" binding:"required"`
	PaymentAmount float64   `json:"payment_amount" binding:"required"`

	PaymentMode string `json:"payment_mode" binding:"required"`
	PaymentType string `json:"payment_type" binding:"required"`

	ReferenceNumber string `json:"reference_number"`
	BankName        string `json:"bank_name"`

	CollectedBy string `json:"collected_by" binding:"required"`

	Remarks string `json:"remarks"`
}

type UpdatePaymentRequest struct {
	PaymentDate      *time.Time `json:"payment_date,omitempty"`
	PaymentAmount    *float64   `json:"payment_amount,omitempty"`
	PaymentMode      *string    `json:"payment_mode,omitempty"`
	PaymentType      *string    `json:"payment_type,omitempty"`
	ReferenceNumber  *string    `json:"reference_number,omitempty"`
	BankName         *string    `json:"bank_name,omitempty"`
	CollectedBy      *string    `json:"collected_by,omitempty"`
	Remarks          *string    `json:"remarks,omitempty"`
}

func NewCreatePaymentRequest() *CreatePaymentRequest {
	return &CreatePaymentRequest{}
}

func NewUpdatePaymentRequest() *UpdatePaymentRequest {
	return &UpdatePaymentRequest{}
}

func (r *CreatePaymentRequest) Validate(c *gin.Context) error {
	return c.ShouldBindJSON(r)
}

func (r *UpdatePaymentRequest) Validate(c *gin.Context) error {
	return c.ShouldBindJSON(r)
}