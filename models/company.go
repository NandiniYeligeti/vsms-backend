package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmailSettings struct {
	SenderName       string `bson:"sender_name" json:"sender_name"`
	SenderEmail      string `bson:"sender_email" json:"sender_email"`
	SMTPHost         string `bson:"smtp_host" json:"smtp_host"`
	SMTPPort         int    `bson:"smtp_port" json:"smtp_port"`
	EncryptionType   string `bson:"encryption_type" json:"encryption_type"`
	EmailUsername    string `bson:"email_username" json:"email_username"`
	EmailPassword    string `bson:"email_password" json:"email_password"`
	EnableEmail      bool   `bson:"enable_email" json:"enable_email"`
	AutoSendReceipt  bool   `bson:"auto_send_receipt" json:"auto_send_receipt"`
	AttachInvoice    bool   `bson:"attach_invoice" json:"attach_invoice"`
}

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

	EmailSettings EmailSettings `bson:"email_settings" json:"email_settings"`

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
		EmailSettings: EmailSettings{
			SMTPPort:       587,
			EncryptionType: "TLS",
		},
		UpdatedAt:     time.Now(),
	}
}
