package models

import (
	"time"
	"vehiclesales/requests"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VehicleType struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	EntityID   string             `bson:"entity_id" json:"entity_id"`
	CompanyID  string             `bson:"company_id" json:"company_id"`
	CategoryID string             `bson:"category_id" json:"category_id"`
	Code       string             `bson:"code" json:"code"`
	Name       string             `bson:"name" json:"name"`
	IsDeleted  bool               `bson:"is_deleted" json:"is_deleted"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
}

type UpdateVehicleType struct {
	CategoryID *string `bson:"category_id,omitempty" json:"category_id,omitempty"`
	Code       *string `bson:"code,omitempty" json:"code,omitempty"`
	Name       *string `bson:"name,omitempty" json:"name,omitempty"`
}

func NewVehicleType(req *requests.CreateVehicleTypeRequest) *VehicleType {
	now := time.Now().UTC()
	id := primitive.NewObjectID()
	return &VehicleType{
		ID:         id,
		EntityID:   id.Hex(),
		CompanyID:  req.CompanyID,
		CategoryID: req.CategoryID,
		Code:       req.Code,
		Name:       req.Name,
		IsDeleted:  false,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}

type VehicleCategory struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	EntityID  string             `bson:"entity_id" json:"entity_id"`
	CompanyID string             `bson:"company_id" json:"company_id"`
	Code      string             `bson:"code" json:"code"`
	Name      string             `bson:"name" json:"name"`
	IsDeleted bool               `bson:"is_deleted" json:"is_deleted"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type UpdateVehicleCategory struct {
	Code *string `bson:"code,omitempty" json:"code,omitempty"`
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
}

func NewVehicleCategory(req *requests.CreateVehicleCategoryRequest) *VehicleCategory {
	now := time.Now().UTC()
	id := primitive.NewObjectID()
	return &VehicleCategory{
		ID:        id,
		EntityID:  id.Hex(),
		CompanyID: req.CompanyID,
		Code:      req.Code,
		Name:      req.Name,
		IsDeleted: false,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

type VehicleAccessory struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	EntityID  string             `bson:"entity_id" json:"entity_id"`
	CompanyID string             `bson:"company_id" json:"company_id"`
	CategoryID string            `bson:"category_id" json:"category_id"`
	TypeID     string            `bson:"type_id" json:"type_id"`
	ModelID    string            `bson:"model_id" json:"model_id"`
	Code       string            `bson:"code" json:"code"`
	Name      string             `bson:"name" json:"name"`
	Price     float64            `bson:"price" json:"price"`
	IsDeleted bool               `bson:"is_deleted" json:"is_deleted"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type UpdateVehicleAccessory struct {
	CategoryID *string  `bson:"category_id,omitempty" json:"category_id,omitempty"`
	TypeID     *string  `bson:"type_id,omitempty" json:"type_id,omitempty"`
	ModelID    *string  `bson:"model_id,omitempty" json:"model_id,omitempty"`
	Code       *string  `bson:"code,omitempty" json:"code,omitempty"`
	Name       *string  `bson:"name,omitempty" json:"name,omitempty"`
	Price      *float64 `bson:"price,omitempty" json:"price,omitempty"`
}

func NewVehicleAccessory(req *requests.CreateVehicleAccessoryRequest) *VehicleAccessory {
	now := time.Now().UTC()
	id := primitive.NewObjectID()
	return &VehicleAccessory{
		ID:         id,
		EntityID:   id.Hex(),
		CompanyID:  req.CompanyID,
		CategoryID: req.CategoryID,
		TypeID:     req.TypeID,
		ModelID:    req.ModelID,
		Code:       req.Code,
		Name:       req.Name,
		Price:      req.Price,
		IsDeleted:  false,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}
