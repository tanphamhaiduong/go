package arguments

// UserLogin ...
type UserLogin struct {
	Username string `graphql:"username" validate:"required"`
	Password string `graphql:"password" validate:"required"`
}
