// @generated
package modules

import "github.com/graphql-go/graphql"

// ICoreRolePermissionHandler ...
type ICoreRolePermissionHandler interface {
}

// ICoreRolePermissionResolver ...
type ICoreRolePermissionResolver interface {
	ForwardParams(p graphql.ResolveParams) (interface{}, error)
}

//go:generate mockery -name=IRolePermissionResolver -output=mocks -case=underscore
//go:generate mockery -name=IRolePermissionHandler -output=mocks -case=underscore
