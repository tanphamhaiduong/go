// @generated
package role

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

//scanRole
func (r *RepositoryImpl) scanRole(row database.IRow, role *models.Role) error {
	err := row.Scan(
		&role.ID,
		&role.Name,
		&role.CompanyID,
		&role.Status,
		&role.CreatedBy,
		&role.UpdatedBy,
	)
	return err
}

// GetByID ...
func (r *RepositoryImpl) GetByID(ctx context.Context, params arguments.RoleGetByID) (models.Role, error) {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Repository GetByID of role")
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
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository GetByID build sql string of role")
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository GetByID selectBuilder error of role")
		return role, err
	}
	output := make(chan models.Role, 1)
	errors := hystrix.Go("roleGetByID", func() error {
		stmt, err := r.db.PrepareContext(ctx, sql)
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository GetByID PrepareContext error of role")
			return err
		}
		row := stmt.QueryRowContext(ctx, args...)
		err = r.scanRole(row, &role)
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository GetByID Scan error of role")
			return err
		}
		output <- role
		return nil
	}, nil)

	select {
	case err := <-errors:
		return role, err
	case out := <-output:
		return out, nil
	}
}

// GetByIDs ...
func (r *RepositoryImpl) GetByIDs(ctx context.Context, params arguments.RoleGetByIDs) ([]models.Role, error) {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Repository GetByIDs of role")
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
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository GetByIDs build sql string of role")
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository GetByIDs selectBuilder error of role")
		return roles, err
	}
	output := make(chan []models.Role, 1)
	errors := hystrix.Go("roleGetByIDs", func() error {
		stmt, err := r.db.PrepareContext(ctx, sql)
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository GetByIDs PrepareContext error of role")
			return err
		}
		rows, err := stmt.QueryContext(ctx, args...)
		defer rows.Close()
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository GetByIDs QueryContext error of role")
			return err
		}
		for rows.Next() {
			role := models.Role{}
			err := r.scanRole(rows, &role)
			if err != nil {
				logger.WithFields(logger.Fields{
					"traceId": ctx.Value(utils.TraceIDKey),
					"Error":   err,
				}).Errorf("Repository GetByIDs Scan error of role")
				return err
			}
			roles = append(roles, role)
		}
		output <- roles
		return nil
	}, nil)

	select {
	case err := <-errors:
		return roles, err
	case out := <-output:
		return out, nil
	}
}

// setArgsToListSelectBuilder ...
func (r *RepositoryImpl) setArgsToListSelectBuilder(ctx context.Context, selectBuilder sq.SelectBuilder, params arguments.RoleList) sq.SelectBuilder {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Repository setArgsToListSelectBuilder of role")
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
	if params.BeginID != 0 {
		selectBuilder = selectBuilder.Where(sq.Gt{"id": params.BeginID})
	}
	if params.EndID != 0 {
		selectBuilder = selectBuilder.Where(sq.Lt{"id": params.EndID})
	}
	if params.PageSize != 0 {
		selectBuilder = selectBuilder.Limit(uint64(params.PageSize))
	}
	return selectBuilder
}

// List ...
func (r *RepositoryImpl) List(ctx context.Context, params arguments.RoleList) ([]models.Role, error) {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Repository List of role")
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
	selectBuilderWithArgs := r.setArgsToListSelectBuilder(ctx, selectBuilder, params)
	sql, args, err := selectBuilderWithArgs.ToSql()
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository List build sql string of role")
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository List selectBuilderWithArgs error of role")
		return roles, err
	}
	output := make(chan []models.Role, 1)
	errors := hystrix.Go("roleList", func() error {
		stmt, err := r.db.PrepareContext(ctx, sql)
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository List PrepareContext error of role")
			return err
		}
		rows, err := stmt.QueryContext(ctx, args...)
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository List QueryContext error of role")
			return err
		}
		defer rows.Close()
		for rows.Next() {
			role := models.Role{}
			err := r.scanRole(rows, &role)
			if err != nil {
				logger.WithFields(logger.Fields{
					"traceId": ctx.Value(utils.TraceIDKey),
					"Error":   err,
				}).Errorf("Repository List Scan error of role")
				return err
			}
			roles = append(roles, role)
		}
		output <- roles
		return nil
	}, nil)

	select {
	case err := <-errors:
		return roles, err
	case out := <-output:
		return out, nil
	}
}

// setArgsToCountSelectBuilder ...
func (r *RepositoryImpl) setArgsToCountSelectBuilder(ctx context.Context, selectBuilder sq.SelectBuilder, params arguments.RoleCount) sq.SelectBuilder {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Repository setArgsToCountSelectBuilder of role")
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
func (r *RepositoryImpl) Count(ctx context.Context, params arguments.RoleCount) (int64, error) {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Repository Count of role")
	var (
		count         int64
		selectBuilder = sq.Select("COUNT(id)").From("role")
	)
	selectBuilderWithArgs := r.setArgsToCountSelectBuilder(ctx, selectBuilder, params)
	sql, args, err := selectBuilderWithArgs.ToSql()
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository Count build sql string of role")
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository Count selectBuilderWithArgs error of role")
		return count, err
	}
	output := make(chan int64, 1)
	errors := hystrix.Go("roleCount", func() error {
		stmt, err := r.db.PrepareContext(ctx, sql)
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository Count PrepareContext error of role")
			return err
		}
		row := stmt.QueryRowContext(ctx, args...)
		err = row.Scan(&count)
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository Count Scan error of role")
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
func (r *RepositoryImpl) Insert(ctx context.Context, params arguments.RoleInsert) (models.Role, error) {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Repository Insert of role")
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
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository Insert build sql string of role")
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository Insert insertBuilder error of role")
		return role, err
	}
	output := make(chan models.Role, 1)
	errors := hystrix.Go("roleInsert", func() error {
		stmt, err := r.db.PrepareContext(ctx, sql)
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository Insert PrepareContext error of role")
			return err
		}
		row, err := stmt.ExecContext(ctx, args...)
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository Insert ExecContext error of role")
			return err
		}
		id, err := row.LastInsertId()
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository Insert LastInsertId error of role")
			return err
		}
		newRole, err := r.GetByID(ctx, arguments.RoleGetByID{ID: id})
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository Insert GetByID error of role")
			return err
		}
		output <- newRole
		return nil
	}, nil)

	select {
	case err := <-errors:
		return role, err
	case out := <-output:
		return out, nil
	}
}

// setArgsToUpdateBuilder ...
func (r *RepositoryImpl) setArgsToUpdateBuilder(ctx context.Context, updateBuilder sq.UpdateBuilder, params arguments.RoleUpdate) sq.UpdateBuilder {
	logger.WithFields(logger.Fields{
		"params": params,
	}).Infof("Repository setArgsToUpdateBuilder of role")
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
func (r *RepositoryImpl) Update(ctx context.Context, params arguments.RoleUpdate) (models.Role, error) {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Repository Update of role")
	var (
		role          models.Role
		updateBuilder = sq.Update("role").Where(sq.Eq{"id": *params.ID})
	)
	updateBuilderWithArgs := r.setArgsToUpdateBuilder(ctx, updateBuilder, params)

	sql, args, err := updateBuilderWithArgs.ToSql()
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository Update build sql string of role")
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository Update updateBuilderWithArgs error of role")
		return role, err
	}
	output := make(chan models.Role, 1)
	errors := hystrix.Go("roleUpdate", func() error {
		stmt, err := r.db.PrepareContext(ctx, sql)
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository Update PrepareContext error of role")
			return err
		}
		result, err := stmt.ExecContext(ctx, args...)
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository Update ExecContext error of role")
			return err
		}
		rowAffected, err := result.RowsAffected()
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository Update RowsAffected error of role")
			return err
		}
		if rowAffected <= 0 {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   fmt.Errorf("error when update record id %d", *params.ID),
			}).Errorf("Repository Update rowAffected <= 0 of role")
			return fmt.Errorf("error when update record id %d", *params.ID)
		}
		newRole, err := r.GetByID(ctx, arguments.RoleGetByID{ID: *params.ID})
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository Update GetByID error of role")
			return err
		}
		output <- newRole
		return nil
	}, nil)

	select {
	case err := <-errors:
		return role, err
	case out := <-output:
		return out, nil
	}
}

// Delete ...
func (r *RepositoryImpl) Delete(ctx context.Context, params arguments.RoleDelete) (int64, error) {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Repository Delete of role")
	var (
		id            int64
		deleteBuilder = sq.Delete("role").Where(sq.Eq{"id": params.ID})
	)
	sql, args, err := deleteBuilder.ToSql()
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository Delete build sql string of role")
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository Delete deleteBuilder error of role")
		return id, err
	}
	output := make(chan int64, 1)
	errors := hystrix.Go("roleDelete", func() error {
		stmt, err := r.db.PrepareContext(ctx, sql)
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository Delete PrepareContext error of role")
			return err
		}
		result, err := stmt.ExecContext(ctx, args...)
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository Delete ExecContext error of role")
			return err
		}
		rowAffected, err := result.RowsAffected()
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository Delete RowsAffected error of role")
			return err
		}
		if rowAffected <= 0 {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   fmt.Errorf("not found record by id %d", params.ID),
			}).Errorf("Repository Update rowAffected <= 0 of role")
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
