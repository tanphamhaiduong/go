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
	log.WithFields(log.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Info("Handler GetByID of company")
	var (
		company models.Company
	)
	if err := validator.Struct(params); err != nil {
		log.WithFields(log.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Error("Handler GetByID validator.Struct error of company")
		return company, err
	}
	company, err := h.company.GetByID(ctx, params)
	if err != nil {
		log.WithFields(log.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Error("Handler GetByID h.company.GetByID error of company")
		return company, err
	}
	return company, nil
}

// GetByIDs ...
func (h *HandlerImpl) GetByIDs(ctx context.Context, params arguments.CompanyGetByIDsArgs) ([]models.Company, error) {
	log.WithFields(log.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Info("Handler GetByIDs of company")
	var (
		companies []models.Company
	)
	if err := validator.Struct(params); err != nil {
		log.WithFields(log.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Error("Handler GetByIDs validator.Struct error of company")
		return companies, err
	}
	companies, err := h.company.GetByIDs(ctx, params)
	if err != nil {
		log.WithFields(log.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Error("Handler GetByIDs h.company.GetByIDs error of company")
		return companies, err
	}
	return companies, nil
}

// Count ...
func (h *HandlerImpl) Count(ctx context.Context, params arguments.CompanyCountArgs) (int64, error) {
	log.WithFields(log.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Info("Handler Count of company")
	var (
		count int64
	)
	if err := validator.Struct(params); err != nil {
		log.WithFields(log.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Error("Handler Count validator.Struct error of company")
		return count, err
	}
	count, err := h.company.Count(ctx, params)
	if err != nil {
		log.WithFields(log.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Error("Handler Count h.company.Count error of company")
		return count, err
	}
	return count, nil
}

// List ...
func (h *HandlerImpl) List(ctx context.Context, params arguments.CompanyListArgs) ([]models.Company, error) {
	log.WithFields(log.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Info("Handler List of company")
	var (
		companies []models.Company
	)
	if err := validator.Struct(params); err != nil {
		log.WithFields(log.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Error("Handler List validator.Struct error of company")
		return companies, err
	}
	companies, err := h.company.List(ctx, params)
	if err != nil {
		log.WithFields(log.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Error("Handler List h.company.List error of company")
		return companies, err
	}
	return companies, nil
}

// Insert ...
func (h *HandlerImpl) Insert(ctx context.Context, params arguments.CompanyInsertArgs) (models.Company, error) {
	log.WithFields(log.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Info("Handler Insert of company")
	var (
		company models.Company
	)
	if err := validator.Struct(params); err != nil {
		log.WithFields(log.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Error("Handler Insert validator.Struct error of company")
		return company, err
	}
	company, err := h.company.Insert(ctx, params)
	if err != nil {
		log.WithFields(log.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Error("Handler Insert h.company.Insert error of company")
		return company, err
	}
	return company, nil
}

// Update ...
func (h *HandlerImpl) Update(ctx context.Context, params arguments.CompanyUpdateArgs) (models.Company, error) {
	log.WithFields(log.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Info("Handler Update of company")
	var (
		company models.Company
	)
	if err := validator.Struct(params); err != nil {
		log.WithFields(log.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Error("Handler Update validator.Struct error of company")
		return company, err
	}
	company, err := h.company.Update(ctx, params)
	if err != nil {
		log.WithFields(log.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Error("Handler Update h.company.Update error of company")
		return company, err
	}
	return company, nil
}

// Delete ...
func (h *HandlerImpl) Delete(ctx context.Context, params arguments.CompanyDeleteArgs) (int64, error) {
	log.WithFields(log.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Info("Handler Delete of company")
	var (
		id int64
	)
	if err := validator.Struct(params); err != nil {
		log.WithFields(log.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Error("Handler Delete validator.Struct error of company")
		return id, err
	}
	id, err := h.company.Delete(ctx, params)
	if err != nil {
		log.WithFields(log.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Error("Handler Delete h.company.Delete error of company")
		return id, err
	}
	return id, nil
}

//go:generate mockery -name=IRepository -output=mocks -case=underscore
