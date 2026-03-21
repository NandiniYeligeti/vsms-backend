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

const VehicleInventoryCollection = "vehicle_inventory"
const VehicleModelCollectionRef = "vehicle_models"

// ================== SERVICE INTERFACE ==================

type VehicleInventoryService interface {
	Create(ctx context.Context, companyCode string, req *requests.CreateVehicleInventoryRequest) (*models.VehicleInventory, error)
	GetAll(ctx context.Context, companyCode string) ([]*models.VehicleInventory, error)
	GetByID(ctx context.Context, companyCode string, id string) (*models.VehicleInventory, error)
	Update(ctx context.Context, companyCode string, id string, req *requests.UpdateVehicleInventoryRequest) (*models.VehicleInventory, error)
	Delete(ctx context.Context, companyCode string, id string) error
}

// ================== SERVICE STRUCT ==================

type vehicleInventoryService struct{}

func NewVehicleInventoryService() VehicleInventoryService {
	return &vehicleInventoryService{}
}

// ================== CREATE INVENTORY ==================

func (s *vehicleInventoryService) Create(
	ctx context.Context,
	companyCode string,
	req *requests.CreateVehicleInventoryRequest,
) (*models.VehicleInventory, error) {

	db := storage.GetMongo()
	database := db.Database(fmt.Sprintf("company_%s", companyCode))

	inventoryCollection := database.Collection(VehicleInventoryCollection)
	modelCollection := database.Collection(VehicleModelCollectionRef)


	modelFilter := bson.M{"entity_id": req.VehicleModelID, "is_deleted": bson.M{"$ne": true}}
	if oid, err := primitive.ObjectIDFromHex(req.VehicleModelID); err == nil {
		modelFilter = bson.M{"_id": oid, "is_deleted": bson.M{"$ne": true}}
	}

	var vehicleModel models.VehicleModel
	err := modelCollection.FindOne(ctx, modelFilter).Decode(&vehicleModel)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("vehicle model not found")
	}
	if err != nil {
		return nil, err
	}

	inventory := models.NewVehicleInventory()
	inventory.Bind(req)

	inventory.Brand = vehicleModel.Brand
	inventory.Model = vehicleModel.Model
	inventory.Variant = vehicleModel.Variant
	inventory.FuelType = vehicleModel.FuelType
	inventory.BasePrice = vehicleModel.BasePrice

	_, err = inventoryCollection.InsertOne(ctx, inventory)
	if err != nil {
		return nil, err
	}

	fmt.Println("Vehicle inventory created successfully")
	return inventory, nil
}

// ================== GET ALL INVENTORY ==================

func (s *vehicleInventoryService) GetAll(
	ctx context.Context,
	companyCode string,
) ([]*models.VehicleInventory, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(VehicleInventoryCollection)

	cursor, err := collection.Find(ctx, bson.M{"is_deleted": bson.M{"$ne": true}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var inventory []*models.VehicleInventory
	if err := cursor.All(ctx, &inventory); err != nil {
		return nil, err
	}

	return inventory, nil
}

// ================== GET INVENTORY BY ID ==================

func (s *vehicleInventoryService) GetByID(
	ctx context.Context,
	companyCode string,
	id string,
) (*models.VehicleInventory, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(VehicleInventoryCollection)

	filter := bson.M{"entity_id": id, "is_deleted": bson.M{"$ne": true}}
	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		filter = bson.M{"_id": oid, "is_deleted": bson.M{"$ne": true}}
	}

	var inventory models.VehicleInventory
	err := collection.FindOne(ctx, filter).Decode(&inventory)

	if err == mongo.ErrNoDocuments {
		return nil, errors.New("vehicle inventory not found")
	}

	return &inventory, err
}

// ================== UPDATE INVENTORY ==================

func (s *vehicleInventoryService) Update(
	ctx context.Context,
	companyCode string,
	id string,
	req *requests.UpdateVehicleInventoryRequest,
) (*models.VehicleInventory, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(VehicleInventoryCollection)

	updateFields := bson.M{}

	if req.Color != nil {
		updateFields["color"] = *req.Color
	}
	if req.ChassisNumber != nil {
		updateFields["chassis_number"] = *req.ChassisNumber
	}
	if req.EngineNumber != nil {
		updateFields["engine_number"] = *req.EngineNumber
	}
	if req.PurchaseDate != nil {
		updateFields["purchase_date"] = *req.PurchaseDate
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
		return nil, errors.New("vehicle inventory not found")
	}

	var updated models.VehicleInventory
	_ = collection.FindOne(ctx, filter).Decode(&updated)

	return &updated, nil
}

// ================== DELETE INVENTORY ==================

func (s *vehicleInventoryService) Delete(
	ctx context.Context,
	companyCode string,
	id string,
) error {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(VehicleInventoryCollection)

	filter := bson.M{"entity_id": id}
	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		filter = bson.M{"_id": oid}
	}

	update := bson.M{
		"$set": bson.M{
			"is_deleted": true,
			"updated_at": time.Now(),
		},
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("vehicle inventory not found")
	}

	fmt.Println("Vehicle inventory soft deleted successfully")
	return nil
}