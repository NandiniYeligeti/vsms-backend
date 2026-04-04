package requests

import (
	"github.com/gin-gonic/gin"
)

type CreateBankMasterRequest struct {
	BankName      string  `json:"bank_name" binding:"required"`
	BranchName    string  `json:"branch_name"`
	ContactPerson string  `json:"contact_person"`
	ContactNumber string  `json:"contact_number"`
	CompanyID     string  `json:"company_id"`
}

func (r *CreateBankMasterRequest) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(r); err != nil {
		return err
	}
	return nil
}

type UpdateBankMasterRequest struct {
	BankName      *string `json:"bank_name,omitempty"`
	BranchName    *string `json:"branch_name,omitempty"`
	ContactPerson *string `json:"contact_person,omitempty"`
	ContactNumber *string `json:"contact_number,omitempty"`
}

func (r *UpdateBankMasterRequest) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(r); err != nil {
		return err
	}
	return nil
}

func NewCreateBankMasterRequest() *CreateBankMasterRequest {
	return &CreateBankMasterRequest{}
}

func NewUpdateBankMasterRequest() *UpdateBankMasterRequest {
	return &UpdateBankMasterRequest{}
}
