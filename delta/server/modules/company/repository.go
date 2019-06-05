package company

import (
	"context"

	"github.com/tanphamhaiduong/go/delta/server/database"
)

// IDatabase ...
type IDatabase interface {
	PrepareContext(ctx context.Context, query string) (database.IStmt, error)
}

type repositoryImpl struct {
	db IDatabase
}

// newRepository ...
func newRepository(db database.IDB) repositoryImpl {
	return repositoryImpl{
		db: db,
	}
}
