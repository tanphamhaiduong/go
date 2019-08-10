// @generated
package role

import (
	"context"
	"errors"

	"github.com/graphql-go/graphql"

	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
)

func (s *RoleResolverTestSuite) TestForwardParams_Success() {
	var (
		sampleID int64 = 1
		params         = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{"id": sampleID}}
		expected       = map[string]interface{}(map[string]interface{}{"id": sampleID})
	)
	actual, err := s.Role.ForwardParams(params)
	s.Nil(err)
	s.Equal(expected, actual)
}

func (s *RoleResolverTestSuite) TestGetByID_Success() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{"id": sampleID}}
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
		ctx            = context.Background()
		sampleID int64 = 1
		params         = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{"id": sampleID}}
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

func (s *RoleResolverTestSuite) TestList_Success() {
	var (
		ctx    = context.Background()
		params = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{"page": 1, "pageSize": 10}, Args: map[string]interface{}{}}
		roles  []models.Role
		args   = arguments.RoleList{
			Page:     1,
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
		ctx    = context.Background()
		params = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{"page": 1, "pageSize": 10}, Args: map[string]interface{}{}}
		roles  []models.Role
		args   = arguments.RoleList{
			Page:     1,
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
		ctx    = context.Background()
		params = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{}}
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
		ctx    = context.Background()
		params = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{}}
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
		ctx    = context.Background()
		params = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{}}
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
		ctx    = context.Background()
		params = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args   = arguments.RoleInsert{}
		role   models.Role
	)
	s.MockIRole.On("Insert", ctx, args).Return(role, errors.New("some errors"))
	actual, err := s.Role.Insert(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *RoleResolverTestSuite) TestUpdate_Success() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = graphql.ResolveParams{
			Context: context.Background(),
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
		ctx            = context.Background()
		sampleID int64 = 1
		params         = graphql.ResolveParams{
			Context: context.Background(),
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

func (s *RoleResolverTestSuite) TestDelete_Success() {
	var (
		ctx    = context.Background()
		params = graphql.ResolveParams{
			Context: context.Background(),
			Source:  map[string]interface{}{},
			Args:    map[string]interface{}{"id": 1},
		}
		args = arguments.RoleDelete{
			ID: 1,
		}
		sampleID int64 = 1
	)
	s.MockIRole.On("Delete", ctx, args).Return(sampleID, nil)
	actual, err := s.Role.Delete(params)
	s.Nil(err)
	s.Equal(sampleID, actual)
}

func (s *RoleResolverTestSuite) TestDelete_Fail() {
	var (
		ctx    = context.Background()
		params = graphql.ResolveParams{
			Context: context.Background(),
			Source:  map[string]interface{}{},
			Args:    map[string]interface{}{"id": 1},
		}
		args = arguments.RoleDelete{
			ID: 1,
		}
		sampleID int64
	)
	s.MockIRole.On("Delete", ctx, args).Return(sampleID, errors.New("some errors"))
	actual, err := s.Role.Delete(params)
	s.Nil(actual)
	s.NotNil(err)
}
