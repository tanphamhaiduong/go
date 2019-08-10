// @generated
package modules

import (
	"context"

	"github.com/graphql-go/graphql"
	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
)

// ICoreRoleHandler ...
type ICoreRoleHandler interface {
	GetByID(ctx context.Context, params arguments.RoleGetByID) (models.Role, error)
	Count(ctx context.Context, params arguments.RoleCount) (int64, error)
	List(ctx context.Context, params arguments.RoleList) ([]models.Role, error)
	Insert(ctx context.Context, params arguments.RoleInsert) (models.Role, error)
	Update(ctx context.Context, params arguments.RoleUpdate) (models.Role, error)
}

// ICoreRoleResolver ...
type ICoreRoleResolver interface {
	ForwardParams(p graphql.ResolveParams) (interface{}, error)
	GetByID(p graphql.ResolveParams) (interface{}, error)
	Count(p graphql.ResolveParams) (interface{}, error)
	List(p graphql.ResolveParams) (interface{}, error)
	Insert(p graphql.ResolveParams) (interface{}, error)
	Update(p graphql.ResolveParams) (interface{}, error)
}

//go:generate mockery -name=IRoleResolver -output=mocks -case=underscore
//go:generate mockery -name=IRoleHandler -output=mocks -case=underscore
