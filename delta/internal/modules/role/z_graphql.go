// @generated
package role

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
	}).Infof("Resolver GetByID of role")
	claims := utils.ExtractClaimsFromContext(param.Context)
	isPermit := r.checkPermission(claims, fmt.Sprintf("Role%s", utils.GetByID))
	if !isPermit {
		logger.WithFields(logger.Fields{
			"traceId": param.Context.Value(utils.TraceIDKey),
			"Error":   goerrors.ErrNotAuthorized(param.Context.Value(utils.TraceIDKey)),
		}).Errorf("Resolver GetByID !isPermit of role")
		return nil, goerrors.ErrNotAuthorized(param.Context.Value(utils.TraceIDKey))
	}
	// parse param
	args := arguments.RoleGetByID{}
	if err := utils.Parse(param.Args, &args); err != nil {
		logger.WithFields(logger.Fields{
			"traceId": param.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver GetByID utils.Parse of role")
		return nil, goerrors.ErrInternalServerError(param.Context.Value(utils.TraceIDKey))
	}
	response, err := r.role.GetByID(param.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": param.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver GetByID r.role.GetByID of role")
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
	}).Infof("Resolver Count of role")
	claims := utils.ExtractClaimsFromContext(params.Context)
	isPermit := r.checkPermission(claims, fmt.Sprintf("Role%s", utils.Count))
	if !isPermit {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   goerrors.ErrNotAuthorized(params.Context.Value(utils.TraceIDKey)),
		}).Errorf("Resolver Count !isPermit of role")
		return nil, goerrors.ErrNotAuthorized(params.Context.Value(utils.TraceIDKey))
	}
	// parse params
	args := arguments.RoleCount{}
	err := utils.Parse(params.Source.(map[string]interface{}), &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver Count utils.Parse of role")
		return nil, goerrors.ErrInternalServerError(params.Context.Value(utils.TraceIDKey))
	}
	response, err := r.role.Count(params.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver Count r.role.Count of role")
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
	}).Infof("Resolver List of role")
	claims := utils.ExtractClaimsFromContext(params.Context)
	isPermit := r.checkPermission(claims, fmt.Sprintf("Role%s", utils.List))
	if !isPermit {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   goerrors.ErrNotAuthorized(params.Context.Value(utils.TraceIDKey)),
		}).Errorf("Resolver List  of role")
		return nil, goerrors.ErrNotAuthorized(params.Context.Value(utils.TraceIDKey))
	}
	// parse params
	args := arguments.RoleList{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver List utils.Parse of role")
		return nil, goerrors.ErrInternalServerError(params.Context.Value(utils.TraceIDKey))
	}
	response, err := r.role.List(params.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver List r.role.List of role")
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
	}).Infof("Resolver Insert of role")
	claims := utils.ExtractClaimsFromContext(params.Context)
	isPermit := r.checkPermission(claims, fmt.Sprintf("Role%s", utils.Insert))
	if !isPermit {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   goerrors.ErrNotAuthorized(params.Context.Value(utils.TraceIDKey)),
		}).Errorf("Resolver Insert !isPermit of role")
		return nil, goerrors.ErrNotAuthorized(params.Context.Value(utils.TraceIDKey))
	}
	// parse params
	args := arguments.RoleInsert{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver Insert utils.Parse of role")
		return nil, goerrors.ErrInternalServerError(params.Context.Value(utils.TraceIDKey))
	}
	response, err := r.role.Insert(params.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver Insert r.role.Insert of role")
		if err == sql.ErrNoRows {
			return nil, goerrors.ErrNotFound(params.Context.Value(utils.TraceIDKey))
		}
		return nil, goerrors.ErrInternalServerError(params.Context.Value(utils.TraceIDKey))
	}
	return response, nil
}

// Update ...
func (r *ResolverImpl) Update(params graphql.ResolveParams) (interface{}, error) {
	logger.WithFields(logger.Fields{
		"traceId": params.Context.Value(utils.TraceIDKey),
		"params":  params,
	}).Infof("Resolver Update of role")
	claims := utils.ExtractClaimsFromContext(params.Context)
	isPermit := r.checkPermission(claims, fmt.Sprintf("Role%s", utils.Update))
	if !isPermit {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   goerrors.ErrNotAuthorized(params.Context.Value(utils.TraceIDKey)),
		}).Errorf("Resolver Update !isPermit of role")
		return nil, goerrors.ErrNotAuthorized(params.Context.Value(utils.TraceIDKey))
	}
	// parse params
	args := arguments.RoleUpdate{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver Update utils.Parse of role")
		return nil, goerrors.ErrInternalServerError(params.Context.Value(utils.TraceIDKey))
	}
	response, err := r.role.Update(params.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver Update r.role.Update of role")
		if err == sql.ErrNoRows {
			return nil, goerrors.ErrNotFound(params.Context.Value(utils.TraceIDKey))
		}
		return nil, goerrors.ErrInternalServerError(params.Context.Value(utils.TraceIDKey))
	}
	return response, nil
}

//go:generate mockery -name=IHandler -output=mocks -case=underscore
