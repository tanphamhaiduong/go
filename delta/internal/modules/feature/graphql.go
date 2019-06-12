package feature

// IHandler ...
type IHandler interface {
	ICoreHandler
}

// ResolverImpl ...
type ResolverImpl struct {
	feature IHandler
}

// NewResolver ...
func NewResolver(feature IHandler) *ResolverImpl {
	return &ResolverImpl{
		feature: feature,
	}
}
