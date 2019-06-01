// @generated
package rolefeature

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/tanphamhaiduong/go/delta/server/arguments"
	"github.com/tanphamhaiduong/go/delta/server/models"
	"github.com/tanphamhaiduong/go/delta/server/utils"
)

// GetByID ...
func (r repositoryImpl) GetByID(ctx context.Context, params arguments.RoleFeatureGetByIDArgs) (models.RoleFeature, error) {
	var (
		rolefeature   = models.RoleFeature{}
		selectBuilder = sq.Select("*").From("role_feature").Where(sq.Eq{"id": params.ID})
	)
	sql, args, err := selectBuilder.ToSql()
	if err != nil {
		return rolefeature, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return rolefeature, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	row.Scan(
		&rolefeature.ID,
		&rolefeature.RoleID,
		&rolefeature.PermissionID,
		&rolefeature.CreatedBy,
		&rolefeature.UpdatedBy,
	)
	return rolefeature, nil
}

// GetByID ...
func (r repositoryImpl) GetByIDs(ctx context.Context, params arguments.RoleFeatureGetByIDsArgs) ([]models.RoleFeature, error) {
	var (
		rolefeatures  = []models.RoleFeature{}
		selectBuilder = sq.Select("*").From("role_feature").Where(sq.Eq{"id": params.IDs})
	)
	sql, args, err := selectBuilder.ToSql()
	if err != nil {
		return rolefeatures, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return rolefeatures, err
	}
	rows, err := stmt.QueryContext(ctx, args...)
	defer rows.Close()
	if err != nil {
		return rolefeatures, err
	}
	for rows.Next() {
		rolefeature := models.RoleFeature{}
		err := rows.Scan(
			&rolefeature.ID,
			&rolefeature.RoleID,
			&rolefeature.PermissionID,
			&rolefeature.CreatedBy,
			&rolefeature.UpdatedBy,
		)
		if err != nil {
			return rolefeatures, err
		}
		rolefeatures = append(rolefeatures, rolefeature)
	}
	return rolefeatures, nil
}

// setArgsToListSelectBuilder ...
func setArgsToListSelectBuilder(selectBuilder sq.SelectBuilder, params arguments.RoleFeatureListArgs) sq.SelectBuilder {
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
func (r repositoryImpl) List(ctx context.Context, params arguments.RoleFeatureListArgs) ([]models.RoleFeature, error) {
	var (
		rolefeatures  = []models.RoleFeature{}
		selectBuilder = sq.Select("*").From("role_feature")
	)
	selectBuilderWithArgs := setArgsToListSelectBuilder(selectBuilder, params)
	sql, args, err := selectBuilderWithArgs.ToSql()
	if err != nil {
		return rolefeatures, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return rolefeatures, err
	}
	rows, err := stmt.QueryContext(ctx, args...)
	defer rows.Close()
	if err != nil {
		return rolefeatures, err
	}
	for rows.Next() {
		rolefeature := models.RoleFeature{}
		err := rows.Scan(
			&rolefeature.ID,
			&rolefeature.RoleID,
			&rolefeature.PermissionID,
			&rolefeature.CreatedBy,
			&rolefeature.UpdatedBy,
		)
		if err != nil {
			return rolefeatures, err
		}
		rolefeatures = append(rolefeatures, rolefeature)
	}
	return rolefeatures, nil
}

// setArgsToCountSelectBuilder ...
func setArgsToCountSelectBuilder(selectBuilder sq.SelectBuilder, params arguments.RoleFeatureCountArgs) sq.SelectBuilder {
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
func (r repositoryImpl) Count(ctx context.Context, params arguments.RoleFeatureCountArgs) (int64, error) {
	var (
		count         int64
		selectBuilder = sq.Select("COUNT(id)").From("role_feature")
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
func (r repositoryImpl) Insert(ctx context.Context, params arguments.RoleFeatureInsertArgs) (models.RoleFeature, error) {
	var (
		rolefeature   = models.RoleFeature{}
		insertBuilder = sq.Insert("role_feature").Columns(
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
		return rolefeature, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return rolefeature, err
	}
	row, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return rolefeature, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		return rolefeature, err
	}
	rolefeature = models.RoleFeature{
		ID:     id,
		RoleID: params.RoleID,

		PermissionID: params.PermissionID,

		CreatedBy: params.CreatedBy,

		UpdatedBy: params.UpdatedBy,
	}
	return rolefeature, nil
}

// setArgsToUpdateBuilder ...
func setArgsToUpdateBuilder(updateBuilder sq.UpdateBuilder, params arguments.RoleFeatureUpdateArgs) sq.UpdateBuilder {
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
func (r repositoryImpl) Update(ctx context.Context, params arguments.RoleFeatureUpdateArgs) (models.RoleFeature, error) {
	var (
		rolefeature   = models.RoleFeature{}
		updateBuilder = sq.Update("role_feature").Where(sq.Eq{"id": *params.ID})
	)
	updateBuilderWithArgs := setArgsToUpdateBuilder(updateBuilder, params)

	sql, args, err := updateBuilderWithArgs.ToSql()
	if err != nil {
		return rolefeature, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return rolefeature, err
	}
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return rolefeature, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return rolefeature, err
	}
	if rowAffected == 0 {
		return rolefeature, fmt.Errorf("error when update record id %d", *params.ID)
	}
	rolefeature = models.RoleFeature{
		ID:           *params.ID,
		RoleID:       *params.RoleID,
		PermissionID: *params.PermissionID,
		CreatedBy:    *params.CreatedBy,
		UpdatedBy:    *params.UpdatedBy,
	}
	return rolefeature, nil
}

// Delete ...
func (r repositoryImpl) Delete(ctx context.Context, params arguments.RoleFeatureDeleteArgs) (int64, error) {
	var (
		id            int64
		deleteBuilder = sq.Delete("role_feature").Where(sq.Eq{"id": params.ID})
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
