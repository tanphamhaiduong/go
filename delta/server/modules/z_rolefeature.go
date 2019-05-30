// @generated
package modules

import (
	"context"

	"github.com/tanphamhaiduong/go/delta/server/arguments"
	"github.com/tanphamhaiduong/go/delta/server/models"
)

// ICoreRoleFeatureHandler ...
type ICoreRoleFeatureHandler interface {
	GetByID(ctx context.Context, params arguments.RoleFeatureGetByIDArgs) (models.RoleFeature, error)
	Count(ctx context.Context, params arguments.RoleFeatureCountArgs) (int64, error)
	List(ctx context.Context, params arguments.RoleFeatureListArgs) ([]models.RoleFeature, error)
	Insert(ctx context.Context, params arguments.RoleFeatureInsertArgs) (models.RoleFeature, error)
	Update(ctx context.Context, params arguments.RoleFeatureUpdateArgs) (models.RoleFeature, error)
	Delete(ctx context.Context, params arguments.RoleFeatureDeleteArgs) (int64, error)
}

// ICoreRoleFeatureResolver ...
type ICoreRoleFeatureResolver interface {
	IResolver
}

//go:generate mockery -name=IRoleFeatureResolver -output=mocks -case=underscore
//go:generate mockery -name=IRoleFeatureHandler -output=mocks -case=underscore
