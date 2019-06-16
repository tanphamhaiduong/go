// @generated
package company

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
	"github.com/tanphamhaiduong/go/delta/internal/validator"
)

// ICoreRepository ...
type ICoreRepository interface {
	GetByID(ctx context.Context, params arguments.CompanyGetByIDArgs) (models.Company, error)
	GetByIDs(ctx context.Context, params arguments.CompanyGetByIDsArgs) ([]models.Company, error)
	List(ctx context.Context, params arguments.CompanyListArgs) ([]models.Company, error)
	Count(ctx context.Context, params arguments.CompanyCountArgs) (int64, error)
	Insert(ctx context.Context, params arguments.CompanyInsertArgs) (models.Company, error)
	Update(ctx context.Context, params arguments.CompanyUpdateArgs) (models.Company, error)
	Delete(ctx context.Context, params arguments.CompanyDeleteArgs) (int64, error)
}

// GetByID ...
func (h *HandlerImpl) GetByID(ctx context.Context, params arguments.CompanyGetByIDArgs) (models.Company, error) {
	log.WithField("params", params).Info("company of Handler GetByID")
	var (
		company models.Company
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("company of Handler GetByID validator.Struct error")
		return company, err
	}
	company, err := h.company.GetByID(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("company of Handler GetByID h.company.GetByID error")
		return company, err
	}
	return company, nil
}

// GetByIDs ...
func (h *HandlerImpl) GetByIDs(ctx context.Context, params arguments.CompanyGetByIDsArgs) ([]models.Company, error) {
	log.WithField("params", params).Info("company of Handler GetByIDs")
	var (
		companies []models.Company
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("company of Handler GetByIDs validator.Struct error")
		return companies, err
	}
	companies, err := h.company.GetByIDs(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("company of Handler GetByIDs h.company.GetByIDs error")
		return companies, err
	}
	return companies, nil
}

// Count ...
func (h *HandlerImpl) Count(ctx context.Context, params arguments.CompanyCountArgs) (int64, error) {
	log.WithField("params", params).Info("company of Handler Count")
	var (
		count int64
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("company of Handler Count validator.Struct error")
		return count, err
	}
	count, err := h.company.Count(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("company of Handler Count h.company.Count error")
		return count, err
	}
	return count, nil
}

// List ...
func (h *HandlerImpl) List(ctx context.Context, params arguments.CompanyListArgs) ([]models.Company, error) {
	log.WithField("params", params).Info("company of Handler List")
	var (
		companies []models.Company
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("company of Handler List validator.Struct error")
		return companies, err
	}
	companies, err := h.company.List(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("company of Handler List h.company.List error")
		return companies, err
	}
	return companies, nil
}

// Insert ...
func (h *HandlerImpl) Insert(ctx context.Context, params arguments.CompanyInsertArgs) (models.Company, error) {
	log.WithField("params", params).Info("company of Handler Insert")
	var (
		company models.Company
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("company of Handler Insert validator.Struct error")
		return company, err
	}
	company, err := h.company.Insert(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("company of Handler Insert h.company.Insert error")
		return company, err
	}
	return company, nil
}

// Update ...
func (h *HandlerImpl) Update(ctx context.Context, params arguments.CompanyUpdateArgs) (models.Company, error) {
	log.WithField("params", params).Info("company of Handler Update")
	var (
		company models.Company
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("company of Handler Update validator.Struct error")
		return company, err
	}
	company, err := h.company.Update(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("company of Handler Update h.company.Update error")
		return company, err
	}
	return company, nil
}

// Delete ...
func (h *HandlerImpl) Delete(ctx context.Context, params arguments.CompanyDeleteArgs) (int64, error) {
	log.WithField("params", params).Info("company of Handler Delete")
	var (
		id int64
	)
	if err := validator.Struct(params); err != nil {
		log.WithField("Error", err).Error("company of Handler Delete validator.Struct error")
		return id, err
	}
	id, err := h.company.Delete(ctx, params)
	if err != nil {
		log.WithField("Error", err).Error("company of Handler Delete h.company.Delete error")
		return id, err
	}
	return id, nil
}

//go:generate mockery -name=IRepository -output=mocks -case=underscore
