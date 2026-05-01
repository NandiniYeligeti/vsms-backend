package models

import (
	"time"
	"vehiclesales/requests"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VehicleModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	EntityID  string             `bson:"entity_id" json:"entity_id"`
	ModelCode string             `bson:"model_code" json:"model_code"`

	CompanyID string `bson:"company_id" json:"company_id"`
	BranchID  string `bson:"branch_id" json:"branch_id"`

	Brand     string   `bson:"brand" json:"brand"`
	Model     string   `bson:"model" json:"model"`
	Variant   string   `bson:"variant" json:"variant"`
	FuelType  []string `bson:"fuel_type" json:"fuel_type"`
	BasePrice float64  `bson:"base_price" json:"base_price"`

	TypeID     string   `bson:"type_id" json:"type_id"`
	CategoryID string   `bson:"category_id" json:"category_id"`
	Colors     []string `bson:"colors" json:"colors"`

	IncentiveType  string  `bson:"incentive_type" json:"incentive_type"`
	IncentiveValue float64 `bson:"incentive_value" json:"incentive_value"`
	ColorCount     int     `bson:"color_count" json:"color_count"`

	// Per-variant spec fields
	Transmission   string  `bson:"transmission" json:"transmission"`
	EngineCC       float64 `bson:"engine_cc" json:"engine_cc"`
	BatteryKWh     float64 `bson:"battery_kwh" json:"battery_kwh"`
	ChargingTime   string  `bson:"charging_time" json:"charging_time"`
	TankCapacity   float64 `bson:"tank_capacity" json:"tank_capacity"`
	AverageMileage float64 `bson:"average_mileage" json:"average_mileage"`

	IsDeleted bool `bson:"is_deleted" json:"is_deleted"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type UpdateVehicleModel struct {
	ModelCode  *string   `bson:"model_code,omitempty" json:"model_code,omitempty"`
	Brand      *string   `bson:"brand,omitempty" json:"brand,omitempty"`
	Model      *string   `bson:"model,omitempty" json:"model,omitempty"`
	Variant    *string   `bson:"variant,omitempty" json:"variant,omitempty"`
	FuelType   *[]string `bson:"fuel_type,omitempty" json:"fuel_type,omitempty"`
	BasePrice  *float64  `bson:"base_price,omitempty" json:"base_price,omitempty"`
	TypeID     *string   `bson:"type_id,omitempty" json:"type_id,omitempty"`
	CategoryID *string   `bson:"category_id,omitempty" json:"category_id,omitempty"`
	Colors     *[]string `bson:"colors,omitempty" json:"colors,omitempty"`

	IncentiveType  *string  `bson:"incentive_type,omitempty" json:"incentive_type,omitempty"`
	IncentiveValue *float64 `bson:"incentive_value,omitempty" json:"incentive_value,omitempty"`
	ColorCount     *int     `bson:"color_count,omitempty" json:"color_count,omitempty"`

	// Per-variant spec fields
	Transmission   *string  `bson:"transmission,omitempty" json:"transmission,omitempty"`
	EngineCC       *float64 `bson:"engine_cc,omitempty" json:"engine_cc,omitempty"`
	BatteryKWh     *float64 `bson:"battery_kwh,omitempty" json:"battery_kwh,omitempty"`
	ChargingTime   *string  `bson:"charging_time,omitempty" json:"charging_time,omitempty"`
	TankCapacity   *float64 `bson:"tank_capacity,omitempty" json:"tank_capacity,omitempty"`
	AverageMileage *float64 `bson:"average_mileage,omitempty" json:"average_mileage,omitempty"`
}

func NewVehicleModel() *VehicleModel {
	now := time.Now().UTC()
	id := primitive.NewObjectID()

	return &VehicleModel{
		ID:        id,
		EntityID:  id.Hex(),
		IsDeleted: false,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (v *VehicleModel) Bind(req *requests.CreateVehicleModelRequest) {
	v.CompanyID = req.CompanyID
	v.BranchID = req.BranchID
	v.ModelCode = req.ModelCode
	v.Brand = req.Brand
	v.Model = req.Model
	v.Variant = req.Variant
	v.FuelType = req.FuelType
	v.BasePrice = req.BasePrice
	v.TypeID = req.TypeID
	v.CategoryID = req.CategoryID
	v.Colors = req.Colors
	v.IncentiveType = req.IncentiveType
	v.IncentiveValue = req.IncentiveValue
	v.ColorCount = req.ColorCount
	v.Transmission = req.Transmission
	v.EngineCC = req.EngineCC
	v.BatteryKWh = req.BatteryKWh
	v.ChargingTime = req.ChargingTime
	v.TankCapacity = req.TankCapacity
	v.AverageMileage = req.AverageMileage
}