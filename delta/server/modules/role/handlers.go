package role

import "database/sql"

// iRepository ...
type iRepository interface {
	iCoreRepository
}

// HandlerImpl ...
type HandlerImpl struct {
	db          *sql.DB
	iRepository iRepository
}

// NewHandler ...
func NewHandler(db *sql.DB) HandlerImpl {
	iRepository := newRepository(db)
	return HandlerImpl{
		db:          db,
		iRepository: iRepository,
	}
}
