package arguments

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/tanphamhaiduong/go/delta/internal/models"
)

// UserLogin ...
type UserLogin struct {
	Username string `graphql:"username" validate:"required"`
	Password string `graphql:"password" validate:"required"`
}

// Claims ...
type Claims struct {
	ID          int64               `json:"id"`
	Username    string              `json:"username"`
	CompanyID   int64               `json:"companyId"`
	Permissions []models.Permission `json:"permissions"`
	jwt.StandardClaims
}
