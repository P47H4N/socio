package helpers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/P47H4N/socio/internals/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secret []byte

func LoadJWT(s string) {
	secret = []byte(s)
}

func GenerateToken(payload *models.TokenBody) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString(secret)
	if err != nil {
		return "", errors.New("Failed to generate JWT Token")
	}
	return token, nil
}

func getKey(tk *jwt.Token) (any, error) {
	if _, ok := tk.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", tk.Header["alg"])
	}
	return secret, nil
}

func ValidateToken(tk string) (*models.TokenBody, error) {
	claims := &models.TokenBody{}
	token, err := jwt.ParseWithClaims(tk, claims, getKey)
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return claims, nil
	}
	return nil, errors.New("Invalid Token.")
}

func GetToken(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", errors.New("No authorization token found.")
	}
	parts := strings.Split(authHeader, " ")
	if len(parts) == 2 && parts[0] == "Bearer" {
		return parts[1], nil
	}
	return "", errors.New("Invalid Token.")
}