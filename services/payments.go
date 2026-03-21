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

const PaymentCollection = "payments"
// const SalesOrderCollection = "sales_orders"

// ================== SERVICE INTERFACE ==================

type PaymentService interface {
	Create(ctx context.Context, companyCode string, req *requests.CreatePaymentRequest) (*models.Payment, error)
	GetAll(ctx context.Context, companyCode string) ([]*models.Payment, error)
	GetByID(ctx context.Context, companyCode string, id string) (*models.Payment, error)
	Update(ctx context.Context, companyCode string, id string, req *requests.UpdatePaymentRequest) (*models.Payment, error)
	Delete(ctx context.Context, companyCode string, id string) error
}

// ================== SERVICE STRUCT ==================

type paymentService struct{}

func NewPaymentService() PaymentService {
	return &paymentService{}
}

// ================== CREATE PAYMENT ==================

func (s *paymentService) Create(
	ctx context.Context,
	companyCode string,
	req *requests.CreatePaymentRequest,
) (*models.Payment, error) {

	db := storage.GetMongo()

	paymentCollection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(PaymentCollection)
	salesCollection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(SalesOrderCollection)

	// fetch sales order
	var order models.SalesOrder
	err := salesCollection.FindOne(ctx, bson.M{
		"entity_id": req.SalesOrderID,
		"is_deleted": bson.M{"$ne": true},
	}).Decode(&order)

	if err == mongo.ErrNoDocuments {
		return nil, errors.New("sales order not found")
	}

	// prevent overpayment
	if req.PaymentAmount > order.BalanceAmount {
		return nil, errors.New("payment exceeds balance amount")
	}

	// create payment
	payment := models.NewPayment()
	payment.Bind(req)

	_, err = paymentCollection.InsertOne(ctx, payment)
	if err != nil {
		return nil, err
	}

	// update sales order balance
	newBalance := order.BalanceAmount - req.PaymentAmount

	updateFields := bson.M{
		"balance_amount": newBalance,
		"updated_at":     time.Now(),
	}

	if newBalance == 0 {
		updateFields["status"] = "fully_paid"
	}

	_, _ = salesCollection.UpdateOne(
		ctx,
		bson.M{"entity_id": req.SalesOrderID},
		bson.M{
			"$set": updateFields,
		},
	)

	fmt.Println("Payment recorded successfully")
	return payment, nil
}

// ================== GET ALL PAYMENTS ==================

func (s *paymentService) GetAll(
	ctx context.Context,
	companyCode string,
) ([]*models.Payment, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(PaymentCollection)

	cursor, err := collection.Find(ctx, bson.M{
		"is_deleted": bson.M{"$ne": true},
	})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var payments []*models.Payment

	if err := cursor.All(ctx, &payments); err != nil {
		return nil, err
	}

	return payments, nil
}

// ================== GET PAYMENT BY ID ==================

func (s *paymentService) GetByID(
	ctx context.Context,
	companyCode string,
	id string,
) (*models.Payment, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(PaymentCollection)

	var payment models.Payment

	err := collection.FindOne(ctx, bson.M{
		"entity_id": id,
		"is_deleted": bson.M{"$ne": true},
	}).Decode(&payment)

	if err == mongo.ErrNoDocuments {
		return nil, errors.New("payment not found")
	}

	return &payment, err
}

func (s *paymentService) Update(
	ctx context.Context,
	companyCode string,
	id string,
	req *requests.UpdatePaymentRequest,
) (*models.Payment, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(PaymentCollection)

	updateFields := bson.M{}

	if req.PaymentDate != nil {
		updateFields["payment_date"] = *req.PaymentDate
	}
	if req.PaymentAmount != nil {
		updateFields["payment_amount"] = *req.PaymentAmount
	}
	if req.PaymentMode != nil {
		updateFields["payment_mode"] = *req.PaymentMode
	}
	if req.PaymentType != nil {
		updateFields["payment_type"] = *req.PaymentType
	}
	if req.ReferenceNumber != nil {
		updateFields["reference_number"] = *req.ReferenceNumber
	}
	if req.BankName != nil {
		updateFields["bank_name"] = *req.BankName
	}
	if req.CollectedBy != nil {
		updateFields["collected_by"] = *req.CollectedBy
	}
	if req.Remarks != nil {
		updateFields["remarks"] = *req.Remarks
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
		return nil, errors.New("payment not found")
	}

	var updated models.Payment
	_ = collection.FindOne(ctx, filter).Decode(&updated)

	return &updated, nil
}

// ================== DELETE PAYMENT ==================

func (s *paymentService) Delete(
	ctx context.Context,
	companyCode string,
	id string,
) error {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(PaymentCollection)

	result, err := collection.UpdateOne(
		ctx,
		bson.M{"entity_id": id},
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
		return errors.New("payment not found")
	}

	fmt.Println("Payment deleted successfully")
	return nil
}