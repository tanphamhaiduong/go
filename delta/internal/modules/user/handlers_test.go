package user

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tanphamhaiduong/go/common/logger"
	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
	permissionMocks "github.com/tanphamhaiduong/go/delta/internal/modules/permission/mocks"
	rolepermissionMocks "github.com/tanphamhaiduong/go/delta/internal/modules/rolepermission/mocks"
	userMocks "github.com/tanphamhaiduong/go/delta/internal/modules/user/mocks"
)

type UserHandlerTestSuite struct {
	suite.Suite
	MockIUser           *userMocks.IRepository
	MockIPermission     *permissionMocks.IHandler
	MockIRolePermission *rolepermissionMocks.IHandler
	User                HandlerImpl
}

func TestUserHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}

func (s *UserHandlerTestSuite) SetupTest() {
	s.MockIUser = &userMocks.IRepository{}
	s.MockIPermission = &permissionMocks.IHandler{}
	s.MockIRolePermission = &rolepermissionMocks.IHandler{}
	handler := NewHandler(s.MockIUser, s.MockIPermission, s.MockIRolePermission)
	s.User = *handler
	s.User.user = s.MockIUser
	s.User.permission = s.MockIPermission
	s.User.rolePermission = s.MockIRolePermission
	logConfig := logger.Configuration{}
	logger.NewLogger(logConfig, logger.InstanceZapLogger)
}

func (s *UserHandlerTestSuite) TestLogin_Success() {
	var (
		ctx   = context.Background()
		param = arguments.UserLogin{
			Username: "developer",
			Password: "1",
		}
		user = models.User{
			Username: "developer",
			Password: "$2a$10$gK6vF.VHzdUqnsviHCcCI.UYyLLv9n7fqUe.kQz9ZxBJM0SNQ21ES",
			RoleID:   1,
		}
		permissionIDs = []int64{1}
		permission    = models.Permission{
			ID:   1,
			Name: "developer",
		}
		rolePermission = models.RolePermission{
			ID:           1,
			RoleID:       user.RoleID,
			PermissionID: permission.ID,
		}
		permissions = []models.Permission{
			permission,
		}
		rolePermissions = []models.RolePermission{
			rolePermission,
		}
	)
	s.MockIUser.On("GetByUsername", ctx, param.Username).Return(user, nil)
	s.MockIRolePermission.On("GetByRoleID", ctx, user.RoleID).Return(rolePermissions, nil)
	s.MockIPermission.On("GetByIDs", ctx, arguments.PermissionGetByIDs{IDs: permissionIDs}).Return(permissions, nil)
	actual, err := s.User.Login(ctx, param)
	s.Nil(err)
	s.NotNil(actual)
}

func (s *UserHandlerTestSuite) TestLogin_Fail() {
	var (
		ctx   = context.Background()
		param = arguments.UserLogin{
			Username: "developer",
			Password: "1",
		}
		user = models.User{
			Username: "developer",
			Password: "$2a$10$gK6vF.VHzdUqnsviHCcCI.UYyLLv9n7fqUe.kQz9ZxBJM0SNQ21ES",
			RoleID:   1,
		}

		permissionIDs = []int64{1}
		expected      = ""
		permission    = models.Permission{
			ID:   1,
			Name: "developer",
		}
		rolePermission = models.RolePermission{
			ID:           1,
			RoleID:       user.RoleID,
			PermissionID: permission.ID,
		}
		permissions = []models.Permission{
			permission,
		}
		rolePermissions = []models.RolePermission{
			rolePermission,
		}
	)
	s.MockIUser.On("GetByUsername", ctx, param.Username).Return(user, errors.New("some errors"))
	s.MockIRolePermission.On("GetByRoleID", ctx, user.RoleID).Return(rolePermissions, errors.New("some errors"))
	s.MockIPermission.On("GetByIDs", ctx, arguments.PermissionGetByIDs{IDs: permissionIDs}).Return(permissions, errors.New("some errors"))
	actual, err := s.User.Login(ctx, param)
	s.Equal(expected, actual)
	s.NotNil(err)
}

func (s *UserHandlerTestSuite) TestLogin_Fail1() {
	var (
		ctx   = context.Background()
		param = arguments.UserLogin{
			Username: "developer",
			Password: "1",
		}
		user = models.User{
			Username: "developer",
			Password: "$2a$10$gK6vF.VHzdUqnsviHCcCI.UYyLLv9n7fqUe.kQz9ZxBJM0SNQ21ES",
			RoleID:   1,
		}
		permissionIDs = []int64{1}
		expected      = ""
		permission    = models.Permission{
			ID:   1,
			Name: "developer",
		}
		rolePermission = models.RolePermission{
			ID:           1,
			RoleID:       user.RoleID,
			PermissionID: permission.ID,
		}
		permissions = []models.Permission{
			permission,
		}
		rolePermissions = []models.RolePermission{
			rolePermission,
		}
	)
	s.MockIUser.On("GetByUsername", ctx, param.Username).Return(user, nil)
	s.MockIRolePermission.On("GetByRoleID", ctx, user.RoleID).Return(rolePermissions, errors.New("some errors"))
	s.MockIPermission.On("GetByIDs", ctx, arguments.PermissionGetByIDs{IDs: permissionIDs}).Return(permissions, errors.New("some errors"))
	actual, err := s.User.Login(ctx, param)
	s.Equal(expected, actual)
	s.NotNil(err)
}

func (s *UserHandlerTestSuite) TestLogin_Fail2() {
	var (
		ctx   = context.Background()
		param = arguments.UserLogin{
			Username: "developer",
			Password: "1",
		}
		user = models.User{
			Username: "developer",
			Password: "$2a$10$gK6vF.VHzdUqnsviHCcCI.UYyLLv9n7fqUe.kQz9ZxBJM0SNQ21ES",
			RoleID:   1,
		}
		permissionIDs = []int64{1}
		expected      = ""
		permission    = models.Permission{
			ID:   1,
			Name: "developer",
		}
		rolePermission = models.RolePermission{
			ID:           1,
			RoleID:       user.RoleID,
			PermissionID: permission.ID,
		}
		permissions = []models.Permission{
			permission,
		}
		rolePermissions = []models.RolePermission{
			rolePermission,
		}
	)
	s.MockIUser.On("GetByUsername", ctx, param.Username).Return(user, nil)
	s.MockIRolePermission.On("GetByRoleID", ctx, user.RoleID).Return(rolePermissions, nil)
	s.MockIPermission.On("GetByIDs", ctx, arguments.PermissionGetByIDs{IDs: permissionIDs}).Return(permissions, errors.New("some errors"))
	actual, err := s.User.Login(ctx, param)
	s.Equal(expected, actual)
	s.NotNil(err)
}

func (s *UserHandlerTestSuite) TestLogin_Fail3() {
	var (
		ctx   = context.Background()
		param = arguments.UserLogin{
			Username: "developer",
			Password: "2",
		}
		user = models.User{
			Username: "developer",
			Password: "123",
			RoleID:   1,
		}
		permissionIDs = []int64{1}
		expected      = ""
		permission    = models.Permission{
			ID:   1,
			Name: "developer",
		}
		rolePermission = models.RolePermission{
			ID:           1,
			RoleID:       user.RoleID,
			PermissionID: permission.ID,
		}
		permissions = []models.Permission{
			permission,
		}
		rolePermissions = []models.RolePermission{
			rolePermission,
		}
	)
	s.MockIUser.On("GetByUsername", ctx, param.Username).Return(user, nil)
	s.MockIRolePermission.On("GetByRoleID", ctx, user.RoleID).Return(rolePermissions, nil)
	s.MockIPermission.On("GetByIDs", ctx, arguments.PermissionGetByIDs{IDs: permissionIDs}).Return(permissions, errors.New("some errors"))
	actual, err := s.User.Login(ctx, param)
	s.Equal(expected, actual)
	s.NotNil(err)
}
