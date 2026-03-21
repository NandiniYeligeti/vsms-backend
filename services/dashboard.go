package services

import (
	"context"
	"fmt"
	"time"
	"vehiclesales/models"
	"vehiclesales/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DashboardService interface {
	GetStats(ctx context.Context, companyCode string) (*models.DashboardStats, error)
}

type dashboardService struct{}

func NewDashboardService() DashboardService {
	return &dashboardService{}
}

func (s *dashboardService) GetStats(ctx context.Context, companyCode string) (*models.DashboardStats, error) {
	db := storage.GetMongo()
	database := db.Database(fmt.Sprintf("company_%s", companyCode))

	stats := &models.DashboardStats{}

	// 1. Total Vehicles In Stock (Available for sale)
	stockCount, _ := database.Collection(VehicleInventoryCollection).CountDocuments(ctx, bson.M{
		"status":     bson.M{"$ne": "Sold"},
		"is_deleted": false,
	})
	stats.TotalVehiclesInStock = stockCount

	// 2. Total Vehicles Sold
	soldCount, _ := database.Collection(VehicleInventoryCollection).CountDocuments(ctx, bson.M{
		"status":     "Sold",
		"is_deleted": false,
	})
	stats.TotalVehiclesSold = soldCount

	// 3. Total Customers
	customerCount, _ := database.Collection(CustomerCollection).CountDocuments(ctx, bson.M{
		"is_deleted": false,
	})
	stats.TotalCustomers = customerCount

	// 4. Financial Metrics (Total Sales Revenue & Pending Payments)
	cursor, err := database.Collection(SalesOrderCollection).Aggregate(ctx, []bson.M{
		{
			"$match": bson.M{"is_deleted": false},
		},
		{
			"$group": bson.M{
				"_id":             nil,
				"totalRevenue":    bson.M{"$sum": "$total_amount"},
				"pendingPayments": bson.M{"$sum": "$balance_amount"},
			},
		},
	})
	if err == nil && cursor.Next(ctx) {
		var result struct {
			TotalRevenue    float64 `bson:"totalRevenue"`
			PendingPayments float64 `bson:"pendingPayments"`
		}
		if err := cursor.Decode(&result); err == nil {
			stats.TotalSalesRevenue = result.TotalRevenue
			stats.TotalPendingPayments = result.PendingPayments
		}
	}

	// 5. Total Pending Loans (Count where loan_amount > 0)
	loanCount, _ := database.Collection(SalesOrderCollection).CountDocuments(ctx, bson.M{
		"loan_amount": bson.M{"$gt": 0},
		"is_deleted":  false,
	})
	stats.TotalPendingLoans = loanCount

	// 6. Recent Sales (Last 5)
	findOptions := options.Find().SetSort(bson.M{"created_at": -1}).SetLimit(5)
	recentCursor, err := database.Collection(SalesOrderCollection).Find(ctx, bson.M{"is_deleted": false}, findOptions)
	if err == nil {
		var recent []*models.SalesOrder
		if err := recentCursor.All(ctx, &recent); err == nil {
			stats.RecentSales = recent
		}
	}

	// 7. Monthly Revenue (Last 6 months)
	sixMonthsAgo := time.Now().AddDate(0, -6, 0)
	monthlyCursor, err := database.Collection(SalesOrderCollection).Aggregate(ctx, []bson.M{
		{
			"$match": bson.M{
				"sale_date":  bson.M{"$gte": sixMonthsAgo},
				"is_deleted": false,
			},
		},
		{
			"$group": bson.M{
				"_id": bson.M{
					"month": bson.M{"$month": "$sale_date"},
					"year":  bson.M{"$year": "$sale_date"},
				},
				"revenue": bson.M{"$sum": "$total_amount"},
			},
		},
		{
			"$sort": bson.M{"_id.year": 1, "_id.month": 1},
		},
	})
	if err == nil {
		for monthlyCursor.Next(ctx) {
			var result struct {
				ID struct {
					Month int `bson:"month"`
					Year  int `bson:"year"`
				} `bson:"_id"`
				Revenue float64 `bson:"revenue"`
			}
			if err := monthlyCursor.Decode(&result); err == nil {
				stats.MonthlyRevenue = append(stats.MonthlyRevenue, models.MonthlyRevenue{
					Month:   fmt.Sprintf("%s", time.Month(result.ID.Month).String()),
					Revenue: result.Revenue,
				})
			}
		}
	}

	// 8. Sales By Model (Pie chart)
	modelCursor, err := database.Collection(SalesOrderCollection).Aggregate(ctx, []bson.M{
		{
			"$match": bson.M{"is_deleted": false},
		},
		{
			"$group": bson.M{
				"_id":   "$model",
				"count": bson.M{"$sum": 1},
			},
		},
		{
			"$limit": 10,
		},
	})
	if err == nil {
		for modelCursor.Next(ctx) {
			var result struct {
				Model string `bson:"_id"`
				Count int64  `bson:"count"`
			}
			if err := modelCursor.Decode(&result); err == nil {
				stats.SalesByModel = append(stats.SalesByModel, models.ModelSales{
					Model: result.Model,
					Count: result.Count,
				})
			}
		}
	}

	return stats, nil
}
