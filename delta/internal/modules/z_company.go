// @generated
package modules

import (
	"context"

	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
)

// ICoreCompanyHandler ...
type ICoreCompanyHandler interface {
	GetByID(ctx context.Context, params arguments.CompanyGetByID) (models.Company, error)
	Count(ctx context.Context, params arguments.CompanyCount) (int64, error)
	List(ctx context.Context, params arguments.CompanyList) ([]models.Company, error)
	Insert(ctx context.Context, params arguments.CompanyInsert) (models.Company, error)
	Update(ctx context.Context, params arguments.CompanyUpdate) (models.Company, error)
	Delete(ctx context.Context, params arguments.CompanyDelete) (int64, error)
}

// ICoreCompanyResolver ...
type ICoreCompanyResolver interface {
	IResolver
}

//go:generate mockery -name=ICompanyResolver -output=mocks -case=underscore
//go:generate mockery -name=ICompanyHandler -output=mocks -case=underscore
