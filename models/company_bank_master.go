package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompanyBankMaster struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	EntityID      string             `bson:"entity_id" json:"entity_id"`
	CompanyID     string             `bson:"company_id" json:"company_id"`
	BankName      string             `bson:"bank_name" json:"bank_name"`
	BranchName    string             `bson:"branch_name" json:"branch_name"`
	AccountNumber string             `bson:"account_number" json:"account_number"`
	IsDefault     bool               `bson:"is_default" json:"is_default"`
	IsDeleted     bool               `bson:"is_deleted" json:"is_deleted"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at" json:"updated_at"`
}

func NewCompanyBankMaster() *CompanyBankMaster {
	now := time.Now().UTC()
	id := primitive.NewObjectID()
	return &CompanyBankMaster{
		ID:        id,
		EntityID:  id.Hex(),
		IsDeleted: false,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
