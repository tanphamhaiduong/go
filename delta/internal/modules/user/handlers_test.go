package user

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tanphamhaiduong/go/delta/internal/modules/user/mocks"
)

type UserHandlerTestSuite struct {
	suite.Suite
	MockIUser *mocks.IRepository
	User      HandlerImpl
}

func TestUserHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}

func (s *UserHandlerTestSuite) SetupTest() {
	s.MockIUser = &mocks.IRepository{}
	handler := NewHandler(s.MockIUser)
	s.User = *handler
	s.User.user = s.MockIUser
}
