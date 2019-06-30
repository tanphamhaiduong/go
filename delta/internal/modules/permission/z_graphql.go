// @generated
package permission

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
		Name:        "Permission",
		Description: "This is type Feature",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "This is permission's id",
			},
			"name": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is permission's name",
			},
			"status": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is permission's active",
			},
			"createdBy": &graphql.Field{
				Type:        graphql.String,
				Description: "This is permission's createdBy",
			},
			"updatedBy": &graphql.Field{
				Type:        graphql.String,
				Description: "This is permission's updatedBy",
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
			Description: "This is permission's id",
		},
		"name": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is permission's name",
		},
		"status": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is permission's active",
		},
		"createdBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is permission's createdBy",
		},
		"updatedBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is permission's updatedBy",
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
		"name": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is permission's name",
		},
		"status": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is permission's active",
		},
		"createdBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is permission's createdBy",
		},
		"updatedBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is permission's updatedBy",
		},
	}

	// UpdateTypeArgs ...
	UpdateTypeArgs = graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"name": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is permission's name",
		},
		"status": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is permission's active",
		},
		"createdBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is permission's createdBy",
		},
		"updatedBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is permission's updatedBy",
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
	GetByID(ctx context.Context, params arguments.PermissionGetByIDArgs) (models.Permission, error)
	GetByIDs(ctx context.Context, params arguments.PermissionGetByIDsArgs) ([]models.Permission, error)
	Count(ctx context.Context, params arguments.PermissionCountArgs) (int64, error)
	List(ctx context.Context, params arguments.PermissionListArgs) ([]models.Permission, error)
	Insert(ctx context.Context, params arguments.PermissionInsertArgs) (models.Permission, error)
	Update(ctx context.Context, params arguments.PermissionUpdateArgs) (models.Permission, error)
	Delete(ctx context.Context, params arguments.PermissionDeleteArgs) (int64, error)
}

// ForwardParams ...
func (r *ResolverImpl) ForwardParams(params graphql.ResolveParams) (interface{}, error) {
	log.WithField("params", params).Info("Resolver ForwardParams of permission")
	return params.Args, nil
}

// GetByID ...
func (r *ResolverImpl) GetByID(params graphql.ResolveParams) (interface{}, error) {
	log.WithField("params", params).Info("Resolver GetByID of permission")
	// parse params
	args := arguments.PermissionGetByIDArgs{}
	if err := utils.Parse(params.Args, &args); err != nil {
		return nil, err
	}
	response, err := r.permission.GetByID(params.Context, args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Count ...
func (r *ResolverImpl) Count(params graphql.ResolveParams) (interface{}, error) {
	log.WithField("params", params).Info("Resolver Count of permission")
	// parse params
	args := arguments.PermissionCountArgs{}
	err := utils.Parse(params.Source.(map[string]interface{}), &args)
	if err != nil {
		return nil, err
	}
	response, err := r.permission.Count(params.Context, args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// List ...
func (r *ResolverImpl) List(params graphql.ResolveParams) (interface{}, error) {
	log.WithField("params", params).Info("Resolver List of permission")
	// parse params
	args := arguments.PermissionListArgs{}
	err := utils.Parse(params.Source.(map[string]interface{}), &args)
	if err != nil {
		log.WithField("Error", err).Error("Resolver List utils.Parse permission")
		return nil, err
	}
	response, err := r.permission.List(params.Context, args)
	if err != nil {
		log.WithField("Error", err).Error("Resolver List r.permission.List permission")
		return nil, err
	}
	return response, nil
}

// Insert ...
func (r *ResolverImpl) Insert(params graphql.ResolveParams) (interface{}, error) {
	log.WithField("params", params).Info("Resolver Insert of permission")
	// parse params
	args := arguments.PermissionInsertArgs{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		log.WithField("Error", err).Error("Resolver Insert utils.Parse permission")
		return nil, err
	}
	response, err := r.permission.Insert(params.Context, args)
	if err != nil {
		log.WithField("Error", err).Error("Resolver Insert r.permission.Insert permission")
		return nil, err
	}
	return response, nil
}

// Update ...
func (r *ResolverImpl) Update(params graphql.ResolveParams) (interface{}, error) {
	log.WithField("params", params).Info("Resolver Update of permission")
	// parse params
	args := arguments.PermissionUpdateArgs{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		log.WithField("Error", err).Error("Resolver Update utils.Parse permission")
		return nil, err
	}
	response, err := r.permission.Update(params.Context, args)
	if err != nil {
		log.WithField("Error", err).Error("Resolver Update r.permission.Update permission")
		return nil, err
	}
	return response, nil
}

// Delete ...
func (r *ResolverImpl) Delete(params graphql.ResolveParams) (interface{}, error) {
	log.WithField("params", params).Info("Resolver Delete of permission")
	// parse params
	args := arguments.PermissionDeleteArgs{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		log.WithField("Error", err).Error("Resolver Delete utils.Parse permission")
		return nil, err
	}
	response, err := r.permission.Delete(params.Context, args)
	if err != nil {
		log.WithField("Error", err).Error("Resolver Delete r.permission.Delete permission")
		return nil, err
	}
	return response, nil
}

//go:generate mockery -name=IHandler -output=mocks -case=underscore
