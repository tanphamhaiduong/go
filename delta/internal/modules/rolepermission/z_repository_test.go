// @generated
package rolepermission

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

func (s *RolePermissionRepositoryTestSuite) TestGetByID_Success() {
	var (
		ctx    = context.Background()
		params = arguments.RolePermissionGetByIDArgs{
			ID: 1,
		}
		rolepermission models.RolePermission
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
	).Return(nil)
	s.MockIRolePermission.On("GetByID", ctx, params).Return(rolepermission, nil)
	actual, err := s.Repository.GetByID(ctx, params)
	s.Nil(err)
	s.Equal(rolepermission, actual)
}

func (s *RolePermissionRepositoryTestSuite) TestGetByID_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.RolePermissionGetByIDArgs{
			ID: 1,
		}
		rolepermission models.RolePermission
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryRowContext", ctx, mock.Anything).Return(s.MockIRow)
	s.MockIRow.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("some errors"))
	s.MockIRolePermission.On("GetByID", ctx, params).Return(rolepermission, errors.New("some errors"))
	actual, err := s.Repository.GetByID(ctx, params)
	s.Equal(rolepermission, actual)
	s.NotNil(err)
}

func (s *RolePermissionRepositoryTestSuite) TestGetByIDs_Success() {
	var (
		ctx    = context.Background()
		params = arguments.RolePermissionGetByIDsArgs{
			IDs: []int64{1, 2},
		}
		rolepermissions []models.RolePermission
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryContext", ctx, mock.Anything, mock.Anything).Return(s.MockIRows, nil)
	s.MockIRows.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockIRolePermission.On("GetByIDs", ctx, params).Return(rolepermissions, nil)
	actual, err := s.Repository.GetByIDs(ctx, params)
	s.Nil(err)
	s.Equal(rolepermissions, actual)
}

func (s *RolePermissionRepositoryTestSuite) TestGetByIDs_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.RolePermissionGetByIDsArgs{
			IDs: []int64{1, 2},
		}
		rolepermissions []models.RolePermission
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryContext", ctx, mock.Anything, mock.Anything).Return(s.MockIRows, errors.New("some errors"))
	s.MockIRows.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("some errors"))
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(false)
	s.MockIRolePermission.On("GetByIDs", ctx, params).Return(rolepermissions, errors.New("some errors"))
	actual, err := s.Repository.GetByIDs(ctx, params)
	s.Equal(rolepermissions, actual)
	s.NotNil(err)
}

func (s *RolePermissionRepositoryTestSuite) TestSetArgsToListSelectBuilder_Success() {
	var (
		params        = arguments.RolePermissionListArgs{}
		selectBuilder = sq.Select("*").From("rolepermission")
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	offset := utils.CalculateOffsetForPage(params.Page, params.PageSize)
	expectedSelectBuilder := selectBuilder.Where(sq.Eq{"id": params.ID}).Where(sq.Eq{"role_id": params.RoleID}).Where(sq.Eq{"permission_id": params.PermissionID}).Where(sq.Eq{"status": params.Status}).Where(sq.Eq{"created_by": params.CreatedBy}).Where(sq.Eq{"updated_by": params.UpdatedBy}).Limit(uint64(params.PageSize)).Offset(uint64(offset))
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
			Page:     1,
			PageSize: 10,
		}
		rolepermissions []models.RolePermission
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
	s.MockIRolePermission.On("List", ctx, params).Return(rolepermissions, nil)
	actual, err := s.Repository.List(ctx, params)
	s.Nil(err)
	s.Equal(rolepermissions, actual)
}

func (s *RolePermissionRepositoryTestSuite) TestList_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.RolePermissionListArgs{
			Page:     1,
			PageSize: 10,
		}
		rolepermissions []models.RolePermission
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
	s.MockIRolePermission.On("List", ctx, params).Return(rolepermissions, errors.New("some errors"))
	actual, err := s.Repository.List(ctx, params)
	s.Equal(rolepermissions, actual)
	s.NotNil(err)
}

func (s *RolePermissionRepositoryTestSuite) TestSetArgsToCountSelectBuilder_Success() {
	var (
		params        = arguments.RolePermissionCountArgs{}
		selectBuilder = sq.Select("COUNT(id)").From("rolepermission")
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	expectedSelectBuilder := selectBuilder.Where(sq.Eq{"id": params.ID}).Where(sq.Eq{"role_id": params.RoleID}).Where(sq.Eq{"permission_id": params.PermissionID}).Where(sq.Eq{"status": params.Status}).Where(sq.Eq{"created_by": params.CreatedBy}).Where(sq.Eq{"updated_by": params.UpdatedBy})
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
		params = arguments.RolePermissionCountArgs{}
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
		params = arguments.RolePermissionCountArgs{}
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
	s.MockIRolePermission.On("List", ctx, params).Return(count, errors.New("some errors"))
	actual, err := s.Repository.Count(ctx, params)
	s.Equal(count, actual)
	s.NotNil(err)
}

func (s *RolePermissionRepositoryTestSuite) TestInsert_Success() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = arguments.RolePermissionInsertArgs{}
		expected models.RolePermission
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	rolepermission := models.RolePermission{
		ID:           sampleID,
		RoleID:       params.RoleID,
		PermissionID: params.PermissionID,
		Status:       params.Status,
		CreatedBy:    params.CreatedBy,
		UpdatedBy:    params.UpdatedBy,
	}
	//Mock Insert
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx,
		params.RoleID,
		params.PermissionID,
		params.Status,
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
			RoleID:       0,
			PermissionID: 0,
			Status:       "mockString",
			CreatedBy:    "mockString",
			UpdatedBy:    "mockString",
		}
		rolepermission models.RolePermission
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("ExecContext", ctx,
		params.RoleID,
		params.PermissionID,
		params.Status,
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
		params = arguments.RolePermissionUpdateArgs{}
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	updateBuilder := sq.Update("rolepermission").Where(sq.Eq{"id": *params.ID})
	expectedSelectBuilder := updateBuilder.Set("role_id", *params.RoleID).Set("permission_id", *params.PermissionID).Set("status", *params.Status).Set("created_by", *params.CreatedBy).Set("updated_by", *params.UpdatedBy)
	actual := s.Repository.setArgsToUpdateBuilder(updateBuilder, params)
	s.Equal(expectedSelectBuilder, actual)
}

func (s *RolePermissionRepositoryTestSuite) TestUpdate_Success() {
	var (
		ctx      = context.Background()
		params   = arguments.RolePermissionUpdateArgs{}
		expected models.RolePermission
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	rolepermission := models.RolePermission{
		ID:           *params.ID,
		RoleID:       *params.RoleID,
		PermissionID: *params.PermissionID,
		Status:       *params.Status,
		CreatedBy:    *params.CreatedBy,
		UpdatedBy:    *params.UpdatedBy,
	}
	// Mock Update
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx,
		*params.RoleID,
		*params.PermissionID,
		*params.Status,
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
			Status:       &mockString,
			CreatedBy:    &mockString,
			UpdatedBy:    &mockString,
		}
		rolepermission models.RolePermission
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
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
		ctx               = context.Background()
		params            = arguments.RolePermissionDeleteArgs{}
		rowEffected int64 = 1
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx, params.ID).Return(s.MockIResult, nil)
	s.MockIResult.On("RowsAffected").Return(params.ID, nil)
	s.MockIRolePermission.On("Delete", ctx, params).Return(rowEffected, nil)
	actual, err := s.Repository.Delete(ctx, params)
	s.Nil(err)
	s.Equal(params.ID, actual)
}

func (s *RolePermissionRepositoryTestSuite) TestDelete_Fail() {
	var (
		ctx         = context.Background()
		params      = arguments.RolePermissionDeleteArgs{}
		rowEffected int64
	)
	err := faker.FakeData(&params)
	if err != nil {
		log.Fatal(err)
	}
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("ExecContext", ctx, params.ID).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("RowsAffected").Return(params.ID, errors.New("some errors"))
	s.MockIRolePermission.On("Delete", ctx, params).Return(rowEffected, errors.New("some errors"))
	actual, err := s.Repository.Delete(ctx, params)
	s.Equal(rowEffected, actual)
	s.NotNil(err)
}
