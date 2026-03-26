package models

import (
	"time"
	"vehiclesales/requests"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VehicleModel struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	EntityID         string             `bson:"entity_id" json:"entity_id"`
	VehicleModelCode string             `bson:"vehicle_model_code" json:"vehicle_model_code"`

	CompanyID string `bson:"company_id" json:"company_id"`
	BranchID  string `bson:"branch_id" json:"branch_id"`

	Brand     string  `bson:"brand" json:"brand"`
	Model     string  `bson:"model" json:"model"`
	Variant   string  `bson:"variant" json:"variant"`
	FuelType  string  `bson:"fuel_type" json:"fuel_type"`
	BasePrice float64 `bson:"base_price" json:"base_price"`

	TypeID     string `bson:"type_id" json:"type_id"`
	CategoryID string `bson:"category_id" json:"category_id"`
	Colors     []string `bson:"colors" json:"colors"`

	IsDeleted bool `bson:"is_deleted" json:"is_deleted"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type UpdateVehicleModel struct {
	Brand      *string  `bson:"brand,omitempty" json:"brand,omitempty"`
	Model      *string  `bson:"model,omitempty" json:"model,omitempty"`
	Variant    *string  `bson:"variant,omitempty" json:"variant,omitempty"`
	FuelType   *string  `bson:"fuel_type,omitempty" json:"fuel_type,omitempty"`
	BasePrice  *float64 `bson:"base_price,omitempty" json:"base_price,omitempty"`
	TypeID     *string  `bson:"type_id,omitempty" json:"type_id,omitempty"`
	CategoryID *string  `bson:"category_id,omitempty" json:"category_id,omitempty"`
	Colors     *[]string `bson:"colors,omitempty" json:"colors,omitempty"`
}

func NewVehicleModel() *VehicleModel {
	now := time.Now().UTC()
	id := primitive.NewObjectID()

	return &VehicleModel{
		ID:               id,
		EntityID:         id.Hex(),
		VehicleModelCode: "VEH" + id.Hex()[18:24],
		IsDeleted:        false,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
}

func (v *VehicleModel) Bind(req *requests.CreateVehicleModelRequest) {
	v.CompanyID = req.CompanyID
	v.BranchID = req.BranchID
	v.Brand = req.Brand
	v.Model = req.Model
	v.Variant = req.Variant
	v.FuelType = req.FuelType
	v.BasePrice = req.BasePrice
	v.TypeID = req.TypeID
	v.CategoryID = req.CategoryID
	v.Colors = req.Colors
}