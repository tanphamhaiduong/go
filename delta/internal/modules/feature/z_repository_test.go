// @generated
package feature

import (
	"context"
	"errors"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/mock"
	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
	"github.com/tanphamhaiduong/go/delta/internal/utils"
)

func (s *FeatureRepositoryTestSuite) TestGetByID_Success() {
	var (
		ctx    = context.Background()
		params = arguments.FeatureGetByIDArgs{
			ID: 1,
		}
		feature models.Feature
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryRowContext", ctx, mock.Anything).Return(s.MockIRow)
	s.MockIRow.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(nil)
	s.MockIFeature.On("GetByID", ctx, params).Return(feature, nil)
	actual, err := s.Repository.GetByID(ctx, params)
	s.Nil(err)
	s.Equal(feature, actual)
}

func (s *FeatureRepositoryTestSuite) TestGetByID_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.FeatureGetByIDArgs{
			ID: 1,
		}
		feature models.Feature
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryRowContext", ctx, mock.Anything).Return(s.MockIRow)
	s.MockIRow.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("some errors"))
	s.MockIFeature.On("GetByID", ctx, params).Return(feature, errors.New("some errors"))
	actual, err := s.Repository.GetByID(ctx, params)
	s.Equal(feature, actual)
	s.NotNil(err)
}

func (s *FeatureRepositoryTestSuite) TestGetByIDs_Success() {
	var (
		ctx    = context.Background()
		params = arguments.FeatureGetByIDsArgs{
			IDs: []int64{1, 2},
		}
		features []models.Feature
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryContext", ctx, mock.Anything, mock.Anything).Return(s.MockIRows, nil)
	s.MockIRows.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockIFeature.On("GetByIDs", ctx, params).Return(features, nil)
	actual, err := s.Repository.GetByIDs(ctx, params)
	s.Nil(err)
	s.Equal(features, actual)
}

func (s *FeatureRepositoryTestSuite) TestGetByIDs_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.FeatureGetByIDsArgs{
			IDs: []int64{1, 2},
		}
		features []models.Feature
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryContext", ctx, mock.Anything, mock.Anything).Return(s.MockIRows, errors.New("some errors"))
	s.MockIRows.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("some errors"))
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockIFeature.On("GetByIDs", ctx, params).Return(features, errors.New("some errors"))
	actual, err := s.Repository.GetByIDs(ctx, params)
	s.Equal(features, actual)
	s.NotNil(err)
}

func (s *FeatureRepositoryTestSuite) TestSetArgsToListSelectBuilder_Success() {
	var (
		params        = arguments.FeatureListArgs{}
		selectBuilder = sq.Select("*").From("feature")
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	offset := utils.CalculateOffsetForPage(params.Page, params.PageSize)
	expectedSelectBuilder := selectBuilder.Where(sq.Eq{"id": params.ID}).Where(sq.Eq{"url": params.URL}).Where(sq.Eq{"meta": params.Meta}).Where(sq.Eq{"company_id": params.CompanyID}).Where(sq.Eq{"status": params.Status}).Where(sq.Eq{"created_by": params.CreatedBy}).Where(sq.Eq{"updated_by": params.UpdatedBy}).Limit(uint64(params.PageSize)).Offset(uint64(offset))
	expectSQL, expectArgs, expectErr := expectedSelectBuilder.ToSql()
	// Actual
	actual := s.Repository.setArgsToListSelectBuilder(selectBuilder, params)
	sql, args, err := actual.ToSql()
	s.Nil(err)
	s.Equal(expectSQL, sql)
	s.Equal(expectArgs, args)
	s.Equal(expectErr, err)
}

func (s *FeatureRepositoryTestSuite) TestList_Success() {
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
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryContext", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(s.MockIRows, nil)
	s.MockIRows.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockIFeature.On("List", ctx, params).Return(features, nil)
	actual, err := s.Repository.List(ctx, params)
	s.Nil(err)
	s.Equal(features, actual)
}

func (s *FeatureRepositoryTestSuite) TestList_Fail() {
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
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryContext", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(s.MockIRows, errors.New("some errors"))
	s.MockIRows.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("some errors"))
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockIFeature.On("List", ctx, params).Return(features, errors.New("some errors"))
	actual, err := s.Repository.List(ctx, params)
	s.Equal(features, actual)
	s.NotNil(err)
}

func (s *FeatureRepositoryTestSuite) TestSetArgsToCountSelectBuilder_Success() {
	var (
		params        = arguments.FeatureCountArgs{}
		selectBuilder = sq.Select("COUNT(id)").From("feature")
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	expectedSelectBuilder := selectBuilder.Where(sq.Eq{"id": params.ID}).Where(sq.Eq{"url": params.URL}).Where(sq.Eq{"meta": params.Meta}).Where(sq.Eq{"company_id": params.CompanyID}).Where(sq.Eq{"status": params.Status}).Where(sq.Eq{"created_by": params.CreatedBy}).Where(sq.Eq{"updated_by": params.UpdatedBy})
	expectSQL, expectArgs, expectErr := expectedSelectBuilder.ToSql()
	// Actual
	actual := s.Repository.setArgsToCountSelectBuilder(selectBuilder, params)
	sql, args, err := actual.ToSql()
	s.Nil(err)
	s.Equal(expectSQL, sql)
	s.Equal(expectArgs, args)
	s.Equal(expectErr, err)
}

func (s *FeatureRepositoryTestSuite) TestCount_Success() {
	var (
		ctx    = context.Background()
		params = arguments.FeatureCountArgs{}
		count  int64
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryRowContext", ctx,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(s.MockIRows, nil)
	s.MockIRows.On("Scan", mock.Anything).Return(nil)
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockIFeature.On("List", ctx, params).Return(count, nil)
	actual, err := s.Repository.Count(ctx, params)
	s.Nil(err)
	s.Equal(count, actual)
}

func (s *FeatureRepositoryTestSuite) TestCount_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.FeatureCountArgs{}
		count  int64
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryRowContext", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(s.MockIRows, errors.New("some errors"))
	s.MockIRows.On("Scan", mock.Anything).Return(errors.New("some errors"))
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockIFeature.On("List", ctx, params).Return(count, errors.New("some errors"))
	actual, err := s.Repository.Count(ctx, params)
	s.Equal(count, actual)
	s.NotNil(err)
}

func (s *FeatureRepositoryTestSuite) TestInsert_Success() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = arguments.FeatureInsertArgs{}
		expected models.Feature
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	feature := models.Feature{
		ID:        sampleID,
		URL:       params.URL,
		Meta:      params.Meta,
		CompanyID: params.CompanyID,
		Status:    params.Status,
		CreatedBy: params.CreatedBy,
		UpdatedBy: params.UpdatedBy,
	}
	//Mock Insert
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx,
		params.URL,
		params.Meta,
		params.CompanyID,
		params.Status,
		params.CreatedBy,
		params.UpdatedBy,
	).Return(s.MockIResult, nil)
	s.MockIResult.On("LastInsertId").Return(sampleID, nil)
	s.MockIFeature.On("Insert", ctx, params).Return(feature, nil)
	// Mock GetByID
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryRowContext", ctx, mock.Anything).Return(s.MockIRow)
	s.MockIRow.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(nil)
	s.MockIFeature.On("GetByID", ctx, params).Return(feature, nil)

	actual, err := s.Repository.Insert(ctx, params)
	s.Nil(err)
	s.Equal(expected, actual)
}

func (s *FeatureRepositoryTestSuite) TestInsert_Fail() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = arguments.FeatureInsertArgs{
			URL:       "mockString",
			Meta:      "mockString",
			CompanyID: 0,
			Status:    "mockString",
			CreatedBy: "mockString",
			UpdatedBy: "mockString",
		}
		feature models.Feature
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("ExecContext", ctx,
		params.URL,
		params.Meta,
		params.CompanyID,
		params.Status,
		params.CreatedBy,
		params.UpdatedBy,
	).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("LastInsertId").Return(sampleID, errors.New("some errors"))
	s.MockIFeature.On("Insert", ctx, params).Return(feature, errors.New("some errors"))
	actual, err := s.Repository.Insert(ctx, params)
	s.Equal(feature, actual)
	s.NotNil(err)
}

func (s *FeatureRepositoryTestSuite) TestSetArgsToUpdateBuilder_Success() {
	var (
		params = arguments.FeatureUpdateArgs{}
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	updateBuilder := sq.Update("feature").Where(sq.Eq{"id": *params.ID})
	expectedSelectBuilder := updateBuilder.Set("url", *params.URL).Set("meta", *params.Meta).Set("company_id", *params.CompanyID).Set("status", *params.Status).Set("created_by", *params.CreatedBy).Set("updated_by", *params.UpdatedBy)
	actual := s.Repository.setArgsToUpdateBuilder(updateBuilder, params)
	s.Equal(expectedSelectBuilder, actual)
}

func (s *FeatureRepositoryTestSuite) TestUpdate_Success() {
	var (
		ctx      = context.Background()
		params   = arguments.FeatureUpdateArgs{}
		expected models.Feature
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	feature := models.Feature{
		ID:        *params.ID,
		URL:       *params.URL,
		Meta:      *params.Meta,
		CompanyID: *params.CompanyID,
		Status:    *params.Status,
		CreatedBy: *params.CreatedBy,
		UpdatedBy: *params.UpdatedBy,
	}
	// Mock Update
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx,
		*params.URL,
		*params.Meta,
		*params.CompanyID,
		*params.Status,
		*params.CreatedBy,
		*params.UpdatedBy,
		*params.ID,
	).Return(s.MockIResult, nil)
	s.MockIResult.On("RowsAffected").Return(*params.ID, nil)
	s.MockIFeature.On("Update", ctx, params).Return(feature, nil)
	// Mock GetByID
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryRowContext", ctx, mock.Anything).Return(s.MockIRow)
	s.MockIRow.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(nil)
	s.MockIFeature.On("GetByID", ctx, params).Return(feature, nil)

	actual, err := s.Repository.Update(ctx, params)
	s.Nil(err)
	s.Equal(expected, actual)
}

func (s *FeatureRepositoryTestSuite) TestUpdate_Fail() {
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
		feature models.Feature
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("ExecContext", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("RowsAffected").Return(sampleID, errors.New("some errors"))
	s.MockIFeature.On("Update", ctx, params).Return(feature, errors.New("some errors"))
	actual, err := s.Repository.Update(ctx, params)
	s.Equal(feature, actual)
	s.NotNil(err)
}

func (s *FeatureRepositoryTestSuite) TestDelete_Success() {
	var (
		ctx               = context.Background()
		params            = arguments.FeatureDeleteArgs{}
		rowEffected int64 = 1
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx, params.ID).Return(s.MockIResult, nil)
	s.MockIResult.On("RowsAffected").Return(params.ID, nil)
	s.MockIFeature.On("Delete", ctx, params).Return(rowEffected, nil)
	actual, err := s.Repository.Delete(ctx, params)
	s.Nil(err)
	s.Equal(params.ID, actual)
}

func (s *FeatureRepositoryTestSuite) TestDelete_Fail() {
	var (
		ctx         = context.Background()
		params      = arguments.FeatureDeleteArgs{}
		rowEffected int64
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("ExecContext", ctx, params.ID).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("RowsAffected").Return(params.ID, errors.New("some errors"))
	s.MockIFeature.On("Delete", ctx, params).Return(rowEffected, errors.New("some errors"))
	actual, err := s.Repository.Delete(ctx, params)
	s.Equal(rowEffected, actual)
	s.NotNil(err)
}
