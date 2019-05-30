package permission

import "database/sql"

// iRepository ...
type iRepository interface {
	iCoreRepository
}

// HandlerImpl ...
type HandlerImpl struct {
	permission iRepository
}

// NewHandler ...
func NewHandler(db *sql.DB) HandlerImpl {
	permission := newRepository(db)
	return HandlerImpl{
		permission: permission,
	}
}
