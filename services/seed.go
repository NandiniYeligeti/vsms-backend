package services

import (
	"context"
	"fmt"
	"time"

	"vehiclesales/models"
	"vehiclesales/storage"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SeedDatabase(ctx context.Context, companyCode string) error {
	db := storage.GetMongo()
	database := db.Database(fmt.Sprintf("company_%s", companyCode))

	// 1. Seed Customer
	customer := models.NewCustomer()
	customer.CustomerName = "Rajesh Kumar"
	customer.MobileNumber = "9876543210"
	customer.Email = "rajesh@email.com"
	customer.City = "Mumbai"
	_, err := database.Collection(CustomerCollection).InsertOne(ctx, customer)
	if err != nil {
		return err
	}

	// 2. Seed Vehicle Models
	vm1 := models.NewVehicleModel()
	vm1.Brand = "Hyundai"
	vm1.Model = "Creta"
	vm1.Variant = "SX(O)"
	vm1.BasePrice = 1799000
	_, err = database.Collection("vehicle_models").InsertOne(ctx, vm1)
	if err != nil {
		return err
	}

	vm2 := models.NewVehicleModel()
	vm2.Brand = "Tata"
	vm2.Model = "Nexon EV"
	vm2.Variant = "Max LR"
	vm2.BasePrice = 1999000
	_, err = database.Collection("vehicle_models").InsertOne(ctx, vm2)
	if err != nil {
		return err
	}

	// 3. Seed Vehicle Inventory
	vi1 := models.NewVehicleInventory()
	vi1.VehicleModelID = vm1.EntityID
	vi1.Brand = vm1.Brand
	vi1.Model = vm1.Model
	vi1.Color = "Phantom Black"
	vi1.Status = "Sold"
	_, err = database.Collection(VehicleInventoryCollection).InsertOne(ctx, vi1)
	if err != nil {
		return err
	}

	vi2 := models.NewVehicleInventory()
	vi2.VehicleModelID = vm2.EntityID
	vi2.Brand = vm2.Brand
	vi2.Model = vm2.Model
	vi2.Color = "Teal Blue"
	vi2.Status = "Available"
	_, err = database.Collection(VehicleInventoryCollection).InsertOne(ctx, vi2)
	if err != nil {
		return err
	}

	vi3 := models.NewVehicleInventory()
	vi3.VehicleModelID = vm1.EntityID
	vi3.Brand = vm1.Brand
	vi3.Model = vm1.Model
	vi3.Color = "Pearl White"
	vi3.Status = "Available"
	_, err = database.Collection(VehicleInventoryCollection).InsertOne(ctx, vi3)
	if err != nil {
		return err
	}

	// 4. Seed Sales Order
	so1 := models.NewSalesOrder()
	so1.CustomerID = customer.EntityID
	so1.CustomerName = customer.CustomerName
	so1.VehicleInventoryID = vi1.EntityID
	so1.Brand = vi1.Brand
	so1.Model = vi1.Model
	so1.ChassisNumber = vi1.ChassisNumber
	so1.TotalAmount = 1931000
	so1.BalanceAmount = 231000
	so1.LoanAmount = 1200000
	so1.Status = "Delivered"
	// Set SaleDate to last month for the chart
	so1.SaleDate = time.Now().AddDate(0, -1, -5)
	
	// manually override ID and createdAt so we can set historic dates reliably for charts
	soid := primitive.NewObjectID()
	so1.ID = soid
	so1.EntityID = soid.Hex()
	so1.SalesOrderCode = "SO" + soid.Hex()[18:24]
	so1.CreatedAt = time.Now().AddDate(0, -1, -5)
	_, err = database.Collection(SalesOrderCollection).InsertOne(ctx, so1)
	if err != nil {
		return err
	}

	so2 := models.NewSalesOrder()
	so2.CustomerID = customer.EntityID
	so2.CustomerName = customer.CustomerName
	so2.TotalAmount = 2100000
	so2.BalanceAmount = 0
	so2.Brand = vm2.Brand
	so2.Model = vm2.Model
	so2.Status = "Confirmed"
	so2.SaleDate = time.Now().AddDate(0, -3, -10)
	
	soid2 := primitive.NewObjectID()
	so2.ID = soid2
	so2.EntityID = soid2.Hex()
	so2.SalesOrderCode = "SO" + soid2.Hex()[18:24]
	so2.CreatedAt = time.Now().AddDate(0, -3, -10)
	_, err = database.Collection(SalesOrderCollection).InsertOne(ctx, so2)
	if err != nil {
		return err
	}
	
	so3 := models.NewSalesOrder()
	so3.CustomerID = customer.EntityID
	so3.CustomerName = customer.CustomerName
	so3.TotalAmount = 900000
	so3.BalanceAmount = 100000
	so3.Brand = vm1.Brand
	so3.Model = vm1.Model
	so3.Status = "Pending"
	so3.SaleDate = time.Now().AddDate(0, 0, -2) // current month
	
	soid3 := primitive.NewObjectID()
	so3.ID = soid3
	so3.EntityID = soid3.Hex()
	so3.SalesOrderCode = "SO" + soid3.Hex()[18:24]
	so3.CreatedAt = time.Now().AddDate(0, 0, -2)
	_, err = database.Collection(SalesOrderCollection).InsertOne(ctx, so3)
	if err != nil {
		return err
	}

	return nil
}
