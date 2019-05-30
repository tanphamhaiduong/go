package rolefeature

import "database/sql"

// iRepository ...
type iRepository interface {
	iCoreRepository
}

// HandlerImpl ...
type HandlerImpl struct {
	rolefeature iRepository
}

// NewHandler ...
func NewHandler(db *sql.DB) HandlerImpl {
	rolefeature := newRepository(db)
	return HandlerImpl{
		rolefeature: rolefeature,
	}
}
