// @generated
package rolepermission

import (
	"context"

	"github.com/graphql-go/graphql"
	log "github.com/sirupsen/logrus"
	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
	"github.com/tanphamhaiduong/go/delta/internal/utils"
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
			"status": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is role's status",
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

	// GetByIDTypeArgs ...
	GetByIDTypeArgs = graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	}

	// ListTypeArgs ...
	ListTypeArgs = graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type:        graphql.Int,
			Description: "This is rolePermission's id",
		},
		"roleId": &graphql.ArgumentConfig{
			Type:        graphql.Int,
			Description: "This is rolePermission's roleId",
		},
		"permissionId": &graphql.ArgumentConfig{
			Type:        graphql.Int,
			Description: "This is rolePermission's permissionId",
		},
		"status": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is role's status",
		},
		"createdBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is rolePermission's createdBy",
		},
		"updatedBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is rolePermission's updatedBy",
		},
		"page": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "This is feature page",
		},
		"pageSize": &graphql.ArgumentConfig{
			Type:         graphql.Int,
			Description:  "This is feature pageSize",
			DefaultValue: 10,
		},
	}

	// InsertTypeArgs ...
	InsertTypeArgs = graphql.FieldConfigArgument{
		"roleId": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "This is rolePermission's roleId",
		},
		"permissionId": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "This is rolePermission's permissionId",
		},
		"status": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is role's status",
		},
		"createdBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is rolePermission's createdBy",
		},
		"updatedBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is rolePermission's updatedBy",
		},
	}

	// UpdateTypeArgs ...
	UpdateTypeArgs = graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"roleId": &graphql.ArgumentConfig{
			Type:        graphql.Int,
			Description: "This is rolePermission's roleId",
		},
		"permissionId": &graphql.ArgumentConfig{
			Type:        graphql.Int,
			Description: "This is rolePermission's permissionId",
		},
		"status": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is role's status",
		},
		"createdBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is rolePermission's createdBy",
		},
		"updatedBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is rolePermission's updatedBy",
		},
	}

	// DeleteTypeArgs ...
	DeleteTypeArgs = graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	}
)

// ICoreHandler ...
type ICoreHandler interface {
	GetByID(ctx context.Context, params arguments.RolePermissionGetByIDArgs) (models.RolePermission, error)
	GetByIDs(ctx context.Context, params arguments.RolePermissionGetByIDsArgs) ([]models.RolePermission, error)
	Count(ctx context.Context, params arguments.RolePermissionCountArgs) (int64, error)
	List(ctx context.Context, params arguments.RolePermissionListArgs) ([]models.RolePermission, error)
	Insert(ctx context.Context, params arguments.RolePermissionInsertArgs) (models.RolePermission, error)
	Update(ctx context.Context, params arguments.RolePermissionUpdateArgs) (models.RolePermission, error)
	Delete(ctx context.Context, params arguments.RolePermissionDeleteArgs) (int64, error)
}

// ForwardParams ...
func (r *ResolverImpl) ForwardParams(params graphql.ResolveParams) (interface{}, error) {
	log.WithField("params", params).Info("Resolver ForwardParams of rolepermission")
	return params.Args, nil
}

// GetByID ...
func (r *ResolverImpl) GetByID(params graphql.ResolveParams) (interface{}, error) {
	log.WithField("params", params).Info("Resolver GetByID of rolepermission")
	// parse params
	args := arguments.RolePermissionGetByIDArgs{}
	if err := utils.Parse(params.Args, &args); err != nil {
		return nil, err
	}
	response, err := r.rolepermission.GetByID(params.Context, args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Count ...
func (r *ResolverImpl) Count(params graphql.ResolveParams) (interface{}, error) {
	log.WithField("params", params).Info("Resolver Count of rolepermission")
	// parse params
	args := arguments.RolePermissionCountArgs{}
	err := utils.Parse(params.Source.(map[string]interface{}), &args)
	if err != nil {
		return nil, err
	}
	response, err := r.rolepermission.Count(params.Context, args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// List ...
func (r *ResolverImpl) List(params graphql.ResolveParams) (interface{}, error) {
	log.WithField("params", params).Info("Resolver List of rolepermission")
	// parse params
	args := arguments.RolePermissionListArgs{}
	err := utils.Parse(params.Source.(map[string]interface{}), &args)
	if err != nil {
		log.WithField("Error", err).Error("Resolver List utils.Parse rolepermission")
		return nil, err
	}
	response, err := r.rolepermission.List(params.Context, args)
	if err != nil {
		log.WithField("Error", err).Error("Resolver List r.rolepermission.List rolepermission")
		return nil, err
	}
	return response, nil
}

// Insert ...
func (r *ResolverImpl) Insert(params graphql.ResolveParams) (interface{}, error) {
	log.WithField("params", params).Info("Resolver Insert of rolepermission")
	// parse params
	args := arguments.RolePermissionInsertArgs{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		log.WithField("Error", err).Error("Resolver Insert utils.Parse rolepermission")
		return nil, err
	}
	response, err := r.rolepermission.Insert(params.Context, args)
	if err != nil {
		log.WithField("Error", err).Error("Resolver Insert r.rolepermission.Insert rolepermission")
		return nil, err
	}
	return response, nil
}

// Update ...
func (r *ResolverImpl) Update(params graphql.ResolveParams) (interface{}, error) {
	log.WithField("params", params).Info("Resolver Update of rolepermission")
	// parse params
	args := arguments.RolePermissionUpdateArgs{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		log.WithField("Error", err).Error("Resolver Update utils.Parse rolepermission")
		return nil, err
	}
	response, err := r.rolepermission.Update(params.Context, args)
	if err != nil {
		log.WithField("Error", err).Error("Resolver Update r.rolepermission.Update rolepermission")
		return nil, err
	}
	return response, nil
}

// Delete ...
func (r *ResolverImpl) Delete(params graphql.ResolveParams) (interface{}, error) {
	log.WithField("params", params).Info("Resolver Delete of rolepermission")
	// parse params
	args := arguments.RolePermissionDeleteArgs{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		log.WithField("Error", err).Error("Resolver Delete utils.Parse rolepermission")
		return nil, err
	}
	response, err := r.rolepermission.Delete(params.Context, args)
	if err != nil {
		log.WithField("Error", err).Error("Resolver Delete r.rolepermission.Delete rolepermission")
		return nil, err
	}
	return response, nil
}

//go:generate mockery -name=IHandler -output=mocks -case=underscore
