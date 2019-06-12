// @generated
package company

import (
	"context"
	"errors"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/mock"
	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
	"github.com/tanphamhaiduong/go/delta/internal/utils"
)

func (s *CompanyRepositoryTestSuite) TestGetByID_Success() {
	var (
		ctx    = context.Background()
		params = arguments.CompanyGetByIDArgs{
			ID: 1,
		}
		company models.Company
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryRowContext", ctx, mock.Anything).Return(s.MockIRow)
	s.MockIRow.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(nil)
	s.MockICompany.On("GetByID", ctx, params).Return(company, nil)
	actual, err := s.Repository.GetByID(ctx, params)
	s.Nil(err)
	s.Equal(company, actual)
}

func (s *CompanyRepositoryTestSuite) TestGetByID_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.CompanyGetByIDArgs{
			ID: 1,
		}
		company models.Company
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryRowContext", ctx, mock.Anything).Return(s.MockIRow)
	s.MockIRow.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("some errors"))
	s.MockICompany.On("GetByID", ctx, params).Return(company, errors.New("some errors"))
	actual, err := s.Repository.GetByID(ctx, params)
	s.Equal(company, actual)
	s.NotNil(err)
}

func (s *CompanyRepositoryTestSuite) TestGetByIDs_Success() {
	var (
		ctx    = context.Background()
		params = arguments.CompanyGetByIDsArgs{
			IDs: []int64{1, 2},
		}
		companies []models.Company
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryContext", ctx, mock.Anything, mock.Anything).Return(s.MockIRows, nil)
	s.MockIRows.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockICompany.On("GetByIDs", ctx, params).Return(companies, nil)
	actual, err := s.Repository.GetByIDs(ctx, params)
	s.Nil(err)
	s.Equal(companies, actual)
}

func (s *CompanyRepositoryTestSuite) TestGetByIDs_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.CompanyGetByIDsArgs{
			IDs: []int64{1, 2},
		}
		companies []models.Company
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryContext", ctx, mock.Anything, mock.Anything).Return(s.MockIRows, errors.New("some errors"))
	s.MockIRows.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("some errors"))
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockICompany.On("GetByIDs", ctx, params).Return(companies, errors.New("some errors"))
	actual, err := s.Repository.GetByIDs(ctx, params)
	s.Equal(companies, actual)
	s.NotNil(err)
}

func (s *CompanyRepositoryTestSuite) TestSetArgsToListSelectBuilder_Success() {
	var (
		params        = arguments.CompanyListArgs{}
		selectBuilder = sq.Select("*").From("company")
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	offset := utils.CalculateOffsetForPage(params.Page, params.PageSize)
	expectedSelectBuilder := selectBuilder.Where(sq.Eq{"id": params.ID}).Where(sq.Like{"name": params.Name}).Where(sq.Eq{"status": params.Status}).Where(sq.Eq{"created_by": params.CreatedBy}).Where(sq.Eq{"updated_by": params.UpdatedBy}).Limit(uint64(params.PageSize)).Offset(uint64(offset))
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
			Page:     1,
			PageSize: 10,
		}
		companies []models.Company
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryContext", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(s.MockIRows, nil)
	s.MockIRows.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
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
			Page:     1,
			PageSize: 10,
		}
		companies []models.Company
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryContext", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(s.MockIRows, errors.New("some errors"))
	s.MockIRows.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("some errors"))
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockICompany.On("List", ctx, params).Return(companies, errors.New("some errors"))
	actual, err := s.Repository.List(ctx, params)
	s.Equal(companies, actual)
	s.NotNil(err)
}

func (s *CompanyRepositoryTestSuite) TestSetArgsToCountSelectBuilder_Success() {
	var (
		params        = arguments.CompanyCountArgs{}
		selectBuilder = sq.Select("COUNT(id)").From("company")
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	expectedSelectBuilder := selectBuilder.Where(sq.Eq{"id": params.ID}).Where(sq.Like{"name": params.Name}).Where(sq.Eq{"status": params.Status}).Where(sq.Eq{"created_by": params.CreatedBy}).Where(sq.Eq{"updated_by": params.UpdatedBy})
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
		params = arguments.CompanyCountArgs{}
		count  int64
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryRowContext", ctx,
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
		params = arguments.CompanyCountArgs{}
		count  int64
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
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
		params         = arguments.CompanyInsertArgs{}
		expected models.Company
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	company := models.Company{
		ID:        sampleID,
		Name:      params.Name,
		Status:    params.Status,
		CreatedBy: params.CreatedBy,
		UpdatedBy: params.UpdatedBy,
	}
	//Mock Insert
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx,
		params.Name,
		params.Status,
		params.CreatedBy,
		params.UpdatedBy,
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
			Name:      "mockString",
			Status:    "mockString",
			CreatedBy: "mockString",
			UpdatedBy: "mockString",
		}
		company models.Company
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("ExecContext", ctx,
		params.Name,
		params.Status,
		params.CreatedBy,
		params.UpdatedBy,
	).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("LastInsertId").Return(sampleID, errors.New("some errors"))
	s.MockICompany.On("Insert", ctx, params).Return(company, errors.New("some errors"))
	actual, err := s.Repository.Insert(ctx, params)
	s.Equal(company, actual)
	s.NotNil(err)
}

func (s *CompanyRepositoryTestSuite) TestSetArgsToUpdateBuilder_Success() {
	var (
		params = arguments.CompanyUpdateArgs{}
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	updateBuilder := sq.Update("company").Where(sq.Eq{"id": *params.ID})
	expectedSelectBuilder := updateBuilder.Set("name", *params.Name).Set("status", *params.Status).Set("created_by", *params.CreatedBy).Set("updated_by", *params.UpdatedBy)
	actual := s.Repository.setArgsToUpdateBuilder(updateBuilder, params)
	s.Equal(expectedSelectBuilder, actual)
}

func (s *CompanyRepositoryTestSuite) TestUpdate_Success() {
	var (
		ctx      = context.Background()
		params   = arguments.CompanyUpdateArgs{}
		expected models.Company
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	company := models.Company{
		ID:        *params.ID,
		Name:      *params.Name,
		Status:    *params.Status,
		CreatedBy: *params.CreatedBy,
		UpdatedBy: *params.UpdatedBy,
	}
	// Mock Update
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx,
		*params.Name,
		*params.Status,
		*params.CreatedBy,
		*params.UpdatedBy,
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
		params           = arguments.CompanyUpdateArgs{
			ID:        &sampleID,
			Name:      &mockString,
			Status:    &mockString,
			CreatedBy: &mockString,
			UpdatedBy: &mockString,
		}
		company models.Company
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
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
		ctx               = context.Background()
		params            = arguments.CompanyDeleteArgs{}
		rowEffected int64 = 1
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx, params.ID).Return(s.MockIResult, nil)
	s.MockIResult.On("RowsAffected").Return(params.ID, nil)
	s.MockICompany.On("Delete", ctx, params).Return(rowEffected, nil)
	actual, err := s.Repository.Delete(ctx, params)
	s.Nil(err)
	s.Equal(params.ID, actual)
}

func (s *CompanyRepositoryTestSuite) TestDelete_Fail() {
	var (
		ctx         = context.Background()
		params      = arguments.CompanyDeleteArgs{}
		rowEffected int64
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("ExecContext", ctx, params.ID).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("RowsAffected").Return(params.ID, errors.New("some errors"))
	s.MockICompany.On("Delete", ctx, params).Return(rowEffected, errors.New("some errors"))
	actual, err := s.Repository.Delete(ctx, params)
	s.Equal(rowEffected, actual)
	s.NotNil(err)
}
