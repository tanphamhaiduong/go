package company

import "database/sql"

type companyImpl struct {
	db *sql.DB
}

// newRepository ...
func newRepository(db *sql.DB) companyImpl {
	return companyImpl{
		db: db,
	}
}
