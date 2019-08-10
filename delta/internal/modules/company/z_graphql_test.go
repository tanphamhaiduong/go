// @generated
package company

import (
	"context"
	"errors"

	"github.com/graphql-go/graphql"

	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
)

func (s *CompanyResolverTestSuite) TestForwardParams_Success() {
	var (
		sampleID int64 = 1
		params         = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{"id": sampleID}}
		expected       = map[string]interface{}(map[string]interface{}{"id": sampleID})
	)
	actual, err := s.Company.ForwardParams(params)
	s.Nil(err)
	s.Equal(expected, actual)
}

func (s *CompanyResolverTestSuite) TestGetByID_Success() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{"id": sampleID}}
		company  models.Company
		args     = arguments.CompanyGetByIDArgs{
			ID: 1,
		}
	)
	s.MockICompany.On("GetByID", ctx, args).Return(company, nil)
	actual, err := s.Company.GetByID(params)
	s.Nil(err)
	s.Equal(company, actual)
}

func (s *CompanyResolverTestSuite) TestGetByID_Fail() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{"id": sampleID}}
		company  models.Company
		args     = arguments.CompanyGetByIDArgs{
			ID: 1,
		}
	)
	s.MockICompany.On("GetByID", ctx, args).Return(company, errors.New("some errors"))
	actual, err := s.Company.GetByID(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *CompanyResolverTestSuite) TestList_Success() {
	var (
		ctx       = context.Background()
		params    = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{"page": 1, "pageSize": 10}, Args: map[string]interface{}{}}
		companies []models.Company
		args      = arguments.CompanyListArgs{
			Page:     1,
			PageSize: 10,
		}
	)
	s.MockICompany.On("List", ctx, args).Return(companies, nil)
	actual, err := s.Company.List(params)
	s.Nil(err)
	s.Equal(companies, actual)
}

func (s *CompanyResolverTestSuite) TestList_Fail() {
	var (
		ctx       = context.Background()
		params    = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{"page": 1, "pageSize": 10}, Args: map[string]interface{}{}}
		companies []models.Company
		args      = arguments.CompanyListArgs{
			Page:     1,
			PageSize: 10,
		}
	)
	s.MockICompany.On("List", ctx, args).Return(companies, errors.New("some errors"))
	actual, err := s.Company.List(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *CompanyResolverTestSuite) TestCount_Success() {
	var (
		ctx    = context.Background()
		params = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args   = arguments.CompanyCountArgs{}
		count  int64
	)
	s.MockICompany.On("Count", ctx, args).Return(count, nil)
	actual, err := s.Company.Count(params)
	s.Nil(err)
	s.Equal(count, actual)
}

func (s *CompanyResolverTestSuite) TestCount_Fail() {
	var (
		ctx    = context.Background()
		params = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args   = arguments.CompanyCountArgs{}
		count  int64
	)
	s.MockICompany.On("Count", ctx, args).Return(count, errors.New("some errors"))
	actual, err := s.Company.Count(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *CompanyResolverTestSuite) TestInsert_Success() {
	var (
		ctx     = context.Background()
		params  = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args    = arguments.CompanyInsertArgs{}
		company models.Company
	)
	// Mock Insert
	s.MockICompany.On("Insert", ctx, args).Return(company, nil)
	// Mock GetByID
	s.MockICompany.On("GetByID", ctx, params).Return(company, nil)

	actual, err := s.Company.Insert(params)
	s.Nil(err)
	s.Equal(company, actual)
}

func (s *CompanyResolverTestSuite) TestInsert_Fail() {
	var (
		ctx     = context.Background()
		params  = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args    = arguments.CompanyInsertArgs{}
		company models.Company
	)
	s.MockICompany.On("Insert", ctx, args).Return(company, errors.New("some errors"))
	actual, err := s.Company.Insert(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *CompanyResolverTestSuite) TestUpdate_Success() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = graphql.ResolveParams{
			Context: context.Background(),
			Source:  map[string]interface{}{},
			Args:    map[string]interface{}{"id": 1},
		}
		args = arguments.CompanyUpdateArgs{
			ID: &sampleID,
		}
		company = models.Company{}
	)
	// Mock for Update
	s.MockICompany.On("Update", ctx, args).Return(company, nil)
	// Mock for LoadByID
	s.MockICompany.On("GetByID", ctx, args).Return(company, nil)
	actual, err := s.Company.Update(params)
	s.Nil(err)
	s.Equal(company, actual)
}

func (s *CompanyResolverTestSuite) TestUpdate_Fail() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = graphql.ResolveParams{
			Context: context.Background(),
			Source:  map[string]interface{}{},
			Args:    map[string]interface{}{"id": 1},
		}
		args = arguments.CompanyUpdateArgs{
			ID: &sampleID,
		}
		company models.Company
	)
	s.MockICompany.On("Update", ctx, args).Return(company, errors.New("some errors"))
	actual, err := s.Company.Update(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *CompanyResolverTestSuite) TestDelete_Success() {
	var (
		ctx    = context.Background()
		params = graphql.ResolveParams{
			Context: context.Background(),
			Source:  map[string]interface{}{},
			Args:    map[string]interface{}{"id": 1},
		}
		args = arguments.CompanyDeleteArgs{
			ID: 1,
		}
		sampleID int64 = 1
	)
	s.MockICompany.On("Delete", ctx, args).Return(sampleID, nil)
	actual, err := s.Company.Delete(params)
	s.Nil(err)
	s.Equal(sampleID, actual)
}

func (s *CompanyResolverTestSuite) TestDelete_Fail() {
	var (
		ctx    = context.Background()
		params = graphql.ResolveParams{
			Context: context.Background(),
			Source:  map[string]interface{}{},
			Args:    map[string]interface{}{"id": 1},
		}
		args = arguments.CompanyDeleteArgs{
			ID: 1,
		}
		sampleID int64
	)
	s.MockICompany.On("Delete", ctx, args).Return(sampleID, errors.New("some errors"))
	actual, err := s.Company.Delete(params)
	s.Nil(actual)
	s.NotNil(err)
}
