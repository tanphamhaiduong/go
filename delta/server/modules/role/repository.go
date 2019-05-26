package role

import "database/sql"

type roleImpl struct {
	db *sql.DB
}

// newRepository ...
func newRepository(db *sql.DB) roleImpl {
	return roleImpl{
		db: db,
	}
}
