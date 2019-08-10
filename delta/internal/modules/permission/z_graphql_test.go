// @generated
package permission

import (
	"context"
	"errors"

	"github.com/graphql-go/graphql"

	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
)

func (s *PermissionResolverTestSuite) TestForwardParams_Success() {
	var (
		sampleID int64 = 1
		params         = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{"id": sampleID}}
		expected       = map[string]interface{}(map[string]interface{}{"id": sampleID})
	)
	actual, err := s.Permission.ForwardParams(params)
	s.Nil(err)
	s.Equal(expected, actual)
}

func (s *PermissionResolverTestSuite) TestGetByID_Success() {
	var (
		ctx              = context.Background()
		sampleID   int64 = 1
		params           = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{"id": sampleID}}
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
		ctx              = context.Background()
		sampleID   int64 = 1
		params           = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{"id": sampleID}}
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

func (s *PermissionResolverTestSuite) TestList_Success() {
	var (
		ctx         = context.Background()
		params      = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{"page": 1, "pageSize": 10}, Args: map[string]interface{}{}}
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
		ctx         = context.Background()
		params      = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{"page": 1, "pageSize": 10}, Args: map[string]interface{}{}}
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
		ctx    = context.Background()
		params = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{}}
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
		ctx    = context.Background()
		params = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{}}
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
		ctx        = context.Background()
		params     = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{}}
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
		ctx        = context.Background()
		params     = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args       = arguments.PermissionInsert{}
		permission models.Permission
	)
	s.MockIPermission.On("Insert", ctx, args).Return(permission, errors.New("some errors"))
	actual, err := s.Permission.Insert(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *PermissionResolverTestSuite) TestUpdate_Success() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = graphql.ResolveParams{
			Context: context.Background(),
			Source:  map[string]interface{}{},
			Args:    map[string]interface{}{"id": 1},
		}
		args = arguments.PermissionUpdate{
			ID: &sampleID,
		}
		permission = models.Permission{}
	)
	// Mock for Update
	s.MockIPermission.On("Update", ctx, args).Return(permission, nil)
	// Mock for LoadByID
	s.MockIPermission.On("GetByID", ctx, args).Return(permission, nil)
	actual, err := s.Permission.Update(params)
	s.Nil(err)
	s.Equal(permission, actual)
}

func (s *PermissionResolverTestSuite) TestUpdate_Fail() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = graphql.ResolveParams{
			Context: context.Background(),
			Source:  map[string]interface{}{},
			Args:    map[string]interface{}{"id": 1},
		}
		args = arguments.PermissionUpdate{
			ID: &sampleID,
		}
		permission models.Permission
	)
	s.MockIPermission.On("Update", ctx, args).Return(permission, errors.New("some errors"))
	actual, err := s.Permission.Update(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *PermissionResolverTestSuite) TestDelete_Success() {
	var (
		ctx    = context.Background()
		params = graphql.ResolveParams{
			Context: context.Background(),
			Source:  map[string]interface{}{},
			Args:    map[string]interface{}{"id": 1},
		}
		args = arguments.PermissionDelete{
			ID: 1,
		}
		sampleID int64 = 1
	)
	s.MockIPermission.On("Delete", ctx, args).Return(sampleID, nil)
	actual, err := s.Permission.Delete(params)
	s.Nil(err)
	s.Equal(sampleID, actual)
}

func (s *PermissionResolverTestSuite) TestDelete_Fail() {
	var (
		ctx    = context.Background()
		params = graphql.ResolveParams{
			Context: context.Background(),
			Source:  map[string]interface{}{},
			Args:    map[string]interface{}{"id": 1},
		}
		args = arguments.PermissionDelete{
			ID: 1,
		}
		sampleID int64
	)
	s.MockIPermission.On("Delete", ctx, args).Return(sampleID, errors.New("some errors"))
	actual, err := s.Permission.Delete(params)
	s.Nil(actual)
	s.NotNil(err)
}
