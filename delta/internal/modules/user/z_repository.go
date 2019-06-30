// @generated
package user

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
func (r *RepositoryImpl) GetByID(ctx context.Context, params arguments.UserGetByIDArgs) (models.User, error) {
	log.WithField("params", params).Info("Repository GetByID of user")
	var (
		user          models.User
		selectBuilder = sq.Select(
			"id",
			"email",
			"name",
			"company_id",
			"status",
			"created_by",
			"updated_by",
		).From("user").Where(sq.Eq{"id": params.ID})
	)
	sql, args, err := selectBuilder.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository GetByID build sql string of user")
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByID selectBuilder error of user")
		return user, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByID PrepareContext error of user")
		return user, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	err = row.Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.CompanyID,
		&user.Status,
		&user.CreatedBy,
		&user.UpdatedBy,
	)
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByID Scan error of user")
		return user, err
	}
	return user, nil
}

// GetByIDs ...
func (r *RepositoryImpl) GetByIDs(ctx context.Context, params arguments.UserGetByIDsArgs) ([]models.User, error) {
	log.WithField("params", params).Info("Repository GetByIDs of user")
	var (
		users         []models.User
		selectBuilder = sq.Select(
			"id",
			"email",
			"name",
			"company_id",
			"status",
			"created_by",
			"updated_by",
		).From("user").Where(sq.Eq{"id": params.IDs})
	)
	sql, args, err := selectBuilder.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository GetByIDs build sql string of user")
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByIDs selectBuilder error of user")
		return users, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByIDs PrepareContext error of user")
		return users, err
	}
	rows, err := stmt.QueryContext(ctx, args...)
	defer rows.Close()
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByIDs QueryContext error of user")
		return users, err
	}
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Name,
			&user.CompanyID,
			&user.Status,
			&user.CreatedBy,
			&user.UpdatedBy,
		)
		if err != nil {
			log.WithField("Error", err).Error("Repository GetByIDs Scan error of user")
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

// setArgsToListSelectBuilder ...
func (r *RepositoryImpl) setArgsToListSelectBuilder(selectBuilder sq.SelectBuilder, params arguments.UserListArgs) sq.SelectBuilder {
	log.WithField("params", params).Info("Repository setArgsToListSelectBuilder of user")
	if params.ID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"id": params.ID})
	}
	if params.Email != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"email": params.Email})
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
func (r *RepositoryImpl) List(ctx context.Context, params arguments.UserListArgs) ([]models.User, error) {
	log.WithField("params", params).Info("Repository List of user")
	var (
		users         []models.User
		selectBuilder = sq.Select(
			"id",
			"email",
			"name",
			"company_id",
			"status",
			"created_by",
			"updated_by",
		).From("user")
	)
	selectBuilderWithArgs := r.setArgsToListSelectBuilder(selectBuilder, params)
	sql, args, err := selectBuilderWithArgs.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository List build sql string of user")
	if err != nil {
		log.WithField("Error", err).Error("Repository List selectBuilderWithArgs error of user")
		return users, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository List PrepareContext error of user")
		return users, err
	}
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		log.WithField("Error", err).Error("Repository List QueryContext error of user")
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Name,
			&user.CompanyID,
			&user.Status,
			&user.CreatedBy,
			&user.UpdatedBy,
		)
		if err != nil {
			log.WithField("Error", err).Error("Repository List Scan error of user")
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

// setArgsToCountSelectBuilder ...
func (r *RepositoryImpl) setArgsToCountSelectBuilder(selectBuilder sq.SelectBuilder, params arguments.UserCountArgs) sq.SelectBuilder {
	log.WithField("params", params).Info("Repository setArgsToCountSelectBuilder of user")
	if params.ID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"id": params.ID})
	}
	if params.Email != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"email": params.Email})
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
func (r *RepositoryImpl) Count(ctx context.Context, params arguments.UserCountArgs) (int64, error) {
	log.WithField("params", params).Info("Repository Count of user")
	var (
		count         int64
		selectBuilder = sq.Select("COUNT(id)").From("user")
	)
	selectBuilderWithArgs := r.setArgsToCountSelectBuilder(selectBuilder, params)
	sql, args, err := selectBuilderWithArgs.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository Count build sql string of user")
	if err != nil {
		log.WithField("Error", err).Error("Repository Count selectBuilderWithArgs error of user")
		return count, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository Count PrepareContext error of user")
		return count, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	err = row.Scan(&count)
	if err != nil {
		log.WithField("Error", err).Error("Repository Count Scan error of user")
		return count, err
	}
	return count, nil
}

// Insert ...
func (r *RepositoryImpl) Insert(ctx context.Context, params arguments.UserInsertArgs) (models.User, error) {
	log.WithField("params", params).Info("Repository Insert of user")
	var (
		user          models.User
		insertBuilder = sq.Insert("user").Columns(
			"email",
			"name",
			"company_id",
			"status",
			"created_by",
			"updated_by",
		).Values(
			params.Email,
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
	}).Info("Repository Insert build sql string of user")
	if err != nil {
		log.WithField("Error", err).Error("Repository Insert insertBuilder error of user")
		return user, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository Insert PrepareContext error of user")
		return user, err
	}
	row, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		log.WithField("Error", err).Error("Repository Insert ExecContext error of user")
		return user, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		log.WithField("Error", err).Error("Repository Insert LastInsertId error of user")
		return user, err
	}
	newUser, err := r.GetByID(ctx, arguments.UserGetByIDArgs{ID: id})
	if err != nil {
		log.WithField("Error", err).Error("Repository Insert GetByID error of user")
		return user, err
	}
	return newUser, nil
}

// setArgsToUpdateBuilder ...
func (r *RepositoryImpl) setArgsToUpdateBuilder(updateBuilder sq.UpdateBuilder, params arguments.UserUpdateArgs) sq.UpdateBuilder {
	log.WithField("params", params).Info("Repository setArgsToUpdateBuilder of user")
	if params.Email != nil {
		updateBuilder = updateBuilder.Set("email", *params.Email)
	}

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
func (r *RepositoryImpl) Update(ctx context.Context, params arguments.UserUpdateArgs) (models.User, error) {
	log.WithField("params", params).Info("Repository Update of user")
	var (
		user          models.User
		updateBuilder = sq.Update("user").Where(sq.Eq{"id": *params.ID})
	)
	updateBuilderWithArgs := r.setArgsToUpdateBuilder(updateBuilder, params)

	sql, args, err := updateBuilderWithArgs.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository Update build sql string of user")
	if err != nil {
		log.WithField("Error", err).Error("Repository Update updateBuilderWithArgs error of user")
		return user, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository Update PrepareContext error of user")
		return user, err
	}
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		log.WithField("Error", err).Error("Repository Update ExecContext error of user")
		return user, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		log.WithField("Error", err).Error("Repository Update RowsAffected error of user")
		return user, err
	}
	if rowAffected <= 0 {
		log.WithField("Error", fmt.Errorf("error when update record id %d", *params.ID)).Error("Repository Update rowAffected <= 0 of user")
		return user, fmt.Errorf("error when update record id %d", *params.ID)
	}
	newUser, err := r.GetByID(ctx, arguments.UserGetByIDArgs{ID: *params.ID})
	if err != nil {
		log.WithField("Error", err).Error("Repository Update GetByID error of user")
		return user, err
	}
	return newUser, nil
}

// Delete ...
func (r *RepositoryImpl) Delete(ctx context.Context, params arguments.UserDeleteArgs) (int64, error) {
	log.WithField("params", params).Info("Repository Delete of user")
	var (
		id            int64
		deleteBuilder = sq.Delete("user").Where(sq.Eq{"id": params.ID})
	)
	sql, args, err := deleteBuilder.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository Delete build sql string of user")
	if err != nil {
		log.WithField("Error", err).Error("Repository Delete deleteBuilder error of user")
		return id, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository Delete PrepareContext error of user")
		return id, err
	}
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		log.WithField("Error", err).Error("Repository Delete ExecContext error of user")
		return id, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		log.WithField("Error", err).Error("Repository Delete RowsAffected error of user")
		return id, err
	}
	if rowAffected <= 0 {
		log.WithField("Error", fmt.Errorf("not found record by id %d", params.ID)).Error("Repository Update rowAffected <= 0 of user")
		return id, fmt.Errorf("not found record by id %d", params.ID)
	}
	return params.ID, nil
}

//go:generate mockery -name=IDatabase -output=mocks -case=underscore
