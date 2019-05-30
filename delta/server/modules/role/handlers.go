package role

import "database/sql"

// iRepository ...
type iRepository interface {
	iCoreRepository
}

// HandlerImpl ...
type HandlerImpl struct {
	role iRepository
}

// NewHandler ...
func NewHandler(db *sql.DB) HandlerImpl {
	role := newRepository(db)
	return HandlerImpl{
		role: role,
	}
}
