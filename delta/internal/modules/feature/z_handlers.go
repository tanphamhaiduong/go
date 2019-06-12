// @generated
package feature

import (
	"context"

	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
	"github.com/tanphamhaiduong/go/delta/internal/validator"
)

// ICoreRepository ...
type ICoreRepository interface {
	GetByID(ctx context.Context, params arguments.FeatureGetByIDArgs) (models.Feature, error)
	GetByIDs(ctx context.Context, params arguments.FeatureGetByIDsArgs) ([]models.Feature, error)
	List(ctx context.Context, params arguments.FeatureListArgs) ([]models.Feature, error)
	Count(ctx context.Context, params arguments.FeatureCountArgs) (int64, error)
	Insert(ctx context.Context, params arguments.FeatureInsertArgs) (models.Feature, error)
	Update(ctx context.Context, params arguments.FeatureUpdateArgs) (models.Feature, error)
	Delete(ctx context.Context, params arguments.FeatureDeleteArgs) (int64, error)
}

// GetByID ...
func (h *HandlerImpl) GetByID(ctx context.Context, params arguments.FeatureGetByIDArgs) (models.Feature, error) {
	var (
		feature models.Feature
	)
	feature, err := h.feature.GetByID(ctx, params)
	if err != nil {
		return feature, err
	}
	return feature, nil
}

// GetByIDs ...
func (h *HandlerImpl) GetByIDs(ctx context.Context, params arguments.FeatureGetByIDsArgs) ([]models.Feature, error) {
	var (
		features []models.Feature
	)
	features, err := h.feature.GetByIDs(ctx, params)
	if err != nil {
		return features, err
	}
	return features, nil
}

// Count ...
func (h *HandlerImpl) Count(ctx context.Context, params arguments.FeatureCountArgs) (int64, error) {
	var (
		count int64
	)
	if err := validator.Struct(params); err != nil {
		return count, err
	}
	count, err := h.feature.Count(ctx, params)
	if err != nil {
		return count, err
	}
	return count, nil
}

// List ...
func (h *HandlerImpl) List(ctx context.Context, params arguments.FeatureListArgs) ([]models.Feature, error) {
	var (
		features []models.Feature
	)
	if err := validator.Struct(params); err != nil {
		return features, err
	}
	features, err := h.feature.List(ctx, params)
	if err != nil {
		return features, err
	}
	return features, nil
}

// Insert ...
func (h *HandlerImpl) Insert(ctx context.Context, params arguments.FeatureInsertArgs) (models.Feature, error) {
	var (
		feature models.Feature
	)
	if err := validator.Struct(params); err != nil {
		return feature, err
	}
	feature, err := h.feature.Insert(ctx, params)
	if err != nil {
		return feature, err
	}
	return feature, nil
}

// Update ...
func (h *HandlerImpl) Update(ctx context.Context, params arguments.FeatureUpdateArgs) (models.Feature, error) {
	var (
		feature models.Feature
	)
	if err := validator.Struct(params); err != nil {
		return feature, err
	}
	feature, err := h.feature.Update(ctx, params)
	if err != nil {
		return feature, err
	}
	return feature, nil
}

// Delete ...
func (h *HandlerImpl) Delete(ctx context.Context, params arguments.FeatureDeleteArgs) (int64, error) {
	var (
		id int64
	)
	if err := validator.Struct(params); err != nil {
		return id, err
	}
	id, err := h.feature.Delete(ctx, params)
	if err != nil {
		return id, err
	}
	return id, nil
}

//go:generate mockery -name=IRepository -output=mocks -case=underscore
