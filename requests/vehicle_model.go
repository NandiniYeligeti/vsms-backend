package requests

import (
	"github.com/gin-gonic/gin"
)

type CreateVehicleModelRequest struct {
	CompanyID  string  `json:"company_id" binding:"required"`
	BranchID   string  `json:"branch_id" binding:"required"`
	TypeID     string  `json:"type_id" binding:"required"`
	CategoryID string  `json:"category_id" binding:"required"`
	Brand      string  `json:"brand" binding:"required"`
	Model      string  `json:"model" binding:"required"`
	Variant    string  `json:"variant" binding:"required"`
	FuelType   string  `json:"fuel_type" binding:"required"`
	BasePrice  float64  `json:"base_price" binding:"required"`
	Colors     []string `json:"colors"`
}

type UpdateVehicleModelRequest struct {
	TypeID     *string  `json:"type_id,omitempty"`
	CategoryID *string  `json:"category_id,omitempty"`
	Brand      *string  `json:"brand,omitempty"`
	Model      *string  `json:"model,omitempty"`
	Variant    *string  `json:"variant,omitempty"`
	FuelType   *string  `json:"fuel_type,omitempty"`
	BasePrice  *float64  `json:"base_price,omitempty"`
	Colors     *[]string `json:"colors,omitempty"`
}

func NewCreateVehicleModelRequest() *CreateVehicleModelRequest {
	return &CreateVehicleModelRequest{}
}

func NewUpdateVehicleModelRequest() *UpdateVehicleModelRequest {
	return &UpdateVehicleModelRequest{}
}

func (r *CreateVehicleModelRequest) Validate(c *gin.Context) error {
	return c.ShouldBindJSON(r)
}

func (r *UpdateVehicleModelRequest) Validate(c *gin.Context) error {
	return c.ShouldBindJSON(r)
}