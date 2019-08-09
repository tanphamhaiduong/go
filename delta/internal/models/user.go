package models

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	User
	jwt.StandardClaims
	Permission Permission
}

type RefreshClaims struct {
	ID int64
	jwt.StandardClaims
}
