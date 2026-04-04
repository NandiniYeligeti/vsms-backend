package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Loan struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	EntityID        string            `bson:"entity_id" json:"entity_id"`
	LoanCode        string            `bson:"loan_code" json:"loan_code"`
	CompanyID       string            `bson:"company_id" json:"company_id"`
	BranchID        string            `bson:"branch_id" json:"branch_id"`

	CustomerID      string            `bson:"customer_id" json:"customer_id"`
	CustomerName    string            `bson:"customer_name" json:"customer_name"`
	SalesOrderID    string            `bson:"sales_order_id" json:"sales_order_id"`
	SalesOrderCode  string            `bson:"sales_order_code" json:"sales_order_code"`

	BankName        string            `bson:"bank_name" json:"bank_name"`
	LoanAmount      float64           `bson:"loan_amount" json:"loan_amount"`
	InterestRate    float64           `bson:"interest_rate" json:"interest_rate"`
	DurationMonths  int               `bson:"duration_months" json:"duration_months"`
	EMIAmount       float64           `bson:"emi_amount" json:"emi_amount"`
	Status          string            `bson:"status" json:"status"` // Applied, Approved, Disbursed, Rejected
	AccountNumber   string            `bson:"account_number" json:"account_number"`
	BankPerson      string            `bson:"bank_person" json:"bank_person"`
	Mobile          string            `bson:"mobile" json:"mobile"`
	DisbursementDate *time.Time        `bson:"disbursement_date" json:"disbursement_date"`
	
	IsDeleted       bool              `bson:"is_deleted" json:"is_deleted"`
	CreatedAt       time.Time         `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time         `bson:"updated_at" json:"updated_at"`
}

func NewLoan() *Loan {
	now := time.Now().UTC()
	id := primitive.NewObjectID()

	return &Loan{
		ID:             id,
		EntityID:       id.Hex(),
		LoanCode:       "LN" + id.Hex()[18:24],
		Status:         "Applied",
		IsDeleted:      false,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
}
