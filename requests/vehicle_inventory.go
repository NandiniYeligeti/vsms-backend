package requests

import (
	"time"

	"github.com/gin-gonic/gin"
)

type CreateVehicleInventoryRequest struct {
	CompanyID      string    `json:"company_id" binding:"required"`
	BranchID       string    `json:"branch_id" binding:"required"`
	VehicleModelID string    `json:"vehicle_model_id" binding:"required"`
	Color          string    `json:"color" binding:"required"`
	ChassisNumber  string    `json:"chassis_number" binding:"required"`
	EngineNumber   string    `json:"engine_number" binding:"required"`
	PurchaseDate   time.Time `json:"purchase_date" binding:"required"`
	MfgYear        string    `json:"mfg_year" binding:"required"`
	InventoryDate  time.Time `json:"inventory_date" binding:"required"`
	PurchasePrice  float64   `json:"purchase_price" binding:"required"`
	Accessories    []string  `json:"accessories"`
	TotalPrice     float64   `json:"total_price"`
	Status         string    `json:"status"`
	SellingPrice   float64   `json:"selling_price"`
}

type UpdateVehicleInventoryRequest struct {
	Color         *string    `json:"color,omitempty"`
	ChassisNumber *string    `json:"chassis_number,omitempty"`
	EngineNumber  *string    `json:"engine_number,omitempty"`
	PurchaseDate  *time.Time `json:"purchase_date,omitempty"`
	MfgYear       *string    `json:"mfg_year,omitempty"`
	InventoryDate *time.Time `json:"inventory_date,omitempty"`
	PurchasePrice *float64   `json:"purchase_price,omitempty"`
	Accessories   *[]string  `json:"accessories,omitempty"`
	TotalPrice    *float64   `json:"total_price,omitempty"`
	Status        *string    `json:"status,omitempty"`
	SellingPrice  *float64   `json:"selling_price,omitempty"`
}

func NewCreateVehicleInventoryRequest() *CreateVehicleInventoryRequest {
	return &CreateVehicleInventoryRequest{}
}

func NewUpdateVehicleInventoryRequest() *UpdateVehicleInventoryRequest {
	return &UpdateVehicleInventoryRequest{}
}

func (r *CreateVehicleInventoryRequest) Validate(c *gin.Context) error {
	return c.ShouldBindJSON(r)
}

func (r *UpdateVehicleInventoryRequest) Validate(c *gin.Context) error {
	return c.ShouldBindJSON(r)
}