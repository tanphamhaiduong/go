package user

import (
	"context"
	"errors"
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/stretchr/testify/suite"
	"github.com/tanphamhaiduong/go/common/logger"
	"github.com/tanphamhaiduong/go/delta/internal/arguments"
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
	logConfig := logger.Configuration{}
	logger.NewLogger(logConfig, logger.InstanceZapLogger)
}

func (s *UserResolverTestSuite) TestLogin_Success() {
	var (
		ctx      = context.Background()
		username = "developer"
		token    = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiVXNlcm5hbWUiOiJkZXZlbG9wZXIiLCJQYXNzd29yZCI6IiQyYSQxMCRnSzZ2Ri5WSHpkVXFuc3ZpSENjQ0kuVVl5TEx2OW43ZnFVZS5rUXo5WnhCSk0wU05RMjFFUyIsIk5hbWUiOiJEZXZlbG9wZXIiLCJEYXRlT2ZCaXJ0aCI6eyJUaW1lIjoiMTk5MS0wOC0wNVQwMDowMDowMFoiLCJWYWxpZCI6dHJ1ZX0sIlJlZmVyZW5jZSI6IjAiLCJBdmF0YXJVcmwiOiJodHRwczovL3Jlcy5jbG91ZGluYXJ5LmNvbS9kdW9uZ3BoYW0tY3JtL2ltYWdlL3VwbG9hZC92MTUxODI3MTg1OC9lNmR1enJ0Yjd1ZGZ3c2o5cW43Ni5qcGciLCJMaWNlbnNlTnVtYmVyIjoiIiwiUGhvbmVOdW1iZXIiOiIiLCJFeHRlbnNpb24iOiIyMDEiLCJUZWxQcm92aWRlciI6Im5vbmUiLCJUZWxBcGkiOiIiLCJTdXBlcnZpc29ySWQiOjAsIlJvbGVJZCI6MSwiQ29tcGFueUlEIjoxLCJTdGF0dXMiOiJhY3RpdmUiLCJDcmVhdGVkQnkiOiIiLCJVcGRhdGVkQnkiOiIiLCJleHAiOjE1NjU0NTMxNzEsIlBlcm1pc3Npb24iOnsiSUQiOjAsIk5hbWUiOiIiLCJEZXNjcmlwdGlvbiI6IiIsIlN0YXR1cyI6IiIsIkNyZWF0ZWRCeSI6IiIsIlVwZGF0ZWRCeSI6IiJ9fQ.A-JzWktCxncWbRg3WRhxh1nJLAikWL9ux4ZU0LVMbz0"
		password = "1"
		params   = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{"username": username, "password": password}}
		args     = arguments.UserLogin{
			Username: username,
			Password: password,
		}
	)
	s.MockIUser.On("Login", ctx, args).Return(token, nil)
	actual, err := s.User.Login(params)
	s.Nil(err)
	s.Equal(token, actual)
}

func (s *UserResolverTestSuite) TestLogin_Fail() {
	var (
		ctx      = context.Background()
		username = "developer"
		token    = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiVXNlcm5hbWUiOiJkZXZlbG9wZXIiLCJQYXNzd29yZCI6IiQyYSQxMCRnSzZ2Ri5WSHpkVXFuc3ZpSENjQ0kuVVl5TEx2OW43ZnFVZS5rUXo5WnhCSk0wU05RMjFFUyIsIk5hbWUiOiJEZXZlbG9wZXIiLCJEYXRlT2ZCaXJ0aCI6eyJUaW1lIjoiMTk5MS0wOC0wNVQwMDowMDowMFoiLCJWYWxpZCI6dHJ1ZX0sIlJlZmVyZW5jZSI6IjAiLCJBdmF0YXJVcmwiOiJodHRwczovL3Jlcy5jbG91ZGluYXJ5LmNvbS9kdW9uZ3BoYW0tY3JtL2ltYWdlL3VwbG9hZC92MTUxODI3MTg1OC9lNmR1enJ0Yjd1ZGZ3c2o5cW43Ni5qcGciLCJMaWNlbnNlTnVtYmVyIjoiIiwiUGhvbmVOdW1iZXIiOiIiLCJFeHRlbnNpb24iOiIyMDEiLCJUZWxQcm92aWRlciI6Im5vbmUiLCJUZWxBcGkiOiIiLCJTdXBlcnZpc29ySWQiOjAsIlJvbGVJZCI6MSwiQ29tcGFueUlEIjoxLCJTdGF0dXMiOiJhY3RpdmUiLCJDcmVhdGVkQnkiOiIiLCJVcGRhdGVkQnkiOiIiLCJleHAiOjE1NjU0NTMxNzEsIlBlcm1pc3Npb24iOnsiSUQiOjAsIk5hbWUiOiIiLCJEZXNjcmlwdGlvbiI6IiIsIlN0YXR1cyI6IiIsIkNyZWF0ZWRCeSI6IiIsIlVwZGF0ZWRCeSI6IiJ9fQ.A-JzWktCxncWbRg3WRhxh1nJLAikWL9ux4ZU0LVMbz0"
		password = "1"
		params   = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{"username": username, "password": password}}
		args     = arguments.UserLogin{
			Username: username,
			Password: password,
		}
	)
	s.MockIUser.On("Login", ctx, args).Return(token, errors.New("some error"))
	actual, err := s.User.Login(params)
	s.Nil(actual)
	s.NotNil(err)
}
