// @generated
package modules

import (
	"context"

	"github.com/tanphamhaiduong/go/delta/server/arguments"
	"github.com/tanphamhaiduong/go/delta/server/models"
)

// ICoreFeatureHandler ...
type ICoreFeatureHandler interface {
	GetByID(ctx context.Context, params arguments.FeatureGetByIDArgs) (models.Feature, error)
	Count(ctx context.Context, params arguments.FeatureCountArgs) (int64, error)
	List(ctx context.Context, params arguments.FeatureListArgs) ([]models.Feature, error)
	Insert(ctx context.Context, params arguments.FeatureInsertArgs) (models.Feature, error)
	Update(ctx context.Context, params arguments.FeatureUpdateArgs) (models.Feature, error)
	Delete(ctx context.Context, params arguments.FeatureDeleteArgs) (int64, error)
}

// ICoreFeatureResolver ...
type ICoreFeatureResolver interface {
	IResolver
}

//go:generate mockery -name=IFeatureResolver -output=mocks -case=underscore
//go:generate mockery -name=IFeatureHandler -output=mocks -case=underscore
