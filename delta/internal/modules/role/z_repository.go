// @generated
package role

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
func (r *RepositoryImpl) GetByID(ctx context.Context, params arguments.RoleGetByIDArgs) (models.Role, error) {
	log.WithField("params", params).Info("Repository GetByID of role")
	var (
		role          models.Role
		selectBuilder = sq.Select(
			"id",
			"name",
			"company_id",
			"status",
			"created_by",
			"updated_by",
		).From("role").Where(sq.Eq{"id": params.ID})
	)
	sql, args, err := selectBuilder.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository GetByID build sql string of role")
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByID selectBuilder error of role")
		return role, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByID PrepareContext error of role")
		return role, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	err = row.Scan(
		&role.ID,
		&role.Name,
		&role.CompanyID,
		&role.Status,
		&role.CreatedBy,
		&role.UpdatedBy,
	)
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByID Scan error of role")
		return role, err
	}
	return role, nil
}

// GetByIDs ...
func (r *RepositoryImpl) GetByIDs(ctx context.Context, params arguments.RoleGetByIDsArgs) ([]models.Role, error) {
	log.WithField("params", params).Info("Repository GetByIDs of role")
	var (
		roles         []models.Role
		selectBuilder = sq.Select(
			"id",
			"name",
			"company_id",
			"status",
			"created_by",
			"updated_by",
		).From("role").Where(sq.Eq{"id": params.IDs})
	)
	sql, args, err := selectBuilder.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository GetByIDs build sql string of role")
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByIDs selectBuilder error of role")
		return roles, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByIDs PrepareContext error of role")
		return roles, err
	}
	rows, err := stmt.QueryContext(ctx, args...)
	defer rows.Close()
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByIDs QueryContext error of role")
		return roles, err
	}
	for rows.Next() {
		role := models.Role{}
		err := rows.Scan(
			&role.ID,
			&role.Name,
			&role.CompanyID,
			&role.Status,
			&role.CreatedBy,
			&role.UpdatedBy,
		)
		if err != nil {
			log.WithField("Error", err).Error("Repository GetByIDs Scan error of role")
			return roles, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}

// setArgsToListSelectBuilder ...
func (r *RepositoryImpl) setArgsToListSelectBuilder(selectBuilder sq.SelectBuilder, params arguments.RoleListArgs) sq.SelectBuilder {
	log.WithField("params", params).Info("Repository setArgsToListSelectBuilder of role")
	if params.ID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"id": params.ID})
	}
	if params.Name != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"name": params.Name})
	}
	if params.CompanyID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"company_id": params.CompanyID})
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
func (r *RepositoryImpl) List(ctx context.Context, params arguments.RoleListArgs) ([]models.Role, error) {
	log.WithField("params", params).Info("Repository List of role")
	var (
		roles         []models.Role
		selectBuilder = sq.Select(
			"id",
			"name",
			"company_id",
			"status",
			"created_by",
			"updated_by",
		).From("role")
	)
	selectBuilderWithArgs := r.setArgsToListSelectBuilder(selectBuilder, params)
	sql, args, err := selectBuilderWithArgs.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository List build sql string of role")
	if err != nil {
		log.WithField("Error", err).Error("Repository List selectBuilderWithArgs error of role")
		return roles, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository List PrepareContext error of role")
		return roles, err
	}
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		log.WithField("Error", err).Error("Repository List QueryContext error of role")
		return roles, err
	}
	defer rows.Close()
	for rows.Next() {
		role := models.Role{}
		err := rows.Scan(
			&role.ID,
			&role.Name,
			&role.CompanyID,
			&role.Status,
			&role.CreatedBy,
			&role.UpdatedBy,
		)
		if err != nil {
			log.WithField("Error", err).Error("Repository List Scan error of role")
			return roles, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}

// setArgsToCountSelectBuilder ...
func (r *RepositoryImpl) setArgsToCountSelectBuilder(selectBuilder sq.SelectBuilder, params arguments.RoleCountArgs) sq.SelectBuilder {
	log.WithField("params", params).Info("Repository setArgsToCountSelectBuilder of role")
	if params.ID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"id": params.ID})
	}
	if params.Name != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"name": params.Name})
	}
	if params.CompanyID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"company_id": params.CompanyID})
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
func (r *RepositoryImpl) Count(ctx context.Context, params arguments.RoleCountArgs) (int64, error) {
	log.WithField("params", params).Info("Repository Count of role")
	var (
		count         int64
		selectBuilder = sq.Select("COUNT(id)").From("role")
	)
	selectBuilderWithArgs := r.setArgsToCountSelectBuilder(selectBuilder, params)
	sql, args, err := selectBuilderWithArgs.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository Count build sql string of role")
	if err != nil {
		log.WithField("Error", err).Error("Repository Count selectBuilderWithArgs error of role")
		return count, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository Count PrepareContext error of role")
		return count, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	err = row.Scan(&count)
	if err != nil {
		log.WithField("Error", err).Error("Repository Count Scan error of role")
		return count, err
	}
	return count, nil
}

// Insert ...
func (r *RepositoryImpl) Insert(ctx context.Context, params arguments.RoleInsertArgs) (models.Role, error) {
	log.WithField("params", params).Info("Repository Insert of role")
	var (
		role          models.Role
		insertBuilder = sq.Insert("role").Columns(
			"name",
			"company_id",
			"status",
			"created_by",
			"updated_by",
		).Values(
			params.Name,
			params.CompanyID,
			params.Status,
			params.CreatedBy,
			params.UpdatedBy,
		)
	)
	sql, args, err := insertBuilder.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository Insert build sql string of role")
	if err != nil {
		log.WithField("Error", err).Error("Repository Insert insertBuilder error of role")
		return role, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository Insert PrepareContext error of role")
		return role, err
	}
	row, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		log.WithField("Error", err).Error("Repository Insert ExecContext error of role")
		return role, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		log.WithField("Error", err).Error("Repository Insert LastInsertId error of role")
		return role, err
	}
	newRole, err := r.GetByID(ctx, arguments.RoleGetByIDArgs{ID: id})
	if err != nil {
		log.WithField("Error", err).Error("Repository Insert GetByID error of role")
		return role, err
	}
	return newRole, nil
}

// setArgsToUpdateBuilder ...
func (r *RepositoryImpl) setArgsToUpdateBuilder(updateBuilder sq.UpdateBuilder, params arguments.RoleUpdateArgs) sq.UpdateBuilder {
	log.WithField("params", params).Info("Repository setArgsToUpdateBuilder of role")
	if params.Name != nil {
		updateBuilder = updateBuilder.Set("name", *params.Name)
	}

	if params.CompanyID != nil {
		updateBuilder = updateBuilder.Set("company_id", *params.CompanyID)
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
func (r *RepositoryImpl) Update(ctx context.Context, params arguments.RoleUpdateArgs) (models.Role, error) {
	log.WithField("params", params).Info("Repository Update of role")
	var (
		role          models.Role
		updateBuilder = sq.Update("role").Where(sq.Eq{"id": *params.ID})
	)
	updateBuilderWithArgs := r.setArgsToUpdateBuilder(updateBuilder, params)

	sql, args, err := updateBuilderWithArgs.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository Update build sql string of role")
	if err != nil {
		log.WithField("Error", err).Error("Repository Update updateBuilderWithArgs error of role")
		return role, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository Update PrepareContext error of role")
		return role, err
	}
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		log.WithField("Error", err).Error("Repository Update ExecContext error of role")
		return role, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		log.WithField("Error", err).Error("Repository Update RowsAffected error of role")
		return role, err
	}
	if rowAffected <= 0 {
		log.WithField("Error", fmt.Errorf("error when update record id %d", *params.ID)).Error("Repository Update rowAffected <= 0 of role")
		return role, fmt.Errorf("error when update record id %d", *params.ID)
	}
	newRole, err := r.GetByID(ctx, arguments.RoleGetByIDArgs{ID: *params.ID})
	if err != nil {
		log.WithField("Error", err).Error("Repository Update GetByID error of role")
		return role, err
	}
	return newRole, nil
}

// Delete ...
func (r *RepositoryImpl) Delete(ctx context.Context, params arguments.RoleDeleteArgs) (int64, error) {
	log.WithField("params", params).Info("Repository Delete of role")
	var (
		id            int64
		deleteBuilder = sq.Delete("role").Where(sq.Eq{"id": params.ID})
	)
	sql, args, err := deleteBuilder.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository Delete build sql string of role")
	if err != nil {
		log.WithField("Error", err).Error("Repository Delete deleteBuilder error of role")
		return id, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository Delete PrepareContext error of role")
		return id, err
	}
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		log.WithField("Error", err).Error("Repository Delete ExecContext error of role")
		return id, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		log.WithField("Error", err).Error("Repository Delete RowsAffected error of role")
		return id, err
	}
	if rowAffected <= 0 {
		log.WithField("Error", fmt.Errorf("not found record by id %d", params.ID)).Error("Repository Update rowAffected <= 0 of role")
		return id, fmt.Errorf("not found record by id %d", params.ID)
	}
	return params.ID, nil
}

//go:generate mockery -name=IDatabase -output=mocks -case=underscore
