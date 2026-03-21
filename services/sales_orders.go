package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"vehiclesales/models"
	"vehiclesales/requests"
	"vehiclesales/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const SalesOrderCollection = "sales_orders"
// const VehicleInventoryCollection = "vehicle_inventory"
// const CustomerCollection = "customers"

// ================== SERVICE INTERFACE ==================

type SalesOrderService interface {
	Create(ctx context.Context, companyCode string, req *requests.CreateSalesOrderRequest) (*models.SalesOrder, error)
	GetAll(ctx context.Context, companyCode string) ([]*models.SalesOrder, error)
	GetByID(ctx context.Context, companyCode string, id string) (*models.SalesOrder, error)
	Update(ctx context.Context, companyCode string, id string, req *requests.UpdateSalesOrderRequest) (*models.SalesOrder, error)
	Delete(ctx context.Context, companyCode string, id string) error
}

// ================== SERVICE STRUCT ==================

type salesOrderService struct{}

func NewSalesOrderService() SalesOrderService {
	return &salesOrderService{}
}

//
// ================== CREATE SALES ORDER ==================
//

func (s *salesOrderService) Create(
	ctx context.Context,
	companyCode string,
	req *requests.CreateSalesOrderRequest,
) (*models.SalesOrder, error) {

	db := storage.GetMongo()

	salesCollection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(SalesOrderCollection)
	vehicleCollection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(VehicleInventoryCollection)
	customerCollection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(CustomerCollection)

	// 1. Check vehicle exists
	var vehicle models.VehicleInventory
	vehicleFilter := bson.M{"entity_id": req.VehicleInventoryID, "is_deleted": false}
	if oid, err := primitive.ObjectIDFromHex(req.VehicleInventoryID); err == nil {
		vehicleFilter = bson.M{"_id": oid, "is_deleted": false}
	}

	err := vehicleCollection.FindOne(ctx, vehicleFilter).Decode(&vehicle)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("vehicle not found")
	}

	// 2. Prevent duplicate sold vehicle
	if vehicle.Status == "sold" {
		return nil, errors.New("vehicle already sold")
	}

	// 3. Fetch customer
	var customer models.Customer
	customerFilter := bson.M{"entity_id": req.CustomerID, "is_deleted": bson.M{"$ne": true}}
	if oid, err := primitive.ObjectIDFromHex(req.CustomerID); err == nil {
		customerFilter = bson.M{"_id": oid, "is_deleted": bson.M{"$ne": true}}
	}

	err = customerCollection.FindOne(ctx, customerFilter).Decode(&customer)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("customer not found")
	}

	// 4. Create sales order object
	order := models.NewSalesOrder()
	order.Bind(req)

	// 5. Populate denormalized details
	order.CustomerName = customer.CustomerName
	order.MobileNumber = customer.MobileNumber
	order.Email = customer.Email
	order.Address = customer.Address

	order.Brand = vehicle.Brand
	order.Model = vehicle.Model
	order.Variant = vehicle.Variant
	order.Color = vehicle.Color
	order.ChassisNumber = vehicle.ChassisNumber
	order.EngineNumber = vehicle.EngineNumber

	// 6. Auto calculation
	order.TotalAmount =
		req.VehiclePrice +
			req.RegistrationCharges +
			req.Insurance +
			req.Accessories

	order.BalanceAmount =
		order.TotalAmount -
			req.DownPayment -
			req.LoanAmount

	// 7. Insert order
	_, err = salesCollection.InsertOne(ctx, order)
	if err != nil {
		return nil, err
	}

	// 8. Mark vehicle sold
	_, _ = vehicleCollection.UpdateOne(
		ctx,
		vehicleFilter,
		bson.M{
			"$set": bson.M{
				"status":     "sold",
				"updated_at": time.Now(),
			},
		},
	)

	fmt.Println("Sales order created successfully")
	return order, nil
}

//
// ================== GET ALL SALES ==================
//

func (s *salesOrderService) GetAll(
	ctx context.Context,
	companyCode string,
) ([]*models.SalesOrder, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(SalesOrderCollection)

	cursor, err := collection.Find(ctx, bson.M{
		"is_deleted": bson.M{"$ne": true},
	})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var orders []*models.SalesOrder

	if err := cursor.All(ctx, &orders); err != nil {
		return nil, err
	}

	return orders, nil
}

//
// ================== GET BY ID ==================
//

func (s *salesOrderService) GetByID(
	ctx context.Context,
	companyCode string,
	id string,
) (*models.SalesOrder, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(SalesOrderCollection)

	var order models.SalesOrder

	filter := bson.M{"entity_id": id, "is_deleted": bson.M{"$ne": true}}
	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		filter = bson.M{"_id": oid, "is_deleted": bson.M{"$ne": true}}
	}

	err := collection.FindOne(ctx, filter).Decode(&order)

	if err == mongo.ErrNoDocuments {
		return nil, errors.New("sales order not found")
	}

	return &order, err
}

//
// ================== UPDATE SALES ORDER ==================
//

func (s *salesOrderService) Update(
	ctx context.Context,
	companyCode string,
	id string,
	req *requests.UpdateSalesOrderRequest,
) (*models.SalesOrder, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(SalesOrderCollection)

	updateFields := bson.M{}

	if req.DeliveryDate != nil {
		updateFields["delivery_date"] = *req.DeliveryDate
	}
	if req.DownPayment != nil {
		updateFields["down_payment"] = *req.DownPayment
	}
	if req.LoanAmount != nil {
		updateFields["loan_amount"] = *req.LoanAmount
	}
	if req.BalanceAmount != nil {
		updateFields["balance_amount"] = *req.BalanceAmount
	}
	if req.Status != nil {
		updateFields["status"] = *req.Status
	}

	updateFields["updated_at"] = time.Now()

	filter := bson.M{"entity_id": id}
	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		filter = bson.M{"_id": oid}
	}

	result, err := collection.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": updateFields},
	)

	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, errors.New("sales order not found")
	}

	var updated models.SalesOrder
	_ = collection.FindOne(ctx, filter).Decode(&updated)

	return &updated, nil
}

//
// ================== DELETE ==================
//

func (s *salesOrderService) Delete(
	ctx context.Context,
	companyCode string,
	id string,
) error {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(SalesOrderCollection)

	filter := bson.M{"entity_id": id}
	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		filter = bson.M{"_id": oid}
	}

	result, err := collection.UpdateOne(
		ctx,
		filter,
		bson.M{
			"$set": bson.M{
				"is_deleted": true,
				"updated_at": time.Now(),
			},
		},
	)

	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("sales order not found")
	}

	fmt.Println("Sales order deleted successfully")
	return nil
}