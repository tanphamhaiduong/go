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
	log.WithField("params", params).Info("user of Repository GetByID")
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
	}).Info("user of Repository GetByID build sql string")
	if err != nil {
		log.WithField("Error", err).Error("user of Repository GetByID selectBuilder error")
		return user, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("user of Repository GetByID PrepareContext error")
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
		log.WithField("Error", err).Error("user of Repository GetByID Scan error")
		return user, err
	}
	return user, nil
}

// GetByIDs ...
func (r *RepositoryImpl) GetByIDs(ctx context.Context, params arguments.UserGetByIDsArgs) ([]models.User, error) {
	log.WithField("params", params).Info("user of Repository GetByIDs")
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
	}).Info("user of Repository GetByIDs build sql string")
	if err != nil {
		log.WithField("Error", err).Error("user of Repository GetByIDs selectBuilder error")
		return users, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("user of Repository GetByIDs PrepareContext error")
		return users, err
	}
	rows, err := stmt.QueryContext(ctx, args...)
	defer rows.Close()
	if err != nil {
		log.WithField("Error", err).Error("user of Repository GetByIDs QueryContext error")
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
			log.WithField("Error", err).Error("user of Repository GetByIDs Scan error")
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

// setArgsToListSelectBuilder ...
func (r *RepositoryImpl) setArgsToListSelectBuilder(selectBuilder sq.SelectBuilder, params arguments.UserListArgs) sq.SelectBuilder {
	log.WithField("params", params).Info("user of Repository setArgsToListSelectBuilder")
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
	log.WithField("params", params).Info("user of Repository List")
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
	}).Info("user of Repository List build sql string")
	if err != nil {
		log.WithField("Error", err).Error("user of Repository List selectBuilderWithArgs error")
		return users, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("user of Repository List PrepareContext error")
		return users, err
	}
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		log.WithField("Error", err).Error("user of Repository List QueryContext error")
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
			log.WithField("Error", err).Error("user of Repository List Scan error")
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

// setArgsToCountSelectBuilder ...
func (r *RepositoryImpl) setArgsToCountSelectBuilder(selectBuilder sq.SelectBuilder, params arguments.UserCountArgs) sq.SelectBuilder {
	log.WithField("params", params).Info("user of Repository setArgsToCountSelectBuilder")
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
	log.WithField("params", params).Info("user of Repository Count")
	var (
		count         int64
		selectBuilder = sq.Select("COUNT(id)").From("user")
	)
	selectBuilderWithArgs := r.setArgsToCountSelectBuilder(selectBuilder, params)
	sql, args, err := selectBuilderWithArgs.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("user of Repository Count build sql string")
	if err != nil {
		log.WithField("Error", err).Error("user of Repository Count selectBuilderWithArgs error")
		return count, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("user of Repository Count PrepareContext error")
		return count, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	err = row.Scan(&count)
	if err != nil {
		log.WithField("Error", err).Error("user of Repository Count Scan error")
		return count, err
	}
	return count, nil
}

// Insert ...
func (r *RepositoryImpl) Insert(ctx context.Context, params arguments.UserInsertArgs) (models.User, error) {
	log.WithField("params", params).Info("user of Repository Insert")
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
	}).Info("user of Repository Insert build sql string")
	if err != nil {
		log.WithField("Error", err).Error("user of Repository Insert insertBuilder error")
		return user, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("user of Repository Insert PrepareContext error")
		return user, err
	}
	row, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		log.WithField("Error", err).Error("user of Repository Insert ExecContext error")
		return user, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		log.WithField("Error", err).Error("user of Repository Insert LastInsertId error")
		return user, err
	}
	newUser, err := r.GetByID(ctx, arguments.UserGetByIDArgs{ID: id})
	if err != nil {
		log.WithField("Error", err).Error("user of Repository Insert GetByID error")
		return user, err
	}
	return newUser, nil
}

// setArgsToUpdateBuilder ...
func (r *RepositoryImpl) setArgsToUpdateBuilder(updateBuilder sq.UpdateBuilder, params arguments.UserUpdateArgs) sq.UpdateBuilder {
	log.WithField("params", params).Info("user of Repository setArgsToUpdateBuilder")
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
	log.WithField("params", params).Info("user of Repository Update")
	var (
		user          models.User
		updateBuilder = sq.Update("user").Where(sq.Eq{"id": *params.ID})
	)
	updateBuilderWithArgs := r.setArgsToUpdateBuilder(updateBuilder, params)

	sql, args, err := updateBuilderWithArgs.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("user of Repository Update build sql string")
	if err != nil {
		log.WithField("Error", err).Error("user of Repository Update updateBuilderWithArgs error")
		return user, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("user of Repository Update PrepareContext error")
		return user, err
	}
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		log.WithField("Error", err).Error("user of Repository Update ExecContext error")
		return user, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		log.WithField("Error", err).Error("user of Repository Update RowsAffected error")
		return user, err
	}
	if rowAffected <= 0 {
		log.WithField("Error", fmt.Errorf("error when update record id %d", *params.ID)).Error("user of Repository Update rowAffected <= 0")
		return user, fmt.Errorf("error when update record id %d", *params.ID)
	}
	newUser, err := r.GetByID(ctx, arguments.UserGetByIDArgs{ID: *params.ID})
	if err != nil {
		log.WithField("Error", err).Error("user of Repository Update GetByID error")
		return user, err
	}
	return newUser, nil
}

// Delete ...
func (r *RepositoryImpl) Delete(ctx context.Context, params arguments.UserDeleteArgs) (int64, error) {
	log.WithField("params", params).Info("user of Repository Delete")
	var (
		id            int64
		deleteBuilder = sq.Delete("user").Where(sq.Eq{"id": params.ID})
	)
	sql, args, err := deleteBuilder.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("user of Repository Delete build sql string")
	if err != nil {
		log.WithField("Error", err).Error("user of Repository Delete deleteBuilder error")
		return id, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("user of Repository Delete PrepareContext error")
		return id, err
	}
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		log.WithField("Error", err).Error("user of Repository Delete ExecContext error")
		return id, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		log.WithField("Error", err).Error("user of Repository Delete RowsAffected error")
		return id, err
	}
	if rowAffected <= 0 {
		log.WithField("Error", fmt.Errorf("not found record by id %d", params.ID)).Error("user of Repository Update rowAffected <= 0")
		return id, fmt.Errorf("not found record by id %d", params.ID)
	}
	return params.ID, nil
}

//go:generate mockery -name=IDatabase -output=mocks -case=underscore
