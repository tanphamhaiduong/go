package arguments

// UserLoginArgs ...
type UserLoginArgs struct {
	Username string `graphql:"username" validate:"required"`
	Password string `graphql:"password" validate:"required"`
}

// GetByUsernameArgs ...
type GetByUsernameArgs struct {
	Username string `graphql:"username" validate:"required"`
}
