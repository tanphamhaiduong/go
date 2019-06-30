package rolepermission

// IRepository ...
type IRepository interface {
	ICoreRepository
}

// HandlerImpl ...
type HandlerImpl struct {
	rolepermission IRepository
}

// NewHandler ...
func NewHandler(rolepermission IRepository) *HandlerImpl {
	return &HandlerImpl{
		rolepermission: rolepermission,
	}
}
