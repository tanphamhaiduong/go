package user

import "database/sql"

type userImpl struct {
	db *sql.DB
}

// newRepository ...
func newRepository(db *sql.DB) userImpl {
	return userImpl{
		db: db,
	}
}
