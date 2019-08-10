// @generated
package modules

import (
	"context"

	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
)

// ICorePermissionHandler ...
type ICorePermissionHandler interface {
	GetByID(ctx context.Context, params arguments.PermissionGetByID) (models.Permission, error)
	Count(ctx context.Context, params arguments.PermissionCount) (int64, error)
	List(ctx context.Context, params arguments.PermissionList) ([]models.Permission, error)
	Insert(ctx context.Context, params arguments.PermissionInsert) (models.Permission, error)
	Update(ctx context.Context, params arguments.PermissionUpdate) (models.Permission, error)
	Delete(ctx context.Context, params arguments.PermissionDelete) (int64, error)
}

// ICorePermissionResolver ...
type ICorePermissionResolver interface {
	IResolver
}

//go:generate mockery -name=IPermissionResolver -output=mocks -case=underscore
//go:generate mockery -name=IPermissionHandler -output=mocks -case=underscore
