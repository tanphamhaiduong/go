// @generated
package role

import (
	"context"
	"database/sql"
	"errors"

	"github.com/graphql-go/graphql"

	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
	"github.com/tanphamhaiduong/go/delta/internal/utils"
)

func (s *RoleResolverTestSuite) TestCheckPermission_True() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "RoleGetByID",
				},
			},
		}
		expected = true
	)
	actual := s.Role.checkPermission(*claims, "RoleGetByID")
	s.Equal(expected, actual)
}

func (s *RoleResolverTestSuite) TestCheckPermission_False() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "RoleList",
				},
			},
		}
		expected = false
	)
	actual := s.Role.checkPermission(*claims, "RoleGetByID")
	s.Equal(expected, actual)
}
func (s *RoleResolverTestSuite) TestGetByID_Success() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "RoleGetByID",
				},
			},
		}
		sampleID int64 = 1
		ctx            = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params         = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{"id": sampleID}}
		role     models.Role
		args     = arguments.RoleGetByID{
			ID: 1,
		}
	)
	s.MockIRole.On("GetByID", ctx, args).Return(role, nil)
	actual, err := s.Role.GetByID(params)
	s.Nil(err)
	s.Equal(role, actual)
}

func (s *RoleResolverTestSuite) TestGetByID_Fail() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "RoleGetByID",
				},
			},
		}
		sampleID int64 = 1
		ctx            = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params         = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{"id": sampleID}}
		role     models.Role
		args     = arguments.RoleGetByID{
			ID: 1,
		}
	)
	s.MockIRole.On("GetByID", ctx, args).Return(role, errors.New("some errors"))
	actual, err := s.Role.GetByID(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *RoleResolverTestSuite) TestGetByID_Fail1() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "RoleGetByID",
				},
			},
		}
		sampleID int64 = 1
		ctx            = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params         = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{"id": sampleID}}
		role     models.Role
		args     = arguments.RoleGetByID{
			ID: 1,
		}
	)
	s.MockIRole.On("GetByID", ctx, args).Return(role, sql.ErrNoRows)
	actual, err := s.Role.GetByID(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *RoleResolverTestSuite) TestGetByID_Fail2() {
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
		role     models.Role
		args     = arguments.RoleGetByID{
			ID: 1,
		}
	)
	s.MockIRole.On("GetByID", ctx, args).Return(role, sql.ErrNoRows)
	actual, err := s.Role.GetByID(params)
	s.Nil(actual)
	s.NotNil(err)
}
func (s *RoleResolverTestSuite) TestList_Success() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "RoleList",
				},
			},
		}
		ctx    = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{"lastId": 1, "pageSize": 10}}
		roles  []models.Role
		args   = arguments.RoleList{
			LastID:   1,
			PageSize: 10,
		}
	)
	s.MockIRole.On("List", ctx, args).Return(roles, nil)
	actual, err := s.Role.List(params)
	s.Nil(err)
	s.Equal(roles, actual)
}

func (s *RoleResolverTestSuite) TestList_Fail() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "RoleList",
				},
			},
		}
		ctx    = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{"lastId": 1, "pageSize": 10}}
		roles  []models.Role
		args   = arguments.RoleList{
			LastID:   1,
			PageSize: 10,
		}
	)
	s.MockIRole.On("List", ctx, args).Return(roles, errors.New("some errors"))
	actual, err := s.Role.List(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *RoleResolverTestSuite) TestList_Fail1() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "RoleList",
				},
			},
		}
		ctx    = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{"lastId": 1, "pageSize": 10}}
		roles  []models.Role
		args   = arguments.RoleList{
			LastID:   1,
			PageSize: 10,
		}
	)
	s.MockIRole.On("List", ctx, args).Return(roles, sql.ErrNoRows)
	actual, err := s.Role.List(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *RoleResolverTestSuite) TestList_Fail2() {
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
		params = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{"lastId": 1, "pageSize": 10}}
		roles  []models.Role
		args   = arguments.RoleList{
			LastID:   1,
			PageSize: 10,
		}
	)
	s.MockIRole.On("List", ctx, args).Return(roles, errors.New("some errors"))
	actual, err := s.Role.List(params)
	s.Nil(actual)
	s.NotNil(err)
}
func (s *RoleResolverTestSuite) TestCount_Success() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "RoleCount",
				},
			},
		}
		ctx    = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args   = arguments.RoleCount{}
		count  int64
	)
	s.MockIRole.On("Count", ctx, args).Return(count, nil)
	actual, err := s.Role.Count(params)
	s.Nil(err)
	s.Equal(count, actual)
}

func (s *RoleResolverTestSuite) TestCount_Fail() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "RoleCount",
				},
			},
		}
		ctx    = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args   = arguments.RoleCount{}
		count  int64
	)
	s.MockIRole.On("Count", ctx, args).Return(count, errors.New("some errors"))
	actual, err := s.Role.Count(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *RoleResolverTestSuite) TestCount_Fail1() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "RoleCount",
				},
			},
		}
		ctx    = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args   = arguments.RoleCount{}
		count  int64
	)
	s.MockIRole.On("Count", ctx, args).Return(count, sql.ErrNoRows)
	actual, err := s.Role.Count(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *RoleResolverTestSuite) TestCount_Fail2() {
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
		args   = arguments.RoleCount{}
		count  int64
	)
	s.MockIRole.On("Count", ctx, args).Return(count, errors.New("some errors"))
	actual, err := s.Role.Count(params)
	s.Nil(actual)
	s.NotNil(err)
}
func (s *RoleResolverTestSuite) TestInsert_Success() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "RoleInsert",
				},
			},
		}
		ctx    = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args   = arguments.RoleInsert{}
		role   models.Role
	)
	// Mock Insert
	s.MockIRole.On("Insert", ctx, args).Return(role, nil)
	// Mock GetByID
	s.MockIRole.On("GetByID", ctx, params).Return(role, nil)

	actual, err := s.Role.Insert(params)
	s.Nil(err)
	s.Equal(role, actual)
}

func (s *RoleResolverTestSuite) TestInsert_Fail() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "RoleInsert",
				},
			},
		}
		ctx    = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args   = arguments.RoleInsert{}
		role   models.Role
	)
	s.MockIRole.On("Insert", ctx, args).Return(role, errors.New("some errors"))
	actual, err := s.Role.Insert(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *RoleResolverTestSuite) TestInsert_Fail1() {
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
		args   = arguments.RoleInsert{}
		role   models.Role
	)
	s.MockIRole.On("Insert", ctx, args).Return(role, errors.New("some errors"))
	actual, err := s.Role.Insert(params)
	s.Nil(actual)
	s.NotNil(err)
}
func (s *RoleResolverTestSuite) TestInsert_Fail2() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "RoleInsert",
				},
			},
		}
		ctx    = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args   = arguments.RoleInsert{}
		role   models.Role
	)
	s.MockIRole.On("Insert", ctx, args).Return(role, sql.ErrNoRows)
	actual, err := s.Role.Insert(params)
	s.Nil(actual)
	s.NotNil(err)
}
func (s *RoleResolverTestSuite) TestUpdate_Success() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "RoleUpdate",
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
		args = arguments.RoleUpdate{
			ID: &sampleID,
		}
		role = models.Role{}
	)
	// Mock for Update
	s.MockIRole.On("Update", ctx, args).Return(role, nil)
	// Mock for LoadByID
	s.MockIRole.On("GetByID", ctx, args).Return(role, nil)
	actual, err := s.Role.Update(params)
	s.Nil(err)
	s.Equal(role, actual)
}

func (s *RoleResolverTestSuite) TestUpdate_Fail() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "RoleUpdate",
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
		args = arguments.RoleUpdate{
			ID: &sampleID,
		}
		role models.Role
	)
	s.MockIRole.On("Update", ctx, args).Return(role, errors.New("some errors"))
	actual, err := s.Role.Update(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *RoleResolverTestSuite) TestUpdate_Fail1() {
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
		args = arguments.RoleUpdate{
			ID: &sampleID,
		}
		role models.Role
	)
	s.MockIRole.On("Update", ctx, args).Return(role, errors.New("some errors"))
	actual, err := s.Role.Update(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *RoleResolverTestSuite) TestUpdate_Fail2() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "RoleUpdate",
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
		args = arguments.RoleUpdate{
			ID: &sampleID,
		}
		role models.Role
	)
	s.MockIRole.On("Update", ctx, args).Return(role, sql.ErrNoRows)
	actual, err := s.Role.Update(params)
	s.Nil(actual)
	s.NotNil(err)
}
