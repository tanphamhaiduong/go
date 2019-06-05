// @generated
package company

import (
	"context"

	"github.com/tanphamhaiduong/go/delta/server/arguments"
	"github.com/tanphamhaiduong/go/delta/server/models"
	"github.com/tanphamhaiduong/go/delta/server/validator"
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
func (h HandlerImpl) GetByID(ctx context.Context, params arguments.CompanyGetByIDArgs) (models.Company, error) {
	var (
		company = models.Company{}
	)
	company, err := h.company.GetByID(ctx, params)
	if err != nil {
		return company, err
	}
	return company, nil
}

// GetByIDs ...
func (h HandlerImpl) GetByIDs(ctx context.Context, params arguments.CompanyGetByIDsArgs) ([]models.Company, error) {
	var (
		companies = []models.Company{}
	)
	companies, err := h.company.GetByIDs(ctx, params)
	if err != nil {
		return companies, err
	}
	return companies, nil
}

// Count ...
func (h HandlerImpl) Count(ctx context.Context, params arguments.CompanyCountArgs) (int64, error) {
	var (
		count int64
	)
	if err := validator.Struct(params); err != nil {
		return count, err
	}
	count, err := h.company.Count(ctx, params)
	if err != nil {
		return count, err
	}
	return count, nil
}

// List ...
func (h HandlerImpl) List(ctx context.Context, params arguments.CompanyListArgs) ([]models.Company, error) {
	var (
		companies = []models.Company{}
	)
	if err := validator.Struct(params); err != nil {
		return companies, err
	}
	companies, err := h.company.List(ctx, params)
	if err != nil {
		return companies, err
	}
	return companies, nil
}

// Insert ...
func (h HandlerImpl) Insert(ctx context.Context, params arguments.CompanyInsertArgs) (models.Company, error) {
	var (
		company = models.Company{}
	)
	if err := validator.Struct(params); err != nil {
		return company, err
	}
	company, err := h.company.Insert(ctx, params)
	if err != nil {
		return company, err
	}
	return company, nil
}

// Update ...
func (h HandlerImpl) Update(ctx context.Context, params arguments.CompanyUpdateArgs) (models.Company, error) {
	var (
		company = models.Company{}
	)
	if err := validator.Struct(params); err != nil {
		return company, err
	}
	company, err := h.company.Update(ctx, params)
	if err != nil {
		return company, err
	}
	return company, nil
}

// Delete ...
func (h HandlerImpl) Delete(ctx context.Context, params arguments.CompanyDeleteArgs) (int64, error) {
	var (
		id int64
	)
	if err := validator.Struct(params); err != nil {
		return id, err
	}
	id, err := h.company.Delete(ctx, params)
	if err != nil {
		return id, err
	}
	return id, nil
}

//go:generate mockery -name=IRepository -output=mocks -case=underscore
