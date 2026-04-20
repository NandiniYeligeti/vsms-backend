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
	"go.mongodb.org/mongo-driver/mongo/options"
)

const VehicleModelCollection = "vehicle_models"

// ================== SERVICE INTERFACE ==================

type VehicleModelService interface {
	Create(ctx context.Context, companyCode string, req *requests.CreateVehicleModelRequest) (*models.VehicleModel, error)
	GetAll(ctx context.Context, companyCode string) ([]*models.VehicleModel, error)
	GetByID(ctx context.Context, companyCode string, id string) (*models.VehicleModel, error)
	Update(ctx context.Context, companyCode string, id string, req *requests.UpdateVehicleModelRequest) (*models.VehicleModel, error)
	Delete(ctx context.Context, companyCode string, id string) error
}

// ================== SERVICE STRUCT ==================

type vehicleModelService struct{}

func NewVehicleModelService() VehicleModelService {
	return &vehicleModelService{}
}

// ================== CREATE VEHICLE MODEL ==================

func (s *vehicleModelService) Create(
	ctx context.Context,
	companyCode string,
	req *requests.CreateVehicleModelRequest,
) (*models.VehicleModel, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(VehicleModelCollection)

	duplicateFilter := bson.M{
		"brand":      req.Brand,
		"model":      req.Model,
		"variant":    req.Variant,
		"is_deleted": bson.M{"$ne": true},
	}

	if count, _ := collection.CountDocuments(ctx, duplicateFilter); count > 0 {
		return nil, errors.New("vehicle model already exists")
	}

	vehicle := models.NewVehicleModel()
	vehicle.Bind(req)

	_, err := collection.InsertOne(ctx, vehicle)
	if err != nil {
		return nil, err
	}

	fmt.Println("Vehicle model created successfully")
	return vehicle, nil
}

// ================== GET ALL VEHICLE MODELS ==================

func (s *vehicleModelService) GetAll(
	ctx context.Context,
	companyCode string,
) ([]*models.VehicleModel, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(VehicleModelCollection)

	opts := options.Find().SetSort(bson.M{"created_at": -1})
	cursor, err := collection.Find(ctx, bson.M{"is_deleted": bson.M{"$ne": true}}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var vehicles []*models.VehicleModel
	if err := cursor.All(ctx, &vehicles); err != nil {
		return nil, err
	}

	return vehicles, nil
}

// ================== GET VEHICLE MODEL BY ID ==================

func (s *vehicleModelService) GetByID(
	ctx context.Context,
	companyCode string,
	id string,
) (*models.VehicleModel, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(VehicleModelCollection)

	filter := bson.M{"entity_id": id, "is_deleted": bson.M{"$ne": true}}
	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		filter = bson.M{"_id": oid, "is_deleted": bson.M{"$ne": true}}
	}

	var vehicle models.VehicleModel
	err := collection.FindOne(ctx, filter).Decode(&vehicle)

	if err == mongo.ErrNoDocuments {
		return nil, errors.New("vehicle model not found")
	}

	return &vehicle, err
}

// ================== UPDATE VEHICLE MODEL ==================

func (s *vehicleModelService) Update(
	ctx context.Context,
	companyCode string,
	id string,
	req *requests.UpdateVehicleModelRequest,
) (*models.VehicleModel, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(VehicleModelCollection)

	updateFields := bson.M{}

	if req.Brand != nil {
		updateFields["brand"] = *req.Brand
	}
	if req.Model != nil {
		updateFields["model"] = *req.Model
	}
	if req.Variant != nil {
		updateFields["variant"] = *req.Variant
	}
	if req.FuelType != nil {
		updateFields["fuel_type"] = *req.FuelType
	}
	if req.BasePrice != nil {
		updateFields["base_price"] = *req.BasePrice
	}
	if req.TypeID != nil {
		updateFields["type_id"] = *req.TypeID
	}
	if req.CategoryID != nil {
		updateFields["category_id"] = *req.CategoryID
	}
	if req.Colors != nil {
		updateFields["colors"] = *req.Colors
	}
	if req.IncentiveType != nil {
		updateFields["incentive_type"] = *req.IncentiveType
	}
	if req.IncentiveValue != nil {
		updateFields["incentive_value"] = *req.IncentiveValue
	}
	if req.ColorCount != nil {
		updateFields["color_count"] = *req.ColorCount
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
		return nil, errors.New("vehicle model not found")
	}

	var updated models.VehicleModel
	_ = collection.FindOne(ctx, filter).Decode(&updated)

	return &updated, nil
}

// ================== DELETE VEHICLE MODEL ==================

func (s *vehicleModelService) Delete(
	ctx context.Context,
	companyCode string,
	id string,
) error {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(VehicleModelCollection)

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
		return errors.New("vehicle model not found")
	}

	fmt.Println("Vehicle model soft deleted successfully")
	return nil
}