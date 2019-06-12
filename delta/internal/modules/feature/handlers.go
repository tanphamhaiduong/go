package feature

// IRepository ...
type IRepository interface {
	ICoreRepository
}

// HandlerImpl ...
type HandlerImpl struct {
	feature IRepository
}

// NewHandler ...
func NewHandler(feature IRepository) *HandlerImpl {
	return &HandlerImpl{
		feature: feature,
	}
}
