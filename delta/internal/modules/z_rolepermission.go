// @generated
package modules

import (
	"context"

	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
)

// ICoreRolePermissionHandler ...
type ICoreRolePermissionHandler interface {
	GetByID(ctx context.Context, params arguments.RolePermissionGetByID) (models.RolePermission, error)
	Count(ctx context.Context, params arguments.RolePermissionCount) (int64, error)
	List(ctx context.Context, params arguments.RolePermissionList) ([]models.RolePermission, error)
	Insert(ctx context.Context, params arguments.RolePermissionInsert) (models.RolePermission, error)
	Update(ctx context.Context, params arguments.RolePermissionUpdate) (models.RolePermission, error)
	Delete(ctx context.Context, params arguments.RolePermissionDelete) (int64, error)
}

// ICoreRolePermissionResolver ...
type ICoreRolePermissionResolver interface {
	IResolver
}

//go:generate mockery -name=IRolePermissionResolver -output=mocks -case=underscore
//go:generate mockery -name=IRolePermissionHandler -output=mocks -case=underscore
