package rolepermission

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tanphamhaiduong/go/common/logger"
	"github.com/tanphamhaiduong/go/delta/internal/models"
	"github.com/tanphamhaiduong/go/delta/internal/modules/rolepermission/mocks"
)

type RolePermissionHandlerTestSuite struct {
	suite.Suite
	MockIRolePermission *mocks.IRepository
	RolePermission      HandlerImpl
}

func TestRolePermissionHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(RolePermissionHandlerTestSuite))
}

func (s *RolePermissionHandlerTestSuite) SetupTest() {
	s.MockIRolePermission = &mocks.IRepository{}
	handler := NewHandler(s.MockIRolePermission)
	s.RolePermission = *handler
	s.RolePermission.rolepermission = s.MockIRolePermission
	logConfig := logger.Configuration{}
	logger.NewLogger(logConfig, logger.InstanceZapLogger)
}

func (s *RolePermissionHandlerTestSuite) TestGetByRoleID_Success() {
	var (
		ctx                   = context.Background()
		roleID          int64 = 1
		rolepermissions []models.RolePermission
	)
	s.MockIRolePermission.On("GetByRoleID", ctx, roleID).Return(rolepermissions, nil)
	actual, err := s.RolePermission.GetByRoleID(ctx, roleID)
	s.Nil(err)
	s.Equal(rolepermissions, actual)
}

func (s *RolePermissionHandlerTestSuite) TestGetByRoleID_Fail() {
	var (
		ctx                   = context.Background()
		roleID          int64 = 1
		rolepermissions []models.RolePermission
	)
	s.MockIRolePermission.On("GetByRoleID", ctx, roleID).Return(rolepermissions, errors.New("some errors"))
	actual, err := s.RolePermission.GetByRoleID(ctx, roleID)
	s.Equal(rolepermissions, actual)
	s.NotNil(err)
}

func (s *RolePermissionHandlerTestSuite) TestGetByRoleID_Fail1() {
	var (
		ctx                   = context.Background()
		roleID          int64 = 1
		rolepermissions []models.RolePermission
	)
	s.MockIRolePermission.On("GetByRoleID", ctx, roleID).Return(rolepermissions, errors.New("some errors"))
	actual, err := s.RolePermission.GetByRoleID(ctx, roleID)
	s.Equal(rolepermissions, actual)
	s.NotNil(err)
}
