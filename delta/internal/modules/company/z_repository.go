// @generated
package company

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

//scanCompany
func (r *RepositoryImpl) scanCompany(row database.IRow, company *models.Company) error {
	err := row.Scan(
		&company.ID,
		&company.Name,
		&company.CompanyCode,
		&company.Status,
		&company.CreatedBy,
		&company.UpdatedBy,
		&company.ApiSecretKey,
	)
	return err
}

// GetByID ...
func (r *RepositoryImpl) GetByID(ctx context.Context, params arguments.CompanyGetByID) (models.Company, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Repository GetByID of company")
	var (
		company       models.Company
		selectBuilder = sq.Select(
			"id",
			"name",
			"company_code",
			"status",
			"created_by",
			"updated_by",
			"api_secret_key",
		).From("company").Where(sq.Eq{"id": params.ID})
	)
	sql, args, err := selectBuilder.ToSql()
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository GetByID build sql string of company")
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository GetByID selectBuilder error of company")
		return company, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository GetByID PrepareContext error of company")
		return company, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	err = r.scanCompany(row, &company)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository GetByID Scan error of company")
		return company, err
	}
	return company, nil
}

// GetByIDs ...
func (r *RepositoryImpl) GetByIDs(ctx context.Context, params arguments.CompanyGetByIDs) ([]models.Company, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Repository GetByIDs of company")
	var (
		companies     []models.Company
		selectBuilder = sq.Select(
			"id",
			"name",
			"company_code",
			"status",
			"created_by",
			"updated_by",
			"api_secret_key",
		).From("company").Where(sq.Eq{"id": params.IDs})
	)
	sql, args, err := selectBuilder.ToSql()
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository GetByIDs build sql string of company")
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository GetByIDs selectBuilder error of company")
		return companies, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository GetByIDs PrepareContext error of company")
		return companies, err
	}
	rows, err := stmt.QueryContext(ctx, args...)
	defer rows.Close()
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository GetByIDs QueryContext error of company")
		return companies, err
	}
	for rows.Next() {
		company := models.Company{}
		err := r.scanCompany(rows, &company)
		if err != nil {
			logger.WithFields(logger.Fields{
				"TraceID": ctx.Value("TraceID"),
				"Error":   err,
			}).Errorf("Repository GetByIDs Scan error of company")
			return companies, err
		}
		companies = append(companies, company)
	}
	return companies, nil
}

// setArgsToListSelectBuilder ...
func (r *RepositoryImpl) setArgsToListSelectBuilder(ctx context.Context, selectBuilder sq.SelectBuilder, params arguments.CompanyList) sq.SelectBuilder {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Repository setArgsToListSelectBuilder of company")
	if params.ID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"id": params.ID})
	}
	if params.Name != "" {
		selectBuilder = selectBuilder.Where(sq.Like{"name": params.Name})
	}
	if params.CompanyCode != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"company_code": params.CompanyCode})
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
	if params.ApiSecretKey != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"api_secret_key": params.ApiSecretKey})
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
func (r *RepositoryImpl) List(ctx context.Context, params arguments.CompanyList) ([]models.Company, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Repository List of company")
	var (
		companies     []models.Company
		selectBuilder = sq.Select(
			"id",
			"name",
			"company_code",
			"status",
			"created_by",
			"updated_by",
			"api_secret_key",
		).From("company")
	)
	selectBuilderWithArgs := r.setArgsToListSelectBuilder(ctx, selectBuilder, params)
	sql, args, err := selectBuilderWithArgs.ToSql()
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository List build sql string of company")
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository List selectBuilderWithArgs error of company")
		return companies, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository List PrepareContext error of company")
		return companies, err
	}
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository List QueryContext error of company")
		return companies, err
	}
	defer rows.Close()
	for rows.Next() {
		company := models.Company{}
		err := r.scanCompany(rows, &company)
		if err != nil {
			logger.WithFields(logger.Fields{
				"TraceID": ctx.Value("TraceID"),
				"Error":   err,
			}).Errorf("Repository List Scan error of company")
			return companies, err
		}
		companies = append(companies, company)
	}
	return companies, nil
}

// setArgsToCountSelectBuilder ...
func (r *RepositoryImpl) setArgsToCountSelectBuilder(ctx context.Context, selectBuilder sq.SelectBuilder, params arguments.CompanyCount) sq.SelectBuilder {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Repository setArgsToCountSelectBuilder of company")
	if params.ID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"id": params.ID})
	}
	if params.Name != "" {
		selectBuilder = selectBuilder.Where(sq.Like{"name": params.Name})
	}
	if params.CompanyCode != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"company_code": params.CompanyCode})
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
	if params.ApiSecretKey != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"api_secret_key": params.ApiSecretKey})
	}
	return selectBuilder
}

// Count ...
func (r *RepositoryImpl) Count(ctx context.Context, params arguments.CompanyCount) (int64, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Repository Count of company")
	var (
		count         int64
		selectBuilder = sq.Select("COUNT(id)").From("company")
	)
	selectBuilderWithArgs := r.setArgsToCountSelectBuilder(ctx, selectBuilder, params)
	sql, args, err := selectBuilderWithArgs.ToSql()
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository Count build sql string of company")
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Count selectBuilderWithArgs error of company")
		return count, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Count PrepareContext error of company")
		return count, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	err = row.Scan(&count)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Count Scan error of company")
		return count, err
	}
	return count, nil
}

// Insert ...
func (r *RepositoryImpl) Insert(ctx context.Context, params arguments.CompanyInsert) (models.Company, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Repository Insert of company")
	var (
		company       models.Company
		insertBuilder = sq.Insert("company").Columns(
			"name",
			"company_code",
			"status",
			"created_by",
			"updated_by",
			"api_secret_key",
		).Values(
			params.Name,
			params.CompanyCode,
			params.Status,
			params.CreatedBy,
			params.UpdatedBy,
			params.ApiSecretKey,
		)
	)
	sql, args, err := insertBuilder.ToSql()
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository Insert build sql string of company")
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Insert insertBuilder error of company")
		return company, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Insert PrepareContext error of company")
		return company, err
	}
	row, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Insert ExecContext error of company")
		return company, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Insert LastInsertId error of company")
		return company, err
	}
	newCompany, err := r.GetByID(ctx, arguments.CompanyGetByID{ID: id})
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Insert GetByID error of company")
		return company, err
	}
	return newCompany, nil
}

// setArgsToUpdateBuilder ...
func (r *RepositoryImpl) setArgsToUpdateBuilder(ctx context.Context, updateBuilder sq.UpdateBuilder, params arguments.CompanyUpdate) sq.UpdateBuilder {
	logger.WithFields(logger.Fields{
		"params": params,
	}).Infof("Repository setArgsToUpdateBuilder of company")
	if params.Name != nil {
		updateBuilder = updateBuilder.Set("name", *params.Name)
	}
	if params.CompanyCode != nil {
		updateBuilder = updateBuilder.Set("company_code", *params.CompanyCode)
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
	if params.ApiSecretKey != nil {
		updateBuilder = updateBuilder.Set("api_secret_key", *params.ApiSecretKey)
	}
	return updateBuilder
}

// Update ...
func (r *RepositoryImpl) Update(ctx context.Context, params arguments.CompanyUpdate) (models.Company, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Repository Update of company")
	var (
		company       models.Company
		updateBuilder = sq.Update("company").Where(sq.Eq{"id": *params.ID})
	)
	updateBuilderWithArgs := r.setArgsToUpdateBuilder(ctx, updateBuilder, params)

	sql, args, err := updateBuilderWithArgs.ToSql()
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository Update build sql string of company")
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Update updateBuilderWithArgs error of company")
		return company, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Update PrepareContext error of company")
		return company, err
	}
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Update ExecContext error of company")
		return company, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Update RowsAffected error of company")
		return company, err
	}
	if rowAffected <= 0 {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   fmt.Errorf("error when update record id %d", *params.ID),
		}).Errorf("Repository Update rowAffected <= 0 of company")
		return company, fmt.Errorf("error when update record id %d", *params.ID)
	}
	newCompany, err := r.GetByID(ctx, arguments.CompanyGetByID{ID: *params.ID})
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Update GetByID error of company")
		return company, err
	}
	return newCompany, nil
}

// Delete ...
func (r *RepositoryImpl) Delete(ctx context.Context, params arguments.CompanyDelete) (int64, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Repository Delete of company")
	var (
		id            int64
		deleteBuilder = sq.Delete("company").Where(sq.Eq{"id": params.ID})
	)
	sql, args, err := deleteBuilder.ToSql()
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository Delete build sql string of company")
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Delete deleteBuilder error of company")
		return id, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Delete PrepareContext error of company")
		return id, err
	}
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Delete ExecContext error of company")
		return id, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Delete RowsAffected error of company")
		return id, err
	}
	if rowAffected <= 0 {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   fmt.Errorf("not found record by id %d", params.ID),
		}).Errorf("Repository Update rowAffected <= 0 of company")
		return id, fmt.Errorf("not found record by id %d", params.ID)
	}
	return params.ID, nil
}

//go:generate mockery -name=IDatabase -output=mocks -case=underscore
