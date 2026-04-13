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

	Showroom string `bson:"showroom" json:"showroom"`
	Branch   string `bson:"branch" json:"branch"`
	Area     string `bson:"area" json:"area"`

	IsInactive   bool       `bson:"is_inactive" json:"is_inactive"`
	InactiveDate *time.Time `bson:"inactive_date,omitempty" json:"inactive_date,omitempty"`

	IsDeleted bool `bson:"is_deleted" json:"is_deleted"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type UpdateSalesperson struct {
	FullName     *string `bson:"full_name,omitempty" json:"full_name,omitempty"`
	MobileNumber *string `bson:"mobile_number,omitempty" json:"mobile_number,omitempty"`
	Email        *string `bson:"email,omitempty" json:"email,omitempty"`
	BranchID     *string `bson:"branch_id,omitempty" json:"branch_id,omitempty"`
	Showroom     *string `bson:"showroom,omitempty" json:"showroom,omitempty"`
	Branch       *string `bson:"branch,omitempty" json:"branch,omitempty"`
	Area         *string `bson:"area,omitempty" json:"area,omitempty"`
	IsInactive   *bool      `bson:"is_inactive,omitempty" json:"is_inactive,omitempty"`
	InactiveDate *time.Time `bson:"inactive_date,omitempty" json:"inactive_date,omitempty"`
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
	s.Showroom = req.Showroom
	s.Branch = req.Branch
	s.Area = req.Area
}