// @generated
package rolepermission

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
	GetByID(ctx context.Context, param arguments.RolePermissionGetByID) (models.RolePermission, error)
	GetByIDs(ctx context.Context, param arguments.RolePermissionGetByIDs) ([]models.RolePermission, error)
	Count(ctx context.Context, params arguments.RolePermissionCount) (int64, error)
	List(ctx context.Context, params arguments.RolePermissionList) ([]models.RolePermission, error)
	Insert(ctx context.Context, params arguments.RolePermissionInsert) (models.RolePermission, error)
	Update(ctx context.Context, params arguments.RolePermissionUpdate) (models.RolePermission, error)
	Delete(ctx context.Context, param arguments.RolePermissionDelete) (int64, error)
}

// ForwardParams ...
func (r *ResolverImpl) ForwardParams(params graphql.ResolveParams) (interface{}, error) {
	logger.WithFields(logger.Fields{
		"TraceID": params.Context.Value("TraceID"),
		"params":  params,
	}).Infof("Resolver ForwardParams of rolepermission")
	return params.Args, nil
}

// GetByID ...
func (r *ResolverImpl) GetByID(param graphql.ResolveParams) (interface{}, error) {
	logger.WithFields(logger.Fields{
		"TraceID": param.Context.Value("TraceID"),
		"param":   param,
	}).Infof("Resolver GetByID of rolepermission")
	// parse param
	args := arguments.RolePermissionGetByID{}
	if err := utils.Parse(param.Args, &args); err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": param.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver Count utils.Parse rolepermission")
		return nil, err
	}
	response, err := r.rolepermission.GetByID(param.Context, args)
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
	}).Infof("Resolver Count of rolepermission")
	// parse params
	args := arguments.RolePermissionCount{}
	err := utils.Parse(params.Source.(map[string]interface{}), &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": params.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver Count utils.Parse rolepermission")
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
	logger.WithFields(logger.Fields{
		"TraceID": params.Context.Value("TraceID"),
		"params":  params,
	}).Infof("Resolver List of rolepermission")
	// parse params
	args := arguments.RolePermissionList{}
	err := utils.Parse(params.Source.(map[string]interface{}), &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": params.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver List utils.Parse rolepermission")
		return nil, err
	}
	response, err := r.rolepermission.List(params.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": params.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver List r.rolepermission.List rolepermission")
		return nil, err
	}
	return response, nil
}

// Insert ...
func (r *ResolverImpl) Insert(params graphql.ResolveParams) (interface{}, error) {
	logger.WithFields(logger.Fields{
		"TraceID": params.Context.Value("TraceID"),
		"params":  params,
	}).Infof("Resolver Insert of rolepermission")
	// parse params
	args := arguments.RolePermissionInsert{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": params.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver Insert utils.Parse rolepermission")
		return nil, err
	}
	response, err := r.rolepermission.Insert(params.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": params.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver Insert r.rolepermission.Insert rolepermission")
		return nil, err
	}
	return response, nil
}

// Update ...
func (r *ResolverImpl) Update(params graphql.ResolveParams) (interface{}, error) {
	logger.WithFields(logger.Fields{
		"TraceID": params.Context.Value("TraceID"),
		"params":  params,
	}).Infof("Resolver Update of rolepermission")
	// parse params
	args := arguments.RolePermissionUpdate{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": params.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver Update utils.Parse rolepermission")
		return nil, err
	}
	response, err := r.rolepermission.Update(params.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": params.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver Update r.rolepermission.Update rolepermission")
		return nil, err
	}
	return response, nil
}

// Delete ...
func (r *ResolverImpl) Delete(param graphql.ResolveParams) (interface{}, error) {
	logger.WithFields(logger.Fields{
		"TraceID": param.Context.Value("TraceID"),
		"param":   param,
	}).Infof("Resolver Delete of rolepermission")
	// parse param
	args := arguments.RolePermissionDelete{}
	err := utils.Parse(param.Args, &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": param.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver Delete utils.Parse rolepermission")
		return nil, err
	}
	response, err := r.rolepermission.Delete(param.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": param.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver Delete r.rolepermission.Delete rolepermission")
		return nil, err
	}
	return response, nil
}

//go:generate mockery -name=IHandler -output=mocks -case=underscore
