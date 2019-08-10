// @generated
package modules

import (
	"context"

	"github.com/graphql-go/graphql"
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
}

// ICoreCompanyResolver ...
type ICoreCompanyResolver interface {
	ForwardParams(p graphql.ResolveParams) (interface{}, error)
	GetByID(p graphql.ResolveParams) (interface{}, error)
	Count(p graphql.ResolveParams) (interface{}, error)
	List(p graphql.ResolveParams) (interface{}, error)
	Insert(p graphql.ResolveParams) (interface{}, error)
	Update(p graphql.ResolveParams) (interface{}, error)
}

//go:generate mockery -name=ICompanyResolver -output=mocks -case=underscore
//go:generate mockery -name=ICompanyHandler -output=mocks -case=underscore
