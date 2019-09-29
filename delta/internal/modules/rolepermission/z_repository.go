// @generated
package rolepermission

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/tanphamhaiduong/go/common/logger"
	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/database"
	"github.com/tanphamhaiduong/go/delta/internal/models"
	"github.com/tanphamhaiduong/go/delta/internal/utils"
)

//scanRolePermission
func (r *RepositoryImpl) scanRolePermission(row database.IRow, rolepermission *models.RolePermission) error {
	err := row.Scan(
		&rolepermission.ID,
		&rolepermission.RoleID,
		&rolepermission.PermissionID,
		&rolepermission.CreatedBy,
		&rolepermission.UpdatedBy,
	)
	return err
}

// GetByID ...
func (r *RepositoryImpl) GetByID(ctx context.Context, params arguments.RolePermissionGetByID) (models.RolePermission, error) {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Repository GetByID of rolepermission")
	var (
		rolepermission models.RolePermission
		selectBuilder  = sq.Select(
			"id",
			"role_id",
			"permission_id",
			"created_by",
			"updated_by",
		).From("role_permission").Where(sq.Eq{"id": params.ID})
	)
	sql, args, err := selectBuilder.ToSql()
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository GetByID build sql string of rolepermission")
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository GetByID selectBuilder error of rolepermission")
		return rolepermission, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository GetByID PrepareContext error of rolepermission")
		return rolepermission, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	err = r.scanRolePermission(row, &rolepermission)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository GetByID Scan error of rolepermission")
		return rolepermission, err
	}
	return rolepermission, nil
}

// GetByIDs ...
func (r *RepositoryImpl) GetByIDs(ctx context.Context, params arguments.RolePermissionGetByIDs) ([]models.RolePermission, error) {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Repository GetByIDs of rolepermission")
	var (
		rolepermissions []models.RolePermission
		selectBuilder   = sq.Select(
			"id",
			"role_id",
			"permission_id",
			"created_by",
			"updated_by",
		).From("role_permission").Where(sq.Eq{"id": params.IDs})
	)
	sql, args, err := selectBuilder.ToSql()
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository GetByIDs build sql string of rolepermission")
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository GetByIDs selectBuilder error of rolepermission")
		return rolepermissions, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository GetByIDs PrepareContext error of rolepermission")
		return rolepermissions, err
	}
	rows, err := stmt.QueryContext(ctx, args...)
	defer rows.Close()
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository GetByIDs QueryContext error of rolepermission")
		return rolepermissions, err
	}
	for rows.Next() {
		rolepermission := models.RolePermission{}
		err := r.scanRolePermission(rows, &rolepermission)
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository GetByIDs Scan error of rolepermission")
			return rolepermissions, err
		}
		rolepermissions = append(rolepermissions, rolepermission)
	}
	return rolepermissions, nil
}

// setArgsToListSelectBuilder ...
func (r *RepositoryImpl) setArgsToListSelectBuilder(ctx context.Context, selectBuilder sq.SelectBuilder, params arguments.RolePermissionList) sq.SelectBuilder {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Repository setArgsToListSelectBuilder of rolepermission")
	if params.ID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"id": params.ID})
	}
	if params.RoleID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"role_id": params.RoleID})
	}
	if params.PermissionID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"permission_id": params.PermissionID})
	}
	if params.CreatedBy != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"created_by": params.CreatedBy})
	}
	if params.UpdatedBy != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"updated_by": params.UpdatedBy})
	}
	if params.PageSize != 0 {
		selectBuilder = selectBuilder.Limit(uint64(params.PageSize))
	}
	selectBuilder = selectBuilder.Where(sq.Gt{"id": params.LastID})
	return selectBuilder
}

// List ...
func (r *RepositoryImpl) List(ctx context.Context, params arguments.RolePermissionList) ([]models.RolePermission, error) {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Repository List of rolepermission")
	var (
		rolepermissions []models.RolePermission
		selectBuilder   = sq.Select(
			"id",
			"role_id",
			"permission_id",
			"created_by",
			"updated_by",
		).From("role_permission")
	)
	selectBuilderWithArgs := r.setArgsToListSelectBuilder(ctx, selectBuilder, params)
	sql, args, err := selectBuilderWithArgs.ToSql()
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository List build sql string of rolepermission")
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository List selectBuilderWithArgs error of rolepermission")
		return rolepermissions, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository List PrepareContext error of rolepermission")
		return rolepermissions, err
	}
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository List QueryContext error of rolepermission")
		return rolepermissions, err
	}
	defer rows.Close()
	for rows.Next() {
		rolepermission := models.RolePermission{}
		err := r.scanRolePermission(rows, &rolepermission)
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository List Scan error of rolepermission")
			return rolepermissions, err
		}
		rolepermissions = append(rolepermissions, rolepermission)
	}
	return rolepermissions, nil
}

// setArgsToCountSelectBuilder ...
func (r *RepositoryImpl) setArgsToCountSelectBuilder(ctx context.Context, selectBuilder sq.SelectBuilder, params arguments.RolePermissionCount) sq.SelectBuilder {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Repository setArgsToCountSelectBuilder of rolepermission")
	if params.ID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"id": params.ID})
	}
	if params.RoleID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"role_id": params.RoleID})
	}
	if params.PermissionID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"permission_id": params.PermissionID})
	}
	if params.CreatedBy != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"created_by": params.CreatedBy})
	}
	if params.UpdatedBy != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"updated_by": params.UpdatedBy})
	}
	return selectBuilder
}

// Count ...
func (r *RepositoryImpl) Count(ctx context.Context, params arguments.RolePermissionCount) (int64, error) {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Repository Count of rolepermission")
	var (
		count         int64
		selectBuilder = sq.Select("COUNT(id)").From("role_permission")
	)
	selectBuilderWithArgs := r.setArgsToCountSelectBuilder(ctx, selectBuilder, params)
	sql, args, err := selectBuilderWithArgs.ToSql()
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository Count build sql string of rolepermission")
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository Count selectBuilderWithArgs error of rolepermission")
		return count, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository Count PrepareContext error of rolepermission")
		return count, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	err = row.Scan(&count)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository Count Scan error of rolepermission")
		return count, err
	}
	return count, nil
}

// Insert ...
func (r *RepositoryImpl) Insert(ctx context.Context, params arguments.RolePermissionInsert) (models.RolePermission, error) {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Repository Insert of rolepermission")
	var (
		rolepermission models.RolePermission
		insertBuilder  = sq.Insert("role_permission").Columns(
			"role_id",
			"permission_id",
			"created_by",
			"updated_by",
		).Values(
			params.RoleID,
			params.PermissionID,
			params.CreatedBy,
			params.UpdatedBy,
		)
	)
	sql, args, err := insertBuilder.ToSql()
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository Insert build sql string of rolepermission")
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository Insert insertBuilder error of rolepermission")
		return rolepermission, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository Insert PrepareContext error of rolepermission")
		return rolepermission, err
	}
	row, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository Insert ExecContext error of rolepermission")
		return rolepermission, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository Insert LastInsertId error of rolepermission")
		return rolepermission, err
	}
	newRolePermission, err := r.GetByID(ctx, arguments.RolePermissionGetByID{ID: id})
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository Insert GetByID error of rolepermission")
		return rolepermission, err
	}
	return newRolePermission, nil
}

// setArgsToUpdateBuilder ...
func (r *RepositoryImpl) setArgsToUpdateBuilder(ctx context.Context, updateBuilder sq.UpdateBuilder, params arguments.RolePermissionUpdate) sq.UpdateBuilder {
	logger.WithFields(logger.Fields{
		"params": params,
	}).Infof("Repository setArgsToUpdateBuilder of rolepermission")
	if params.RoleID != nil {
		updateBuilder = updateBuilder.Set("role_id", *params.RoleID)
	}
	if params.PermissionID != nil {
		updateBuilder = updateBuilder.Set("permission_id", *params.PermissionID)
	}
	if params.CreatedBy != nil {
		updateBuilder = updateBuilder.Set("created_by", *params.CreatedBy)
	}
	if params.UpdatedBy != nil {
		updateBuilder = updateBuilder.Set("updated_by", *params.UpdatedBy)
	}
	return updateBuilder
}

// Update ...
func (r *RepositoryImpl) Update(ctx context.Context, params arguments.RolePermissionUpdate) (models.RolePermission, error) {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Repository Update of rolepermission")
	var (
		rolepermission models.RolePermission
		updateBuilder  = sq.Update("role_permission").Where(sq.Eq{"id": *params.ID})
	)
	updateBuilderWithArgs := r.setArgsToUpdateBuilder(ctx, updateBuilder, params)

	sql, args, err := updateBuilderWithArgs.ToSql()
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository Update build sql string of rolepermission")
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository Update updateBuilderWithArgs error of rolepermission")
		return rolepermission, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository Update PrepareContext error of rolepermission")
		return rolepermission, err
	}
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository Update ExecContext error of rolepermission")
		return rolepermission, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository Update RowsAffected error of rolepermission")
		return rolepermission, err
	}
	if rowAffected <= 0 {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   fmt.Errorf("error when update record id %d", *params.ID),
		}).Errorf("Repository Update rowAffected <= 0 of rolepermission")
		return rolepermission, fmt.Errorf("error when update record id %d", *params.ID)
	}
	newRolePermission, err := r.GetByID(ctx, arguments.RolePermissionGetByID{ID: *params.ID})
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository Update GetByID error of rolepermission")
		return rolepermission, err
	}
	return newRolePermission, nil
}

// Delete ...
func (r *RepositoryImpl) Delete(ctx context.Context, params arguments.RolePermissionDelete) (int64, error) {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Repository Delete of rolepermission")
	var (
		id            int64
		deleteBuilder = sq.Delete("role_permission").Where(sq.Eq{"id": params.ID})
	)
	sql, args, err := deleteBuilder.ToSql()
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository Delete build sql string of rolepermission")
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository Delete deleteBuilder error of rolepermission")
		return id, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository Delete PrepareContext error of rolepermission")
		return id, err
	}
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository Delete ExecContext error of rolepermission")
		return id, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository Delete RowsAffected error of rolepermission")
		return id, err
	}
	if rowAffected <= 0 {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   fmt.Errorf("not found record by id %d", params.ID),
		}).Errorf("Repository Update rowAffected <= 0 of rolepermission")
		return id, fmt.Errorf("not found record by id %d", params.ID)
	}
	return params.ID, nil
}

//go:generate mockery -name=IDatabase -output=mocks -case=underscore
