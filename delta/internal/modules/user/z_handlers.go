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
	log.WithField("params", params).Info("Handler GetByID of user")
	var (
		user models.User
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("Handler GetByID validator.Struct error of user")
		return user, err
	}
	user, err := h.user.GetByID(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("Handler GetByID h.user.GetByID error of user")
		return user, err
	}
	return user, nil
}

// GetByIDs ...
func (h *HandlerImpl) GetByIDs(ctx context.Context, params arguments.UserGetByIDsArgs) ([]models.User, error) {
	log.WithField("params", params).Info("Handler GetByIDs of user")
	var (
		users []models.User
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("Handler GetByIDs validator.Struct error of user")
		return users, err
	}
	users, err := h.user.GetByIDs(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("Handler GetByIDs h.user.GetByIDs error of user")
		return users, err
	}
	return users, nil
}

// Count ...
func (h *HandlerImpl) Count(ctx context.Context, params arguments.UserCountArgs) (int64, error) {
	log.WithField("params", params).Info("Handler Count of user")
	var (
		count int64
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("Handler Count validator.Struct error of user")
		return count, err
	}
	count, err := h.user.Count(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("Handler Count h.user.Count error of user")
		return count, err
	}
	return count, nil
}

// List ...
func (h *HandlerImpl) List(ctx context.Context, params arguments.UserListArgs) ([]models.User, error) {
	log.WithField("params", params).Info("Handler List of user")
	var (
		users []models.User
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("Handler List validator.Struct error of user")
		return users, err
	}
	users, err := h.user.List(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("Handler List h.user.List error of user")
		return users, err
	}
	return users, nil
}

// Insert ...
func (h *HandlerImpl) Insert(ctx context.Context, params arguments.UserInsertArgs) (models.User, error) {
	log.WithField("params", params).Info("Handler Insert of user")
	var (
		user models.User
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("Handler Insert validator.Struct error of user")
		return user, err
	}
	user, err := h.user.Insert(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("Handler Insert h.user.Insert error of user")
		return user, err
	}
	return user, nil
}

// Update ...
func (h *HandlerImpl) Update(ctx context.Context, params arguments.UserUpdateArgs) (models.User, error) {
	log.WithField("params", params).Info("Handler Update of user")
	var (
		user models.User
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("Handler Update validator.Struct error of user")
		return user, err
	}
	user, err := h.user.Update(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("Handler Update h.user.Update error of user")
		return user, err
	}
	return user, nil
}

// Delete ...
func (h *HandlerImpl) Delete(ctx context.Context, params arguments.UserDeleteArgs) (int64, error) {
	log.WithField("params", params).Info("Handler Delete of user")
	var (
		id int64
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("Handler Delete validator.Struct error of user")
		return id, err
	}
	id, err := h.user.Delete(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("Handler Delete h.user.Delete error of user")
		return id, err
	}
	return id, nil
}

//go:generate mockery -name=IRepository -output=mocks -case=underscore
