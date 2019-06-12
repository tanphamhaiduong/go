// @generated
package feature

import (
	"context"
	"errors"
	"log"

	"github.com/bxcodec/faker"

	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
)

func (s *FeatureHandlerTestSuite) TestGetByID_Success() {
	var (
		ctx    = context.Background()
		params = arguments.FeatureGetByIDArgs{
			ID: 1,
		}
		feature = models.Feature{}
	)
	s.MockIFeature.On("GetByID", ctx, params).Return(feature, nil)
	actual, err := s.Feature.GetByID(ctx, params)
	s.Nil(err)
	s.Equal(feature, actual)
}

func (s *FeatureHandlerTestSuite) TestGetByID_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.FeatureGetByIDArgs{
			ID: 1,
		}
		feature = models.Feature{}
	)
	s.MockIFeature.On("GetByID", ctx, params).Return(feature, errors.New("some errors"))
	actual, err := s.Feature.GetByID(ctx, params)
	s.Equal(feature, actual)
	s.NotNil(err)
}

func (s *FeatureHandlerTestSuite) TestGetByIDs_Success() {
	var (
		ctx    = context.Background()
		params = arguments.FeatureGetByIDsArgs{
			IDs: []int64{1, 2},
		}
		features []models.Feature
	)
	s.MockIFeature.On("GetByIDs", ctx, params).Return(features, nil)
	actual, err := s.Feature.GetByIDs(ctx, params)
	s.Nil(err)
	s.Equal(features, actual)
}

func (s *FeatureHandlerTestSuite) TestGetByIDs_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.FeatureGetByIDsArgs{
			IDs: []int64{1, 2},
		}
		features []models.Feature
	)
	s.MockIFeature.On("GetByIDs", ctx, params).Return(features, errors.New("some errors"))
	actual, err := s.Feature.GetByIDs(ctx, params)
	s.Equal(features, actual)
	s.NotNil(err)
}

func (s *FeatureHandlerTestSuite) TestList_Success() {
	var (
		ctx      = context.Background()
		params   = arguments.FeatureListArgs{}
		features []models.Feature
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	params.Status = "active"
	params.Page = 1
	params.PageSize = 10
	s.MockIFeature.On("List", ctx, params).Return(features, nil)
	actual, err := s.Feature.List(ctx, params)
	s.Nil(err)
	s.Equal(features, actual)
}

func (s *FeatureHandlerTestSuite) TestList_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.FeatureListArgs{
			Page:     1,
			PageSize: 10,
		}
		features []models.Feature
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIFeature.On("List", ctx, params).Return(features, errors.New("some errors"))
	actual, err := s.Feature.List(ctx, params)
	s.Equal(features, actual)
	s.NotNil(err)
}

func (s *FeatureHandlerTestSuite) TestCount_Success() {
	var (
		ctx    = context.Background()
		params = arguments.FeatureCountArgs{}
		count  int64
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	params.Status = "active"
	s.MockIFeature.On("Count", ctx, params).Return(count, nil)
	actual, err := s.Feature.Count(ctx, params)
	s.Nil(err)
	s.Equal(count, actual)
}

func (s *FeatureHandlerTestSuite) TestCount_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.FeatureCountArgs{}
		count  int64
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIFeature.On("Count", ctx, params).Return(count, errors.New("some errors"))
	actual, err := s.Feature.Count(ctx, params)
	s.Equal(count, actual)
	s.NotNil(err)
}

func (s *FeatureHandlerTestSuite) TestInsert_Success() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = arguments.FeatureInsertArgs{}
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	params.Status = "active"
	feature := models.Feature{
		ID:        sampleID,
		URL:       params.URL,
		Meta:      params.Meta,
		CompanyID: params.CompanyID,
		Status:    params.Status,
		CreatedBy: params.CreatedBy,
		UpdatedBy: params.UpdatedBy,
	}
	// Mock Insert
	s.MockIFeature.On("Insert", ctx, params).Return(feature, nil)
	// Mock GetByID
	s.MockIFeature.On("GetByID", ctx, params).Return(feature, nil)

	actual, err := s.Feature.Insert(ctx, params)
	s.Nil(err)
	s.Equal(feature, actual)
}

func (s *FeatureHandlerTestSuite) TestInsert_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.FeatureInsertArgs{
			URL:       "mockString",
			Meta:      "mockString",
			CompanyID: 0,
			Status:    "mockString",
			CreatedBy: "mockString",
			UpdatedBy: "mockString",
		}
		feature = models.Feature{}
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIFeature.On("Insert", ctx, params).Return(feature, errors.New("some errors"))
	actual, err := s.Feature.Insert(ctx, params)
	s.Equal(feature, actual)
	s.NotNil(err)
}

func (s *FeatureHandlerTestSuite) TestUpdate_Success() {
	var (
		ctx    = context.Background()
		params = arguments.FeatureUpdateArgs{}
		status = "active"
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	params.Status = &status
	feature := models.Feature{
		ID:        *params.ID,
		URL:       *params.URL,
		Meta:      *params.Meta,
		CompanyID: *params.CompanyID,
		Status:    *params.Status,
		CreatedBy: *params.CreatedBy,
		UpdatedBy: *params.UpdatedBy,
	}
	// Mock Insert
	s.MockIFeature.On("Update", ctx, params).Return(feature, nil)
	// Mock GetByID
	s.MockIFeature.On("GetByID", ctx, params).Return(feature, nil)

	actual, err := s.Feature.Update(ctx, params)
	s.Nil(err)
	s.Equal(feature, actual)
}

func (s *FeatureHandlerTestSuite) TestUpdate_Fail() {
	var (
		ctx              = context.Background()
		sampleID   int64 = 1
		mockString       = "mockString"
		params           = arguments.FeatureUpdateArgs{
			ID:        &sampleID,
			URL:       &mockString,
			Meta:      &mockString,
			CompanyID: &sampleID,
			Status:    &mockString,
			CreatedBy: &mockString,
			UpdatedBy: &mockString,
		}
		feature = models.Feature{}
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIFeature.On("Update", ctx, params).Return(feature, errors.New("some errors"))
	actual, err := s.Feature.Update(ctx, params)
	s.Equal(feature, actual)
	s.NotNil(err)
}

func (s *FeatureHandlerTestSuite) TestDelete_Success() {
	var (
		ctx    = context.Background()
		params = arguments.FeatureDeleteArgs{}
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIFeature.On("Delete", ctx, params).Return(params.ID, nil)
	actual, err := s.Feature.Delete(ctx, params)
	s.Nil(err)
	s.Equal(params.ID, actual)
}

func (s *FeatureHandlerTestSuite) TestDelete_Fail() {
	var (
		ctx         = context.Background()
		params      = arguments.FeatureDeleteArgs{}
		rowEffected int64
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIFeature.On("Delete", ctx, params).Return(rowEffected, errors.New("some errors"))
	actual, err := s.Feature.Delete(ctx, params)
	s.Equal(rowEffected, actual)
	s.NotNil(err)
}
