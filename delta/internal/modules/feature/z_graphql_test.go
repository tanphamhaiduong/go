// @generated
package feature

import (
	"context"
	"errors"

	"github.com/graphql-go/graphql"

	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
)

func (s *FeatureResolverTestSuite) TestGetByID_Success() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{"id": sampleID}}
		feature  models.Feature
		args     = arguments.FeatureGetByIDArgs{
			ID: 1,
		}
	)
	s.MockIFeature.On("GetByID", ctx, args).Return(feature, nil)
	actual, err := s.Feature.GetByID(params)
	s.Nil(err)
	s.Equal(feature, actual)
}

func (s *FeatureResolverTestSuite) TestGetByID_Fail() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{"id": sampleID}}
		feature  models.Feature
		args     = arguments.FeatureGetByIDArgs{
			ID: 1,
		}
	)
	s.MockIFeature.On("GetByID", ctx, args).Return(feature, errors.New("some errors"))
	actual, err := s.Feature.GetByID(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *FeatureResolverTestSuite) TestList_Success() {
	var (
		ctx      = context.Background()
		params   = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{"page": 1, "pageSize": 10}, Args: map[string]interface{}{}}
		features []models.Feature
		args     = arguments.FeatureListArgs{
			Page:     1,
			PageSize: 10,
		}
	)
	s.MockIFeature.On("List", ctx, args).Return(features, nil)
	actual, err := s.Feature.List(params)
	s.Nil(err)
	s.Equal(features, actual)
}

func (s *FeatureResolverTestSuite) TestList_Fail() {
	var (
		ctx      = context.Background()
		params   = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{"page": 1, "pageSize": 10}, Args: map[string]interface{}{}}
		features []models.Feature
		args     = arguments.FeatureListArgs{
			Page:     1,
			PageSize: 10,
		}
	)
	s.MockIFeature.On("List", ctx, args).Return(features, errors.New("some errors"))
	actual, err := s.Feature.List(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *FeatureResolverTestSuite) TestCount_Success() {
	var (
		ctx    = context.Background()
		params = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args   = arguments.FeatureCountArgs{}
		count  int64
	)
	s.MockIFeature.On("Count", ctx, args).Return(count, nil)
	actual, err := s.Feature.Count(params)
	s.Nil(err)
	s.Equal(count, actual)
}

func (s *FeatureResolverTestSuite) TestCount_Fail() {
	var (
		ctx    = context.Background()
		params = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args   = arguments.FeatureCountArgs{}
		count  int64
	)
	s.MockIFeature.On("Count", ctx, args).Return(count, errors.New("some errors"))
	actual, err := s.Feature.Count(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *FeatureResolverTestSuite) TestInsert_Success() {
	var (
		ctx     = context.Background()
		params  = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args    = arguments.FeatureInsertArgs{}
		feature models.Feature
	)
	// Mock Insert
	s.MockIFeature.On("Insert", ctx, args).Return(feature, nil)
	// Mock GetByID
	s.MockIFeature.On("GetByID", ctx, params).Return(feature, nil)

	actual, err := s.Feature.Insert(params)
	s.Nil(err)
	s.Equal(feature, actual)
}

func (s *FeatureResolverTestSuite) TestInsert_Fail() {
	var (
		ctx     = context.Background()
		params  = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args    = arguments.FeatureInsertArgs{}
		feature models.Feature
	)
	s.MockIFeature.On("Insert", ctx, args).Return(feature, errors.New("some errors"))
	actual, err := s.Feature.Insert(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *FeatureResolverTestSuite) TestUpdate_Success() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = graphql.ResolveParams{
			Context: context.Background(),
			Source:  map[string]interface{}{},
			Args:    map[string]interface{}{"id": 1},
		}
		args = arguments.FeatureUpdateArgs{
			ID: &sampleID,
		}
		feature = models.Feature{}
	)
	// Mock for Update
	s.MockIFeature.On("Update", ctx, args).Return(feature, nil)
	// Mock for LoadByID
	s.MockIFeature.On("GetByID", ctx, args).Return(feature, nil)
	actual, err := s.Feature.Update(params)
	s.Nil(err)
	s.Equal(feature, actual)
}

func (s *FeatureResolverTestSuite) TestUpdate_Fail() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = graphql.ResolveParams{
			Context: context.Background(),
			Source:  map[string]interface{}{},
			Args:    map[string]interface{}{"id": 1},
		}
		args = arguments.FeatureUpdateArgs{
			ID: &sampleID,
		}
		feature models.Feature
	)
	s.MockIFeature.On("Update", ctx, args).Return(feature, errors.New("some errors"))
	actual, err := s.Feature.Update(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *FeatureResolverTestSuite) TestDelete_Success() {
	var (
		ctx    = context.Background()
		params = graphql.ResolveParams{
			Context: context.Background(),
			Source:  map[string]interface{}{},
			Args:    map[string]interface{}{"id": 1},
		}
		args = arguments.FeatureDeleteArgs{
			ID: 1,
		}
		sampleID int64 = 1
	)
	s.MockIFeature.On("Delete", ctx, args).Return(sampleID, nil)
	actual, err := s.Feature.Delete(params)
	s.Nil(err)
	s.Equal(sampleID, actual)
}

func (s *FeatureResolverTestSuite) TestDelete_Fail() {
	var (
		ctx    = context.Background()
		params = graphql.ResolveParams{
			Context: context.Background(),
			Source:  map[string]interface{}{},
			Args:    map[string]interface{}{"id": 1},
		}
		args = arguments.FeatureDeleteArgs{
			ID: 1,
		}
		sampleID int64
	)
	s.MockIFeature.On("Delete", ctx, args).Return(sampleID, errors.New("some errors"))
	actual, err := s.Feature.Delete(params)
	s.Nil(actual)
	s.NotNil(err)
}
