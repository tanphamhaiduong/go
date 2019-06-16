// @generated
package company

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

	// DeleteTypeArgs ...
	DeleteTypeArgs = graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	}
)

// ICoreHandler ...
type ICoreHandler interface {
	GetByID(ctx context.Context, params arguments.CompanyGetByIDArgs) (models.Company, error)
	GetByIDs(ctx context.Context, params arguments.CompanyGetByIDsArgs) ([]models.Company, error)
	Count(ctx context.Context, params arguments.CompanyCountArgs) (int64, error)
	List(ctx context.Context, params arguments.CompanyListArgs) ([]models.Company, error)
	Insert(ctx context.Context, params arguments.CompanyInsertArgs) (models.Company, error)
	Update(ctx context.Context, params arguments.CompanyUpdateArgs) (models.Company, error)
	Delete(ctx context.Context, params arguments.CompanyDeleteArgs) (int64, error)
}

// ForwardParams ...
func (r *ResolverImpl) ForwardParams(params graphql.ResolveParams) (interface{}, error) {
	log.WithField("params", params).Info("company of Resolver ForwardParams")
	return params.Args, nil
}

// GetByID ...
func (r *ResolverImpl) GetByID(params graphql.ResolveParams) (interface{}, error) {
	log.WithField("params", params).Info("company of Resolver GetByID")
	// parse params
	args := arguments.CompanyGetByIDArgs{}
	if err := utils.Parse(params.Args, &args); err != nil {
		return nil, err
	}
	response, err := r.company.GetByID(params.Context, args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Count ...
func (r *ResolverImpl) Count(params graphql.ResolveParams) (interface{}, error) {
	log.WithField("params", params).Info("company of Resolver Count")
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
	log.WithField("params", params).Info("company of Resolver List")
	// parse params
	args := arguments.CompanyListArgs{}
	err := utils.Parse(params.Source.(map[string]interface{}), &args)
	if err != nil {
		log.WithField("Error", err).Error("company of Resolver List utils.Parse")
		return nil, err
	}
	response, err := r.company.List(params.Context, args)
	if err != nil {
		log.WithField("Error", err).Error("company of Resolver List r.company.List")
		return nil, err
	}
	return response, nil
}

// Insert ...
func (r *ResolverImpl) Insert(params graphql.ResolveParams) (interface{}, error) {
	log.WithField("params", params).Info("company of Resolver Insert")
	// parse params
	args := arguments.CompanyInsertArgs{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		log.WithField("Error", err).Error("company of Resolver Insert utils.Parse")
		return nil, err
	}
	response, err := r.company.Insert(params.Context, args)
	if err != nil {
		log.WithField("Error", err).Error("company of Resolver Insert r.company.Insert")
		return nil, err
	}
	return response, nil
}

// Update ...
func (r *ResolverImpl) Update(params graphql.ResolveParams) (interface{}, error) {
	log.WithField("params", params).Info("company of Resolver Update")
	// parse params
	args := arguments.CompanyUpdateArgs{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		log.WithField("Error", err).Error("company of Resolver Update utils.Parse")
		return nil, err
	}
	response, err := r.company.Update(params.Context, args)
	if err != nil {
		log.WithField("Error", err).Error("company of Resolver Update r.company.Update")
		return nil, err
	}
	return response, nil
}

// Delete ...
func (r *ResolverImpl) Delete(params graphql.ResolveParams) (interface{}, error) {
	log.WithField("params", params).Info("company of Resolver Delete")
	// parse params
	args := arguments.CompanyDeleteArgs{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		log.WithField("Error", err).Error("company of Resolver Delete utils.Parse")
		return nil, err
	}
	response, err := r.company.Delete(params.Context, args)
	if err != nil {
		log.WithField("Error", err).Error("company of Resolver Delete r.company.Delete")
		return nil, err
	}
	return response, nil
}

//go:generate mockery -name=IHandler -output=mocks -case=underscore
