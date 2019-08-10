// @generated
package user

import (
	"context"
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/mock"
	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
	"github.com/tanphamhaiduong/go/delta/internal/utils"
)

func (s *UserRepositoryTestSuite) TestGetByID_Success() {
	var (
		ctx   = context.Background()
		param = arguments.UserGetByIDArgs{
			ID: 1,
		}
		user models.User
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryRowContext", ctx, param.ID).Return(s.MockIRow)
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
	actual, err := s.Repository.GetByID(ctx, param)
	s.Nil(err)
	s.Equal(user, actual)
}

func (s *UserRepositoryTestSuite) TestGetByID_Fail() {
	var (
		ctx   = context.Background()
		param = arguments.UserGetByIDArgs{
			ID: 1,
		}
		user models.User
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryRowContext", ctx, param.ID).Return(s.MockIRow)
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
	).Return(errors.New("some errors"))
	actual, err := s.Repository.GetByID(ctx, param)
	s.Equal(user, actual)
	s.NotNil(err)
}

func (s *UserRepositoryTestSuite) TestGetByID_Fail1() {
	var (
		ctx   = context.Background()
		param = arguments.UserGetByIDArgs{
			ID: 1,
		}
		user models.User
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryRowContext", ctx, param.ID).Return(s.MockIRow)
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
	).Return(errors.New("some errors"))
	actual, err := s.Repository.GetByID(ctx, param)
	s.Equal(user, actual)
	s.NotNil(err)
}

func (s *UserRepositoryTestSuite) TestGetByIDs_Success() {
	var (
		ctx   = context.Background()
		param = arguments.UserGetByIDsArgs{
			IDs: []int64{1, 2},
		}
		users []models.User
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryContext", ctx, mock.Anything, mock.Anything).Return(s.MockIRows, nil)
	s.MockIRows.On("Scan",
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
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	actual, err := s.Repository.GetByIDs(ctx, param)
	s.Nil(err)
	s.Equal(users, actual)
}

func (s *UserRepositoryTestSuite) TestGetByIDs_Fail() {
	var (
		ctx   = context.Background()
		param = arguments.UserGetByIDsArgs{
			IDs: []int64{1, 2},
		}
		users []models.User
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryContext", ctx, mock.Anything, mock.Anything).Return(s.MockIRows, errors.New("some errors"))
	s.MockIRows.On("Scan",
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
	).Return(errors.New("some errors"))
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	actual, err := s.Repository.GetByIDs(ctx, param)
	s.Equal(users, actual)
	s.NotNil(err)
}

func (s *UserRepositoryTestSuite) TestGetByIDs_Fail1() {
	var (
		ctx   = context.Background()
		param = arguments.UserGetByIDsArgs{
			IDs: []int64{1, 2},
		}
		users []models.User
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryContext", ctx, mock.Anything, mock.Anything).Return(s.MockIRows, errors.New("some errors"))
	s.MockIRows.On("Scan",
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
	).Return(errors.New("some errors"))
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	actual, err := s.Repository.GetByIDs(ctx, param)
	s.Equal(users, actual)
	s.NotNil(err)
}

func (s *UserRepositoryTestSuite) TestGetByIDs_Fail2() {
	var (
		ctx   = context.Background()
		param = arguments.UserGetByIDsArgs{
			IDs: []int64{1, 2},
		}
		users []models.User
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryContext", ctx, mock.Anything, mock.Anything).Return(s.MockIRows, nil)
	s.MockIRows.On("Scan",
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
	).Return(errors.New("some errors"))
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(true)
	actual, err := s.Repository.GetByIDs(ctx, param)
	s.Equal(users, actual)
	s.NotNil(err)
}

func (s *UserRepositoryTestSuite) TestSetArgsToListSelectBuilder_Success() {
	var (
		ctx    = context.Background()
		params = arguments.UserListArgs{
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
			SupervisorId:  1,
			RoleId:        1,
			CompanyID:     1,
			Status:        "active",
			CreatedBy:     "mockString",
			UpdatedBy:     "mockString",
			Page:          1,
			PageSize:      10,
		}
		selectBuilder = sq.Select("*").From("user")
	)
	offset := utils.CalculateOffsetForPage(params.Page, params.PageSize)
	expectedSelectBuilder := selectBuilder.Where(sq.Eq{"id": params.ID}).Where(sq.Eq{"username": params.Username}).Where(sq.Eq{"password": params.Password}).Where(sq.Eq{"name": params.Name}).Where(sq.Eq{"date_of_birth": params.DateOfBirth}).Where(sq.Eq{"reference": params.Reference}).Where(sq.Eq{"avatar_url": params.AvatarUrl}).Where(sq.Eq{"license_number": params.LicenseNumber}).Where(sq.Eq{"phone_number": params.PhoneNumber}).Where(sq.Eq{"extension": params.Extension}).Where(sq.Eq{"tel_provider": params.TelProvider}).Where(sq.Eq{"tel_api": params.TelApi}).Where(sq.Eq{"supervisor_id": params.SupervisorId}).Where(sq.Eq{"role_id": params.RoleId}).Where(sq.Eq{"company_id": params.CompanyID}).Where(sq.Eq{"status": params.Status}).Where(sq.Eq{"created_by": params.CreatedBy}).Where(sq.Eq{"updated_by": params.UpdatedBy}).Limit(uint64(params.PageSize)).Offset(uint64(offset))
	expectSQL, expectArgs, expectErr := expectedSelectBuilder.ToSql()
	// Actual
	actual := s.Repository.setArgsToListSelectBuilder(ctx, selectBuilder, params)
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
			SupervisorId:  1,
			RoleId:        1,
			CompanyID:     1,
			Status:        "active",
			CreatedBy:     "mockString",
			UpdatedBy:     "mockString",
			Page:          1,
			PageSize:      10,
		}
		users []models.User
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryContext", ctx,
		params.ID,
		params.Username,
		params.Password,
		params.Name,
		params.DateOfBirth.Time,
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
		mock.Anything,
		mock.Anything,
	).Return(s.MockIRows, nil)
	s.MockIRows.On("Scan",
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
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	actual, err := s.Repository.List(ctx, params)
	s.Nil(err)
	s.Equal(users, actual)
}

func (s *UserRepositoryTestSuite) TestList_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.UserListArgs{
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
			SupervisorId:  1,
			RoleId:        1,
			CompanyID:     1,
			Status:        "active",
			CreatedBy:     "mockString",
			UpdatedBy:     "mockString",
			Page:          1,
			PageSize:      10,
		}
		users []models.User
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryContext", ctx,
		params.ID,
		params.Username,
		params.Password,
		params.Name,
		params.DateOfBirth.Time,
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
		mock.Anything,
		mock.Anything,
	).Return(s.MockIRows, errors.New("some errors"))
	s.MockIRows.On("Scan",
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
	).Return(errors.New("some errors"))
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	actual, err := s.Repository.List(ctx, params)
	s.Equal(users, actual)
	s.NotNil(err)
}

func (s *UserRepositoryTestSuite) TestList_Fail1() {
	var (
		ctx    = context.Background()
		params = arguments.UserListArgs{
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
			SupervisorId:  1,
			RoleId:        1,
			CompanyID:     1,
			Status:        "active",
			CreatedBy:     "mockString",
			UpdatedBy:     "mockString",
			Page:          1,
			PageSize:      10,
		}
		users []models.User
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryContext", ctx,
		params.ID,
		params.Username,
		params.Password,
		params.Name,
		params.DateOfBirth.Time,
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
		mock.Anything,
		mock.Anything,
	).Return(s.MockIRows, errors.New("some errors"))
	s.MockIRows.On("Scan",
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
	).Return(errors.New("some errors"))
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	actual, err := s.Repository.List(ctx, params)
	s.Equal(users, actual)
	s.NotNil(err)
}

func (s *UserRepositoryTestSuite) TestList_Fail2() {
	var (
		ctx    = context.Background()
		params = arguments.UserListArgs{
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
			SupervisorId:  1,
			RoleId:        1,
			CompanyID:     1,
			Status:        "active",
			CreatedBy:     "mockString",
			UpdatedBy:     "mockString",
			Page:          1,
			PageSize:      10,
		}
		users []models.User
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryContext", ctx,
		params.ID,
		params.Username,
		params.Password,
		params.Name,
		params.DateOfBirth.Time,
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
		mock.Anything,
		mock.Anything,
	).Return(s.MockIRows, nil)
	s.MockIRows.On("Scan",
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
	).Return(errors.New("some errors"))
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(true)
	actual, err := s.Repository.List(ctx, params)
	s.Equal(users, actual)
	s.NotNil(err)
}

func (s *UserRepositoryTestSuite) TestSetArgsToCountSelectBuilder_Success() {
	var (
		ctx    = context.Background()
		params = arguments.UserCountArgs{
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
			SupervisorId:  1,
			RoleId:        1,
			CompanyID:     1,
			Status:        "active",
			CreatedBy:     "mockString",
			UpdatedBy:     "mockString",
		}
		selectBuilder = sq.Select("COUNT(id)").From("user")
	)
	expectedSelectBuilder := selectBuilder.Where(sq.Eq{"id": params.ID}).Where(sq.Eq{"username": params.Username}).Where(sq.Eq{"password": params.Password}).Where(sq.Eq{"name": params.Name}).Where(sq.Eq{"date_of_birth": params.DateOfBirth}).Where(sq.Eq{"reference": params.Reference}).Where(sq.Eq{"avatar_url": params.AvatarUrl}).Where(sq.Eq{"license_number": params.LicenseNumber}).Where(sq.Eq{"phone_number": params.PhoneNumber}).Where(sq.Eq{"extension": params.Extension}).Where(sq.Eq{"tel_provider": params.TelProvider}).Where(sq.Eq{"tel_api": params.TelApi}).Where(sq.Eq{"supervisor_id": params.SupervisorId}).Where(sq.Eq{"role_id": params.RoleId}).Where(sq.Eq{"company_id": params.CompanyID}).Where(sq.Eq{"status": params.Status}).Where(sq.Eq{"created_by": params.CreatedBy}).Where(sq.Eq{"updated_by": params.UpdatedBy})
	expectSQL, expectArgs, expectErr := expectedSelectBuilder.ToSql()
	// Actual
	actual := s.Repository.setArgsToCountSelectBuilder(ctx, selectBuilder, params)
	sql, args, err := actual.ToSql()
	s.Nil(err)
	s.Equal(expectSQL, sql)
	s.Equal(expectArgs, args)
	s.Equal(expectErr, err)
}

func (s *UserRepositoryTestSuite) TestCount_Success() {
	var (
		ctx    = context.Background()
		params = arguments.UserCountArgs{
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
			SupervisorId:  1,
			RoleId:        1,
			CompanyID:     1,
			Status:        "active",
			CreatedBy:     "mockString",
			UpdatedBy:     "mockString",
		}
		count int64
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryRowContext", ctx,
		params.ID,
		params.Username,
		params.Password,
		params.Name,
		params.DateOfBirth.Time,
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
	).Return(s.MockIRows, nil)
	s.MockIRows.On("Scan", mock.Anything).Return(nil)
	actual, err := s.Repository.Count(ctx, params)
	s.Nil(err)
	s.Equal(count, actual)
}

func (s *UserRepositoryTestSuite) TestCount_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.UserCountArgs{
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
			SupervisorId:  1,
			RoleId:        1,
			CompanyID:     1,
			Status:        "active",
			CreatedBy:     "mockString",
			UpdatedBy:     "mockString",
		}
		count int64
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryRowContext", ctx,
		params.ID,
		params.Username,
		params.Password,
		params.Name,
		params.DateOfBirth.Time,
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
	).Return(s.MockIRows, errors.New("some errors"))
	s.MockIRows.On("Scan", mock.Anything).Return(errors.New("some errors"))
	actual, err := s.Repository.Count(ctx, params)
	s.Equal(count, actual)
	s.NotNil(err)
}

func (s *UserRepositoryTestSuite) TestCount_Fail1() {
	var (
		ctx    = context.Background()
		params = arguments.UserCountArgs{
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
			SupervisorId:  1,
			RoleId:        1,
			CompanyID:     1,
			Status:        "active",
			CreatedBy:     "mockString",
			UpdatedBy:     "mockString",
		}
		count int64
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryRowContext", ctx,
		params.ID,
		params.Username,
		params.Password,
		params.Name,
		params.DateOfBirth.Time,
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
	).Return(s.MockIRows, errors.New("some errors"))
	s.MockIRows.On("Scan", mock.Anything).Return(errors.New("some errors"))
	actual, err := s.Repository.Count(ctx, params)
	s.Equal(count, actual)
	s.NotNil(err)
}

func (s *UserRepositoryTestSuite) TestCount_Fail2() {
	var (
		ctx    = context.Background()
		params = arguments.UserCountArgs{
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
			SupervisorId:  1,
			RoleId:        1,
			CompanyID:     1,
			Status:        "active",
			CreatedBy:     "mockString",
			UpdatedBy:     "mockString",
		}
		count int64
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryRowContext", ctx,
		params.ID,
		params.Username,
		params.Password,
		params.Name,
		params.DateOfBirth.Time,
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
	).Return(s.MockIRows, nil)
	s.MockIRows.On("Scan", mock.Anything).Return(errors.New("some errors"))
	actual, err := s.Repository.Count(ctx, params)
	s.Equal(count, actual)
	s.NotNil(err)
}

func (s *UserRepositoryTestSuite) TestInsert_Success() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = arguments.UserInsertArgs{
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
			SupervisorId:  1,
			RoleId:        1,
			CompanyID:     1,
			Status:        "active",
			CreatedBy:     "mockString",
			UpdatedBy:     "mockString",
		}
		expected models.User
	)
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

	actual, err := s.Repository.Insert(ctx, params)
	s.Nil(err)
	s.Equal(expected, actual)
}

func (s *UserRepositoryTestSuite) TestInsert_Fail() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = arguments.UserInsertArgs{
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
			SupervisorId:  1,
			RoleId:        1,
			CompanyID:     1,
			Status:        "active",
			CreatedBy:     "mockString",
			UpdatedBy:     "mockString",
		}
		user models.User
	)
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
	actual, err := s.Repository.Insert(ctx, params)
	s.Equal(user, actual)
	s.NotNil(err)
}

func (s *UserRepositoryTestSuite) TestInsert_Fail1() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = arguments.UserInsertArgs{
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
			SupervisorId:  1,
			RoleId:        1,
			CompanyID:     1,
			Status:        "active",
			CreatedBy:     "mockString",
			UpdatedBy:     "mockString",
		}
		user models.User
	)
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
	).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("LastInsertId").Return(sampleID, errors.New("some errors"))
	actual, err := s.Repository.Insert(ctx, params)
	s.Equal(user, actual)
	s.NotNil(err)
}

func (s *UserRepositoryTestSuite) TestInsert_Fail2() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = arguments.UserInsertArgs{
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
			SupervisorId:  1,
			RoleId:        1,
			CompanyID:     1,
			Status:        "active",
			CreatedBy:     "mockString",
			UpdatedBy:     "mockString",
		}
		user models.User
	)
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
	s.MockIResult.On("LastInsertId").Return(sampleID, errors.New("some errors"))
	actual, err := s.Repository.Insert(ctx, params)
	s.Equal(user, actual)
	s.NotNil(err)
}

func (s *UserRepositoryTestSuite) TestInsert_Fail3() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = arguments.UserInsertArgs{
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
			SupervisorId:  1,
			RoleId:        1,
			CompanyID:     1,
			Status:        "active",
			CreatedBy:     "mockString",
			UpdatedBy:     "mockString",
		}
		expected models.User
	)
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
	).Return(errors.New("some errors"))

	actual, err := s.Repository.Insert(ctx, params)
	s.Equal(expected, actual)
	s.NotNil(err)
}

func (s *UserRepositoryTestSuite) TestSetArgsToUpdateBuilder_Success() {
	var (
		ctx              = context.Background()
		sampleID   int64 = 1
		mockString       = "mockString"
		status           = "active"
		params           = arguments.UserUpdateArgs{
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
			SupervisorId:  &sampleID,
			RoleId:        &sampleID,
			CompanyID:     &sampleID,
			Status:        &status,
			CreatedBy:     &mockString,
			UpdatedBy:     &mockString,
		}
	)
	updateBuilder := sq.Update("user").Where(sq.Eq{"id": *params.ID})
	expectedSelectBuilder := updateBuilder.Set("username", *params.Username).Set("password", *params.Password).Set("name", *params.Name).Set("date_of_birth", *params.DateOfBirth).Set("reference", *params.Reference).Set("avatar_url", *params.AvatarUrl).Set("license_number", *params.LicenseNumber).Set("phone_number", *params.PhoneNumber).Set("extension", *params.Extension).Set("tel_provider", *params.TelProvider).Set("tel_api", *params.TelApi).Set("supervisor_id", *params.SupervisorId).Set("role_id", *params.RoleId).Set("company_id", *params.CompanyID).Set("status", *params.Status).Set("created_by", *params.CreatedBy).Set("updated_by", *params.UpdatedBy)
	actual := s.Repository.setArgsToUpdateBuilder(ctx, updateBuilder, params)
	s.Equal(expectedSelectBuilder, actual)
}

func (s *UserRepositoryTestSuite) TestUpdate_Success() {
	var (
		ctx               = context.Background()
		sampleID    int64 = 1
		rowEffected int64 = 1
		mockString        = "mockString"
		status            = "active"
		params            = arguments.UserUpdateArgs{
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
			SupervisorId:  &sampleID,
			RoleId:        &sampleID,
			CompanyID:     &sampleID,
			Status:        &status,
			CreatedBy:     &mockString,
			UpdatedBy:     &mockString,
		}
		expected models.User
	)
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
	s.MockIResult.On("RowsAffected").Return(rowEffected, nil)
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

	actual, err := s.Repository.Update(ctx, params)
	s.Nil(err)
	s.Equal(expected, actual)
}

func (s *UserRepositoryTestSuite) TestUpdate_Fail() {
	var (
		ctx               = context.Background()
		sampleID    int64 = 1
		rowEffected int64
		mockString  = "mockString"
		status      = "active"
		params      = arguments.UserUpdateArgs{
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
			SupervisorId:  &sampleID,
			RoleId:        &sampleID,
			CompanyID:     &sampleID,
			Status:        &status,
			CreatedBy:     &mockString,
			UpdatedBy:     &mockString,
		}
		user models.User
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
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
	).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("RowsAffected").Return(rowEffected, errors.New("some errors"))
	actual, err := s.Repository.Update(ctx, params)
	s.Equal(user, actual)
	s.NotNil(err)
}

func (s *UserRepositoryTestSuite) TestUpdate_Fail1() {
	var (
		ctx               = context.Background()
		sampleID    int64 = 1
		rowEffected int64
		mockString  = "mockString"
		status      = "active"
		params      = arguments.UserUpdateArgs{
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
			SupervisorId:  &sampleID,
			RoleId:        &sampleID,
			CompanyID:     &sampleID,
			Status:        &status,
			CreatedBy:     &mockString,
			UpdatedBy:     &mockString,
		}
		user models.User
	)
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
	).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("RowsAffected").Return(rowEffected, errors.New("some errors"))
	actual, err := s.Repository.Update(ctx, params)
	s.Equal(user, actual)
	s.NotNil(err)
}

func (s *UserRepositoryTestSuite) TestUpdate_Fail2() {
	var (
		ctx               = context.Background()
		sampleID    int64 = 1
		rowEffected int64
		mockString  = "mockString"
		status      = "active"
		params      = arguments.UserUpdateArgs{
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
			SupervisorId:  &sampleID,
			RoleId:        &sampleID,
			CompanyID:     &sampleID,
			Status:        &status,
			CreatedBy:     &mockString,
			UpdatedBy:     &mockString,
		}
		user models.User
	)
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
	s.MockIResult.On("RowsAffected").Return(rowEffected, errors.New("some errors"))
	actual, err := s.Repository.Update(ctx, params)
	s.Equal(user, actual)
	s.NotNil(err)
}

func (s *UserRepositoryTestSuite) TestUpdate_Fail3() {
	var (
		ctx               = context.Background()
		sampleID    int64 = 1
		rowEffected int64
		mockString  = "mockString"
		status      = "active"
		params      = arguments.UserUpdateArgs{
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
			SupervisorId:  &sampleID,
			RoleId:        &sampleID,
			CompanyID:     &sampleID,
			Status:        &status,
			CreatedBy:     &mockString,
			UpdatedBy:     &mockString,
		}
		user models.User
	)
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
	s.MockIResult.On("RowsAffected").Return(rowEffected, nil)
	actual, err := s.Repository.Update(ctx, params)
	s.Equal(user, actual)
	s.NotNil(err)
}

func (s *UserRepositoryTestSuite) TestUpdate_Fail4() {
	var (
		ctx              = context.Background()
		sampleID   int64 = 1
		mockString       = "mockString"
		status           = "active"
		params           = arguments.UserUpdateArgs{
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
			SupervisorId:  &sampleID,
			RoleId:        &sampleID,
			CompanyID:     &sampleID,
			Status:        &status,
			CreatedBy:     &mockString,
			UpdatedBy:     &mockString,
		}
		user models.User
	)
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
	).Return(errors.New("some errors"))

	actual, err := s.Repository.Update(ctx, params)
	s.Equal(user, actual)
	s.NotNil(err)
}

func (s *UserRepositoryTestSuite) TestDelete_Success() {
	var (
		ctx   = context.Background()
		param = arguments.UserDeleteArgs{
			ID: 1,
		}
		rowEffected int64 = 1
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx, param.ID).Return(s.MockIResult, nil)
	s.MockIResult.On("RowsAffected").Return(rowEffected, nil)
	actual, err := s.Repository.Delete(ctx, param)
	s.Nil(err)
	s.Equal(param.ID, actual)
}

func (s *UserRepositoryTestSuite) TestDelete_Fail() {
	var (
		ctx   = context.Background()
		param = arguments.UserDeleteArgs{
			ID: 1,
		}
		rowEffected int64
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("ExecContext", ctx, param.ID).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("RowsAffected").Return(rowEffected, errors.New("some errors"))
	actual, err := s.Repository.Delete(ctx, param)
	s.Equal(rowEffected, actual)
	s.NotNil(err)
}

func (s *UserRepositoryTestSuite) TestDelete_Fail1() {
	var (
		ctx   = context.Background()
		param = arguments.UserDeleteArgs{
			ID: 1,
		}
		rowEffected int64
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx, param.ID).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("RowsAffected").Return(rowEffected, errors.New("some errors"))
	actual, err := s.Repository.Delete(ctx, param)
	s.Equal(rowEffected, actual)
	s.NotNil(err)
}

func (s *UserRepositoryTestSuite) TestDelete_Fail2() {
	var (
		ctx   = context.Background()
		param = arguments.UserDeleteArgs{
			ID: 1,
		}
		rowEffected int64
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx, param.ID).Return(s.MockIResult, nil)
	s.MockIResult.On("RowsAffected").Return(rowEffected, errors.New("some errors"))
	actual, err := s.Repository.Delete(ctx, param)
	s.Equal(rowEffected, actual)
	s.NotNil(err)
}

func (s *UserRepositoryTestSuite) TestDelete_Fail3() {
	var (
		ctx   = context.Background()
		param = arguments.UserDeleteArgs{
			ID: 1,
		}
		rowEffected int64
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx, param.ID).Return(s.MockIResult, nil)
	s.MockIResult.On("RowsAffected").Return(rowEffected, nil)
	actual, err := s.Repository.Delete(ctx, param)
	s.Equal(rowEffected, actual)
	s.NotNil(err)
}
