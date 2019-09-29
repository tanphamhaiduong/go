// @generated
package company

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
		Name:        "Company",
		Description: "This is type Feature",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "This is company's id",
			},
			"name": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is company's name",
			},
			"status": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is company's status",
			},
			"createdBy": &graphql.Field{
				Type:        graphql.String,
				Description: "This is company's createdBy",
			},
			"updatedBy": &graphql.Field{
				Type:        graphql.String,
				Description: "This is company's updatedBy",
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
			Description: "This is company's id",
		},
		"name": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is company's name",
		},
		"status": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is company's status",
		},
		"createdBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is company's createdBy",
		},
		"updatedBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is company's updatedBy",
		},
		"page": &graphql.ArgumentConfig{
			Type:         graphql.Int,
			Description:  "This is feature page",
			DefaultValue: 1,
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
			Description: "This is company's name",
		},
		"status": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is company's status",
		},
		"createdBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is company's createdBy",
		},
		"updatedBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is company's updatedBy",
		},
	}
	// UpdateTypeArgs ...
	UpdateTypeArgs = graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"name": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is company's name",
		},
		"status": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is company's status",
		},
		"createdBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is company's createdBy",
		},
		"updatedBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is company's updatedBy",
		},
	}
)

// ICoreHandler ...
type ICoreHandler interface {
	GetByID(ctx context.Context, param arguments.CompanyGetByID) (models.Company, error)
	GetByIDs(ctx context.Context, param arguments.CompanyGetByIDs) ([]models.Company, error)
	Count(ctx context.Context, params arguments.CompanyCount) (int64, error)
	List(ctx context.Context, params arguments.CompanyList) ([]models.Company, error)
	Insert(ctx context.Context, params arguments.CompanyInsert) (models.Company, error)
	Update(ctx context.Context, params arguments.CompanyUpdate) (models.Company, error)
	Delete(ctx context.Context, param arguments.CompanyDelete) (int64, error)
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
	}).Infof("Resolver GetByID of company")
	claims := utils.ExtractClaimsFromContext(param.Context)
	isPermit := r.checkPermission(claims, fmt.Sprintf("Company%s", utils.GetByID))
	if !isPermit {
		logger.WithFields(logger.Fields{
			"traceId": param.Context.Value(utils.TraceIDKey),
			"Error":   goerrors.ErrNotAuthorized(param.Context.Value(utils.TraceIDKey)),
		}).Errorf("Resolver GetByID !isPermit of company")
		return nil, goerrors.ErrNotAuthorized(param.Context.Value(utils.TraceIDKey))
	}
	// parse param
	args := arguments.CompanyGetByID{}
	if err := utils.Parse(param.Args, &args); err != nil {
		logger.WithFields(logger.Fields{
			"traceId": param.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver GetByID utils.Parse of company")
		return nil, goerrors.ErrInternalServerError(param.Context.Value(utils.TraceIDKey))
	}
	response, err := r.company.GetByID(param.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": param.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver GetByID r.company.GetByID of company")
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
	}).Infof("Resolver Count of company")
	claims := utils.ExtractClaimsFromContext(params.Context)
	isPermit := r.checkPermission(claims, fmt.Sprintf("Company%s", utils.Count))
	if !isPermit {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   goerrors.ErrNotAuthorized(params.Context.Value(utils.TraceIDKey)),
		}).Errorf("Resolver Count !isPermit of company")
		return nil, goerrors.ErrNotAuthorized(params.Context.Value(utils.TraceIDKey))
	}
	// parse params
	args := arguments.CompanyCount{}
	err := utils.Parse(params.Source.(map[string]interface{}), &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver Count utils.Parse of company")
		return nil, goerrors.ErrInternalServerError(params.Context.Value(utils.TraceIDKey))
	}
	response, err := r.company.Count(params.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver Count r.company.Count of company")
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
	}).Infof("Resolver List of company")
	claims := utils.ExtractClaimsFromContext(params.Context)
	isPermit := r.checkPermission(claims, fmt.Sprintf("Company%s", utils.List))
	if !isPermit {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   goerrors.ErrNotAuthorized(params.Context.Value(utils.TraceIDKey)),
		}).Errorf("Resolver List  of company")
		return nil, goerrors.ErrNotAuthorized(params.Context.Value(utils.TraceIDKey))
	}
	// parse params
	args := arguments.CompanyList{}
	err := utils.Parse(params.Source.(map[string]interface{}), &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver List utils.Parse of company")
		return nil, goerrors.ErrInternalServerError(params.Context.Value(utils.TraceIDKey))
	}
	response, err := r.company.List(params.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver List r.company.List of company")
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
	}).Infof("Resolver Insert of company")
	claims := utils.ExtractClaimsFromContext(params.Context)
	isPermit := r.checkPermission(claims, fmt.Sprintf("Company%s", utils.Insert))
	if !isPermit {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   goerrors.ErrNotAuthorized(params.Context.Value(utils.TraceIDKey)),
		}).Errorf("Resolver Insert !isPermit of company")
		return nil, goerrors.ErrNotAuthorized(params.Context.Value(utils.TraceIDKey))
	}
	// parse params
	args := arguments.CompanyInsert{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver Insert utils.Parse of company")
		return nil, goerrors.ErrInternalServerError(params.Context.Value(utils.TraceIDKey))
	}
	response, err := r.company.Insert(params.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver Insert r.company.Insert of company")
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
	}).Infof("Resolver Update of company")
	claims := utils.ExtractClaimsFromContext(params.Context)
	isPermit := r.checkPermission(claims, fmt.Sprintf("Company%s", utils.Update))
	if !isPermit {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   goerrors.ErrNotAuthorized(params.Context.Value(utils.TraceIDKey)),
		}).Errorf("Resolver Update !isPermit of company")
		return nil, goerrors.ErrNotAuthorized(params.Context.Value(utils.TraceIDKey))
	}
	// parse params
	args := arguments.CompanyUpdate{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver Update utils.Parse of company")
		return nil, goerrors.ErrInternalServerError(params.Context.Value(utils.TraceIDKey))
	}
	response, err := r.company.Update(params.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver Update r.company.Update of company")
		if err == sql.ErrNoRows {
			return nil, goerrors.ErrNotFound(params.Context.Value(utils.TraceIDKey))
		}
		return nil, goerrors.ErrInternalServerError(params.Context.Value(utils.TraceIDKey))
	}
	return response, nil
}

//go:generate mockery -name=IHandler -output=mocks -case=underscore
