// @generated
package permission

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/tanphamhaiduong/go/common/goerrors"
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
		"lastId": &graphql.ArgumentConfig{
			Type:         graphql.Int,
			Description:  "This is feature page",
			DefaultValue: 0,
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

// GetByID ...
func (r *ResolverImpl) GetByID(param graphql.ResolveParams) (interface{}, error) {
	logger.WithFields(logger.Fields{
		"traceId": param.Context.Value(utils.TraceIDKey),
		"param":   param,
	}).Infof("Resolver GetByID of permission")
	claims := utils.ExtractClaimsFromContext(param.Context)
	isPermit := r.checkPermission(claims, fmt.Sprintf("Permission%s", utils.GetByID))
	if !isPermit {
		logger.WithFields(logger.Fields{
			"traceId": param.Context.Value(utils.TraceIDKey),
			"Error":   goerrors.ErrNotAuthorized(param.Context.Value(utils.TraceIDKey)),
		}).Errorf("Resolver GetByID !isPermit of permission")
		return nil, goerrors.ErrNotAuthorized(param.Context.Value(utils.TraceIDKey))
	}
	// parse param
	args := arguments.PermissionGetByID{}
	if err := utils.Parse(param.Args, &args); err != nil {
		logger.WithFields(logger.Fields{
			"traceId": param.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver GetByID utils.Parse of permission")
		return nil, goerrors.ErrInternalServerError(param.Context.Value(utils.TraceIDKey))
	}
	response, err := r.permission.GetByID(param.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": param.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver GetByID r.permission.GetByID of permission")
		if err == sql.ErrNoRows {
			return nil, goerrors.ErrNotFound(param.Context.Value(utils.TraceIDKey))
		}
		return nil, goerrors.ErrInternalServerError(param.Context.Value(utils.TraceIDKey))
	}
	return response, nil
}

// Count ...
func (r *ResolverImpl) Count(params graphql.ResolveParams) (interface{}, error) {
	logger.WithFields(logger.Fields{
		"traceId": params.Context.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Resolver Count of permission")
	claims := utils.ExtractClaimsFromContext(params.Context)
	isPermit := r.checkPermission(claims, fmt.Sprintf("Permission%s", utils.Count))
	if !isPermit {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   goerrors.ErrNotAuthorized(params.Context.Value(utils.TraceIDKey)),
		}).Errorf("Resolver Count !isPermit of permission")
		return nil, goerrors.ErrNotAuthorized(params.Context.Value(utils.TraceIDKey))
	}
	// parse params
	args := arguments.PermissionCount{}
	err := utils.Parse(params.Source.(map[string]interface{}), &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver Count utils.Parse of permission")
		return nil, goerrors.ErrInternalServerError(params.Context.Value(utils.TraceIDKey))
	}
	response, err := r.permission.Count(params.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver Count r.permission.Count of permission")
		if err == sql.ErrNoRows {
			return nil, goerrors.ErrNotFound(params.Context.Value(utils.TraceIDKey))
		}
		return nil, goerrors.ErrInternalServerError(params.Context.Value(utils.TraceIDKey))
	}
	return response, nil
}

// List ...
func (r *ResolverImpl) List(params graphql.ResolveParams) (interface{}, error) {
	logger.WithFields(logger.Fields{
		"traceId": params.Context.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Resolver List of permission")
	claims := utils.ExtractClaimsFromContext(params.Context)
	isPermit := r.checkPermission(claims, fmt.Sprintf("Permission%s", utils.List))
	if !isPermit {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   goerrors.ErrNotAuthorized(params.Context.Value(utils.TraceIDKey)),
		}).Errorf("Resolver List  of permission")
		return nil, goerrors.ErrNotAuthorized(params.Context.Value(utils.TraceIDKey))
	}
	// parse params
	args := arguments.PermissionList{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver List utils.Parse of permission")
		return nil, goerrors.ErrInternalServerError(params.Context.Value(utils.TraceIDKey))
	}
	response, err := r.permission.List(params.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver List r.permission.List of permission")
		if err == sql.ErrNoRows {
			return nil, goerrors.ErrNotFound(params.Context.Value(utils.TraceIDKey))
		}
		return nil, goerrors.ErrInternalServerError(params.Context.Value(utils.TraceIDKey))
	}
	return response, nil
}

// Insert ...
func (r *ResolverImpl) Insert(params graphql.ResolveParams) (interface{}, error) {
	logger.WithFields(logger.Fields{
		"traceId": params.Context.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Resolver Insert of permission")
	claims := utils.ExtractClaimsFromContext(params.Context)
	isPermit := r.checkPermission(claims, fmt.Sprintf("Permission%s", utils.Insert))
	if !isPermit {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   goerrors.ErrNotAuthorized(params.Context.Value(utils.TraceIDKey)),
		}).Errorf("Resolver Insert !isPermit of permission")
		return nil, goerrors.ErrNotAuthorized(params.Context.Value(utils.TraceIDKey))
	}
	// parse params
	args := arguments.PermissionInsert{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver Insert utils.Parse of permission")
		return nil, goerrors.ErrInternalServerError(params.Context.Value(utils.TraceIDKey))
	}
	response, err := r.permission.Insert(params.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver Insert r.permission.Insert of permission")
		if err == sql.ErrNoRows {
			return nil, goerrors.ErrNotFound(params.Context.Value(utils.TraceIDKey))
		}
		return nil, goerrors.ErrInternalServerError(params.Context.Value(utils.TraceIDKey))
	}
	return response, nil
}

//go:generate mockery -name=IHandler -output=mocks -case=underscore
