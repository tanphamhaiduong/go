package rolepermission

import (
	"context"

	"github.com/tanphamhaiduong/go/delta/internal/models"
)

// IHandler ...
type IHandler interface {
	ICoreHandler
	GetByRoleID(ctx context.Context, roleID int64) ([]models.RolePermission, error)
}

// ResolverImpl ...
type ResolverImpl struct {
	rolepermission IHandler
}

// NewResolver ...
func NewResolver(rolepermission IHandler) *ResolverImpl {
	return &ResolverImpl{
		rolepermission: rolepermission,
	}
}
