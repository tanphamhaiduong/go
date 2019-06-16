// @generated
package company

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
func (r *RepositoryImpl) GetByID(ctx context.Context, params arguments.CompanyGetByIDArgs) (models.Company, error) {
	log.WithField("params", params).Info("company of Repository GetByID")
	var (
		company       models.Company
		selectBuilder = sq.Select(
			"id",
			"name",
			"status",
			"created_by",
			"updated_by",
		).From("company").Where(sq.Eq{"id": params.ID})
	)
	sql, args, err := selectBuilder.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("company of Repository GetByID build sql string")
	if err != nil {
		log.WithField("Error", err).Error("company of Repository GetByID selectBuilder error")
		return company, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("company of Repository GetByID PrepareContext error")
		return company, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	err = row.Scan(
		&company.ID,
		&company.Name,
		&company.Status,
		&company.CreatedBy,
		&company.UpdatedBy,
	)
	if err != nil {
		log.WithField("Error", err).Error("company of Repository GetByID Scan error")
		return company, err
	}
	return company, nil
}

// GetByIDs ...
func (r *RepositoryImpl) GetByIDs(ctx context.Context, params arguments.CompanyGetByIDsArgs) ([]models.Company, error) {
	log.WithField("params", params).Info("company of Repository GetByIDs")
	var (
		companies     []models.Company
		selectBuilder = sq.Select(
			"id",
			"name",
			"status",
			"created_by",
			"updated_by",
		).From("company").Where(sq.Eq{"id": params.IDs})
	)
	sql, args, err := selectBuilder.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("company of Repository GetByIDs build sql string")
	if err != nil {
		log.WithField("Error", err).Error("company of Repository GetByIDs selectBuilder error")
		return companies, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("company of Repository GetByIDs PrepareContext error")
		return companies, err
	}
	rows, err := stmt.QueryContext(ctx, args...)
	defer rows.Close()
	if err != nil {
		log.WithField("Error", err).Error("company of Repository GetByIDs QueryContext error")
		return companies, err
	}
	for rows.Next() {
		company := models.Company{}
		err := rows.Scan(
			&company.ID,
			&company.Name,
			&company.Status,
			&company.CreatedBy,
			&company.UpdatedBy,
		)
		if err != nil {
			log.WithField("Error", err).Error("company of Repository GetByIDs Scan error")
			return companies, err
		}
		companies = append(companies, company)
	}
	return companies, nil
}

// setArgsToListSelectBuilder ...
func (r *RepositoryImpl) setArgsToListSelectBuilder(selectBuilder sq.SelectBuilder, params arguments.CompanyListArgs) sq.SelectBuilder {
	log.WithField("params", params).Info("company of Repository setArgsToListSelectBuilder")
	if params.ID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"id": params.ID})
	}
	if params.Name != "" {
		selectBuilder = selectBuilder.Where(sq.Like{"name": params.Name})
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
func (r *RepositoryImpl) List(ctx context.Context, params arguments.CompanyListArgs) ([]models.Company, error) {
	log.WithField("params", params).Info("company of Repository List")
	var (
		companies     []models.Company
		selectBuilder = sq.Select(
			"id",
			"name",
			"status",
			"created_by",
			"updated_by",
		).From("company")
	)
	selectBuilderWithArgs := r.setArgsToListSelectBuilder(selectBuilder, params)
	sql, args, err := selectBuilderWithArgs.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("company of Repository List build sql string")
	if err != nil {
		log.WithField("Error", err).Error("company of Repository List selectBuilderWithArgs error")
		return companies, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("company of Repository List PrepareContext error")
		return companies, err
	}
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		log.WithField("Error", err).Error("company of Repository List QueryContext error")
		return companies, err
	}
	defer rows.Close()
	for rows.Next() {
		company := models.Company{}
		err := rows.Scan(
			&company.ID,
			&company.Name,
			&company.Status,
			&company.CreatedBy,
			&company.UpdatedBy,
		)
		if err != nil {
			log.WithField("Error", err).Error("company of Repository List Scan error")
			return companies, err
		}
		companies = append(companies, company)
	}
	return companies, nil
}

// setArgsToCountSelectBuilder ...
func (r *RepositoryImpl) setArgsToCountSelectBuilder(selectBuilder sq.SelectBuilder, params arguments.CompanyCountArgs) sq.SelectBuilder {
	log.WithField("params", params).Info("company of Repository setArgsToCountSelectBuilder")
	if params.ID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"id": params.ID})
	}
	if params.Name != "" {
		selectBuilder = selectBuilder.Where(sq.Like{"name": params.Name})
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
func (r *RepositoryImpl) Count(ctx context.Context, params arguments.CompanyCountArgs) (int64, error) {
	log.WithField("params", params).Info("company of Repository Count")
	var (
		count         int64
		selectBuilder = sq.Select("COUNT(id)").From("company")
	)
	selectBuilderWithArgs := r.setArgsToCountSelectBuilder(selectBuilder, params)
	sql, args, err := selectBuilderWithArgs.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("company of Repository Count build sql string")
	if err != nil {
		log.WithField("Error", err).Error("company of Repository Count selectBuilderWithArgs error")
		return count, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("company of Repository Count PrepareContext error")
		return count, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	err = row.Scan(&count)
	if err != nil {
		log.WithField("Error", err).Error("company of Repository Count Scan error")
		return count, err
	}
	return count, nil
}

// Insert ...
func (r *RepositoryImpl) Insert(ctx context.Context, params arguments.CompanyInsertArgs) (models.Company, error) {
	log.WithField("params", params).Info("company of Repository Insert")
	var (
		company       models.Company
		insertBuilder = sq.Insert("company").Columns(
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
	}).Info("company of Repository Insert build sql string")
	if err != nil {
		log.WithField("Error", err).Error("company of Repository Insert insertBuilder error")
		return company, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("company of Repository Insert PrepareContext error")
		return company, err
	}
	row, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		log.WithField("Error", err).Error("company of Repository Insert ExecContext error")
		return company, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		log.WithField("Error", err).Error("company of Repository Insert LastInsertId error")
		return company, err
	}
	newCompany, err := r.GetByID(ctx, arguments.CompanyGetByIDArgs{ID: id})
	if err != nil {
		log.WithField("Error", err).Error("company of Repository Insert GetByID error")
		return company, err
	}
	return newCompany, nil
}

// setArgsToUpdateBuilder ...
func (r *RepositoryImpl) setArgsToUpdateBuilder(updateBuilder sq.UpdateBuilder, params arguments.CompanyUpdateArgs) sq.UpdateBuilder {
	log.WithField("params", params).Info("company of Repository setArgsToUpdateBuilder")
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
func (r *RepositoryImpl) Update(ctx context.Context, params arguments.CompanyUpdateArgs) (models.Company, error) {
	log.WithField("params", params).Info("company of Repository Update")
	var (
		company       models.Company
		updateBuilder = sq.Update("company").Where(sq.Eq{"id": *params.ID})
	)
	updateBuilderWithArgs := r.setArgsToUpdateBuilder(updateBuilder, params)

	sql, args, err := updateBuilderWithArgs.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("company of Repository Update build sql string")
	if err != nil {
		log.WithField("Error", err).Error("company of Repository Update updateBuilderWithArgs error")
		return company, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("company of Repository Update PrepareContext error")
		return company, err
	}
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		log.WithField("Error", err).Error("company of Repository Update ExecContext error")
		return company, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		log.WithField("Error", err).Error("company of Repository Update RowsAffected error")
		return company, err
	}
	if rowAffected <= 0 {
		log.WithField("Error", fmt.Errorf("error when update record id %d", *params.ID)).Error("company of Repository Update rowAffected <= 0")
		return company, fmt.Errorf("error when update record id %d", *params.ID)
	}
	newCompany, err := r.GetByID(ctx, arguments.CompanyGetByIDArgs{ID: *params.ID})
	if err != nil {
		log.WithField("Error", err).Error("company of Repository Update GetByID error")
		return company, err
	}
	return newCompany, nil
}

// Delete ...
func (r *RepositoryImpl) Delete(ctx context.Context, params arguments.CompanyDeleteArgs) (int64, error) {
	log.WithField("params", params).Info("company of Repository Delete")
	var (
		id            int64
		deleteBuilder = sq.Delete("company").Where(sq.Eq{"id": params.ID})
	)
	sql, args, err := deleteBuilder.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("company of Repository Delete build sql string")
	if err != nil {
		log.WithField("Error", err).Error("company of Repository Delete deleteBuilder error")
		return id, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("company of Repository Delete PrepareContext error")
		return id, err
	}
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		log.WithField("Error", err).Error("company of Repository Delete ExecContext error")
		return id, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		log.WithField("Error", err).Error("company of Repository Delete RowsAffected error")
		return id, err
	}
	if rowAffected <= 0 {
		log.WithField("Error", fmt.Errorf("not found record by id %d", params.ID)).Error("company of Repository Update rowAffected <= 0")
		return id, fmt.Errorf("not found record by id %d", params.ID)
	}
	return params.ID, nil
}

//go:generate mockery -name=IDatabase -output=mocks -case=underscore
