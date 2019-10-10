// @generated
package role

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/mock"
	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
)

func (s *RoleRepositoryTestSuite) TestGetByID_Success() {
	var (
		ctx   = context.Background()
		param = arguments.RoleGetByID{
			ID: 1,
		}
		role models.Role
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
	actual, err := s.Repository.GetByID(ctx, param)
	s.Nil(err)
	s.Equal(role, actual)
}

func (s *RoleRepositoryTestSuite) TestGetByID_Fail() {
	var (
		ctx   = context.Background()
		param = arguments.RoleGetByID{
			ID: 1,
		}
		role models.Role
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
	actual, err := s.Repository.GetByID(ctx, param)
	s.Equal(role, actual)
	s.NotNil(err)
}

func (s *RoleRepositoryTestSuite) TestGetByID_Fail1() {
	var (
		ctx   = context.Background()
		param = arguments.RoleGetByID{
			ID: 1,
		}
		role models.Role
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
	).Return(errors.New("some errors"))
	actual, err := s.Repository.GetByID(ctx, param)
	s.Equal(role, actual)
	s.NotNil(err)
}

func (s *RoleRepositoryTestSuite) TestGetByIDs_Success() {
	var (
		ctx   = context.Background()
		param = arguments.RoleGetByIDs{
			IDs: []int64{1, 2},
		}
		roles []models.Role
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
	actual, err := s.Repository.GetByIDs(ctx, param)
	s.Nil(err)
	s.Equal(roles, actual)
}

func (s *RoleRepositoryTestSuite) TestGetByIDs_Fail() {
	var (
		ctx   = context.Background()
		param = arguments.RoleGetByIDs{
			IDs: []int64{1, 2},
		}
		roles []models.Role
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
	actual, err := s.Repository.GetByIDs(ctx, param)
	s.Equal(roles, actual)
	s.NotNil(err)
}

func (s *RoleRepositoryTestSuite) TestGetByIDs_Fail1() {
	var (
		ctx   = context.Background()
		param = arguments.RoleGetByIDs{
			IDs: []int64{1, 2},
		}
		roles []models.Role
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
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
	actual, err := s.Repository.GetByIDs(ctx, param)
	s.Equal(roles, actual)
	s.NotNil(err)
}

func (s *RoleRepositoryTestSuite) TestGetByIDs_Fail2() {
	var (
		ctx   = context.Background()
		param = arguments.RoleGetByIDs{
			IDs: []int64{1, 2},
		}
		roles []models.Role
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
	).Return(errors.New("some errors"))
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(true)
	actual, err := s.Repository.GetByIDs(ctx, param)
	s.Equal(roles, actual)
	s.NotNil(err)
}

func (s *RoleRepositoryTestSuite) TestSetArgsToListSelectBuilder_Success() {
	var (
		ctx    = context.Background()
		params = arguments.RoleList{
			ID:        1,
			Name:      "mockString",
			CompanyID: 1,
			Status:    "active",
			CreatedBy: "mockString",
			UpdatedBy: "mockString",
			LastID:    1,
			PageSize:  10,
		}
		selectBuilder = sq.Select("*").From("role")
	)
	expectedSelectBuilder := selectBuilder.Where(sq.Eq{"id": params.ID}).Where(sq.Eq{"name": params.Name}).Where(sq.Eq{"company_id": params.CompanyID}).Where(sq.Eq{"status": params.Status}).Where(sq.Eq{"created_by": params.CreatedBy}).Where(sq.Eq{"updated_by": params.UpdatedBy}).Limit(uint64(params.PageSize)).Where(sq.Gt{"id": params.LastID})
	expectSQL, expectArgs, expectErr := expectedSelectBuilder.ToSql()
	// Actual
	actual := s.Repository.setArgsToListSelectBuilder(ctx, selectBuilder, params)
	sql, args, err := actual.ToSql()
	s.Nil(err)
	s.Equal(expectSQL, sql)
	s.Equal(expectArgs, args)
	s.Equal(expectErr, err)
}

func (s *RoleRepositoryTestSuite) TestList_Success() {
	var (
		ctx    = context.Background()
		params = arguments.RoleList{
			ID:        1,
			Name:      "mockString",
			CompanyID: 1,
			Status:    "active",
			CreatedBy: "mockString",
			UpdatedBy: "mockString",
			PageSize:  10,
		}
		roles []models.Role
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryContext", ctx,
		params.ID,
		params.Name,
		params.CompanyID,
		params.Status,
		params.CreatedBy,
		params.UpdatedBy,
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
	actual, err := s.Repository.List(ctx, params)
	s.Nil(err)
	s.Equal(roles, actual)
}

func (s *RoleRepositoryTestSuite) TestList_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.RoleList{
			ID:        1,
			Name:      "mockString",
			CompanyID: 1,
			Status:    "active",
			CreatedBy: "mockString",
			UpdatedBy: "mockString",
			PageSize:  10,
		}
		roles []models.Role
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryContext", ctx,
		params.ID,
		params.Name,
		params.CompanyID,
		params.Status,
		params.CreatedBy,
		params.UpdatedBy,
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
	actual, err := s.Repository.List(ctx, params)
	s.Equal(roles, actual)
	s.NotNil(err)
}

func (s *RoleRepositoryTestSuite) TestList_Fail1() {
	var (
		ctx    = context.Background()
		params = arguments.RoleList{
			ID:        1,
			Name:      "mockString",
			CompanyID: 1,
			Status:    "active",
			CreatedBy: "mockString",
			UpdatedBy: "mockString",
			PageSize:  10,
		}
		roles []models.Role
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryContext", ctx,
		params.ID,
		params.Name,
		params.CompanyID,
		params.Status,
		params.CreatedBy,
		params.UpdatedBy,
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
	actual, err := s.Repository.List(ctx, params)
	s.Equal(roles, actual)
	s.NotNil(err)
}

func (s *RoleRepositoryTestSuite) TestList_Fail2() {
	var (
		ctx    = context.Background()
		params = arguments.RoleList{
			ID:        1,
			Name:      "mockString",
			CompanyID: 1,
			Status:    "active",
			CreatedBy: "mockString",
			UpdatedBy: "mockString",
		}
		roles []models.Role
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryContext", ctx,
		params.ID,
		params.Name,
		params.CompanyID,
		params.Status,
		params.CreatedBy,
		params.UpdatedBy,
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
	).Return(errors.New("some errors"))
	s.MockIRows.On("Close").Return(nil)
	s.MockIRows.On("Next").Return(true)
	actual, err := s.Repository.List(ctx, params)
	s.Equal(roles, actual)
	s.NotNil(err)
}

func (s *RoleRepositoryTestSuite) TestSetArgsToCountSelectBuilder_Success() {
	var (
		ctx    = context.Background()
		params = arguments.RoleCount{
			ID:        1,
			Name:      "mockString",
			CompanyID: 1,
			Status:    "active",
			CreatedBy: "mockString",
			UpdatedBy: "mockString",
		}
		selectBuilder = sq.Select("COUNT(id)").From("role")
	)
	expectedSelectBuilder := selectBuilder.Where(sq.Eq{"id": params.ID}).Where(sq.Eq{"name": params.Name}).Where(sq.Eq{"company_id": params.CompanyID}).Where(sq.Eq{"status": params.Status}).Where(sq.Eq{"created_by": params.CreatedBy}).Where(sq.Eq{"updated_by": params.UpdatedBy})
	expectSQL, expectArgs, expectErr := expectedSelectBuilder.ToSql()
	// Actual
	actual := s.Repository.setArgsToCountSelectBuilder(ctx, selectBuilder, params)
	sql, args, err := actual.ToSql()
	s.Nil(err)
	s.Equal(expectSQL, sql)
	s.Equal(expectArgs, args)
	s.Equal(expectErr, err)
}

func (s *RoleRepositoryTestSuite) TestCount_Success() {
	var (
		ctx    = context.Background()
		params = arguments.RoleCount{
			ID:        1,
			Name:      "mockString",
			CompanyID: 1,
			Status:    "active",
			CreatedBy: "mockString",
			UpdatedBy: "mockString",
		}
		count int64
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryRowContext", ctx,
		params.ID,
		params.Name,
		params.CompanyID,
		params.Status,
		params.CreatedBy,
		params.UpdatedBy,
	).Return(s.MockIRows, nil)
	s.MockIRows.On("Scan", mock.Anything).Return(nil)
	actual, err := s.Repository.Count(ctx, params)
	s.Nil(err)
	s.Equal(count, actual)
}

func (s *RoleRepositoryTestSuite) TestCount_Fail() {
	var (
		ctx    = context.Background()
		params = arguments.RoleCount{
			ID:        1,
			Name:      "mockString",
			CompanyID: 1,
			Status:    "active",
			CreatedBy: "mockString",
			UpdatedBy: "mockString",
		}
		count int64
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryRowContext", ctx,
		params.ID,
		params.Name,
		params.CompanyID,
		params.Status,
		params.CreatedBy,
		params.UpdatedBy,
	).Return(s.MockIRows, errors.New("some errors"))
	s.MockIRows.On("Scan", mock.Anything).Return(errors.New("some errors"))
	actual, err := s.Repository.Count(ctx, params)
	s.Equal(count, actual)
	s.NotNil(err)
}

func (s *RoleRepositoryTestSuite) TestCount_Fail1() {
	var (
		ctx    = context.Background()
		params = arguments.RoleCount{
			ID:        1,
			Name:      "mockString",
			CompanyID: 1,
			Status:    "active",
			CreatedBy: "mockString",
			UpdatedBy: "mockString",
		}
		count int64
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryRowContext", ctx,
		params.ID,
		params.Name,
		params.CompanyID,
		params.Status,
		params.CreatedBy,
		params.UpdatedBy,
	).Return(s.MockIRows, errors.New("some errors"))
	s.MockIRows.On("Scan", mock.Anything).Return(errors.New("some errors"))
	actual, err := s.Repository.Count(ctx, params)
	s.Equal(count, actual)
	s.NotNil(err)
}

func (s *RoleRepositoryTestSuite) TestCount_Fail2() {
	var (
		ctx    = context.Background()
		params = arguments.RoleCount{
			ID:        1,
			Name:      "mockString",
			CompanyID: 1,
			Status:    "active",
			CreatedBy: "mockString",
			UpdatedBy: "mockString",
		}
		count int64
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryRowContext", ctx,
		params.ID,
		params.Name,
		params.CompanyID,
		params.Status,
		params.CreatedBy,
		params.UpdatedBy,
	).Return(s.MockIRows, nil)
	s.MockIRows.On("Scan", mock.Anything).Return(errors.New("some errors"))
	actual, err := s.Repository.Count(ctx, params)
	s.Equal(count, actual)
	s.NotNil(err)
}

func (s *RoleRepositoryTestSuite) TestInsert_Success() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = arguments.RoleInsert{
			Name:      "mockString",
			CompanyID: 1,
			Status:    "active",
			CreatedBy: "mockString",
			UpdatedBy: "mockString",
		}
		expected models.Role
	)
	//Mock Insert
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx,
		params.Name,
		params.CompanyID,
		params.Status,
		params.CreatedBy,
		params.UpdatedBy,
	).Return(s.MockIResult, nil)
	s.MockIResult.On("LastInsertId").Return(sampleID, nil)
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

	actual, err := s.Repository.Insert(ctx, params)
	s.Nil(err)
	s.Equal(expected, actual)
}

func (s *RoleRepositoryTestSuite) TestInsert_Fail() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = arguments.RoleInsert{
			Name:      "mockString",
			CompanyID: 1,
			Status:    "active",
			CreatedBy: "mockString",
			UpdatedBy: "mockString",
		}
		role models.Role
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("ExecContext", ctx,
		params.Name,
		params.CompanyID,
		params.Status,
		params.CreatedBy,
		params.UpdatedBy,
	).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("LastInsertId").Return(sampleID, errors.New("some errors"))
	actual, err := s.Repository.Insert(ctx, params)
	s.Equal(role, actual)
	s.NotNil(err)
}

func (s *RoleRepositoryTestSuite) TestInsert_Fail1() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = arguments.RoleInsert{
			Name:      "mockString",
			CompanyID: 1,
			Status:    "active",
			CreatedBy: "mockString",
			UpdatedBy: "mockString",
		}
		role models.Role
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx,
		params.Name,
		params.CompanyID,
		params.Status,
		params.CreatedBy,
		params.UpdatedBy,
	).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("LastInsertId").Return(sampleID, errors.New("some errors"))
	actual, err := s.Repository.Insert(ctx, params)
	s.Equal(role, actual)
	s.NotNil(err)
}

func (s *RoleRepositoryTestSuite) TestInsert_Fail2() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = arguments.RoleInsert{
			Name:      "mockString",
			CompanyID: 1,
			Status:    "active",
			CreatedBy: "mockString",
			UpdatedBy: "mockString",
		}
		role models.Role
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx,
		params.Name,
		params.CompanyID,
		params.Status,
		params.CreatedBy,
		params.UpdatedBy,
	).Return(s.MockIResult, nil)
	s.MockIResult.On("LastInsertId").Return(sampleID, errors.New("some errors"))
	actual, err := s.Repository.Insert(ctx, params)
	s.Equal(role, actual)
	s.NotNil(err)
}

func (s *RoleRepositoryTestSuite) TestInsert_Fail3() {
	var (
		ctx            = context.Background()
		sampleID int64 = 1
		params         = arguments.RoleInsert{
			Name:      "mockString",
			CompanyID: 1,
			Status:    "active",
			CreatedBy: "mockString",
			UpdatedBy: "mockString",
		}
		expected models.Role
	)
	//Mock Insert
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx,
		params.Name,
		params.CompanyID,
		params.Status,
		params.CreatedBy,
		params.UpdatedBy,
	).Return(s.MockIResult, nil)
	s.MockIResult.On("LastInsertId").Return(sampleID, nil)
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
	).Return(errors.New("some errors"))

	actual, err := s.Repository.Insert(ctx, params)
	s.Equal(expected, actual)
	s.NotNil(err)
}

func (s *RoleRepositoryTestSuite) TestSetArgsToUpdateBuilder_Success() {
	var (
		ctx              = context.Background()
		sampleID   int64 = 1
		mockString       = "mockString"
		status           = "active"
		params           = arguments.RoleUpdate{
			ID:        &sampleID,
			Name:      &mockString,
			CompanyID: &sampleID,
			Status:    &status,
			CreatedBy: &mockString,
			UpdatedBy: &mockString,
		}
	)
	updateBuilder := sq.Update("role").Where(sq.Eq{"id": *params.ID})
	expectedSelectBuilder := updateBuilder.Set("name", *params.Name).Set("company_id", *params.CompanyID).Set("status", *params.Status).Set("created_by", *params.CreatedBy).Set("updated_by", *params.UpdatedBy)
	actual := s.Repository.setArgsToUpdateBuilder(ctx, updateBuilder, params)
	s.Equal(expectedSelectBuilder, actual)
}

func (s *RoleRepositoryTestSuite) TestUpdate_Success() {
	var (
		ctx               = context.Background()
		sampleID    int64 = 1
		rowEffected int64 = 1
		mockString        = "mockString"
		status            = "active"
		params            = arguments.RoleUpdate{
			ID:        &sampleID,
			Name:      &mockString,
			CompanyID: &sampleID,
			Status:    &status,
			CreatedBy: &mockString,
			UpdatedBy: &mockString,
		}
		expected models.Role
	)
	// Mock Update
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx,
		*params.Name,
		*params.CompanyID,
		*params.Status,
		*params.CreatedBy,
		*params.UpdatedBy,
		*params.ID,
	).Return(s.MockIResult, nil)
	s.MockIResult.On("RowsAffected").Return(rowEffected, nil)
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

	actual, err := s.Repository.Update(ctx, params)
	s.Nil(err)
	s.Equal(expected, actual)
}

func (s *RoleRepositoryTestSuite) TestUpdate_Fail() {
	var (
		ctx               = context.Background()
		sampleID    int64 = 1
		rowEffected int64
		mockString  = "mockString"
		status      = "active"
		params      = arguments.RoleUpdate{
			ID:        &sampleID,
			Name:      &mockString,
			CompanyID: &sampleID,
			Status:    &status,
			CreatedBy: &mockString,
			UpdatedBy: &mockString,
		}
		role models.Role
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("ExecContext", ctx,
		*params.Name,
		*params.CompanyID,
		*params.Status,
		*params.CreatedBy,
		*params.UpdatedBy,
		*params.ID,
	).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("RowsAffected").Return(rowEffected, errors.New("some errors"))
	actual, err := s.Repository.Update(ctx, params)
	s.Equal(role, actual)
	s.NotNil(err)
}

func (s *RoleRepositoryTestSuite) TestUpdate_Fail1() {
	var (
		ctx               = context.Background()
		sampleID    int64 = 1
		rowEffected int64
		mockString  = "mockString"
		status      = "active"
		params      = arguments.RoleUpdate{
			ID:        &sampleID,
			Name:      &mockString,
			CompanyID: &sampleID,
			Status:    &status,
			CreatedBy: &mockString,
			UpdatedBy: &mockString,
		}
		role models.Role
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx,
		*params.Name,
		*params.CompanyID,
		*params.Status,
		*params.CreatedBy,
		*params.UpdatedBy,
		*params.ID,
	).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("RowsAffected").Return(rowEffected, errors.New("some errors"))
	actual, err := s.Repository.Update(ctx, params)
	s.Equal(role, actual)
	s.NotNil(err)
}

func (s *RoleRepositoryTestSuite) TestUpdate_Fail2() {
	var (
		ctx               = context.Background()
		sampleID    int64 = 1
		rowEffected int64
		mockString  = "mockString"
		status      = "active"
		params      = arguments.RoleUpdate{
			ID:        &sampleID,
			Name:      &mockString,
			CompanyID: &sampleID,
			Status:    &status,
			CreatedBy: &mockString,
			UpdatedBy: &mockString,
		}
		role models.Role
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx,
		*params.Name,
		*params.CompanyID,
		*params.Status,
		*params.CreatedBy,
		*params.UpdatedBy,
		*params.ID,
	).Return(s.MockIResult, nil)
	s.MockIResult.On("RowsAffected").Return(rowEffected, errors.New("some errors"))
	actual, err := s.Repository.Update(ctx, params)
	s.Equal(role, actual)
	s.NotNil(err)
}

func (s *RoleRepositoryTestSuite) TestUpdate_Fail3() {
	var (
		ctx               = context.Background()
		sampleID    int64 = 1
		rowEffected int64
		mockString  = "mockString"
		status      = "active"
		params      = arguments.RoleUpdate{
			ID:        &sampleID,
			Name:      &mockString,
			CompanyID: &sampleID,
			Status:    &status,
			CreatedBy: &mockString,
			UpdatedBy: &mockString,
		}
		role models.Role
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx,
		*params.Name,
		*params.CompanyID,
		*params.Status,
		*params.CreatedBy,
		*params.UpdatedBy,
		*params.ID,
	).Return(s.MockIResult, nil)
	s.MockIResult.On("RowsAffected").Return(rowEffected, nil)
	actual, err := s.Repository.Update(ctx, params)
	s.Equal(role, actual)
	s.NotNil(err)
}

func (s *RoleRepositoryTestSuite) TestUpdate_Fail4() {
	var (
		ctx              = context.Background()
		sampleID   int64 = 1
		mockString       = "mockString"
		status           = "active"
		params           = arguments.RoleUpdate{
			ID:        &sampleID,
			Name:      &mockString,
			CompanyID: &sampleID,
			Status:    &status,
			CreatedBy: &mockString,
			UpdatedBy: &mockString,
		}
		role models.Role
	)
	// Mock Update
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx,
		*params.Name,
		*params.CompanyID,
		*params.Status,
		*params.CreatedBy,
		*params.UpdatedBy,
		*params.ID,
	).Return(s.MockIResult, nil)
	s.MockIResult.On("RowsAffected").Return(*params.ID, nil)
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
	).Return(errors.New("some errors"))

	actual, err := s.Repository.Update(ctx, params)
	s.Equal(role, actual)
	s.NotNil(err)
}

func (s *RoleRepositoryTestSuite) TestDelete_Success() {
	var (
		ctx   = context.Background()
		param = arguments.RoleDelete{
			ID: 1,
		}
		rowEffected int64 = 1
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx, param.ID).Return(s.MockIResult, nil)
	s.MockIResult.On("RowsAffected").Return(rowEffected, nil)
	actual, err := s.Repository.Delete(ctx, param)
	s.Nil(err)
	s.Equal(param.ID, actual)
}

func (s *RoleRepositoryTestSuite) TestDelete_Fail() {
	var (
		ctx   = context.Background()
		param = arguments.RoleDelete{
			ID: 1,
		}
		rowEffected int64
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("ExecContext", ctx, param.ID).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("RowsAffected").Return(rowEffected, errors.New("some errors"))
	actual, err := s.Repository.Delete(ctx, param)
	s.Equal(rowEffected, actual)
	s.NotNil(err)
}

func (s *RoleRepositoryTestSuite) TestDelete_Fail1() {
	var (
		ctx   = context.Background()
		param = arguments.RoleDelete{
			ID: 1,
		}
		rowEffected int64
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx, param.ID).Return(s.MockIResult, errors.New("some errors"))
	s.MockIResult.On("RowsAffected").Return(rowEffected, errors.New("some errors"))
	actual, err := s.Repository.Delete(ctx, param)
	s.Equal(rowEffected, actual)
	s.NotNil(err)
}

func (s *RoleRepositoryTestSuite) TestDelete_Fail2() {
	var (
		ctx   = context.Background()
		param = arguments.RoleDelete{
			ID: 1,
		}
		rowEffected int64
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx, param.ID).Return(s.MockIResult, nil)
	s.MockIResult.On("RowsAffected").Return(rowEffected, errors.New("some errors"))
	actual, err := s.Repository.Delete(ctx, param)
	s.Equal(rowEffected, actual)
	s.NotNil(err)
}

func (s *RoleRepositoryTestSuite) TestDelete_Fail3() {
	var (
		ctx   = context.Background()
		param = arguments.RoleDelete{
			ID: 1,
		}
		rowEffected int64
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("ExecContext", ctx, param.ID).Return(s.MockIResult, nil)
	s.MockIResult.On("RowsAffected").Return(rowEffected, nil)
	actual, err := s.Repository.Delete(ctx, param)
	s.Equal(rowEffected, actual)
	s.NotNil(err)
}
