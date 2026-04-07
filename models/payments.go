package models

import (
	"time"
	"vehiclesales/requests"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payment struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	EntityID    string             `bson:"entity_id" json:"entity_id"`
	PaymentCode string             `bson:"payment_code" json:"payment_code"`

	CompanyID string `bson:"company_id" json:"company_id"`
	BranchID  string `bson:"branch_id" json:"branch_id"`

	CustomerID    string `bson:"customer_id" json:"customer_id"`
	SalesOrderID  string `bson:"sales_order_id" json:"sales_order_id"`
	InvoiceNumber string `bson:"invoice_number" json:"invoice_number"`

	PaymentDate   time.Time `bson:"payment_date" json:"payment_date"`
	PaymentAmount float64   `bson:"payment_amount" json:"payment_amount"`

	PaymentMode string `bson:"payment_mode" json:"payment_mode"`
	PaymentType string `bson:"payment_type" json:"payment_type"`

	ReferenceNumber string `bson:"reference_number" json:"reference_number"`
	BankName        string `bson:"bank_name" json:"bank_name"`

	CollectedBy string `bson:"collected_by" json:"collected_by"`

	Remarks string `bson:"remarks" json:"remarks"`
	EmailStatus string `bson:"email_status" json:"email_status"` // Sent, Failed, Pending

	IsDeleted bool `bson:"is_deleted" json:"is_deleted"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type UpdatePayment struct {
	PaymentAmount    *float64   `bson:"payment_amount,omitempty" json:"payment_amount,omitempty"`
	PaymentMode      *string    `bson:"payment_mode,omitempty" json:"payment_mode,omitempty"`
	PaymentType      *string    `bson:"payment_type,omitempty" json:"payment_type,omitempty"`
	ReferenceNumber  *string    `bson:"reference_number,omitempty" json:"reference_number,omitempty"`
	BankName         *string    `bson:"bank_name,omitempty" json:"bank_name,omitempty"`
	CollectedBy      *string    `bson:"collected_by,omitempty" json:"collected_by,omitempty"`
	Remarks          *string    `bson:"remarks,omitempty" json:"remarks,omitempty"`
	PaymentDate      *time.Time `bson:"payment_date,omitempty" json:"payment_date,omitempty"`
}

func NewPayment() *Payment {
	now := time.Now().UTC()
	id := primitive.NewObjectID()

	return &Payment{
		ID:          id,
		EntityID:    id.Hex(),
		PaymentCode: "PAY" + id.Hex()[18:24],
		IsDeleted:   false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (p *Payment) Bind(req *requests.CreatePaymentRequest) {
	p.CompanyID = req.CompanyID
	p.BranchID = req.BranchID
	p.CustomerID = req.CustomerID
	p.SalesOrderID = req.SalesOrderID
	p.InvoiceNumber = req.InvoiceNumber

	p.PaymentDate = req.PaymentDate
	p.PaymentAmount = req.PaymentAmount

	p.PaymentMode = req.PaymentMode
	p.PaymentType = req.PaymentType

	p.ReferenceNumber = req.ReferenceNumber
	p.BankName = req.BankName

	p.CollectedBy = req.CollectedBy
	p.Remarks = req.Remarks
}