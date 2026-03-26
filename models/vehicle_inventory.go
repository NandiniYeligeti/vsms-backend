package models

import (
	"time"
	"vehiclesales/requests"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VehicleInventory struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	EntityID      string             `bson:"entity_id" json:"entity_id"`
	InventoryCode string             `bson:"inventory_code" json:"inventory_code"`

	CompanyID string `bson:"company_id" json:"company_id"`
	BranchID  string `bson:"branch_id" json:"branch_id"`

	VehicleModelID string `bson:"vehicle_model_id" json:"vehicle_model_id"`

	Brand     string  `bson:"brand" json:"brand"`
	Model     string  `bson:"model" json:"model"`
	Variant   string  `bson:"variant" json:"variant"`
	FuelType  string  `bson:"fuel_type" json:"fuel_type"`
	BasePrice float64 `bson:"base_price" json:"base_price"`

	Accessories []string `bson:"accessories" json:"accessories"`
	TotalPrice  float64  `bson:"total_price" json:"total_price"`
	SellingPrice float64 `bson:"selling_price" json:"selling_price"`

	Color         string    `bson:"color" json:"color"`
	ChassisNumber string    `bson:"chassis_number" json:"chassis_number"`
	EngineNumber  string    `bson:"engine_number" json:"engine_number"`
	PurchaseDate  time.Time `bson:"purchase_date" json:"purchase_date"`

	Status string `bson:"status" json:"status"`

	IsDeleted bool `bson:"is_deleted" json:"is_deleted"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type UpdateVehicleInventory struct {
	Color         *string    `bson:"color,omitempty" json:"color,omitempty"`
	ChassisNumber *string    `bson:"chassis_number,omitempty" json:"chassis_number,omitempty"`
	EngineNumber  *string    `bson:"engine_number,omitempty" json:"engine_number,omitempty"`
	PurchaseDate  *time.Time `bson:"purchase_date,omitempty" json:"purchase_date,omitempty"`
	SellingPrice  *float64   `bson:"selling_price,omitempty" json:"selling_price,omitempty"`
}

func NewVehicleInventory() *VehicleInventory {
	now := time.Now().UTC()
	id := primitive.NewObjectID()

	return &VehicleInventory{
		ID:            id,
		EntityID:      id.Hex(),
		InventoryCode: "INV" + id.Hex()[18:24],
		Status:        "Available",
		IsDeleted:     false,
		CreatedAt:     now,
		UpdatedAt:     now,
	}
}

func (v *VehicleInventory) Bind(req *requests.CreateVehicleInventoryRequest) {
	v.CompanyID = req.CompanyID
	v.BranchID = req.BranchID
	v.VehicleModelID = req.VehicleModelID
	v.Color = req.Color
	v.ChassisNumber = req.ChassisNumber
	v.EngineNumber = req.EngineNumber
	v.PurchaseDate = req.PurchaseDate
	v.Accessories = req.Accessories
	v.TotalPrice = req.TotalPrice
	if req.Status != "" {
		v.Status = req.Status
	}
	v.SellingPrice = req.SellingPrice
}