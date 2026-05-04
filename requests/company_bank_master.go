package requests

import (
	"github.com/gin-gonic/gin"
)

type CreateCompanyBankMasterRequest struct {
	BankName      string `json:"bank_name" binding:"required"`
	BranchName    string `json:"branch_name" binding:"required"`
	AccountNumber string `json:"account_number" binding:"required"`
	IsDefault     bool   `json:"is_default"`
}

func (r *CreateCompanyBankMasterRequest) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(r); err != nil {
		return err
	}
	return nil
}

type UpdateCompanyBankMasterRequest struct {
	BankName      *string `json:"bank_name,omitempty"`
	BranchName    *string `json:"branch_name,omitempty"`
	AccountNumber *string `json:"account_number,omitempty"`
	IsDefault     *bool   `json:"is_default,omitempty"`
}

func (r *UpdateCompanyBankMasterRequest) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(r); err != nil {
		return err
	}
	return nil
}
