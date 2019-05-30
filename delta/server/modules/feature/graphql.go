package feature

import "database/sql"

// IHandler ...
type IHandler interface {
	ICoreHandler
}

// ResolverImpl ...
type ResolverImpl struct {
	feature IHandler
}

// NewResolver ...
func NewResolver(db *sql.DB) ResolverImpl {
	feature := NewHandler(db)
	return ResolverImpl{
		feature: feature,
	}
}
