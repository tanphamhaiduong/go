package role

// IRepository ...
type IRepository interface {
	ICoreRepository
}

// HandlerImpl ...
type HandlerImpl struct {
	role IRepository
}

// NewHandler ...
func NewHandler(role IRepository) *HandlerImpl {
	return &HandlerImpl{
		role: role,
	}
}
