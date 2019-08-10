// @generated
package rolepermission

import (
	"context"

	"github.com/tanphamhaiduong/go/common/logger"
	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
	"github.com/tanphamhaiduong/go/delta/internal/validator"
)

// ICoreRepository ...
type ICoreRepository interface {
	GetByID(ctx context.Context, param arguments.RolePermissionGetByID) (models.RolePermission, error)
	GetByIDs(ctx context.Context, param arguments.RolePermissionGetByIDs) ([]models.RolePermission, error)
	List(ctx context.Context, params arguments.RolePermissionList) ([]models.RolePermission, error)
	Count(ctx context.Context, params arguments.RolePermissionCount) (int64, error)
	Insert(ctx context.Context, params arguments.RolePermissionInsert) (models.RolePermission, error)
	Update(ctx context.Context, params arguments.RolePermissionUpdate) (models.RolePermission, error)
	Delete(ctx context.Context, param arguments.RolePermissionDelete) (int64, error)
}

// GetByID ...
func (h *HandlerImpl) GetByID(ctx context.Context, params arguments.RolePermissionGetByID) (models.RolePermission, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Handler GetByID of rolepermission")
	var (
		rolepermission models.RolePermission
	)
	if err := validator.Struct(params); err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler GetByID validator.Struct error of rolepermission")
		return rolepermission, err
	}
	rolepermission, err := h.rolepermission.GetByID(ctx, params)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler GetByID h.rolepermission.GetByID error of rolepermission")
		return rolepermission, err
	}
	return rolepermission, nil
}

// GetByIDs ...
func (h *HandlerImpl) GetByIDs(ctx context.Context, params arguments.RolePermissionGetByIDs) ([]models.RolePermission, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Handler GetByIDs of rolepermission")
	var (
		rolepermissions []models.RolePermission
	)
	if err := validator.Struct(params); err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler GetByIDs validator.Struct error of rolepermission")
		return rolepermissions, err
	}
	rolepermissions, err := h.rolepermission.GetByIDs(ctx, params)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler GetByIDs h.rolepermission.GetByIDs error of rolepermission")
		return rolepermissions, err
	}
	return rolepermissions, nil
}

// Count ...
func (h *HandlerImpl) Count(ctx context.Context, params arguments.RolePermissionCount) (int64, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Handler Count of rolepermission")
	var (
		count int64
	)
	if err := validator.Struct(params); err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler Count validator.Struct error of rolepermission")
		return count, err
	}
	count, err := h.rolepermission.Count(ctx, params)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler Count h.rolepermission.Count error of rolepermission")
		return count, err
	}
	return count, nil
}

// List ...
func (h *HandlerImpl) List(ctx context.Context, params arguments.RolePermissionList) ([]models.RolePermission, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Handler List of rolepermission")
	var (
		rolepermissions []models.RolePermission
	)
	if err := validator.Struct(params); err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler List validator.Struct error of rolepermission")
		return rolepermissions, err
	}
	rolepermissions, err := h.rolepermission.List(ctx, params)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler List h.rolepermission.List error of rolepermission")
		return rolepermissions, err
	}
	return rolepermissions, nil
}

// Insert ...
func (h *HandlerImpl) Insert(ctx context.Context, params arguments.RolePermissionInsert) (models.RolePermission, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Handler Insert of rolepermission")
	var (
		rolepermission models.RolePermission
	)
	if err := validator.Struct(params); err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler Insert validator.Struct error of rolepermission")
		return rolepermission, err
	}
	rolepermission, err := h.rolepermission.Insert(ctx, params)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler Insert h.rolepermission.Insert error of rolepermission")
		return rolepermission, err
	}
	return rolepermission, nil
}

// Update ...
func (h *HandlerImpl) Update(ctx context.Context, params arguments.RolePermissionUpdate) (models.RolePermission, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Handler Update of rolepermission")
	var (
		rolepermission models.RolePermission
	)
	if err := validator.Struct(params); err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler Update validator.Struct error of rolepermission")
		return rolepermission, err
	}
	rolepermission, err := h.rolepermission.Update(ctx, params)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler Update h.rolepermission.Update error of rolepermission")
		return rolepermission, err
	}
	return rolepermission, nil
}

// Delete ...
func (h *HandlerImpl) Delete(ctx context.Context, param arguments.RolePermissionDelete) (int64, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"param":   param,
	}).Infof("Handler Delete of rolepermission")
	var (
		id int64
	)
	if err := validator.Struct(param); err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler Delete validator.Struct error of rolepermission")
		return id, err
	}
	id, err := h.rolepermission.Delete(ctx, param)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Handler Delete h.rolepermission.Delete error of rolepermission")
		return id, err
	}
	return id, nil
}

//go:generate mockery -name=IRepository -output=mocks -case=underscore
