// @generated
package company

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/afex/hystrix-go/hystrix"
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
		&company.Status,
		&company.CreatedBy,
		&company.UpdatedBy,
	)
	return err
}

// GetByID ...
func (r *RepositoryImpl) GetByID(ctx context.Context, params arguments.CompanyGetByID) (models.Company, error) {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Repository GetByID of company")
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
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository GetByID build sql string of company")
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository GetByID selectBuilder error of company")
		return company, err
	}
	output := make(chan models.Company, 1)
	errors := hystrix.Go("companyGetByID", func() error {
		stmt, err := r.db.PrepareContext(ctx, sql)
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository GetByID PrepareContext error of company")
			return err
		}
		row := stmt.QueryRowContext(ctx, args...)
		err = r.scanCompany(row, &company)
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository GetByID Scan error of company")
			return err
		}
		output <- company
		return nil
	}, nil)

	select {
	case err := <-errors:
		return company, err
	case out := <-output:
		return out, nil
	}
}

// GetByIDs ...
func (r *RepositoryImpl) GetByIDs(ctx context.Context, params arguments.CompanyGetByIDs) ([]models.Company, error) {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Repository GetByIDs of company")
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
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository GetByIDs build sql string of company")
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository GetByIDs selectBuilder error of company")
		return companies, err
	}
	output := make(chan []models.Company, 1)
	errors := hystrix.Go("companyGetByIDs", func() error {
		stmt, err := r.db.PrepareContext(ctx, sql)
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository GetByIDs PrepareContext error of company")
			return err
		}
		rows, err := stmt.QueryContext(ctx, args...)
		defer rows.Close()
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository GetByIDs QueryContext error of company")
			return err
		}
		for rows.Next() {
			company := models.Company{}
			err := r.scanCompany(rows, &company)
			if err != nil {
				logger.WithFields(logger.Fields{
					"traceId": ctx.Value(utils.TraceIDKey),
					"Error":   err,
				}).Errorf("Repository GetByIDs Scan error of company")
				return err
			}
			companies = append(companies, company)
		}
		output <- companies
		return nil
	}, nil)

	select {
	case err := <-errors:
		return companies, err
	case out := <-output:
		return out, nil
	}
}

// setArgsToListSelectBuilder ...
func (r *RepositoryImpl) setArgsToListSelectBuilder(ctx context.Context, selectBuilder sq.SelectBuilder, params arguments.CompanyList) sq.SelectBuilder {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Repository setArgsToListSelectBuilder of company")
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
	if params.LastID != 0 {
		selectBuilder = selectBuilder.Where(sq.Gt{"id": params.LastID})
	}
	return selectBuilder
}

// List ...
func (r *RepositoryImpl) List(ctx context.Context, params arguments.CompanyList) ([]models.Company, error) {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Repository List of company")
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
	selectBuilderWithArgs := r.setArgsToListSelectBuilder(ctx, selectBuilder, params)
	sql, args, err := selectBuilderWithArgs.ToSql()
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository List build sql string of company")
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository List selectBuilderWithArgs error of company")
		return companies, err
	}
	output := make(chan []models.Company, 1)
	errors := hystrix.Go("companyList", func() error {
		stmt, err := r.db.PrepareContext(ctx, sql)
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository List PrepareContext error of company")
			return err
		}
		rows, err := stmt.QueryContext(ctx, args...)
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository List QueryContext error of company")
			return err
		}
		defer rows.Close()
		for rows.Next() {
			company := models.Company{}
			err := r.scanCompany(rows, &company)
			if err != nil {
				logger.WithFields(logger.Fields{
					"traceId": ctx.Value(utils.TraceIDKey),
					"Error":   err,
				}).Errorf("Repository List Scan error of company")
				return err
			}
			companies = append(companies, company)
		}
		output <- companies
		return nil
	}, nil)

	select {
	case err := <-errors:
		return companies, err
	case out := <-output:
		return out, nil
	}
}

// setArgsToCountSelectBuilder ...
func (r *RepositoryImpl) setArgsToCountSelectBuilder(ctx context.Context, selectBuilder sq.SelectBuilder, params arguments.CompanyCount) sq.SelectBuilder {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Repository setArgsToCountSelectBuilder of company")
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
func (r *RepositoryImpl) Count(ctx context.Context, params arguments.CompanyCount) (int64, error) {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Repository Count of company")
	var (
		count         int64
		selectBuilder = sq.Select("COUNT(id)").From("company")
	)
	selectBuilderWithArgs := r.setArgsToCountSelectBuilder(ctx, selectBuilder, params)
	sql, args, err := selectBuilderWithArgs.ToSql()
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository Count build sql string of company")
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository Count selectBuilderWithArgs error of company")
		return count, err
	}
	output := make(chan int64, 1)
	errors := hystrix.Go("companyCount", func() error {
		stmt, err := r.db.PrepareContext(ctx, sql)
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository Count PrepareContext error of company")
			return err
		}
		row := stmt.QueryRowContext(ctx, args...)
		err = row.Scan(&count)
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository Count Scan error of company")
			return err
		}
		output <- count
		return nil
	}, nil)

	select {
	case err := <-errors:
		return count, err
	case out := <-output:
		return out, nil
	}
}

// Insert ...
func (r *RepositoryImpl) Insert(ctx context.Context, params arguments.CompanyInsert) (models.Company, error) {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Repository Insert of company")
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
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository Insert build sql string of company")
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository Insert insertBuilder error of company")
		return company, err
	}
	output := make(chan models.Company, 1)
	errors := hystrix.Go("companyInsert", func() error {
		stmt, err := r.db.PrepareContext(ctx, sql)
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository Insert PrepareContext error of company")
			return err
		}
		row, err := stmt.ExecContext(ctx, args...)
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository Insert ExecContext error of company")
			return err
		}
		id, err := row.LastInsertId()
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository Insert LastInsertId error of company")
			return err
		}
		newCompany, err := r.GetByID(ctx, arguments.CompanyGetByID{ID: id})
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository Insert GetByID error of company")
			return err
		}
		output <- newCompany
		return nil
	}, nil)

	select {
	case err := <-errors:
		return company, err
	case out := <-output:
		return out, nil
	}
}

// setArgsToUpdateBuilder ...
func (r *RepositoryImpl) setArgsToUpdateBuilder(ctx context.Context, updateBuilder sq.UpdateBuilder, params arguments.CompanyUpdate) sq.UpdateBuilder {
	logger.WithFields(logger.Fields{
		"params": params,
	}).Infof("Repository setArgsToUpdateBuilder of company")
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
func (r *RepositoryImpl) Update(ctx context.Context, params arguments.CompanyUpdate) (models.Company, error) {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Repository Update of company")
	var (
		company       models.Company
		updateBuilder = sq.Update("company").Where(sq.Eq{"id": *params.ID})
	)
	updateBuilderWithArgs := r.setArgsToUpdateBuilder(ctx, updateBuilder, params)

	sql, args, err := updateBuilderWithArgs.ToSql()
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository Update build sql string of company")
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository Update updateBuilderWithArgs error of company")
		return company, err
	}
	output := make(chan models.Company, 1)
	errors := hystrix.Go("companyUpdate", func() error {
		stmt, err := r.db.PrepareContext(ctx, sql)
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository Update PrepareContext error of company")
			return err
		}
		result, err := stmt.ExecContext(ctx, args...)
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository Update ExecContext error of company")
			return err
		}
		rowAffected, err := result.RowsAffected()
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository Update RowsAffected error of company")
			return err
		}
		if rowAffected <= 0 {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   fmt.Errorf("error when update record id %d", *params.ID),
			}).Errorf("Repository Update rowAffected <= 0 of company")
			return fmt.Errorf("error when update record id %d", *params.ID)
		}
		newCompany, err := r.GetByID(ctx, arguments.CompanyGetByID{ID: *params.ID})
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository Update GetByID error of company")
			return err
		}
		output <- newCompany
		return nil
	}, nil)

	select {
	case err := <-errors:
		return company, err
	case out := <-output:
		return out, nil
	}
}

// Delete ...
func (r *RepositoryImpl) Delete(ctx context.Context, params arguments.CompanyDelete) (int64, error) {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Repository Delete of company")
	var (
		id            int64
		deleteBuilder = sq.Delete("company").Where(sq.Eq{"id": params.ID})
	)
	sql, args, err := deleteBuilder.ToSql()
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository Delete build sql string of company")
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository Delete deleteBuilder error of company")
		return id, err
	}
	output := make(chan int64, 1)
	errors := hystrix.Go("companyDelete", func() error {
		stmt, err := r.db.PrepareContext(ctx, sql)
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository Delete PrepareContext error of company")
			return err
		}
		result, err := stmt.ExecContext(ctx, args...)
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository Delete ExecContext error of company")
			return err
		}
		rowAffected, err := result.RowsAffected()
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository Delete RowsAffected error of company")
			return err
		}
		if rowAffected <= 0 {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   fmt.Errorf("not found record by id %d", params.ID),
			}).Errorf("Repository Update rowAffected <= 0 of company")
			return fmt.Errorf("not found record by id %d", params.ID)
		}
		output <- params.ID
		return nil
	}, nil)

	select {
	case err := <-errors:
		return id, err
	case out := <-output:
		return out, nil
	}
}

//go:generate mockery -name=IDatabase -output=mocks -case=underscore
