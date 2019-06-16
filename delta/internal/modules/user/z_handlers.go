// @generated
package user

import (
	"context"

	log "github.com/sirupsen/logrus"
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
	log.WithField("params", params).Info("user of Handler GetByID")
	var (
		user models.User
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("user of Handler GetByID validator.Struct error")
		return user, err
	}
	user, err := h.user.GetByID(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("user of Handler GetByID h.user.GetByID error")
		return user, err
	}
	return user, nil
}

// GetByIDs ...
func (h *HandlerImpl) GetByIDs(ctx context.Context, params arguments.UserGetByIDsArgs) ([]models.User, error) {
	log.WithField("params", params).Info("user of Handler GetByIDs")
	var (
		users []models.User
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("user of Handler GetByIDs validator.Struct error")
		return users, err
	}
	users, err := h.user.GetByIDs(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("user of Handler GetByIDs h.user.GetByIDs error")
		return users, err
	}
	return users, nil
}

// Count ...
func (h *HandlerImpl) Count(ctx context.Context, params arguments.UserCountArgs) (int64, error) {
	log.WithField("params", params).Info("user of Handler Count")
	var (
		count int64
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("user of Handler Count validator.Struct error")
		return count, err
	}
	count, err := h.user.Count(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("user of Handler Count h.user.Count error")
		return count, err
	}
	return count, nil
}

// List ...
func (h *HandlerImpl) List(ctx context.Context, params arguments.UserListArgs) ([]models.User, error) {
	log.WithField("params", params).Info("user of Handler List")
	var (
		users []models.User
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("user of Handler List validator.Struct error")
		return users, err
	}
	users, err := h.user.List(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("user of Handler List h.user.List error")
		return users, err
	}
	return users, nil
}

// Insert ...
func (h *HandlerImpl) Insert(ctx context.Context, params arguments.UserInsertArgs) (models.User, error) {
	log.WithField("params", params).Info("user of Handler Insert")
	var (
		user models.User
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("user of Handler Insert validator.Struct error")
		return user, err
	}
	user, err := h.user.Insert(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("user of Handler Insert h.user.Insert error")
		return user, err
	}
	return user, nil
}

// Update ...
func (h *HandlerImpl) Update(ctx context.Context, params arguments.UserUpdateArgs) (models.User, error) {
	log.WithField("params", params).Info("user of Handler Update")
	var (
		user models.User
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("user of Handler Update validator.Struct error")
		return user, err
	}
	user, err := h.user.Update(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("user of Handler Update h.user.Update error")
		return user, err
	}
	return user, nil
}

// Delete ...
func (h *HandlerImpl) Delete(ctx context.Context, params arguments.UserDeleteArgs) (int64, error) {
	log.WithField("params", params).Info("user of Handler Delete")
	var (
		id int64
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("user of Handler Delete validator.Struct error")
		return id, err
	}
	id, err := h.user.Delete(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("user of Handler Delete h.user.Delete error")
		return id, err
	}
	return id, nil
}

//go:generate mockery -name=IRepository -output=mocks -case=underscore
