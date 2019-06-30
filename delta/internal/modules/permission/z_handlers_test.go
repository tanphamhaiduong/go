// @generated
package permission

import (
	"context"
	"errors"
	"log"

	"github.com/bxcodec/faker"

	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
)

func (s *PermissionHandlerTestSuite) TestGetByID_Success() {
	var (
		ctx    = context.Background()
		params = arguments.PermissionGetByIDArgs{
			ID: 1,
		}
		permission = models.Permission{}
	)
	s.MockIPermission.On("GetByID", ctx, params).Return(permission, nil)
	actual, err := s.Permission.GetByID(ctx, params)
	s.Nil(err)
	s.Equal(permission, actual)
}

func (s *PermissionHandlerTestSuite) TestGetByID_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.PermissionGetByIDArgs{
			ID: 1,
		}
		permission = models.Permission{}
	)
	s.MockIPermission.On("GetByID", ctx, params).Return(permission, errors.New("some errors"))
	actual, err := s.Permission.GetByID(ctx, params)
	s.Equal(permission, actual)
	s.NotNil(err)
}

func (s *PermissionHandlerTestSuite) TestGetByIDs_Success() {
	var (
		ctx    = context.Background()
		params = arguments.PermissionGetByIDsArgs{
			IDs: []int64{1, 2},
		}
		permissions []models.Permission
	)
	s.MockIPermission.On("GetByIDs", ctx, params).Return(permissions, nil)
	actual, err := s.Permission.GetByIDs(ctx, params)
	s.Nil(err)
	s.Equal(permissions, actual)
}

func (s *PermissionHandlerTestSuite) TestGetByIDs_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.PermissionGetByIDsArgs{
			IDs: []int64{1, 2},
		}
		permissions []models.Permission
	)
	s.MockIPermission.On("GetByIDs", ctx, params).Return(permissions, errors.New("some errors"))
	actual, err := s.Permission.GetByIDs(ctx, params)
	s.Equal(permissions, actual)
	s.NotNil(err)
}

func (s *PermissionHandlerTestSuite) TestList_Success() {
	var (
		ctx         = context.Background()
		params      = arguments.PermissionListArgs{}
		permissions []models.Permission
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	params.Status = "active"
	params.Page = 1
	params.PageSize = 10
	s.MockIPermission.On("List", ctx, params).Return(permissions, nil)
	actual, err := s.Permission.List(ctx, params)
	s.Nil(err)
	s.Equal(permissions, actual)
}

func (s *PermissionHandlerTestSuite) TestList_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.PermissionListArgs{
			Page:     1,
			PageSize: 10,
		}
		permissions []models.Permission
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIPermission.On("List", ctx, params).Return(permissions, errors.New("some errors"))
	actual, err := s.Permission.List(ctx, params)
	s.Equal(permissions, actual)
	s.NotNil(err)
}

func (s *PermissionHandlerTestSuite) TestCount_Success() {
	var (
		ctx    = context.Background()
		params = arguments.PermissionCountArgs{}
		count  int64
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	params.Status = "active"
	s.MockIPermission.On("Count", ctx, params).Return(count, nil)
	actual, err := s.Permission.Count(ctx, params)
	s.Nil(err)
	s.Equal(count, actual)
}

func (s *PermissionHandlerTestSuite) TestCount_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.PermissionCountArgs{}
		count  int64
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIPermission.On("Count", ctx, params).Return(count, errors.New("some errors"))
	actual, err := s.Permission.Count(ctx, params)
	s.Equal(count, actual)
	s.NotNil(err)
}

func (s *PermissionHandlerTestSuite) TestInsert_Success() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = arguments.PermissionInsertArgs{}
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	params.Status = "active"
	permission := models.Permission{
		ID:        sampleID,
		Name:      params.Name,
		Status:    params.Status,
		CreatedBy: params.CreatedBy,
		UpdatedBy: params.UpdatedBy,
	}
	// Mock Insert
	s.MockIPermission.On("Insert", ctx, params).Return(permission, nil)
	// Mock GetByID
	s.MockIPermission.On("GetByID", ctx, params).Return(permission, nil)

	actual, err := s.Permission.Insert(ctx, params)
	s.Nil(err)
	s.Equal(permission, actual)
}

func (s *PermissionHandlerTestSuite) TestInsert_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.PermissionInsertArgs{
			Name:      "mockString",
			Status:    "mockString",
			CreatedBy: "mockString",
			UpdatedBy: "mockString",
		}
		permission = models.Permission{}
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIPermission.On("Insert", ctx, params).Return(permission, errors.New("some errors"))
	actual, err := s.Permission.Insert(ctx, params)
	s.Equal(permission, actual)
	s.NotNil(err)
}

func (s *PermissionHandlerTestSuite) TestUpdate_Success() {
	var (
		ctx    = context.Background()
		params = arguments.PermissionUpdateArgs{}
		status = "active"
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	params.Status = &status
	permission := models.Permission{
		ID:        *params.ID,
		Name:      *params.Name,
		Status:    *params.Status,
		CreatedBy: *params.CreatedBy,
		UpdatedBy: *params.UpdatedBy,
	}
	// Mock Insert
	s.MockIPermission.On("Update", ctx, params).Return(permission, nil)
	// Mock GetByID
	s.MockIPermission.On("GetByID", ctx, params).Return(permission, nil)

	actual, err := s.Permission.Update(ctx, params)
	s.Nil(err)
	s.Equal(permission, actual)
}

func (s *PermissionHandlerTestSuite) TestUpdate_Fail() {
	var (
		ctx              = context.Background()
		sampleID   int64 = 1
		mockString       = "mockString"
		params           = arguments.PermissionUpdateArgs{
			ID:        &sampleID,
			Name:      &mockString,
			Status:    &mockString,
			CreatedBy: &mockString,
			UpdatedBy: &mockString,
		}
		permission = models.Permission{}
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIPermission.On("Update", ctx, params).Return(permission, errors.New("some errors"))
	actual, err := s.Permission.Update(ctx, params)
	s.Equal(permission, actual)
	s.NotNil(err)
}

func (s *PermissionHandlerTestSuite) TestDelete_Success() {
	var (
		ctx    = context.Background()
		params = arguments.PermissionDeleteArgs{}
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIPermission.On("Delete", ctx, params).Return(params.ID, nil)
	actual, err := s.Permission.Delete(ctx, params)
	s.Nil(err)
	s.Equal(params.ID, actual)
}

func (s *PermissionHandlerTestSuite) TestDelete_Fail() {
	var (
		ctx         = context.Background()
		params      = arguments.PermissionDeleteArgs{}
		rowEffected int64
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIPermission.On("Delete", ctx, params).Return(rowEffected, errors.New("some errors"))
	actual, err := s.Permission.Delete(ctx, params)
	s.Equal(rowEffected, actual)
	s.NotNil(err)
}
