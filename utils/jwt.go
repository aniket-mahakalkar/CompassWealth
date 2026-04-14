package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWTKey = []byte(getSecret())

func getSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "compass_wealth_base_secret_key"
	}
	return secret
}

type Claims struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(id uint, email string) (string, error) {
	expirationTime := time.Now().Add(72 * time.Hour)
	claims := &Claims{
		ID:    id,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWTKey)

	return tokenString, err
}
