// @generated
package rolepermission

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
	"github.com/tanphamhaiduong/go/delta/internal/validator"
)

// ICoreRepository ...
type ICoreRepository interface {
	GetByID(ctx context.Context, params arguments.RolePermissionGetByIDArgs) (models.RolePermission, error)
	GetByIDs(ctx context.Context, params arguments.RolePermissionGetByIDsArgs) ([]models.RolePermission, error)
	List(ctx context.Context, params arguments.RolePermissionListArgs) ([]models.RolePermission, error)
	Count(ctx context.Context, params arguments.RolePermissionCountArgs) (int64, error)
	Insert(ctx context.Context, params arguments.RolePermissionInsertArgs) (models.RolePermission, error)
	Update(ctx context.Context, params arguments.RolePermissionUpdateArgs) (models.RolePermission, error)
	Delete(ctx context.Context, params arguments.RolePermissionDeleteArgs) (int64, error)
}

// GetByID ...
func (h *HandlerImpl) GetByID(ctx context.Context, params arguments.RolePermissionGetByIDArgs) (models.RolePermission, error) {
	log.WithFields(log.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Info("Handler GetByID of rolepermission")
	var (
		rolepermission models.RolePermission
	)
	if err := validator.Struct(params); err != nil {
		log.WithFields(log.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Error("Handler GetByID validator.Struct error of rolepermission")
		return rolepermission, err
	}
	rolepermission, err := h.rolepermission.GetByID(ctx, params)
	if err != nil {
		log.WithFields(log.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Error("Handler GetByID h.rolepermission.GetByID error of rolepermission")
		return rolepermission, err
	}
	return rolepermission, nil
}

// GetByIDs ...
func (h *HandlerImpl) GetByIDs(ctx context.Context, params arguments.RolePermissionGetByIDsArgs) ([]models.RolePermission, error) {
	log.WithFields(log.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Info("Handler GetByIDs of rolepermission")
	var (
		rolepermissions []models.RolePermission
	)
	if err := validator.Struct(params); err != nil {
		log.WithFields(log.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Error("Handler GetByIDs validator.Struct error of rolepermission")
		return rolepermissions, err
	}
	rolepermissions, err := h.rolepermission.GetByIDs(ctx, params)
	if err != nil {
		log.WithFields(log.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Error("Handler GetByIDs h.rolepermission.GetByIDs error of rolepermission")
		return rolepermissions, err
	}
	return rolepermissions, nil
}

// Count ...
func (h *HandlerImpl) Count(ctx context.Context, params arguments.RolePermissionCountArgs) (int64, error) {
	log.WithFields(log.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Info("Handler Count of rolepermission")
	var (
		count int64
	)
	if err := validator.Struct(params); err != nil {
		log.WithFields(log.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Error("Handler Count validator.Struct error of rolepermission")
		return count, err
	}
	count, err := h.rolepermission.Count(ctx, params)
	if err != nil {
		log.WithFields(log.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Error("Handler Count h.rolepermission.Count error of rolepermission")
		return count, err
	}
	return count, nil
}

// List ...
func (h *HandlerImpl) List(ctx context.Context, params arguments.RolePermissionListArgs) ([]models.RolePermission, error) {
	log.WithFields(log.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Info("Handler List of rolepermission")
	var (
		rolepermissions []models.RolePermission
	)
	if err := validator.Struct(params); err != nil {
		log.WithFields(log.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Error("Handler List validator.Struct error of rolepermission")
		return rolepermissions, err
	}
	rolepermissions, err := h.rolepermission.List(ctx, params)
	if err != nil {
		log.WithFields(log.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Error("Handler List h.rolepermission.List error of rolepermission")
		return rolepermissions, err
	}
	return rolepermissions, nil
}

// Insert ...
func (h *HandlerImpl) Insert(ctx context.Context, params arguments.RolePermissionInsertArgs) (models.RolePermission, error) {
	log.WithFields(log.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Info("Handler Insert of rolepermission")
	var (
		rolepermission models.RolePermission
	)
	if err := validator.Struct(params); err != nil {
		log.WithFields(log.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Error("Handler Insert validator.Struct error of rolepermission")
		return rolepermission, err
	}
	rolepermission, err := h.rolepermission.Insert(ctx, params)
	if err != nil {
		log.WithFields(log.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Error("Handler Insert h.rolepermission.Insert error of rolepermission")
		return rolepermission, err
	}
	return rolepermission, nil
}

// Update ...
func (h *HandlerImpl) Update(ctx context.Context, params arguments.RolePermissionUpdateArgs) (models.RolePermission, error) {
	log.WithFields(log.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Info("Handler Update of rolepermission")
	var (
		rolepermission models.RolePermission
	)
	if err := validator.Struct(params); err != nil {
		log.WithFields(log.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Error("Handler Update validator.Struct error of rolepermission")
		return rolepermission, err
	}
	rolepermission, err := h.rolepermission.Update(ctx, params)
	if err != nil {
		log.WithFields(log.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Error("Handler Update h.rolepermission.Update error of rolepermission")
		return rolepermission, err
	}
	return rolepermission, nil
}

// Delete ...
func (h *HandlerImpl) Delete(ctx context.Context, params arguments.RolePermissionDeleteArgs) (int64, error) {
	log.WithFields(log.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Info("Handler Delete of rolepermission")
	var (
		id int64
	)
	if err := validator.Struct(params); err != nil {
		log.WithFields(log.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Error("Handler Delete validator.Struct error of rolepermission")
		return id, err
	}
	id, err := h.rolepermission.Delete(ctx, params)
	if err != nil {
		log.WithFields(log.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Error("Handler Delete h.rolepermission.Delete error of rolepermission")
		return id, err
	}
	return id, nil
}

//go:generate mockery -name=IRepository -output=mocks -case=underscore
