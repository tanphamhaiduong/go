// @generated
package modules

import (
	"context"

	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
)

// ICorePermissionHandler ...
type ICorePermissionHandler interface {
	GetByID(ctx context.Context, params arguments.PermissionGetByIDArgs) (models.Permission, error)
	Count(ctx context.Context, params arguments.PermissionCountArgs) (int64, error)
	List(ctx context.Context, params arguments.PermissionListArgs) ([]models.Permission, error)
	Insert(ctx context.Context, params arguments.PermissionInsertArgs) (models.Permission, error)
	Update(ctx context.Context, params arguments.PermissionUpdateArgs) (models.Permission, error)
	Delete(ctx context.Context, params arguments.PermissionDeleteArgs) (int64, error)
}

// ICorePermissionResolver ...
type ICorePermissionResolver interface {
	IResolver
}

//go:generate mockery -name=IPermissionResolver -output=mocks -case=underscore
//go:generate mockery -name=IPermissionHandler -output=mocks -case=underscore
