// @generated
package permission

import (
	"context"

	"github.com/tanphamhaiduong/go/common/logger"
	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
	"github.com/tanphamhaiduong/go/delta/internal/validator"
)

// ICoreRepository ...
type ICoreRepository interface {
	GetByID(ctx context.Context, param arguments.PermissionGetByIDArgs) (models.Permission, error)
	GetByIDs(ctx context.Context, param arguments.PermissionGetByIDsArgs) ([]models.Permission, error)
	List(ctx context.Context, params arguments.PermissionListArgs) ([]models.Permission, error)
	Count(ctx context.Context, params arguments.PermissionCountArgs) (int64, error)
	Insert(ctx context.Context, params arguments.PermissionInsertArgs) (models.Permission, error)
	Update(ctx context.Context, params arguments.PermissionUpdateArgs) (models.Permission, error)
	Delete(ctx context.Context, param arguments.PermissionDeleteArgs) (int64, error)
}

// GetByID ...
func (h *HandlerImpl) GetByID(ctx context.Context, params arguments.PermissionGetByIDArgs) (models.Permission, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Handler GetByID of permission")
	var (
		permission models.Permission
	)
	if err := validator.Struct(params); err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler GetByID validator.Struct error of permission")
		return permission, err
	}
	permission, err := h.permission.GetByID(ctx, params)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler GetByID h.permission.GetByID error of permission")
		return permission, err
	}
	return permission, nil
}

// GetByIDs ...
func (h *HandlerImpl) GetByIDs(ctx context.Context, params arguments.PermissionGetByIDsArgs) ([]models.Permission, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Handler GetByIDs of permission")
	var (
		permissions []models.Permission
	)
	if err := validator.Struct(params); err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler GetByIDs validator.Struct error of permission")
		return permissions, err
	}
	permissions, err := h.permission.GetByIDs(ctx, params)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler GetByIDs h.permission.GetByIDs error of permission")
		return permissions, err
	}
	return permissions, nil
}

// Count ...
func (h *HandlerImpl) Count(ctx context.Context, params arguments.PermissionCountArgs) (int64, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Handler Count of permission")
	var (
		count int64
	)
	if err := validator.Struct(params); err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler Count validator.Struct error of permission")
		return count, err
	}
	count, err := h.permission.Count(ctx, params)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler Count h.permission.Count error of permission")
		return count, err
	}
	return count, nil
}

// List ...
func (h *HandlerImpl) List(ctx context.Context, params arguments.PermissionListArgs) ([]models.Permission, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Handler List of permission")
	var (
		permissions []models.Permission
	)
	if err := validator.Struct(params); err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler List validator.Struct error of permission")
		return permissions, err
	}
	permissions, err := h.permission.List(ctx, params)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler List h.permission.List error of permission")
		return permissions, err
	}
	return permissions, nil
}

// Insert ...
func (h *HandlerImpl) Insert(ctx context.Context, params arguments.PermissionInsertArgs) (models.Permission, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Handler Insert of permission")
	var (
		permission models.Permission
	)
	if err := validator.Struct(params); err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler Insert validator.Struct error of permission")
		return permission, err
	}
	permission, err := h.permission.Insert(ctx, params)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler Insert h.permission.Insert error of permission")
		return permission, err
	}
	return permission, nil
}

// Update ...
func (h *HandlerImpl) Update(ctx context.Context, params arguments.PermissionUpdateArgs) (models.Permission, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Handler Update of permission")
	var (
		permission models.Permission
	)
	if err := validator.Struct(params); err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler Update validator.Struct error of permission")
		return permission, err
	}
	permission, err := h.permission.Update(ctx, params)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler Update h.permission.Update error of permission")
		return permission, err
	}
	return permission, nil
}

// Delete ...
func (h *HandlerImpl) Delete(ctx context.Context, param arguments.PermissionDeleteArgs) (int64, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"param":   param,
	}).Infof("Handler Delete of permission")
	var (
		id int64
	)
	if err := validator.Struct(param); err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler Delete validator.Struct error of permission")
		return id, err
	}
	id, err := h.permission.Delete(ctx, param)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler Delete h.permission.Delete error of permission")
		return id, err
	}
	return id, nil
}

//go:generate mockery -name=IRepository -output=mocks -case=underscore
