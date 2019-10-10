// @generated
package permission

import (
	"context"
	"errors"

	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
)

func (s *PermissionHandlerTestSuite) TestGetByID_Success() {
	var (
		ctx   = context.Background()
		param = arguments.PermissionGetByID{
			ID: 1,
		}
		permission = models.Permission{}
	)
	s.MockIPermission.On("GetByID", ctx, param).Return(permission, nil)
	actual, err := s.Permission.GetByID(ctx, param)
	s.Nil(err)
	s.Equal(permission, actual)
}

func (s *PermissionHandlerTestSuite) TestGetByID_Fail() {
	var (
		ctx   = context.Background()
		param = arguments.PermissionGetByID{
			ID: 1,
		}
		permission = models.Permission{}
	)
	s.MockIPermission.On("GetByID", ctx, param).Return(permission, errors.New("some errors"))
	actual, err := s.Permission.GetByID(ctx, param)
	s.Equal(permission, actual)
	s.NotNil(err)
}

func (s *PermissionHandlerTestSuite) TestGetByID_Fail1() {
	var (
		ctx   = context.Background()
		param = arguments.PermissionGetByID{
			ID: 0,
		}
		permission = models.Permission{}
	)
	s.MockIPermission.On("GetByID", ctx, param).Return(permission, errors.New("some errors"))
	actual, err := s.Permission.GetByID(ctx, param)
	s.Equal(permission, actual)
	s.NotNil(err)
}

func (s *PermissionHandlerTestSuite) TestGetByIDs_Success() {
	var (
		ctx   = context.Background()
		param = arguments.PermissionGetByIDs{
			IDs: []int64{1, 2},
		}
		permissions []models.Permission
	)
	s.MockIPermission.On("GetByIDs", ctx, param).Return(permissions, nil)
	actual, err := s.Permission.GetByIDs(ctx, param)
	s.Nil(err)
	s.Equal(permissions, actual)
}

func (s *PermissionHandlerTestSuite) TestGetByIDs_Fail() {
	var (
		ctx   = context.Background()
		param = arguments.PermissionGetByIDs{
			IDs: []int64{1, 2},
		}
		permissions []models.Permission
	)
	s.MockIPermission.On("GetByIDs", ctx, param).Return(permissions, errors.New("some errors"))
	actual, err := s.Permission.GetByIDs(ctx, param)
	s.Equal(permissions, actual)
	s.NotNil(err)
}

func (s *PermissionHandlerTestSuite) TestGetByIDs_Fail1() {
	var (
		ctx         = context.Background()
		param       = arguments.PermissionGetByIDs{}
		permissions []models.Permission
	)
	s.MockIPermission.On("GetByIDs", ctx, param).Return(permissions, errors.New("some errors"))
	actual, err := s.Permission.GetByIDs(ctx, param)
	s.Equal(permissions, actual)
	s.NotNil(err)
}

func (s *PermissionHandlerTestSuite) TestList_Success() {
	var (
		ctx    = context.Background()
		params = arguments.PermissionList{
			ID:          1,
			Name:        "mockString",
			Description: "mockString",
			Status:      "active",
			CreatedBy:   "mockString",
			UpdatedBy:   "mockString",
			LastID:      1,
			PageSize:    10,
		}
		permissions []models.Permission
	)
	s.MockIPermission.On("List", ctx, params).Return(permissions, nil)
	actual, err := s.Permission.List(ctx, params)
	s.Nil(err)
	s.Equal(permissions, actual)
}

func (s *PermissionHandlerTestSuite) TestList_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.PermissionList{
			ID:          1,
			Name:        "mockString",
			Description: "mockString",
			Status:      "active",
			CreatedBy:   "mockString",
			UpdatedBy:   "mockString",
			LastID:      1,
			PageSize:    10,
		}
		permissions []models.Permission
	)
	s.MockIPermission.On("List", ctx, params).Return(permissions, errors.New("some errors"))
	actual, err := s.Permission.List(ctx, params)
	s.Equal(permissions, actual)
	s.NotNil(err)
}

func (s *PermissionHandlerTestSuite) TestList_Fail1() {
	var (
		ctx    = context.Background()
		params = arguments.PermissionList{
			ID:          0,
			Name:        "mockString",
			Description: "mockString",
			Status:      "active",
			CreatedBy:   "mockString",
			UpdatedBy:   "mockString",
			LastID:      1,
			PageSize:    10,
		}
		permissions []models.Permission
	)
	s.MockIPermission.On("List", ctx, params).Return(permissions, errors.New("some errors"))
	actual, err := s.Permission.List(ctx, params)
	s.Equal(permissions, actual)
	s.NotNil(err)
}

func (s *PermissionHandlerTestSuite) TestCount_Success() {
	var (
		ctx    = context.Background()
		params = arguments.PermissionCount{
			ID:          1,
			Name:        "mockString",
			Description: "mockString",
			Status:      "active",
			CreatedBy:   "mockString",
			UpdatedBy:   "mockString",
		}
		count int64
	)
	s.MockIPermission.On("Count", ctx, params).Return(count, nil)
	actual, err := s.Permission.Count(ctx, params)
	s.Nil(err)
	s.Equal(count, actual)
}

func (s *PermissionHandlerTestSuite) TestCount_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.PermissionCount{
			ID:          1,
			Name:        "mockString",
			Description: "mockString",
			Status:      "active",
			CreatedBy:   "mockString",
			UpdatedBy:   "mockString",
		}
		count int64
	)
	s.MockIPermission.On("Count", ctx, params).Return(count, errors.New("some errors"))
	actual, err := s.Permission.Count(ctx, params)
	s.Equal(count, actual)
	s.NotNil(err)
}

func (s *PermissionHandlerTestSuite) TestCount_Fail1() {
	var (
		ctx    = context.Background()
		params = arguments.PermissionCount{
			ID:          0,
			Name:        "mockString",
			Description: "mockString",
			Status:      "active",
			CreatedBy:   "mockString",
			UpdatedBy:   "mockString",
		}
		count int64
	)
	s.MockIPermission.On("Count", ctx, params).Return(count, errors.New("some errors"))
	actual, err := s.Permission.Count(ctx, params)
	s.Equal(count, actual)
	s.NotNil(err)
}

func (s *PermissionHandlerTestSuite) TestInsert_Success() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = arguments.PermissionInsert{
			Name:        "mockString",
			Description: "mockString",
			Status:      "active",
			CreatedBy:   "mockString",
			UpdatedBy:   "mockString",
		}
	)
	permission := models.Permission{
		ID:          sampleID,
		Name:        params.Name,
		Description: params.Description,
		Status:      params.Status,
		CreatedBy:   params.CreatedBy,
		UpdatedBy:   params.UpdatedBy,
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
		params = arguments.PermissionInsert{
			Name:        "mockString",
			Description: "mockString",
			Status:      "active",
			CreatedBy:   "mockString",
			UpdatedBy:   "mockString",
		}
		permission = models.Permission{}
	)
	s.MockIPermission.On("Insert", ctx, params).Return(permission, errors.New("some errors"))
	actual, err := s.Permission.Insert(ctx, params)
	s.Equal(permission, actual)
	s.NotNil(err)
}

func (s *PermissionHandlerTestSuite) TestInsert_Fail1() {
	var (
		ctx    = context.Background()
		params = arguments.PermissionInsert{
			Name:        "mockString",
			Description: "mockString",
			Status:      "active",
			CreatedBy:   "mockString",
			UpdatedBy:   "mockString",
		}
		permission = models.Permission{}
	)
	s.MockIPermission.On("Insert", ctx, params).Return(permission, errors.New("some errors"))
	actual, err := s.Permission.Insert(ctx, params)
	s.Equal(permission, actual)
	s.NotNil(err)
}

func (s *PermissionHandlerTestSuite) TestUpdate_Success() {
	var (
		ctx              = context.Background()
		sampleID   int64 = 1
		mockString       = "mockString"
		status           = "active"
		params           = arguments.PermissionUpdate{
			ID:          &sampleID,
			Name:        &mockString,
			Description: &mockString,
			Status:      &status,
			CreatedBy:   &mockString,
			UpdatedBy:   &mockString,
		}
	)
	permission := models.Permission{
		ID:          *params.ID,
		Name:        *params.Name,
		Description: *params.Description,
		Status:      *params.Status,
		CreatedBy:   *params.CreatedBy,
		UpdatedBy:   *params.UpdatedBy,
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
		status           = "active"
		params           = arguments.PermissionUpdate{
			ID:          &sampleID,
			Name:        &mockString,
			Description: &mockString,
			Status:      &status,
			CreatedBy:   &mockString,
			UpdatedBy:   &mockString,
		}
		permission = models.Permission{}
	)
	s.MockIPermission.On("Update", ctx, params).Return(permission, errors.New("some errors"))
	actual, err := s.Permission.Update(ctx, params)
	s.Equal(permission, actual)
	s.NotNil(err)
}

func (s *PermissionHandlerTestSuite) TestUpdate_Fail1() {
	var (
		ctx        = context.Background()
		sampleID   int64
		mockString = "mockString"
		status     = "active"
		params     = arguments.PermissionUpdate{
			ID:          &sampleID,
			Name:        &mockString,
			Description: &mockString,
			Status:      &status,
			CreatedBy:   &mockString,
			UpdatedBy:   &mockString,
		}
		permission = models.Permission{}
	)
	s.MockIPermission.On("Update", ctx, params).Return(permission, errors.New("some errors"))
	actual, err := s.Permission.Update(ctx, params)
	s.Equal(permission, actual)
	s.NotNil(err)
}

func (s *PermissionHandlerTestSuite) TestDelete_Success() {
	var (
		ctx   = context.Background()
		param = arguments.PermissionDelete{
			ID: 1,
		}
	)
	s.MockIPermission.On("Delete", ctx, param).Return(param.ID, nil)
	actual, err := s.Permission.Delete(ctx, param)
	s.Nil(err)
	s.Equal(param.ID, actual)
}

func (s *PermissionHandlerTestSuite) TestDelete_Fail() {
	var (
		ctx   = context.Background()
		param = arguments.PermissionDelete{
			ID: 1,
		}
		rowEffected int64
	)
	s.MockIPermission.On("Delete", ctx, param).Return(rowEffected, errors.New("some errors"))
	actual, err := s.Permission.Delete(ctx, param)
	s.Equal(rowEffected, actual)
	s.NotNil(err)
}

func (s *PermissionHandlerTestSuite) TestDelete_Fail1() {
	var (
		ctx   = context.Background()
		param = arguments.PermissionDelete{
			ID: 0,
		}
		rowEffected int64
	)
	s.MockIPermission.On("Delete", ctx, param).Return(rowEffected, errors.New("some errors"))
	actual, err := s.Permission.Delete(ctx, param)
	s.Equal(rowEffected, actual)
	s.NotNil(err)
}
