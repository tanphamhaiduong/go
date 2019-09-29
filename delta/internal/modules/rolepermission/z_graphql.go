// @generated
package rolepermission

import (
	"context"

	"github.com/graphql-go/graphql"
	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
)

var (
	// Type ...
	Type = graphql.NewObject(graphql.ObjectConfig{
		Name:        "RolePermission",
		Description: "This is type Feature",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "This is rolePermission's id",
			},
			"roleId": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "This is rolePermission's roleId",
			},
			"permissionId": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "This is rolePermission's permissionId",
			},
			"createdBy": &graphql.Field{
				Type:        graphql.String,
				Description: "This is rolePermission's createdBy",
			},
			"updatedBy": &graphql.Field{
				Type:        graphql.String,
				Description: "This is rolePermission's updatedBy",
			},
		},
	})
)

// ICoreHandler ...
type ICoreHandler interface {
	GetByID(ctx context.Context, param arguments.RolePermissionGetByID) (models.RolePermission, error)
	GetByIDs(ctx context.Context, param arguments.RolePermissionGetByIDs) ([]models.RolePermission, error)
	Count(ctx context.Context, params arguments.RolePermissionCount) (int64, error)
	List(ctx context.Context, params arguments.RolePermissionList) ([]models.RolePermission, error)
	Insert(ctx context.Context, params arguments.RolePermissionInsert) (models.RolePermission, error)
	Update(ctx context.Context, params arguments.RolePermissionUpdate) (models.RolePermission, error)
	Delete(ctx context.Context, param arguments.RolePermissionDelete) (int64, error)
}

func (r *ResolverImpl) checkPermission(claims models.Claims, method string) bool {
	var isPermit bool
	for _, permission := range claims.Permissions {
		isPermit = permission.Name == method
		if isPermit {
			break
		}
	}
	return isPermit
}

//go:generate mockery -name=IHandler -output=mocks -case=underscore
