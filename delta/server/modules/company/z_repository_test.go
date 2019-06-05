package company

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/mock"
	"github.com/tanphamhaiduong/go/delta/server/arguments"
	"github.com/tanphamhaiduong/go/delta/server/models"
	"github.com/tanphamhaiduong/go/delta/server/utils"
)

func (s *CompanyRepositoryTestSuite) TestGetByID_Success() {
	var (
		ctx    = context.Background()
		params = arguments.CompanyGetByIDArgs{
			ID: 1,
		}
		company = models.Company{}
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryRowContext", ctx, mock.Anything).Return(s.MockIRow)
	s.MockIRow.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
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
		company = models.Company{}
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
		companies = []models.Company{}
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
		companies = []models.Company{}
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
		params = arguments.CompanyListArgs{
			ID:        1,
			Name:      "mockString",
			Status:    "mockString",
			CreatedBy: "mockString",
			UpdatedBy: "mockString",
			Page:      1,
			PageSize:  10,
		}
		offset                = utils.CalculateOffsetForPage(params.Page, params.PageSize)
		selectBuilder         = sq.Select("*").From("company")
		expectedSelectBuilder = selectBuilder.Where(sq.Eq{"id": params.ID}).Where(sq.Like{"name": params.Name}).Where(sq.Eq{"status": params.Status}).Where(sq.Eq{"created_by": params.CreatedBy}).Where(sq.Eq{"updated_by": params.UpdatedBy}).Limit(uint64(params.PageSize)).Offset(uint64(offset))
	)
	actual := s.Repository.setArgsToListSelectBuilder(selectBuilder, params)
	s.Equal(expectedSelectBuilder, actual)
}

func (s *CompanyRepositoryTestSuite) TestList_Success() {
	var (
		ctx    = context.Background()
		params = arguments.CompanyListArgs{
			ID:        1,
			Name:      "mockString",
			Status:    "mockString",
			CreatedBy: "mockString",
			UpdatedBy: "mockString",
			Page:      1,
			PageSize:  10,
		}
		companies = []models.Company{}
	)
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
			ID:        1,
			Name:      "mockString",
			Status:    "mockString",
			CreatedBy: "mockString",
			UpdatedBy: "mockString",
		}
		companies = []models.Company{}
	)
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
		params = arguments.CompanyCountArgs{
			ID:        1,
			Name:      "mockString",
			Status:    "mockString",
			CreatedBy: "mockString",
			UpdatedBy: "mockString",
		}
		selectBuilder         = sq.Select("COUNT(id)").From("company")
		expectedSelectBuilder = selectBuilder.Where(sq.Eq{"id": params.ID}).Where(sq.Like{"name": params.Name}).Where(sq.Eq{"status": params.Status}).Where(sq.Eq{"created_by": params.CreatedBy}).Where(sq.Eq{"updated_by": params.UpdatedBy})
	)
	actual := s.Repository.setArgsToCountSelectBuilder(selectBuilder, params)
	s.Equal(expectedSelectBuilder, actual)
}

func (s *CompanyRepositoryTestSuite) TestCount_Success() {
	var (
		ctx    = context.Background()
		params = arguments.CompanyCountArgs{
			ID:        1,
			Name:      "mockString",
			Status:    "mockString",
			CreatedBy: "mockString",
			UpdatedBy: "mockString",
		}
		count int64
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryRowContext", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(s.MockIRows, nil)
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
			ID:        1,
			Name:      "mockString",
			Status:    "mockString",
			CreatedBy: "mockString",
			UpdatedBy: "mockString",
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
			Name:      "mockString",
			Status:    "mockString",
			CreatedBy: "mockString",
			UpdatedBy: "mockString",
		}
		company = models.Company{
			ID:        sampleID,
			Name:      "mockString",
			Status:    "mockString",
			CreatedBy: "mockString",
			UpdatedBy: "mockString",
		}
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx, params.Name, params.Status, params.CreatedBy, params.UpdatedBy).Return(s.MockIResult, nil)
	s.MockIResult.On("LastInsertId").Return(sampleID, nil)
	s.MockICompany.On("Insert", ctx, params).Return(company, nil)
	actual, err := s.Repository.Insert(ctx, params)
	s.Nil(err)
	s.Equal(company, actual)
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
		company = models.Company{}
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("ExecContext", ctx, params.Name, params.Status, params.CreatedBy, params.UpdatedBy).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("LastInsertId").Return(sampleID, errors.New("some errors"))
	s.MockICompany.On("Insert", ctx, params).Return(company, errors.New("some errors"))
	actual, err := s.Repository.Insert(ctx, params)
	s.Equal(company, actual)
	s.NotNil(err)
}

func (s *CompanyRepositoryTestSuite) TestUpdate_Success() {
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
		company = models.Company{
			ID:        sampleID,
			Name:      mockString,
			Status:    mockString,
			CreatedBy: mockString,
			UpdatedBy: mockString,
		}
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(s.MockIResult, nil)
	s.MockIResult.On("RowsAffected").Return(sampleID, nil)
	s.MockICompany.On("Update", ctx, params).Return(company, nil)
	actual, err := s.Repository.Update(ctx, params)
	s.Nil(err)
	s.Equal(company, actual)
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
		company = models.Company{}
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("ExecContext", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("RowsAffected").Return(sampleID, errors.New("some errors"))
	s.MockICompany.On("Update", ctx, params).Return(company, errors.New("some errors"))
	actual, err := s.Repository.Update(ctx, params)
	s.Equal(company, actual)
	s.NotNil(err)
}
