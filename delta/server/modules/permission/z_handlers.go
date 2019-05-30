// @generated
package permission

import (
	"context"

	"github.com/tanphamhaiduong/go/delta/server/arguments"
	"github.com/tanphamhaiduong/go/delta/server/models"
	"github.com/tanphamhaiduong/go/delta/server/validator"
)

// iCoreRepository ...
type iCoreRepository interface {
	GetByID(ctx context.Context, params arguments.PermissionGetByIDArgs) (models.Permission, error)
	GetByIDs(ctx context.Context, params arguments.PermissionGetByIDsArgs) ([]models.Permission, error)
	List(ctx context.Context, params arguments.PermissionListArgs) ([]models.Permission, error)
	Count(ctx context.Context, params arguments.PermissionCountArgs) (int64, error)
	Insert(ctx context.Context, params arguments.PermissionInsertArgs) (models.Permission, error)
	Update(ctx context.Context, params arguments.PermissionUpdateArgs) (models.Permission, error)
	Delete(ctx context.Context, params arguments.PermissionDeleteArgs) (int64, error)
}

// GetByID ...
func (h HandlerImpl) GetByID(ctx context.Context, params arguments.PermissionGetByIDArgs) (models.Permission, error) {
	var (
		permission = models.Permission{}
	)
	permission, err := h.permission.GetByID(ctx, params)
	if err != nil {
		return permission, err
	}
	return permission, nil
}

// GetByIDs ...
func (h HandlerImpl) GetByIDs(ctx context.Context, params arguments.PermissionGetByIDsArgs) ([]models.Permission, error) {
	var (
		permissions = []models.Permission{}
	)
	permissions, err := h.permission.GetByIDs(ctx, params)
	if err != nil {
		return permissions, err
	}
	return permissions, nil
}

// Count ...
func (h HandlerImpl) Count(ctx context.Context, params arguments.PermissionCountArgs) (int64, error) {
	var (
		count int64
	)
	if err := validator.Struct(params); err != nil {
		return count, err
	}
	count, err := h.permission.Count(ctx, params)
	if err != nil {
		return count, err
	}
	return count, nil
}

// List ...
func (h HandlerImpl) List(ctx context.Context, params arguments.PermissionListArgs) ([]models.Permission, error) {
	var (
		permissions = []models.Permission{}
	)
	if err := validator.Struct(params); err != nil {
		return permissions, err
	}
	permissions, err := h.permission.List(ctx, params)
	if err != nil {
		return permissions, err
	}
	return permissions, nil
}

// Insert ...
func (h HandlerImpl) Insert(ctx context.Context, params arguments.PermissionInsertArgs) (models.Permission, error) {
	var (
		permission = models.Permission{}
	)
	if err := validator.Struct(params); err != nil {
		return permission, err
	}
	permission, err := h.permission.Insert(ctx, params)
	if err != nil {
		return permission, err
	}
	return permission, nil
}

// Update ...
func (h HandlerImpl) Update(ctx context.Context, params arguments.PermissionUpdateArgs) (models.Permission, error) {
	var (
		permission = models.Permission{}
	)
	if err := validator.Struct(params); err != nil {
		return permission, err
	}
	permission, err := h.permission.Update(ctx, params)
	if err != nil {
		return permission, err
	}
	return permission, nil
}

// Delete ...
func (h HandlerImpl) Delete(ctx context.Context, params arguments.PermissionDeleteArgs) (int64, error) {
	var (
		id int64
	)
	if err := validator.Struct(params); err != nil {
		return id, err
	}
	id, err := h.permission.Delete(ctx, params)
	if err != nil {
		return id, err
	}
	return id, nil
}

//go:generate mockery -name=iRepository -output=mocks -case=underscore
