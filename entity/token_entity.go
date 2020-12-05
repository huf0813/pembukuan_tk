package entity

import "github.com/dgrijalva/jwt-go"

type Token struct {
	Username   string `json:"username"`
	UserTypeID int    `json:"user_type_id"`
	UserID     int    `json:"user_id"`
	jwt.StandardClaims
}

type TokenExtract struct {
	Username   string `json:"username"`
	UserTypeID int    `json:"user_type_id"`
	UserID     int    `json:"user_id"`
	Token      string `json:"token"`
}

type TokenReq struct {
	Token string `json:"token"`
}
