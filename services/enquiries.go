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

const EnquiryCollection = "enquiries"

type EnquiryService interface {
	Create(ctx context.Context, companyCode string, req *requests.CreateEnquiryRequest) (*models.Enquiry, error)
	GetAll(ctx context.Context, companyCode string) ([]*models.Enquiry, error)
	GetByID(ctx context.Context, companyCode string, id string) (*models.Enquiry, error)
	Update(ctx context.Context, companyCode string, id string, req *requests.UpdateEnquiryRequest) (*models.Enquiry, error)
	Delete(ctx context.Context, companyCode string, id string) error
}

type enquiryService struct{}

func NewEnquiryService() EnquiryService {
	return &enquiryService{}
}

func (s *enquiryService) Create(
	ctx context.Context,
	companyCode string,
	req *requests.CreateEnquiryRequest,
) (*models.Enquiry, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(EnquiryCollection)

	enquiry := models.NewEnquiry()
	enquiry.Bind(req)

	_, err := collection.InsertOne(ctx, enquiry)
	if err != nil {
		return nil, err
	}

	return enquiry, nil
}

func (s *enquiryService) GetAll(
	ctx context.Context,
	companyCode string,
) ([]*models.Enquiry, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(EnquiryCollection)

	opts := options.Find().SetSort(bson.M{"created_at": -1})
	cursor, err := collection.Find(ctx, bson.M{"is_deleted": bson.M{"$ne": true}}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var enquiries []*models.Enquiry
	if err := cursor.All(ctx, &enquiries); err != nil {
		return nil, err
	}

	return enquiries, nil
}

func (s *enquiryService) GetByID(
	ctx context.Context,
	companyCode string,
	id string,
) (*models.Enquiry, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(EnquiryCollection)

	filter := bson.M{"entity_id": id, "is_deleted": bson.M{"$ne": true}}
	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		filter = bson.M{"_id": oid, "is_deleted": bson.M{"$ne": true}}
	}

	var enquiry models.Enquiry
	err := collection.FindOne(ctx, filter).Decode(&enquiry)

	if err == mongo.ErrNoDocuments {
		return nil, errors.New("enquiry not found")
	}

	return &enquiry, err
}

func (s *enquiryService) Update(
	ctx context.Context,
	companyCode string,
	id string,
	req *requests.UpdateEnquiryRequest,
) (*models.Enquiry, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(EnquiryCollection)

	updateFields := bson.M{}

	if req.Name != nil {
		updateFields["name"] = *req.Name
	}
	if req.Mobile != nil {
		updateFields["mobile"] = *req.Mobile
	}
	if req.Email != nil {
		updateFields["email"] = *req.Email
	}
	if req.Vehicle != nil {
		updateFields["vehicle"] = *req.Vehicle
	}
	if req.Budget != nil {
		updateFields["budget"] = *req.Budget
	}
	if req.Salesperson != nil {
		updateFields["salesperson"] = *req.Salesperson
	}
	if req.FollowUps != nil {
		var fus []models.FollowUp
		for _, f := range *req.FollowUps {
			fus = append(fus, models.FollowUp{
				Date:     f.Date,
				Remark:   f.Remark,
				Status:   f.Status,
				NextDate: f.NextDate,
			})
		}
		updateFields["follow_ups"] = fus
	}
	if req.IsConverted != nil {
		updateFields["is_converted"] = *req.IsConverted
	}

	filter := bson.M{"entity_id": id}
	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		filter = bson.M{"_id": oid}
	}

	updateFields["updated_at"] = time.Now()

	result, err := collection.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": updateFields},
	)

	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, errors.New("enquiry not found")
	}

	var updated models.Enquiry
	_ = collection.FindOne(ctx, filter).Decode(&updated)

	return &updated, nil
}

func (s *enquiryService) Delete(
	ctx context.Context,
	companyCode string,
	id string,
) error {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(EnquiryCollection)

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
		return errors.New("enquiry not found")
	}

	return nil
}
