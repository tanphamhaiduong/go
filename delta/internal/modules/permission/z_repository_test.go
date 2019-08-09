// @generated
package permission

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/mock"
	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
	"github.com/tanphamhaiduong/go/delta/internal/utils"
)

func (s *PermissionRepositoryTestSuite) TestGetByID_Success() {
	var (
		ctx   = context.Background()
		param = arguments.PermissionGetByIDArgs{
			ID: 1,
		}
		permission models.Permission
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryRowContext", ctx, param.ID).Return(s.MockIRow)
	s.MockIRow.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(nil)
	s.MockIPermission.On("GetByID", ctx, param).Return(permission, nil)
	actual, err := s.Repository.GetByID(ctx, param)
	s.Nil(err)
	s.Equal(permission, actual)
}

func (s *PermissionRepositoryTestSuite) TestGetByID_Fail() {
	var (
		ctx   = context.Background()
		param = arguments.PermissionGetByIDArgs{
			ID: 1,
		}
		permission models.Permission
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryRowContext", ctx, param.ID).Return(s.MockIRow)
	s.MockIRow.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(errors.New("some errors"))
	s.MockIPermission.On("GetByID", ctx, param).Return(permission, errors.New("some errors"))
	actual, err := s.Repository.GetByID(ctx, param)
	s.Equal(permission, actual)
	s.NotNil(err)
}

func (s *PermissionRepositoryTestSuite) TestGetByIDs_Success() {
	var (
		ctx   = context.Background()
		param = arguments.PermissionGetByIDsArgs{
			IDs: []int64{1, 2},
		}
		permissions []models.Permission
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryContext", ctx, mock.Anything, mock.Anything).Return(s.MockIRows, nil)
	s.MockIRows.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(nil)
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockIPermission.On("GetByIDs", ctx, param).Return(permissions, nil)
	actual, err := s.Repository.GetByIDs(ctx, param)
	s.Nil(err)
	s.Equal(permissions, actual)
}

func (s *PermissionRepositoryTestSuite) TestGetByIDs_Fail() {
	var (
		ctx   = context.Background()
		param = arguments.PermissionGetByIDsArgs{
			IDs: []int64{1, 2},
		}
		permissions []models.Permission
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryContext", ctx, mock.Anything, mock.Anything).Return(s.MockIRows, errors.New("some errors"))
	s.MockIRows.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(errors.New("some errors"))
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockIPermission.On("GetByIDs", ctx, param).Return(permissions, errors.New("some errors"))
	actual, err := s.Repository.GetByIDs(ctx, param)
	s.Equal(permissions, actual)
	s.NotNil(err)
}

func (s *PermissionRepositoryTestSuite) TestSetArgsToListSelectBuilder_Success() {
	var (
		ctx    = context.Background()
		params = arguments.PermissionListArgs{
			ID:          1,
			Name:        "mockString",
			Description: "mockString",
			Status:      "active",
			CreatedBy:   "mockString",
			UpdatedBy:   "mockString",
			Page:        1,
			PageSize:    10,
		}
		selectBuilder = sq.Select("*").From("permission")
	)
	offset := utils.CalculateOffsetForPage(params.Page, params.PageSize)
	expectedSelectBuilder := selectBuilder.Where(sq.Eq{"id": params.ID}).Where(sq.Eq{"name": params.Name}).Where(sq.Eq{"description": params.Description}).Where(sq.Eq{"status": params.Status}).Where(sq.Eq{"created_by": params.CreatedBy}).Where(sq.Eq{"updated_by": params.UpdatedBy}).Limit(uint64(params.PageSize)).Offset(uint64(offset))
	expectSQL, expectArgs, expectErr := expectedSelectBuilder.ToSql()
	// Actual
	actual := s.Repository.setArgsToListSelectBuilder(ctx, selectBuilder, params)
	sql, args, err := actual.ToSql()
	s.Nil(err)
	s.Equal(expectSQL, sql)
	s.Equal(expectArgs, args)
	s.Equal(expectErr, err)
}

func (s *PermissionRepositoryTestSuite) TestList_Success() {
	var (
		ctx    = context.Background()
		params = arguments.PermissionListArgs{
			ID:          1,
			Name:        "mockString",
			Description: "mockString",
			Status:      "active",
			CreatedBy:   "mockString",
			UpdatedBy:   "mockString",
			Page:        1,
			PageSize:    10,
		}
		permissions []models.Permission
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryContext", ctx,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(s.MockIRows, nil)
	s.MockIRows.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(nil)
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockIPermission.On("List", ctx, params).Return(permissions, nil)
	actual, err := s.Repository.List(ctx, params)
	s.Nil(err)
	s.Equal(permissions, actual)
}

func (s *PermissionRepositoryTestSuite) TestList_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.PermissionListArgs{
			ID:          1,
			Name:        "mockString",
			Description: "mockString",
			Status:      "active",
			CreatedBy:   "mockString",
			UpdatedBy:   "mockString",
			Page:        1,
			PageSize:    10,
		}
		permissions []models.Permission
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryContext", ctx,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(s.MockIRows, errors.New("some errors"))
	s.MockIRows.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(errors.New("some errors"))
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockIPermission.On("List", ctx, params).Return(permissions, errors.New("some errors"))
	actual, err := s.Repository.List(ctx, params)
	s.Equal(permissions, actual)
	s.NotNil(err)
}

func (s *PermissionRepositoryTestSuite) TestSetArgsToCountSelectBuilder_Success() {
	var (
		ctx    = context.Background()
		params = arguments.PermissionCountArgs{
			ID:          1,
			Name:        "mockString",
			Description: "mockString",
			Status:      "active",
			CreatedBy:   "mockString",
			UpdatedBy:   "mockString",
		}
		selectBuilder = sq.Select("COUNT(id)").From("permission")
	)
	expectedSelectBuilder := selectBuilder.Where(sq.Eq{"id": params.ID}).Where(sq.Eq{"name": params.Name}).Where(sq.Eq{"description": params.Description}).Where(sq.Eq{"status": params.Status}).Where(sq.Eq{"created_by": params.CreatedBy}).Where(sq.Eq{"updated_by": params.UpdatedBy})
	expectSQL, expectArgs, expectErr := expectedSelectBuilder.ToSql()
	// Actual
	actual := s.Repository.setArgsToCountSelectBuilder(ctx, selectBuilder, params)
	sql, args, err := actual.ToSql()
	s.Nil(err)
	s.Equal(expectSQL, sql)
	s.Equal(expectArgs, args)
	s.Equal(expectErr, err)
}

func (s *PermissionRepositoryTestSuite) TestCount_Success() {
	var (
		ctx    = context.Background()
		params = arguments.PermissionCountArgs{
			ID:          1,
			Name:        "mockString",
			Description: "mockString",
			Status:      "active",
			CreatedBy:   "mockString",
			UpdatedBy:   "mockString",
		}
		count int64
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryRowContext", ctx,
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
	s.MockIPermission.On("List", ctx, params).Return(count, nil)
	actual, err := s.Repository.Count(ctx, params)
	s.Nil(err)
	s.Equal(count, actual)
}

func (s *PermissionRepositoryTestSuite) TestCount_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.PermissionCountArgs{
			ID:          1,
			Name:        "mockString",
			Description: "mockString",
			Status:      "active",
			CreatedBy:   "mockString",
			UpdatedBy:   "mockString",
		}
		count int64
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryRowContext", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(s.MockIRows, errors.New("some errors"))
	s.MockIRows.On("Scan", mock.Anything).Return(errors.New("some errors"))
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockIPermission.On("List", ctx, params).Return(count, errors.New("some errors"))
	actual, err := s.Repository.Count(ctx, params)
	s.Equal(count, actual)
	s.NotNil(err)
}

func (s *PermissionRepositoryTestSuite) TestInsert_Success() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = arguments.PermissionInsertArgs{
			Name:        "mockString",
			Description: "mockString",
			Status:      "active",
			CreatedBy:   "mockString",
			UpdatedBy:   "mockString",
		}
		expected models.Permission
	)
	permission := models.Permission{
		ID:          sampleID,
		Name:        params.Name,
		Description: params.Description,
		Status:      params.Status,
		CreatedBy:   params.CreatedBy,
		UpdatedBy:   params.UpdatedBy,
	}
	//Mock Insert
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx,
		params.Name,
		params.Description,
		params.Status,
		params.CreatedBy,
		params.UpdatedBy,
	).Return(s.MockIResult, nil)
	s.MockIResult.On("LastInsertId").Return(sampleID, nil)
	s.MockIPermission.On("Insert", ctx, params).Return(permission, nil)
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
	).Return(nil)
	s.MockIPermission.On("GetByID", ctx, params).Return(permission, nil)

	actual, err := s.Repository.Insert(ctx, params)
	s.Nil(err)
	s.Equal(expected, actual)
}

func (s *PermissionRepositoryTestSuite) TestInsert_Fail() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = arguments.PermissionInsertArgs{
			Name:        "mockString",
			Description: "mockString",
			Status:      "active",
			CreatedBy:   "mockString",
			UpdatedBy:   "mockString",
		}
		permission models.Permission
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("ExecContext", ctx,
		params.Name,
		params.Description,
		params.Status,
		params.CreatedBy,
		params.UpdatedBy,
	).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("LastInsertId").Return(sampleID, errors.New("some errors"))
	s.MockIPermission.On("Insert", ctx, params).Return(permission, errors.New("some errors"))
	actual, err := s.Repository.Insert(ctx, params)
	s.Equal(permission, actual)
	s.NotNil(err)
}

func (s *PermissionRepositoryTestSuite) TestSetArgsToUpdateBuilder_Success() {
	var (
		ctx              = context.Background()
		sampleID   int64 = 1
		mockString       = "mockString"
		status           = "active"
		params           = arguments.PermissionUpdateArgs{
			ID:          &sampleID,
			Name:        &mockString,
			Description: &mockString,
			Status:      &status,
			CreatedBy:   &mockString,
			UpdatedBy:   &mockString,
		}
	)
	updateBuilder := sq.Update("permission").Where(sq.Eq{"id": *params.ID})
	expectedSelectBuilder := updateBuilder.Set("name", *params.Name).Set("description", *params.Description).Set("status", *params.Status).Set("created_by", *params.CreatedBy).Set("updated_by", *params.UpdatedBy)
	actual := s.Repository.setArgsToUpdateBuilder(ctx, updateBuilder, params)
	s.Equal(expectedSelectBuilder, actual)
}

func (s *PermissionRepositoryTestSuite) TestUpdate_Success() {
	var (
		ctx              = context.Background()
		sampleID   int64 = 1
		mockString       = "mockString"
		status           = "active"
		params           = arguments.PermissionUpdateArgs{
			ID:          &sampleID,
			Name:        &mockString,
			Description: &mockString,
			Status:      &status,
			CreatedBy:   &mockString,
			UpdatedBy:   &mockString,
		}
		expected models.Permission
	)
	permission := models.Permission{
		ID:          *params.ID,
		Name:        *params.Name,
		Description: *params.Description,
		Status:      *params.Status,
		CreatedBy:   *params.CreatedBy,
		UpdatedBy:   *params.UpdatedBy,
	}
	// Mock Update
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx,
		*params.Name,
		*params.Description,
		*params.Status,
		*params.CreatedBy,
		*params.UpdatedBy,
		*params.ID,
	).Return(s.MockIResult, nil)
	s.MockIResult.On("RowsAffected").Return(*params.ID, nil)
	s.MockIPermission.On("Update", ctx, params).Return(permission, nil)
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
	).Return(nil)
	s.MockIPermission.On("GetByID", ctx, params).Return(permission, nil)

	actual, err := s.Repository.Update(ctx, params)
	s.Nil(err)
	s.Equal(expected, actual)
}

func (s *PermissionRepositoryTestSuite) TestUpdate_Fail() {
	var (
		ctx              = context.Background()
		sampleID   int64 = 1
		mockString       = "mockString"
		status           = "active"
		params           = arguments.PermissionUpdateArgs{
			ID:          &sampleID,
			Name:        &mockString,
			Description: &mockString,
			Status:      &status,
			CreatedBy:   &mockString,
			UpdatedBy:   &mockString,
		}
		permission models.Permission
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("ExecContext", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("RowsAffected").Return(sampleID, errors.New("some errors"))
	s.MockIPermission.On("Update", ctx, params).Return(permission, errors.New("some errors"))
	actual, err := s.Repository.Update(ctx, params)
	s.Equal(permission, actual)
	s.NotNil(err)
}

func (s *PermissionRepositoryTestSuite) TestDelete_Success() {
	var (
		ctx   = context.Background()
		param = arguments.PermissionDeleteArgs{
			ID: 1,
		}
		rowEffected int64 = 1
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx, param.ID).Return(s.MockIResult, nil)
	s.MockIResult.On("RowsAffected").Return(param.ID, nil)
	s.MockIPermission.On("Delete", ctx, param).Return(rowEffected, nil)
	actual, err := s.Repository.Delete(ctx, param)
	s.Nil(err)
	s.Equal(param.ID, actual)
}

func (s *PermissionRepositoryTestSuite) TestDelete_Fail() {
	var (
		ctx   = context.Background()
		param = arguments.PermissionDeleteArgs{
			ID: 1,
		}
		rowEffected int64
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("ExecContext", ctx, param.ID).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("RowsAffected").Return(param.ID, errors.New("some errors"))
	s.MockIPermission.On("Delete", ctx, param).Return(rowEffected, errors.New("some errors"))
	actual, err := s.Repository.Delete(ctx, param)
	s.Equal(rowEffected, actual)
	s.NotNil(err)
}
