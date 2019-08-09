// @generated
package company

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/mock"
	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
	"github.com/tanphamhaiduong/go/delta/internal/utils"
)

func (s *CompanyRepositoryTestSuite) TestGetByID_Success() {
	var (
		ctx   = context.Background()
		param = arguments.CompanyGetByIDArgs{
			ID: 1,
		}
		company models.Company
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryRowContext", ctx, param.ID).Return(s.MockIRow)
	s.MockIRow.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(nil)
	s.MockICompany.On("GetByID", ctx, param).Return(company, nil)
	actual, err := s.Repository.GetByID(ctx, param)
	s.Nil(err)
	s.Equal(company, actual)
}

func (s *CompanyRepositoryTestSuite) TestGetByID_Fail() {
	var (
		ctx   = context.Background()
		param = arguments.CompanyGetByIDArgs{
			ID: 1,
		}
		company models.Company
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryRowContext", ctx, param.ID).Return(s.MockIRow)
	s.MockIRow.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(errors.New("some errors"))
	s.MockICompany.On("GetByID", ctx, param).Return(company, errors.New("some errors"))
	actual, err := s.Repository.GetByID(ctx, param)
	s.Equal(company, actual)
	s.NotNil(err)
}

func (s *CompanyRepositoryTestSuite) TestGetByIDs_Success() {
	var (
		ctx   = context.Background()
		param = arguments.CompanyGetByIDsArgs{
			IDs: []int64{1, 2},
		}
		companies []models.Company
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryContext", ctx, mock.Anything, mock.Anything).Return(s.MockIRows, nil)
	s.MockIRows.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(nil)
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockICompany.On("GetByIDs", ctx, param).Return(companies, nil)
	actual, err := s.Repository.GetByIDs(ctx, param)
	s.Nil(err)
	s.Equal(companies, actual)
}

func (s *CompanyRepositoryTestSuite) TestGetByIDs_Fail() {
	var (
		ctx   = context.Background()
		param = arguments.CompanyGetByIDsArgs{
			IDs: []int64{1, 2},
		}
		companies []models.Company
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryContext", ctx, mock.Anything, mock.Anything).Return(s.MockIRows, errors.New("some errors"))
	s.MockIRows.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(errors.New("some errors"))
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockICompany.On("GetByIDs", ctx, param).Return(companies, errors.New("some errors"))
	actual, err := s.Repository.GetByIDs(ctx, param)
	s.Equal(companies, actual)
	s.NotNil(err)
}

func (s *CompanyRepositoryTestSuite) TestSetArgsToListSelectBuilder_Success() {
	var (
		params = arguments.CompanyListArgs{
			ID:           1,
			Name:         "mockString",
			CompanyCode:  "mockString",
			Status:       "active",
			CreatedBy:    "mockString",
			UpdatedBy:    "mockString",
			ApiSecretKey: "mockString",
			Page:         1,
			PageSize:     10,
		}
		selectBuilder = sq.Select("*").From("company")
	)
	offset := utils.CalculateOffsetForPage(params.Page, params.PageSize)
	expectedSelectBuilder := selectBuilder.Where(sq.Eq{"id": params.ID}).Where(sq.Like{"name": params.Name}).Where(sq.Eq{"company_code": params.CompanyCode}).Where(sq.Eq{"status": params.Status}).Where(sq.Eq{"created_by": params.CreatedBy}).Where(sq.Eq{"updated_by": params.UpdatedBy}).Where(sq.Eq{"api_secret_key": params.ApiSecretKey}).Limit(uint64(params.PageSize)).Offset(uint64(offset))
	expectSQL, expectArgs, expectErr := expectedSelectBuilder.ToSql()
	// Actual
	actual := s.Repository.setArgsToListSelectBuilder(selectBuilder, params)
	sql, args, err := actual.ToSql()
	s.Nil(err)
	s.Equal(expectSQL, sql)
	s.Equal(expectArgs, args)
	s.Equal(expectErr, err)
}

func (s *CompanyRepositoryTestSuite) TestList_Success() {
	var (
		ctx    = context.Background()
		params = arguments.CompanyListArgs{
			ID:           1,
			Name:         "mockString",
			CompanyCode:  "mockString",
			Status:       "active",
			CreatedBy:    "mockString",
			UpdatedBy:    "mockString",
			ApiSecretKey: "mockString",
			Page:         1,
			PageSize:     10,
		}
		companies []models.Company
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryContext", ctx,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(s.MockIRows, nil)
	s.MockIRows.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(nil)
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockICompany.On("List", ctx, params).Return(companies, nil)
	actual, err := s.Repository.List(ctx, params)
	s.Nil(err)
	s.Equal(companies, actual)
}

func (s *CompanyRepositoryTestSuite) TestList_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.CompanyListArgs{
			ID:           1,
			Name:         "mockString",
			CompanyCode:  "mockString",
			Status:       "active",
			CreatedBy:    "mockString",
			UpdatedBy:    "mockString",
			ApiSecretKey: "mockString",
			Page:         1,
			PageSize:     10,
		}
		companies []models.Company
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryContext", ctx,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(s.MockIRows, errors.New("some errors"))
	s.MockIRows.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(errors.New("some errors"))
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockICompany.On("List", ctx, params).Return(companies, errors.New("some errors"))
	actual, err := s.Repository.List(ctx, params)
	s.Equal(companies, actual)
	s.NotNil(err)
}

func (s *CompanyRepositoryTestSuite) TestSetArgsToCountSelectBuilder_Success() {
	var (
		params = arguments.CompanyCountArgs{
			ID:           1,
			Name:         "mockString",
			CompanyCode:  "mockString",
			Status:       "active",
			CreatedBy:    "mockString",
			UpdatedBy:    "mockString",
			ApiSecretKey: "mockString",
		}
		selectBuilder = sq.Select("COUNT(id)").From("company")
	)
	expectedSelectBuilder := selectBuilder.Where(sq.Eq{"id": params.ID}).Where(sq.Like{"name": params.Name}).Where(sq.Eq{"company_code": params.CompanyCode}).Where(sq.Eq{"status": params.Status}).Where(sq.Eq{"created_by": params.CreatedBy}).Where(sq.Eq{"updated_by": params.UpdatedBy}).Where(sq.Eq{"api_secret_key": params.ApiSecretKey})
	expectSQL, expectArgs, expectErr := expectedSelectBuilder.ToSql()
	// Actual
	actual := s.Repository.setArgsToCountSelectBuilder(selectBuilder, params)
	sql, args, err := actual.ToSql()
	s.Nil(err)
	s.Equal(expectSQL, sql)
	s.Equal(expectArgs, args)
	s.Equal(expectErr, err)
}

func (s *CompanyRepositoryTestSuite) TestCount_Success() {
	var (
		ctx    = context.Background()
		params = arguments.CompanyCountArgs{
			ID:           1,
			Name:         "mockString",
			CompanyCode:  "mockString",
			Status:       "active",
			CreatedBy:    "mockString",
			UpdatedBy:    "mockString",
			ApiSecretKey: "mockString",
		}
		count int64
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryRowContext", ctx,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(s.MockIRows, nil)
	s.MockIRows.On("Scan", mock.Anything).Return(nil)
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockICompany.On("List", ctx, params).Return(count, nil)
	actual, err := s.Repository.Count(ctx, params)
	s.Nil(err)
	s.Equal(count, actual)
}

func (s *CompanyRepositoryTestSuite) TestCount_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.CompanyCountArgs{
			ID:           1,
			Name:         "mockString",
			CompanyCode:  "mockString",
			Status:       "active",
			CreatedBy:    "mockString",
			UpdatedBy:    "mockString",
			ApiSecretKey: "mockString",
		}
		count int64
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryRowContext", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(s.MockIRows, errors.New("some errors"))
	s.MockIRows.On("Scan", mock.Anything).Return(errors.New("some errors"))
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockICompany.On("List", ctx, params).Return(count, errors.New("some errors"))
	actual, err := s.Repository.Count(ctx, params)
	s.Equal(count, actual)
	s.NotNil(err)
}

func (s *CompanyRepositoryTestSuite) TestInsert_Success() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = arguments.CompanyInsertArgs{
			Name:         "mockString",
			CompanyCode:  "mockString",
			Status:       "active",
			CreatedBy:    "mockString",
			UpdatedBy:    "mockString",
			ApiSecretKey: "mockString",
		}
		expected models.Company
	)
	company := models.Company{
		ID:           sampleID,
		Name:         params.Name,
		CompanyCode:  params.CompanyCode,
		Status:       params.Status,
		CreatedBy:    params.CreatedBy,
		UpdatedBy:    params.UpdatedBy,
		ApiSecretKey: params.ApiSecretKey,
	}
	//Mock Insert
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx,
		params.Name,
		params.CompanyCode,
		params.Status,
		params.CreatedBy,
		params.UpdatedBy,
		params.ApiSecretKey,
	).Return(s.MockIResult, nil)
	s.MockIResult.On("LastInsertId").Return(sampleID, nil)
	s.MockICompany.On("Insert", ctx, params).Return(company, nil)
	// Mock GetByID
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryRowContext", ctx, mock.Anything).Return(s.MockIRow)
	s.MockIRow.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(nil)
	s.MockICompany.On("GetByID", ctx, params).Return(company, nil)

	actual, err := s.Repository.Insert(ctx, params)
	s.Nil(err)
	s.Equal(expected, actual)
}

func (s *CompanyRepositoryTestSuite) TestInsert_Fail() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = arguments.CompanyInsertArgs{
			Name:         "mockString",
			CompanyCode:  "mockString",
			Status:       "active",
			CreatedBy:    "mockString",
			UpdatedBy:    "mockString",
			ApiSecretKey: "mockString",
		}
		company models.Company
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("ExecContext", ctx,
		params.Name,
		params.CompanyCode,
		params.Status,
		params.CreatedBy,
		params.UpdatedBy,
		params.ApiSecretKey,
	).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("LastInsertId").Return(sampleID, errors.New("some errors"))
	s.MockICompany.On("Insert", ctx, params).Return(company, errors.New("some errors"))
	actual, err := s.Repository.Insert(ctx, params)
	s.Equal(company, actual)
	s.NotNil(err)
}

func (s *CompanyRepositoryTestSuite) TestSetArgsToUpdateBuilder_Success() {
	var (
		sampleID   int64 = 1
		mockString       = "mockString"
		status           = "active"
		params           = arguments.CompanyUpdateArgs{
			ID:           &sampleID,
			Name:         &mockString,
			CompanyCode:  &mockString,
			Status:       &status,
			CreatedBy:    &mockString,
			UpdatedBy:    &mockString,
			ApiSecretKey: &mockString,
		}
	)
	updateBuilder := sq.Update("company").Where(sq.Eq{"id": *params.ID})
	expectedSelectBuilder := updateBuilder.Set("name", *params.Name).Set("company_code", *params.CompanyCode).Set("status", *params.Status).Set("created_by", *params.CreatedBy).Set("updated_by", *params.UpdatedBy).Set("api_secret_key", *params.ApiSecretKey)
	actual := s.Repository.setArgsToUpdateBuilder(updateBuilder, params)
	s.Equal(expectedSelectBuilder, actual)
}

func (s *CompanyRepositoryTestSuite) TestUpdate_Success() {
	var (
		ctx              = context.Background()
		sampleID   int64 = 1
		mockString       = "mockString"
		status           = "active"
		params           = arguments.CompanyUpdateArgs{
			ID:           &sampleID,
			Name:         &mockString,
			CompanyCode:  &mockString,
			Status:       &status,
			CreatedBy:    &mockString,
			UpdatedBy:    &mockString,
			ApiSecretKey: &mockString,
		}
		expected models.Company
	)
	company := models.Company{
		ID:           *params.ID,
		Name:         *params.Name,
		CompanyCode:  *params.CompanyCode,
		Status:       *params.Status,
		CreatedBy:    *params.CreatedBy,
		UpdatedBy:    *params.UpdatedBy,
		ApiSecretKey: *params.ApiSecretKey,
	}
	// Mock Update
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx,
		*params.Name,
		*params.CompanyCode,
		*params.Status,
		*params.CreatedBy,
		*params.UpdatedBy,
		*params.ApiSecretKey,
		*params.ID,
	).Return(s.MockIResult, nil)
	s.MockIResult.On("RowsAffected").Return(*params.ID, nil)
	s.MockICompany.On("Update", ctx, params).Return(company, nil)
	// Mock GetByID
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryRowContext", ctx, mock.Anything).Return(s.MockIRow)
	s.MockIRow.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(nil)
	s.MockICompany.On("GetByID", ctx, params).Return(company, nil)

	actual, err := s.Repository.Update(ctx, params)
	s.Nil(err)
	s.Equal(expected, actual)
}

func (s *CompanyRepositoryTestSuite) TestUpdate_Fail() {
	var (
		ctx              = context.Background()
		sampleID   int64 = 1
		mockString       = "mockString"
		status           = "active"
		params           = arguments.CompanyUpdateArgs{
			ID:           &sampleID,
			Name:         &mockString,
			CompanyCode:  &mockString,
			Status:       &status,
			CreatedBy:    &mockString,
			UpdatedBy:    &mockString,
			ApiSecretKey: &mockString,
		}
		company models.Company
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("ExecContext", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("RowsAffected").Return(sampleID, errors.New("some errors"))
	s.MockICompany.On("Update", ctx, params).Return(company, errors.New("some errors"))
	actual, err := s.Repository.Update(ctx, params)
	s.Equal(company, actual)
	s.NotNil(err)
}

func (s *CompanyRepositoryTestSuite) TestDelete_Success() {
	var (
		ctx   = context.Background()
		param = arguments.CompanyDeleteArgs{
			ID: 1,
		}
		rowEffected int64 = 1
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx, param.ID).Return(s.MockIResult, nil)
	s.MockIResult.On("RowsAffected").Return(param.ID, nil)
	s.MockICompany.On("Delete", ctx, param).Return(rowEffected, nil)
	actual, err := s.Repository.Delete(ctx, param)
	s.Nil(err)
	s.Equal(param.ID, actual)
}

func (s *CompanyRepositoryTestSuite) TestDelete_Fail() {
	var (
		ctx   = context.Background()
		param = arguments.CompanyDeleteArgs{
			ID: 1,
		}
		rowEffected int64
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("ExecContext", ctx, param.ID).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("RowsAffected").Return(param.ID, errors.New("some errors"))
	s.MockICompany.On("Delete", ctx, param).Return(rowEffected, errors.New("some errors"))
	actual, err := s.Repository.Delete(ctx, param)
	s.Equal(rowEffected, actual)
	s.NotNil(err)
}
