// @generated
package company

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

//scanCompany
func (r *RepositoryImpl) scanCompany(row database.IRow, company *models.Company) error {
	err := row.Scan(
		&company.ID,
		&company.Name,
		&company.Status,
		&company.CreatedBy,
		&company.UpdatedBy,
	)
	return err
}

// GetByID ...
func (r *RepositoryImpl) GetByID(ctx context.Context, params arguments.CompanyGetByIDArgs) (models.Company, error) {
	log.WithField("params", params).Info("Repository GetByID of company")
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
	}).Info("Repository GetByID build sql string of company")
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByID selectBuilder error of company")
		return company, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByID PrepareContext error of company")
		return company, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	err = r.scanCompany(row, &company)
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByID Scan error of company")
		return company, err
	}
	return company, nil
}

// GetByIDs ...
func (r *RepositoryImpl) GetByIDs(ctx context.Context, params arguments.CompanyGetByIDsArgs) ([]models.Company, error) {
	log.WithField("params", params).Info("Repository GetByIDs of company")
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
	}).Info("Repository GetByIDs build sql string of company")
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByIDs selectBuilder error of company")
		return companies, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByIDs PrepareContext error of company")
		return companies, err
	}
	rows, err := stmt.QueryContext(ctx, args...)
	defer rows.Close()
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByIDs QueryContext error of company")
		return companies, err
	}
	for rows.Next() {
		company := models.Company{}
		err := r.scanCompany(rows, &company)
		if err != nil {
			log.WithField("Error", err).Error("Repository GetByIDs Scan error of company")
			return companies, err
		}
		companies = append(companies, company)
	}
	return companies, nil
}

// setArgsToListSelectBuilder ...
func (r *RepositoryImpl) setArgsToListSelectBuilder(selectBuilder sq.SelectBuilder, params arguments.CompanyListArgs) sq.SelectBuilder {
	log.WithField("params", params).Info("Repository setArgsToListSelectBuilder of company")
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
	log.WithField("params", params).Info("Repository List of company")
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
	}).Info("Repository List build sql string of company")
	if err != nil {
		log.WithField("Error", err).Error("Repository List selectBuilderWithArgs error of company")
		return companies, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository List PrepareContext error of company")
		return companies, err
	}
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		log.WithField("Error", err).Error("Repository List QueryContext error of company")
		return companies, err
	}
	defer rows.Close()
	for rows.Next() {
		company := models.Company{}
		err := r.scanCompany(rows, &company)
		if err != nil {
			log.WithField("Error", err).Error("Repository List Scan error of company")
			return companies, err
		}
		companies = append(companies, company)
	}
	return companies, nil
}

// setArgsToCountSelectBuilder ...
func (r *RepositoryImpl) setArgsToCountSelectBuilder(selectBuilder sq.SelectBuilder, params arguments.CompanyCountArgs) sq.SelectBuilder {
	log.WithField("params", params).Info("Repository setArgsToCountSelectBuilder of company")
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
	log.WithField("params", params).Info("Repository Count of company")
	var (
		count         int64
		selectBuilder = sq.Select("COUNT(id)").From("company")
	)
	selectBuilderWithArgs := r.setArgsToCountSelectBuilder(selectBuilder, params)
	sql, args, err := selectBuilderWithArgs.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository Count build sql string of company")
	if err != nil {
		log.WithField("Error", err).Error("Repository Count selectBuilderWithArgs error of company")
		return count, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository Count PrepareContext error of company")
		return count, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	err = row.Scan(&count)
	if err != nil {
		log.WithField("Error", err).Error("Repository Count Scan error of company")
		return count, err
	}
	return count, nil
}

// Insert ...
func (r *RepositoryImpl) Insert(ctx context.Context, params arguments.CompanyInsertArgs) (models.Company, error) {
	log.WithField("params", params).Info("Repository Insert of company")
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
	}).Info("Repository Insert build sql string of company")
	if err != nil {
		log.WithField("Error", err).Error("Repository Insert insertBuilder error of company")
		return company, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository Insert PrepareContext error of company")
		return company, err
	}
	row, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		log.WithField("Error", err).Error("Repository Insert ExecContext error of company")
		return company, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		log.WithField("Error", err).Error("Repository Insert LastInsertId error of company")
		return company, err
	}
	newCompany, err := r.GetByID(ctx, arguments.CompanyGetByIDArgs{ID: id})
	if err != nil {
		log.WithField("Error", err).Error("Repository Insert GetByID error of company")
		return company, err
	}
	return newCompany, nil
}

// setArgsToUpdateBuilder ...
func (r *RepositoryImpl) setArgsToUpdateBuilder(updateBuilder sq.UpdateBuilder, params arguments.CompanyUpdateArgs) sq.UpdateBuilder {
	log.WithField("params", params).Info("Repository setArgsToUpdateBuilder of company")
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
	log.WithField("params", params).Info("Repository Update of company")
	var (
		company       models.Company
		updateBuilder = sq.Update("company").Where(sq.Eq{"id": *params.ID})
	)
	updateBuilderWithArgs := r.setArgsToUpdateBuilder(updateBuilder, params)

	sql, args, err := updateBuilderWithArgs.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository Update build sql string of company")
	if err != nil {
		log.WithField("Error", err).Error("Repository Update updateBuilderWithArgs error of company")
		return company, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository Update PrepareContext error of company")
		return company, err
	}
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		log.WithField("Error", err).Error("Repository Update ExecContext error of company")
		return company, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		log.WithField("Error", err).Error("Repository Update RowsAffected error of company")
		return company, err
	}
	if rowAffected <= 0 {
		log.WithField("Error", fmt.Errorf("error when update record id %d", *params.ID)).Error("Repository Update rowAffected <= 0 of company")
		return company, fmt.Errorf("error when update record id %d", *params.ID)
	}
	newCompany, err := r.GetByID(ctx, arguments.CompanyGetByIDArgs{ID: *params.ID})
	if err != nil {
		log.WithField("Error", err).Error("Repository Update GetByID error of company")
		return company, err
	}
	return newCompany, nil
}

// Delete ...
func (r *RepositoryImpl) Delete(ctx context.Context, params arguments.CompanyDeleteArgs) (int64, error) {
	log.WithField("params", params).Info("Repository Delete of company")
	var (
		id            int64
		deleteBuilder = sq.Delete("company").Where(sq.Eq{"id": params.ID})
	)
	sql, args, err := deleteBuilder.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository Delete build sql string of company")
	if err != nil {
		log.WithField("Error", err).Error("Repository Delete deleteBuilder error of company")
		return id, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository Delete PrepareContext error of company")
		return id, err
	}
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		log.WithField("Error", err).Error("Repository Delete ExecContext error of company")
		return id, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		log.WithField("Error", err).Error("Repository Delete RowsAffected error of company")
		return id, err
	}
	if rowAffected <= 0 {
		log.WithField("Error", fmt.Errorf("not found record by id %d", params.ID)).Error("Repository Update rowAffected <= 0 of company")
		return id, fmt.Errorf("not found record by id %d", params.ID)
	}
	return params.ID, nil
}

//go:generate mockery -name=IDatabase -output=mocks -case=underscore
