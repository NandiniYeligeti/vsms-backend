package requests

import (
	"time"

	"github.com/gin-gonic/gin"
)

type CreateLoanRequest struct {
	CompanyID      string  `json:"company_id" binding:"required"`
	BranchID       string  `json:"branch_id" binding:"required"`
	CustomerID     string  `json:"customer_id" binding:"required"`
	SalesOrderID   string  `json:"sales_order_id" binding:"required"`
	BankName       string  `json:"bank_name" binding:"required"`
	LoanAmount     float64 `json:"loan_amount" binding:"required"`
	InterestRate   float64 `json:"interest_rate" binding:"required"`
	DurationMonths int     `json:"duration_months" binding:"required"`
	EMIAmount      float64 `json:"emi_amount" binding:"required"`
	Status         string  `json:"status" binding:"omitempty"`
	AccountNumber  string  `json:"account_number" binding:"omitempty"`
	BankPerson     string  `json:"bank_person" binding:"omitempty"`
	Mobile         string  `json:"mobile" binding:"omitempty"`
}

type UpdateLoanRequest struct {
	BankName        *string    `json:"bank_name,omitempty"`
	LoanAmount      *float64   `json:"loan_amount,omitempty"`
	InterestRate    *float64   `json:"interest_rate,omitempty"`
	DurationMonths  *int       `json:"duration_months,omitempty"`
	EMIAmount       *float64   `json:"emi_amount,omitempty"`
	Status          *string    `json:"status,omitempty"`
	AccountNumber   *string    `json:"account_number,omitempty"`
	BankPerson      *string    `json:"bank_person,omitempty"`
	Mobile          *string    `json:"mobile,omitempty"`
	DisbursementDate *time.Time `json:"disbursement_date,omitempty"`
}

func NewCreateLoanRequest() *CreateLoanRequest {
	return &CreateLoanRequest{}
}

func NewUpdateLoanRequest() *UpdateLoanRequest {
	return &UpdateLoanRequest{}
}

func (r *CreateLoanRequest) Validate(c *gin.Context) error {
	return c.ShouldBindJSON(r)
}

func (r *UpdateLoanRequest) Validate(c *gin.Context) error {
	return c.ShouldBindJSON(r)
}
