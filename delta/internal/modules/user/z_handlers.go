// @generated
package user

import (
	"context"

	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
	"github.com/tanphamhaiduong/go/delta/internal/validator"
)

// ICoreRepository ...
type ICoreRepository interface {
	GetByID(ctx context.Context, params arguments.UserGetByIDArgs) (models.User, error)
	GetByIDs(ctx context.Context, params arguments.UserGetByIDsArgs) ([]models.User, error)
	List(ctx context.Context, params arguments.UserListArgs) ([]models.User, error)
	Count(ctx context.Context, params arguments.UserCountArgs) (int64, error)
	Insert(ctx context.Context, params arguments.UserInsertArgs) (models.User, error)
	Update(ctx context.Context, params arguments.UserUpdateArgs) (models.User, error)
	Delete(ctx context.Context, params arguments.UserDeleteArgs) (int64, error)
}

// GetByID ...
func (h *HandlerImpl) GetByID(ctx context.Context, params arguments.UserGetByIDArgs) (models.User, error) {
	var (
		user models.User
	)
	user, err := h.user.GetByID(ctx, params)
	if err != nil {
		return user, err
	}
	return user, nil
}

// GetByIDs ...
func (h *HandlerImpl) GetByIDs(ctx context.Context, params arguments.UserGetByIDsArgs) ([]models.User, error) {
	var (
		users []models.User
	)
	users, err := h.user.GetByIDs(ctx, params)
	if err != nil {
		return users, err
	}
	return users, nil
}

// Count ...
func (h *HandlerImpl) Count(ctx context.Context, params arguments.UserCountArgs) (int64, error) {
	var (
		count int64
	)
	if err := validator.Struct(params); err != nil {
		return count, err
	}
	count, err := h.user.Count(ctx, params)
	if err != nil {
		return count, err
	}
	return count, nil
}

// List ...
func (h *HandlerImpl) List(ctx context.Context, params arguments.UserListArgs) ([]models.User, error) {
	var (
		users []models.User
	)
	if err := validator.Struct(params); err != nil {
		return users, err
	}
	users, err := h.user.List(ctx, params)
	if err != nil {
		return users, err
	}
	return users, nil
}

// Insert ...
func (h *HandlerImpl) Insert(ctx context.Context, params arguments.UserInsertArgs) (models.User, error) {
	var (
		user models.User
	)
	if err := validator.Struct(params); err != nil {
		return user, err
	}
	user, err := h.user.Insert(ctx, params)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Update ...
func (h *HandlerImpl) Update(ctx context.Context, params arguments.UserUpdateArgs) (models.User, error) {
	var (
		user models.User
	)
	if err := validator.Struct(params); err != nil {
		return user, err
	}
	user, err := h.user.Update(ctx, params)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Delete ...
func (h *HandlerImpl) Delete(ctx context.Context, params arguments.UserDeleteArgs) (int64, error) {
	var (
		id int64
	)
	if err := validator.Struct(params); err != nil {
		return id, err
	}
	id, err := h.user.Delete(ctx, params)
	if err != nil {
		return id, err
	}
	return id, nil
}

//go:generate mockery -name=IRepository -output=mocks -case=underscore
