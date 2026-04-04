package requests

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type CreateVehicleTypeRequest struct {
	CompanyID  string `json:"company_id"`
	CategoryID string `json:"category_id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
}

func (r *CreateVehicleTypeRequest) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(r); err != nil {
		return err
	}
	if r.CategoryID == "" {
		return errors.New("category is required")
	}
	if r.Code == "" {
		return errors.New("type code is required")
	}
	if r.Name == "" {
		return errors.New("type name is required")
	}
	return nil
}

type CreateVehicleCategoryRequest struct {
	CompanyID string `json:"company_id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
}

func (r *CreateVehicleCategoryRequest) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(r); err != nil {
		return err
	}
	if r.Code == "" {
		return errors.New("category code is required")
	}
	if r.Name == "" {
		return errors.New("category name is required")
	}
	return nil
}

type CreateVehicleAccessoryRequest struct {
	CompanyID  string  `json:"company_id"`
	CategoryID string  `json:"category_id"`
	TypeID     string  `json:"type_id"`
	ModelID    string  `json:"model_id"`
	Code       string  `json:"code"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
}

func (r *CreateVehicleAccessoryRequest) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(r); err != nil {
		return err
	}
	if r.CategoryID == "" || r.TypeID == "" || r.ModelID == "" {
		return errors.New("category, type, and model are required")
	}
	if r.Name == "" {
		return errors.New("name is required")
	}
	if r.Price < 0 {
		return errors.New("price cannot be negative")
	}
	return nil
}
