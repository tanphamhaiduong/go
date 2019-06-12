package company

// IRepository ...
type IRepository interface {
	ICoreRepository
}

// HandlerImpl ...
type HandlerImpl struct {
	company IRepository
}

// NewHandler ...
func NewHandler(company IRepository) *HandlerImpl {
	return &HandlerImpl{
		company: company,
	}
}
