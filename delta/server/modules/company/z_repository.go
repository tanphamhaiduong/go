// @generated
package company

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/tanphamhaiduong/go/delta/server/arguments"
	"github.com/tanphamhaiduong/go/delta/server/models"
	"github.com/tanphamhaiduong/go/delta/server/utils"
)

// GetByID ...
func (r repositoryImpl) GetByID(ctx context.Context, params arguments.CompanyGetByIDArgs) (models.Company, error) {
	var (
		company       = models.Company{}
		selectBuilder = sq.Select("*").From("company").Where(sq.Eq{"id": params.ID})
	)
	sql, args, err := selectBuilder.ToSql()
	if err != nil {
		return company, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return company, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	row.Scan(
		&company.ID,
		&company.Name,
		&company.Status,
		&company.CreatedBy,
		&company.UpdatedBy,
	)
	return company, nil
}

// GetByID ...
func (r repositoryImpl) GetByIDs(ctx context.Context, params arguments.CompanyGetByIDsArgs) ([]models.Company, error) {
	var (
		companies     = []models.Company{}
		selectBuilder = sq.Select("*").From("company").Where(sq.Eq{"id": params.IDs})
	)
	sql, args, err := selectBuilder.ToSql()
	if err != nil {
		return companies, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return companies, err
	}
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
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
			return companies, err
		}
		companies = append(companies, company)
	}
	return companies, nil
}

// setArgsToListSelectBuilder ...
func setArgsToListSelectBuilder(selectBuilder sq.SelectBuilder, params arguments.CompanyListArgs) sq.SelectBuilder {
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
func (r repositoryImpl) List(ctx context.Context, params arguments.CompanyListArgs) ([]models.Company, error) {
	var (
		companies     = []models.Company{}
		selectBuilder = sq.Select("*").From("company")
	)
	selectBuilderWithArgs := setArgsToListSelectBuilder(selectBuilder, params)
	sql, args, err := selectBuilderWithArgs.ToSql()
	if err != nil {
		return companies, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return companies, err
	}
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
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
			return companies, err
		}
		companies = append(companies, company)
	}
	return companies, nil
}

// setArgsToCountSelectBuilder ...
func setArgsToCountSelectBuilder(selectBuilder sq.SelectBuilder, params arguments.CompanyCountArgs) sq.SelectBuilder {
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
func (r repositoryImpl) Count(ctx context.Context, params arguments.CompanyCountArgs) (int64, error) {
	var (
		count         int64
		selectBuilder = sq.Select("COUNT(id)").From("company")
	)
	selectBuilderWithArgs := setArgsToCountSelectBuilder(selectBuilder, params)
	sql, args, err := selectBuilderWithArgs.ToSql()
	if err != nil {
		return count, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return count, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	err = row.Scan(&count)
	if err != nil {
		return count, err
	}
	return count, nil
}

// Insert ...
func (r repositoryImpl) Insert(ctx context.Context, params arguments.CompanyInsertArgs) (models.Company, error) {
	var (
		company       = models.Company{}
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
	if err != nil {
		return company, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return company, err
	}
	row, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return company, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		return company, err
	}
	company = models.Company{
		ID:   id,
		Name: params.Name,

		Status: params.Status,

		CreatedBy: params.CreatedBy,

		UpdatedBy: params.UpdatedBy,
	}
	return company, nil
}

// setArgsToUpdateBuilder ...
func setArgsToUpdateBuilder(updateBuilder sq.UpdateBuilder, params arguments.CompanyUpdateArgs) sq.UpdateBuilder {
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
func (r repositoryImpl) Update(ctx context.Context, params arguments.CompanyUpdateArgs) (models.Company, error) {
	var (
		company       = models.Company{}
		updateBuilder = sq.Update("company").Where(sq.Eq{"id": *params.ID})
	)
	updateBuilderWithArgs := setArgsToUpdateBuilder(updateBuilder, params)

	sql, args, err := updateBuilderWithArgs.ToSql()
	if err != nil {
		return company, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return company, err
	}
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return company, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return company, err
	}
	if rowAffected == 0 {
		return company, fmt.Errorf("error when update record id %d", *params.ID)
	}
	company = models.Company{
		ID:        *params.ID,
		Name:      *params.Name,
		Status:    *params.Status,
		CreatedBy: *params.CreatedBy,
		UpdatedBy: *params.UpdatedBy,
	}
	return company, nil
}

// Delete ...
func (r repositoryImpl) Delete(ctx context.Context, params arguments.CompanyDeleteArgs) (int64, error) {
	var (
		id            int64
		deleteBuilder = sq.Delete("company").Where(sq.Eq{"id": params.ID})
	)
	sql, args, err := deleteBuilder.ToSql()
	if err != nil {
		return id, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return id, err
	}
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return id, err
	}
	rowAffected, err := result.RowsAffected()
	if rowAffected == 0 {
		return id, fmt.Errorf("not found record by id %d", params.ID)
	}
	return params.ID, nil
}

//go:generate mockery -name=iDatabase -output=mocks -case=underscore
