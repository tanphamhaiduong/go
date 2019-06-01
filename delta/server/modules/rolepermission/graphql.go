package rolepermission

import "database/sql"

// IHandler ...
type IHandler interface {
	ICoreHandler
}

// ResolverImpl ...
type ResolverImpl struct {
	rolepermission IHandler
}

// NewResolver ...
func NewResolver(db *sql.DB) ResolverImpl {
	rolepermission := NewHandler(db)
	return ResolverImpl{
		rolepermission: rolepermission,
	}
}
