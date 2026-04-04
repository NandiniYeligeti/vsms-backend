package requests

import (
	"errors"
	"github.com/gin-gonic/gin"
)

type CreateCompanyMasterRequest struct {
	CompanyID string `json:"company_id"`
	Type      string `json:"type"` // "Showroom", "Branch", "Area"
	Name      string `json:"name"`
}

func (r *CreateCompanyMasterRequest) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(r); err != nil {
		return err
	}
	if r.Type == "" {
		return errors.New("type is required")
	}
	if r.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
