// @generated
package permission

import (
	"context"
	"database/sql"
	"errors"

	"github.com/graphql-go/graphql"

	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
	"github.com/tanphamhaiduong/go/delta/internal/utils"
)

func (s *PermissionResolverTestSuite) TestCheckPermission_True() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "PermissionGetByID",
				},
			},
		}
		expected = true
	)
	actual := s.Permission.checkPermission(*claims, "PermissionGetByID")
	s.Equal(expected, actual)
}

func (s *PermissionResolverTestSuite) TestCheckPermission_False() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "PermissionList",
				},
			},
		}
		expected = false
	)
	actual := s.Permission.checkPermission(*claims, "PermissionGetByID")
	s.Equal(expected, actual)
}
func (s *PermissionResolverTestSuite) TestGetByID_Success() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "PermissionGetByID",
				},
			},
		}
		sampleID   int64 = 1
		ctx              = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params           = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{"id": sampleID}}
		permission models.Permission
		args       = arguments.PermissionGetByID{
			ID: 1,
		}
	)
	s.MockIPermission.On("GetByID", ctx, args).Return(permission, nil)
	actual, err := s.Permission.GetByID(params)
	s.Nil(err)
	s.Equal(permission, actual)
}

func (s *PermissionResolverTestSuite) TestGetByID_Fail() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "PermissionGetByID",
				},
			},
		}
		sampleID   int64 = 1
		ctx              = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params           = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{"id": sampleID}}
		permission models.Permission
		args       = arguments.PermissionGetByID{
			ID: 1,
		}
	)
	s.MockIPermission.On("GetByID", ctx, args).Return(permission, errors.New("some errors"))
	actual, err := s.Permission.GetByID(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *PermissionResolverTestSuite) TestGetByID_Fail1() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "PermissionGetByID",
				},
			},
		}
		sampleID   int64 = 1
		ctx              = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params           = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{"id": sampleID}}
		permission models.Permission
		args       = arguments.PermissionGetByID{
			ID: 1,
		}
	)
	s.MockIPermission.On("GetByID", ctx, args).Return(permission, sql.ErrNoRows)
	actual, err := s.Permission.GetByID(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *PermissionResolverTestSuite) TestGetByID_Fail2() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "",
				},
			},
		}
		sampleID   int64 = 1
		ctx              = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params           = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{"id": sampleID}}
		permission models.Permission
		args       = arguments.PermissionGetByID{
			ID: 1,
		}
	)
	s.MockIPermission.On("GetByID", ctx, args).Return(permission, sql.ErrNoRows)
	actual, err := s.Permission.GetByID(params)
	s.Nil(actual)
	s.NotNil(err)
}
func (s *PermissionResolverTestSuite) TestList_Success() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "PermissionList",
				},
			},
		}
		ctx         = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params      = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{"page": 1, "pageSize": 10}, Args: map[string]interface{}{}}
		permissions []models.Permission
		args        = arguments.PermissionList{
			Page:     1,
			PageSize: 10,
		}
	)
	s.MockIPermission.On("List", ctx, args).Return(permissions, nil)
	actual, err := s.Permission.List(params)
	s.Nil(err)
	s.Equal(permissions, actual)
}

func (s *PermissionResolverTestSuite) TestList_Fail() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "PermissionList",
				},
			},
		}
		ctx         = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params      = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{"page": 1, "pageSize": 10}, Args: map[string]interface{}{}}
		permissions []models.Permission
		args        = arguments.PermissionList{
			Page:     1,
			PageSize: 10,
		}
	)
	s.MockIPermission.On("List", ctx, args).Return(permissions, errors.New("some errors"))
	actual, err := s.Permission.List(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *PermissionResolverTestSuite) TestList_Fail1() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "PermissionList",
				},
			},
		}
		ctx         = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params      = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{"page": 1, "pageSize": 10}, Args: map[string]interface{}{}}
		permissions []models.Permission
		args        = arguments.PermissionList{
			Page:     1,
			PageSize: 10,
		}
	)
	s.MockIPermission.On("List", ctx, args).Return(permissions, sql.ErrNoRows)
	actual, err := s.Permission.List(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *PermissionResolverTestSuite) TestList_Fail2() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "",
				},
			},
		}
		ctx         = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params      = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{"page": 1, "pageSize": 10}, Args: map[string]interface{}{}}
		permissions []models.Permission
		args        = arguments.PermissionList{
			Page:     1,
			PageSize: 10,
		}
	)
	s.MockIPermission.On("List", ctx, args).Return(permissions, errors.New("some errors"))
	actual, err := s.Permission.List(params)
	s.Nil(actual)
	s.NotNil(err)
}
func (s *PermissionResolverTestSuite) TestCount_Success() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "PermissionCount",
				},
			},
		}
		ctx    = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args   = arguments.PermissionCount{}
		count  int64
	)
	s.MockIPermission.On("Count", ctx, args).Return(count, nil)
	actual, err := s.Permission.Count(params)
	s.Nil(err)
	s.Equal(count, actual)
}

func (s *PermissionResolverTestSuite) TestCount_Fail() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "PermissionCount",
				},
			},
		}
		ctx    = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args   = arguments.PermissionCount{}
		count  int64
	)
	s.MockIPermission.On("Count", ctx, args).Return(count, errors.New("some errors"))
	actual, err := s.Permission.Count(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *PermissionResolverTestSuite) TestCount_Fail1() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "PermissionCount",
				},
			},
		}
		ctx    = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args   = arguments.PermissionCount{}
		count  int64
	)
	s.MockIPermission.On("Count", ctx, args).Return(count, sql.ErrNoRows)
	actual, err := s.Permission.Count(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *PermissionResolverTestSuite) TestCount_Fail2() {
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
		args   = arguments.PermissionCount{}
		count  int64
	)
	s.MockIPermission.On("Count", ctx, args).Return(count, errors.New("some errors"))
	actual, err := s.Permission.Count(params)
	s.Nil(actual)
	s.NotNil(err)
}
func (s *PermissionResolverTestSuite) TestInsert_Success() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "PermissionInsert",
				},
			},
		}
		ctx        = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params     = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args       = arguments.PermissionInsert{}
		permission models.Permission
	)
	// Mock Insert
	s.MockIPermission.On("Insert", ctx, args).Return(permission, nil)
	// Mock GetByID
	s.MockIPermission.On("GetByID", ctx, params).Return(permission, nil)

	actual, err := s.Permission.Insert(params)
	s.Nil(err)
	s.Equal(permission, actual)
}

func (s *PermissionResolverTestSuite) TestInsert_Fail() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "PermissionInsert",
				},
			},
		}
		ctx        = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params     = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args       = arguments.PermissionInsert{}
		permission models.Permission
	)
	s.MockIPermission.On("Insert", ctx, args).Return(permission, errors.New("some errors"))
	actual, err := s.Permission.Insert(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *PermissionResolverTestSuite) TestInsert_Fail1() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "",
				},
			},
		}
		ctx        = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params     = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args       = arguments.PermissionInsert{}
		permission models.Permission
	)
	s.MockIPermission.On("Insert", ctx, args).Return(permission, errors.New("some errors"))
	actual, err := s.Permission.Insert(params)
	s.Nil(actual)
	s.NotNil(err)
}
func (s *PermissionResolverTestSuite) TestInsert_Fail2() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "PermissionInsert",
				},
			},
		}
		ctx        = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params     = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args       = arguments.PermissionInsert{}
		permission models.Permission
	)
	s.MockIPermission.On("Insert", ctx, args).Return(permission, sql.ErrNoRows)
	actual, err := s.Permission.Insert(params)
	s.Nil(actual)
	s.NotNil(err)
}
