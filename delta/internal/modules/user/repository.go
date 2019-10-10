package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/afex/hystrix-go/hystrix"
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

// GetByUsername ...
func (r *RepositoryImpl) GetByUsername(ctx context.Context, username string) (models.User, error) {
	logger.WithFields(logger.Fields{
		"traceId":  ctx.Value(utils.TraceIDKey),
		"username": username,
	}).Infof("Repository GetByUsername of user")
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
		).From("user").Where(sq.And{sq.Eq{"username": username}, sq.Eq{"status": "active"}})
	)
	sql, args, err := selectBuilder.ToSql()
	logger.WithFields(logger.Fields{
		"traceId": ctx.Value(utils.TraceIDKey),
		"SQL":     sql,
		"Args":    args,
	}).Infof("Repository GetByUsername build sql string of user")
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": ctx.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Repository GetByUsername selectBuilder error of user")
		return user, err
	}
	output := make(chan models.User, 1)
	errors := hystrix.Go("userGetByUsername", func() error {
		stmt, err := r.db.PrepareContext(ctx, sql)
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository GetByUsername PrepareContext error of user")
			return err
		}
		row := stmt.QueryRowContext(ctx, args...)
		err = r.scanUser(row, &user)
		if err != nil {
			logger.WithFields(logger.Fields{
				"traceId": ctx.Value(utils.TraceIDKey),
				"Error":   err,
			}).Errorf("Repository GetByUsername Scan error of user")
			return err
		}
		output <- user
		return nil
	}, nil)

	select {
	case err := <-errors:
		return user, err
	case out := <-output:
		return out, nil
	}
}
