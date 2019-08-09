// @generated
package user

import (
	"context"
	"errors"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/bxcodec/faker"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/mock"
	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
	"github.com/tanphamhaiduong/go/delta/internal/utils"
)

func (s *UserRepositoryTestSuite) TestGetByID_Success() {
	var (
		ctx    = context.Background()
		params = arguments.UserGetByIDArgs{
			ID: 1,
		}
		user models.User
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryRowContext", ctx, mock.Anything).Return(s.MockIRow)
	s.MockIRow.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(nil)
	s.MockIUser.On("GetByID", ctx, params).Return(user, nil)
	actual, err := s.Repository.GetByID(ctx, params)
	s.Nil(err)
	s.Equal(user, actual)
}

func (s *UserRepositoryTestSuite) TestGetByID_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.UserGetByIDArgs{
			ID: 1,
		}
		user models.User
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryRowContext", ctx, mock.Anything).Return(s.MockIRow)
	s.MockIRow.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("some errors"))
	s.MockIUser.On("GetByID", ctx, params).Return(user, errors.New("some errors"))
	actual, err := s.Repository.GetByID(ctx, params)
	s.Equal(user, actual)
	s.NotNil(err)
}

func (s *UserRepositoryTestSuite) TestGetByIDs_Success() {
	var (
		ctx    = context.Background()
		params = arguments.UserGetByIDsArgs{
			IDs: []int64{1, 2},
		}
		users []models.User
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryContext", ctx, mock.Anything, mock.Anything).Return(s.MockIRows, nil)
	s.MockIRows.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockIUser.On("GetByIDs", ctx, params).Return(users, nil)
	actual, err := s.Repository.GetByIDs(ctx, params)
	s.Nil(err)
	s.Equal(users, actual)
}

func (s *UserRepositoryTestSuite) TestGetByIDs_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.UserGetByIDsArgs{
			IDs: []int64{1, 2},
		}
		users []models.User
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryContext", ctx, mock.Anything, mock.Anything).Return(s.MockIRows, errors.New("some errors"))
	s.MockIRows.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("some errors"))
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockIUser.On("GetByIDs", ctx, params).Return(users, errors.New("some errors"))
	actual, err := s.Repository.GetByIDs(ctx, params)
	s.Equal(users, actual)
	s.NotNil(err)
}

func (s *UserRepositoryTestSuite) TestSetArgsToListSelectBuilder_Success() {
	var (
		params        = arguments.UserListArgs{}
		selectBuilder = sq.Select("*").From("user")
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	offset := utils.CalculateOffsetForPage(params.Page, params.PageSize)
	expectedSelectBuilder := selectBuilder.Where(sq.Eq{"id": params.ID}).Where(sq.Eq{"username": params.Username}).Where(sq.Eq{"password": params.Password}).Where(sq.Eq{"name": params.Name}).Where(sq.Eq{"date_of_birth": params.DateOfBirth}).Where(sq.Eq{"reference": params.Reference}).Where(sq.Eq{"avatar_url": params.AvatarUrl}).Where(sq.Eq{"license_number": params.LicenseNumber}).Where(sq.Eq{"phone_number": params.PhoneNumber}).Where(sq.Eq{"extension": params.Extension}).Where(sq.Eq{"tel_provider": params.TelProvider}).Where(sq.Eq{"tel_api": params.TelApi}).Where(sq.Eq{"supervisor_id": params.SupervisorId}).Where(sq.Eq{"role_id": params.RoleId}).Where(sq.Eq{"company_id": params.CompanyID}).Where(sq.Eq{"status": params.Status}).Where(sq.Eq{"created_by": params.CreatedBy}).Where(sq.Eq{"updated_by": params.UpdatedBy}).Limit(uint64(params.PageSize)).Offset(uint64(offset))
	expectSQL, expectArgs, expectErr := expectedSelectBuilder.ToSql()
	// Actual
	actual := s.Repository.setArgsToListSelectBuilder(selectBuilder, params)
	sql, args, err := actual.ToSql()
	s.Nil(err)
	s.Equal(expectSQL, sql)
	s.Equal(expectArgs, args)
	s.Equal(expectErr, err)
}

func (s *UserRepositoryTestSuite) TestList_Success() {
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
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryContext", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(s.MockIRows, nil)
	s.MockIRows.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockIUser.On("List", ctx, params).Return(users, nil)
	actual, err := s.Repository.List(ctx, params)
	s.Nil(err)
	s.Equal(users, actual)
}

func (s *UserRepositoryTestSuite) TestList_Fail() {
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
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryContext", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(s.MockIRows, errors.New("some errors"))
	s.MockIRows.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("some errors"))
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockIUser.On("List", ctx, params).Return(users, errors.New("some errors"))
	actual, err := s.Repository.List(ctx, params)
	s.Equal(users, actual)
	s.NotNil(err)
}

func (s *UserRepositoryTestSuite) TestSetArgsToCountSelectBuilder_Success() {
	var (
		params        = arguments.UserCountArgs{}
		selectBuilder = sq.Select("COUNT(id)").From("user")
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	expectedSelectBuilder := selectBuilder.Where(sq.Eq{"id": params.ID}).Where(sq.Eq{"username": params.Username}).Where(sq.Eq{"password": params.Password}).Where(sq.Eq{"name": params.Name}).Where(sq.Eq{"date_of_birth": params.DateOfBirth}).Where(sq.Eq{"reference": params.Reference}).Where(sq.Eq{"avatar_url": params.AvatarUrl}).Where(sq.Eq{"license_number": params.LicenseNumber}).Where(sq.Eq{"phone_number": params.PhoneNumber}).Where(sq.Eq{"extension": params.Extension}).Where(sq.Eq{"tel_provider": params.TelProvider}).Where(sq.Eq{"tel_api": params.TelApi}).Where(sq.Eq{"supervisor_id": params.SupervisorId}).Where(sq.Eq{"role_id": params.RoleId}).Where(sq.Eq{"company_id": params.CompanyID}).Where(sq.Eq{"status": params.Status}).Where(sq.Eq{"created_by": params.CreatedBy}).Where(sq.Eq{"updated_by": params.UpdatedBy})
	expectSQL, expectArgs, expectErr := expectedSelectBuilder.ToSql()
	// Actual
	actual := s.Repository.setArgsToCountSelectBuilder(selectBuilder, params)
	sql, args, err := actual.ToSql()
	s.Nil(err)
	s.Equal(expectSQL, sql)
	s.Equal(expectArgs, args)
	s.Equal(expectErr, err)
}

func (s *UserRepositoryTestSuite) TestCount_Success() {
	var (
		ctx    = context.Background()
		params = arguments.UserCountArgs{}
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
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(s.MockIRows, nil)
	s.MockIRows.On("Scan", mock.Anything).Return(nil)
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockIUser.On("List", ctx, params).Return(count, nil)
	actual, err := s.Repository.Count(ctx, params)
	s.Nil(err)
	s.Equal(count, actual)
}

func (s *UserRepositoryTestSuite) TestCount_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.UserCountArgs{}
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
	s.MockIUser.On("List", ctx, params).Return(count, errors.New("some errors"))
	actual, err := s.Repository.Count(ctx, params)
	s.Equal(count, actual)
	s.NotNil(err)
}

func (s *UserRepositoryTestSuite) TestInsert_Success() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = arguments.UserInsertArgs{}
		expected models.User
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
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
	//Mock Insert
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx,
		params.Username,
		params.Password,
		params.Name,
		params.DateOfBirth,
		params.Reference,
		params.AvatarUrl,
		params.LicenseNumber,
		params.PhoneNumber,
		params.Extension,
		params.TelProvider,
		params.TelApi,
		params.SupervisorId,
		params.RoleId,
		params.CompanyID,
		params.Status,
		params.CreatedBy,
		params.UpdatedBy,
	).Return(s.MockIResult, nil)
	s.MockIResult.On("LastInsertId").Return(sampleID, nil)
	s.MockIUser.On("Insert", ctx, params).Return(user, nil)
	// Mock GetByID
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryRowContext", ctx, mock.Anything).Return(s.MockIRow)
	s.MockIRow.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(nil)
	s.MockIUser.On("GetByID", ctx, params).Return(user, nil)

	actual, err := s.Repository.Insert(ctx, params)
	s.Nil(err)
	s.Equal(expected, actual)
}

func (s *UserRepositoryTestSuite) TestInsert_Fail() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = arguments.UserInsertArgs{
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
		user models.User
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("ExecContext", ctx,
		params.Username,
		params.Password,
		params.Name,
		params.DateOfBirth,
		params.Reference,
		params.AvatarUrl,
		params.LicenseNumber,
		params.PhoneNumber,
		params.Extension,
		params.TelProvider,
		params.TelApi,
		params.SupervisorId,
		params.RoleId,
		params.CompanyID,
		params.Status,
		params.CreatedBy,
		params.UpdatedBy,
	).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("LastInsertId").Return(sampleID, errors.New("some errors"))
	s.MockIUser.On("Insert", ctx, params).Return(user, errors.New("some errors"))
	actual, err := s.Repository.Insert(ctx, params)
	s.Equal(user, actual)
	s.NotNil(err)
}

func (s *UserRepositoryTestSuite) TestSetArgsToUpdateBuilder_Success() {
	var (
		params = arguments.UserUpdateArgs{}
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	updateBuilder := sq.Update("user").Where(sq.Eq{"id": *params.ID})
	expectedSelectBuilder := updateBuilder.Set("username", *params.Username).Set("password", *params.Password).Set("name", *params.Name).Set("date_of_birth", *params.DateOfBirth).Set("reference", *params.Reference).Set("avatar_url", *params.AvatarUrl).Set("license_number", *params.LicenseNumber).Set("phone_number", *params.PhoneNumber).Set("extension", *params.Extension).Set("tel_provider", *params.TelProvider).Set("tel_api", *params.TelApi).Set("supervisor_id", *params.SupervisorId).Set("role_id", *params.RoleId).Set("company_id", *params.CompanyID).Set("status", *params.Status).Set("created_by", *params.CreatedBy).Set("updated_by", *params.UpdatedBy)
	actual := s.Repository.setArgsToUpdateBuilder(updateBuilder, params)
	s.Equal(expectedSelectBuilder, actual)
}

func (s *UserRepositoryTestSuite) TestUpdate_Success() {
	var (
		ctx      = context.Background()
		params   = arguments.UserUpdateArgs{}
		expected models.User
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
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
	// Mock Update
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx,
		*params.Username,
		*params.Password,
		*params.Name,
		*params.DateOfBirth,
		*params.Reference,
		*params.AvatarUrl,
		*params.LicenseNumber,
		*params.PhoneNumber,
		*params.Extension,
		*params.TelProvider,
		*params.TelApi,
		*params.SupervisorId,
		*params.RoleId,
		*params.CompanyID,
		*params.Status,
		*params.CreatedBy,
		*params.UpdatedBy,
		*params.ID,
	).Return(s.MockIResult, nil)
	s.MockIResult.On("RowsAffected").Return(*params.ID, nil)
	s.MockIUser.On("Update", ctx, params).Return(user, nil)
	// Mock GetByID
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryRowContext", ctx, mock.Anything).Return(s.MockIRow)
	s.MockIRow.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(nil)
	s.MockIUser.On("GetByID", ctx, params).Return(user, nil)

	actual, err := s.Repository.Update(ctx, params)
	s.Nil(err)
	s.Equal(expected, actual)
}

func (s *UserRepositoryTestSuite) TestUpdate_Fail() {
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
		user models.User
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("ExecContext", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("RowsAffected").Return(sampleID, errors.New("some errors"))
	s.MockIUser.On("Update", ctx, params).Return(user, errors.New("some errors"))
	actual, err := s.Repository.Update(ctx, params)
	s.Equal(user, actual)
	s.NotNil(err)
}

func (s *UserRepositoryTestSuite) TestDelete_Success() {
	var (
		ctx               = context.Background()
		params            = arguments.UserDeleteArgs{}
		rowEffected int64 = 1
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx, params.ID).Return(s.MockIResult, nil)
	s.MockIResult.On("RowsAffected").Return(params.ID, nil)
	s.MockIUser.On("Delete", ctx, params).Return(rowEffected, nil)
	actual, err := s.Repository.Delete(ctx, params)
	s.Nil(err)
	s.Equal(params.ID, actual)
}

func (s *UserRepositoryTestSuite) TestDelete_Fail() {
	var (
		ctx         = context.Background()
		params      = arguments.UserDeleteArgs{}
		rowEffected int64
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("ExecContext", ctx, params.ID).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("RowsAffected").Return(params.ID, errors.New("some errors"))
	s.MockIUser.On("Delete", ctx, params).Return(rowEffected, errors.New("some errors"))
	actual, err := s.Repository.Delete(ctx, params)
	s.Equal(rowEffected, actual)
	s.NotNil(err)
}
