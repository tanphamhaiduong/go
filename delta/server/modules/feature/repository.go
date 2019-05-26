package feature

import "database/sql"

type featureImpl struct {
	db *sql.DB
}

// newRepository ...
func newRepository(db *sql.DB) featureImpl {
	return featureImpl{
		db: db,
	}
}
