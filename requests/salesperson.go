package requests

import (
	"github.com/gin-gonic/gin"
)

type CreateSalespersonRequest struct {
	CompanyID string `json:"company_id" binding:"required"`
	BranchID  string `json:"branch_id" binding:"required"`

	FullName     string `json:"full_name" binding:"required"`
	MobileNumber string `json:"mobile_number" binding:"required"`
	Email        string `json:"email" binding:"omitempty,email"`

	CommissionType  string  `json:"commissionType" binding:"required"`
	CommissionValue float64 `json:"commissionValue" binding:"required"`
	Vehicle         string  `json:"vehicle,omitempty"`
}

type UpdateSalespersonRequest struct {
	FullName     *string `json:"full_name,omitempty"`
	MobileNumber *string `json:"mobile_number,omitempty"`
	Email        *string `json:"email,omitempty" binding:"omitempty,email"`
	BranchID     *string `json:"branch_id,omitempty"`

	CommissionType  *string  `json:"commissionType,omitempty"`
	CommissionValue *float64 `json:"commissionValue,omitempty"`
	Vehicle         *string  `json:"vehicle,omitempty"`
}

func NewCreateSalespersonRequest() *CreateSalespersonRequest {
	return &CreateSalespersonRequest{}
}

func NewUpdateSalespersonRequest() *UpdateSalespersonRequest {
	return &UpdateSalespersonRequest{}
}

func (r *CreateSalespersonRequest) Validate(c *gin.Context) error {
	return c.ShouldBindJSON(r)
}

func (r *UpdateSalespersonRequest) Validate(c *gin.Context) error {
	return c.ShouldBindJSON(r)
}