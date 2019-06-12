package database

import (
	"context"
	"database/sql"
)

// IRow ...
type IRow interface {
	Scan(dest ...interface{}) error
}

type rowImpl struct {
	row *sql.Row
	IRow
}

func (s *rowImpl) Scan(dest ...interface{}) error {
	return s.row.Scan(dest)
}

// IRows ...
type IRows interface {
	Next() bool
	Close() error
	IRow
}

type rowsImpl struct {
	rows *sql.Rows
	IRows
}

func (s *rowsImpl) Next() bool {
	return s.rows.Next()
}

func (s *rowsImpl) Close() error {
	return s.rows.Close()
}

// IResult ...
type IResult interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type resultImpl struct {
	result sql.Result
	IResult
}

func (s *resultImpl) LastInsertId() (int64, error) {
	return s.result.LastInsertId()
}

func (s *resultImpl) RowsAffected() (int64, error) {
	return s.result.RowsAffected()
}

// IStmt ...
type IStmt interface {
	ExecContext(ctx context.Context, args ...interface{}) (IResult, error)
	QueryRowContext(ctx context.Context, args ...interface{}) IRow
	QueryContext(ctx context.Context, args ...interface{}) (IRows, error)
}

type stmtImpl struct {
	stmt *sql.Stmt
	IStmt
}

func (s *stmtImpl) ExecContext(ctx context.Context, args ...interface{}) (IResult, error) {
	return s.stmt.ExecContext(ctx, args...)
}

func (s *stmtImpl) QueryRowContext(ctx context.Context, args ...interface{}) IRow {
	return s.stmt.QueryRowContext(ctx, args...)
}

func (s *stmtImpl) QueryContext(ctx context.Context, args ...interface{}) (IRows, error) {
	return s.stmt.QueryContext(ctx, args...)
}

// IDB ...
type IDB interface {
	Ping() error
	Close() error
	PrepareContext(ctx context.Context, query string) (IStmt, error)
}

type sqlImpl struct {
	db *sql.DB
	IDB
}

func (s *sqlImpl) Close() error {
	return s.db.Close()
}

func (s *sqlImpl) PrepareContext(ctx context.Context, query string) (IStmt, error) {
	stmt, err := s.db.PrepareContext(ctx, query)
	return &stmtImpl{stmt: stmt}, err
}

// NewSQLlib ....
func NewSQLlib(driverName, connectionString string) (IDB, error) {
	sqlDb, err := sql.Open(driverName, connectionString)
	sqlDb.SetConnMaxLifetime(0)
	sqlDb.SetMaxIdleConns(50)
	sqlDb.SetMaxOpenConns(50)
	if err != nil {
		return nil, err
	}
	err = sqlDb.Ping()
	if err != nil {
		return nil, err
	}
	return &sqlImpl{db: sqlDb}, nil
}

//go:generate mockery -name=IRow -output=mocks -case=underscore
//go:generate mockery -name=IRows -output=mocks -case=underscore
//go:generate mockery -name=IResult -output=mocks -case=underscore
//go:generate mockery -name=IStmt -output=mocks -case=underscore
//go:generate mockery -name=IDB -output=mocks -case=underscore
