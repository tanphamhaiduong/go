package user

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tanphamhaiduong/go/delta/internal/modules/user/mocks"
)

type UserResolverTestSuite struct {
	suite.Suite
	MockIUser *mocks.IHandler
	User      ResolverImpl
}

func TestUserResolverTestSuite(t *testing.T) {
	suite.Run(t, new(UserResolverTestSuite))
}

func (s *UserResolverTestSuite) SetupTest() {
	s.MockIUser = &mocks.IHandler{}
	resolver := NewResolver(s.MockIUser)
	s.User = *resolver
	s.User.user = s.MockIUser
}
