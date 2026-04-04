package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BankMaster struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	EntityID      string             `bson:"entity_id" json:"entity_id"`
	CompanyID     string             `bson:"company_id" json:"company_id"`
	BankName      string             `bson:"bank_name" json:"bank_name"`
	BranchName    string             `bson:"branch_name" json:"branch_name"`
	ContactPerson string             `bson:"contact_person" json:"contact_person"`
	ContactNumber string             `bson:"contact_number" json:"contact_number"`
	IsDeleted     bool               `bson:"is_deleted" json:"is_deleted"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at" json:"updated_at"`
}

func NewBankMaster() *BankMaster {
	now := time.Now().UTC()
	id := primitive.NewObjectID()
	return &BankMaster{
		ID:        id,
		EntityID:  id.Hex(),
		IsDeleted: false,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
