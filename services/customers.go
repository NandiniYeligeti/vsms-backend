package services

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"sort"
	"time"

	"vehiclesales/models"
	"vehiclesales/requests"
	"vehiclesales/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const CustomerCollection = "customers"

// ================== PHONE VALIDATION ==================

var customerPhoneRegex = regexp.MustCompile(`^\d{10}$`)

func isValidCustomerPhone(phone string) bool {
	return customerPhoneRegex.MatchString(phone)
}

// ================== SERVICE INTERFACE ==================

type CustomerService interface {
	Create(ctx context.Context, companyCode string, req *requests.CreateCustomerRequest) (*models.Customer, error)
	GetAll(ctx context.Context, companyCode string) ([]*models.Customer, error)
	GetByID(ctx context.Context, companyCode string, id string) (*models.Customer, error)
	Update(ctx context.Context, companyCode string, id string, req *requests.UpdateCustomerRequest) (*models.Customer, error)
	Delete(ctx context.Context, companyCode string, id string) error
	GetLedger(ctx context.Context, companyCode string, customerId string) ([]*models.LedgerEntry, error)
}

// ================== SERVICE STRUCT ==================

type customerService struct{}

func NewCustomerService() CustomerService {
	return &customerService{}
}

// ================== CREATE CUSTOMER ==================

func (s *customerService) Create(
	ctx context.Context,
	companyCode string,
	req *requests.CreateCustomerRequest,
) (*models.Customer, error) {

	if !isValidCustomerPhone(req.MobileNumber) {
		return nil, errors.New("phone number must be exactly 10 digits")
	}

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(CustomerCollection)

	customer := models.NewCustomer()
	customer.Bind(req)

	_, err := collection.InsertOne(ctx, customer)
	if err != nil {
		return nil, err
	}

	fmt.Println("Customer created successfully")
	return customer, nil
}

// ================== GET ALL CUSTOMERS ==================

func (s *customerService) GetAll(
	ctx context.Context,
	companyCode string,
) ([]*models.Customer, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(CustomerCollection)

	cursor, err := collection.Find(ctx, bson.M{"is_deleted": bson.M{"$ne": true}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var customers []*models.Customer
	if err := cursor.All(ctx, &customers); err != nil {
		return nil, err
	}

	return customers, nil
}

// ================== GET CUSTOMER BY ID ==================

func (s *customerService) GetByID(
	ctx context.Context,
	companyCode string,
	id string,
) (*models.Customer, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(CustomerCollection)

	filter := bson.M{"entity_id": id, "is_deleted": bson.M{"$ne": true}}
	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		filter = bson.M{"_id": oid, "is_deleted": bson.M{"$ne": true}}
	}

	var customer models.Customer
	err := collection.FindOne(ctx, filter).Decode(&customer)

	if err == mongo.ErrNoDocuments {
		return nil, errors.New("customer not found")
	}

	return &customer, err
}

// ================== UPDATE CUSTOMER ==================

func (s *customerService) Update(
	ctx context.Context,
	companyCode string,
	id string,
	req *requests.UpdateCustomerRequest,
) (*models.Customer, error) {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(CustomerCollection)

	updateFields := bson.M{}

	if req.CustomerName != nil {
		updateFields["customer_name"] = *req.CustomerName
	}
	if req.MobileNumber != nil {
		if !isValidCustomerPhone(*req.MobileNumber) {
			return nil, errors.New("phone number must be exactly 10 digits")
		}
		updateFields["mobile_number"] = *req.MobileNumber
	}
	if req.Email != nil {
		updateFields["email"] = *req.Email
	}
	if req.Address != nil {
		updateFields["address"] = *req.Address
	}
	if req.City != nil {
		updateFields["city"] = *req.City
	}
	if req.State != nil {
		updateFields["state"] = *req.State
	}
	if req.Pincode != nil {
		updateFields["pincode"] = *req.Pincode
	}
	if req.Photo != nil {
		updateFields["photo"] = *req.Photo
	}
	if req.AadhaarCardNo != nil {
		updateFields["aadhaar_card_no"] = *req.AadhaarCardNo
	}
	if req.PanCardNo != nil {
		updateFields["pan_card_no"] = *req.PanCardNo
	}

	filter := bson.M{"entity_id": id}
	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		filter = bson.M{"_id": oid}
	}

	// Handle Documents appending
	if len(req.Documents) > 0 {
		var existing models.Customer
		err := collection.FindOne(ctx, filter).Decode(&existing)
		if err == nil {
			docList := existing.Documents
			for _, d := range req.Documents {
				docList = append(docList, d.Filename)
			}
			updateFields["documents"] = docList
		}
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
		return nil, errors.New("customer not found")
	}

	var updated models.Customer
	_ = collection.FindOne(ctx, filter).Decode(&updated)

	return &updated, nil
}

// ================== DELETE CUSTOMER ==================

func (s *customerService) Delete(
	ctx context.Context,
	companyCode string,
	id string,
) error {

	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(CustomerCollection)

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
		return errors.New("customer not found")
	}

	fmt.Println("Customer soft deleted successfully")
	return nil
}

// ================== GET CUSTOMER LEDGER ==================

func (s *customerService) GetLedger(
	ctx context.Context,
	companyCode string,
	customerId string,
) ([]*models.LedgerEntry, error) {

	db := storage.GetMongo()
	dbName := fmt.Sprintf("company_%s", companyCode)

	salesCollection := db.Database(dbName).Collection("sales_orders")
	paymentsCollection := db.Database(dbName).Collection("payments")

	// Fetch Sales Orders
	sCursor, err := salesCollection.Find(ctx, bson.M{"customer_id": customerId, "is_deleted": bson.M{"$ne": true}})
	if err != nil {
		return nil, err
	}
	defer sCursor.Close(ctx)
	var sales []*models.SalesOrder
	if err := sCursor.All(ctx, &sales); err != nil {
		return nil, err
	}

	// Fetch Payments
	pCursor, err := paymentsCollection.Find(ctx, bson.M{"customer_id": customerId, "is_deleted": bson.M{"$ne": true}})
	if err != nil {
		return nil, err
	}
	defer pCursor.Close(ctx)
	var payments []*models.Payment
	if err := pCursor.All(ctx, &payments); err != nil {
		return nil, err
	}

	var entries []*models.LedgerEntry

	// Add Sales Orders as Debits
	for _, s := range sales {
		vehicleStr := fmt.Sprintf("%s %s", s.Brand, s.Model)
		if s.ChassisNumber != "" {
			vehicleStr += fmt.Sprintf(" (%s)", s.ChassisNumber)
		}

		entries = append(entries, &models.LedgerEntry{
			ID:          s.EntityID,
			Date:        s.SaleDate,
			Description: "Vehicle Sale",
			Debit:       s.TotalAmount,
			Credit:      0,
			VehicleName: vehicleStr,
			VehicleID:   s.EntityID,
		})
	}

	// Add Payments as Credits
	for _, p := range payments {
		vName := "Other/Unassigned"
		vID := p.SalesOrderID
		if vID == "" {
			vID = "unassigned"
		} else {
			for _, s := range sales {
				if s.EntityID == vID || s.ID.Hex() == vID {
					vName = fmt.Sprintf("%s %s", s.Brand, s.Model)
					if s.ChassisNumber != "" {
						vName += fmt.Sprintf(" (%s)", s.ChassisNumber)
					}
					vID = s.EntityID
					break
				}
			}
		}

		entries = append(entries, &models.LedgerEntry{
			ID:          p.EntityID,
			Date:        p.PaymentDate,
			Description: fmt.Sprintf("%s — %s", p.PaymentType, p.PaymentMode),
			Debit:       0,
			Credit:      p.PaymentAmount,
			VehicleName: vName,
			VehicleID:   vID,
		})
	}

	// Sort by date
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Date.Before(entries[j].Date)
	})

	// Calculate balance
	var balance float64
	for i := range entries {
		balance += entries[i].Debit
		balance -= entries[i].Credit
		entries[i].Balance = balance
	}

	return entries, nil
}