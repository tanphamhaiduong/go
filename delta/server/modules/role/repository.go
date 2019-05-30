package role

import (
	"context"
	"database/sql"
)

type iDatabase interface {
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

type repositoryImpl struct {
	db iDatabase
}

// newRepository ...
func newRepository(db *sql.DB) repositoryImpl {
	return repositoryImpl{
		db: db,
	}
}
