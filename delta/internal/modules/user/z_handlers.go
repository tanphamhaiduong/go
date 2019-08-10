// @generated
package user

import (
	"context"

	"github.com/tanphamhaiduong/go/common/logger"
	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
	"github.com/tanphamhaiduong/go/delta/internal/validator"
)

// ICoreRepository ...
type ICoreRepository interface {
	GetByID(ctx context.Context, param arguments.UserGetByID) (models.User, error)
	GetByIDs(ctx context.Context, param arguments.UserGetByIDs) ([]models.User, error)
	List(ctx context.Context, params arguments.UserList) ([]models.User, error)
	Count(ctx context.Context, params arguments.UserCount) (int64, error)
	Insert(ctx context.Context, params arguments.UserInsert) (models.User, error)
	Update(ctx context.Context, params arguments.UserUpdate) (models.User, error)
	Delete(ctx context.Context, param arguments.UserDelete) (int64, error)
}

// GetByID ...
func (h *HandlerImpl) GetByID(ctx context.Context, params arguments.UserGetByID) (models.User, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Handler GetByID of user")
	var (
		user models.User
	)
	if err := validator.Struct(params); err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler GetByID validator.Struct error of user")
		return user, err
	}
	user, err := h.user.GetByID(ctx, params)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler GetByID h.user.GetByID error of user")
		return user, err
	}
	return user, nil
}

// GetByIDs ...
func (h *HandlerImpl) GetByIDs(ctx context.Context, params arguments.UserGetByIDs) ([]models.User, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Handler GetByIDs of user")
	var (
		users []models.User
	)
	if err := validator.Struct(params); err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler GetByIDs validator.Struct error of user")
		return users, err
	}
	users, err := h.user.GetByIDs(ctx, params)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler GetByIDs h.user.GetByIDs error of user")
		return users, err
	}
	return users, nil
}

// Count ...
func (h *HandlerImpl) Count(ctx context.Context, params arguments.UserCount) (int64, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Handler Count of user")
	var (
		count int64
	)
	if err := validator.Struct(params); err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler Count validator.Struct error of user")
		return count, err
	}
	count, err := h.user.Count(ctx, params)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler Count h.user.Count error of user")
		return count, err
	}
	return count, nil
}

// List ...
func (h *HandlerImpl) List(ctx context.Context, params arguments.UserList) ([]models.User, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Handler List of user")
	var (
		users []models.User
	)
	if err := validator.Struct(params); err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler List validator.Struct error of user")
		return users, err
	}
	users, err := h.user.List(ctx, params)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler List h.user.List error of user")
		return users, err
	}
	return users, nil
}

// Insert ...
func (h *HandlerImpl) Insert(ctx context.Context, params arguments.UserInsert) (models.User, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Handler Insert of user")
	var (
		user models.User
	)
	if err := validator.Struct(params); err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler Insert validator.Struct error of user")
		return user, err
	}
	user, err := h.user.Insert(ctx, params)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler Insert h.user.Insert error of user")
		return user, err
	}
	return user, nil
}

// Update ...
func (h *HandlerImpl) Update(ctx context.Context, params arguments.UserUpdate) (models.User, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Handler Update of user")
	var (
		user models.User
	)
	if err := validator.Struct(params); err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler Update validator.Struct error of user")
		return user, err
	}
	user, err := h.user.Update(ctx, params)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler Update h.user.Update error of user")
		return user, err
	}
	return user, nil
}

// Delete ...
func (h *HandlerImpl) Delete(ctx context.Context, param arguments.UserDelete) (int64, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"param":   param,
	}).Infof("Handler Delete of user")
	var (
		id int64
	)
	if err := validator.Struct(param); err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler Delete validator.Struct error of user")
		return id, err
	}
	id, err := h.user.Delete(ctx, param)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler Delete h.user.Delete error of user")
		return id, err
	}
	return id, nil
}

//go:generate mockery -name=IRepository -output=mocks -case=underscore
