package company

import (
	"github.com/tanphamhaiduong/go/delta/server/database"
)

// IHandler ...
type IHandler interface {
	ICoreHandler
}

// ResolverImpl ...
type ResolverImpl struct {
	company IHandler
}

// NewResolver ...
func NewResolver(db database.IDB) ResolverImpl {
	company := NewHandler(db)
	return ResolverImpl{
		company: company,
	}
}
