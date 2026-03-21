package models

import "time"

type DashboardStats struct {
	TotalVehiclesInStock int64   `json:"total_vehicles_in_stock"`
	TotalVehiclesSold    int64   `json:"total_vehicles_sold"`
	TotalCustomers       int64   `json:"total_customers"`
	TotalSalesRevenue    float64 `json:"total_sales_revenue"`
	TotalPendingPayments float64 `json:"total_pending_payments"`
	TotalPendingLoans    int64   `json:"total_pending_loans"`

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
