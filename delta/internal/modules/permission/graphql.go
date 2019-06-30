package permission

// IHandler ...
type IHandler interface {
	ICoreHandler
}

// ResolverImpl ...
type ResolverImpl struct {
	permission IHandler
}

// NewResolver ...
func NewResolver(permission IHandler) *ResolverImpl {
	return &ResolverImpl{
		permission: permission,
	}
}
