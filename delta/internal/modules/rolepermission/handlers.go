package rolepermission

import (
	"context"
	"errors"

	"github.com/tanphamhaiduong/go/common/logger"
	"github.com/tanphamhaiduong/go/delta/internal/models"
	"github.com/tanphamhaiduong/go/delta/internal/utils"
)

// IRepository ...
type IRepository interface {
	ICoreRepository
	GetByRoleID(ctx context.Context, roleID int64) ([]models.RolePermission, error)
}

// HandlerImpl ...
type HandlerImpl struct {
	rolepermission IRepository
}

// NewHandler ...
func NewHandler(rolepermission IRepository) *HandlerImpl {
	return &HandlerImpl{
		rolepermission: rolepermission,
	}
}

// GetByRoleID ...
func (h *HandlerImpl) GetByRoleID(ctx context.Context, roleID int64) ([]models.RolePermission, error) {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"roleID":  roleID,
	}).Infof("Handler GetByRoleID of rolepermission")
	var (
		rolepermissions []models.RolePermission
	)
	if roleID == 0 {
		err := errors.New("Wrong role ID")
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Handler GetByRoleID validator.Struct error of rolepermission")
		return rolepermissions, err
	}
	rolepermissions, err := h.rolepermission.GetByRoleID(ctx, roleID)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Handler GetByRoleID h.rolepermission.GetByRoleID error of rolepermission")
		return rolepermissions, err
	}
	return rolepermissions, nil
}
