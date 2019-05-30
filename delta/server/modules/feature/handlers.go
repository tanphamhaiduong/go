package feature

import "database/sql"

// iRepository ...
type iRepository interface {
	iCoreRepository
}

// HandlerImpl ...
type HandlerImpl struct {
	feature iRepository
}

// NewHandler ...
func NewHandler(db *sql.DB) HandlerImpl {
	feature := newRepository(db)
	return HandlerImpl{
		feature: feature,
	}
}
