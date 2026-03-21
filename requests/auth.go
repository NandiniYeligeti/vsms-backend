package requests

import (
	"errors"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *LoginRequest) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(r); err != nil {
		return err
	}
	if r.Email == "" || r.Password == "" {
		return errors.New("email and password are required")
	}
	return nil
}

type CreateCompanyRequest struct {
	CompanyName    string `json:"company_name"`
	CompanyCode    string `json:"company_code"`
	AdminEmail     string `json:"admin_email"`
	AdminPassword  string `json:"admin_password"`
}

func (r *CreateCompanyRequest) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(r); err != nil {
		return err
	}
	if r.CompanyName == "" || r.CompanyCode == "" || r.AdminEmail == "" || r.AdminPassword == "" {
		return errors.New("company_name, company_code, admin_email, and admin_password are required")
	}
	return nil
}
