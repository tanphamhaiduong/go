// @generated
package rolepermission

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	log "github.com/sirupsen/logrus"
	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
	"github.com/tanphamhaiduong/go/delta/internal/utils"
)

// GetByID ...
func (r *RepositoryImpl) GetByID(ctx context.Context, params arguments.RolePermissionGetByIDArgs) (models.RolePermission, error) {
	log.WithField("params", params).Info("Repository GetByID of rolepermission")
	var (
		rolepermission models.RolePermission
		selectBuilder  = sq.Select(
			"id",
			"role_id",
			"permission_id",
			"status",
			"created_by",
			"updated_by",
		).From("role_permission").Where(sq.Eq{"id": params.ID})
	)
	sql, args, err := selectBuilder.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository GetByID build sql string of rolepermission")
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByID selectBuilder error of rolepermission")
		return rolepermission, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByID PrepareContext error of rolepermission")
		return rolepermission, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	err = row.Scan(
		&rolepermission.ID,
		&rolepermission.RoleID,
		&rolepermission.PermissionID,
		&rolepermission.Status,
		&rolepermission.CreatedBy,
		&rolepermission.UpdatedBy,
	)
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByID Scan error of rolepermission")
		return rolepermission, err
	}
	return rolepermission, nil
}

// GetByIDs ...
func (r *RepositoryImpl) GetByIDs(ctx context.Context, params arguments.RolePermissionGetByIDsArgs) ([]models.RolePermission, error) {
	log.WithField("params", params).Info("Repository GetByIDs of rolepermission")
	var (
		rolepermissions []models.RolePermission
		selectBuilder   = sq.Select(
			"id",
			"role_id",
			"permission_id",
			"status",
			"created_by",
			"updated_by",
		).From("role_permission").Where(sq.Eq{"id": params.IDs})
	)
	sql, args, err := selectBuilder.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository GetByIDs build sql string of rolepermission")
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByIDs selectBuilder error of rolepermission")
		return rolepermissions, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByIDs PrepareContext error of rolepermission")
		return rolepermissions, err
	}
	rows, err := stmt.QueryContext(ctx, args...)
	defer rows.Close()
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByIDs QueryContext error of rolepermission")
		return rolepermissions, err
	}
	for rows.Next() {
		rolepermission := models.RolePermission{}
		err := rows.Scan(
			&rolepermission.ID,
			&rolepermission.RoleID,
			&rolepermission.PermissionID,
			&rolepermission.Status,
			&rolepermission.CreatedBy,
			&rolepermission.UpdatedBy,
		)
		if err != nil {
			log.WithField("Error", err).Error("Repository GetByIDs Scan error of rolepermission")
			return rolepermissions, err
		}
		rolepermissions = append(rolepermissions, rolepermission)
	}
	return rolepermissions, nil
}

// setArgsToListSelectBuilder ...
func (r *RepositoryImpl) setArgsToListSelectBuilder(selectBuilder sq.SelectBuilder, params arguments.RolePermissionListArgs) sq.SelectBuilder {
	log.WithField("params", params).Info("Repository setArgsToListSelectBuilder of rolepermission")
	if params.ID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"id": params.ID})
	}
	if params.RoleID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"role_id": params.RoleID})
	}
	if params.PermissionID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"permission_id": params.PermissionID})
	}
	if params.Status != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"status": params.Status})
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
	if params.Page != 0 {
		offset := utils.CalculateOffsetForPage(params.Page, params.PageSize)
		selectBuilder = selectBuilder.Offset(uint64(offset))
	}
	return selectBuilder
}

// List ...
func (r *RepositoryImpl) List(ctx context.Context, params arguments.RolePermissionListArgs) ([]models.RolePermission, error) {
	log.WithField("params", params).Info("Repository List of rolepermission")
	var (
		rolepermissions []models.RolePermission
		selectBuilder   = sq.Select(
			"id",
			"role_id",
			"permission_id",
			"status",
			"created_by",
			"updated_by",
		).From("role_permission")
	)
	selectBuilderWithArgs := r.setArgsToListSelectBuilder(selectBuilder, params)
	sql, args, err := selectBuilderWithArgs.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository List build sql string of rolepermission")
	if err != nil {
		log.WithField("Error", err).Error("Repository List selectBuilderWithArgs error of rolepermission")
		return rolepermissions, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository List PrepareContext error of rolepermission")
		return rolepermissions, err
	}
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		log.WithField("Error", err).Error("Repository List QueryContext error of rolepermission")
		return rolepermissions, err
	}
	defer rows.Close()
	for rows.Next() {
		rolepermission := models.RolePermission{}
		err := rows.Scan(
			&rolepermission.ID,
			&rolepermission.RoleID,
			&rolepermission.PermissionID,
			&rolepermission.Status,
			&rolepermission.CreatedBy,
			&rolepermission.UpdatedBy,
		)
		if err != nil {
			log.WithField("Error", err).Error("Repository List Scan error of rolepermission")
			return rolepermissions, err
		}
		rolepermissions = append(rolepermissions, rolepermission)
	}
	return rolepermissions, nil
}

// setArgsToCountSelectBuilder ...
func (r *RepositoryImpl) setArgsToCountSelectBuilder(selectBuilder sq.SelectBuilder, params arguments.RolePermissionCountArgs) sq.SelectBuilder {
	log.WithField("params", params).Info("Repository setArgsToCountSelectBuilder of rolepermission")
	if params.ID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"id": params.ID})
	}
	if params.RoleID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"role_id": params.RoleID})
	}
	if params.PermissionID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"permission_id": params.PermissionID})
	}
	if params.Status != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"status": params.Status})
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
func (r *RepositoryImpl) Count(ctx context.Context, params arguments.RolePermissionCountArgs) (int64, error) {
	log.WithField("params", params).Info("Repository Count of rolepermission")
	var (
		count         int64
		selectBuilder = sq.Select("COUNT(id)").From("role_permission")
	)
	selectBuilderWithArgs := r.setArgsToCountSelectBuilder(selectBuilder, params)
	sql, args, err := selectBuilderWithArgs.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository Count build sql string of rolepermission")
	if err != nil {
		log.WithField("Error", err).Error("Repository Count selectBuilderWithArgs error of rolepermission")
		return count, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository Count PrepareContext error of rolepermission")
		return count, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	err = row.Scan(&count)
	if err != nil {
		log.WithField("Error", err).Error("Repository Count Scan error of rolepermission")
		return count, err
	}
	return count, nil
}

// Insert ...
func (r *RepositoryImpl) Insert(ctx context.Context, params arguments.RolePermissionInsertArgs) (models.RolePermission, error) {
	log.WithField("params", params).Info("Repository Insert of rolepermission")
	var (
		rolepermission models.RolePermission
		insertBuilder  = sq.Insert("role_permission").Columns(
			"role_id",
			"permission_id",
			"status",
			"created_by",
			"updated_by",
		).Values(
			params.RoleID,
			params.PermissionID,
			params.Status,
			params.CreatedBy,
			params.UpdatedBy,
		)
	)
	sql, args, err := insertBuilder.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository Insert build sql string of rolepermission")
	if err != nil {
		log.WithField("Error", err).Error("Repository Insert insertBuilder error of rolepermission")
		return rolepermission, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository Insert PrepareContext error of rolepermission")
		return rolepermission, err
	}
	row, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		log.WithField("Error", err).Error("Repository Insert ExecContext error of rolepermission")
		return rolepermission, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		log.WithField("Error", err).Error("Repository Insert LastInsertId error of rolepermission")
		return rolepermission, err
	}
	newRolePermission, err := r.GetByID(ctx, arguments.RolePermissionGetByIDArgs{ID: id})
	if err != nil {
		log.WithField("Error", err).Error("Repository Insert GetByID error of rolepermission")
		return rolepermission, err
	}
	return newRolePermission, nil
}

// setArgsToUpdateBuilder ...
func (r *RepositoryImpl) setArgsToUpdateBuilder(updateBuilder sq.UpdateBuilder, params arguments.RolePermissionUpdateArgs) sq.UpdateBuilder {
	log.WithField("params", params).Info("Repository setArgsToUpdateBuilder of rolepermission")
	if params.RoleID != nil {
		updateBuilder = updateBuilder.Set("role_id", *params.RoleID)
	}

	if params.PermissionID != nil {
		updateBuilder = updateBuilder.Set("permission_id", *params.PermissionID)
	}

	if params.Status != nil {
		updateBuilder = updateBuilder.Set("status", *params.Status)
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
func (r *RepositoryImpl) Update(ctx context.Context, params arguments.RolePermissionUpdateArgs) (models.RolePermission, error) {
	log.WithField("params", params).Info("Repository Update of rolepermission")
	var (
		rolepermission models.RolePermission
		updateBuilder  = sq.Update("role_permission").Where(sq.Eq{"id": *params.ID})
	)
	updateBuilderWithArgs := r.setArgsToUpdateBuilder(updateBuilder, params)

	sql, args, err := updateBuilderWithArgs.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository Update build sql string of rolepermission")
	if err != nil {
		log.WithField("Error", err).Error("Repository Update updateBuilderWithArgs error of rolepermission")
		return rolepermission, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository Update PrepareContext error of rolepermission")
		return rolepermission, err
	}
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		log.WithField("Error", err).Error("Repository Update ExecContext error of rolepermission")
		return rolepermission, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		log.WithField("Error", err).Error("Repository Update RowsAffected error of rolepermission")
		return rolepermission, err
	}
	if rowAffected <= 0 {
		log.WithField("Error", fmt.Errorf("error when update record id %d", *params.ID)).Error("Repository Update rowAffected <= 0 of rolepermission")
		return rolepermission, fmt.Errorf("error when update record id %d", *params.ID)
	}
	newRolePermission, err := r.GetByID(ctx, arguments.RolePermissionGetByIDArgs{ID: *params.ID})
	if err != nil {
		log.WithField("Error", err).Error("Repository Update GetByID error of rolepermission")
		return rolepermission, err
	}
	return newRolePermission, nil
}

// Delete ...
func (r *RepositoryImpl) Delete(ctx context.Context, params arguments.RolePermissionDeleteArgs) (int64, error) {
	log.WithField("params", params).Info("Repository Delete of rolepermission")
	var (
		id            int64
		deleteBuilder = sq.Delete("role_permission").Where(sq.Eq{"id": params.ID})
	)
	sql, args, err := deleteBuilder.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository Delete build sql string of rolepermission")
	if err != nil {
		log.WithField("Error", err).Error("Repository Delete deleteBuilder error of rolepermission")
		return id, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository Delete PrepareContext error of rolepermission")
		return id, err
	}
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		log.WithField("Error", err).Error("Repository Delete ExecContext error of rolepermission")
		return id, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		log.WithField("Error", err).Error("Repository Delete RowsAffected error of rolepermission")
		return id, err
	}
	if rowAffected <= 0 {
		log.WithField("Error", fmt.Errorf("not found record by id %d", params.ID)).Error("Repository Update rowAffected <= 0 of rolepermission")
		return id, fmt.Errorf("not found record by id %d", params.ID)
	}
	return params.ID, nil
}

//go:generate mockery -name=IDatabase -output=mocks -case=underscore