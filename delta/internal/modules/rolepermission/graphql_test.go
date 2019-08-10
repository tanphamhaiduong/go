package rolepermission

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tanphamhaiduong/go/common/logger"
	"github.com/tanphamhaiduong/go/delta/internal/modules/rolepermission/mocks"
)

type RolePermissionResolverTestSuite struct {
	suite.Suite
	MockIRolePermission *mocks.IHandler
	RolePermission      ResolverImpl
}

func TestRolePermissionResolverTestSuite(t *testing.T) {
	suite.Run(t, new(RolePermissionResolverTestSuite))
}

func (s *RolePermissionResolverTestSuite) SetupTest() {
	s.MockIRolePermission = &mocks.IHandler{}
	resolver := NewResolver(s.MockIRolePermission)
	s.RolePermission = *resolver
	s.RolePermission.rolepermission = s.MockIRolePermission
	logConfig := logger.Configuration{}
	logger.NewLogger(logConfig, logger.InstanceZapLogger)
}
