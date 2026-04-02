package services

import (
	"context"
	"fmt"
	"time"
	"vehiclesales/models"
	"vehiclesales/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const CompanySettingsCollection = "company_settings"

type CompanySettingsService interface {
	Get(ctx context.Context, companyCode string) (*models.CompanySettings, error)
	Update(ctx context.Context, companyCode string, settings *models.CompanySettings) error
}

type companySettingsService struct{}

func NewCompanySettingsService() CompanySettingsService {
	return &companySettingsService{}
}

func (s *companySettingsService) Get(ctx context.Context, companyCode string) (*models.CompanySettings, error) {
	db := storage.GetMongo()
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(CompanySettingsCollection)

	var settings models.CompanySettings
	err := collection.FindOne(ctx, bson.M{}).Decode(&settings)
	if err != nil {
		// If not found, return a default one
		return models.NewCompanySettings(companyCode), nil
	}

	return &settings, nil
}

func (s *companySettingsService) Update(ctx context.Context, companyCode string, settings *models.CompanySettings) error {
	db := storage.GetMongo()
	
	// 1. Update company-specific settings
	collection := db.Database(fmt.Sprintf("company_%s", companyCode)).Collection(CompanySettingsCollection)
	settings.CompanyID = companyCode
	settings.UpdatedAt = time.Now()
	opts := options.Replace().SetUpsert(true)
	if _, err := collection.ReplaceOne(ctx, bson.M{}, settings, opts); err != nil {
		return err
	}

	// 2. Sync changes to master users collection for Super Admin view
	masterUsersColl := db.Database("vsms_master").Collection("users")
	_, err := masterUsersColl.UpdateMany(ctx, 
		bson.M{"company_code": companyCode},
		bson.M{"$set": bson.M{
			"company_name": settings.CompanyName,
			"logo_url":     settings.LogoURL,
			"updated_at":   time.Now(),
		}},
	)
	
	return err
}
