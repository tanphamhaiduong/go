package company

import "database/sql"

// IHandler ...
type IHandler interface {
	ICoreHandler
}

// ResolverImpl ...
type ResolverImpl struct {
	company IHandler
}

// NewResolver ...
func NewResolver(db *sql.DB) ResolverImpl {
	company := NewHandler(db)
	return ResolverImpl{
		company: company,
	}
}
