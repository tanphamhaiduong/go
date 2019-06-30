package role

// IHandler ...
type IHandler interface {
	ICoreHandler
}

// ResolverImpl ...
type ResolverImpl struct {
	role IHandler
}

// NewResolver ...
func NewResolver(role IHandler) *ResolverImpl {
	return &ResolverImpl{
		role: role,
	}
}
