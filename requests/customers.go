package requests

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type CreateCustomerRequest struct {
	CompanyID     string                `json:"company_id" form:"company_id" binding:"required"`
	BranchID      string                `json:"branch_id" form:"branch_id" binding:"required"`
	CustomerName  string                `json:"customer_name" form:"customer_name" binding:"required"`
	MobileNumber  string                `json:"mobile_number" form:"mobile_number" binding:"required"`
	Email         string                `json:"email" form:"email"`
	Address       string                `json:"address" form:"address" binding:"required"`
	City          string                `json:"city" form:"city" binding:"required"`
	State         string                `json:"state" form:"state" binding:"required"`
	Pincode       string                `json:"pincode" form:"pincode" binding:"required"`
	Photo         *multipart.FileHeader `json:"photo" form:"photo"`
	AadhaarCardNo string                `json:"aadhaar_card_no" form:"aadhaar_card_no"`
	PanCardNo     string                `json:"pan_card_no" form:"pan_card_no"`
}

type UpdateCustomerRequest struct {
	CustomerName  *string               `json:"customer_name,omitempty" form:"customer_name"`
	MobileNumber  *string               `json:"mobile_number,omitempty" form:"mobile_number"`
	Email         *string               `json:"email,omitempty" form:"email"`
	Address       *string               `json:"address,omitempty" form:"address"`
	City          *string               `json:"city,omitempty" form:"city"`
	State         *string               `json:"state,omitempty" form:"state"`
	Pincode       *string               `json:"pincode,omitempty" form:"pincode"`
	Photo         *multipart.FileHeader `json:"photo,omitempty" form:"photo"`
	AadhaarCardNo *string               `json:"aadhaar_card_no,omitempty" form:"aadhaar_card_no"`
	PanCardNo     *string                 `json:"pan_card_no,omitempty" form:"pan_card_no"`
	Documents     []*multipart.FileHeader `json:"documents,omitempty" form:"documents"`
}

func NewCreateCustomerRequest() *CreateCustomerRequest {
	return &CreateCustomerRequest{}
}

func NewUpdateCustomerRequest() *UpdateCustomerRequest {
	return &UpdateCustomerRequest{}
}
func (r *CreateCustomerRequest) Validate(c *gin.Context) error {
	return c.ShouldBind(r)
}
func (r *UpdateCustomerRequest) Validate(c *gin.Context) error {
	return c.ShouldBind(r)
}