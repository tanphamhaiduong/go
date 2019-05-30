package role

import "database/sql"

// IHandler ...
type IHandler interface {
	ICoreHandler
}

// ResolverImpl ...
type ResolverImpl struct {
	role IHandler
}

// NewResolver ...
func NewResolver(db *sql.DB) ResolverImpl {
	role := NewHandler(db)
	return ResolverImpl{
		role: role,
	}
}
