package rolefeature

import "database/sql"

// IHandler ...
type IHandler interface {
	ICoreHandler
}

// ResolverImpl ...
type ResolverImpl struct {
	rolefeature IHandler
}

// NewResolver ...
func NewResolver(db *sql.DB) ResolverImpl {
	rolefeature := NewHandler(db)
	return ResolverImpl{
		rolefeature: rolefeature,
	}
}
