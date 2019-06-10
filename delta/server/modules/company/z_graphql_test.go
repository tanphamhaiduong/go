// @generated
package company

import (
	"context"
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/stretchr/testify/mock"

	"github.com/tanphamhaiduong/go/delta/server/models"
)

func (s *CompanyResolverTestSuite) TestGetByID_Success() {
	var (
		ctx           = context.Background()
		args    int64 = 1
		params        = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{"id": args}}
		company models.Company
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryRowContext", ctx, args).Return(s.MockIRow)
	s.MockIRow.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	s.MockIHandlerCompany.On("GetByID", ctx, params).Return(company, nil)
	s.MockIRepositoryCompany.On("GetByID", ctx, params).Return(company, nil)
	actual, err := s.Company.GetByID(params)
	s.Nil(err)
	s.Equal(company, actual)
}

func (s *CompanyResolverTestSuite) TestGetByID_Fail() {
	var (
		ctx           = context.Background()
		args    int64 = 1
		params        = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{"id": args}}
		company models.Company
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryRowContext", ctx, args).Return(s.MockIRow)
	s.MockIRow.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("some errors"))
	s.MockIHandlerCompany.On("GetByID", ctx, params).Return(company, errors.New("some errors"))
	s.MockIRepositoryCompany.On("GetByID", ctx, params).Return(company, errors.New("some errors"))
	actual, err := s.Company.GetByID(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *CompanyResolverTestSuite) TestList_Success() {
	var (
		ctx       = context.Background()
		params    = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{"page": 1, "pageSize": 10}, Args: map[string]interface{}{}}
		companies []models.Company
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryContext", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(s.MockIRows, nil)
	s.MockIRows.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockIHandlerCompany.On("List", ctx, params).Return(companies, nil)
	s.MockIRepositoryCompany.On("List", ctx, params).Return(companies, nil)
	actual, err := s.Company.List(params)
	s.Nil(err)
	s.Equal(companies, actual)
}

func (s *CompanyResolverTestSuite) TestList_Fail() {
	var (
		ctx       = context.Background()
		params    = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{"page": 1, "pageSize": 10}, Args: map[string]interface{}{}}
		companies []models.Company
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryContext", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(s.MockIRows, errors.New("some errors"))
	s.MockIRows.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("some errors"))
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockIHandlerCompany.On("List", ctx, params).Return(companies, errors.New("some errors"))
	s.MockIRepositoryCompany.On("List", ctx, params).Return(companies, errors.New("some errors"))
	actual, err := s.Company.List(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *CompanyResolverTestSuite) TestCount_Success() {
	var (
		ctx    = context.Background()
		params = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		count  int64
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryRowContext", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(s.MockIRows, nil)
	s.MockIRows.On("Scan", mock.Anything).Return(nil)
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockIHandlerCompany.On("Count", ctx, params).Return(count, nil)
	s.MockIRepositoryCompany.On("Count", ctx, params).Return(count, nil)
	actual, err := s.Company.Count(params)
	s.Nil(err)
	s.Equal(count, actual)
}

func (s *CompanyResolverTestSuite) TestCount_Fail() {
	var (
		ctx    = context.Background()
		params = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		count  int64
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryRowContext", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(s.MockIRows, errors.New("some errors"))
	s.MockIRows.On("Scan", mock.Anything).Return(errors.New("some errors"))
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockIHandlerCompany.On("Count", ctx, params).Return(count, errors.New("some errors"))
	s.MockIRepositoryCompany.On("Count", ctx, params).Return(count, errors.New("some errors"))
	actual, err := s.Company.Count(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *CompanyResolverTestSuite) TestInsert_Success() {
	var (
		ctx            = context.Background()
		args     int64 = 1
		sampleID int64 = 1
		params         = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		company  models.Company
	)
	// Mock Insert
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(s.MockIResult, nil)
	s.MockIResult.On("LastInsertId").Return(sampleID, nil)
	s.MockIHandlerCompany.On("Insert", ctx, params).Return(company, nil)
	s.MockIRepositoryCompany.On("Insert", ctx, params).Return(company, nil)
	// Mock GetByID
	s.MockIStmt.On("QueryRowContext", ctx, args).Return(s.MockIRow)
	s.MockIRow.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	s.MockIHandlerCompany.On("GetByID", ctx, params).Return(company, nil)
	s.MockIRepositoryCompany.On("GetByID", ctx, params).Return(company, nil)

	actual, err := s.Company.Insert(params)
	s.Nil(err)
	s.Equal(company, actual)
}

func (s *CompanyResolverTestSuite) TestInsert_Fail() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		company  models.Company
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("ExecContext", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("LastInsertId").Return(sampleID, errors.New("some errors"))
	s.MockIHandlerCompany.On("Insert", ctx, params).Return(company, errors.New("some errors"))
	s.MockIRepositoryCompany.On("Insert", ctx, params).Return(company, errors.New("some errors"))
	actual, err := s.Company.Insert(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *CompanyResolverTestSuite) TestUpdate_Success() {
	var (
		ctx          = context.Background()
		args   int64 = 1
		params       = graphql.ResolveParams{
			Context: context.Background(),
			Source:  map[string]interface{}{},
			Args:    map[string]interface{}{"id": 1, "createdBy": "tanphamhaiduong@gmail.com"},
		}
		rowEffected int64 = 1
		company     models.Company
	)
	// Mock for Update
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(s.MockIResult, nil)
	s.MockIResult.On("RowsAffected").Return(rowEffected, nil)
	s.MockIHandlerCompany.On("Update", ctx, params).Return(company, nil)
	s.MockIRepositoryCompany.On("Update", ctx, params).Return(company, nil)
	// Mock for LoadByID
	s.MockIStmt.On("QueryRowContext", ctx, args).Return(s.MockIRow)
	s.MockIRow.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	s.MockIHandlerCompany.On("GetByID", ctx, params).Return(company, nil)
	s.MockIRepositoryCompany.On("GetByID", ctx, params).Return(company, nil)
	actual, err := s.Company.Update(params)
	s.Nil(err)
	s.Equal(company, actual)
}

func (s *CompanyResolverTestSuite) TestUpdate_Fail() {
	var (
		ctx    = context.Background()
		params = graphql.ResolveParams{
			Context: context.Background(),
			Source:  map[string]interface{}{},
			Args:    map[string]interface{}{"id": 1},
		}
		company models.Company
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("ExecContext", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(s.MockIResult, errors.New("some errors"))
	s.MockIHandlerCompany.On("Update", ctx, params).Return(company, errors.New("some errors"))
	s.MockIRepositoryCompany.On("Update", ctx, params).Return(company, errors.New("some errors"))
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
		rowEffected int64 = 1
		expectedID  int64 = 1
	)

	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx, rowEffected).Return(s.MockIResult, nil)
	s.MockIResult.On("RowsAffected").Return(rowEffected, nil)
	s.MockIHandlerCompany.On("Delete", ctx, params).Return(rowEffected, nil)
	s.MockIRepositoryCompany.On("Delete", ctx, params).Return(rowEffected, nil)
	actual, err := s.Company.Delete(params)
	s.Nil(err)
	s.Equal(expectedID, actual)
}

func (s *CompanyResolverTestSuite) TestDelete_Fail() {
	var (
		ctx    = context.Background()
		params = graphql.ResolveParams{
			Context: context.Background(),
			Source:  map[string]interface{}{},
			Args:    map[string]interface{}{"id": 1},
		}
		rowEffected int64
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("ExecContext", ctx, rowEffected).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("RowsAffected").Return(rowEffected, errors.New("some errors"))
	s.MockIHandlerCompany.On("Delete", ctx, params).Return(rowEffected, errors.New("some errors"))
	s.MockIRepositoryCompany.On("Delete", ctx, params).Return(rowEffected, errors.New("some errors"))
	actual, err := s.Company.Delete(params)
	s.Nil(actual)
	s.NotNil(err)
}
