package middleware

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Email       string `json:"email"`
	Role        string `json:"role"`
	CompanyCode string `json:"company_code"`
	jwt.RegisteredClaims
}

var jwtSecret []byte

func InitJWT() {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Println("⚠️  WARNING: JWT_SECRET not set, using insecure default. Set it in production!")
		secret = "dev-only-insecure-secret-change-me"
	}
	jwtSecret = []byte(secret)
}

func GenerateJWT(email, role, companyCode string) (string, error) {
	claims := Claims{
		Email:       email,
		Role:        role,
		CompanyCode: companyCode,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseAndValidateJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}
	return claims, nil
}
