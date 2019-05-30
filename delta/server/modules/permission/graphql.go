package permission

import "database/sql"

// IHandler ...
type IHandler interface {
	ICoreHandler
}

// ResolverImpl ...
type ResolverImpl struct {
	permission IHandler
}

// NewResolver ...
func NewResolver(db *sql.DB) ResolverImpl {
	permission := NewHandler(db)
	return ResolverImpl{
		permission: permission,
	}
}
