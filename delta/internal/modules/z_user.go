// @generated
package modules

import (
	"context"

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
	Delete(ctx context.Context, params arguments.UserDelete) (int64, error)
}

// ICoreUserResolver ...
type ICoreUserResolver interface {
	IResolver
}

//go:generate mockery -name=IUserResolver -output=mocks -case=underscore
//go:generate mockery -name=IUserHandler -output=mocks -case=underscore
