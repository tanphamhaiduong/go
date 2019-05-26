// @generated

package role

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/tanphamhaiduong/go/delta/server/arguments"
	"github.com/tanphamhaiduong/go/delta/server/models"
	"github.com/tanphamhaiduong/go/delta/server/utils"
)

// getByID ...
func (r roleImpl) getByID(ctx context.Context, params arguments.RoleGetByIDArgs) (models.Role, error) {
	var (
		role          = models.Role{}
		selectBuilder = sq.Select("*").From("role").Where(sq.Eq{"id": params.ID})
	)
	sql, args, err := selectBuilder.ToSql()
	if err != nil {
		return role, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return role, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	row.Scan(
		&role.ID,
		&role.Name,
		&role.CompanyID,
		&role.Status,
		&role.CreatedBy,
		&role.UpdatedBy,
	)
	return role, nil
}

// setArgsToListSelectBuilder ...
func setArgsToListSelectBuilder(selectBuilder sq.SelectBuilder, params arguments.RoleListArgs) sq.SelectBuilder {
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
	if params.PageSize != 0 {
		selectBuilder = selectBuilder.Limit(uint64(params.PageSize))
	}
	if params.Page != 0 {
		offset := utils.CalculateOffsetForPage(params.Page, params.PageSize)
		selectBuilder = selectBuilder.Offset(uint64(offset))
	}
	return selectBuilder
}

// list ...
func (r roleImpl) list(ctx context.Context, params arguments.RoleListArgs) ([]models.Role, error) {
	var (
		roles         = []models.Role{}
		selectBuilder = sq.Select("*").From("role")
	)
	selectBuilderWithArgs := setArgsToListSelectBuilder(selectBuilder, params)
	sql, args, err := selectBuilderWithArgs.ToSql()
	if err != nil {
		return roles, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return roles, err
	}
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return roles, err
	}
	for rows.Next() {
		role := models.Role{}
		err := rows.Scan(
			&role.ID,
			&role.Name,
			&role.CompanyID,
			&role.Status,
			&role.CreatedBy,
			&role.UpdatedBy,
		)
		if err != nil {
			return roles, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}

// setArgsToCountSelectBuilder ...
func setArgsToCountSelectBuilder(selectBuilder sq.SelectBuilder, params arguments.RoleCountArgs) sq.SelectBuilder {
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

// count ...
func (r roleImpl) count(ctx context.Context, params arguments.RoleCountArgs) (int64, error) {
	var (
		count         int64
		selectBuilder = sq.Select("COUNT(id)").From("role")
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

// insert ...
func (r roleImpl) insert(ctx context.Context, params arguments.RoleInsertArgs) (models.Role, error) {
	var (
		role          = models.Role{}
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
	if err != nil {
		return role, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return role, err
	}
	row, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return role, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		return role, err
	}
	role = models.Role{
		ID: id, Name: params.Name,
		CompanyID: params.CompanyID,
		Status:    params.Status,
		CreatedBy: params.CreatedBy,
		UpdatedBy: params.UpdatedBy,
	}
	return role, nil
}

// setArgsToUpdateBuilder ...
func setArgsToUpdateBuilder(updateBuilder sq.UpdateBuilder, params arguments.RoleUpdateArgs) sq.UpdateBuilder {
	if params.Name != "" {
		updateBuilder = updateBuilder.Set("name", params.Name)
	}
	if params.CompanyID != 0 {
		updateBuilder = updateBuilder.Set("company_id", params.CompanyID)
	}
	if params.Status != "" {
		updateBuilder = updateBuilder.Set("status", params.Status)
	}
	if params.CreatedBy != "" {
		updateBuilder = updateBuilder.Set("created_by", params.CreatedBy)
	}
	if params.UpdatedBy != "" {
		updateBuilder = updateBuilder.Set("updated_by", params.UpdatedBy)
	}

	return updateBuilder
}

// update ...
func (r roleImpl) update(ctx context.Context, params arguments.RoleUpdateArgs) (models.Role, error) {
	var (
		role          = models.Role{}
		updateBuilder = sq.Update("role").Where(sq.Eq{"id": params.ID})
	)
	updateBuilderWithArgs := setArgsToUpdateBuilder(updateBuilder, params)

	sql, args, err := updateBuilderWithArgs.ToSql()
	if err != nil {
		return role, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return role, err
	}
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return role, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return role, err
	}
	if rowAffected == 0 {
		return role, fmt.Errorf("not found record by id %d", params.ID)
	}
	role = models.Role{
		ID:        params.ID,
		Name:      params.Name,
		CompanyID: params.CompanyID,
		Status:    params.Status,
		CreatedBy: params.CreatedBy,
		UpdatedBy: params.UpdatedBy,
	}
	return role, nil
}

// delete ...
func (r roleImpl) delete(ctx context.Context, params arguments.RoleDeleteArgs) (int64, error) {
	var (
		id            int64
		deleteBuilder = sq.Delete("role").Where(sq.Eq{"id": params.ID})
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
