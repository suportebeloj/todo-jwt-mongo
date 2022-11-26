package utils

import (
	"encoding/base64"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
	"todo-jwt-mongo/internal/core/authentication/models"
)

type Claims struct {
	Username   string
	Group      string
	Permission int8
	jwt.RegisteredClaims
}

func NewToken(user *models.User) (string, error) {
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		Username:   user.Username,
		Group:      user.Permission.Group,
		Permission: user.Permission.Level,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	tokenBase64 := EncodeToBase64([]byte(tokenString))

	return tokenBase64, nil
}

func VerifyToken(tokenString string) error {
	decodedToken, err := DecodeFromBase64(tokenString)

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(decodedToken, claims, func(tkn *jwt.Token) (interface{}, error) {
		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return errors.New("signature invalid")
		}
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}

func EncodeToBase64(jwtToken []byte) string {
	return base64.StdEncoding.EncodeToString(jwtToken)
}

func DecodeFromBase64(base64Hash string) (string, error) {
	token, err := base64.StdEncoding.DecodeString(base64Hash)
	if err != nil {
		return "", err
	}

	return string(token), nil
}
