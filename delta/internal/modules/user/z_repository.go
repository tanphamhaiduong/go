// @generated
package user

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

//scanUser
func (r *RepositoryImpl) scanUser(row database.IRow, user *models.User) error {
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Name,
		&user.DateOfBirth,
		&user.Reference,
		&user.AvatarUrl,
		&user.LicenseNumber,
		&user.PhoneNumber,
		&user.Extension,
		&user.TelProvider,
		&user.TelApi,
		&user.SupervisorID,
		&user.RoleID,
		&user.CompanyID,
		&user.Status,
		&user.CreatedBy,
		&user.UpdatedBy,
	)
	return err
}

// GetByID ...
func (r *RepositoryImpl) GetByID(ctx context.Context, params arguments.UserGetByID) (models.User, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Repository GetByID of user")
	var (
		user          models.User
		selectBuilder = sq.Select(
			"id",
			"username",
			"password",
			"name",
			"date_of_birth",
			"reference",
			"avatar_url",
			"license_number",
			"phone_number",
			"extension",
			"tel_provider",
			"tel_api",
			"supervisor_id",
			"role_id",
			"company_id",
			"status",
			"created_by",
			"updated_by",
		).From("user").Where(sq.Eq{"id": params.ID})
	)
	sql, args, err := selectBuilder.ToSql()
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository GetByID build sql string of user")
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository GetByID selectBuilder error of user")
		return user, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository GetByID PrepareContext error of user")
		return user, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	err = r.scanUser(row, &user)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository GetByID Scan error of user")
		return user, err
	}
	return user, nil
}

// GetByIDs ...
func (r *RepositoryImpl) GetByIDs(ctx context.Context, params arguments.UserGetByIDs) ([]models.User, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Repository GetByIDs of user")
	var (
		users         []models.User
		selectBuilder = sq.Select(
			"id",
			"username",
			"password",
			"name",
			"date_of_birth",
			"reference",
			"avatar_url",
			"license_number",
			"phone_number",
			"extension",
			"tel_provider",
			"tel_api",
			"supervisor_id",
			"role_id",
			"company_id",
			"status",
			"created_by",
			"updated_by",
		).From("user").Where(sq.Eq{"id": params.IDs})
	)
	sql, args, err := selectBuilder.ToSql()
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository GetByIDs build sql string of user")
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository GetByIDs selectBuilder error of user")
		return users, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository GetByIDs PrepareContext error of user")
		return users, err
	}
	rows, err := stmt.QueryContext(ctx, args...)
	defer rows.Close()
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository GetByIDs QueryContext error of user")
		return users, err
	}
	for rows.Next() {
		user := models.User{}
		err := r.scanUser(rows, &user)
		if err != nil {
			logger.WithFields(logger.Fields{
				"TraceID": ctx.Value("TraceID"),
				"Error":   err,
			}).Errorf("Repository GetByIDs Scan error of user")
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

// setArgsToListSelectBuilder ...
func (r *RepositoryImpl) setArgsToListSelectBuilder(ctx context.Context, selectBuilder sq.SelectBuilder, params arguments.UserList) sq.SelectBuilder {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Repository setArgsToListSelectBuilder of user")
	if params.ID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"id": params.ID})
	}
	if params.Username != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"username": params.Username})
	}
	if params.Password != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"password": params.Password})
	}
	if params.Name != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"name": params.Name})
	}
	if value, _ := params.DateOfBirth.Value(); value != nil {
		selectBuilder = selectBuilder.Where(sq.Eq{"date_of_birth": params.DateOfBirth})
	}
	if params.Reference != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"reference": params.Reference})
	}
	if params.AvatarUrl != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"avatar_url": params.AvatarUrl})
	}
	if params.LicenseNumber != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"license_number": params.LicenseNumber})
	}
	if params.PhoneNumber != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"phone_number": params.PhoneNumber})
	}
	if params.Extension != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"extension": params.Extension})
	}
	if params.TelProvider != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"tel_provider": params.TelProvider})
	}
	if params.TelApi != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"tel_api": params.TelApi})
	}
	if params.SupervisorID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"supervisor_id": params.SupervisorID})
	}
	if params.RoleID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"role_id": params.RoleID})
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
func (r *RepositoryImpl) List(ctx context.Context, params arguments.UserList) ([]models.User, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Repository List of user")
	var (
		users         []models.User
		selectBuilder = sq.Select(
			"id",
			"username",
			"password",
			"name",
			"date_of_birth",
			"reference",
			"avatar_url",
			"license_number",
			"phone_number",
			"extension",
			"tel_provider",
			"tel_api",
			"supervisor_id",
			"role_id",
			"company_id",
			"status",
			"created_by",
			"updated_by",
		).From("user")
	)
	selectBuilderWithArgs := r.setArgsToListSelectBuilder(ctx, selectBuilder, params)
	sql, args, err := selectBuilderWithArgs.ToSql()
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository List build sql string of user")
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository List selectBuilderWithArgs error of user")
		return users, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository List PrepareContext error of user")
		return users, err
	}
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository List QueryContext error of user")
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		user := models.User{}
		err := r.scanUser(rows, &user)
		if err != nil {
			logger.WithFields(logger.Fields{
				"TraceID": ctx.Value("TraceID"),
				"Error":   err,
			}).Errorf("Repository List Scan error of user")
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

// setArgsToCountSelectBuilder ...
func (r *RepositoryImpl) setArgsToCountSelectBuilder(ctx context.Context, selectBuilder sq.SelectBuilder, params arguments.UserCount) sq.SelectBuilder {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Repository setArgsToCountSelectBuilder of user")
	if params.ID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"id": params.ID})
	}
	if params.Username != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"username": params.Username})
	}
	if params.Password != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"password": params.Password})
	}
	if params.Name != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"name": params.Name})
	}
	if value, _ := params.DateOfBirth.Value(); value != nil {
		selectBuilder = selectBuilder.Where(sq.Eq{"date_of_birth": params.DateOfBirth})
	}
	if params.Reference != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"reference": params.Reference})
	}
	if params.AvatarUrl != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"avatar_url": params.AvatarUrl})
	}
	if params.LicenseNumber != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"license_number": params.LicenseNumber})
	}
	if params.PhoneNumber != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"phone_number": params.PhoneNumber})
	}
	if params.Extension != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"extension": params.Extension})
	}
	if params.TelProvider != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"tel_provider": params.TelProvider})
	}
	if params.TelApi != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"tel_api": params.TelApi})
	}
	if params.SupervisorID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"supervisor_id": params.SupervisorID})
	}
	if params.RoleID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"role_id": params.RoleID})
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
func (r *RepositoryImpl) Count(ctx context.Context, params arguments.UserCount) (int64, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Repository Count of user")
	var (
		count         int64
		selectBuilder = sq.Select("COUNT(id)").From("user")
	)
	selectBuilderWithArgs := r.setArgsToCountSelectBuilder(ctx, selectBuilder, params)
	sql, args, err := selectBuilderWithArgs.ToSql()
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository Count build sql string of user")
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Count selectBuilderWithArgs error of user")
		return count, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Count PrepareContext error of user")
		return count, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	err = row.Scan(&count)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Count Scan error of user")
		return count, err
	}
	return count, nil
}

// Insert ...
func (r *RepositoryImpl) Insert(ctx context.Context, params arguments.UserInsert) (models.User, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Repository Insert of user")
	var (
		user          models.User
		insertBuilder = sq.Insert("user").Columns(
			"username",
			"password",
			"name",
			"date_of_birth",
			"reference",
			"avatar_url",
			"license_number",
			"phone_number",
			"extension",
			"tel_provider",
			"tel_api",
			"supervisor_id",
			"role_id",
			"company_id",
			"status",
			"created_by",
			"updated_by",
		).Values(
			params.Username,
			params.Password,
			params.Name,
			params.DateOfBirth,
			params.Reference,
			params.AvatarUrl,
			params.LicenseNumber,
			params.PhoneNumber,
			params.Extension,
			params.TelProvider,
			params.TelApi,
			params.SupervisorID,
			params.RoleID,
			params.CompanyID,
			params.Status,
			params.CreatedBy,
			params.UpdatedBy,
		)
	)
	sql, args, err := insertBuilder.ToSql()
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository Insert build sql string of user")
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Insert insertBuilder error of user")
		return user, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Insert PrepareContext error of user")
		return user, err
	}
	row, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Insert ExecContext error of user")
		return user, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Insert LastInsertId error of user")
		return user, err
	}
	newUser, err := r.GetByID(ctx, arguments.UserGetByID{ID: id})
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Insert GetByID error of user")
		return user, err
	}
	return newUser, nil
}

// setArgsToUpdateBuilder ...
func (r *RepositoryImpl) setArgsToUpdateBuilder(ctx context.Context, updateBuilder sq.UpdateBuilder, params arguments.UserUpdate) sq.UpdateBuilder {
	logger.WithFields(logger.Fields{
		"params": params,
	}).Infof("Repository setArgsToUpdateBuilder of user")
	if params.Username != nil {
		updateBuilder = updateBuilder.Set("username", *params.Username)
	}
	if params.Password != nil {
		updateBuilder = updateBuilder.Set("password", *params.Password)
	}
	if params.Name != nil {
		updateBuilder = updateBuilder.Set("name", *params.Name)
	}
	if value, _ := params.DateOfBirth.Value(); value != nil {
		updateBuilder = updateBuilder.Set("date_of_birth", *params.DateOfBirth)
	}
	if params.Reference != nil {
		updateBuilder = updateBuilder.Set("reference", *params.Reference)
	}
	if params.AvatarUrl != nil {
		updateBuilder = updateBuilder.Set("avatar_url", *params.AvatarUrl)
	}
	if params.LicenseNumber != nil {
		updateBuilder = updateBuilder.Set("license_number", *params.LicenseNumber)
	}
	if params.PhoneNumber != nil {
		updateBuilder = updateBuilder.Set("phone_number", *params.PhoneNumber)
	}
	if params.Extension != nil {
		updateBuilder = updateBuilder.Set("extension", *params.Extension)
	}
	if params.TelProvider != nil {
		updateBuilder = updateBuilder.Set("tel_provider", *params.TelProvider)
	}
	if params.TelApi != nil {
		updateBuilder = updateBuilder.Set("tel_api", *params.TelApi)
	}
	if params.SupervisorID != nil {
		updateBuilder = updateBuilder.Set("supervisor_id", *params.SupervisorID)
	}
	if params.RoleID != nil {
		updateBuilder = updateBuilder.Set("role_id", *params.RoleID)
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
func (r *RepositoryImpl) Update(ctx context.Context, params arguments.UserUpdate) (models.User, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Repository Update of user")
	var (
		user          models.User
		updateBuilder = sq.Update("user").Where(sq.Eq{"id": *params.ID})
	)
	updateBuilderWithArgs := r.setArgsToUpdateBuilder(ctx, updateBuilder, params)

	sql, args, err := updateBuilderWithArgs.ToSql()
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository Update build sql string of user")
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Update updateBuilderWithArgs error of user")
		return user, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Update PrepareContext error of user")
		return user, err
	}
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Update ExecContext error of user")
		return user, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Update RowsAffected error of user")
		return user, err
	}
	if rowAffected <= 0 {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   fmt.Errorf("error when update record id %d", *params.ID),
		}).Errorf("Repository Update rowAffected <= 0 of user")
		return user, fmt.Errorf("error when update record id %d", *params.ID)
	}
	newUser, err := r.GetByID(ctx, arguments.UserGetByID{ID: *params.ID})
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Update GetByID error of user")
		return user, err
	}
	return newUser, nil
}

// Delete ...
func (r *RepositoryImpl) Delete(ctx context.Context, params arguments.UserDelete) (int64, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"params":  params,
	}).Infof("Repository Delete of user")
	var (
		id            int64
		deleteBuilder = sq.Delete("user").Where(sq.Eq{"id": params.ID})
	)
	sql, args, err := deleteBuilder.ToSql()
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository Delete build sql string of user")
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Delete deleteBuilder error of user")
		return id, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Delete PrepareContext error of user")
		return id, err
	}
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Delete ExecContext error of user")
		return id, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   err,
		}).Errorf("Repository Delete RowsAffected error of user")
		return id, err
	}
	if rowAffected <= 0 {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"Error":   fmt.Errorf("not found record by id %d", params.ID),
		}).Errorf("Repository Update rowAffected <= 0 of user")
		return id, fmt.Errorf("not found record by id %d", params.ID)
	}
	return params.ID, nil
}

//go:generate mockery -name=IDatabase -output=mocks -case=underscore
