// @generated
package company

import (
	"context"
	"errors"

	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
)

func (s *CompanyHandlerTestSuite) TestGetByID_Success() {
	var (
		ctx   = context.Background()
		param = arguments.CompanyGetByID{
			ID: 1,
		}
		company = models.Company{}
	)
	s.MockICompany.On("GetByID", ctx, param).Return(company, nil)
	actual, err := s.Company.GetByID(ctx, param)
	s.Nil(err)
	s.Equal(company, actual)
}

func (s *CompanyHandlerTestSuite) TestGetByID_Fail() {
	var (
		ctx   = context.Background()
		param = arguments.CompanyGetByID{
			ID: 1,
		}
		company = models.Company{}
	)
	s.MockICompany.On("GetByID", ctx, param).Return(company, errors.New("some errors"))
	actual, err := s.Company.GetByID(ctx, param)
	s.Equal(company, actual)
	s.NotNil(err)
}

func (s *CompanyHandlerTestSuite) TestGetByID_Fail1() {
	var (
		ctx   = context.Background()
		param = arguments.CompanyGetByID{
			ID: 0,
		}
		company = models.Company{}
	)
	s.MockICompany.On("GetByID", ctx, param).Return(company, errors.New("some errors"))
	actual, err := s.Company.GetByID(ctx, param)
	s.Equal(company, actual)
	s.NotNil(err)
}

func (s *CompanyHandlerTestSuite) TestGetByIDs_Success() {
	var (
		ctx   = context.Background()
		param = arguments.CompanyGetByIDs{
			IDs: []int64{1, 2},
		}
		companies []models.Company
	)
	s.MockICompany.On("GetByIDs", ctx, param).Return(companies, nil)
	actual, err := s.Company.GetByIDs(ctx, param)
	s.Nil(err)
	s.Equal(companies, actual)
}

func (s *CompanyHandlerTestSuite) TestGetByIDs_Fail() {
	var (
		ctx   = context.Background()
		param = arguments.CompanyGetByIDs{
			IDs: []int64{1, 2},
		}
		companies []models.Company
	)
	s.MockICompany.On("GetByIDs", ctx, param).Return(companies, errors.New("some errors"))
	actual, err := s.Company.GetByIDs(ctx, param)
	s.Equal(companies, actual)
	s.NotNil(err)
}

func (s *CompanyHandlerTestSuite) TestGetByIDs_Fail1() {
	var (
		ctx       = context.Background()
		param     = arguments.CompanyGetByIDs{}
		companies []models.Company
	)
	s.MockICompany.On("GetByIDs", ctx, param).Return(companies, errors.New("some errors"))
	actual, err := s.Company.GetByIDs(ctx, param)
	s.Equal(companies, actual)
	s.NotNil(err)
}

func (s *CompanyHandlerTestSuite) TestList_Success() {
	var (
		ctx    = context.Background()
		params = arguments.CompanyList{
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
	s.MockICompany.On("List", ctx, params).Return(companies, nil)
	actual, err := s.Company.List(ctx, params)
	s.Nil(err)
	s.Equal(companies, actual)
}

func (s *CompanyHandlerTestSuite) TestList_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.CompanyList{
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
	s.MockICompany.On("List", ctx, params).Return(companies, errors.New("some errors"))
	actual, err := s.Company.List(ctx, params)
	s.Equal(companies, actual)
	s.NotNil(err)
}

func (s *CompanyHandlerTestSuite) TestList_Fail1() {
	var (
		ctx    = context.Background()
		params = arguments.CompanyList{
			ID:           0,
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
	s.MockICompany.On("List", ctx, params).Return(companies, errors.New("some errors"))
	actual, err := s.Company.List(ctx, params)
	s.Equal(companies, actual)
	s.NotNil(err)
}

func (s *CompanyHandlerTestSuite) TestCount_Success() {
	var (
		ctx    = context.Background()
		params = arguments.CompanyCount{
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
	s.MockICompany.On("Count", ctx, params).Return(count, nil)
	actual, err := s.Company.Count(ctx, params)
	s.Nil(err)
	s.Equal(count, actual)
}

func (s *CompanyHandlerTestSuite) TestCount_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.CompanyCount{
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
	s.MockICompany.On("Count", ctx, params).Return(count, errors.New("some errors"))
	actual, err := s.Company.Count(ctx, params)
	s.Equal(count, actual)
	s.NotNil(err)
}

func (s *CompanyHandlerTestSuite) TestCount_Fail1() {
	var (
		ctx    = context.Background()
		params = arguments.CompanyCount{
			ID:           0,
			Name:         "mockString",
			CompanyCode:  "mockString",
			Status:       "active",
			CreatedBy:    "mockString",
			UpdatedBy:    "mockString",
			ApiSecretKey: "mockString",
		}
		count int64
	)
	s.MockICompany.On("Count", ctx, params).Return(count, errors.New("some errors"))
	actual, err := s.Company.Count(ctx, params)
	s.Equal(count, actual)
	s.NotNil(err)
}

func (s *CompanyHandlerTestSuite) TestInsert_Success() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = arguments.CompanyInsert{
			Name:         "mockString",
			CompanyCode:  "mockString",
			Status:       "active",
			CreatedBy:    "mockString",
			UpdatedBy:    "mockString",
			ApiSecretKey: "mockString",
		}
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
	// Mock Insert
	s.MockICompany.On("Insert", ctx, params).Return(company, nil)
	// Mock GetByID
	s.MockICompany.On("GetByID", ctx, params).Return(company, nil)

	actual, err := s.Company.Insert(ctx, params)
	s.Nil(err)
	s.Equal(company, actual)
}

func (s *CompanyHandlerTestSuite) TestInsert_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.CompanyInsert{
			Name:         "mockString",
			CompanyCode:  "mockString",
			Status:       "active",
			CreatedBy:    "mockString",
			UpdatedBy:    "mockString",
			ApiSecretKey: "mockString",
		}
		company = models.Company{}
	)
	s.MockICompany.On("Insert", ctx, params).Return(company, errors.New("some errors"))
	actual, err := s.Company.Insert(ctx, params)
	s.Equal(company, actual)
	s.NotNil(err)
}

func (s *CompanyHandlerTestSuite) TestInsert_Fail1() {
	var (
		ctx    = context.Background()
		params = arguments.CompanyInsert{
			Name:         "mockString",
			CompanyCode:  "mockString",
			Status:       "active",
			CreatedBy:    "mockString",
			UpdatedBy:    "mockString",
			ApiSecretKey: "mockString",
		}
		company = models.Company{}
	)
	s.MockICompany.On("Insert", ctx, params).Return(company, errors.New("some errors"))
	actual, err := s.Company.Insert(ctx, params)
	s.Equal(company, actual)
	s.NotNil(err)
}

func (s *CompanyHandlerTestSuite) TestUpdate_Success() {
	var (
		ctx              = context.Background()
		sampleID   int64 = 1
		mockString       = "mockString"
		status           = "active"
		params           = arguments.CompanyUpdate{
			ID:           &sampleID,
			Name:         &mockString,
			CompanyCode:  &mockString,
			Status:       &status,
			CreatedBy:    &mockString,
			UpdatedBy:    &mockString,
			ApiSecretKey: &mockString,
		}
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
	// Mock Insert
	s.MockICompany.On("Update", ctx, params).Return(company, nil)
	// Mock GetByID
	s.MockICompany.On("GetByID", ctx, params).Return(company, nil)

	actual, err := s.Company.Update(ctx, params)
	s.Nil(err)
	s.Equal(company, actual)
}

func (s *CompanyHandlerTestSuite) TestUpdate_Fail() {
	var (
		ctx              = context.Background()
		sampleID   int64 = 1
		mockString       = "mockString"
		status           = "active"
		params           = arguments.CompanyUpdate{
			ID:           &sampleID,
			Name:         &mockString,
			CompanyCode:  &mockString,
			Status:       &status,
			CreatedBy:    &mockString,
			UpdatedBy:    &mockString,
			ApiSecretKey: &mockString,
		}
		company = models.Company{}
	)
	s.MockICompany.On("Update", ctx, params).Return(company, errors.New("some errors"))
	actual, err := s.Company.Update(ctx, params)
	s.Equal(company, actual)
	s.NotNil(err)
}

func (s *CompanyHandlerTestSuite) TestUpdate_Fail1() {
	var (
		ctx        = context.Background()
		sampleID   int64
		mockString = "mockString"
		status     = "active"
		params     = arguments.CompanyUpdate{
			ID:           &sampleID,
			Name:         &mockString,
			CompanyCode:  &mockString,
			Status:       &status,
			CreatedBy:    &mockString,
			UpdatedBy:    &mockString,
			ApiSecretKey: &mockString,
		}
		company = models.Company{}
	)
	s.MockICompany.On("Update", ctx, params).Return(company, errors.New("some errors"))
	actual, err := s.Company.Update(ctx, params)
	s.Equal(company, actual)
	s.NotNil(err)
}

func (s *CompanyHandlerTestSuite) TestDelete_Success() {
	var (
		ctx   = context.Background()
		param = arguments.CompanyDelete{
			ID: 1,
		}
	)
	s.MockICompany.On("Delete", ctx, param).Return(param.ID, nil)
	actual, err := s.Company.Delete(ctx, param)
	s.Nil(err)
	s.Equal(param.ID, actual)
}

func (s *CompanyHandlerTestSuite) TestDelete_Fail() {
	var (
		ctx   = context.Background()
		param = arguments.CompanyDelete{
			ID: 1,
		}
		rowEffected int64
	)
	s.MockICompany.On("Delete", ctx, param).Return(rowEffected, errors.New("some errors"))
	actual, err := s.Company.Delete(ctx, param)
	s.Equal(rowEffected, actual)
	s.NotNil(err)
}

func (s *CompanyHandlerTestSuite) TestDelete_Fail1() {
	var (
		ctx   = context.Background()
		param = arguments.CompanyDelete{
			ID: 0,
		}
		rowEffected int64
	)
	s.MockICompany.On("Delete", ctx, param).Return(rowEffected, errors.New("some errors"))
	actual, err := s.Company.Delete(ctx, param)
	s.Equal(rowEffected, actual)
	s.NotNil(err)
}
