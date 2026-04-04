package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompanyMaster struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	EntityID  string             `bson:"entity_id" json:"entity_id"`
	CompanyID string             `bson:"company_id" json:"company_id"`
	Type      string             `bson:"type" json:"type"` // "Showroom", "Branch", "Area"
	Name      string             `bson:"name" json:"name"`
	IsDeleted bool               `bson:"is_deleted" json:"is_deleted"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

func NewCompanyMaster(companyID, masterType, name string) *CompanyMaster {
	now := time.Now().UTC()
	id := primitive.NewObjectID()
	return &CompanyMaster{
		ID:        id,
		EntityID:  id.Hex(),
		CompanyID: companyID,
		Type:      masterType,
		Name:      name,
		IsDeleted: false,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
