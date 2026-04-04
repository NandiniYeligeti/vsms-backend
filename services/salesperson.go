package services

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"time"

	"vehiclesales/models"
	"vehiclesales/requests"
	"vehiclesales/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const SalespersonCollection = "salespersons"

// ================== PHONE VALIDATION ==================

var salespersonPhoneRegex = regexp.MustCompile(`^\d{10}$`)

func isValidSalespersonPhone(phone string) bool {
	return salespersonPhoneRegex.MatchString(phone)
}

// ================== SERVICE INTERFACE ==================

type SalespersonService interface {
	Create(ctx context.Context, companyCode string, req *requests.CreateSalespersonRequest) (*models.Salesperson, error)
	GetAll(ctx context.Context, companyCode string) ([]*models.Salesperson, error)
	GetByID(ctx context.Context, companyCode string, id string) (*models.Salesperson, error)
	Update(ctx context.Context, companyCode string, id string, req *requests.UpdateSalespersonRequest) (*models.Salesperson, error)
	Delete(ctx context.Context, companyCode string, id string) error
}

// ================== SERVICE STRUCT ==================

type salespersonService struct{}

func NewSalespersonService() SalespersonService {
	return &salespersonService{}
}

// ================== CREATE SALESPERSON ==================

func (s *salespersonService) Create(
	ctx context.Context,
	companyCode string,
	req *requests.CreateSalespersonRequest,
) (*models.Salesperson, error) {

	if !isValidSalespersonPhone(req.MobileNumber) {
		return nil, errors.New("phone number must be exactly 10 digits")
	}

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(SalespersonCollection)

	if count, _ := collection.CountDocuments(ctx, bson.M{"mobile_number": req.MobileNumber}); count > 0 {
		return nil, errors.New("phone already exists")
	}

	salesperson := models.NewSalesperson()
	salesperson.Bind(req)

	_, err := collection.InsertOne(ctx, salesperson)
	if err != nil {
		return nil, err
	}

	fmt.Println("Salesperson created successfully")
	return salesperson, nil
}

// ================== GET ALL SALESPERSONS ==================

func (s *salespersonService) GetAll(
	ctx context.Context,
	companyCode string,
) ([]*models.Salesperson, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(SalespersonCollection)

	cursor, err := collection.Find(ctx, bson.M{"is_deleted": bson.M{"$ne": true}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var salespersons []*models.Salesperson
	if err := cursor.All(ctx, &salespersons); err != nil {
		return nil, err
	}

	return salespersons, nil
}

// ================== GET SALESPERSON BY ID ==================

func (s *salespersonService) GetByID(
	ctx context.Context,
	companyCode string,
	id string,
) (*models.Salesperson, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(SalespersonCollection)

	filter := bson.M{"entity_id": id, "is_deleted": bson.M{"$ne": true}}
	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		filter = bson.M{"_id": oid, "is_deleted": bson.M{"$ne": true}}
	}

	var salesperson models.Salesperson
	err := collection.FindOne(ctx, filter).Decode(&salesperson)

	if err == mongo.ErrNoDocuments {
		return nil, errors.New("salesperson not found")
	}

	return &salesperson, err
}

// ================== UPDATE SALESPERSON ==================

func (s *salespersonService) Update(
	ctx context.Context,
	companyCode string,
	id string,
	req *requests.UpdateSalespersonRequest,
) (*models.Salesperson, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(SalespersonCollection)

	updateFields := bson.M{}

	if req.FullName != nil {
		updateFields["full_name"] = *req.FullName
	}

	if req.MobileNumber != nil {
		if !isValidSalespersonPhone(*req.MobileNumber) {
			return nil, errors.New("phone number must be exactly 10 digits")
		}
		updateFields["mobile_number"] = *req.MobileNumber
	}

	if req.Email != nil {
		updateFields["email"] = *req.Email
	}

	if req.BranchID != nil {
		updateFields["branch_id"] = *req.BranchID
	}

	if req.Showroom != nil {
		updateFields["showroom"] = *req.Showroom
	}

	if req.Branch != nil {
		updateFields["branch"] = *req.Branch
	}

	if req.Area != nil {
		updateFields["area"] = *req.Area
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
		return nil, errors.New("salesperson not found")
	}

	var updated models.Salesperson
	_ = collection.FindOne(ctx, filter).Decode(&updated)

	return &updated, nil
}

// ================== DELETE SALESPERSON ==================

func (s *salespersonService) Delete(
	ctx context.Context,
	companyCode string,
	id string,
) error {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(SalespersonCollection)

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
		return errors.New("salesperson not found")
	}

	fmt.Println("Salesperson soft deleted successfully")
	return nil
}