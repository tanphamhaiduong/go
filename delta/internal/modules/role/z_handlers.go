// @generated
package role

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
	"github.com/tanphamhaiduong/go/delta/internal/validator"
)

// ICoreRepository ...
type ICoreRepository interface {
	GetByID(ctx context.Context, params arguments.RoleGetByIDArgs) (models.Role, error)
	GetByIDs(ctx context.Context, params arguments.RoleGetByIDsArgs) ([]models.Role, error)
	List(ctx context.Context, params arguments.RoleListArgs) ([]models.Role, error)
	Count(ctx context.Context, params arguments.RoleCountArgs) (int64, error)
	Insert(ctx context.Context, params arguments.RoleInsertArgs) (models.Role, error)
	Update(ctx context.Context, params arguments.RoleUpdateArgs) (models.Role, error)
	Delete(ctx context.Context, params arguments.RoleDeleteArgs) (int64, error)
}

// GetByID ...
func (h *HandlerImpl) GetByID(ctx context.Context, params arguments.RoleGetByIDArgs) (models.Role, error) {
	log.WithField("params", params).Info("Handler GetByID of role")
	var (
		role models.Role
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("Handler GetByID validator.Struct error of role")
		return role, err
	}
	role, err := h.role.GetByID(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("Handler GetByID h.role.GetByID error of role")
		return role, err
	}
	return role, nil
}

// GetByIDs ...
func (h *HandlerImpl) GetByIDs(ctx context.Context, params arguments.RoleGetByIDsArgs) ([]models.Role, error) {
	log.WithField("params", params).Info("Handler GetByIDs of role")
	var (
		roles []models.Role
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("Handler GetByIDs validator.Struct error of role")
		return roles, err
	}
	roles, err := h.role.GetByIDs(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("Handler GetByIDs h.role.GetByIDs error of role")
		return roles, err
	}
	return roles, nil
}

// Count ...
func (h *HandlerImpl) Count(ctx context.Context, params arguments.RoleCountArgs) (int64, error) {
	log.WithField("params", params).Info("Handler Count of role")
	var (
		count int64
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("Handler Count validator.Struct error of role")
		return count, err
	}
	count, err := h.role.Count(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("Handler Count h.role.Count error of role")
		return count, err
	}
	return count, nil
}

// List ...
func (h *HandlerImpl) List(ctx context.Context, params arguments.RoleListArgs) ([]models.Role, error) {
	log.WithField("params", params).Info("Handler List of role")
	var (
		roles []models.Role
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("Handler List validator.Struct error of role")
		return roles, err
	}
	roles, err := h.role.List(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("Handler List h.role.List error of role")
		return roles, err
	}
	return roles, nil
}

// Insert ...
func (h *HandlerImpl) Insert(ctx context.Context, params arguments.RoleInsertArgs) (models.Role, error) {
	log.WithField("params", params).Info("Handler Insert of role")
	var (
		role models.Role
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("Handler Insert validator.Struct error of role")
		return role, err
	}
	role, err := h.role.Insert(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("Handler Insert h.role.Insert error of role")
		return role, err
	}
	return role, nil
}

// Update ...
func (h *HandlerImpl) Update(ctx context.Context, params arguments.RoleUpdateArgs) (models.Role, error) {
	log.WithField("params", params).Info("Handler Update of role")
	var (
		role models.Role
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("Handler Update validator.Struct error of role")
		return role, err
	}
	role, err := h.role.Update(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("Handler Update h.role.Update error of role")
		return role, err
	}
	return role, nil
}

// Delete ...
func (h *HandlerImpl) Delete(ctx context.Context, params arguments.RoleDeleteArgs) (int64, error) {
	log.WithField("params", params).Info("Handler Delete of role")
	var (
		id int64
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("Handler Delete validator.Struct error of role")
		return id, err
	}
	id, err := h.role.Delete(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("Handler Delete h.role.Delete error of role")
		return id, err
	}
	return id, nil
}

//go:generate mockery -name=IRepository -output=mocks -case=underscore
