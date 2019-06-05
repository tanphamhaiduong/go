package company

import (
	"github.com/tanphamhaiduong/go/delta/server/database"
)

// IRepository ...
type IRepository interface {
	ICoreRepository
}

// HandlerImpl ...
type HandlerImpl struct {
	company IRepository
}

// NewHandler ...
func NewHandler(db database.IDB) HandlerImpl {
	company := newRepository(db)
	return HandlerImpl{
		company: company,
	}
}
