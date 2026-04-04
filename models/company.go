package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompanySettings struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CompanyID     string             `bson:"company_id" json:"company_id"`
	CompanyName   string             `bson:"company_name" json:"company_name"`
	LogoURL       string             `bson:"logo_url" json:"logo_url"`
	GSTNumber     string             `bson:"gst_number" json:"gst_number"`
	Address       string             `bson:"address" json:"address"`
	Phone         string             `bson:"phone" json:"phone"`
	Email         string             `bson:"email" json:"email"`
	InvoicePrefix string             `bson:"invoice_prefix" json:"invoice_prefix"`
	InvoiceSuffix string             `bson:"invoice_suffix" json:"invoice_suffix"`
	SalesPrefix   string             `bson:"sales_prefix" json:"sales_prefix"`
	SalesSuffix   string             `bson:"sales_suffix" json:"sales_suffix"`
	Currency      string             `bson:"currency" json:"currency"`
	Timezone      string             `bson:"timezone" json:"timezone"`

	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

func NewCompanySettings(companyCode string) *CompanySettings {
	return &CompanySettings{
		ID:            primitive.NewObjectID(),
		CompanyID:     companyCode,
		CompanyName:   "AutoDesk Motors",
		LogoURL:       "",
		Currency:      "INR",
		Timezone:      "Asia/Kolkata",
		UpdatedAt:     time.Now(),
	}
}
