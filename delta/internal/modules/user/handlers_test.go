package user

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tanphamhaiduong/go/common/logger"
	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
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
		paramsFetchUser = arguments.UserGetByUsername{Username: param.Username}
		user            = models.User{
			Username: "developer",
			Password: "$2a$10$gK6vF.VHzdUqnsviHCcCI.UYyLLv9n7fqUe.kQz9ZxBJM0SNQ21ES",
		}
	)
	s.MockIUser.On("GetByUsername", ctx, paramsFetchUser).Return(user, nil)
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
		paramsFetchUser = arguments.UserGetByUsername{Username: param.Username}
		user            = models.User{
			Username: "developer",
			Password: "$2a$10$gK6vF.VHzdUqnsviHCcCI.UYyLLv9n7fqUe.kQz9ZxBJM0SNQ21ES",
		}
		expected = ""
	)
	s.MockIUser.On("GetByUsername", ctx, paramsFetchUser).Return(user, errors.New("some errors"))
	actual, err := s.User.Login(ctx, param)
	s.Equal(expected, actual)
	s.NotNil(err)
}

func (s *UserHandlerTestSuite) TestLogin_Fail2() {
	var (
		ctx   = context.Background()
		param = arguments.UserLogin{
			Username: "developer",
			Password: "2",
		}
		paramsFetchUser = arguments.UserGetByUsername{Username: param.Username}
		user            = models.User{
			Username: "developer",
			Password: "$2a$10$gK6vF.VHzdUqnsviHCcCI.UYyLLv9n7fqUe.kQz9ZxBJM0SNQ21ES",
		}
		expected = ""
	)
	s.MockIUser.On("GetByUsername", ctx, paramsFetchUser).Return(user, nil)
	actual, err := s.User.Login(ctx, param)
	s.Equal(expected, actual)
	s.NotNil(err)
}
