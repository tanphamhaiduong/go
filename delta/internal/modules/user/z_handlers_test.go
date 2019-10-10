// @generated
package user

import (
	"context"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
)

func (s *UserHandlerTestSuite) TestGetByID_Success() {
	var (
		ctx   = context.Background()
		param = arguments.UserGetByID{
			ID: 1,
		}
		user = models.User{}
	)
	s.MockIUser.On("GetByID", ctx, param).Return(user, nil)
	actual, err := s.User.GetByID(ctx, param)
	s.Nil(err)
	s.Equal(user, actual)
}

func (s *UserHandlerTestSuite) TestGetByID_Fail() {
	var (
		ctx   = context.Background()
		param = arguments.UserGetByID{
			ID: 1,
		}
		user = models.User{}
	)
	s.MockIUser.On("GetByID", ctx, param).Return(user, errors.New("some errors"))
	actual, err := s.User.GetByID(ctx, param)
	s.Equal(user, actual)
	s.NotNil(err)
}

func (s *UserHandlerTestSuite) TestGetByID_Fail1() {
	var (
		ctx   = context.Background()
		param = arguments.UserGetByID{
			ID: 0,
		}
		user = models.User{}
	)
	s.MockIUser.On("GetByID", ctx, param).Return(user, errors.New("some errors"))
	actual, err := s.User.GetByID(ctx, param)
	s.Equal(user, actual)
	s.NotNil(err)
}

func (s *UserHandlerTestSuite) TestGetByIDs_Success() {
	var (
		ctx   = context.Background()
		param = arguments.UserGetByIDs{
			IDs: []int64{1, 2},
		}
		users []models.User
	)
	s.MockIUser.On("GetByIDs", ctx, param).Return(users, nil)
	actual, err := s.User.GetByIDs(ctx, param)
	s.Nil(err)
	s.Equal(users, actual)
}

func (s *UserHandlerTestSuite) TestGetByIDs_Fail() {
	var (
		ctx   = context.Background()
		param = arguments.UserGetByIDs{
			IDs: []int64{1, 2},
		}
		users []models.User
	)
	s.MockIUser.On("GetByIDs", ctx, param).Return(users, errors.New("some errors"))
	actual, err := s.User.GetByIDs(ctx, param)
	s.Equal(users, actual)
	s.NotNil(err)
}

func (s *UserHandlerTestSuite) TestGetByIDs_Fail1() {
	var (
		ctx   = context.Background()
		param = arguments.UserGetByIDs{}
		users []models.User
	)
	s.MockIUser.On("GetByIDs", ctx, param).Return(users, errors.New("some errors"))
	actual, err := s.User.GetByIDs(ctx, param)
	s.Equal(users, actual)
	s.NotNil(err)
}

func (s *UserHandlerTestSuite) TestList_Success() {
	var (
		ctx    = context.Background()
		params = arguments.UserList{
			ID:       1,
			Username: "mockString",
			Password: "mockString",
			Name:     "mockString",
			DateOfBirth: mysql.NullTime{
				Time:  time.Time{},
				Valid: true,
			},
			Reference:     "mockString",
			AvatarUrl:     "mockString",
			LicenseNumber: "mockString",
			PhoneNumber:   "mockString",
			Extension:     "mockString",
			TelProvider:   "mockString",
			TelApi:        "mockString",
			SupervisorID:  1,
			RoleID:        1,
			CompanyID:     1,
			Status:        "active",
			CreatedBy:     "mockString",
			UpdatedBy:     "mockString",
			LastID:        1,
			PageSize:      10,
		}
		users []models.User
	)
	s.MockIUser.On("List", ctx, params).Return(users, nil)
	actual, err := s.User.List(ctx, params)
	s.Nil(err)
	s.Equal(users, actual)
}

func (s *UserHandlerTestSuite) TestList_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.UserList{
			ID:       1,
			Username: "mockString",
			Password: "mockString",
			Name:     "mockString",
			DateOfBirth: mysql.NullTime{
				Time:  time.Time{},
				Valid: true,
			},
			Reference:     "mockString",
			AvatarUrl:     "mockString",
			LicenseNumber: "mockString",
			PhoneNumber:   "mockString",
			Extension:     "mockString",
			TelProvider:   "mockString",
			TelApi:        "mockString",
			SupervisorID:  1,
			RoleID:        1,
			CompanyID:     1,
			Status:        "active",
			CreatedBy:     "mockString",
			UpdatedBy:     "mockString",
			LastID:        1,
			PageSize:      10,
		}
		users []models.User
	)
	s.MockIUser.On("List", ctx, params).Return(users, errors.New("some errors"))
	actual, err := s.User.List(ctx, params)
	s.Equal(users, actual)
	s.NotNil(err)
}

func (s *UserHandlerTestSuite) TestList_Fail1() {
	var (
		ctx    = context.Background()
		params = arguments.UserList{
			ID:       0,
			Username: "mockString",
			Password: "mockString",
			Name:     "mockString",
			DateOfBirth: mysql.NullTime{
				Time:  time.Time{},
				Valid: true,
			},
			Reference:     "mockString",
			AvatarUrl:     "mockString",
			LicenseNumber: "mockString",
			PhoneNumber:   "mockString",
			Extension:     "mockString",
			TelProvider:   "mockString",
			TelApi:        "mockString",
			SupervisorID:  0,
			RoleID:        0,
			CompanyID:     0,
			Status:        "active",
			CreatedBy:     "mockString",
			UpdatedBy:     "mockString",
			LastID:        1,
			PageSize:      10,
		}
		users []models.User
	)
	s.MockIUser.On("List", ctx, params).Return(users, errors.New("some errors"))
	actual, err := s.User.List(ctx, params)
	s.Equal(users, actual)
	s.NotNil(err)
}

func (s *UserHandlerTestSuite) TestCount_Success() {
	var (
		ctx    = context.Background()
		params = arguments.UserCount{
			ID:       1,
			Username: "mockString",
			Password: "mockString",
			Name:     "mockString",
			DateOfBirth: mysql.NullTime{
				Time:  time.Time{},
				Valid: true,
			},
			Reference:     "mockString",
			AvatarUrl:     "mockString",
			LicenseNumber: "mockString",
			PhoneNumber:   "mockString",
			Extension:     "mockString",
			TelProvider:   "mockString",
			TelApi:        "mockString",
			SupervisorID:  1,
			RoleID:        1,
			CompanyID:     1,
			Status:        "active",
			CreatedBy:     "mockString",
			UpdatedBy:     "mockString",
		}
		count int64
	)
	s.MockIUser.On("Count", ctx, params).Return(count, nil)
	actual, err := s.User.Count(ctx, params)
	s.Nil(err)
	s.Equal(count, actual)
}

func (s *UserHandlerTestSuite) TestCount_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.UserCount{
			ID:       1,
			Username: "mockString",
			Password: "mockString",
			Name:     "mockString",
			DateOfBirth: mysql.NullTime{
				Time:  time.Time{},
				Valid: true,
			},
			Reference:     "mockString",
			AvatarUrl:     "mockString",
			LicenseNumber: "mockString",
			PhoneNumber:   "mockString",
			Extension:     "mockString",
			TelProvider:   "mockString",
			TelApi:        "mockString",
			SupervisorID:  1,
			RoleID:        1,
			CompanyID:     1,
			Status:        "active",
			CreatedBy:     "mockString",
			UpdatedBy:     "mockString",
		}
		count int64
	)
	s.MockIUser.On("Count", ctx, params).Return(count, errors.New("some errors"))
	actual, err := s.User.Count(ctx, params)
	s.Equal(count, actual)
	s.NotNil(err)
}

func (s *UserHandlerTestSuite) TestCount_Fail1() {
	var (
		ctx    = context.Background()
		params = arguments.UserCount{
			ID:       0,
			Username: "mockString",
			Password: "mockString",
			Name:     "mockString",
			DateOfBirth: mysql.NullTime{
				Time:  time.Time{},
				Valid: true,
			},
			Reference:     "mockString",
			AvatarUrl:     "mockString",
			LicenseNumber: "mockString",
			PhoneNumber:   "mockString",
			Extension:     "mockString",
			TelProvider:   "mockString",
			TelApi:        "mockString",
			SupervisorID:  0,
			RoleID:        0,
			CompanyID:     0,
			Status:        "active",
			CreatedBy:     "mockString",
			UpdatedBy:     "mockString",
		}
		count int64
	)
	s.MockIUser.On("Count", ctx, params).Return(count, errors.New("some errors"))
	actual, err := s.User.Count(ctx, params)
	s.Equal(count, actual)
	s.NotNil(err)
}

func (s *UserHandlerTestSuite) TestInsert_Success() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = arguments.UserInsert{
			Username: "mockString",
			Password: "mockString",
			Name:     "mockString",
			DateOfBirth: mysql.NullTime{
				Time:  time.Time{},
				Valid: true,
			},
			Reference:     "mockString",
			AvatarUrl:     "mockString",
			LicenseNumber: "mockString",
			PhoneNumber:   "mockString",
			Extension:     "mockString",
			TelProvider:   "mockString",
			TelApi:        "mockString",
			SupervisorID:  1,
			RoleID:        1,
			CompanyID:     1,
			Status:        "active",
			CreatedBy:     "mockString",
			UpdatedBy:     "mockString",
		}
	)
	user := models.User{
		ID:            sampleID,
		Username:      params.Username,
		Password:      params.Password,
		Name:          params.Name,
		DateOfBirth:   params.DateOfBirth,
		Reference:     params.Reference,
		AvatarUrl:     params.AvatarUrl,
		LicenseNumber: params.LicenseNumber,
		PhoneNumber:   params.PhoneNumber,
		Extension:     params.Extension,
		TelProvider:   params.TelProvider,
		TelApi:        params.TelApi,
		SupervisorID:  params.SupervisorID,
		RoleID:        params.RoleID,
		CompanyID:     params.CompanyID,
		Status:        params.Status,
		CreatedBy:     params.CreatedBy,
		UpdatedBy:     params.UpdatedBy,
	}
	// Mock Insert
	s.MockIUser.On("Insert", ctx, params).Return(user, nil)
	// Mock GetByID
	s.MockIUser.On("GetByID", ctx, params).Return(user, nil)

	actual, err := s.User.Insert(ctx, params)
	s.Nil(err)
	s.Equal(user, actual)
}

func (s *UserHandlerTestSuite) TestInsert_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.UserInsert{
			Username: "mockString",
			Password: "mockString",
			Name:     "mockString",
			DateOfBirth: mysql.NullTime{
				Time:  time.Time{},
				Valid: true,
			},
			Reference:     "mockString",
			AvatarUrl:     "mockString",
			LicenseNumber: "mockString",
			PhoneNumber:   "mockString",
			Extension:     "mockString",
			TelProvider:   "mockString",
			TelApi:        "mockString",
			SupervisorID:  1,
			RoleID:        1,
			CompanyID:     1,
			Status:        "active",
			CreatedBy:     "mockString",
			UpdatedBy:     "mockString",
		}
		user = models.User{}
	)
	s.MockIUser.On("Insert", ctx, params).Return(user, errors.New("some errors"))
	actual, err := s.User.Insert(ctx, params)
	s.Equal(user, actual)
	s.NotNil(err)
}

func (s *UserHandlerTestSuite) TestInsert_Fail1() {
	var (
		ctx    = context.Background()
		params = arguments.UserInsert{
			Username: "mockString",
			Password: "mockString",
			Name:     "mockString",
			DateOfBirth: mysql.NullTime{
				Time:  time.Time{},
				Valid: true,
			},
			Reference:     "mockString",
			AvatarUrl:     "mockString",
			LicenseNumber: "mockString",
			PhoneNumber:   "mockString",
			Extension:     "mockString",
			TelProvider:   "mockString",
			TelApi:        "mockString",
			SupervisorID:  0,
			RoleID:        0,
			CompanyID:     0,
			Status:        "active",
			CreatedBy:     "mockString",
			UpdatedBy:     "mockString",
		}
		user = models.User{}
	)
	s.MockIUser.On("Insert", ctx, params).Return(user, errors.New("some errors"))
	actual, err := s.User.Insert(ctx, params)
	s.Equal(user, actual)
	s.NotNil(err)
}

func (s *UserHandlerTestSuite) TestUpdate_Success() {
	var (
		ctx              = context.Background()
		sampleID   int64 = 1
		mockString       = "mockString"
		status           = "active"
		params           = arguments.UserUpdate{
			ID:       &sampleID,
			Username: &mockString,
			Password: &mockString,
			Name:     &mockString,
			DateOfBirth: &mysql.NullTime{
				Time:  time.Time{},
				Valid: true,
			},
			Reference:     &mockString,
			AvatarUrl:     &mockString,
			LicenseNumber: &mockString,
			PhoneNumber:   &mockString,
			Extension:     &mockString,
			TelProvider:   &mockString,
			TelApi:        &mockString,
			SupervisorID:  &sampleID,
			RoleID:        &sampleID,
			CompanyID:     &sampleID,
			Status:        &status,
			CreatedBy:     &mockString,
			UpdatedBy:     &mockString,
		}
	)
	user := models.User{
		ID:            *params.ID,
		Username:      *params.Username,
		Password:      *params.Password,
		Name:          *params.Name,
		DateOfBirth:   *params.DateOfBirth,
		Reference:     *params.Reference,
		AvatarUrl:     *params.AvatarUrl,
		LicenseNumber: *params.LicenseNumber,
		PhoneNumber:   *params.PhoneNumber,
		Extension:     *params.Extension,
		TelProvider:   *params.TelProvider,
		TelApi:        *params.TelApi,
		SupervisorID:  *params.SupervisorID,
		RoleID:        *params.RoleID,
		CompanyID:     *params.CompanyID,
		Status:        *params.Status,
		CreatedBy:     *params.CreatedBy,
		UpdatedBy:     *params.UpdatedBy,
	}
	// Mock Insert
	s.MockIUser.On("Update", ctx, params).Return(user, nil)
	// Mock GetByID
	s.MockIUser.On("GetByID", ctx, params).Return(user, nil)

	actual, err := s.User.Update(ctx, params)
	s.Nil(err)
	s.Equal(user, actual)
}

func (s *UserHandlerTestSuite) TestUpdate_Fail() {
	var (
		ctx              = context.Background()
		sampleID   int64 = 1
		mockString       = "mockString"
		status           = "active"
		params           = arguments.UserUpdate{
			ID:       &sampleID,
			Username: &mockString,
			Password: &mockString,
			Name:     &mockString,
			DateOfBirth: &mysql.NullTime{
				Time:  time.Time{},
				Valid: true,
			},
			Reference:     &mockString,
			AvatarUrl:     &mockString,
			LicenseNumber: &mockString,
			PhoneNumber:   &mockString,
			Extension:     &mockString,
			TelProvider:   &mockString,
			TelApi:        &mockString,
			SupervisorID:  &sampleID,
			RoleID:        &sampleID,
			CompanyID:     &sampleID,
			Status:        &status,
			CreatedBy:     &mockString,
			UpdatedBy:     &mockString,
		}
		user = models.User{}
	)
	s.MockIUser.On("Update", ctx, params).Return(user, errors.New("some errors"))
	actual, err := s.User.Update(ctx, params)
	s.Equal(user, actual)
	s.NotNil(err)
}

func (s *UserHandlerTestSuite) TestUpdate_Fail1() {
	var (
		ctx        = context.Background()
		sampleID   int64
		mockString = "mockString"
		status     = "active"
		params     = arguments.UserUpdate{
			ID:       &sampleID,
			Username: &mockString,
			Password: &mockString,
			Name:     &mockString,
			DateOfBirth: &mysql.NullTime{
				Time:  time.Time{},
				Valid: true,
			},
			Reference:     &mockString,
			AvatarUrl:     &mockString,
			LicenseNumber: &mockString,
			PhoneNumber:   &mockString,
			Extension:     &mockString,
			TelProvider:   &mockString,
			TelApi:        &mockString,
			SupervisorID:  &sampleID,
			RoleID:        &sampleID,
			CompanyID:     &sampleID,
			Status:        &status,
			CreatedBy:     &mockString,
			UpdatedBy:     &mockString,
		}
		user = models.User{}
	)
	s.MockIUser.On("Update", ctx, params).Return(user, errors.New("some errors"))
	actual, err := s.User.Update(ctx, params)
	s.Equal(user, actual)
	s.NotNil(err)
}

func (s *UserHandlerTestSuite) TestDelete_Success() {
	var (
		ctx   = context.Background()
		param = arguments.UserDelete{
			ID: 1,
		}
	)
	s.MockIUser.On("Delete", ctx, param).Return(param.ID, nil)
	actual, err := s.User.Delete(ctx, param)
	s.Nil(err)
	s.Equal(param.ID, actual)
}

func (s *UserHandlerTestSuite) TestDelete_Fail() {
	var (
		ctx   = context.Background()
		param = arguments.UserDelete{
			ID: 1,
		}
		rowEffected int64
	)
	s.MockIUser.On("Delete", ctx, param).Return(rowEffected, errors.New("some errors"))
	actual, err := s.User.Delete(ctx, param)
	s.Equal(rowEffected, actual)
	s.NotNil(err)
}

func (s *UserHandlerTestSuite) TestDelete_Fail1() {
	var (
		ctx   = context.Background()
		param = arguments.UserDelete{
			ID: 0,
		}
		rowEffected int64
	)
	s.MockIUser.On("Delete", ctx, param).Return(rowEffected, errors.New("some errors"))
	actual, err := s.User.Delete(ctx, param)
	s.Equal(rowEffected, actual)
	s.NotNil(err)
}
