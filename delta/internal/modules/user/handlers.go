package user

// IRepository ...
type IRepository interface {
	ICoreRepository
}

// HandlerImpl ...
type HandlerImpl struct {
	user IRepository
}

// NewHandler ...
func NewHandler(user IRepository) *HandlerImpl {
	return &HandlerImpl{
		user: user,
	}
}
