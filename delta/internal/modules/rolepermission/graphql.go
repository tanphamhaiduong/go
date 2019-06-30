package rolepermission

// IHandler ...
type IHandler interface {
	ICoreHandler
}

// ResolverImpl ...
type ResolverImpl struct {
	rolepermission IHandler
}

// NewResolver ...
func NewResolver(rolepermission IHandler) *ResolverImpl {
	return &ResolverImpl{
		rolepermission: rolepermission,
	}
}
