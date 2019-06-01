// @generated
package modules

import (
	"context"

	"github.com/tanphamhaiduong/go/delta/server/arguments"
	"github.com/tanphamhaiduong/go/delta/server/models"
)

// ICoreRolePermissionHandler ...
type ICoreRolePermissionHandler interface {
	GetByID(ctx context.Context, params arguments.RolePermissionGetByIDArgs) (models.RolePermission, error)
	Count(ctx context.Context, params arguments.RolePermissionCountArgs) (int64, error)
	List(ctx context.Context, params arguments.RolePermissionListArgs) ([]models.RolePermission, error)
	Insert(ctx context.Context, params arguments.RolePermissionInsertArgs) (models.RolePermission, error)
	Update(ctx context.Context, params arguments.RolePermissionUpdateArgs) (models.RolePermission, error)
	Delete(ctx context.Context, params arguments.RolePermissionDeleteArgs) (int64, error)
}

// ICoreRolePermissionResolver ...
type ICoreRolePermissionResolver interface {
	IResolver
}

//go:generate mockery -name=IRolePermissionResolver -output=mocks -case=underscore
//go:generate mockery -name=IRolePermissionHandler -output=mocks -case=underscore
