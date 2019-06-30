package permission

// IRepository ...
type IRepository interface {
	ICoreRepository
}

// HandlerImpl ...
type HandlerImpl struct {
	permission IRepository
}

// NewHandler ...
func NewHandler(permission IRepository) *HandlerImpl {
	return &HandlerImpl{
		permission: permission,
	}
}
