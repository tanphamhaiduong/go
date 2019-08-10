package models

import (
	"github.com/dgrijalva/jwt-go"
)

// Claims ...
type Claims struct {
	ID          int64        `json:"id"`
	Username    string       `json:"username"`
	CompanyID   int64        `json:"companyId"`
	Permissions []Permission `json:"permissions"`
	jwt.StandardClaims
}

// ContextKey ...
type ContextKey string
