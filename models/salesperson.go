package models

import (
	"time"
	"vehiclesales/requests"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Salesperson struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	EntityID        string             `bson:"entity_id" json:"entity_id"`
	SalespersonCode string             `bson:"salesperson_code" json:"salesperson_code"`

	CompanyID string `bson:"company_id" json:"company_id"`
	BranchID  string `bson:"branch_id" json:"branch_id"`

	FullName     string `bson:"full_name" json:"full_name"`
	MobileNumber string `bson:"mobile_number" json:"mobile_number"`
	Email        string `bson:"email" json:"email"`

	CommissionType  string  `bson:"commission_type" json:"commissionType"`
	CommissionValue float64 `bson:"commission_value" json:"commissionValue"`
	Vehicle         string  `bson:"vehicle,omitempty" json:"vehicle,omitempty"`

	IsDeleted bool `bson:"is_deleted" json:"is_deleted"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type UpdateSalesperson struct {
	FullName     *string `bson:"full_name,omitempty" json:"full_name,omitempty"`
	MobileNumber *string `bson:"mobile_number,omitempty" json:"mobile_number,omitempty"`
	Email        *string `bson:"email,omitempty" json:"email,omitempty"`
	BranchID     *string `bson:"branch_id,omitempty" json:"branch_id,omitempty"`

	CommissionType  *string  `bson:"commission_type,omitempty" json:"commissionType,omitempty"`
	CommissionValue *float64 `bson:"commission_value,omitempty" json:"commissionValue,omitempty"`
	Vehicle         *string  `bson:"vehicle,omitempty" json:"vehicle,omitempty"`
}

func NewSalesperson() *Salesperson {
	now := time.Now().UTC()
	id := primitive.NewObjectID()

	return &Salesperson{
		ID:              id,
		EntityID:        id.Hex(),
		SalespersonCode: "SALE" + id.Hex()[18:24],
		IsDeleted:       false,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
}

func (s *Salesperson) Bind(req *requests.CreateSalespersonRequest) {
	s.CompanyID = req.CompanyID
	s.BranchID = req.BranchID
	s.FullName = req.FullName
	s.MobileNumber = req.MobileNumber
	s.Email = req.Email

	s.CommissionType = req.CommissionType
	s.CommissionValue = req.CommissionValue
	s.Vehicle = req.Vehicle
}