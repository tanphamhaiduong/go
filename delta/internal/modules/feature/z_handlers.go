// @generated
package feature

import (
	"context"

	log "github.com/sirupsen/logrus"
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
	log.WithField("params", params).Info("feature of Handler GetByID")
	var (
		feature models.Feature
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("feature of Handler GetByID validator.Struct error")
		return feature, err
	}
	feature, err := h.feature.GetByID(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("feature of Handler GetByID h.feature.GetByID error")
		return feature, err
	}
	return feature, nil
}

// GetByIDs ...
func (h *HandlerImpl) GetByIDs(ctx context.Context, params arguments.FeatureGetByIDsArgs) ([]models.Feature, error) {
	log.WithField("params", params).Info("feature of Handler GetByIDs")
	var (
		features []models.Feature
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("feature of Handler GetByIDs validator.Struct error")
		return features, err
	}
	features, err := h.feature.GetByIDs(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("feature of Handler GetByIDs h.feature.GetByIDs error")
		return features, err
	}
	return features, nil
}

// Count ...
func (h *HandlerImpl) Count(ctx context.Context, params arguments.FeatureCountArgs) (int64, error) {
	log.WithField("params", params).Info("feature of Handler Count")
	var (
		count int64
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("feature of Handler Count validator.Struct error")
		return count, err
	}
	count, err := h.feature.Count(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("feature of Handler Count h.feature.Count error")
		return count, err
	}
	return count, nil
}

// List ...
func (h *HandlerImpl) List(ctx context.Context, params arguments.FeatureListArgs) ([]models.Feature, error) {
	log.WithField("params", params).Info("feature of Handler List")
	var (
		features []models.Feature
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("feature of Handler List validator.Struct error")
		return features, err
	}
	features, err := h.feature.List(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("feature of Handler List h.feature.List error")
		return features, err
	}
	return features, nil
}

// Insert ...
func (h *HandlerImpl) Insert(ctx context.Context, params arguments.FeatureInsertArgs) (models.Feature, error) {
	log.WithField("params", params).Info("feature of Handler Insert")
	var (
		feature models.Feature
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("feature of Handler Insert validator.Struct error")
		return feature, err
	}
	feature, err := h.feature.Insert(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("feature of Handler Insert h.feature.Insert error")
		return feature, err
	}
	return feature, nil
}

// Update ...
func (h *HandlerImpl) Update(ctx context.Context, params arguments.FeatureUpdateArgs) (models.Feature, error) {
	log.WithField("params", params).Info("feature of Handler Update")
	var (
		feature models.Feature
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("feature of Handler Update validator.Struct error")
		return feature, err
	}
	feature, err := h.feature.Update(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("feature of Handler Update h.feature.Update error")
		return feature, err
	}
	return feature, nil
}

// Delete ...
func (h *HandlerImpl) Delete(ctx context.Context, params arguments.FeatureDeleteArgs) (int64, error) {
	log.WithField("params", params).Info("feature of Handler Delete")
	var (
		id int64
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("feature of Handler Delete validator.Struct error")
		return id, err
	}
	id, err := h.feature.Delete(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("feature of Handler Delete h.feature.Delete error")
		return id, err
	}
	return id, nil
}

//go:generate mockery -name=IRepository -output=mocks -case=underscore
