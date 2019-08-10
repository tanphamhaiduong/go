// @generated
package rolepermission

import (
	"context"
	"errors"

	"github.com/graphql-go/graphql"

	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
)

func (s *RolePermissionResolverTestSuite) TestForwardParams_Success() {
	var (
		sampleID int64 = 1
		params         = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{"id": sampleID}}
		expected       = map[string]interface{}(map[string]interface{}{"id": sampleID})
	)
	actual, err := s.RolePermission.ForwardParams(params)
	s.Nil(err)
	s.Equal(expected, actual)
}

func (s *RolePermissionResolverTestSuite) TestGetByID_Success() {
	var (
		ctx                  = context.Background()
		sampleID       int64 = 1
		params               = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{"id": sampleID}}
		rolepermission models.RolePermission
		args           = arguments.RolePermissionGetByIDArgs{
			ID: 1,
		}
	)
	s.MockIRolePermission.On("GetByID", ctx, args).Return(rolepermission, nil)
	actual, err := s.RolePermission.GetByID(params)
	s.Nil(err)
	s.Equal(rolepermission, actual)
}

func (s *RolePermissionResolverTestSuite) TestGetByID_Fail() {
	var (
		ctx                  = context.Background()
		sampleID       int64 = 1
		params               = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{"id": sampleID}}
		rolepermission models.RolePermission
		args           = arguments.RolePermissionGetByIDArgs{
			ID: 1,
		}
	)
	s.MockIRolePermission.On("GetByID", ctx, args).Return(rolepermission, errors.New("some errors"))
	actual, err := s.RolePermission.GetByID(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *RolePermissionResolverTestSuite) TestList_Success() {
	var (
		ctx             = context.Background()
		params          = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{"page": 1, "pageSize": 10}, Args: map[string]interface{}{}}
		rolepermissions []models.RolePermission
		args            = arguments.RolePermissionListArgs{
			Page:     1,
			PageSize: 10,
		}
	)
	s.MockIRolePermission.On("List", ctx, args).Return(rolepermissions, nil)
	actual, err := s.RolePermission.List(params)
	s.Nil(err)
	s.Equal(rolepermissions, actual)
}

func (s *RolePermissionResolverTestSuite) TestList_Fail() {
	var (
		ctx             = context.Background()
		params          = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{"page": 1, "pageSize": 10}, Args: map[string]interface{}{}}
		rolepermissions []models.RolePermission
		args            = arguments.RolePermissionListArgs{
			Page:     1,
			PageSize: 10,
		}
	)
	s.MockIRolePermission.On("List", ctx, args).Return(rolepermissions, errors.New("some errors"))
	actual, err := s.RolePermission.List(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *RolePermissionResolverTestSuite) TestCount_Success() {
	var (
		ctx    = context.Background()
		params = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args   = arguments.RolePermissionCountArgs{}
		count  int64
	)
	s.MockIRolePermission.On("Count", ctx, args).Return(count, nil)
	actual, err := s.RolePermission.Count(params)
	s.Nil(err)
	s.Equal(count, actual)
}

func (s *RolePermissionResolverTestSuite) TestCount_Fail() {
	var (
		ctx    = context.Background()
		params = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args   = arguments.RolePermissionCountArgs{}
		count  int64
	)
	s.MockIRolePermission.On("Count", ctx, args).Return(count, errors.New("some errors"))
	actual, err := s.RolePermission.Count(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *RolePermissionResolverTestSuite) TestInsert_Success() {
	var (
		ctx            = context.Background()
		params         = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args           = arguments.RolePermissionInsertArgs{}
		rolepermission models.RolePermission
	)
	// Mock Insert
	s.MockIRolePermission.On("Insert", ctx, args).Return(rolepermission, nil)
	// Mock GetByID
	s.MockIRolePermission.On("GetByID", ctx, params).Return(rolepermission, nil)

	actual, err := s.RolePermission.Insert(params)
	s.Nil(err)
	s.Equal(rolepermission, actual)
}

func (s *RolePermissionResolverTestSuite) TestInsert_Fail() {
	var (
		ctx            = context.Background()
		params         = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args           = arguments.RolePermissionInsertArgs{}
		rolepermission models.RolePermission
	)
	s.MockIRolePermission.On("Insert", ctx, args).Return(rolepermission, errors.New("some errors"))
	actual, err := s.RolePermission.Insert(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *RolePermissionResolverTestSuite) TestUpdate_Success() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = graphql.ResolveParams{
			Context: context.Background(),
			Source:  map[string]interface{}{},
			Args:    map[string]interface{}{"id": 1},
		}
		args = arguments.RolePermissionUpdateArgs{
			ID: &sampleID,
		}
		rolepermission = models.RolePermission{}
	)
	// Mock for Update
	s.MockIRolePermission.On("Update", ctx, args).Return(rolepermission, nil)
	// Mock for LoadByID
	s.MockIRolePermission.On("GetByID", ctx, args).Return(rolepermission, nil)
	actual, err := s.RolePermission.Update(params)
	s.Nil(err)
	s.Equal(rolepermission, actual)
}

func (s *RolePermissionResolverTestSuite) TestUpdate_Fail() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = graphql.ResolveParams{
			Context: context.Background(),
			Source:  map[string]interface{}{},
			Args:    map[string]interface{}{"id": 1},
		}
		args = arguments.RolePermissionUpdateArgs{
			ID: &sampleID,
		}
		rolepermission models.RolePermission
	)
	s.MockIRolePermission.On("Update", ctx, args).Return(rolepermission, errors.New("some errors"))
	actual, err := s.RolePermission.Update(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *RolePermissionResolverTestSuite) TestDelete_Success() {
	var (
		ctx    = context.Background()
		params = graphql.ResolveParams{
			Context: context.Background(),
			Source:  map[string]interface{}{},
			Args:    map[string]interface{}{"id": 1},
		}
		args = arguments.RolePermissionDeleteArgs{
			ID: 1,
		}
		sampleID int64 = 1
	)
	s.MockIRolePermission.On("Delete", ctx, args).Return(sampleID, nil)
	actual, err := s.RolePermission.Delete(params)
	s.Nil(err)
	s.Equal(sampleID, actual)
}

func (s *RolePermissionResolverTestSuite) TestDelete_Fail() {
	var (
		ctx    = context.Background()
		params = graphql.ResolveParams{
			Context: context.Background(),
			Source:  map[string]interface{}{},
			Args:    map[string]interface{}{"id": 1},
		}
		args = arguments.RolePermissionDeleteArgs{
			ID: 1,
		}
		sampleID int64
	)
	s.MockIRolePermission.On("Delete", ctx, args).Return(sampleID, errors.New("some errors"))
	actual, err := s.RolePermission.Delete(params)
	s.Nil(actual)
	s.NotNil(err)
}
