// @generated
package rolepermission

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/tanphamhaiduong/go/delta/server/arguments"
	"github.com/tanphamhaiduong/go/delta/server/models"
	"github.com/tanphamhaiduong/go/delta/server/utils"
)

// GetByID ...
func (r repositoryImpl) GetByID(ctx context.Context, params arguments.RolePermissionGetByIDArgs) (models.RolePermission, error) {
	var (
		rolepermission = models.RolePermission{}
		selectBuilder  = sq.Select("*").From("role_permission").Where(sq.Eq{"id": params.ID})
	)
	sql, args, err := selectBuilder.ToSql()
	if err != nil {
		return rolepermission, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return rolepermission, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	row.Scan(
		&rolepermission.ID,
		&rolepermission.RoleID,
		&rolepermission.PermissionID,
		&rolepermission.CreatedBy,
		&rolepermission.UpdatedBy,
	)
	return rolepermission, nil
}

// GetByID ...
func (r repositoryImpl) GetByIDs(ctx context.Context, params arguments.RolePermissionGetByIDsArgs) ([]models.RolePermission, error) {
	var (
		rolepermissions = []models.RolePermission{}
		selectBuilder   = sq.Select("*").From("role_permission").Where(sq.Eq{"id": params.IDs})
	)
	sql, args, err := selectBuilder.ToSql()
	if err != nil {
		return rolepermissions, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return rolepermissions, err
	}
	rows, err := stmt.QueryContext(ctx, args...)
	defer rows.Close()
	if err != nil {
		return rolepermissions, err
	}
	for rows.Next() {
		rolepermission := models.RolePermission{}
		err := rows.Scan(
			&rolepermission.ID,
			&rolepermission.RoleID,
			&rolepermission.PermissionID,
			&rolepermission.CreatedBy,
			&rolepermission.UpdatedBy,
		)
		if err != nil {
			return rolepermissions, err
		}
		rolepermissions = append(rolepermissions, rolepermission)
	}
	return rolepermissions, nil
}

// setArgsToListSelectBuilder ...
func setArgsToListSelectBuilder(selectBuilder sq.SelectBuilder, params arguments.RolePermissionListArgs) sq.SelectBuilder {
	if params.ID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"id": params.ID})
	}
	if params.RoleID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"role_id": params.RoleID})
	}
	if params.PermissionID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"permission_id": params.PermissionID})
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
func (r repositoryImpl) List(ctx context.Context, params arguments.RolePermissionListArgs) ([]models.RolePermission, error) {
	var (
		rolepermissions = []models.RolePermission{}
		selectBuilder   = sq.Select("*").From("role_permission")
	)
	selectBuilderWithArgs := setArgsToListSelectBuilder(selectBuilder, params)
	sql, args, err := selectBuilderWithArgs.ToSql()
	if err != nil {
		return rolepermissions, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return rolepermissions, err
	}
	rows, err := stmt.QueryContext(ctx, args...)
	defer rows.Close()
	if err != nil {
		return rolepermissions, err
	}
	for rows.Next() {
		rolepermission := models.RolePermission{}
		err := rows.Scan(
			&rolepermission.ID,
			&rolepermission.RoleID,
			&rolepermission.PermissionID,
			&rolepermission.CreatedBy,
			&rolepermission.UpdatedBy,
		)
		if err != nil {
			return rolepermissions, err
		}
		rolepermissions = append(rolepermissions, rolepermission)
	}
	return rolepermissions, nil
}

// setArgsToCountSelectBuilder ...
func setArgsToCountSelectBuilder(selectBuilder sq.SelectBuilder, params arguments.RolePermissionCountArgs) sq.SelectBuilder {
	if params.ID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"id": params.ID})
	}
	if params.RoleID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"role_id": params.RoleID})
	}
	if params.PermissionID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"permission_id": params.PermissionID})
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
func (r repositoryImpl) Count(ctx context.Context, params arguments.RolePermissionCountArgs) (int64, error) {
	var (
		count         int64
		selectBuilder = sq.Select("COUNT(id)").From("role_permission")
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
func (r repositoryImpl) Insert(ctx context.Context, params arguments.RolePermissionInsertArgs) (models.RolePermission, error) {
	var (
		rolepermission = models.RolePermission{}
		insertBuilder  = sq.Insert("role_permission").Columns(
			"role_id",
			"permission_id",
			"created_by",
			"updated_by",
		).Values(
			params.RoleID,
			params.PermissionID,
			params.CreatedBy,
			params.UpdatedBy,
		)
	)
	sql, args, err := insertBuilder.ToSql()
	if err != nil {
		return rolepermission, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return rolepermission, err
	}
	row, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return rolepermission, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		return rolepermission, err
	}
	rolepermission = models.RolePermission{
		ID:     id,
		RoleID: params.RoleID,

		PermissionID: params.PermissionID,

		CreatedBy: params.CreatedBy,

		UpdatedBy: params.UpdatedBy,
	}
	return rolepermission, nil
}

// setArgsToUpdateBuilder ...
func setArgsToUpdateBuilder(updateBuilder sq.UpdateBuilder, params arguments.RolePermissionUpdateArgs) sq.UpdateBuilder {
	if params.RoleID != nil {
		updateBuilder = updateBuilder.Set("role_id", *params.RoleID)
	}

	if params.PermissionID != nil {
		updateBuilder = updateBuilder.Set("permission_id", *params.PermissionID)
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
func (r repositoryImpl) Update(ctx context.Context, params arguments.RolePermissionUpdateArgs) (models.RolePermission, error) {
	var (
		rolepermission = models.RolePermission{}
		updateBuilder  = sq.Update("role_permission").Where(sq.Eq{"id": *params.ID})
	)
	updateBuilderWithArgs := setArgsToUpdateBuilder(updateBuilder, params)

	sql, args, err := updateBuilderWithArgs.ToSql()
	if err != nil {
		return rolepermission, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return rolepermission, err
	}
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return rolepermission, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return rolepermission, err
	}
	if rowAffected == 0 {
		return rolepermission, fmt.Errorf("error when update record id %d", *params.ID)
	}
	rolepermission = models.RolePermission{
		ID:           *params.ID,
		RoleID:       *params.RoleID,
		PermissionID: *params.PermissionID,
		CreatedBy:    *params.CreatedBy,
		UpdatedBy:    *params.UpdatedBy,
	}
	return rolepermission, nil
}

// Delete ...
func (r repositoryImpl) Delete(ctx context.Context, params arguments.RolePermissionDeleteArgs) (int64, error) {
	var (
		id            int64
		deleteBuilder = sq.Delete("role_permission").Where(sq.Eq{"id": params.ID})
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
