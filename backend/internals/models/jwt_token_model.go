package models

import "github.com/golang-jwt/jwt/v5"

type TokenBody struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	FullName string `json:"fullname"`
	jwt.RegisteredClaims
}
