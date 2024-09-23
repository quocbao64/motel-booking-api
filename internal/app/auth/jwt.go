package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

type JWTClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(email string) (string, error) {
	expiresAt := jwt.NewNumericDate(time.Now().Add(time.Minute + 15))
	issuedAt := jwt.NewNumericDate(time.Now())
	claims := JWTClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  issuedAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func VerifyToken(tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if !token.Valid {
		return err
	}

	return err
}
