package models

import (
	"time"
	"vehiclesales/requests"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

	SaleDate     time.Time `bson:"sale_date" json:"sale_date"`
	DeliveryDate time.Time `bson:"delivery_date" json:"delivery_date"`

	VehiclePrice         float64 `bson:"vehicle_price" json:"vehicle_price"`
	RegistrationCharges  float64 `bson:"registration_charges" json:"registration_charges"`
	Insurance            float64 `bson:"insurance" json:"insurance"`
	Accessories          float64 `bson:"accessories" json:"accessories"`

	TotalAmount   float64 `bson:"total_amount" json:"total_amount"`
	DownPayment   float64 `bson:"down_payment" json:"down_payment"`
	LoanAmount    float64 `bson:"loan_amount" json:"loan_amount"`
	BalanceAmount float64 `bson:"balance_amount" json:"balance_amount"`

	Status string `bson:"status" json:"status"`

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
}

func NewSalesOrder() *SalesOrder {
	now := time.Now().UTC()
	id := primitive.NewObjectID()

	return &SalesOrder{
		ID:             id,
		EntityID:       id.Hex(),
		SalesOrderCode: "SO" + id.Hex()[18:24],
		Status:         "Pending",
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

	s.TotalAmount = req.TotalAmount
	s.DownPayment = req.DownPayment
	s.LoanAmount = req.LoanAmount
	s.BalanceAmount = req.BalanceAmount
}