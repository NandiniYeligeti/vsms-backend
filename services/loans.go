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

const LoanCollection = "loans"
const SalesOrderCollectionRef = "sales_orders"
const CustomerCollectionRef = "customers"

type LoanService interface {
	Create(ctx context.Context, companyCode string, req *requests.CreateLoanRequest) (*models.Loan, error)
	GetAll(ctx context.Context, companyCode string) ([]*models.Loan, error)
	GetByID(ctx context.Context, companyCode string, id string) (*models.Loan, error)
	Update(ctx context.Context, companyCode string, id string, req *requests.UpdateLoanRequest) (*models.Loan, error)
	Delete(ctx context.Context, companyCode string, id string) error
}

type loanService struct{}

func NewLoanService() LoanService {
	return &loanService{}
}

func (s *loanService) Create(
	ctx context.Context,
	companyCode string,
	req *requests.CreateLoanRequest,
) (*models.Loan, error) {

	db := storage.GetMongo()
	dbName := fmt.Sprintf("company_%s", companyCode)
	database := db.Database(dbName)

	loanCol := database.Collection(LoanCollection)
	soCol := database.Collection(SalesOrderCollectionRef)
	custCol := database.Collection(CustomerCollectionRef)

	// Verify Customer
	var customer models.Customer
	cFilter := bson.M{"entity_id": req.CustomerID, "is_deleted": bson.M{"$ne": true}}
	if oid, err := primitive.ObjectIDFromHex(req.CustomerID); err == nil {
		cFilter = bson.M{"_id": oid, "is_deleted": bson.M{"$ne": true}}
	}
	err := custCol.FindOne(ctx, cFilter).Decode(&customer)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("customer not found")
	}

	// Verify Sales Order
	var so models.SalesOrder
	soFilter := bson.M{"entity_id": req.SalesOrderID, "is_deleted": bson.M{"$ne": true}}
	if oid, err := primitive.ObjectIDFromHex(req.SalesOrderID); err == nil {
		soFilter = bson.M{"_id": oid, "is_deleted": bson.M{"$ne": true}}
	}
	err = soCol.FindOne(ctx, soFilter).Decode(&so)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("sales order not found")
	}

	loan := models.NewLoan()
	loan.CompanyID = req.CompanyID
	loan.BranchID = req.BranchID
	loan.CustomerID = req.CustomerID
	loan.CustomerName = customer.CustomerName
	loan.SalesOrderID = req.SalesOrderID
	loan.SalesOrderCode = so.SalesOrderCode
	loan.BankName = req.BankName
	loan.LoanAmount = req.LoanAmount
	loan.InterestRate = req.InterestRate
	loan.DurationMonths = req.DurationMonths
	loan.EMIAmount = req.EMIAmount
	if req.Status != "" {
		loan.Status = req.Status
	}
	loan.AccountNumber = req.AccountNumber
	loan.BankPerson = req.BankPerson
	loan.Mobile = req.Mobile

	_, err = loanCol.InsertOne(ctx, loan)
	if err != nil {
		return nil, err
	}

	return loan, nil
}

func (s *loanService) GetAll(
	ctx context.Context,
	companyCode string,
) ([]*models.Loan, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(LoanCollection)

	opts := options.Find().SetSort(bson.M{"created_at": -1})
	cursor, err := collection.Find(ctx, bson.M{"is_deleted": bson.M{"$ne": true}}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	result := []*models.Loan{}
	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *loanService) GetByID(
	ctx context.Context,
	companyCode string,
	id string,
) (*models.Loan, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(LoanCollection)

	filter := bson.M{"entity_id": id, "is_deleted": bson.M{"$ne": true}}
	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		filter = bson.M{"_id": oid, "is_deleted": bson.M{"$ne": true}}
	}

	var loan models.Loan
	err := collection.FindOne(ctx, filter).Decode(&loan)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("loan not found")
	}

	return &loan, err
}

func (s *loanService) Update(
	ctx context.Context,
	companyCode string,
	id string,
	req *requests.UpdateLoanRequest,
) (*models.Loan, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(LoanCollection)

	updateFields := bson.M{}
	if req.BankName != nil {
		updateFields["bank_name"] = *req.BankName
	}
	if req.LoanAmount != nil {
		updateFields["loan_amount"] = *req.LoanAmount
	}
	if req.InterestRate != nil {
		updateFields["interest_rate"] = *req.InterestRate
	}
	if req.DurationMonths != nil {
		updateFields["duration_months"] = *req.DurationMonths
	}
	if req.EMIAmount != nil {
		updateFields["emi_amount"] = *req.EMIAmount
	}
	if req.Status != nil {
		updateFields["status"] = *req.Status
	}
	if req.AccountNumber != nil {
		updateFields["account_number"] = *req.AccountNumber
	}
	if req.DisbursementDate != nil {
		updateFields["disbursement_date"] = *req.DisbursementDate
	}
	if req.BankPerson != nil {
		updateFields["bank_person"] = *req.BankPerson
	}
	if req.Mobile != nil {
		updateFields["mobile"] = *req.Mobile
	}
	updateFields["updated_at"] = time.Now()

	filter := bson.M{"entity_id": id}
	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		filter = bson.M{"_id": oid}
	}

	result, err := collection.UpdateOne(ctx, filter, bson.M{"$set": updateFields})
	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, errors.New("loan not found")
	}

	var updated models.Loan
	_ = collection.FindOne(ctx, filter).Decode(&updated)

	return &updated, nil
}

func (s *loanService) Delete(
	ctx context.Context,
	companyCode string,
	id string,
) error {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(LoanCollection)

	filter := bson.M{"entity_id": id}
	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		filter = bson.M{"_id": oid}
	}

	update := bson.M{"$set": bson.M{"is_deleted": true, "updated_at": time.Now()}}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("loan not found")
	}

	return nil
}
