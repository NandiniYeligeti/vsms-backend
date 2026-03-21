package middleware

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Email       string `json:"email"`
	Role        string `json:"role"`
	CompanyCode string `json:"company_code"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte("your-very-secret-key") // Use os.Getenv in production

func GenerateJWT(email, role, companyCode string) (string, error) {
	claims := Claims{
		Email:       email,
		Role:        role,
		CompanyCode: companyCode,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
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
