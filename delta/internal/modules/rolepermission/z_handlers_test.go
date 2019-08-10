// @generated
package rolepermission

import (
	"context"
	"errors"

	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
)

func (s *RolePermissionHandlerTestSuite) TestGetByID_Success() {
	var (
		ctx   = context.Background()
		param = arguments.RolePermissionGetByIDArgs{
			ID: 1,
		}
		rolepermission = models.RolePermission{}
	)
	s.MockIRolePermission.On("GetByID", ctx, param).Return(rolepermission, nil)
	actual, err := s.RolePermission.GetByID(ctx, param)
	s.Nil(err)
	s.Equal(rolepermission, actual)
}

func (s *RolePermissionHandlerTestSuite) TestGetByID_Fail() {
	var (
		ctx   = context.Background()
		param = arguments.RolePermissionGetByIDArgs{
			ID: 1,
		}
		rolepermission = models.RolePermission{}
	)
	s.MockIRolePermission.On("GetByID", ctx, param).Return(rolepermission, errors.New("some errors"))
	actual, err := s.RolePermission.GetByID(ctx, param)
	s.Equal(rolepermission, actual)
	s.NotNil(err)
}

func (s *RolePermissionHandlerTestSuite) TestGetByID_Fail1() {
	var (
		ctx   = context.Background()
		param = arguments.RolePermissionGetByIDArgs{
			ID: 0,
		}
		rolepermission = models.RolePermission{}
	)
	s.MockIRolePermission.On("GetByID", ctx, param).Return(rolepermission, errors.New("some errors"))
	actual, err := s.RolePermission.GetByID(ctx, param)
	s.Equal(rolepermission, actual)
	s.NotNil(err)
}

func (s *RolePermissionHandlerTestSuite) TestGetByIDs_Success() {
	var (
		ctx   = context.Background()
		param = arguments.RolePermissionGetByIDsArgs{
			IDs: []int64{1, 2},
		}
		rolepermissions []models.RolePermission
	)
	s.MockIRolePermission.On("GetByIDs", ctx, param).Return(rolepermissions, nil)
	actual, err := s.RolePermission.GetByIDs(ctx, param)
	s.Nil(err)
	s.Equal(rolepermissions, actual)
}

func (s *RolePermissionHandlerTestSuite) TestGetByIDs_Fail() {
	var (
		ctx   = context.Background()
		param = arguments.RolePermissionGetByIDsArgs{
			IDs: []int64{1, 2},
		}
		rolepermissions []models.RolePermission
	)
	s.MockIRolePermission.On("GetByIDs", ctx, param).Return(rolepermissions, errors.New("some errors"))
	actual, err := s.RolePermission.GetByIDs(ctx, param)
	s.Equal(rolepermissions, actual)
	s.NotNil(err)
}

func (s *RolePermissionHandlerTestSuite) TestGetByIDs_Fail1() {
	var (
		ctx             = context.Background()
		param           = arguments.RolePermissionGetByIDsArgs{}
		rolepermissions []models.RolePermission
	)
	s.MockIRolePermission.On("GetByIDs", ctx, param).Return(rolepermissions, errors.New("some errors"))
	actual, err := s.RolePermission.GetByIDs(ctx, param)
	s.Equal(rolepermissions, actual)
	s.NotNil(err)
}

func (s *RolePermissionHandlerTestSuite) TestList_Success() {
	var (
		ctx    = context.Background()
		params = arguments.RolePermissionListArgs{
			ID:           1,
			RoleID:       1,
			PermissionID: 1,
			CreatedBy:    "mockString",
			UpdatedBy:    "mockString",
			Page:         1,
			PageSize:     10,
		}
		rolepermissions []models.RolePermission
	)
	s.MockIRolePermission.On("List", ctx, params).Return(rolepermissions, nil)
	actual, err := s.RolePermission.List(ctx, params)
	s.Nil(err)
	s.Equal(rolepermissions, actual)
}

func (s *RolePermissionHandlerTestSuite) TestList_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.RolePermissionListArgs{
			ID:           1,
			RoleID:       1,
			PermissionID: 1,
			CreatedBy:    "mockString",
			UpdatedBy:    "mockString",
			Page:         1,
			PageSize:     10,
		}
		rolepermissions []models.RolePermission
	)
	s.MockIRolePermission.On("List", ctx, params).Return(rolepermissions, errors.New("some errors"))
	actual, err := s.RolePermission.List(ctx, params)
	s.Equal(rolepermissions, actual)
	s.NotNil(err)
}

func (s *RolePermissionHandlerTestSuite) TestList_Fail1() {
	var (
		ctx    = context.Background()
		params = arguments.RolePermissionListArgs{
			ID:           0,
			RoleID:       0,
			PermissionID: 0,
			CreatedBy:    "mockString",
			UpdatedBy:    "mockString",
			Page:         1,
			PageSize:     10,
		}
		rolepermissions []models.RolePermission
	)
	s.MockIRolePermission.On("List", ctx, params).Return(rolepermissions, errors.New("some errors"))
	actual, err := s.RolePermission.List(ctx, params)
	s.Equal(rolepermissions, actual)
	s.NotNil(err)
}

func (s *RolePermissionHandlerTestSuite) TestCount_Success() {
	var (
		ctx    = context.Background()
		params = arguments.RolePermissionCountArgs{
			ID:           1,
			RoleID:       1,
			PermissionID: 1,
			CreatedBy:    "mockString",
			UpdatedBy:    "mockString",
		}
		count int64
	)
	s.MockIRolePermission.On("Count", ctx, params).Return(count, nil)
	actual, err := s.RolePermission.Count(ctx, params)
	s.Nil(err)
	s.Equal(count, actual)
}

func (s *RolePermissionHandlerTestSuite) TestCount_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.RolePermissionCountArgs{
			ID:           1,
			RoleID:       1,
			PermissionID: 1,
			CreatedBy:    "mockString",
			UpdatedBy:    "mockString",
		}
		count int64
	)
	s.MockIRolePermission.On("Count", ctx, params).Return(count, errors.New("some errors"))
	actual, err := s.RolePermission.Count(ctx, params)
	s.Equal(count, actual)
	s.NotNil(err)
}

func (s *RolePermissionHandlerTestSuite) TestCount_Fail1() {
	var (
		ctx    = context.Background()
		params = arguments.RolePermissionCountArgs{
			ID:           0,
			RoleID:       0,
			PermissionID: 0,
			CreatedBy:    "mockString",
			UpdatedBy:    "mockString",
		}
		count int64
	)
	s.MockIRolePermission.On("Count", ctx, params).Return(count, errors.New("some errors"))
	actual, err := s.RolePermission.Count(ctx, params)
	s.Equal(count, actual)
	s.NotNil(err)
}

func (s *RolePermissionHandlerTestSuite) TestInsert_Success() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = arguments.RolePermissionInsertArgs{
			RoleID:       1,
			PermissionID: 1,
			CreatedBy:    "mockString",
			UpdatedBy:    "mockString",
		}
	)
	rolepermission := models.RolePermission{
		ID:           sampleID,
		RoleID:       params.RoleID,
		PermissionID: params.PermissionID,
		CreatedBy:    params.CreatedBy,
		UpdatedBy:    params.UpdatedBy,
	}
	// Mock Insert
	s.MockIRolePermission.On("Insert", ctx, params).Return(rolepermission, nil)
	// Mock GetByID
	s.MockIRolePermission.On("GetByID", ctx, params).Return(rolepermission, nil)

	actual, err := s.RolePermission.Insert(ctx, params)
	s.Nil(err)
	s.Equal(rolepermission, actual)
}

func (s *RolePermissionHandlerTestSuite) TestInsert_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.RolePermissionInsertArgs{
			RoleID:       1,
			PermissionID: 1,
			CreatedBy:    "mockString",
			UpdatedBy:    "mockString",
		}
		rolepermission = models.RolePermission{}
	)
	s.MockIRolePermission.On("Insert", ctx, params).Return(rolepermission, errors.New("some errors"))
	actual, err := s.RolePermission.Insert(ctx, params)
	s.Equal(rolepermission, actual)
	s.NotNil(err)
}

func (s *RolePermissionHandlerTestSuite) TestInsert_Fail1() {
	var (
		ctx    = context.Background()
		params = arguments.RolePermissionInsertArgs{
			RoleID:       0,
			PermissionID: 0,
			CreatedBy:    "mockString",
			UpdatedBy:    "mockString",
		}
		rolepermission = models.RolePermission{}
	)
	s.MockIRolePermission.On("Insert", ctx, params).Return(rolepermission, errors.New("some errors"))
	actual, err := s.RolePermission.Insert(ctx, params)
	s.Equal(rolepermission, actual)
	s.NotNil(err)
}

func (s *RolePermissionHandlerTestSuite) TestUpdate_Success() {
	var (
		ctx              = context.Background()
		sampleID   int64 = 1
		mockString       = "mockString"
		params           = arguments.RolePermissionUpdateArgs{
			ID:           &sampleID,
			RoleID:       &sampleID,
			PermissionID: &sampleID,
			CreatedBy:    &mockString,
			UpdatedBy:    &mockString,
		}
	)
	rolepermission := models.RolePermission{
		ID:           *params.ID,
		RoleID:       *params.RoleID,
		PermissionID: *params.PermissionID,
		CreatedBy:    *params.CreatedBy,
		UpdatedBy:    *params.UpdatedBy,
	}
	// Mock Insert
	s.MockIRolePermission.On("Update", ctx, params).Return(rolepermission, nil)
	// Mock GetByID
	s.MockIRolePermission.On("GetByID", ctx, params).Return(rolepermission, nil)

	actual, err := s.RolePermission.Update(ctx, params)
	s.Nil(err)
	s.Equal(rolepermission, actual)
}

func (s *RolePermissionHandlerTestSuite) TestUpdate_Fail() {
	var (
		ctx              = context.Background()
		sampleID   int64 = 1
		mockString       = "mockString"
		params           = arguments.RolePermissionUpdateArgs{
			ID:           &sampleID,
			RoleID:       &sampleID,
			PermissionID: &sampleID,
			CreatedBy:    &mockString,
			UpdatedBy:    &mockString,
		}
		rolepermission = models.RolePermission{}
	)
	s.MockIRolePermission.On("Update", ctx, params).Return(rolepermission, errors.New("some errors"))
	actual, err := s.RolePermission.Update(ctx, params)
	s.Equal(rolepermission, actual)
	s.NotNil(err)
}

func (s *RolePermissionHandlerTestSuite) TestUpdate_Fail1() {
	var (
		ctx        = context.Background()
		sampleID   int64
		mockString = "mockString"
		params     = arguments.RolePermissionUpdateArgs{
			ID:           &sampleID,
			RoleID:       &sampleID,
			PermissionID: &sampleID,
			CreatedBy:    &mockString,
			UpdatedBy:    &mockString,
		}
		rolepermission = models.RolePermission{}
	)
	s.MockIRolePermission.On("Update", ctx, params).Return(rolepermission, errors.New("some errors"))
	actual, err := s.RolePermission.Update(ctx, params)
	s.Equal(rolepermission, actual)
	s.NotNil(err)
}

func (s *RolePermissionHandlerTestSuite) TestDelete_Success() {
	var (
		ctx   = context.Background()
		param = arguments.RolePermissionDeleteArgs{
			ID: 1,
		}
	)
	s.MockIRolePermission.On("Delete", ctx, param).Return(param.ID, nil)
	actual, err := s.RolePermission.Delete(ctx, param)
	s.Nil(err)
	s.Equal(param.ID, actual)
}

func (s *RolePermissionHandlerTestSuite) TestDelete_Fail() {
	var (
		ctx   = context.Background()
		param = arguments.RolePermissionDeleteArgs{
			ID: 1,
		}
		rowEffected int64
	)
	s.MockIRolePermission.On("Delete", ctx, param).Return(rowEffected, errors.New("some errors"))
	actual, err := s.RolePermission.Delete(ctx, param)
	s.Equal(rowEffected, actual)
	s.NotNil(err)
}

func (s *RolePermissionHandlerTestSuite) TestDelete_Fail1() {
	var (
		ctx   = context.Background()
		param = arguments.RolePermissionDeleteArgs{
			ID: 0,
		}
		rowEffected int64
	)
	s.MockIRolePermission.On("Delete", ctx, param).Return(rowEffected, errors.New("some errors"))
	actual, err := s.RolePermission.Delete(ctx, param)
	s.Equal(rowEffected, actual)
	s.NotNil(err)
}
