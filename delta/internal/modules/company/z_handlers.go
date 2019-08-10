// @generated
package company

import (
	"context"

	"github.com/tanphamhaiduong/go/common/logger"
	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
	"github.com/tanphamhaiduong/go/delta/internal/utils"
	"github.com/tanphamhaiduong/go/delta/internal/validator"
)

// ICoreRepository ...
type ICoreRepository interface {
	GetByID(ctx context.Context, param arguments.CompanyGetByID) (models.Company, error)
	GetByIDs(ctx context.Context, param arguments.CompanyGetByIDs) ([]models.Company, error)
	List(ctx context.Context, params arguments.CompanyList) ([]models.Company, error)
	Count(ctx context.Context, params arguments.CompanyCount) (int64, error)
	Insert(ctx context.Context, params arguments.CompanyInsert) (models.Company, error)
	Update(ctx context.Context, params arguments.CompanyUpdate) (models.Company, error)
	Delete(ctx context.Context, param arguments.CompanyDelete) (int64, error)
}

// GetByID ...
func (h *HandlerImpl) GetByID(ctx context.Context, params arguments.CompanyGetByID) (models.Company, error) {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Handler GetByID of company")
	var (
		company models.Company
	)
	if err := validator.Struct(params); err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Handler GetByID validator.Struct error of company")
		return company, err
	}
	company, err := h.company.GetByID(ctx, params)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Handler GetByID h.company.GetByID error of company")
		return company, err
	}
	return company, nil
}

// GetByIDs ...
func (h *HandlerImpl) GetByIDs(ctx context.Context, params arguments.CompanyGetByIDs) ([]models.Company, error) {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Handler GetByIDs of company")
	var (
		companies []models.Company
	)
	if err := validator.Struct(params); err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Handler GetByIDs validator.Struct error of company")
		return companies, err
	}
	companies, err := h.company.GetByIDs(ctx, params)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Handler GetByIDs h.company.GetByIDs error of company")
		return companies, err
	}
	return companies, nil
}

// Count ...
func (h *HandlerImpl) Count(ctx context.Context, params arguments.CompanyCount) (int64, error) {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Handler Count of company")
	var (
		count int64
	)
	if err := validator.Struct(params); err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Handler Count validator.Struct error of company")
		return count, err
	}
	count, err := h.company.Count(ctx, params)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Handler Count h.company.Count error of company")
		return count, err
	}
	return count, nil
}

// List ...
func (h *HandlerImpl) List(ctx context.Context, params arguments.CompanyList) ([]models.Company, error) {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Handler List of company")
	var (
		companies []models.Company
	)
	if err := validator.Struct(params); err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Handler List validator.Struct error of company")
		return companies, err
	}
	companies, err := h.company.List(ctx, params)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Handler List h.company.List error of company")
		return companies, err
	}
	return companies, nil
}

// Insert ...
func (h *HandlerImpl) Insert(ctx context.Context, params arguments.CompanyInsert) (models.Company, error) {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Handler Insert of company")
	var (
		company models.Company
	)
	if err := validator.Struct(params); err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Handler Insert validator.Struct error of company")
		return company, err
	}
	company, err := h.company.Insert(ctx, params)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Handler Insert h.company.Insert error of company")
		return company, err
	}
	return company, nil
}

// Update ...
func (h *HandlerImpl) Update(ctx context.Context, params arguments.CompanyUpdate) (models.Company, error) {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Handler Update of company")
	var (
		company models.Company
	)
	if err := validator.Struct(params); err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Handler Update validator.Struct error of company")
		return company, err
	}
	company, err := h.company.Update(ctx, params)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Handler Update h.company.Update error of company")
		return company, err
	}
	return company, nil
}

// Delete ...
func (h *HandlerImpl) Delete(ctx context.Context, param arguments.CompanyDelete) (int64, error) {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"param":   param,
	}).Infof("Handler Delete of company")
	var (
		id int64
	)
	if err := validator.Struct(param); err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Handler Delete validator.Struct error of company")
		return id, err
	}
	id, err := h.company.Delete(ctx, param)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Handler Delete h.company.Delete error of company")
		return id, err
	}
	return id, nil
}

//go:generate mockery -name=IRepository -output=mocks -case=underscore
