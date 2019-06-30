package permission

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tanphamhaiduong/go/delta/internal/modules/permission/mocks"
)

type PermissionResolverTestSuite struct {
	suite.Suite
	MockIPermission *mocks.IHandler
	Permission      ResolverImpl
}

func TestPermissionResolverTestSuite(t *testing.T) {
	suite.Run(t, new(PermissionResolverTestSuite))
}

func (s *PermissionResolverTestSuite) SetupTest() {
	s.MockIPermission = &mocks.IHandler{}
	resolver := NewResolver(s.MockIPermission)
	s.Permission = *resolver
	s.Permission.permission = s.MockIPermission
}
