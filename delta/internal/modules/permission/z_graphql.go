// @generated
package permission

import (
	"context"

	"github.com/graphql-go/graphql"
	"github.com/tanphamhaiduong/go/common/logger"
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
			"description": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is permission's description",
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
		"description": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is permission's description",
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
		"description": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is permission's description",
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
		"description": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is permission's description",
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
	GetByID(ctx context.Context, param arguments.PermissionGetByID) (models.Permission, error)
	GetByIDs(ctx context.Context, param arguments.PermissionGetByIDs) ([]models.Permission, error)
	Count(ctx context.Context, params arguments.PermissionCount) (int64, error)
	List(ctx context.Context, params arguments.PermissionList) ([]models.Permission, error)
	Insert(ctx context.Context, params arguments.PermissionInsert) (models.Permission, error)
	Update(ctx context.Context, params arguments.PermissionUpdate) (models.Permission, error)
	Delete(ctx context.Context, param arguments.PermissionDelete) (int64, error)
}

// ForwardParams ...
func (r *ResolverImpl) ForwardParams(params graphql.ResolveParams) (interface{}, error) {
	logger.WithFields(logger.Fields{
		"TraceID": params.Context.Value("TraceID"),
		"params":  params,
	}).Infof("Resolver ForwardParams of permission")
	return params.Args, nil
}

// GetByID ...
func (r *ResolverImpl) GetByID(param graphql.ResolveParams) (interface{}, error) {
	logger.WithFields(logger.Fields{
		"TraceID": param.Context.Value("TraceID"),
		"param":   param,
	}).Infof("Resolver GetByID of permission")
	// parse param
	args := arguments.PermissionGetByID{}
	if err := utils.Parse(param.Args, &args); err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": param.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver Count utils.Parse permission")
		return nil, err
	}
	response, err := r.permission.GetByID(param.Context, args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Count ...
func (r *ResolverImpl) Count(params graphql.ResolveParams) (interface{}, error) {
	logger.WithFields(logger.Fields{
		"TraceID": params.Context.Value("TraceID"),
		"params":  params,
	}).Infof("Resolver Count of permission")
	// parse params
	args := arguments.PermissionCount{}
	err := utils.Parse(params.Source.(map[string]interface{}), &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": params.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver Count utils.Parse permission")
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
	logger.WithFields(logger.Fields{
		"TraceID": params.Context.Value("TraceID"),
		"params":  params,
	}).Infof("Resolver List of permission")
	// parse params
	args := arguments.PermissionList{}
	err := utils.Parse(params.Source.(map[string]interface{}), &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": params.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver List utils.Parse permission")
		return nil, err
	}
	response, err := r.permission.List(params.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": params.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver List r.permission.List permission")
		return nil, err
	}
	return response, nil
}

// Insert ...
func (r *ResolverImpl) Insert(params graphql.ResolveParams) (interface{}, error) {
	logger.WithFields(logger.Fields{
		"TraceID": params.Context.Value("TraceID"),
		"params":  params,
	}).Infof("Resolver Insert of permission")
	// parse params
	args := arguments.PermissionInsert{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": params.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver Insert utils.Parse permission")
		return nil, err
	}
	response, err := r.permission.Insert(params.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": params.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver Insert r.permission.Insert permission")
		return nil, err
	}
	return response, nil
}

// Update ...
func (r *ResolverImpl) Update(params graphql.ResolveParams) (interface{}, error) {
	logger.WithFields(logger.Fields{
		"TraceID": params.Context.Value("TraceID"),
		"params":  params,
	}).Infof("Resolver Update of permission")
	// parse params
	args := arguments.PermissionUpdate{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": params.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver Update utils.Parse permission")
		return nil, err
	}
	response, err := r.permission.Update(params.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": params.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver Update r.permission.Update permission")
		return nil, err
	}
	return response, nil
}

// Delete ...
func (r *ResolverImpl) Delete(param graphql.ResolveParams) (interface{}, error) {
	logger.WithFields(logger.Fields{
		"TraceID": param.Context.Value("TraceID"),
		"param":   param,
	}).Infof("Resolver Delete of permission")
	// parse param
	args := arguments.PermissionDelete{}
	err := utils.Parse(param.Args, &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": param.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver Delete utils.Parse permission")
		return nil, err
	}
	response, err := r.permission.Delete(param.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": param.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver Delete r.permission.Delete permission")
		return nil, err
	}
	return response, nil
}

//go:generate mockery -name=IHandler -output=mocks -case=underscore
