// @generated
package permission

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	log "github.com/sirupsen/logrus"
	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/database"
	"github.com/tanphamhaiduong/go/delta/internal/models"
	"github.com/tanphamhaiduong/go/delta/internal/utils"
)

//scanPermission
func (r *RepositoryImpl) scanPermission(row database.IRow, permission *models.Permission) error {
	err := row.Scan(
		&permission.ID,
		&permission.Name,
		&permission.Status,
		&permission.CreatedBy,
		&permission.UpdatedBy,
	)
	return err
}

// GetByID ...
func (r *RepositoryImpl) GetByID(ctx context.Context, params arguments.PermissionGetByIDArgs) (models.Permission, error) {
	log.WithField("params", params).Info("Repository GetByID of permission")
	var (
		permission    models.Permission
		selectBuilder = sq.Select(
			"id",
			"name",
			"status",
			"created_by",
			"updated_by",
		).From("permission").Where(sq.Eq{"id": params.ID})
	)
	sql, args, err := selectBuilder.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository GetByID build sql string of permission")
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByID selectBuilder error of permission")
		return permission, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByID PrepareContext error of permission")
		return permission, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	err = r.scanPermission(row, &permission)
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByID Scan error of permission")
		return permission, err
	}
	return permission, nil
}

// GetByIDs ...
func (r *RepositoryImpl) GetByIDs(ctx context.Context, params arguments.PermissionGetByIDsArgs) ([]models.Permission, error) {
	log.WithField("params", params).Info("Repository GetByIDs of permission")
	var (
		permissions   []models.Permission
		selectBuilder = sq.Select(
			"id",
			"name",
			"status",
			"created_by",
			"updated_by",
		).From("permission").Where(sq.Eq{"id": params.IDs})
	)
	sql, args, err := selectBuilder.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository GetByIDs build sql string of permission")
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByIDs selectBuilder error of permission")
		return permissions, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByIDs PrepareContext error of permission")
		return permissions, err
	}
	rows, err := stmt.QueryContext(ctx, args...)
	defer rows.Close()
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByIDs QueryContext error of permission")
		return permissions, err
	}
	for rows.Next() {
		permission := models.Permission{}
		err := r.scanPermission(rows, &permission)
		if err != nil {
			log.WithField("Error", err).Error("Repository GetByIDs Scan error of permission")
			return permissions, err
		}
		permissions = append(permissions, permission)
	}
	return permissions, nil
}

// setArgsToListSelectBuilder ...
func (r *RepositoryImpl) setArgsToListSelectBuilder(selectBuilder sq.SelectBuilder, params arguments.PermissionListArgs) sq.SelectBuilder {
	log.WithField("params", params).Info("Repository setArgsToListSelectBuilder of permission")
	if params.ID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"id": params.ID})
	}
	if params.Name != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"name": params.Name})
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
func (r *RepositoryImpl) List(ctx context.Context, params arguments.PermissionListArgs) ([]models.Permission, error) {
	log.WithField("params", params).Info("Repository List of permission")
	var (
		permissions   []models.Permission
		selectBuilder = sq.Select(
			"id",
			"name",
			"status",
			"created_by",
			"updated_by",
		).From("permission")
	)
	selectBuilderWithArgs := r.setArgsToListSelectBuilder(selectBuilder, params)
	sql, args, err := selectBuilderWithArgs.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository List build sql string of permission")
	if err != nil {
		log.WithField("Error", err).Error("Repository List selectBuilderWithArgs error of permission")
		return permissions, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository List PrepareContext error of permission")
		return permissions, err
	}
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		log.WithField("Error", err).Error("Repository List QueryContext error of permission")
		return permissions, err
	}
	defer rows.Close()
	for rows.Next() {
		permission := models.Permission{}
		err := r.scanPermission(rows, &permission)
		if err != nil {
			log.WithField("Error", err).Error("Repository List Scan error of permission")
			return permissions, err
		}
		permissions = append(permissions, permission)
	}
	return permissions, nil
}

// setArgsToCountSelectBuilder ...
func (r *RepositoryImpl) setArgsToCountSelectBuilder(selectBuilder sq.SelectBuilder, params arguments.PermissionCountArgs) sq.SelectBuilder {
	log.WithField("params", params).Info("Repository setArgsToCountSelectBuilder of permission")
	if params.ID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"id": params.ID})
	}
	if params.Name != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"name": params.Name})
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
func (r *RepositoryImpl) Count(ctx context.Context, params arguments.PermissionCountArgs) (int64, error) {
	log.WithField("params", params).Info("Repository Count of permission")
	var (
		count         int64
		selectBuilder = sq.Select("COUNT(id)").From("permission")
	)
	selectBuilderWithArgs := r.setArgsToCountSelectBuilder(selectBuilder, params)
	sql, args, err := selectBuilderWithArgs.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository Count build sql string of permission")
	if err != nil {
		log.WithField("Error", err).Error("Repository Count selectBuilderWithArgs error of permission")
		return count, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository Count PrepareContext error of permission")
		return count, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	err = row.Scan(&count)
	if err != nil {
		log.WithField("Error", err).Error("Repository Count Scan error of permission")
		return count, err
	}
	return count, nil
}

// Insert ...
func (r *RepositoryImpl) Insert(ctx context.Context, params arguments.PermissionInsertArgs) (models.Permission, error) {
	log.WithField("params", params).Info("Repository Insert of permission")
	var (
		permission    models.Permission
		insertBuilder = sq.Insert("permission").Columns(
			"name",
			"status",
			"created_by",
			"updated_by",
		).Values(
			params.Name,
			params.Status,
			params.CreatedBy,
			params.UpdatedBy,
		)
	)
	sql, args, err := insertBuilder.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository Insert build sql string of permission")
	if err != nil {
		log.WithField("Error", err).Error("Repository Insert insertBuilder error of permission")
		return permission, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository Insert PrepareContext error of permission")
		return permission, err
	}
	row, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		log.WithField("Error", err).Error("Repository Insert ExecContext error of permission")
		return permission, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		log.WithField("Error", err).Error("Repository Insert LastInsertId error of permission")
		return permission, err
	}
	newPermission, err := r.GetByID(ctx, arguments.PermissionGetByIDArgs{ID: id})
	if err != nil {
		log.WithField("Error", err).Error("Repository Insert GetByID error of permission")
		return permission, err
	}
	return newPermission, nil
}

// setArgsToUpdateBuilder ...
func (r *RepositoryImpl) setArgsToUpdateBuilder(updateBuilder sq.UpdateBuilder, params arguments.PermissionUpdateArgs) sq.UpdateBuilder {
	log.WithField("params", params).Info("Repository setArgsToUpdateBuilder of permission")
	if params.Name != nil {
		updateBuilder = updateBuilder.Set("name", *params.Name)
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
func (r *RepositoryImpl) Update(ctx context.Context, params arguments.PermissionUpdateArgs) (models.Permission, error) {
	log.WithField("params", params).Info("Repository Update of permission")
	var (
		permission    models.Permission
		updateBuilder = sq.Update("permission").Where(sq.Eq{"id": *params.ID})
	)
	updateBuilderWithArgs := r.setArgsToUpdateBuilder(updateBuilder, params)

	sql, args, err := updateBuilderWithArgs.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository Update build sql string of permission")
	if err != nil {
		log.WithField("Error", err).Error("Repository Update updateBuilderWithArgs error of permission")
		return permission, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository Update PrepareContext error of permission")
		return permission, err
	}
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		log.WithField("Error", err).Error("Repository Update ExecContext error of permission")
		return permission, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		log.WithField("Error", err).Error("Repository Update RowsAffected error of permission")
		return permission, err
	}
	if rowAffected <= 0 {
		log.WithField("Error", fmt.Errorf("error when update record id %d", *params.ID)).Error("Repository Update rowAffected <= 0 of permission")
		return permission, fmt.Errorf("error when update record id %d", *params.ID)
	}
	newPermission, err := r.GetByID(ctx, arguments.PermissionGetByIDArgs{ID: *params.ID})
	if err != nil {
		log.WithField("Error", err).Error("Repository Update GetByID error of permission")
		return permission, err
	}
	return newPermission, nil
}

// Delete ...
func (r *RepositoryImpl) Delete(ctx context.Context, params arguments.PermissionDeleteArgs) (int64, error) {
	log.WithField("params", params).Info("Repository Delete of permission")
	var (
		id            int64
		deleteBuilder = sq.Delete("permission").Where(sq.Eq{"id": params.ID})
	)
	sql, args, err := deleteBuilder.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository Delete build sql string of permission")
	if err != nil {
		log.WithField("Error", err).Error("Repository Delete deleteBuilder error of permission")
		return id, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository Delete PrepareContext error of permission")
		return id, err
	}
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		log.WithField("Error", err).Error("Repository Delete ExecContext error of permission")
		return id, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		log.WithField("Error", err).Error("Repository Delete RowsAffected error of permission")
		return id, err
	}
	if rowAffected <= 0 {
		log.WithField("Error", fmt.Errorf("not found record by id %d", params.ID)).Error("Repository Update rowAffected <= 0 of permission")
		return id, fmt.Errorf("not found record by id %d", params.ID)
	}
	return params.ID, nil
}

//go:generate mockery -name=IDatabase -output=mocks -case=underscore
