package role

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tanphamhaiduong/go/delta/internal/modules/role/mocks"
)

type RoleResolverTestSuite struct {
	suite.Suite
	MockIRole *mocks.IHandler
	Role      ResolverImpl
}

func TestRoleResolverTestSuite(t *testing.T) {
	suite.Run(t, new(RoleResolverTestSuite))
}

func (s *RoleResolverTestSuite) SetupTest() {
	s.MockIRole = &mocks.IHandler{}
	resolver := NewResolver(s.MockIRole)
	s.Role = *resolver
	s.Role.role = s.MockIRole
}
