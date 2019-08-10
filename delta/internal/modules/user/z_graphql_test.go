// @generated
package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/graphql-go/graphql"

	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
	"github.com/tanphamhaiduong/go/delta/internal/utils"
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

func (s *UserResolverTestSuite) TestCheckPermission_True() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "UserGetByID",
				},
			},
		}
		expected = true
	)
	actual := s.User.checkPermission(*claims, "UserGetByID")
	s.Equal(expected, actual)
}

func (s *UserResolverTestSuite) TestCheckPermission_False() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "UserList",
				},
			},
		}
		expected = false
	)
	actual := s.User.checkPermission(*claims, "UserGetByID")
	s.Equal(expected, actual)
}
func (s *UserResolverTestSuite) TestGetByID_Success() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "UserGetByID",
				},
			},
		}
		sampleID int64 = 1
		ctx            = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params         = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{"id": sampleID}}
		user     models.User
		args     = arguments.UserGetByID{
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
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "UserGetByID",
				},
			},
		}
		sampleID int64 = 1
		ctx            = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params         = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{"id": sampleID}}
		user     models.User
		args     = arguments.UserGetByID{
			ID: 1,
		}
	)
	s.MockIUser.On("GetByID", ctx, args).Return(user, errors.New("some errors"))
	actual, err := s.User.GetByID(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *UserResolverTestSuite) TestGetByID_Fail1() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "UserGetByID",
				},
			},
		}
		sampleID int64 = 1
		ctx            = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params         = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{"id": sampleID}}
		user     models.User
		args     = arguments.UserGetByID{
			ID: 1,
		}
	)
	s.MockIUser.On("GetByID", ctx, args).Return(user, sql.ErrNoRows)
	actual, err := s.User.GetByID(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *UserResolverTestSuite) TestGetByID_Fail2() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "",
				},
			},
		}
		sampleID int64 = 1
		ctx            = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params         = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{"id": sampleID}}
		user     models.User
		args     = arguments.UserGetByID{
			ID: 1,
		}
	)
	s.MockIUser.On("GetByID", ctx, args).Return(user, sql.ErrNoRows)
	actual, err := s.User.GetByID(params)
	s.Nil(actual)
	s.NotNil(err)
}
func (s *UserResolverTestSuite) TestList_Success() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "UserList",
				},
			},
		}
		ctx    = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{"page": 1, "pageSize": 10}, Args: map[string]interface{}{}}
		users  []models.User
		args   = arguments.UserList{
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
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "UserList",
				},
			},
		}
		ctx    = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{"page": 1, "pageSize": 10}, Args: map[string]interface{}{}}
		users  []models.User
		args   = arguments.UserList{
			Page:     1,
			PageSize: 10,
		}
	)
	s.MockIUser.On("List", ctx, args).Return(users, errors.New("some errors"))
	actual, err := s.User.List(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *UserResolverTestSuite) TestList_Fail1() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "UserList",
				},
			},
		}
		ctx    = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{"page": 1, "pageSize": 10}, Args: map[string]interface{}{}}
		users  []models.User
		args   = arguments.UserList{
			Page:     1,
			PageSize: 10,
		}
	)
	s.MockIUser.On("List", ctx, args).Return(users, sql.ErrNoRows)
	actual, err := s.User.List(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *UserResolverTestSuite) TestList_Fail2() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "",
				},
			},
		}
		ctx    = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{"page": 1, "pageSize": 10}, Args: map[string]interface{}{}}
		users  []models.User
		args   = arguments.UserList{
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
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "UserCount",
				},
			},
		}
		ctx    = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args   = arguments.UserCount{}
		count  int64
	)
	s.MockIUser.On("Count", ctx, args).Return(count, nil)
	actual, err := s.User.Count(params)
	s.Nil(err)
	s.Equal(count, actual)
}

func (s *UserResolverTestSuite) TestCount_Fail() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "UserCount",
				},
			},
		}
		ctx    = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args   = arguments.UserCount{}
		count  int64
	)
	s.MockIUser.On("Count", ctx, args).Return(count, errors.New("some errors"))
	actual, err := s.User.Count(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *UserResolverTestSuite) TestCount_Fail1() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "UserCount",
				},
			},
		}
		ctx    = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args   = arguments.UserCount{}
		count  int64
	)
	s.MockIUser.On("Count", ctx, args).Return(count, sql.ErrNoRows)
	actual, err := s.User.Count(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *UserResolverTestSuite) TestCount_Fail2() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "",
				},
			},
		}
		ctx    = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args   = arguments.UserCount{}
		count  int64
	)
	s.MockIUser.On("Count", ctx, args).Return(count, errors.New("some errors"))
	actual, err := s.User.Count(params)
	s.Nil(actual)
	s.NotNil(err)
}
func (s *UserResolverTestSuite) TestInsert_Success() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "UserInsert",
				},
			},
		}
		ctx    = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args   = arguments.UserInsert{}
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
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "UserInsert",
				},
			},
		}
		ctx    = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args   = arguments.UserInsert{}
		user   models.User
	)
	s.MockIUser.On("Insert", ctx, args).Return(user, errors.New("some errors"))
	actual, err := s.User.Insert(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *UserResolverTestSuite) TestInsert_Fail1() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "",
				},
			},
		}
		ctx    = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args   = arguments.UserInsert{}
		user   models.User
	)
	s.MockIUser.On("Insert", ctx, args).Return(user, errors.New("some errors"))
	actual, err := s.User.Insert(params)
	s.Nil(actual)
	s.NotNil(err)
}
func (s *UserResolverTestSuite) TestInsert_Fail2() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "UserInsert",
				},
			},
		}
		ctx    = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args   = arguments.UserInsert{}
		user   models.User
	)
	s.MockIUser.On("Insert", ctx, args).Return(user, sql.ErrNoRows)
	actual, err := s.User.Insert(params)
	s.Nil(actual)
	s.NotNil(err)
}
func (s *UserResolverTestSuite) TestUpdate_Success() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "UserUpdate",
				},
			},
		}
		ctx            = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		sampleID int64 = 1
		params         = graphql.ResolveParams{
			Context: ctx,
			Source:  map[string]interface{}{},
			Args:    map[string]interface{}{"id": 1},
		}
		args = arguments.UserUpdate{
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
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "UserUpdate",
				},
			},
		}
		ctx            = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		sampleID int64 = 1
		params         = graphql.ResolveParams{
			Context: ctx,
			Source:  map[string]interface{}{},
			Args:    map[string]interface{}{"id": 1},
		}
		args = arguments.UserUpdate{
			ID: &sampleID,
		}
		user models.User
	)
	s.MockIUser.On("Update", ctx, args).Return(user, errors.New("some errors"))
	actual, err := s.User.Update(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *UserResolverTestSuite) TestUpdate_Fail1() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "",
				},
			},
		}
		ctx            = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		sampleID int64 = 1
		params         = graphql.ResolveParams{
			Context: ctx,
			Source:  map[string]interface{}{},
			Args:    map[string]interface{}{"id": 1},
		}
		args = arguments.UserUpdate{
			ID: &sampleID,
		}
		user models.User
	)
	s.MockIUser.On("Update", ctx, args).Return(user, errors.New("some errors"))
	actual, err := s.User.Update(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *UserResolverTestSuite) TestUpdate_Fail2() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "UserUpdate",
				},
			},
		}
		ctx            = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		sampleID int64 = 1
		params         = graphql.ResolveParams{
			Context: ctx,
			Source:  map[string]interface{}{},
			Args:    map[string]interface{}{"id": 1},
		}
		args = arguments.UserUpdate{
			ID: &sampleID,
		}
		user models.User
	)
	s.MockIUser.On("Update", ctx, args).Return(user, sql.ErrNoRows)
	actual, err := s.User.Update(params)
	s.Nil(actual)
	s.NotNil(err)
}
