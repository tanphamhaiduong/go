// @generated
package rolefeature

import (
	"context"

	"github.com/tanphamhaiduong/go/delta/server/arguments"
	"github.com/tanphamhaiduong/go/delta/server/models"
	"github.com/tanphamhaiduong/go/delta/server/validator"
)

// iCoreRepository ...
type iCoreRepository interface {
	GetByID(ctx context.Context, params arguments.RoleFeatureGetByIDArgs) (models.RoleFeature, error)
	GetByIDs(ctx context.Context, params arguments.RoleFeatureGetByIDsArgs) ([]models.RoleFeature, error)
	List(ctx context.Context, params arguments.RoleFeatureListArgs) ([]models.RoleFeature, error)
	Count(ctx context.Context, params arguments.RoleFeatureCountArgs) (int64, error)
	Insert(ctx context.Context, params arguments.RoleFeatureInsertArgs) (models.RoleFeature, error)
	Update(ctx context.Context, params arguments.RoleFeatureUpdateArgs) (models.RoleFeature, error)
	Delete(ctx context.Context, params arguments.RoleFeatureDeleteArgs) (int64, error)
}

// GetByID ...
func (h HandlerImpl) GetByID(ctx context.Context, params arguments.RoleFeatureGetByIDArgs) (models.RoleFeature, error) {
	var (
		rolefeature = models.RoleFeature{}
	)
	rolefeature, err := h.rolefeature.GetByID(ctx, params)
	if err != nil {
		return rolefeature, err
	}
	return rolefeature, nil
}

// GetByIDs ...
func (h HandlerImpl) GetByIDs(ctx context.Context, params arguments.RoleFeatureGetByIDsArgs) ([]models.RoleFeature, error) {
	var (
		rolefeatures = []models.RoleFeature{}
	)
	rolefeatures, err := h.rolefeature.GetByIDs(ctx, params)
	if err != nil {
		return rolefeatures, err
	}
	return rolefeatures, nil
}

// Count ...
func (h HandlerImpl) Count(ctx context.Context, params arguments.RoleFeatureCountArgs) (int64, error) {
	var (
		count int64
	)
	if err := validator.Struct(params); err != nil {
		return count, err
	}
	count, err := h.rolefeature.Count(ctx, params)
	if err != nil {
		return count, err
	}
	return count, nil
}

// List ...
func (h HandlerImpl) List(ctx context.Context, params arguments.RoleFeatureListArgs) ([]models.RoleFeature, error) {
	var (
		rolefeatures = []models.RoleFeature{}
	)
	if err := validator.Struct(params); err != nil {
		return rolefeatures, err
	}
	rolefeatures, err := h.rolefeature.List(ctx, params)
	if err != nil {
		return rolefeatures, err
	}
	return rolefeatures, nil
}

// Insert ...
func (h HandlerImpl) Insert(ctx context.Context, params arguments.RoleFeatureInsertArgs) (models.RoleFeature, error) {
	var (
		rolefeature = models.RoleFeature{}
	)
	if err := validator.Struct(params); err != nil {
		return rolefeature, err
	}
	rolefeature, err := h.rolefeature.Insert(ctx, params)
	if err != nil {
		return rolefeature, err
	}
	return rolefeature, nil
}

// Update ...
func (h HandlerImpl) Update(ctx context.Context, params arguments.RoleFeatureUpdateArgs) (models.RoleFeature, error) {
	var (
		rolefeature = models.RoleFeature{}
	)
	if err := validator.Struct(params); err != nil {
		return rolefeature, err
	}
	rolefeature, err := h.rolefeature.Update(ctx, params)
	if err != nil {
		return rolefeature, err
	}
	return rolefeature, nil
}

// Delete ...
func (h HandlerImpl) Delete(ctx context.Context, params arguments.RoleFeatureDeleteArgs) (int64, error) {
	var (
		id int64
	)
	if err := validator.Struct(params); err != nil {
		return id, err
	}
	id, err := h.rolefeature.Delete(ctx, params)
	if err != nil {
		return id, err
	}
	return id, nil
}

//go:generate mockery -name=iRepository -output=mocks -case=underscore
