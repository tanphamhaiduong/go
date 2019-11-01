package company

import (
	"context"

	"github.com/tanphamhaiduong/go/delta/internal/database"
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
