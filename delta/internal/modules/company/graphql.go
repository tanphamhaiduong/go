package company

// IHandler ...
type IHandler interface {
	ICoreHandler
}

// ResolverImpl ...
type ResolverImpl struct {
	company IHandler
}

// NewResolver ...
func NewResolver(company IHandler) *ResolverImpl {
	return &ResolverImpl{
		company: company,
	}
}
