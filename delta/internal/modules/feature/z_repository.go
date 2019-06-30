// @generated
package feature

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
func (r *RepositoryImpl) GetByID(ctx context.Context, params arguments.FeatureGetByIDArgs) (models.Feature, error) {
	log.WithField("params", params).Info("Repository GetByID of feature")
	var (
		feature       models.Feature
		selectBuilder = sq.Select(
			"id",
			"url",
			"meta",
			"company_id",
			"status",
			"created_by",
			"updated_by",
		).From("feature").Where(sq.Eq{"id": params.ID})
	)
	sql, args, err := selectBuilder.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository GetByID build sql string of feature")
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByID selectBuilder error of feature")
		return feature, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByID PrepareContext error of feature")
		return feature, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	err = row.Scan(
		&feature.ID,
		&feature.URL,
		&feature.Meta,
		&feature.CompanyID,
		&feature.Status,
		&feature.CreatedBy,
		&feature.UpdatedBy,
	)
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByID Scan error of feature")
		return feature, err
	}
	return feature, nil
}

// GetByIDs ...
func (r *RepositoryImpl) GetByIDs(ctx context.Context, params arguments.FeatureGetByIDsArgs) ([]models.Feature, error) {
	log.WithField("params", params).Info("Repository GetByIDs of feature")
	var (
		features      []models.Feature
		selectBuilder = sq.Select(
			"id",
			"url",
			"meta",
			"company_id",
			"status",
			"created_by",
			"updated_by",
		).From("feature").Where(sq.Eq{"id": params.IDs})
	)
	sql, args, err := selectBuilder.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository GetByIDs build sql string of feature")
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByIDs selectBuilder error of feature")
		return features, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByIDs PrepareContext error of feature")
		return features, err
	}
	rows, err := stmt.QueryContext(ctx, args...)
	defer rows.Close()
	if err != nil {
		log.WithField("Error", err).Error("Repository GetByIDs QueryContext error of feature")
		return features, err
	}
	for rows.Next() {
		feature := models.Feature{}
		err := rows.Scan(
			&feature.ID,
			&feature.URL,
			&feature.Meta,
			&feature.CompanyID,
			&feature.Status,
			&feature.CreatedBy,
			&feature.UpdatedBy,
		)
		if err != nil {
			log.WithField("Error", err).Error("Repository GetByIDs Scan error of feature")
			return features, err
		}
		features = append(features, feature)
	}
	return features, nil
}

// setArgsToListSelectBuilder ...
func (r *RepositoryImpl) setArgsToListSelectBuilder(selectBuilder sq.SelectBuilder, params arguments.FeatureListArgs) sq.SelectBuilder {
	log.WithField("params", params).Info("Repository setArgsToListSelectBuilder of feature")
	if params.ID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"id": params.ID})
	}
	if params.URL != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"url": params.URL})
	}
	if params.Meta != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"meta": params.Meta})
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
func (r *RepositoryImpl) List(ctx context.Context, params arguments.FeatureListArgs) ([]models.Feature, error) {
	log.WithField("params", params).Info("Repository List of feature")
	var (
		features      []models.Feature
		selectBuilder = sq.Select(
			"id",
			"url",
			"meta",
			"company_id",
			"status",
			"created_by",
			"updated_by",
		).From("feature")
	)
	selectBuilderWithArgs := r.setArgsToListSelectBuilder(selectBuilder, params)
	sql, args, err := selectBuilderWithArgs.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository List build sql string of feature")
	if err != nil {
		log.WithField("Error", err).Error("Repository List selectBuilderWithArgs error of feature")
		return features, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository List PrepareContext error of feature")
		return features, err
	}
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		log.WithField("Error", err).Error("Repository List QueryContext error of feature")
		return features, err
	}
	defer rows.Close()
	for rows.Next() {
		feature := models.Feature{}
		err := rows.Scan(
			&feature.ID,
			&feature.URL,
			&feature.Meta,
			&feature.CompanyID,
			&feature.Status,
			&feature.CreatedBy,
			&feature.UpdatedBy,
		)
		if err != nil {
			log.WithField("Error", err).Error("Repository List Scan error of feature")
			return features, err
		}
		features = append(features, feature)
	}
	return features, nil
}

// setArgsToCountSelectBuilder ...
func (r *RepositoryImpl) setArgsToCountSelectBuilder(selectBuilder sq.SelectBuilder, params arguments.FeatureCountArgs) sq.SelectBuilder {
	log.WithField("params", params).Info("Repository setArgsToCountSelectBuilder of feature")
	if params.ID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"id": params.ID})
	}
	if params.URL != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"url": params.URL})
	}
	if params.Meta != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"meta": params.Meta})
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
func (r *RepositoryImpl) Count(ctx context.Context, params arguments.FeatureCountArgs) (int64, error) {
	log.WithField("params", params).Info("Repository Count of feature")
	var (
		count         int64
		selectBuilder = sq.Select("COUNT(id)").From("feature")
	)
	selectBuilderWithArgs := r.setArgsToCountSelectBuilder(selectBuilder, params)
	sql, args, err := selectBuilderWithArgs.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository Count build sql string of feature")
	if err != nil {
		log.WithField("Error", err).Error("Repository Count selectBuilderWithArgs error of feature")
		return count, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository Count PrepareContext error of feature")
		return count, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	err = row.Scan(&count)
	if err != nil {
		log.WithField("Error", err).Error("Repository Count Scan error of feature")
		return count, err
	}
	return count, nil
}

// Insert ...
func (r *RepositoryImpl) Insert(ctx context.Context, params arguments.FeatureInsertArgs) (models.Feature, error) {
	log.WithField("params", params).Info("Repository Insert of feature")
	var (
		feature       models.Feature
		insertBuilder = sq.Insert("feature").Columns(
			"url",
			"meta",
			"company_id",
			"status",
			"created_by",
			"updated_by",
		).Values(
			params.URL,
			params.Meta,
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
	}).Info("Repository Insert build sql string of feature")
	if err != nil {
		log.WithField("Error", err).Error("Repository Insert insertBuilder error of feature")
		return feature, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository Insert PrepareContext error of feature")
		return feature, err
	}
	row, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		log.WithField("Error", err).Error("Repository Insert ExecContext error of feature")
		return feature, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		log.WithField("Error", err).Error("Repository Insert LastInsertId error of feature")
		return feature, err
	}
	newFeature, err := r.GetByID(ctx, arguments.FeatureGetByIDArgs{ID: id})
	if err != nil {
		log.WithField("Error", err).Error("Repository Insert GetByID error of feature")
		return feature, err
	}
	return newFeature, nil
}

// setArgsToUpdateBuilder ...
func (r *RepositoryImpl) setArgsToUpdateBuilder(updateBuilder sq.UpdateBuilder, params arguments.FeatureUpdateArgs) sq.UpdateBuilder {
	log.WithField("params", params).Info("Repository setArgsToUpdateBuilder of feature")
	if params.URL != nil {
		updateBuilder = updateBuilder.Set("url", *params.URL)
	}

	if params.Meta != nil {
		updateBuilder = updateBuilder.Set("meta", *params.Meta)
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
func (r *RepositoryImpl) Update(ctx context.Context, params arguments.FeatureUpdateArgs) (models.Feature, error) {
	log.WithField("params", params).Info("Repository Update of feature")
	var (
		feature       models.Feature
		updateBuilder = sq.Update("feature").Where(sq.Eq{"id": *params.ID})
	)
	updateBuilderWithArgs := r.setArgsToUpdateBuilder(updateBuilder, params)

	sql, args, err := updateBuilderWithArgs.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository Update build sql string of feature")
	if err != nil {
		log.WithField("Error", err).Error("Repository Update updateBuilderWithArgs error of feature")
		return feature, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository Update PrepareContext error of feature")
		return feature, err
	}
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		log.WithField("Error", err).Error("Repository Update ExecContext error of feature")
		return feature, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		log.WithField("Error", err).Error("Repository Update RowsAffected error of feature")
		return feature, err
	}
	if rowAffected <= 0 {
		log.WithField("Error", fmt.Errorf("error when update record id %d", *params.ID)).Error("Repository Update rowAffected <= 0 of feature")
		return feature, fmt.Errorf("error when update record id %d", *params.ID)
	}
	newFeature, err := r.GetByID(ctx, arguments.FeatureGetByIDArgs{ID: *params.ID})
	if err != nil {
		log.WithField("Error", err).Error("Repository Update GetByID error of feature")
		return feature, err
	}
	return newFeature, nil
}

// Delete ...
func (r *RepositoryImpl) Delete(ctx context.Context, params arguments.FeatureDeleteArgs) (int64, error) {
	log.WithField("params", params).Info("Repository Delete of feature")
	var (
		id            int64
		deleteBuilder = sq.Delete("feature").Where(sq.Eq{"id": params.ID})
	)
	sql, args, err := deleteBuilder.ToSql()
	log.WithFields(log.Fields{
		"SQL":  sql,
		"Args": args,
	}).Info("Repository Delete build sql string of feature")
	if err != nil {
		log.WithField("Error", err).Error("Repository Delete deleteBuilder error of feature")
		return id, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.WithField("Error", err).Error("Repository Delete PrepareContext error of feature")
		return id, err
	}
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		log.WithField("Error", err).Error("Repository Delete ExecContext error of feature")
		return id, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		log.WithField("Error", err).Error("Repository Delete RowsAffected error of feature")
		return id, err
	}
	if rowAffected <= 0 {
		log.WithField("Error", fmt.Errorf("not found record by id %d", params.ID)).Error("Repository Update rowAffected <= 0 of feature")
		return id, fmt.Errorf("not found record by id %d", params.ID)
	}
	return params.ID, nil
}

//go:generate mockery -name=IDatabase -output=mocks -case=underscore
