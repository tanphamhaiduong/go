// @generated

package modules

import (
	"context"

	"github.com/tanphamhaiduong/go/delta/server/arguments"
	"github.com/tanphamhaiduong/go/delta/server/models"
)

// ICoreCompanyHandler ...
type ICoreCompanyHandler interface {
	GetByID(ctx context.Context, params arguments.CompanyGetByIDArgs) (models.Company, error)
	Count(ctx context.Context, params arguments.CompanyCountArgs) (int64, error)
	List(ctx context.Context, params arguments.CompanyListArgs) ([]models.Company, error)
	Insert(ctx context.Context, params arguments.CompanyInsertArgs) (models.Company, error)
	Update(ctx context.Context, params arguments.CompanyUpdateArgs) (models.Company, error)
	Delete(ctx context.Context, params arguments.CompanyDeleteArgs) (int64, error)
}

// ICoreCompanyResolver ...
type ICoreCompanyResolver interface {
	IResolver
}

//go:generate mockery -name=ICompanyResolver -output=mocks -case=underscore
//go:generate mockery -name=ICompanyHandler -output=mocks -case=underscore
