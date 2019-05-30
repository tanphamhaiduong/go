package user

import "database/sql"

// IHandler ...
type IHandler interface {
	ICoreHandler
}

// ResolverImpl ...
type ResolverImpl struct {
	user IHandler
}

// NewResolver ...
func NewResolver(db *sql.DB) ResolverImpl {
	user := NewHandler(db)
	return ResolverImpl{
		user: user,
	}
}
