package modules

import (
	"github.com/graphql-go/graphql"
)

// IUserResolver ...
type IUserResolver interface {
	ICoreUserResolver
	Login(param graphql.ResolveParams) (interface{}, error)
}

// IUserHandler ...
type IUserHandler interface {
	ICoreUserHandler
}
