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

const BankMasterCollection = "bank_master"

type BankMasterService interface {
	Create(ctx context.Context, companyCode string, req *requests.CreateBankMasterRequest) (*models.BankMaster, error)
	GetAll(ctx context.Context, companyCode string) ([]*models.BankMaster, error)
	Update(ctx context.Context, companyCode string, id string, req *requests.UpdateBankMasterRequest) (*models.BankMaster, error)
	Delete(ctx context.Context, companyCode string, id string) error
}

type bankMasterService struct{}

func NewBankMasterService() BankMasterService {
	return &bankMasterService{}
}

func (s *bankMasterService) Create(ctx context.Context, companyCode string, req *requests.CreateBankMasterRequest) (*models.BankMaster, error) {
	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(BankMasterCollection)

	bank := models.NewBankMaster()
	bank.CompanyID = req.CompanyID
	bank.BankName = req.BankName
	bank.BranchName = req.BranchName
	bank.ContactPerson = req.ContactPerson
	bank.ContactNumber = req.ContactNumber

	_, err := collection.InsertOne(ctx, bank)
	return bank, err
}

func (s *bankMasterService) GetAll(ctx context.Context, companyCode string) ([]*models.BankMaster, error) {
	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(BankMasterCollection)

	cursor, err := collection.Find(ctx, bson.M{"is_deleted": bson.M{"$ne": true}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var banks []*models.BankMaster
	if err := cursor.All(ctx, &banks); err != nil {
		return nil, err
	}
	return banks, nil
}

func (s *bankMasterService) Update(ctx context.Context, companyCode string, id string, req *requests.UpdateBankMasterRequest) (*models.BankMaster, error) {
	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(BankMasterCollection)

	filter := bson.M{"entity_id": id}
	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		filter = bson.M{"$or": []bson.M{{"entity_id": id}, {"_id": oid}}}
	}

	update := bson.M{}
	if req.BankName != nil {
		update["bank_name"] = *req.BankName
	}
	if req.BranchName != nil {
		update["branch_name"] = *req.BranchName
	}
	if req.ContactPerson != nil {
		update["contact_person"] = *req.ContactPerson
	}
	if req.ContactNumber != nil {
		update["contact_number"] = *req.ContactNumber
	}
	update["updated_at"] = time.Now()

	_, err := collection.UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return nil, err
	}

	var updated models.BankMaster
	err = collection.FindOne(ctx, filter).Decode(&updated)
	return &updated, err
}

func (s *bankMasterService) Delete(ctx context.Context, companyCode string, id string) error {
	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(BankMasterCollection)

	filter := bson.M{"entity_id": id}
	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		filter = bson.M{"$or": []bson.M{{"entity_id": id}, {"_id": oid}}}
	}

	_, err := collection.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"is_deleted": true, "updated_at": time.Now()}})
	return err
}
