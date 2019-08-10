package rolepermission

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/tanphamhaiduong/go/common/logger"
	"github.com/tanphamhaiduong/go/delta/internal/database"
	"github.com/tanphamhaiduong/go/delta/internal/models"
	"github.com/tanphamhaiduong/go/delta/internal/utils"
)

// IDatabase ...
type IDatabase interface {
	Ping() error
	Close() error
	PrepareContext(ctx context.Context, query string) (database.IStmt, error)
}

// RepositoryImpl ...
type RepositoryImpl struct {
	db IDatabase
}

// NewRepository ...
func NewRepository(db database.IDB) *RepositoryImpl {
	return &RepositoryImpl{
		db: db,
	}
}

// GetByRoleID ...
func (r *RepositoryImpl) GetByRoleID(ctx context.Context, roleID int64) ([]models.RolePermission, error) {
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"roleID":  roleID,
	}).Infof("Repository GetByRoleID of rolepermission")
	var (
		rolepermissions []models.RolePermission
		selectBuilder   = sq.Select(
			"id",
			"role_id",
			"permission_id",
			"created_by",
			"updated_by",
		).From("role_permission").Where(sq.Eq{"role_id": roleID})
	)
	sql, args, err := selectBuilder.ToSql()
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository GetByRoleID build sql string of rolepermission")
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository GetByRoleID selectBuilder error of rolepermission")
		return rolepermissions, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository GetByRoleID PrepareContext error of rolepermission")
		return rolepermissions, err
	}
	rows, err := stmt.QueryContext(ctx, args...)
	defer rows.Close()
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository GetByRoleID QueryContext error of rolepermission")
		return rolepermissions, err
	}
	for rows.Next() {
		rolepermission := models.RolePermission{}
		err := r.scanRolePermission(rows, &rolepermission)
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository GetByRoleID Scan error of rolepermission")
			return rolepermissions, err
		}
		rolepermissions = append(rolepermissions, rolepermission)
	}
	return rolepermissions, nil
}
