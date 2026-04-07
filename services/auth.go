package services

import (
	"context"
	"errors"
	"time"

	"vehiclesales/middleware"
	"vehiclesales/models"
	"vehiclesales/requests"
	"vehiclesales/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const MasterDatabase = "vsms_master"
const UsersCollection = "users"

type AuthResponse struct {
	User  *models.User `json:"user"`
	Token string       `json:"token"`
}

type AuthService interface {
	Login(ctx context.Context, req *requests.LoginRequest) (*AuthResponse, error)
	CreateCompany(ctx context.Context, req *requests.CreateCompanyRequest) (*models.User, error)
	GetCompanies(ctx context.Context) ([]*models.User, error)
	CreateUser(ctx context.Context, req *requests.CreateUserRequest, companyCode string, companyName string) (*models.User, error)
	GetUsers(ctx context.Context, companyCode string) ([]*models.User, error)
	UpdateUserMenus(ctx context.Context, userID string, menus []string, permissions []requests.MenuPermission) error
	DeleteUser(ctx context.Context, userID string) error
}

type authService struct{}

func NewAuthService() AuthService {
	return &authService{}
}

func (s *authService) Login(ctx context.Context, req *requests.LoginRequest) (*AuthResponse, error) {
	// Super Admin Hardcoded Login
	if req.Email == "NandiniY" && req.Password == "Nihnan@853" {
		superAdmin := &models.User{
			ID:          primitive.NewObjectID(),
			Email:       req.Email,
			Role:        "super_admin",
			CompanyCode: "SUPER", // Just a placeholder
			CompanyName: "System Admin",
		}
		token, err := middleware.GenerateJWT(superAdmin.Email, superAdmin.Role, superAdmin.CompanyCode)
		if err != nil {
			return nil, err
		}
		return &AuthResponse{
			User:  superAdmin,
			Token: token,
		}, nil
	}

	// Normal Company Admin Login
	db := storage.GetMongo()
	masterDB := db.Database(MasterDatabase)

	var user models.User
	err := masterDB.Collection(UsersCollection).FindOne(ctx, bson.M{
		"email":      req.Email,
		"is_deleted": false,
	}).Decode(&user)

	if err == mongo.ErrNoDocuments {
		return nil, errors.New("invalid email or password")
	}
	if err != nil {
		return nil, err
	}

	// In a real app, use bcrypt.CompareHashAndPassword instead of plaintext check
	if user.Password != req.Password {
		return nil, errors.New("invalid email or password")
	}

	token, err := middleware.GenerateJWT(user.Email, user.Role, user.CompanyCode)
	if err != nil {
		return nil, err
	}
	return &AuthResponse{
		User:  &user,
		Token: token,
	}, nil
}

func (s *authService) CreateCompany(ctx context.Context, req *requests.CreateCompanyRequest) (*models.User, error) {
	db := storage.GetMongo()
	masterDB := db.Database(MasterDatabase)
	usersColl := masterDB.Collection(UsersCollection)

	// Check if company code already exists
	count, _ := usersColl.CountDocuments(ctx, bson.M{"company_code": req.CompanyCode, "is_deleted": false})
	if count > 0 {
		return nil, errors.New("company code already exists")
	}

	// Check if admin email already exists globally
	count, _ = usersColl.CountDocuments(ctx, bson.M{"email": req.AdminEmail, "is_deleted": false})
	if count > 0 {
		return nil, errors.New("admin email already exists")
	}

	// Create user
	now := time.Now()
	user := models.User{
		ID:          primitive.NewObjectID(),
		Email:       req.AdminEmail,
		Password:    req.AdminPassword, // In a real app, hash this with bcrypt
		Role:        "admin",           // They are the admin of their respective company
		CompanyCode: req.CompanyCode,
		CompanyName: req.CompanyName,
		IsDeleted:   false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	_, err := usersColl.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	// Create initial database & a dummy collection (optional, mongodb creates on fly)
	// Seed their database with default records? Left for future enhancements.

	return &user, nil
}

func (s *authService) GetCompanies(ctx context.Context) ([]*models.User, error) {
	db := storage.GetMongo()
	masterDB := db.Database(MasterDatabase)

	cursor, err := masterDB.Collection(UsersCollection).Find(ctx, bson.M{"is_deleted": false, "role": "admin"})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []*models.User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *authService) CreateUser(ctx context.Context, req *requests.CreateUserRequest, companyCode string, companyName string) (*models.User, error) {
	db := storage.GetMongo()
	masterDB := db.Database(MasterDatabase)
	usersColl := masterDB.Collection(UsersCollection)

	// Check if email already exists
	count, _ := usersColl.CountDocuments(ctx, bson.M{"email": req.Email, "is_deleted": false})
	if count > 0 {
		return nil, errors.New("email already exists")
	}

	// Look up company name from admin record if not provided
	if companyName == "" {
		var admin models.User
		err := usersColl.FindOne(ctx, bson.M{
			"company_code": companyCode,
			"role":         "admin",
			"is_deleted":   false,
		}).Decode(&admin)
		if err == nil {
			companyName = admin.CompanyName
		}
	}

	now := time.Now()
	user := models.User{
		ID:          primitive.NewObjectID(),
		Username:    req.Username,
		Email:       req.Email,
		Password:    req.Password, // In a real app, hash this with bcrypt
		Role:        "user",
		CompanyCode: companyCode,
		CompanyName: companyName,
		Menus:       req.Menus,
		Permissions: req.Permissions,
		IsDeleted:   false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	_, err := usersColl.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *authService) GetUsers(ctx context.Context, companyCode string) ([]*models.User, error) {
	db := storage.GetMongo()
	masterDB := db.Database(MasterDatabase)

	cursor, err := masterDB.Collection(UsersCollection).Find(ctx, bson.M{
		"is_deleted":    false,
		"role":          "user",
		"company_code":  companyCode,
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []*models.User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *authService) UpdateUserMenus(ctx context.Context, userID string, menus []string, permissions []requests.MenuPermission) error {
	db := storage.GetMongo()
	masterDB := db.Database(MasterDatabase)

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("invalid user ID")
	}

	_, err = masterDB.Collection(UsersCollection).UpdateOne(ctx,
		bson.M{"_id": objID, "role": "user"},
		bson.M{"$set": bson.M{
			"menus":       menus,
			"permissions": permissions,
			"updated_at":  time.Now(),
		}},
	)
	return err
}

func (s *authService) DeleteUser(ctx context.Context, userID string) error {
	db := storage.GetMongo()
	masterDB := db.Database(MasterDatabase)

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("invalid user ID")
	}

	_, err = masterDB.Collection(UsersCollection).UpdateOne(ctx,
		bson.M{"_id": objID, "role": "user"},
		bson.M{"$set": bson.M{"is_deleted": true, "updated_at": time.Now()}},
	)
	if err != nil {
		return err
	}

	return nil
}
