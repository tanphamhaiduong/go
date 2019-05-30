// @generated
package role

import (
	"context"

	"github.com/tanphamhaiduong/go/delta/server/arguments"
	"github.com/tanphamhaiduong/go/delta/server/models"
	"github.com/tanphamhaiduong/go/delta/server/validator"
)

// iCoreRepository ...
type iCoreRepository interface {
	GetByID(ctx context.Context, params arguments.RoleGetByIDArgs) (models.Role, error)
	GetByIDs(ctx context.Context, params arguments.RoleGetByIDsArgs) ([]models.Role, error)
	List(ctx context.Context, params arguments.RoleListArgs) ([]models.Role, error)
	Count(ctx context.Context, params arguments.RoleCountArgs) (int64, error)
	Insert(ctx context.Context, params arguments.RoleInsertArgs) (models.Role, error)
	Update(ctx context.Context, params arguments.RoleUpdateArgs) (models.Role, error)
	Delete(ctx context.Context, params arguments.RoleDeleteArgs) (int64, error)
}

// GetByID ...
func (h HandlerImpl) GetByID(ctx context.Context, params arguments.RoleGetByIDArgs) (models.Role, error) {
	var (
		role = models.Role{}
	)
	role, err := h.role.GetByID(ctx, params)
	if err != nil {
		return role, err
	}
	return role, nil
}

// GetByIDs ...
func (h HandlerImpl) GetByIDs(ctx context.Context, params arguments.RoleGetByIDsArgs) ([]models.Role, error) {
	var (
		roles = []models.Role{}
	)
	roles, err := h.role.GetByIDs(ctx, params)
	if err != nil {
		return roles, err
	}
	return roles, nil
}

// Count ...
func (h HandlerImpl) Count(ctx context.Context, params arguments.RoleCountArgs) (int64, error) {
	var (
		count int64
	)
	if err := validator.Struct(params); err != nil {
		return count, err
	}
	count, err := h.role.Count(ctx, params)
	if err != nil {
		return count, err
	}
	return count, nil
}

// List ...
func (h HandlerImpl) List(ctx context.Context, params arguments.RoleListArgs) ([]models.Role, error) {
	var (
		roles = []models.Role{}
	)
	if err := validator.Struct(params); err != nil {
		return roles, err
	}
	roles, err := h.role.List(ctx, params)
	if err != nil {
		return roles, err
	}
	return roles, nil
}

// Insert ...
func (h HandlerImpl) Insert(ctx context.Context, params arguments.RoleInsertArgs) (models.Role, error) {
	var (
		role = models.Role{}
	)
	if err := validator.Struct(params); err != nil {
		return role, err
	}
	role, err := h.role.Insert(ctx, params)
	if err != nil {
		return role, err
	}
	return role, nil
}

// Update ...
func (h HandlerImpl) Update(ctx context.Context, params arguments.RoleUpdateArgs) (models.Role, error) {
	var (
		role = models.Role{}
	)
	if err := validator.Struct(params); err != nil {
		return role, err
	}
	role, err := h.role.Update(ctx, params)
	if err != nil {
		return role, err
	}
	return role, nil
}

// Delete ...
func (h HandlerImpl) Delete(ctx context.Context, params arguments.RoleDeleteArgs) (int64, error) {
	var (
		id int64
	)
	if err := validator.Struct(params); err != nil {
		return id, err
	}
	id, err := h.role.Delete(ctx, params)
	if err != nil {
		return id, err
	}
	return id, nil
}

//go:generate mockery -name=iRepository -output=mocks -case=underscore
