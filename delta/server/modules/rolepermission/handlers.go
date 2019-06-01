package rolepermission

import "database/sql"

// iRepository ...
type iRepository interface {
	iCoreRepository
}

// HandlerImpl ...
type HandlerImpl struct {
	rolepermission iRepository
}

// NewHandler ...
func NewHandler(db *sql.DB) HandlerImpl {
	rolepermission := newRepository(db)
	return HandlerImpl{
		rolepermission: rolepermission,
	}
}
