package models

import "time"

type DashboardStats struct {
	TotalVehiclesInStock int64   `json:"total_vehicles_in_stock"`
	TotalVehiclesSold    int64   `json:"total_vehicles_sold"`
	TotalCustomers       int64   `json:"total_customers"`
	TotalSalesRevenue    float64 `json:"total_sales_revenue"`
	TotalPendingPayments float64 `json:"total_pending_payments"`
	TotalPendingLoans    int64   `json:"total_pending_loans"`
	RegistrationPending  int64   `json:"registration_pending"`
	RegistrationInProcess int64  `json:"registration_in_process"`
	DeliveryPending      int64   `json:"delivery_pending"`
	InsurancePending     int64   `json:"insurance_pending"`
	IncentivePending     int64   `json:"incentive_pending"`
	IncentivePaid        int64   `json:"incentive_paid"`
	TodayFollowUps       int64   `json:"today_follow_ups"`
	FollowUpList         []*Enquiry `json:"follow_up_list"`

	MonthlyRevenue []MonthlyRevenue `json:"monthly_revenue"`
	SalesByModel   []ModelSales     `json:"sales_by_model"`
	RecentSales    []*SalesOrder     `json:"recent_sales"`
}

type MonthlyRevenue struct {
	Month   string  `json:"month"`
	Revenue float64 `json:"revenue"`
}

type ModelSales struct {
	Model string `json:"model"`
	Count int64  `json:"count"`
}

type RecentSale struct {
	OrderID      string    `json:"order_id"`
	CustomerName string    `json:"customer_name"`
	VehicleModel string    `json:"vehicle_model"`
	VIN          string    `json:"vin"`
	Amount       float64   `json:"amount"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
}
