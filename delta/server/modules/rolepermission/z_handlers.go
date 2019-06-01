// @generated
package rolepermission

import (
	"context"

	"github.com/tanphamhaiduong/go/delta/server/arguments"
	"github.com/tanphamhaiduong/go/delta/server/models"
	"github.com/tanphamhaiduong/go/delta/server/validator"
)

// iCoreRepository ...
type iCoreRepository interface {
	GetByID(ctx context.Context, params arguments.RolePermissionGetByIDArgs) (models.RolePermission, error)
	GetByIDs(ctx context.Context, params arguments.RolePermissionGetByIDsArgs) ([]models.RolePermission, error)
	List(ctx context.Context, params arguments.RolePermissionListArgs) ([]models.RolePermission, error)
	Count(ctx context.Context, params arguments.RolePermissionCountArgs) (int64, error)
	Insert(ctx context.Context, params arguments.RolePermissionInsertArgs) (models.RolePermission, error)
	Update(ctx context.Context, params arguments.RolePermissionUpdateArgs) (models.RolePermission, error)
	Delete(ctx context.Context, params arguments.RolePermissionDeleteArgs) (int64, error)
}

// GetByID ...
func (h HandlerImpl) GetByID(ctx context.Context, params arguments.RolePermissionGetByIDArgs) (models.RolePermission, error) {
	var (
		rolepermission = models.RolePermission{}
	)
	rolepermission, err := h.rolepermission.GetByID(ctx, params)
	if err != nil {
		return rolepermission, err
	}
	return rolepermission, nil
}

// GetByIDs ...
func (h HandlerImpl) GetByIDs(ctx context.Context, params arguments.RolePermissionGetByIDsArgs) ([]models.RolePermission, error) {
	var (
		rolepermissions = []models.RolePermission{}
	)
	rolepermissions, err := h.rolepermission.GetByIDs(ctx, params)
	if err != nil {
		return rolepermissions, err
	}
	return rolepermissions, nil
}

// Count ...
func (h HandlerImpl) Count(ctx context.Context, params arguments.RolePermissionCountArgs) (int64, error) {
	var (
		count int64
	)
	if err := validator.Struct(params); err != nil {
		return count, err
	}
	count, err := h.rolepermission.Count(ctx, params)
	if err != nil {
		return count, err
	}
	return count, nil
}

// List ...
func (h HandlerImpl) List(ctx context.Context, params arguments.RolePermissionListArgs) ([]models.RolePermission, error) {
	var (
		rolepermissions = []models.RolePermission{}
	)
	if err := validator.Struct(params); err != nil {
		return rolepermissions, err
	}
	rolepermissions, err := h.rolepermission.List(ctx, params)
	if err != nil {
		return rolepermissions, err
	}
	return rolepermissions, nil
}

// Insert ...
func (h HandlerImpl) Insert(ctx context.Context, params arguments.RolePermissionInsertArgs) (models.RolePermission, error) {
	var (
		rolepermission = models.RolePermission{}
	)
	if err := validator.Struct(params); err != nil {
		return rolepermission, err
	}
	rolepermission, err := h.rolepermission.Insert(ctx, params)
	if err != nil {
		return rolepermission, err
	}
	return rolepermission, nil
}

// Update ...
func (h HandlerImpl) Update(ctx context.Context, params arguments.RolePermissionUpdateArgs) (models.RolePermission, error) {
	var (
		rolepermission = models.RolePermission{}
	)
	if err := validator.Struct(params); err != nil {
		return rolepermission, err
	}
	rolepermission, err := h.rolepermission.Update(ctx, params)
	if err != nil {
		return rolepermission, err
	}
	return rolepermission, nil
}

// Delete ...
func (h HandlerImpl) Delete(ctx context.Context, params arguments.RolePermissionDeleteArgs) (int64, error) {
	var (
		id int64
	)
	if err := validator.Struct(params); err != nil {
		return id, err
	}
	id, err := h.rolepermission.Delete(ctx, params)
	if err != nil {
		return id, err
	}
	return id, nil
}

//go:generate mockery -name=iRepository -output=mocks -case=underscore
