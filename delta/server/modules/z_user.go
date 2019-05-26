// @generated

package modules

import (
	"context"

	"github.com/tanphamhaiduong/go/delta/server/arguments"
	"github.com/tanphamhaiduong/go/delta/server/models"
)

// ICoreUserHandler ...
type ICoreUserHandler interface {
	GetByID(ctx context.Context, params arguments.UserGetByIDArgs) (models.User, error)
	Count(ctx context.Context, params arguments.UserCountArgs) (int64, error)
	List(ctx context.Context, params arguments.UserListArgs) ([]models.User, error)
	Insert(ctx context.Context, params arguments.UserInsertArgs) (models.User, error)
	Update(ctx context.Context, params arguments.UserUpdateArgs) (models.User, error)
	Delete(ctx context.Context, params arguments.UserDeleteArgs) (int64, error)
}

// ICoreUserResolver ...
type ICoreUserResolver interface {
	IResolver
}

//go:generate mockery -name=IUserResolver -output=mocks -case=underscore
//go:generate mockery -name=IUserHandler -output=mocks -case=underscore
