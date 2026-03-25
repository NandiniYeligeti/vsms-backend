package models

import (
	"vehiclesales/requests"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	EntityID     string             `bson:"entity_id" json:"entity_id"`
	CustomerCode string             `bson:"customer_code" json:"customer_code"`

	CompanyID string `bson:"company_id" json:"company_id"`
	BranchID  string `bson:"branch_id" json:"branch_id"`

	CustomerName  string `bson:"customer_name" json:"customer_name"`
	MobileNumber  string `bson:"mobile_number" json:"mobile_number"`
	Email         string `bson:"email" json:"email"`
	Address       string `bson:"address" json:"address"`
	City          string `bson:"city" json:"city"`
	State         string `bson:"state" json:"state"`
	Pincode       string `bson:"pincode" json:"pincode"`
	Photo         string `bson:"photo" json:"photo"`
	AadhaarCardNo string `bson:"aadhaar_card_no" json:"aadhaar_card_no"`
	PanCardNo     string `bson:"pan_card_no" json:"pan_card_no"`

	IsDeleted bool `bson:"is_deleted" json:"is_deleted"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type LedgerEntry struct {
	ID          string    `json:"id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Debit       float64   `json:"debit"`
	Credit      float64   `json:"credit"`
	Balance     float64   `json:"balance"`
	VehicleName string    `json:"vehicle_name"`
	VehicleID   string    `json:"vehicle_id"`
}

type UpdateCustomer struct {
	CustomerName  *string `bson:"customer_name,omitempty" json:"customer_name,omitempty"`
	MobileNumber  *string `bson:"mobile_number,omitempty" json:"mobile_number,omitempty"`
	Email         *string `bson:"email,omitempty" json:"email,omitempty"`
	Address       *string `bson:"address,omitempty" json:"address,omitempty"`
	City          *string `bson:"city,omitempty" json:"city,omitempty"`
	State         *string `bson:"state,omitempty" json:"state,omitempty"`
	Pincode       *string `bson:"pincode,omitempty" json:"pincode,omitempty"`
	Photo         *string `bson:"photo,omitempty" json:"photo,omitempty"`
	AadhaarCardNo *string `bson:"aadhaar_card_no,omitempty" json:"aadhaar_card_no,omitempty"`
	PanCardNo     *string `bson:"pan_card_no,omitempty" json:"pan_card_no,omitempty"`
}

func NewCustomer() *Customer {
	now := time.Now().UTC()
	id := primitive.NewObjectID()

	return &Customer{
		ID:           id,
		EntityID:     id.Hex(),
		CustomerCode: "CUST" + id.Hex()[18:24],
		IsDeleted:    false,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

func (c *Customer) Bind(req *requests.CreateCustomerRequest) {
	c.CompanyID = req.CompanyID
	c.BranchID = req.BranchID
	c.CustomerName = req.CustomerName
	c.MobileNumber = req.MobileNumber
	c.Email = req.Email
	c.Address = req.Address
	c.City = req.City
	c.State = req.State
	c.Pincode = req.Pincode
	if req.Photo != nil {
		c.Photo = req.Photo.Filename
	}
	c.AadhaarCardNo = req.AadhaarCardNo
	c.PanCardNo = req.PanCardNo
}