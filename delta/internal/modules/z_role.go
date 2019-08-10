// @generated
package modules

import (
	"context"

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
	Delete(ctx context.Context, params arguments.RoleDelete) (int64, error)
}

// ICoreRoleResolver ...
type ICoreRoleResolver interface {
	IResolver
}

//go:generate mockery -name=IRoleResolver -output=mocks -case=underscore
//go:generate mockery -name=IRoleHandler -output=mocks -case=underscore
