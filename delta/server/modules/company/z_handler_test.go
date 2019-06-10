// @generated
package company

import (
	"context"
	"errors"
	"log"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/mock"

	"github.com/tanphamhaiduong/go/delta/server/arguments"
	"github.com/tanphamhaiduong/go/delta/server/models"
)

func (s *CompanyHandlerTestSuite) TestGetByID_Success() {
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
	actual, err := s.Company.GetByID(ctx, params)
	s.Nil(err)
	s.Equal(company, actual)
}

func (s *CompanyHandlerTestSuite) TestGetByID_Fail() {
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
	actual, err := s.Company.GetByID(ctx, params)
	s.Equal(company, actual)
	s.NotNil(err)
}

func (s *CompanyHandlerTestSuite) TestGetByIDs_Success() {
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
	actual, err := s.Company.GetByIDs(ctx, params)
	s.Nil(err)
	s.Equal(companies, actual)
}

func (s *CompanyHandlerTestSuite) TestGetByIDs_Fail() {
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
	actual, err := s.Company.GetByIDs(ctx, params)
	s.Equal(companies, actual)
	s.NotNil(err)
}

func (s *CompanyHandlerTestSuite) TestList_Success() {
	var (
		ctx       = context.Background()
		params    = arguments.CompanyListArgs{}
		companies []models.Company
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	params.Status = "active"
	params.Page = 1
	params.PageSize = 10
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryContext", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(s.MockIRows, nil)
	s.MockIRows.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockICompany.On("List", ctx, params).Return(companies, nil)
	actual, err := s.Company.List(ctx, params)
	s.Nil(err)
	s.Equal(companies, actual)
}

func (s *CompanyHandlerTestSuite) TestList_Fail() {
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
	actual, err := s.Company.List(ctx, params)
	s.Equal(companies, actual)
	s.NotNil(err)
}

func (s *CompanyHandlerTestSuite) TestCount_Success() {
	var (
		ctx    = context.Background()
		params = arguments.CompanyCountArgs{}
		count  int64
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	params.Status = "active"
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryRowContext", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(s.MockIRows, nil)
	s.MockIRows.On("Scan", mock.Anything).Return(nil)
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockICompany.On("List", ctx, params).Return(count, nil)
	actual, err := s.Company.Count(ctx, params)
	s.Nil(err)
	s.Equal(count, actual)
}

func (s *CompanyHandlerTestSuite) TestCount_Fail() {
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
	actual, err := s.Company.Count(ctx, params)
	s.Equal(count, actual)
	s.NotNil(err)
}

func (s *CompanyHandlerTestSuite) TestInsert_Success() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		args     int64 = 1
		params         = arguments.CompanyInsertArgs{}
		expected models.Company
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	params.Status = "active"
	company := models.Company{
		ID:        sampleID,
		Name:      params.Name,
		Status:    params.Status,
		CreatedBy: params.CreatedBy,
		UpdatedBy: params.UpdatedBy,
	}
	// Mock Insert
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx, params.Name, params.Status, params.CreatedBy, params.UpdatedBy).Return(s.MockIResult, nil)
	s.MockIResult.On("LastInsertId").Return(sampleID, nil)
	s.MockICompany.On("Insert", ctx, params).Return(company, nil)
	// Mock GetByID
	s.MockIStmt.On("QueryRowContext", ctx, args).Return(s.MockIRow)
	s.MockIRow.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	s.MockICompany.On("GetByID", ctx, params).Return(company, nil)

	actual, err := s.Company.Insert(ctx, params)
	s.Nil(err)
	s.Equal(expected, actual)
}

func (s *CompanyHandlerTestSuite) TestInsert_Fail() {
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
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("ExecContext", ctx, params.Name, params.Status, params.CreatedBy, params.UpdatedBy).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("LastInsertId").Return(sampleID, errors.New("some errors"))
	s.MockICompany.On("Insert", ctx, params).Return(company, errors.New("some errors"))
	actual, err := s.Company.Insert(ctx, params)
	s.Equal(company, actual)
	s.NotNil(err)
}

func (s *CompanyHandlerTestSuite) TestUpdate_Success() {
	var (
		ctx      = context.Background()
		params   = arguments.CompanyUpdateArgs{}
		status   = "active"
		expected models.Company
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	params.Status = &status
	company := models.Company{
		ID:        *params.ID,
		Name:      *params.Name,
		Status:    *params.Status,
		CreatedBy: *params.CreatedBy,
		UpdatedBy: *params.UpdatedBy,
	}
	// Mock Insert
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(s.MockIResult, nil)
	s.MockIResult.On("RowsAffected").Return(*params.ID, nil)
	s.MockICompany.On("Update", ctx, params).Return(company, nil)
	// Mock GetByID
	s.MockIStmt.On("QueryRowContext", ctx, *params.ID).Return(s.MockIRow)
	s.MockIRow.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	s.MockICompany.On("GetByID", ctx, params).Return(company, nil)

	actual, err := s.Company.Update(ctx, params)
	s.Nil(err)
	s.Equal(expected, actual)
}

func (s *CompanyHandlerTestSuite) TestUpdate_Fail() {
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
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("ExecContext", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("RowsAffected").Return(sampleID, errors.New("some errors"))
	s.MockICompany.On("Update", ctx, params).Return(company, errors.New("some errors"))
	actual, err := s.Company.Update(ctx, params)
	s.Equal(company, actual)
	s.NotNil(err)
}

func (s *CompanyHandlerTestSuite) TestDelete_Success() {
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
	actual, err := s.Company.Delete(ctx, params)
	s.Nil(err)
	s.Equal(params.ID, actual)
}

func (s *CompanyHandlerTestSuite) TestDelete_Fail() {
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
	actual, err := s.Company.Delete(ctx, params)
	s.Equal(rowEffected, actual)
	s.NotNil(err)
}
