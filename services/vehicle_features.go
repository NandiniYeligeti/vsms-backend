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
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	VehicleTypeCollection      = "vehicle_types"
	VehicleCategoryCollection  = "vehicle_categories"
	VehicleAccessoryCollection = "vehicle_accessories"
)

type VehicleFeatureService interface {
	CreateType(ctx context.Context, companyCode string, req *requests.CreateVehicleTypeRequest) (*models.VehicleType, error)
	GetAllTypes(ctx context.Context, companyCode string) ([]*models.VehicleType, error)
	DeleteType(ctx context.Context, companyCode string, id string) error

	CreateCategory(ctx context.Context, companyCode string, req *requests.CreateVehicleCategoryRequest) (*models.VehicleCategory, error)
	GetAllCategories(ctx context.Context, companyCode string) ([]*models.VehicleCategory, error)
	DeleteCategory(ctx context.Context, companyCode string, id string) error

	CreateAccessory(ctx context.Context, companyCode string, req *requests.CreateVehicleAccessoryRequest) (*models.VehicleAccessory, error)
	GetAllAccessories(ctx context.Context, companyCode string) ([]*models.VehicleAccessory, error)
	DeleteAccessory(ctx context.Context, companyCode string, id string) error
}

type vehicleFeatureService struct{}

func NewVehicleFeatureService() VehicleFeatureService {
	return &vehicleFeatureService{}
}

func (s *vehicleFeatureService) CreateType(ctx context.Context, companyCode string, req *requests.CreateVehicleTypeRequest) (*models.VehicleType, error) {
	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(VehicleTypeCollection)
	if count, _ := collection.CountDocuments(ctx, bson.M{"name": req.Name}); count > 0 {
		return nil, errors.New("type already exists")
	}
	item := models.NewVehicleType(req)
	if _, err := collection.InsertOne(ctx, item); err != nil {
		return nil, err
	}
	return item, nil
}
func (s *vehicleFeatureService) GetAllTypes(ctx context.Context, companyCode string) ([]*models.VehicleType, error) {
	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(VehicleTypeCollection)
	opts := options.Find().SetSort(bson.M{"created_at": -1})
	cursor, err := collection.Find(ctx, bson.M{"is_deleted": bson.M{"$ne": true}}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var items []*models.VehicleType
	if err := cursor.All(ctx, &items); err != nil {
		return nil, err
	}
	return items, nil
}
func (s *vehicleFeatureService) DeleteType(ctx context.Context, companyCode string, id string) error {
	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(VehicleTypeCollection)
	filter := bson.M{"entity_id": id}
	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		filter = bson.M{"_id": oid}
	}
	_, err := collection.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"is_deleted": true, "updated_at": time.Now()}})
	return err
}

func (s *vehicleFeatureService) CreateCategory(ctx context.Context, companyCode string, req *requests.CreateVehicleCategoryRequest) (*models.VehicleCategory, error) {
	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(VehicleCategoryCollection)
	if count, _ := collection.CountDocuments(ctx, bson.M{"name": req.Name}); count > 0 {
		return nil, errors.New("category already exists")
	}
	item := models.NewVehicleCategory(req)
	if _, err := collection.InsertOne(ctx, item); err != nil {
		return nil, err
	}
	return item, nil
}
func (s *vehicleFeatureService) GetAllCategories(ctx context.Context, companyCode string) ([]*models.VehicleCategory, error) {
	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(VehicleCategoryCollection)
	opts := options.Find().SetSort(bson.M{"created_at": -1})
	cursor, err := collection.Find(ctx, bson.M{"is_deleted": bson.M{"$ne": true}}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var items []*models.VehicleCategory
	if err := cursor.All(ctx, &items); err != nil {
		return nil, err
	}
	return items, nil
}
func (s *vehicleFeatureService) DeleteCategory(ctx context.Context, companyCode string, id string) error {
	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(VehicleCategoryCollection)
	filter := bson.M{"entity_id": id}
	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		filter = bson.M{"_id": oid}
	}
	_, err := collection.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"is_deleted": true, "updated_at": time.Now()}})
	return err
}

func (s *vehicleFeatureService) CreateAccessory(ctx context.Context, companyCode string, req *requests.CreateVehicleAccessoryRequest) (*models.VehicleAccessory, error) {
	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(VehicleAccessoryCollection)
	if count, _ := collection.CountDocuments(ctx, bson.M{"name": req.Name}); count > 0 {
		return nil, errors.New("accessory already exists")
	}
	item := models.NewVehicleAccessory(req)
	if _, err := collection.InsertOne(ctx, item); err != nil {
		return nil, err
	}
	return item, nil
}
func (s *vehicleFeatureService) GetAllAccessories(ctx context.Context, companyCode string) ([]*models.VehicleAccessory, error) {
	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(VehicleAccessoryCollection)
	opts := options.Find().SetSort(bson.M{"created_at": -1})
	cursor, err := collection.Find(ctx, bson.M{"is_deleted": bson.M{"$ne": true}}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var items []*models.VehicleAccessory
	if err := cursor.All(ctx, &items); err != nil {
		return nil, err
	}
	return items, nil
}
func (s *vehicleFeatureService) DeleteAccessory(ctx context.Context, companyCode string, id string) error {
	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(VehicleAccessoryCollection)
	filter := bson.M{"entity_id": id}
	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		filter = bson.M{"_id": oid}
	}
	_, err := collection.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"is_deleted": true, "updated_at": time.Now()}})
	return err
}
