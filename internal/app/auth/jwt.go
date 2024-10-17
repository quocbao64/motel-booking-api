package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

type JWTClaims struct {
	Phone string `json:"phone"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(phone, role string) (string, error) {
	expiresAt := jwt.NewNumericDate(time.Now().Add(time.Minute * 60 * 24))
	issuedAt := jwt.NewNumericDate(time.Now())
	claims := JWTClaims{
		Phone: phone,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  issuedAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func VerifyToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if token == nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
