package rolepermission

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/tanphamhaiduong/go/common/logger"
	mocksDB "github.com/tanphamhaiduong/go/delta/internal/database/mocks"
	"github.com/tanphamhaiduong/go/delta/internal/models"
	"github.com/tanphamhaiduong/go/delta/internal/modules/rolepermission/mocks"
)

type RolePermissionRepositoryTestSuite struct {
	suite.Suite
	MockIDB             *mocksDB.IDB
	MockIStmt           *mocksDB.IStmt
	MockIRow            *mocksDB.IRow
	MockIRows           *mocksDB.IRows
	MockIResult         *mocksDB.IResult
	MockIRolePermission *mocks.IRepository
	Repository          RepositoryImpl
}

func TestRolePermissionRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RolePermissionRepositoryTestSuite))
}

func (s *RolePermissionRepositoryTestSuite) SetupTest() {
	mockIDB := &mocksDB.IDB{}
	s.MockIDB = mockIDB
	s.MockIStmt = &mocksDB.IStmt{}
	s.MockIRow = &mocksDB.IRow{}
	s.MockIRows = &mocksDB.IRows{}
	s.MockIResult = &mocksDB.IResult{}
	s.MockIRolePermission = &mocks.IRepository{}
	repository := NewRepository(mockIDB)
	s.Repository = *repository
	logConfig := logger.Configuration{}
	logger.NewLogger(logConfig, logger.InstanceZapLogger)
}

func (s *RolePermissionRepositoryTestSuite) TestGetByRoleID_Success() {
	var (
		ctx                   = context.Background()
		roleID          int64 = 1
		rolepermissions []models.RolePermission
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryContext", ctx, mock.Anything, mock.Anything).Return(s.MockIRows, nil)
	s.MockIRows.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(nil)
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	actual, err := s.Repository.GetByRoleID(ctx, roleID)
	s.Nil(err)
	s.Equal(rolepermissions, actual)
}

func (s *RolePermissionRepositoryTestSuite) TestGetByRoleID_Fail() {
	var (
		ctx                   = context.Background()
		roleID          int64 = 1
		rolepermissions []models.RolePermission
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryContext", ctx, mock.Anything, mock.Anything).Return(s.MockIRows, errors.New("some errors"))
	s.MockIRows.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(errors.New("some errors"))
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	actual, err := s.Repository.GetByRoleID(ctx, roleID)
	s.Equal(rolepermissions, actual)
	s.NotNil(err)
}

func (s *RolePermissionRepositoryTestSuite) TestGetByRoleID_Fail1() {
	var (
		ctx                   = context.Background()
		roleID          int64 = 1
		rolepermissions []models.RolePermission
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryContext", ctx, mock.Anything, mock.Anything).Return(s.MockIRows, errors.New("some errors"))
	s.MockIRows.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(errors.New("some errors"))
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	actual, err := s.Repository.GetByRoleID(ctx, roleID)
	s.Equal(rolepermissions, actual)
	s.NotNil(err)
}

func (s *RolePermissionRepositoryTestSuite) TestGetByRoleID_Fail2() {
	var (
		ctx                   = context.Background()
		roleID          int64 = 1
		rolepermissions []models.RolePermission
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryContext", ctx, mock.Anything, mock.Anything).Return(s.MockIRows, nil)
	s.MockIRows.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(errors.New("some errors"))
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(true)
	actual, err := s.Repository.GetByRoleID(ctx, roleID)
	s.Equal(rolepermissions, actual)
	s.NotNil(err)
}
