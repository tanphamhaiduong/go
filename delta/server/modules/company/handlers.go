package company

import "database/sql"

// iRepository ...
type iRepository interface {
	iCoreRepository
}

// HandlerImpl ...
type HandlerImpl struct {
	company iRepository
}

// NewHandler ...
func NewHandler(db *sql.DB) HandlerImpl {
	company := newRepository(db)
	return HandlerImpl{
		company: company,
	}
}
