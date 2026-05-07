package models

import (
	"time"
	"vehiclesales/requests"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IncentiveLog struct {
	Action      string    `bson:"action" json:"action"`             // Generated, Edited, Paid
	Description string    `bson:"description" json:"description"`
	Timestamp   time.Time `bson:"timestamp" json:"timestamp"`
}

type SalesOrder struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	EntityID       string            `bson:"entity_id" json:"entity_id"`
	SalesOrderCode string            `bson:"sales_order_code" json:"sales_order_code"`

	CompanyID string `bson:"company_id" json:"company_id"`
	BranchID  string `bson:"branch_id" json:"branch_id"`

	CustomerID string `bson:"customer_id" json:"customer_id"`

	CustomerName string `bson:"customer_name" json:"customer_name"`
	MobileNumber string `bson:"mobile_number" json:"mobile_number"`
	Email        string `bson:"email" json:"email"`
	Address      string `bson:"address" json:"address"`

	VehicleInventoryID string `bson:"vehicle_inventory_id" json:"vehicle_inventory_id"`

	Brand         string  `bson:"brand" json:"brand"`
	Model         string  `bson:"model" json:"model"`
	Variant       string  `bson:"variant" json:"variant"`
	Color         string  `bson:"color" json:"color"`
	ChassisNumber string  `bson:"chassis_number" json:"chassis_number"`
	EngineNumber  string  `bson:"engine_number" json:"engine_number"`

	SalespersonID string `bson:"salesperson_id" json:"salesperson_id"`
	SalespersonName string `bson:"salesperson_name" json:"salesperson_name"`

	SaleDate     time.Time `bson:"sale_date" json:"sale_date"`
	DeliveryDate time.Time `bson:"delivery_date" json:"delivery_date"`

	VehiclePrice         float64 `bson:"vehicle_price" json:"vehicle_price"`
	RegistrationCharges  float64 `bson:"registration_charges" json:"registration_charges"`
	Insurance            float64 `bson:"insurance" json:"insurance"`
	Accessories          float64 `bson:"accessories" json:"accessories"`
	DiscountType         string  `bson:"discount_type" json:"discount_type"`
	DiscountValue        float64 `bson:"discount_value" json:"discount_value"`
	DiscountReason       string  `bson:"discount_reason" json:"discount_reason"`
	DiscountAmount       float64 `bson:"discount_amount" json:"discount_amount"`

	TotalAmount   float64 `bson:"total_amount" json:"total_amount"`
	DownPayment   float64 `bson:"down_payment" json:"down_payment"`
	LoanAmount    float64 `bson:"loan_amount" json:"loan_amount"`
	BalanceAmount float64 `bson:"balance_amount" json:"balance_amount"`
	PaymentType   string  `bson:"payment_type" json:"payment_type"`
	PaymentMode string  `bson:"payment_mode" json:"payment_mode"`
	LoanStatus  string  `bson:"loan_status" json:"loan_status"`
	UTRNumber   string  `bson:"utr_number" json:"utr_number"`

	Status string `bson:"status" json:"status"`
	DeliveryStatus string `bson:"delivery_status" json:"delivery_status"` // Pending, Ready, Delivered, Cancelled
	ActualDeliveryDate *time.Time `bson:"actual_delivery_date" json:"actual_delivery_date"`
	EmailStatus string `bson:"email_status" json:"email_status"` // Sent, Failed, Pending

	// Registration Details
	RegistrationStatus        string     `bson:"registration_status" json:"registration_status"` // Pending, In Process, Completed
	RegistrationApplicationNo string     `bson:"registration_application_no" json:"registration_application_no"`
	RegistrationTempNo        string     `bson:"registration_temp_no" json:"registration_temp_no"`
	RegistrationAgentName     string     `bson:"registration_agent_name" json:"registration_agent_name"`
	RegistrationAgentMobile   string     `bson:"registration_agent_mobile" json:"registration_agent_mobile"`
	RegistrationDate          *time.Time `bson:"registration_date" json:"registration_date"`
	RegistrationDocumentUrl   string     `bson:"registration_document_url" json:"registration_document_url"`
	VehicleNumber             string     `bson:"vehicle_number" json:"vehicle_number"`
	RTO                       string     `bson:"rto" json:"rto"`

	// Insurance Details
	InsuranceCompany   string     `bson:"insurance_company" json:"insurance_company"`
	InsurancePolicyNo  string     `bson:"insurance_policy_no" json:"insurance_policy_no"`
	InsuranceStartDate *time.Time `bson:"insurance_start_date" json:"insurance_start_date"`
	InsuranceEndDate   *time.Time `bson:"insurance_end_date" json:"insurance_end_date"`

	IncentiveAmount float64   `bson:"incentive_amount" json:"incentive_amount"`
	IncentiveStatus string    `bson:"incentive_status" json:"incentive_status"`
	IncentiveDate   time.Time `bson:"incentive_date" json:"incentive_date"`

	IncentivePaymentMethod   string `bson:"incentive_payment_method" json:"incentive_payment_method"`
	IncentiveReferenceNumber string `bson:"incentive_reference_number" json:"incentive_reference_number"`

	IncentiveLogs []IncentiveLog `bson:"incentive_logs" json:"incentive_logs"`

	IsDeleted bool `bson:"is_deleted" json:"is_deleted"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type UpdateSalesOrder struct {
	DeliveryDate *time.Time `bson:"delivery_date,omitempty" json:"delivery_date,omitempty"`
	DownPayment  *float64   `bson:"down_payment,omitempty" json:"down_payment,omitempty"`
	LoanAmount   *float64   `bson:"loan_amount,omitempty" json:"loan_amount,omitempty"`
	BalanceAmount *float64  `bson:"balance_amount,omitempty" json:"balance_amount,omitempty"`
	Status       *string    `bson:"status,omitempty" json:"status,omitempty"`
	IncentiveAmount *float64 `bson:"incentive_amount,omitempty" json:"incentive_amount,omitempty"`
	IncentiveStatus *string  `bson:"incentive_status,omitempty" json:"incentive_status,omitempty"`
	IncentivePaymentMethod   *string `bson:"incentive_payment_method,omitempty" json:"incentive_payment_method,omitempty"`
	IncentiveReferenceNumber *string `bson:"incentive_reference_number,omitempty" json:"incentive_reference_number,omitempty"`
	DeliveryStatus *string `bson:"delivery_status,omitempty" json:"delivery_status,omitempty"`
	ActualDeliveryDate *time.Time `bson:"actual_delivery_date,omitempty" json:"actual_delivery_date,omitempty"`

	// Registration Details
	RegistrationStatus        *string    `bson:"registration_status,omitempty" json:"registration_status,omitempty"`
	RegistrationApplicationNo *string    `bson:"registration_application_no,omitempty" json:"registration_application_no,omitempty"`
	RegistrationTempNo        *string    `bson:"registration_temp_no,omitempty" json:"registration_temp_no,omitempty"`
	RegistrationAgentName     *string    `bson:"registration_agent_name,omitempty" json:"registration_agent_name,omitempty"`
	RegistrationAgentMobile   *string    `bson:"registration_agent_mobile,omitempty" json:"registration_agent_mobile,omitempty"`
	RegistrationDate          *time.Time `bson:"registration_date,omitempty" json:"registration_date,omitempty"`
	RegistrationDocumentUrl   *string    `bson:"registration_document_url,omitempty" json:"registration_document_url,omitempty"`
	VehicleNumber             *string    `bson:"vehicle_number,omitempty" json:"vehicle_number,omitempty"`
	RTO                       *string    `bson:"rto,omitempty" json:"rto,omitempty"`

	// Insurance Details
	InsuranceCompany   *string    `bson:"insurance_company,omitempty" json:"insurance_company,omitempty"`
	InsurancePolicyNo  *string    `bson:"insurance_policy_no,omitempty" json:"insurance_policy_no,omitempty"`
	InsuranceStartDate *time.Time `bson:"insurance_start_date,omitempty" json:"insurance_start_date,omitempty"`
	InsuranceEndDate   *time.Time `bson:"insurance_end_date,omitempty" json:"insurance_end_date,omitempty"`
}

func NewSalesOrder() *SalesOrder {
	now := time.Now().UTC()
	id := primitive.NewObjectID()

	return &SalesOrder{
		ID:             id,
		EntityID:       id.Hex(),
		SalesOrderCode: id.Hex()[18:24], // Running number only; prefix/suffix applied by service
		Status:         "Pending",
		DeliveryStatus: "Pending",
		RegistrationStatus: "Pending",
		IsDeleted:      false,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
}

func (s *SalesOrder) Bind(req *requests.CreateSalesOrderRequest) {
	s.CompanyID = req.CompanyID
	s.BranchID = req.BranchID
	s.CustomerID = req.CustomerID
	s.VehicleInventoryID = req.VehicleInventoryID
	s.SalespersonID = req.SalespersonID

	s.SaleDate = req.SaleDate
	s.DeliveryDate = req.DeliveryDate

	s.VehiclePrice = req.VehiclePrice
	s.RegistrationCharges = req.RegistrationCharges
	s.Insurance = req.Insurance
	s.Accessories = req.Accessories

	s.DiscountType = req.DiscountType
	s.DiscountValue = req.DiscountValue
	s.DiscountReason = req.DiscountReason
	s.DiscountAmount = req.DiscountAmount

	s.TotalAmount = req.TotalAmount
	s.DownPayment = req.DownPayment
	s.LoanAmount = req.LoanAmount
	s.BalanceAmount = req.BalanceAmount
	s.PaymentType = req.PaymentType
	s.PaymentMode = req.PaymentMode
	s.LoanStatus = req.LoanStatus
	s.UTRNumber = req.UTRNumber

	// Registration Details
	s.RegistrationStatus = req.RegistrationStatus
	s.RegistrationApplicationNo = req.RegistrationApplicationNo
	s.RegistrationTempNo = req.RegistrationTempNo
	s.RegistrationAgentName = req.RegistrationAgentName
	s.RegistrationAgentMobile = req.RegistrationAgentMobile
	s.RegistrationDate = req.RegistrationDate
	s.RegistrationDocumentUrl = req.RegistrationDocumentUrl
	s.VehicleNumber = req.VehicleNumber
	s.RTO = req.RTO

	// Insurance Details
	s.InsuranceCompany = req.InsuranceCompany
	s.InsurancePolicyNo = req.InsurancePolicyNo
	s.InsuranceStartDate = req.InsuranceStartDate
	s.InsuranceEndDate = req.InsuranceEndDate
}