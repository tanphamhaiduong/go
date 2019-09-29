// @generated
package modules

import (
	"context"

	"github.com/graphql-go/graphql"
	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
)

// ICorePermissionHandler ...
type ICorePermissionHandler interface {
	GetByID(ctx context.Context, params arguments.PermissionGetByID) (models.Permission, error)
	Count(ctx context.Context, params arguments.PermissionCount) (int64, error)
	List(ctx context.Context, params arguments.PermissionList) ([]models.Permission, error)
	Insert(ctx context.Context, params arguments.PermissionInsert) (models.Permission, error)
}

// ICorePermissionResolver ...
type ICorePermissionResolver interface {
	GetByID(p graphql.ResolveParams) (interface{}, error)
	Count(p graphql.ResolveParams) (interface{}, error)
	List(p graphql.ResolveParams) (interface{}, error)
	Insert(p graphql.ResolveParams) (interface{}, error)
}

//go:generate mockery -name=IPermissionResolver -output=mocks -case=underscore
//go:generate mockery -name=IPermissionHandler -output=mocks -case=underscore
