// @generated
package user

import (
	"context"
	"errors"

	"github.com/graphql-go/graphql"

	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
)

func (s *UserResolverTestSuite) TestForwardParams_Success() {
	var (
		sampleID int64 = 1
		params         = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{"id": sampleID}}
		expected       = map[string]interface{}(map[string]interface{}{"id": sampleID})
	)
	actual, err := s.User.ForwardParams(params)
	s.Nil(err)
	s.Equal(expected, actual)
}

func (s *UserResolverTestSuite) TestGetByID_Success() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{"id": sampleID}}
		user     models.User
		args     = arguments.UserGetByIDArgs{
			ID: 1,
		}
	)
	s.MockIUser.On("GetByID", ctx, args).Return(user, nil)
	actual, err := s.User.GetByID(params)
	s.Nil(err)
	s.Equal(user, actual)
}

func (s *UserResolverTestSuite) TestGetByID_Fail() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{"id": sampleID}}
		user     models.User
		args     = arguments.UserGetByIDArgs{
			ID: 1,
		}
	)
	s.MockIUser.On("GetByID", ctx, args).Return(user, errors.New("some errors"))
	actual, err := s.User.GetByID(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *UserResolverTestSuite) TestList_Success() {
	var (
		ctx    = context.Background()
		params = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{"page": 1, "pageSize": 10}, Args: map[string]interface{}{}}
		users  []models.User
		args   = arguments.UserListArgs{
			Page:     1,
			PageSize: 10,
		}
	)
	s.MockIUser.On("List", ctx, args).Return(users, nil)
	actual, err := s.User.List(params)
	s.Nil(err)
	s.Equal(users, actual)
}

func (s *UserResolverTestSuite) TestList_Fail() {
	var (
		ctx    = context.Background()
		params = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{"page": 1, "pageSize": 10}, Args: map[string]interface{}{}}
		users  []models.User
		args   = arguments.UserListArgs{
			Page:     1,
			PageSize: 10,
		}
	)
	s.MockIUser.On("List", ctx, args).Return(users, errors.New("some errors"))
	actual, err := s.User.List(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *UserResolverTestSuite) TestCount_Success() {
	var (
		ctx    = context.Background()
		params = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args   = arguments.UserCountArgs{}
		count  int64
	)
	s.MockIUser.On("Count", ctx, args).Return(count, nil)
	actual, err := s.User.Count(params)
	s.Nil(err)
	s.Equal(count, actual)
}

func (s *UserResolverTestSuite) TestCount_Fail() {
	var (
		ctx    = context.Background()
		params = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args   = arguments.UserCountArgs{}
		count  int64
	)
	s.MockIUser.On("Count", ctx, args).Return(count, errors.New("some errors"))
	actual, err := s.User.Count(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *UserResolverTestSuite) TestInsert_Success() {
	var (
		ctx    = context.Background()
		params = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args   = arguments.UserInsertArgs{}
		user   models.User
	)
	// Mock Insert
	s.MockIUser.On("Insert", ctx, args).Return(user, nil)
	// Mock GetByID
	s.MockIUser.On("GetByID", ctx, params).Return(user, nil)

	actual, err := s.User.Insert(params)
	s.Nil(err)
	s.Equal(user, actual)
}

func (s *UserResolverTestSuite) TestInsert_Fail() {
	var (
		ctx    = context.Background()
		params = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args   = arguments.UserInsertArgs{}
		user   models.User
	)
	s.MockIUser.On("Insert", ctx, args).Return(user, errors.New("some errors"))
	actual, err := s.User.Insert(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *UserResolverTestSuite) TestUpdate_Success() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = graphql.ResolveParams{
			Context: context.Background(),
			Source:  map[string]interface{}{},
			Args:    map[string]interface{}{"id": 1},
		}
		args = arguments.UserUpdateArgs{
			ID: &sampleID,
		}
		user = models.User{}
	)
	// Mock for Update
	s.MockIUser.On("Update", ctx, args).Return(user, nil)
	// Mock for LoadByID
	s.MockIUser.On("GetByID", ctx, args).Return(user, nil)
	actual, err := s.User.Update(params)
	s.Nil(err)
	s.Equal(user, actual)
}

func (s *UserResolverTestSuite) TestUpdate_Fail() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = graphql.ResolveParams{
			Context: context.Background(),
			Source:  map[string]interface{}{},
			Args:    map[string]interface{}{"id": 1},
		}
		args = arguments.UserUpdateArgs{
			ID: &sampleID,
		}
		user models.User
	)
	s.MockIUser.On("Update", ctx, args).Return(user, errors.New("some errors"))
	actual, err := s.User.Update(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *UserResolverTestSuite) TestDelete_Success() {
	var (
		ctx    = context.Background()
		params = graphql.ResolveParams{
			Context: context.Background(),
			Source:  map[string]interface{}{},
			Args:    map[string]interface{}{"id": 1},
		}
		args = arguments.UserDeleteArgs{
			ID: 1,
		}
		sampleID int64 = 1
	)
	s.MockIUser.On("Delete", ctx, args).Return(sampleID, nil)
	actual, err := s.User.Delete(params)
	s.Nil(err)
	s.Equal(sampleID, actual)
}

func (s *UserResolverTestSuite) TestDelete_Fail() {
	var (
		ctx    = context.Background()
		params = graphql.ResolveParams{
			Context: context.Background(),
			Source:  map[string]interface{}{},
			Args:    map[string]interface{}{"id": 1},
		}
		args = arguments.UserDeleteArgs{
			ID: 1,
		}
		sampleID int64
	)
	s.MockIUser.On("Delete", ctx, args).Return(sampleID, errors.New("some errors"))
	actual, err := s.User.Delete(params)
	s.Nil(actual)
	s.NotNil(err)
}
