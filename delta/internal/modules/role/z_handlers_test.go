// @generated
package role

import (
	"context"
	"errors"
	"log"

	"github.com/bxcodec/faker"

	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
)

func (s *RoleHandlerTestSuite) TestGetByID_Success() {
	var (
		ctx    = context.Background()
		params = arguments.RoleGetByIDArgs{
			ID: 1,
		}
		role = models.Role{}
	)
	s.MockIRole.On("GetByID", ctx, params).Return(role, nil)
	actual, err := s.Role.GetByID(ctx, params)
	s.Nil(err)
	s.Equal(role, actual)
}

func (s *RoleHandlerTestSuite) TestGetByID_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.RoleGetByIDArgs{
			ID: 1,
		}
		role = models.Role{}
	)
	s.MockIRole.On("GetByID", ctx, params).Return(role, errors.New("some errors"))
	actual, err := s.Role.GetByID(ctx, params)
	s.Equal(role, actual)
	s.NotNil(err)
}

func (s *RoleHandlerTestSuite) TestGetByIDs_Success() {
	var (
		ctx    = context.Background()
		params = arguments.RoleGetByIDsArgs{
			IDs: []int64{1, 2},
		}
		roles []models.Role
	)
	s.MockIRole.On("GetByIDs", ctx, params).Return(roles, nil)
	actual, err := s.Role.GetByIDs(ctx, params)
	s.Nil(err)
	s.Equal(roles, actual)
}

func (s *RoleHandlerTestSuite) TestGetByIDs_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.RoleGetByIDsArgs{
			IDs: []int64{1, 2},
		}
		roles []models.Role
	)
	s.MockIRole.On("GetByIDs", ctx, params).Return(roles, errors.New("some errors"))
	actual, err := s.Role.GetByIDs(ctx, params)
	s.Equal(roles, actual)
	s.NotNil(err)
}

func (s *RoleHandlerTestSuite) TestList_Success() {
	var (
		ctx    = context.Background()
		params = arguments.RoleListArgs{}
		roles  []models.Role
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	params.Status = "active"
	params.Page = 1
	params.PageSize = 10
	s.MockIRole.On("List", ctx, params).Return(roles, nil)
	actual, err := s.Role.List(ctx, params)
	s.Nil(err)
	s.Equal(roles, actual)
}

func (s *RoleHandlerTestSuite) TestList_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.RoleListArgs{
			Page:     1,
			PageSize: 10,
		}
		roles []models.Role
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIRole.On("List", ctx, params).Return(roles, errors.New("some errors"))
	actual, err := s.Role.List(ctx, params)
	s.Equal(roles, actual)
	s.NotNil(err)
}

func (s *RoleHandlerTestSuite) TestCount_Success() {
	var (
		ctx    = context.Background()
		params = arguments.RoleCountArgs{}
		count  int64
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	params.Status = "active"
	s.MockIRole.On("Count", ctx, params).Return(count, nil)
	actual, err := s.Role.Count(ctx, params)
	s.Nil(err)
	s.Equal(count, actual)
}

func (s *RoleHandlerTestSuite) TestCount_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.RoleCountArgs{}
		count  int64
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIRole.On("Count", ctx, params).Return(count, errors.New("some errors"))
	actual, err := s.Role.Count(ctx, params)
	s.Equal(count, actual)
	s.NotNil(err)
}

func (s *RoleHandlerTestSuite) TestInsert_Success() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = arguments.RoleInsertArgs{}
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	params.Status = "active"
	role := models.Role{
		ID:        sampleID,
		Name:      params.Name,
		CompanyID: params.CompanyID,
		Status:    params.Status,
		CreatedBy: params.CreatedBy,
		UpdatedBy: params.UpdatedBy,
	}
	// Mock Insert
	s.MockIRole.On("Insert", ctx, params).Return(role, nil)
	// Mock GetByID
	s.MockIRole.On("GetByID", ctx, params).Return(role, nil)

	actual, err := s.Role.Insert(ctx, params)
	s.Nil(err)
	s.Equal(role, actual)
}

func (s *RoleHandlerTestSuite) TestInsert_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.RoleInsertArgs{
			Name:      "mockString",
			CompanyID: 0,
			Status:    "mockString",
			CreatedBy: "mockString",
			UpdatedBy: "mockString",
		}
		role = models.Role{}
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIRole.On("Insert", ctx, params).Return(role, errors.New("some errors"))
	actual, err := s.Role.Insert(ctx, params)
	s.Equal(role, actual)
	s.NotNil(err)
}

func (s *RoleHandlerTestSuite) TestUpdate_Success() {
	var (
		ctx    = context.Background()
		params = arguments.RoleUpdateArgs{}
		status = "active"
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	params.Status = &status
	role := models.Role{
		ID:        *params.ID,
		Name:      *params.Name,
		CompanyID: *params.CompanyID,
		Status:    *params.Status,
		CreatedBy: *params.CreatedBy,
		UpdatedBy: *params.UpdatedBy,
	}
	// Mock Insert
	s.MockIRole.On("Update", ctx, params).Return(role, nil)
	// Mock GetByID
	s.MockIRole.On("GetByID", ctx, params).Return(role, nil)

	actual, err := s.Role.Update(ctx, params)
	s.Nil(err)
	s.Equal(role, actual)
}

func (s *RoleHandlerTestSuite) TestUpdate_Fail() {
	var (
		ctx              = context.Background()
		sampleID   int64 = 1
		mockString       = "mockString"
		params           = arguments.RoleUpdateArgs{
			ID:        &sampleID,
			Name:      &mockString,
			CompanyID: &sampleID,
			Status:    &mockString,
			CreatedBy: &mockString,
			UpdatedBy: &mockString,
		}
		role = models.Role{}
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIRole.On("Update", ctx, params).Return(role, errors.New("some errors"))
	actual, err := s.Role.Update(ctx, params)
	s.Equal(role, actual)
	s.NotNil(err)
}

func (s *RoleHandlerTestSuite) TestDelete_Success() {
	var (
		ctx    = context.Background()
		params = arguments.RoleDeleteArgs{}
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIRole.On("Delete", ctx, params).Return(params.ID, nil)
	actual, err := s.Role.Delete(ctx, params)
	s.Nil(err)
	s.Equal(params.ID, actual)
}

func (s *RoleHandlerTestSuite) TestDelete_Fail() {
	var (
		ctx         = context.Background()
		params      = arguments.RoleDeleteArgs{}
		rowEffected int64
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIRole.On("Delete", ctx, params).Return(rowEffected, errors.New("some errors"))
	actual, err := s.Role.Delete(ctx, params)
	s.Equal(rowEffected, actual)
	s.NotNil(err)
}