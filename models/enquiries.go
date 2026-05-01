package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"vehiclesales/requests"
)

type FollowUp struct {
	Date     string `bson:"date" json:"date"`
	Remark   string `bson:"remark" json:"remark"`
	Status   string `bson:"status" json:"status"`
	NextDate string `bson:"next_date" json:"next_date"`
}

type Enquiry struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	EntityID    string             `bson:"entity_id" json:"entity_id"`
	CompanyID   string             `bson:"company_id" json:"company_id"`
	BranchID    string             `bson:"branch_id" json:"branch_id"`
	Name        string             `bson:"name" json:"name"`
	Mobile      string             `bson:"mobile" json:"mobile"`
	Email       string             `bson:"email" json:"email"`
	Vehicle     string             `bson:"vehicle" json:"vehicle"`
	Budget      string             `bson:"budget" json:"budget"`
	Salesperson string             `bson:"salesperson" json:"salesperson"`
	FollowUps   []FollowUp         `bson:"follow_ups" json:"follow_ups"`
	IsDeleted   bool               `bson:"is_deleted" json:"is_deleted"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

type UpdateEnquiry struct {
	Name        *string     `bson:"name,omitempty" json:"name,omitempty"`
	Mobile      *string     `bson:"mobile,omitempty" json:"mobile,omitempty"`
	Email       *string     `bson:"email,omitempty" json:"email,omitempty"`
	Vehicle     *string     `bson:"vehicle,omitempty" json:"vehicle,omitempty"`
	Budget      *string     `bson:"budget,omitempty" json:"budget,omitempty"`
	Salesperson *string     `bson:"salesperson,omitempty" json:"salesperson,omitempty"`
	FollowUps   *[]FollowUp `bson:"follow_ups,omitempty" json:"follow_ups,omitempty"`
}

func NewEnquiry() *Enquiry {
	now := time.Now().UTC()
	id := primitive.NewObjectID()

	return &Enquiry{
		ID:        id,
		EntityID:  id.Hex(),
		FollowUps: []FollowUp{},
		IsDeleted: false,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (e *Enquiry) Bind(req *requests.CreateEnquiryRequest) {
	e.CompanyID = req.CompanyID
	e.BranchID = req.BranchID
	e.Name = req.Name
	e.Mobile = req.Mobile
	e.Email = req.Email
	e.Vehicle = req.Vehicle
	e.Budget = req.Budget
	e.Salesperson = req.Salesperson
}
