// @generated
package user

import (
	"context"
	"errors"
	"log"

	"github.com/bxcodec/faker"
	"github.com/go-sql-driver/mysql"

	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
)

func (s *UserHandlerTestSuite) TestGetByID_Success() {
	var (
		ctx    = context.Background()
		params = arguments.UserGetByIDArgs{
			ID: 1,
		}
		user = models.User{}
	)
	s.MockIUser.On("GetByID", ctx, params).Return(user, nil)
	actual, err := s.User.GetByID(ctx, params)
	s.Nil(err)
	s.Equal(user, actual)
}

func (s *UserHandlerTestSuite) TestGetByID_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.UserGetByIDArgs{
			ID: 1,
		}
		user = models.User{}
	)
	s.MockIUser.On("GetByID", ctx, params).Return(user, errors.New("some errors"))
	actual, err := s.User.GetByID(ctx, params)
	s.Equal(user, actual)
	s.NotNil(err)
}

func (s *UserHandlerTestSuite) TestGetByIDs_Success() {
	var (
		ctx    = context.Background()
		params = arguments.UserGetByIDsArgs{
			IDs: []int64{1, 2},
		}
		users []models.User
	)
	s.MockIUser.On("GetByIDs", ctx, params).Return(users, nil)
	actual, err := s.User.GetByIDs(ctx, params)
	s.Nil(err)
	s.Equal(users, actual)
}

func (s *UserHandlerTestSuite) TestGetByIDs_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.UserGetByIDsArgs{
			IDs: []int64{1, 2},
		}
		users []models.User
	)
	s.MockIUser.On("GetByIDs", ctx, params).Return(users, errors.New("some errors"))
	actual, err := s.User.GetByIDs(ctx, params)
	s.Equal(users, actual)
	s.NotNil(err)
}

func (s *UserHandlerTestSuite) TestList_Success() {
	var (
		ctx    = context.Background()
		params = arguments.UserListArgs{}
		users  []models.User
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	params.Status = "active"
	params.Page = 1
	params.PageSize = 10
	s.MockIUser.On("List", ctx, params).Return(users, nil)
	actual, err := s.User.List(ctx, params)
	s.Nil(err)
	s.Equal(users, actual)
}

func (s *UserHandlerTestSuite) TestList_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.UserListArgs{
			Page:     1,
			PageSize: 10,
		}
		users []models.User
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIUser.On("List", ctx, params).Return(users, errors.New("some errors"))
	actual, err := s.User.List(ctx, params)
	s.Equal(users, actual)
	s.NotNil(err)
}

func (s *UserHandlerTestSuite) TestCount_Success() {
	var (
		ctx    = context.Background()
		params = arguments.UserCountArgs{}
		count  int64
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	params.Status = "active"
	s.MockIUser.On("Count", ctx, params).Return(count, nil)
	actual, err := s.User.Count(ctx, params)
	s.Nil(err)
	s.Equal(count, actual)
}

func (s *UserHandlerTestSuite) TestCount_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.UserCountArgs{}
		count  int64
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIUser.On("Count", ctx, params).Return(count, errors.New("some errors"))
	actual, err := s.User.Count(ctx, params)
	s.Equal(count, actual)
	s.NotNil(err)
}

func (s *UserHandlerTestSuite) TestInsert_Success() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = arguments.UserInsertArgs{}
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	params.Status = "active"
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
		SupervisorId:  params.SupervisorId,
		RoleId:        params.RoleId,
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
		params = arguments.UserInsertArgs{
			Username:      "mockString",
			Password:      "mockString",
			Name:          "mockString",
			DateOfBirth:   mysql.NullTime{},
			Reference:     "mockString",
			AvatarUrl:     "mockString",
			LicenseNumber: "mockString",
			PhoneNumber:   "mockString",
			Extension:     "mockString",
			TelProvider:   "mockString",
			TelApi:        "mockString",
			SupervisorId:  0,
			RoleId:        0,
			CompanyID:     0,
			Status:        "mockString",
			CreatedBy:     "mockString",
			UpdatedBy:     "mockString",
		}
		user = models.User{}
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIUser.On("Insert", ctx, params).Return(user, errors.New("some errors"))
	actual, err := s.User.Insert(ctx, params)
	s.Equal(user, actual)
	s.NotNil(err)
}

func (s *UserHandlerTestSuite) TestUpdate_Success() {
	var (
		ctx    = context.Background()
		params = arguments.UserUpdateArgs{}
		status = "active"
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	params.Status = &status
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
		SupervisorId:  *params.SupervisorId,
		RoleId:        *params.RoleId,
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
		params           = arguments.UserUpdateArgs{
			ID:            &sampleID,
			Username:      &mockString,
			Password:      &mockString,
			Name:          &mockString,
			DateOfBirth:   &mysql.NullTime{},
			Reference:     &mockString,
			AvatarUrl:     &mockString,
			LicenseNumber: &mockString,
			PhoneNumber:   &mockString,
			Extension:     &mockString,
			TelProvider:   &mockString,
			TelApi:        &mockString,
			SupervisorId:  &sampleID,
			RoleId:        &sampleID,
			CompanyID:     &sampleID,
			Status:        &mockString,
			CreatedBy:     &mockString,
			UpdatedBy:     &mockString,
		}
		user = models.User{}
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIUser.On("Update", ctx, params).Return(user, errors.New("some errors"))
	actual, err := s.User.Update(ctx, params)
	s.Equal(user, actual)
	s.NotNil(err)
}

func (s *UserHandlerTestSuite) TestDelete_Success() {
	var (
		ctx    = context.Background()
		params = arguments.UserDeleteArgs{}
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIUser.On("Delete", ctx, params).Return(params.ID, nil)
	actual, err := s.User.Delete(ctx, params)
	s.Nil(err)
	s.Equal(params.ID, actual)
}

func (s *UserHandlerTestSuite) TestDelete_Fail() {
	var (
		ctx         = context.Background()
		params      = arguments.UserDeleteArgs{}
		rowEffected int64
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIUser.On("Delete", ctx, params).Return(rowEffected, errors.New("some errors"))
	actual, err := s.User.Delete(ctx, params)
	s.Equal(rowEffected, actual)
	s.NotNil(err)
}
