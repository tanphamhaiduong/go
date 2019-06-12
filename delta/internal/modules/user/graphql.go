package user

// IHandler ...
type IHandler interface {
	ICoreHandler
}

// ResolverImpl ...
type ResolverImpl struct {
	user IHandler
}

// NewResolver ...
func NewResolver(user IHandler) *ResolverImpl {
	return &ResolverImpl{
		user: user,
	}
}
