// @generated
package role

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
		Name:        "Role",
		Description: "This is type Feature",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "This is role's id",
			},
			"name": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is role's name",
			},
			"companyId": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "This is role's companyId",
			},
			"status": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is role's status",
			},
			"createdBy": &graphql.Field{
				Type:        graphql.String,
				Description: "This is role's createdBy",
			},
			"updatedBy": &graphql.Field{
				Type:        graphql.String,
				Description: "This is role's updatedBy",
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
			Description: "This is role's id",
		},
		"name": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is role's name",
		},
		"companyId": &graphql.ArgumentConfig{
			Type:        graphql.Int,
			Description: "This is role's companyId",
		},
		"status": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is role's status",
		},
		"createdBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is role's createdBy",
		},
		"updatedBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is role's updatedBy",
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
			Description: "This is role's name",
		},
		"companyId": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "This is role's companyId",
		},
		"status": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is role's status",
		},
		"createdBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is role's createdBy",
		},
		"updatedBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is role's updatedBy",
		},
	}

	// UpdateTypeArgs ...
	UpdateTypeArgs = graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"name": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is role's name",
		},
		"companyId": &graphql.ArgumentConfig{
			Type:        graphql.Int,
			Description: "This is role's companyId",
		},
		"status": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is role's status",
		},
		"createdBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is role's createdBy",
		},
		"updatedBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is role's updatedBy",
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
	GetByID(ctx context.Context, param arguments.RoleGetByID) (models.Role, error)
	GetByIDs(ctx context.Context, param arguments.RoleGetByIDs) ([]models.Role, error)
	Count(ctx context.Context, params arguments.RoleCount) (int64, error)
	List(ctx context.Context, params arguments.RoleList) ([]models.Role, error)
	Insert(ctx context.Context, params arguments.RoleInsert) (models.Role, error)
	Update(ctx context.Context, params arguments.RoleUpdate) (models.Role, error)
	Delete(ctx context.Context, param arguments.RoleDelete) (int64, error)
}

// ForwardParams ...
func (r *ResolverImpl) ForwardParams(params graphql.ResolveParams) (interface{}, error) {
	logger.WithFields(logger.Fields{
		"TraceID": params.Context.Value("TraceID"),
		"params":  params,
	}).Infof("Resolver ForwardParams of role")
	return params.Args, nil
}

// GetByID ...
func (r *ResolverImpl) GetByID(param graphql.ResolveParams) (interface{}, error) {
	logger.WithFields(logger.Fields{
		"TraceID": param.Context.Value("TraceID"),
		"param":   param,
	}).Infof("Resolver GetByID of role")
	// parse param
	args := arguments.RoleGetByID{}
	if err := utils.Parse(param.Args, &args); err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": param.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver Count utils.Parse role")
		return nil, err
	}
	response, err := r.role.GetByID(param.Context, args)
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
	}).Infof("Resolver Count of role")
	// parse params
	args := arguments.RoleCount{}
	err := utils.Parse(params.Source.(map[string]interface{}), &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": params.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver Count utils.Parse role")
		return nil, err
	}
	response, err := r.role.Count(params.Context, args)
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
	}).Infof("Resolver List of role")
	// parse params
	args := arguments.RoleList{}
	err := utils.Parse(params.Source.(map[string]interface{}), &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": params.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver List utils.Parse role")
		return nil, err
	}
	response, err := r.role.List(params.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": params.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver List r.role.List role")
		return nil, err
	}
	return response, nil
}

// Insert ...
func (r *ResolverImpl) Insert(params graphql.ResolveParams) (interface{}, error) {
	logger.WithFields(logger.Fields{
		"TraceID": params.Context.Value("TraceID"),
		"params":  params,
	}).Infof("Resolver Insert of role")
	// parse params
	args := arguments.RoleInsert{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": params.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver Insert utils.Parse role")
		return nil, err
	}
	response, err := r.role.Insert(params.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": params.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver Insert r.role.Insert role")
		return nil, err
	}
	return response, nil
}

// Update ...
func (r *ResolverImpl) Update(params graphql.ResolveParams) (interface{}, error) {
	logger.WithFields(logger.Fields{
		"TraceID": params.Context.Value("TraceID"),
		"params":  params,
	}).Infof("Resolver Update of role")
	// parse params
	args := arguments.RoleUpdate{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": params.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver Update utils.Parse role")
		return nil, err
	}
	response, err := r.role.Update(params.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": params.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver Update r.role.Update role")
		return nil, err
	}
	return response, nil
}

// Delete ...
func (r *ResolverImpl) Delete(param graphql.ResolveParams) (interface{}, error) {
	logger.WithFields(logger.Fields{
		"TraceID": param.Context.Value("TraceID"),
		"param":   param,
	}).Infof("Resolver Delete of role")
	// parse param
	args := arguments.RoleDelete{}
	err := utils.Parse(param.Args, &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": param.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver Delete utils.Parse role")
		return nil, err
	}
	response, err := r.role.Delete(param.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": param.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver Delete r.role.Delete role")
		return nil, err
	}
	return response, nil
}

//go:generate mockery -name=IHandler -output=mocks -case=underscore
