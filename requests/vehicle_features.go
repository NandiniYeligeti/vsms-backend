package requests

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type CreateVehicleTypeRequest struct {
	CompanyID string `json:"company_id"`
	Name      string `json:"name"`
}

func (r *CreateVehicleTypeRequest) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(r); err != nil {
		return err
	}
	if r.Name == "" {
		return errors.New("name is required")
	}
	return nil
}

type CreateVehicleCategoryRequest struct {
	CompanyID string `json:"company_id"`
	Name      string `json:"name"`
}

func (r *CreateVehicleCategoryRequest) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(r); err != nil {
		return err
	}
	if r.Name == "" {
		return errors.New("name is required")
	}
	return nil
}

type CreateVehicleAccessoryRequest struct {
	CompanyID string  `json:"company_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
}

func (r *CreateVehicleAccessoryRequest) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(r); err != nil {
		return err
	}
	if r.Name == "" {
		return errors.New("name is required")
	}
	if r.Price < 0 {
		return errors.New("price cannot be negative")
	}
	return nil
}
