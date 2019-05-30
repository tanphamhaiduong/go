package user

import "database/sql"

// iRepository ...
type iRepository interface {
	iCoreRepository
}

// HandlerImpl ...
type HandlerImpl struct {
	user iRepository
}

// NewHandler ...
func NewHandler(db *sql.DB) HandlerImpl {
	user := newRepository(db)
	return HandlerImpl{
		user: user,
	}
}
