package model

import "github.com/dgrijalva/jwt-go"

type Token struct {
	Username   string `json:"username"`
	UserTypeID int    `json:"user_type_id"`
	UserID     int    `json:"user_id"`
	jwt.StandardClaims
}