// @generated

package modules

import (
	"context"

	"github.com/tanphamhaiduong/go/delta/server/arguments"
	"github.com/tanphamhaiduong/go/delta/server/models"
)

// ICoreRoleHandler ...
type ICoreRoleHandler interface {
	GetByID(ctx context.Context, params arguments.RoleGetByIDArgs) (models.Role, error)
	Count(ctx context.Context, params arguments.RoleCountArgs) (int64, error)
	List(ctx context.Context, params arguments.RoleListArgs) ([]models.Role, error)
	Insert(ctx context.Context, params arguments.RoleInsertArgs) (models.Role, error)
	Update(ctx context.Context, params arguments.RoleUpdateArgs) (models.Role, error)
	Delete(ctx context.Context, params arguments.RoleDeleteArgs) (int64, error)
}

// ICoreRoleResolver ...
type ICoreRoleResolver interface {
	IResolver
}

//go:generate mockery -name=IRoleResolver -output=mocks -case=underscore
//go:generate mockery -name=IRoleHandler -output=mocks -case=underscore
