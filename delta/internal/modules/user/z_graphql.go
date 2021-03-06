// @generated
package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/tanphamhaiduong/go/common/goerrors"
	"github.com/tanphamhaiduong/go/common/logger"
	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/customgraphql"
	"github.com/tanphamhaiduong/go/delta/internal/models"
	"github.com/tanphamhaiduong/go/delta/internal/utils"
)

var (
	// Type ...
	Type = graphql.NewObject(graphql.ObjectConfig{
		Name:        "User",
		Description: "This is type Feature",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "This is user's id",
			},
			"username": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is user's username",
			},
			"password": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is user's password",
			},
			"name": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is user's name",
			},
			"dateOfBirth": &graphql.Field{
				Type:        customgraphql.NullDateTime,
				Description: "This is user's name",
			},
			"reference": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is user's reference",
			},
			"avatarUrl": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is user's avatar url",
			},
			"licenseNumber": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is user's license number",
			},
			"phoneNumber": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is user's name",
			},
			"extension": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is user's extension",
			},
			"telProvider": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is user's tel provider",
			},
			"telApi": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is user's tel api",
			},
			"supervisorId": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "This is user's tel supervisorId",
			},
			"roleId": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "This is user's tel RoleID",
			},
			"companyId": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "This is user's companyId",
			},
			"status": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is user's status",
			},
			"createdBy": &graphql.Field{
				Type:        graphql.String,
				Description: "This is user's createdBy",
			},
			"updatedBy": &graphql.Field{
				Type:        graphql.String,
				Description: "This is user's updatedBy",
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
			Description: "This is user's id",
		},
		"username": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's username",
		},
		"password": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's password",
		},
		"name": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's name",
		},
		"dateOfBirth": &graphql.ArgumentConfig{
			Type:        customgraphql.NullDateTime,
			Description: "This is user's name",
		},
		"reference": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's reference",
		},
		"avatarUrl": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's avatar url",
		},
		"licenseNumber": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's license number",
		},
		"phoneNumber": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's name",
		},
		"extension": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's extension",
		},
		"telProvider": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's tel provider",
		},
		"telApi": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's tel api",
		},
		"supervisorId": &graphql.ArgumentConfig{
			Type:        graphql.Int,
			Description: "This is user's tel supervisorId",
		},
		"roleId": &graphql.ArgumentConfig{
			Type:        graphql.Int,
			Description: "This is user's tel RoleID",
		},
		"companyId": &graphql.ArgumentConfig{
			Type:        graphql.Int,
			Description: "This is user's companyId",
		},
		"status": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's status",
		},
		"createdBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's createdBy",
		},
		"updatedBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's updatedBy",
		},
		"beginId": &graphql.ArgumentConfig{
			Type:         graphql.Int,
			Description:  "This is ID greater than",
			DefaultValue: 0,
		},
		"endId": &graphql.ArgumentConfig{
			Type:         graphql.Int,
			Description:  "This is ID less than",
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
		"username": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is user's username",
		},
		"password": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is user's password",
		},
		"name": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is user's name",
		},
		"dateOfBirth": &graphql.ArgumentConfig{
			Type:        customgraphql.NullDateTime,
			Description: "This is user's name",
		},
		"reference": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is user's reference",
		},
		"avatarUrl": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is user's avatar url",
		},
		"licenseNumber": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is user's license number",
		},
		"phoneNumber": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is user's name",
		},
		"extension": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is user's extension",
		},
		"telProvider": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is user's tel provider",
		},
		"telApi": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is user's tel api",
		},
		"supervisorId": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "This is user's tel supervisorId",
		},
		"roleId": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "This is user's tel RoleID",
		},
		"companyId": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "This is user's companyId",
		},
		"status": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is user's status",
		},
		"createdBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's createdBy",
		},
		"updatedBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's updatedBy",
		},
	}
	// UpdateTypeArgs ...
	UpdateTypeArgs = graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"username": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's username",
		},
		"password": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's password",
		},
		"name": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's name",
		},
		"dateOfBirth": &graphql.ArgumentConfig{
			Type:        customgraphql.NullDateTime,
			Description: "This is user's name",
		},
		"reference": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's reference",
		},
		"avatarUrl": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's avatar url",
		},
		"licenseNumber": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's license number",
		},
		"phoneNumber": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's name",
		},
		"extension": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's extension",
		},
		"telProvider": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's tel provider",
		},
		"telApi": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's tel api",
		},
		"supervisorId": &graphql.ArgumentConfig{
			Type:        graphql.Int,
			Description: "This is user's tel supervisorId",
		},
		"roleId": &graphql.ArgumentConfig{
			Type:        graphql.Int,
			Description: "This is user's tel RoleID",
		},
		"companyId": &graphql.ArgumentConfig{
			Type:        graphql.Int,
			Description: "This is user's companyId",
		},
		"status": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's status",
		},
		"createdBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's createdBy",
		},
		"updatedBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's updatedBy",
		},
	}
)

// ICoreHandler ...
type ICoreHandler interface {
	GetByID(ctx context.Context, param arguments.UserGetByID) (models.User, error)
	GetByIDs(ctx context.Context, param arguments.UserGetByIDs) ([]models.User, error)
	Count(ctx context.Context, params arguments.UserCount) (int64, error)
	List(ctx context.Context, params arguments.UserList) ([]models.User, error)
	Insert(ctx context.Context, params arguments.UserInsert) (models.User, error)
	Update(ctx context.Context, params arguments.UserUpdate) (models.User, error)
	Delete(ctx context.Context, param arguments.UserDelete) (int64, error)
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
	}).Infof("Resolver GetByID of user")
	claims := utils.ExtractClaimsFromContext(param.Context)
	isPermit := r.checkPermission(claims, fmt.Sprintf("User%s", utils.GetByID))
	if !isPermit {
		logger.WithFields(logger.Fields{
			"traceId": param.Context.Value(utils.TraceIDKey),
			"Error":   goerrors.ErrNotAuthorized(param.Context.Value(utils.TraceIDKey)),
		}).Errorf("Resolver GetByID !isPermit of user")
		return nil, goerrors.ErrNotAuthorized(param.Context.Value(utils.TraceIDKey))
	}
	// parse param
	args := arguments.UserGetByID{}
	if err := utils.Parse(param.Args, &args); err != nil {
		logger.WithFields(logger.Fields{
			"traceId": param.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver GetByID utils.Parse of user")
		return nil, goerrors.ErrInternalServerError(param.Context.Value(utils.TraceIDKey))
	}
	response, err := r.user.GetByID(param.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": param.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver GetByID r.user.GetByID of user")
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
	}).Infof("Resolver Count of user")
	claims := utils.ExtractClaimsFromContext(params.Context)
	isPermit := r.checkPermission(claims, fmt.Sprintf("User%s", utils.Count))
	if !isPermit {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   goerrors.ErrNotAuthorized(params.Context.Value(utils.TraceIDKey)),
		}).Errorf("Resolver Count !isPermit of user")
		return nil, goerrors.ErrNotAuthorized(params.Context.Value(utils.TraceIDKey))
	}
	// parse params
	args := arguments.UserCount{}
	err := utils.Parse(params.Source.(map[string]interface{}), &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver Count utils.Parse of user")
		return nil, goerrors.ErrInternalServerError(params.Context.Value(utils.TraceIDKey))
	}
	response, err := r.user.Count(params.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver Count r.user.Count of user")
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
	}).Infof("Resolver List of user")
	claims := utils.ExtractClaimsFromContext(params.Context)
	isPermit := r.checkPermission(claims, fmt.Sprintf("User%s", utils.List))
	if !isPermit {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   goerrors.ErrNotAuthorized(params.Context.Value(utils.TraceIDKey)),
		}).Errorf("Resolver List  of user")
		return nil, goerrors.ErrNotAuthorized(params.Context.Value(utils.TraceIDKey))
	}
	// parse params
	args := arguments.UserList{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver List utils.Parse of user")
		return nil, goerrors.ErrInternalServerError(params.Context.Value(utils.TraceIDKey))
	}
	response, err := r.user.List(params.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver List r.user.List of user")
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
	}).Infof("Resolver Insert of user")
	claims := utils.ExtractClaimsFromContext(params.Context)
	isPermit := r.checkPermission(claims, fmt.Sprintf("User%s", utils.Insert))
	if !isPermit {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   goerrors.ErrNotAuthorized(params.Context.Value(utils.TraceIDKey)),
		}).Errorf("Resolver Insert !isPermit of user")
		return nil, goerrors.ErrNotAuthorized(params.Context.Value(utils.TraceIDKey))
	}
	// parse params
	args := arguments.UserInsert{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver Insert utils.Parse of user")
		return nil, goerrors.ErrInternalServerError(params.Context.Value(utils.TraceIDKey))
	}
	response, err := r.user.Insert(params.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver Insert r.user.Insert of user")
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
	}).Infof("Resolver Update of user")
	claims := utils.ExtractClaimsFromContext(params.Context)
	isPermit := r.checkPermission(claims, fmt.Sprintf("User%s", utils.Update))
	if !isPermit {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   goerrors.ErrNotAuthorized(params.Context.Value(utils.TraceIDKey)),
		}).Errorf("Resolver Update !isPermit of user")
		return nil, goerrors.ErrNotAuthorized(params.Context.Value(utils.TraceIDKey))
	}
	// parse params
	args := arguments.UserUpdate{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver Update utils.Parse of user")
		return nil, goerrors.ErrInternalServerError(params.Context.Value(utils.TraceIDKey))
	}
	response, err := r.user.Update(params.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": params.Context.Value(utils.TraceIDKey),
			"Error":   err,
		}).Errorf("Resolver Update r.user.Update of user")
		if err == sql.ErrNoRows {
			return nil, goerrors.ErrNotFound(params.Context.Value(utils.TraceIDKey))
		}
		return nil, goerrors.ErrInternalServerError(params.Context.Value(utils.TraceIDKey))
	}
	return response, nil
}

//go:generate mockery -name=IHandler -output=mocks -case=underscore
