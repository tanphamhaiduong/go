// @generated
package rolepermission

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/mock"
	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
	"github.com/tanphamhaiduong/go/delta/internal/utils"
)

func (s *RolePermissionRepositoryTestSuite) TestGetByID_Success() {
	var (
		ctx   = context.Background()
		param = arguments.RolePermissionGetByIDArgs{
			ID: 1,
		}
		rolepermission models.RolePermission
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryRowContext", ctx, param.ID).Return(s.MockIRow)
	s.MockIRow.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(nil)
	s.MockIRolePermission.On("GetByID", ctx, param).Return(rolepermission, nil)
	actual, err := s.Repository.GetByID(ctx, param)
	s.Nil(err)
	s.Equal(rolepermission, actual)
}

func (s *RolePermissionRepositoryTestSuite) TestGetByID_Fail() {
	var (
		ctx   = context.Background()
		param = arguments.RolePermissionGetByIDArgs{
			ID: 1,
		}
		rolepermission models.RolePermission
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryRowContext", ctx, param.ID).Return(s.MockIRow)
	s.MockIRow.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(errors.New("some errors"))
	s.MockIRolePermission.On("GetByID", ctx, param).Return(rolepermission, errors.New("some errors"))
	actual, err := s.Repository.GetByID(ctx, param)
	s.Equal(rolepermission, actual)
	s.NotNil(err)
}

func (s *RolePermissionRepositoryTestSuite) TestGetByIDs_Success() {
	var (
		ctx   = context.Background()
		param = arguments.RolePermissionGetByIDsArgs{
			IDs: []int64{1, 2},
		}
		rolepermissions []models.RolePermission
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryContext", ctx, mock.Anything, mock.Anything).Return(s.MockIRows, nil)
	s.MockIRows.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(nil)
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockIRolePermission.On("GetByIDs", ctx, param).Return(rolepermissions, nil)
	actual, err := s.Repository.GetByIDs(ctx, param)
	s.Nil(err)
	s.Equal(rolepermissions, actual)
}

func (s *RolePermissionRepositoryTestSuite) TestGetByIDs_Fail() {
	var (
		ctx   = context.Background()
		param = arguments.RolePermissionGetByIDsArgs{
			IDs: []int64{1, 2},
		}
		rolepermissions []models.RolePermission
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryContext", ctx, mock.Anything, mock.Anything).Return(s.MockIRows, errors.New("some errors"))
	s.MockIRows.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(errors.New("some errors"))
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockIRolePermission.On("GetByIDs", ctx, param).Return(rolepermissions, errors.New("some errors"))
	actual, err := s.Repository.GetByIDs(ctx, param)
	s.Equal(rolepermissions, actual)
	s.NotNil(err)
}

func (s *RolePermissionRepositoryTestSuite) TestSetArgsToListSelectBuilder_Success() {
	var (
		params = arguments.RolePermissionListArgs{
			ID:           1,
			RoleID:       1,
			PermissionID: 1,
			CreatedBy:    "mockString",
			UpdatedBy:    "mockString",
			Page:         1,
			PageSize:     10,
		}
		selectBuilder = sq.Select("*").From("rolepermission")
	)
	offset := utils.CalculateOffsetForPage(params.Page, params.PageSize)
	expectedSelectBuilder := selectBuilder.Where(sq.Eq{"id": params.ID}).Where(sq.Eq{"role_id": params.RoleID}).Where(sq.Eq{"permission_id": params.PermissionID}).Where(sq.Eq{"created_by": params.CreatedBy}).Where(sq.Eq{"updated_by": params.UpdatedBy}).Limit(uint64(params.PageSize)).Offset(uint64(offset))
	expectSQL, expectArgs, expectErr := expectedSelectBuilder.ToSql()
	// Actual
	actual := s.Repository.setArgsToListSelectBuilder(selectBuilder, params)
	sql, args, err := actual.ToSql()
	s.Nil(err)
	s.Equal(expectSQL, sql)
	s.Equal(expectArgs, args)
	s.Equal(expectErr, err)
}

func (s *RolePermissionRepositoryTestSuite) TestList_Success() {
	var (
		ctx    = context.Background()
		params = arguments.RolePermissionListArgs{
			ID:           1,
			RoleID:       1,
			PermissionID: 1,
			CreatedBy:    "mockString",
			UpdatedBy:    "mockString",
			Page:         1,
			PageSize:     10,
		}
		rolepermissions []models.RolePermission
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
	).Return(s.MockIRows, nil)
	s.MockIRows.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(nil)
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockIRolePermission.On("List", ctx, params).Return(rolepermissions, nil)
	actual, err := s.Repository.List(ctx, params)
	s.Nil(err)
	s.Equal(rolepermissions, actual)
}

func (s *RolePermissionRepositoryTestSuite) TestList_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.RolePermissionListArgs{
			ID:           1,
			RoleID:       1,
			PermissionID: 1,
			CreatedBy:    "mockString",
			UpdatedBy:    "mockString",
			Page:         1,
			PageSize:     10,
		}
		rolepermissions []models.RolePermission
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
	).Return(s.MockIRows, errors.New("some errors"))
	s.MockIRows.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(errors.New("some errors"))
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockIRolePermission.On("List", ctx, params).Return(rolepermissions, errors.New("some errors"))
	actual, err := s.Repository.List(ctx, params)
	s.Equal(rolepermissions, actual)
	s.NotNil(err)
}

func (s *RolePermissionRepositoryTestSuite) TestSetArgsToCountSelectBuilder_Success() {
	var (
		params = arguments.RolePermissionCountArgs{
			ID:           1,
			RoleID:       1,
			PermissionID: 1,
			CreatedBy:    "mockString",
			UpdatedBy:    "mockString",
		}
		selectBuilder = sq.Select("COUNT(id)").From("rolepermission")
	)
	expectedSelectBuilder := selectBuilder.Where(sq.Eq{"id": params.ID}).Where(sq.Eq{"role_id": params.RoleID}).Where(sq.Eq{"permission_id": params.PermissionID}).Where(sq.Eq{"created_by": params.CreatedBy}).Where(sq.Eq{"updated_by": params.UpdatedBy})
	expectSQL, expectArgs, expectErr := expectedSelectBuilder.ToSql()
	// Actual
	actual := s.Repository.setArgsToCountSelectBuilder(selectBuilder, params)
	sql, args, err := actual.ToSql()
	s.Nil(err)
	s.Equal(expectSQL, sql)
	s.Equal(expectArgs, args)
	s.Equal(expectErr, err)
}

func (s *RolePermissionRepositoryTestSuite) TestCount_Success() {
	var (
		ctx    = context.Background()
		params = arguments.RolePermissionCountArgs{
			ID:           1,
			RoleID:       1,
			PermissionID: 1,
			CreatedBy:    "mockString",
			UpdatedBy:    "mockString",
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
	).Return(s.MockIRows, nil)
	s.MockIRows.On("Scan", mock.Anything).Return(nil)
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockIRolePermission.On("List", ctx, params).Return(count, nil)
	actual, err := s.Repository.Count(ctx, params)
	s.Nil(err)
	s.Equal(count, actual)
}

func (s *RolePermissionRepositoryTestSuite) TestCount_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.RolePermissionCountArgs{
			ID:           1,
			RoleID:       1,
			PermissionID: 1,
			CreatedBy:    "mockString",
			UpdatedBy:    "mockString",
		}
		count int64
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryRowContext", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(s.MockIRows, errors.New("some errors"))
	s.MockIRows.On("Scan", mock.Anything).Return(errors.New("some errors"))
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockIRolePermission.On("List", ctx, params).Return(count, errors.New("some errors"))
	actual, err := s.Repository.Count(ctx, params)
	s.Equal(count, actual)
	s.NotNil(err)
}

func (s *RolePermissionRepositoryTestSuite) TestInsert_Success() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = arguments.RolePermissionInsertArgs{
			RoleID:       1,
			PermissionID: 1,
			CreatedBy:    "mockString",
			UpdatedBy:    "mockString",
		}
		expected models.RolePermission
	)
	rolepermission := models.RolePermission{
		ID:           sampleID,
		RoleID:       params.RoleID,
		PermissionID: params.PermissionID,
		CreatedBy:    params.CreatedBy,
		UpdatedBy:    params.UpdatedBy,
	}
	//Mock Insert
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx,
		params.RoleID,
		params.PermissionID,
		params.CreatedBy,
		params.UpdatedBy,
	).Return(s.MockIResult, nil)
	s.MockIResult.On("LastInsertId").Return(sampleID, nil)
	s.MockIRolePermission.On("Insert", ctx, params).Return(rolepermission, nil)
	// Mock GetByID
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryRowContext", ctx, mock.Anything).Return(s.MockIRow)
	s.MockIRow.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(nil)
	s.MockIRolePermission.On("GetByID", ctx, params).Return(rolepermission, nil)

	actual, err := s.Repository.Insert(ctx, params)
	s.Nil(err)
	s.Equal(expected, actual)
}

func (s *RolePermissionRepositoryTestSuite) TestInsert_Fail() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = arguments.RolePermissionInsertArgs{
			RoleID:       1,
			PermissionID: 1,
			CreatedBy:    "mockString",
			UpdatedBy:    "mockString",
		}
		rolepermission models.RolePermission
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("ExecContext", ctx,
		params.RoleID,
		params.PermissionID,
		params.CreatedBy,
		params.UpdatedBy,
	).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("LastInsertId").Return(sampleID, errors.New("some errors"))
	s.MockIRolePermission.On("Insert", ctx, params).Return(rolepermission, errors.New("some errors"))
	actual, err := s.Repository.Insert(ctx, params)
	s.Equal(rolepermission, actual)
	s.NotNil(err)
}

func (s *RolePermissionRepositoryTestSuite) TestSetArgsToUpdateBuilder_Success() {
	var (
		sampleID   int64 = 1
		mockString       = "mockString"
		params           = arguments.RolePermissionUpdateArgs{
			ID:           &sampleID,
			RoleID:       &sampleID,
			PermissionID: &sampleID,
			CreatedBy:    &mockString,
			UpdatedBy:    &mockString,
		}
	)
	updateBuilder := sq.Update("rolepermission").Where(sq.Eq{"id": *params.ID})
	expectedSelectBuilder := updateBuilder.Set("role_id", *params.RoleID).Set("permission_id", *params.PermissionID).Set("created_by", *params.CreatedBy).Set("updated_by", *params.UpdatedBy)
	actual := s.Repository.setArgsToUpdateBuilder(updateBuilder, params)
	s.Equal(expectedSelectBuilder, actual)
}

func (s *RolePermissionRepositoryTestSuite) TestUpdate_Success() {
	var (
		ctx              = context.Background()
		sampleID   int64 = 1
		mockString       = "mockString"
		params           = arguments.RolePermissionUpdateArgs{
			ID:           &sampleID,
			RoleID:       &sampleID,
			PermissionID: &sampleID,
			CreatedBy:    &mockString,
			UpdatedBy:    &mockString,
		}
		expected models.RolePermission
	)
	rolepermission := models.RolePermission{
		ID:           *params.ID,
		RoleID:       *params.RoleID,
		PermissionID: *params.PermissionID,
		CreatedBy:    *params.CreatedBy,
		UpdatedBy:    *params.UpdatedBy,
	}
	// Mock Update
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx,
		*params.RoleID,
		*params.PermissionID,
		*params.CreatedBy,
		*params.UpdatedBy,
		*params.ID,
	).Return(s.MockIResult, nil)
	s.MockIResult.On("RowsAffected").Return(*params.ID, nil)
	s.MockIRolePermission.On("Update", ctx, params).Return(rolepermission, nil)
	// Mock GetByID
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryRowContext", ctx, mock.Anything).Return(s.MockIRow)
	s.MockIRow.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(nil)
	s.MockIRolePermission.On("GetByID", ctx, params).Return(rolepermission, nil)

	actual, err := s.Repository.Update(ctx, params)
	s.Nil(err)
	s.Equal(expected, actual)
}

func (s *RolePermissionRepositoryTestSuite) TestUpdate_Fail() {
	var (
		ctx              = context.Background()
		sampleID   int64 = 1
		mockString       = "mockString"
		params           = arguments.RolePermissionUpdateArgs{
			ID:           &sampleID,
			RoleID:       &sampleID,
			PermissionID: &sampleID,
			CreatedBy:    &mockString,
			UpdatedBy:    &mockString,
		}
		rolepermission models.RolePermission
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("ExecContext", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("RowsAffected").Return(sampleID, errors.New("some errors"))
	s.MockIRolePermission.On("Update", ctx, params).Return(rolepermission, errors.New("some errors"))
	actual, err := s.Repository.Update(ctx, params)
	s.Equal(rolepermission, actual)
	s.NotNil(err)
}

func (s *RolePermissionRepositoryTestSuite) TestDelete_Success() {
	var (
		ctx   = context.Background()
		param = arguments.RolePermissionDeleteArgs{
			ID: 1,
		}
		rowEffected int64 = 1
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx, param.ID).Return(s.MockIResult, nil)
	s.MockIResult.On("RowsAffected").Return(param.ID, nil)
	s.MockIRolePermission.On("Delete", ctx, param).Return(rowEffected, nil)
	actual, err := s.Repository.Delete(ctx, param)
	s.Nil(err)
	s.Equal(param.ID, actual)
}

func (s *RolePermissionRepositoryTestSuite) TestDelete_Fail() {
	var (
		ctx   = context.Background()
		param = arguments.RolePermissionDeleteArgs{
			ID: 1,
		}
		rowEffected int64
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("ExecContext", ctx, param.ID).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("RowsAffected").Return(param.ID, errors.New("some errors"))
	s.MockIRolePermission.On("Delete", ctx, param).Return(rowEffected, errors.New("some errors"))
	actual, err := s.Repository.Delete(ctx, param)
	s.Equal(rowEffected, actual)
	s.NotNil(err)
}
