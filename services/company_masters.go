package services

import (
	"context"
	"fmt"
	"time"
	"vehiclesales/models"
	"vehiclesales/requests"
	"vehiclesales/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CompanyMastersCollection = "company_masters"
)

type CompanyMasterService interface {
	Create(ctx context.Context, companyCode string, req *requests.CreateCompanyMasterRequest) (*models.CompanyMaster, error)
	GetAll(ctx context.Context, companyCode string) ([]*models.CompanyMaster, error)
	Delete(ctx context.Context, companyCode string, id string) error
}

type companyMasterService struct{}

func NewCompanyMasterService() CompanyMasterService {
	return &companyMasterService{}
}

func (s *companyMasterService) Create(ctx context.Context, companyCode string, req *requests.CreateCompanyMasterRequest) (*models.CompanyMaster, error) {
	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(CompanyMastersCollection)
	item := models.NewCompanyMaster(companyCode, req.Type, req.Name)
	if _, err := collection.InsertOne(ctx, item); err != nil {
		return nil, err
	}
	return item, nil
}

func (s *companyMasterService) GetAll(ctx context.Context, companyCode string) ([]*models.CompanyMaster, error) {
	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(CompanyMastersCollection)
	cursor, err := collection.Find(ctx, bson.M{"is_deleted": bson.M{"$ne": true}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var items []*models.CompanyMaster
	if err := cursor.All(ctx, &items); err != nil {
		return nil, err
	}
	return items, nil
}

func (s *companyMasterService) Delete(ctx context.Context, companyCode string, id string) error {
	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(CompanyMastersCollection)
	filter := bson.M{"entity_id": id}
	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		filter = bson.M{"_id": oid}
	}
	_, err := collection.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"is_deleted": true, "updated_at": time.Now()}})
	return err
}
