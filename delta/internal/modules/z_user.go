// @generated
package modules

import (
	"context"

	"github.com/graphql-go/graphql"
	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
)

// ICoreUserHandler ...
type ICoreUserHandler interface {
	GetByID(ctx context.Context, params arguments.UserGetByID) (models.User, error)
	Count(ctx context.Context, params arguments.UserCount) (int64, error)
	List(ctx context.Context, params arguments.UserList) ([]models.User, error)
	Insert(ctx context.Context, params arguments.UserInsert) (models.User, error)
	Update(ctx context.Context, params arguments.UserUpdate) (models.User, error)
}

// ICoreUserResolver ...
type ICoreUserResolver interface {
	ForwardParams(p graphql.ResolveParams) (interface{}, error)
	GetByID(p graphql.ResolveParams) (interface{}, error)
	Count(p graphql.ResolveParams) (interface{}, error)
	List(p graphql.ResolveParams) (interface{}, error)
	Insert(p graphql.ResolveParams) (interface{}, error)
	Update(p graphql.ResolveParams) (interface{}, error)
}

//go:generate mockery -name=IUserResolver -output=mocks -case=underscore
//go:generate mockery -name=IUserHandler -output=mocks -case=underscore
