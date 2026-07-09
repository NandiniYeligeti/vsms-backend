package services

import (
	"context"
	"errors"
	"os"
	"time"

	"vehiclesales/models"
	"vehiclesales/storage"
	"vehiclesales/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// EnsureSuperAdmin checks whether a super_admin user exists and seeds one from environment variables.
func EnsureSuperAdmin(ctx context.Context) error {
	db := storage.GetMongo()
	masterDB := db.Database(MasterDatabase)
	usersColl := masterDB.Collection(UsersCollection)

	count, err := usersColl.CountDocuments(ctx, bson.M{"role": "super_admin", "is_deleted": false})
	if err != nil {
		return err
	}
	if count > 0 {
		return nil
	}

	superAdminEmail := os.Getenv("SUPER_ADMIN_EMAIL")
	superAdminPassword := os.Getenv("SUPER_ADMIN_PASSWORD")
	if superAdminEmail == "" || superAdminPassword == "" {
		return errors.New("SUPER_ADMIN_EMAIL and SUPER_ADMIN_PASSWORD must be set when no super_admin exists")
	}

	emailConflict, err := usersColl.CountDocuments(ctx, bson.M{"email": superAdminEmail, "is_deleted": false})
	if err != nil {
		return err
	}
	if emailConflict > 0 {
		return errors.New("cannot create super_admin: a user with the configured SUPER_ADMIN_EMAIL already exists")
	}

	hashedPassword, err := utils.HashPassword(superAdminPassword)
	if err != nil {
		return err
	}

	now := time.Now()
	superAdmin := models.User{
		ID:          primitive.NewObjectID(),
		Email:       superAdminEmail,
		Password:    hashedPassword,
		Role:        "super_admin",
		CompanyCode: "SUPER",
		CompanyName: "System Admin",
		IsDeleted:   false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	_, err = usersColl.InsertOne(ctx, superAdmin)
	return err
}
