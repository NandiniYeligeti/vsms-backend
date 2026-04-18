package models

import (
	"time"
	"vehiclesales/requests"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username    string             `bson:"username" json:"username"`
	Email       string             `bson:"email" json:"email"`
	Password    string             `bson:"password" json:"-"` // never return password in JSON
	Role        string             `bson:"role" json:"role"`
	CompanyCode string             `bson:"company_code" json:"company_code"`
	CompanyName string             `bson:"company_name" json:"company_name"`
	LogoURL     string             `bson:"logo_url" json:"logo_url"`
	Menus       []string           `bson:"menus" json:"menus"`
	Permissions []requests.MenuPermission `bson:"permissions" json:"permissions"`
	Branches    []string                  `bson:"branches" json:"branches"`
	Showrooms   []string                  `bson:"showrooms" json:"showrooms"`
	Areas       []string                  `bson:"areas" json:"areas"`
	IsDeleted   bool                      `bson:"is_deleted" json:"is_deleted"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}
