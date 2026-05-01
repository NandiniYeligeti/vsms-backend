package requests

import "github.com/gin-gonic/gin"

type FollowUpRequest struct {
	Date     string `json:"date"`
	Remark   string `json:"remark"`
	Status   string `json:"status"`
	NextDate string `json:"next_date"`
}

type CreateEnquiryRequest struct {
	CompanyID   string `json:"company_id" binding:"required"`
	BranchID    string `json:"branch_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Mobile      string `json:"mobile" binding:"required"`
	Email       string `json:"email"`
	Vehicle     string `json:"vehicle"`
	Budget      string `json:"budget"`
	Salesperson string `json:"salesperson"`
}

type UpdateEnquiryRequest struct {
	Name        *string            `json:"name,omitempty"`
	Mobile      *string            `json:"mobile,omitempty"`
	Email       *string            `json:"email,omitempty"`
	Vehicle     *string            `json:"vehicle,omitempty"`
	Budget      *string            `json:"budget,omitempty"`
	Salesperson *string            `json:"salesperson,omitempty"`
	FollowUps   *[]FollowUpRequest `json:"follow_ups,omitempty"`
}

func (r *CreateEnquiryRequest) Validate(c *gin.Context) error {
	return c.ShouldBindJSON(r)
}

func (r *UpdateEnquiryRequest) Validate(c *gin.Context) error {
	return c.ShouldBindJSON(r)
}
