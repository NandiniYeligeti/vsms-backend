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
	"go.mongodb.org/mongo-driver/mongo/options"
)

const CompanyBankMasterCollection = "company_bank_master"

type CompanyBankMasterService interface {
	Create(ctx context.Context, companyCode string, req *requests.CreateCompanyBankMasterRequest) (*models.CompanyBankMaster, error)
	GetAll(ctx context.Context, companyCode string) ([]*models.CompanyBankMaster, error)
	Update(ctx context.Context, companyCode string, id string, req *requests.UpdateCompanyBankMasterRequest) (*models.CompanyBankMaster, error)
	Delete(ctx context.Context, companyCode string, id string) error
}

type companyBankMasterService struct{}

func NewCompanyBankMasterService() CompanyBankMasterService {
	return &companyBankMasterService{}
}

func (s *companyBankMasterService) Create(ctx context.Context, companyCode string, req *requests.CreateCompanyBankMasterRequest) (*models.CompanyBankMaster, error) {
	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(CompanyBankMasterCollection)

	bank := models.NewCompanyBankMaster()
	bank.CompanyID = companyCode
	bank.BankName = req.BankName
	bank.BranchName = req.BranchName
	bank.AccountNumber = req.AccountNumber
	bank.IsDefault = req.IsDefault

	if bank.IsDefault {
		// Reset other defaults
		collection.UpdateMany(ctx, bson.M{"company_id": companyCode}, bson.M{"$set": bson.M{"is_default": false}})
	}

	_, err := collection.InsertOne(ctx, bank)
	return bank, err
}

func (s *companyBankMasterService) GetAll(ctx context.Context, companyCode string) ([]*models.CompanyBankMaster, error) {
	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(CompanyBankMasterCollection)

	opts := options.Find().SetSort(bson.M{"created_at": -1})
	cursor, err := collection.Find(ctx, bson.M{"is_deleted": bson.M{"$ne": true}}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var banks []*models.CompanyBankMaster
	if err := cursor.All(ctx, &banks); err != nil {
		return nil, err
	}
	return banks, nil
}

func (s *companyBankMasterService) Update(ctx context.Context, companyCode string, id string, req *requests.UpdateCompanyBankMasterRequest) (*models.CompanyBankMaster, error) {
	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(CompanyBankMasterCollection)

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
	if req.AccountNumber != nil {
		update["account_number"] = *req.AccountNumber
	}
	if req.IsDefault != nil {
		update["is_default"] = *req.IsDefault
		if *req.IsDefault {
			// Reset other defaults
			collection.UpdateMany(ctx, bson.M{"company_id": companyCode}, bson.M{"$set": bson.M{"is_default": false}})
		}
	}
	update["updated_at"] = time.Now()

	_, err := collection.UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return nil, err
	}

	var updated models.CompanyBankMaster
	err = collection.FindOne(ctx, filter).Decode(&updated)
	return &updated, err
}

func (s *companyBankMasterService) Delete(ctx context.Context, companyCode string, id string) error {
	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(CompanyBankMasterCollection)

	filter := bson.M{"entity_id": id}
	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		filter = bson.M{"$or": []bson.M{{"entity_id": id}, {"_id": oid}}}
	}

	_, err := collection.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"is_deleted": true, "updated_at": time.Now()}})
	return err
}
