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

	Showroom string `json:"showroom"`
	Branch   string `json:"branch"`
	Area     string `json:"area"`
}

type UpdateSalespersonRequest struct {
	FullName     *string `json:"full_name,omitempty"`
	MobileNumber *string `json:"mobile_number,omitempty"`
	Email        *string `json:"email,omitempty" binding:"omitempty,email"`
	BranchID     *string `json:"branch_id,omitempty"`
	Showroom     *string `json:"showroom,omitempty"`
	Branch       *string `json:"branch,omitempty"`
	Area         *string `json:"area,omitempty"`
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