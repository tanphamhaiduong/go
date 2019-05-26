package rolefeature

import "database/sql"

type rolefeatureImpl struct {
	db *sql.DB
}

// newRepository ...
func newRepository(db *sql.DB) rolefeatureImpl {
	return rolefeatureImpl{
		db: db,
	}
}
