package requests

import (
	"errors"
	"github.com/gin-gonic/gin"
)

type MenuPermission struct {
	MenuID    string `json:"menu_id" bson:"menu_id"`
	CanView   bool   `json:"can_view" bson:"can_view"`
	CanAdd    bool   `json:"can_add" bson:"can_add"`
	CanEdit   bool   `json:"can_edit" bson:"can_edit"`
	CanDelete bool   `json:"can_delete" bson:"can_delete"`
}

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

type CreateUserRequest struct {
	Username    string           `json:"username"`
	Email       string           `json:"email"`
	Password    string           `json:"password"`
	Menus       []string         `json:"menus"`
	Permissions []MenuPermission `json:"permissions"`
	Branches    []string         `json:"branches"`
	Showrooms   []string         `json:"showrooms"`
	Areas       []string         `json:"areas"`
}

func (r *CreateUserRequest) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(r); err != nil {
		return err
	}
	if r.Username == "" || r.Email == "" || r.Password == "" {
		return errors.New("username, email, and password are required")
	}
	return nil
}

type UpdateUserMenusRequest struct {
	Menus       []string         `json:"menus"`
	Permissions []MenuPermission `json:"permissions"`
	Branches    []string         `json:"branches"`
	Showrooms   []string         `json:"showrooms"`
	Areas       []string         `json:"areas"`
}

func (r *UpdateUserMenusRequest) Validate(c *gin.Context) error {
	return c.ShouldBindJSON(r)
}

type UpdatePasswordRequest struct {
	Password string `json:"password"`
}

func (r *UpdatePasswordRequest) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(r); err != nil {
		return err
	}
	if r.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

func (r *ForgotPasswordRequest) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(r); err != nil {
		return err
	}
	if r.Email == "" {
		return errors.New("email is required")
	}
	return nil
}

