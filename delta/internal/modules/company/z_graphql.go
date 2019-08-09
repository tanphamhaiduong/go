// @generated
package company

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
			"companyCode": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is company's code",
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
			"apiSecretKey": &graphql.Field{
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
		"companyCode": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is company's code",
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
		"apiSecretKey": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is company's updatedBy",
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
			Description: "This is company's name",
		},
		"companyCode": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is company's code",
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
		"apiSecretKey": &graphql.ArgumentConfig{
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
		"companyCode": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is company's code",
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
		"apiSecretKey": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is company's updatedBy",
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
	GetByID(ctx context.Context, param arguments.CompanyGetByIDArgs) (models.Company, error)
	GetByIDs(ctx context.Context, param arguments.CompanyGetByIDsArgs) ([]models.Company, error)
	Count(ctx context.Context, params arguments.CompanyCountArgs) (int64, error)
	List(ctx context.Context, params arguments.CompanyListArgs) ([]models.Company, error)
	Insert(ctx context.Context, params arguments.CompanyInsertArgs) (models.Company, error)
	Update(ctx context.Context, params arguments.CompanyUpdateArgs) (models.Company, error)
	Delete(ctx context.Context, param arguments.CompanyDeleteArgs) (int64, error)
}

// ForwardParams ...
func (r *ResolverImpl) ForwardParams(params graphql.ResolveParams) (interface{}, error) {
	logger.WithFields(logger.Fields{
		"TraceID": params.Context.Value("TraceID"),
		"params":  params,
	}).Infof("Resolver ForwardParams of company")
	return params.Args, nil
}

// GetByID ...
func (r *ResolverImpl) GetByID(param graphql.ResolveParams) (interface{}, error) {
	logger.WithFields(logger.Fields{
		"TraceID": param.Context.Value("TraceID"),
		"param":   param,
	}).Infof("Resolver GetByID of company")
	// parse param
	args := arguments.CompanyGetByIDArgs{}
	if err := utils.Parse(param.Args, &args); err != nil {
		return nil, err
	}
	response, err := r.company.GetByID(param.Context, args)
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
	}).Infof("Resolver Count of company")
	// parse params
	args := arguments.CompanyCountArgs{}
	err := utils.Parse(params.Source.(map[string]interface{}), &args)
	if err != nil {
		return nil, err
	}
	response, err := r.company.Count(params.Context, args)
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
	}).Infof("Resolver List of company")
	// parse params
	args := arguments.CompanyListArgs{}
	err := utils.Parse(params.Source.(map[string]interface{}), &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": params.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver List utils.Parse company")
		return nil, err
	}
	response, err := r.company.List(params.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": params.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver List r.company.List company")
		return nil, err
	}
	return response, nil
}

// Insert ...
func (r *ResolverImpl) Insert(params graphql.ResolveParams) (interface{}, error) {
	logger.WithFields(logger.Fields{
		"TraceID": params.Context.Value("TraceID"),
		"params":  params,
	}).Infof("Resolver Insert of company")
	// parse params
	args := arguments.CompanyInsertArgs{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": params.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver Insert utils.Parse company")
		return nil, err
	}
	response, err := r.company.Insert(params.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": params.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver Insert r.company.Insert company")
		return nil, err
	}
	return response, nil
}

// Update ...
func (r *ResolverImpl) Update(params graphql.ResolveParams) (interface{}, error) {
	logger.WithFields(logger.Fields{
		"TraceID": params.Context.Value("TraceID"),
		"params":  params,
	}).Infof("Resolver Update of company")
	// parse params
	args := arguments.CompanyUpdateArgs{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": params.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver Update utils.Parse company")
		return nil, err
	}
	response, err := r.company.Update(params.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": params.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver Update r.company.Update company")
		return nil, err
	}
	return response, nil
}

// Delete ...
func (r *ResolverImpl) Delete(param graphql.ResolveParams) (interface{}, error) {
	logger.WithFields(logger.Fields{
		"TraceID": param.Context.Value("TraceID"),
		"param":   param,
	}).Infof("Resolver Delete of company")
	// parse param
	args := arguments.CompanyDeleteArgs{}
	err := utils.Parse(param.Args, &args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": param.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver Delete utils.Parse company")
		return nil, err
	}
	response, err := r.company.Delete(param.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": param.Context.Value("TraceID"),
			"Error":   err,
		}).Errorf("Resolver Delete r.company.Delete company")
		return nil, err
	}
	return response, nil
}

//go:generate mockery -name=IHandler -output=mocks -case=underscore
